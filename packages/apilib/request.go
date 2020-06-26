package apilib

import (
	"errors"
	"fmt"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	valuetransaction "github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/transaction"
	nodeapi "github.com/iotaledger/goshimmer/packages/waspconn/apilib"
	"github.com/iotaledger/wasp/packages/sctransaction"
	"github.com/iotaledger/wasp/packages/sctransaction/txbuilder"
	"github.com/iotaledger/wasp/packages/util"
	"strconv"
)

type RequestBlockJson struct {
	Address     string            `json:"address"`
	RequestCode uint16            `json:"request_code"`
	Amount      int64             `json:"amount"` // minimum 1i
	Vars        map[string]string `json:"vars"`
}

func CreateRequestTransaction(node string, senderSigScheme signaturescheme.SignatureScheme, reqsJson []*RequestBlockJson) (*sctransaction.Transaction, error) {
	var err error
	if len(reqsJson) == 0 {
		return nil, errors.New("CreateRequestTransaction: must be at least 1 request block")
	}
	senderAddr := senderSigScheme.Address()
	allOuts, err := nodeapi.GetAccountOutputs(node, &senderAddr)
	if err != nil {
		return nil, fmt.Errorf("can't get outputs from the node: %v", err)
	}
	fmt.Printf("[[[[]]]] CreateRequestTransaction:\nall inputs: %s sender addr: %s\nnode: %s\n",
		util.InputBalancesToString(allOuts), senderAddr.String(), node)

	txb, err := txbuilder.NewFromOutputBalances(allOuts)
	if err != nil {
		return nil, err
	}

	for _, reqBlkJson := range reqsJson {
		reqBlk, err := requestBlockFromJson(reqBlkJson)
		if err != nil {
			return nil, err
		}
		// add request block
		if err = txb.AddRequestBlock(reqBlk); err != nil {
			return nil, err
		}
	}
	tx, err := txb.Build(false)
	if err != nil {
		return nil, err
	}
	tx.Sign(senderSigScheme)

	MustNotNullInputs(tx.Transaction)
	fmt.Printf("[[[[]]]] create request tx:\n%s\n", tx.String())

	return tx, nil
}

func requestBlockFromJson(reqBlkJson *RequestBlockJson) (*sctransaction.RequestBlock, error) {
	var err error
	addr, err := address.FromBase58(reqBlkJson.Address)
	if err != nil {
		return nil, err
	}
	ret := sctransaction.NewRequestBlock(addr, sctransaction.RequestCode(reqBlkJson.RequestCode))

	for k, v := range reqBlkJson.Vars {
		n, err := strconv.Atoi(v)
		if err != nil {
			ret.Args().SetString(k, v)
		} else {
			ret.Args().SetInt64(k, int64(n))
		}
	}
	return ret, nil
}

func MustNotNullInputs(tx *valuetransaction.Transaction) {
	var nilid valuetransaction.ID
	tx.Inputs().ForEach(func(outputId valuetransaction.OutputID) bool {
		if outputId.TransactionID() == nilid {
			panic(fmt.Sprintf("nil input in txid %s", tx.ID().String()))
		}
		return true
	})
}
