package cluster

import (
	"bytes"
	"fmt"
	"time"

	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/vm/builtinvm/blob"
	"github.com/iotaledger/wasp/plugins/wasmtimevm"

	"github.com/iotaledger/goshimmer/client/wallet/packages/seed"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/client"
	"github.com/iotaledger/wasp/client/chainclient"
	"github.com/iotaledger/wasp/client/multiclient"
	"github.com/iotaledger/wasp/client/scclient"
	"github.com/iotaledger/wasp/packages/coret"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/sctransaction"
	"github.com/iotaledger/wasp/packages/vm/builtinvm/root"
)

type Chain struct {
	Description string

	OriginatorSeed *seed.Seed

	CommitteeNodes []int
	AccessNodes    []int
	Quorum         uint16
	Address        address.Address

	ChainID coret.ChainID
	Color   balance.Color

	Cluster *Cluster
}

func (ch *Chain) ChainAddress() *address.Address {
	r := address.Address(ch.ChainID)
	return &r
}

func (ch *Chain) ContractID(contractHname coret.Hname) coret.ContractID {
	return coret.NewContractID(ch.ChainID, contractHname)
}

func (ch *Chain) AllNodes() []int {
	r := make([]int, 0)
	r = append(r, ch.CommitteeNodes...)
	r = append(r, ch.AccessNodes...)
	return r
}

func (ch *Chain) CommitteeApi() []string {
	return ch.Cluster.WaspHosts(ch.CommitteeNodes, (*WaspNodeConfig).ApiHost)
}

func (ch *Chain) OriginatorAddress() *address.Address {
	addr := ch.OriginatorSeed.Address(0).Address
	return &addr
}

func (ch *Chain) OriginatorID() *coret.AgentID {
	ret := coret.NewAgentIDFromAddress(*ch.OriginatorAddress())
	return &ret
}

func (ch *Chain) OriginatorSigScheme() signaturescheme.SignatureScheme {
	return signaturescheme.ED25519(*ch.OriginatorSeed.KeyPair(0))
}

func (ch *Chain) OriginatorClient() *chainclient.Client {
	return ch.Client(ch.OriginatorSigScheme())
}

func (ch *Chain) Client(sigScheme signaturescheme.SignatureScheme) *chainclient.Client {
	return chainclient.New(
		ch.Cluster.NodeClient,
		ch.Cluster.WaspClient(ch.CommitteeNodes[0]),
		ch.ChainID,
		sigScheme,
	)
}

func (ch *Chain) SCClient(contractHname coret.Hname, sigScheme signaturescheme.SignatureScheme) *scclient.SCClient {
	return scclient.New(ch.Client(sigScheme), contractHname)
}

func (ch *Chain) CommitteeMultiClient() *multiclient.MultiClient {
	return multiclient.New(ch.CommitteeApi())
}

func (ch *Chain) WithSCState(hname coret.Hname, f func(host string, blockIndex uint32, state codec.ImmutableMustCodec) bool) bool {
	pass := true
	for i, host := range ch.CommitteeApi() {
		if !ch.Cluster.Config.Nodes[i].IsUp() {
			continue
		}
		contractID := coret.NewContractID(ch.ChainID, hname)
		actual, err := ch.Cluster.WaspClient(i).DumpSCState(&contractID)
		if client.IsNotFound(err) {
			pass = false
			fmt.Printf("   FAIL: state does not exist\n")
			continue
		}
		if err != nil {
			panic(err)
		}
		if !f(host, actual.Index, codec.NewMustCodec(actual.Variables)) {
			pass = false
		}
	}
	return pass
}

func (ch *Chain) DeployContract(name string, progHashStr string, description string, initParams map[string]interface{}) (*sctransaction.Transaction, error) {
	programHash, err := hashing.HashValueFromBase58(progHashStr)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		root.ParamName:        name,
		root.ParamProgramHash: programHash,
		root.ParamDescription: description,
	}
	for k, v := range initParams {
		params[k] = v
	}
	tx, err := ch.OriginatorClient().PostRequest(
		root.Interface.Hname(),
		coret.Hn(root.FuncDeployContract),
		chainclient.PostRequestParams{
			Args: codec.EncodeDictFromMap(params),
		},
	)
	if err != nil {
		return nil, err
	}

	err = ch.CommitteeMultiClient().WaitUntilAllRequestsProcessed(tx, 30*time.Second)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (ch *Chain) DeployWasmContract(name string, description string, progBinary []byte, initParams map[string]interface{}) (*sctransaction.Transaction, hashing.HashValue, error) {
	// upload binary to the chain
	blobFieldValues := map[string]interface{}{
		blob.VarFieldVMType:             wasmtimevm.VMType,
		blob.VarFieldProgramBinary:      progBinary,
		blob.VarFieldProgramDescription: description,
	}
	programHash := blob.MustGetBlobHash(codec.NewCodec(codec.EncodeDictFromMap(blobFieldValues)))

	reqTx, err := ch.OriginatorClient().PostRequest(
		blob.Interface.Hname(),
		coret.Hn(blob.FuncStoreBlob),
		chainclient.PostRequestParams{
			Args: codec.EncodeDictFromMap(blobFieldValues),
		},
	)
	err = ch.CommitteeMultiClient().WaitUntilAllRequestsProcessed(reqTx, 30*time.Second)
	if err != nil {
		return nil, *hashing.NilHash, err
	}
	progBinaryBack, err := ch.GetBlobFieldValue(programHash, blob.VarFieldProgramBinary)
	if err != nil {
		return nil, *hashing.NilHash, err
	}
	if !bytes.Equal(progBinary, progBinaryBack) {
		return nil, *hashing.NilHash, fmt.Errorf("!bytes.Equal(progBinary, progBinaryBack)")
	}
	fmt.Printf("---- blob installed correctly len = %d\n", len(progBinaryBack))

	params := make(map[string]interface{})
	for k, v := range initParams {
		params[k] = v
	}
	params[root.ParamName] = name
	params[root.ParamProgramHash] = programHash
	params[root.ParamDescription] = description

	tx, err := ch.OriginatorClient().PostRequest(
		root.Interface.Hname(),
		coret.Hn(root.FuncDeployContract),
		chainclient.PostRequestParams{
			Args: codec.EncodeDictFromMap(params),
		},
	)
	if err != nil {
		return nil, *hashing.NilHash, err
	}

	err = ch.CommitteeMultiClient().WaitUntilAllRequestsProcessed(tx, 30*time.Second)
	if err != nil {
		return nil, *hashing.NilHash, err
	}

	return tx, programHash, nil
}

func (ch *Chain) GetBlobFieldValue(blobHash hashing.HashValue, field string) ([]byte, error) {
	v, err := ch.Cluster.WaspClient(0).CallView(
		ch.ContractID(blob.Interface.Hname()),
		blob.FuncGetBlobField,
		dict.FromGoMap(map[kv.Key][]byte{
			blob.ParamHash:  blobHash[:],
			blob.ParamField: []byte(field),
		}),
	)
	if err != nil {
		return nil, err
	}
	if v.IsEmpty() {
		return nil, nil
	}
	ret, err := v.Get(blob.ParamBytes)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
