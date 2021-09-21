package tests

import (
	"testing"
	"time"

	"github.com/iotaledger/wasp/client/chainclient"
	"github.com/iotaledger/wasp/contracts/native/inccounter"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/iscp/colored"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/stretchr/testify/require"
)

const numRequests = 100000

func TestSpamOnledger(t *testing.T) {
	testutil.SkipHeavy(t)
	env := setupAdvancedInccounterTest(t, 1, []int{0})

	keyPair, _ := env.getOrCreateAddress()
	myClient := env.chain.SCClient(iscp.Hn(incCounterSCName), keyPair)

	for i := 0; i < numRequests; i++ {
		args := chainclient.NewPostRequestParams().WithIotas(1)
		_, err := myClient.PostRequest(inccounter.FuncIncCounter.Name, *args)
		require.NoError(t, err)
	}
	// TODO check blocklog
}

func TestSpamOffledger(t *testing.T) {
	testutil.SkipHeavy(t)
	// single wasp node committee, to test if publishing can break state transitions
	env := setupAdvancedInccounterTest(t, 1, []int{0})

	// deposit funds for offledger requests
	keyPair, myAddress := env.getOrCreateAddress()
	myAgentID := iscp.NewAgentID(myAddress, 0)

	accountsClient := env.chain.SCClient(accounts.Contract.Hname(), keyPair)
	_, err := accountsClient.PostRequest(accounts.FuncDeposit.Name, chainclient.PostRequestParams{
		Transfer: colored.NewBalancesForIotas(1000000),
	})
	require.NoError(t, err)

	waitUntil(t, env.balanceOnChainIotaEquals(myAgentID, 1000000), util.MakeRange(0, 1), 60*time.Second, "send 100i")

	myClient := env.chain.SCClient(iscp.Hn(incCounterSCName), keyPair)

	for i := 0; i < numRequests; i++ {
		_, err = myClient.PostOffLedgerRequest(inccounter.FuncIncCounter.Name, chainclient.PostRequestParams{Nonce: uint64(i + 1)})
		require.NoError(t, err)
	}
	// TODO check blocklog
}
