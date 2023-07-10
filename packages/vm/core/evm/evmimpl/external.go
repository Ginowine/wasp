package evmimpl

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/subrealm"
	"github.com/iotaledger/wasp/packages/vm/core/evm/emulator"
)

func Nonce(state kv.KVStoreReader, addr common.Address) uint64 {
	emulatorState := evmStateSubrealmR(state)
	stateDBStore := subrealm.NewReadOnly(emulatorState, emulator.KeyStateDB)
	return emulator.GetNonce(stateDBStore, addr)
}

func CheckNonce(state kv.KVStore, addr common.Address, nonce uint64) error {
	expected := Nonce(state, addr)
	if nonce != expected {
		return fmt.Errorf("Invalid nonce, expected %d, got %d", expected, nonce)
	}
	return nil
}