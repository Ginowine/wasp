// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package coreaccounts

import (
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
)

type ImmutableFoundryCreateNewResults struct {
	Proxy wasmtypes.Proxy
}

// serial number of the newly created foundry
func (s ImmutableFoundryCreateNewResults) FoundrySN() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.Proxy.Root(ResultFoundrySN))
}

type MutableFoundryCreateNewResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableFoundryCreateNewResults() MutableFoundryCreateNewResults {
	return MutableFoundryCreateNewResults{Proxy: wasmlib.NewResultsProxy()}
}

// serial number of the newly created foundry
func (s MutableFoundryCreateNewResults) FoundrySN() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.Proxy.Root(ResultFoundrySN))
}

type MapUint32ToImmutableBool struct {
	Proxy wasmtypes.Proxy
}

func (m MapUint32ToImmutableBool) GetBool(key uint32) wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(m.Proxy.Key(wasmtypes.Uint32ToBytes(key)))
}

type ImmutableAccountFoundriesResults struct {
	Proxy wasmtypes.Proxy
}

// foundry serial numbers owned by the given account
func (s ImmutableAccountFoundriesResults) Foundries() MapUint32ToImmutableBool {
	return MapUint32ToImmutableBool(s)
}

type MapUint32ToMutableBool struct {
	Proxy wasmtypes.Proxy
}

func (m MapUint32ToMutableBool) Clear() {
	m.Proxy.ClearMap()
}

func (m MapUint32ToMutableBool) GetBool(key uint32) wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(m.Proxy.Key(wasmtypes.Uint32ToBytes(key)))
}

type MutableAccountFoundriesResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableAccountFoundriesResults() MutableAccountFoundriesResults {
	return MutableAccountFoundriesResults{Proxy: wasmlib.NewResultsProxy()}
}

// foundry serial numbers owned by the given account
func (s MutableAccountFoundriesResults) Foundries() MapUint32ToMutableBool {
	return MapUint32ToMutableBool(s)
}

type ImmutableAccountNFTAmountResults struct {
	Proxy wasmtypes.Proxy
}

// amount of NFTs owned by the account
func (s ImmutableAccountNFTAmountResults) Amount() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.Proxy.Root(ResultAmount))
}

type MutableAccountNFTAmountResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableAccountNFTAmountResults() MutableAccountNFTAmountResults {
	return MutableAccountNFTAmountResults{Proxy: wasmlib.NewResultsProxy()}
}

// amount of NFTs owned by the account
func (s MutableAccountNFTAmountResults) Amount() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.Proxy.Root(ResultAmount))
}

type ImmutableAccountNFTAmountInCollectionResults struct {
	Proxy wasmtypes.Proxy
}

// amount of NFTs in collection owned by the account
func (s ImmutableAccountNFTAmountInCollectionResults) Amount() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.Proxy.Root(ResultAmount))
}

type MutableAccountNFTAmountInCollectionResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableAccountNFTAmountInCollectionResults() MutableAccountNFTAmountInCollectionResults {
	return MutableAccountNFTAmountInCollectionResults{Proxy: wasmlib.NewResultsProxy()}
}

// amount of NFTs in collection owned by the account
func (s MutableAccountNFTAmountInCollectionResults) Amount() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.Proxy.Root(ResultAmount))
}

type ArrayOfImmutableNftID struct {
	Proxy wasmtypes.Proxy
}

func (a ArrayOfImmutableNftID) Length() uint32 {
	return a.Proxy.Length()
}

func (a ArrayOfImmutableNftID) GetNftID(index uint32) wasmtypes.ScImmutableNftID {
	return wasmtypes.NewScImmutableNftID(a.Proxy.Index(index))
}

type ImmutableAccountNFTsResults struct {
	Proxy wasmtypes.Proxy
}

// NFT IDs owned by the account
func (s ImmutableAccountNFTsResults) NftIDs() ArrayOfImmutableNftID {
	return ArrayOfImmutableNftID{Proxy: s.Proxy.Root(ResultNftIDs)}
}

type ArrayOfMutableNftID struct {
	Proxy wasmtypes.Proxy
}

func (a ArrayOfMutableNftID) AppendNftID() wasmtypes.ScMutableNftID {
	return wasmtypes.NewScMutableNftID(a.Proxy.Append())
}

func (a ArrayOfMutableNftID) Clear() {
	a.Proxy.ClearArray()
}

func (a ArrayOfMutableNftID) Length() uint32 {
	return a.Proxy.Length()
}

func (a ArrayOfMutableNftID) GetNftID(index uint32) wasmtypes.ScMutableNftID {
	return wasmtypes.NewScMutableNftID(a.Proxy.Index(index))
}

type MutableAccountNFTsResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableAccountNFTsResults() MutableAccountNFTsResults {
	return MutableAccountNFTsResults{Proxy: wasmlib.NewResultsProxy()}
}

// NFT IDs owned by the account
func (s MutableAccountNFTsResults) NftIDs() ArrayOfMutableNftID {
	return ArrayOfMutableNftID{Proxy: s.Proxy.Root(ResultNftIDs)}
}

type ImmutableAccountNFTsInCollectionResults struct {
	Proxy wasmtypes.Proxy
}

// NFT IDs in collection owned by the account
func (s ImmutableAccountNFTsInCollectionResults) NftIDs() ArrayOfImmutableNftID {
	return ArrayOfImmutableNftID{Proxy: s.Proxy.Root(ResultNftIDs)}
}

type MutableAccountNFTsInCollectionResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableAccountNFTsInCollectionResults() MutableAccountNFTsInCollectionResults {
	return MutableAccountNFTsInCollectionResults{Proxy: wasmlib.NewResultsProxy()}
}

// NFT IDs in collection owned by the account
func (s MutableAccountNFTsInCollectionResults) NftIDs() ArrayOfMutableNftID {
	return ArrayOfMutableNftID{Proxy: s.Proxy.Root(ResultNftIDs)}
}

type MapAgentIDToImmutableBool struct {
	Proxy wasmtypes.Proxy
}

func (m MapAgentIDToImmutableBool) GetBool(key wasmtypes.ScAgentID) wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(m.Proxy.Key(wasmtypes.AgentIDToBytes(key)))
}

type ImmutableAccountsResults struct {
	Proxy wasmtypes.Proxy
}

// agent IDs
func (s ImmutableAccountsResults) AllAccounts() MapAgentIDToImmutableBool {
	return MapAgentIDToImmutableBool(s)
}

type MapAgentIDToMutableBool struct {
	Proxy wasmtypes.Proxy
}

func (m MapAgentIDToMutableBool) Clear() {
	m.Proxy.ClearMap()
}

func (m MapAgentIDToMutableBool) GetBool(key wasmtypes.ScAgentID) wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(m.Proxy.Key(wasmtypes.AgentIDToBytes(key)))
}

type MutableAccountsResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableAccountsResults() MutableAccountsResults {
	return MutableAccountsResults{Proxy: wasmlib.NewResultsProxy()}
}

// agent IDs
func (s MutableAccountsResults) AllAccounts() MapAgentIDToMutableBool {
	return MapAgentIDToMutableBool(s)
}

type MapTokenIDToImmutableBigInt struct {
	Proxy wasmtypes.Proxy
}

func (m MapTokenIDToImmutableBigInt) GetBigInt(key wasmtypes.ScTokenID) wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(m.Proxy.Key(wasmtypes.TokenIDToBytes(key)))
}

type ImmutableBalanceResults struct {
	Proxy wasmtypes.Proxy
}

// balance per token ID, zero length indicates base token
func (s ImmutableBalanceResults) Balances() MapTokenIDToImmutableBigInt {
	return MapTokenIDToImmutableBigInt(s)
}

type MapTokenIDToMutableBigInt struct {
	Proxy wasmtypes.Proxy
}

func (m MapTokenIDToMutableBigInt) Clear() {
	m.Proxy.ClearMap()
}

func (m MapTokenIDToMutableBigInt) GetBigInt(key wasmtypes.ScTokenID) wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(m.Proxy.Key(wasmtypes.TokenIDToBytes(key)))
}

type MutableBalanceResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableBalanceResults() MutableBalanceResults {
	return MutableBalanceResults{Proxy: wasmlib.NewResultsProxy()}
}

// balance per token ID, zero length indicates base token
func (s MutableBalanceResults) Balances() MapTokenIDToMutableBigInt {
	return MapTokenIDToMutableBigInt(s)
}

type ImmutableBalanceBaseTokenResults struct {
	Proxy wasmtypes.Proxy
}

// amount of base tokens in the account
func (s ImmutableBalanceBaseTokenResults) Balance() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.Proxy.Root(ResultBalance))
}

type MutableBalanceBaseTokenResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableBalanceBaseTokenResults() MutableBalanceBaseTokenResults {
	return MutableBalanceBaseTokenResults{Proxy: wasmlib.NewResultsProxy()}
}

// amount of base tokens in the account
func (s MutableBalanceBaseTokenResults) Balance() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.Proxy.Root(ResultBalance))
}

type ImmutableBalanceNativeTokenResults struct {
	Proxy wasmtypes.Proxy
}

// amount of native tokens in the account
func (s ImmutableBalanceNativeTokenResults) Tokens() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.Proxy.Root(ResultTokens))
}

type MutableBalanceNativeTokenResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableBalanceNativeTokenResults() MutableBalanceNativeTokenResults {
	return MutableBalanceNativeTokenResults{Proxy: wasmlib.NewResultsProxy()}
}

// amount of native tokens in the account
func (s MutableBalanceNativeTokenResults) Tokens() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.Proxy.Root(ResultTokens))
}

type ImmutableGetAccountNonceResults struct {
	Proxy wasmtypes.Proxy
}

// account nonce
func (s ImmutableGetAccountNonceResults) AccountNonce() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.Proxy.Root(ResultAccountNonce))
}

type MutableGetAccountNonceResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetAccountNonceResults() MutableGetAccountNonceResults {
	return MutableGetAccountNonceResults{Proxy: wasmlib.NewResultsProxy()}
}

// account nonce
func (s MutableGetAccountNonceResults) AccountNonce() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.Proxy.Root(ResultAccountNonce))
}

type MapTokenIDToImmutableBool struct {
	Proxy wasmtypes.Proxy
}

func (m MapTokenIDToImmutableBool) GetBool(key wasmtypes.ScTokenID) wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(m.Proxy.Key(wasmtypes.TokenIDToBytes(key)))
}

type ImmutableGetNativeTokenIDRegistryResults struct {
	Proxy wasmtypes.Proxy
}

// token IDs
func (s ImmutableGetNativeTokenIDRegistryResults) Mapping() MapTokenIDToImmutableBool {
	return MapTokenIDToImmutableBool(s)
}

type MapTokenIDToMutableBool struct {
	Proxy wasmtypes.Proxy
}

func (m MapTokenIDToMutableBool) Clear() {
	m.Proxy.ClearMap()
}

func (m MapTokenIDToMutableBool) GetBool(key wasmtypes.ScTokenID) wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(m.Proxy.Key(wasmtypes.TokenIDToBytes(key)))
}

type MutableGetNativeTokenIDRegistryResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetNativeTokenIDRegistryResults() MutableGetNativeTokenIDRegistryResults {
	return MutableGetNativeTokenIDRegistryResults{Proxy: wasmlib.NewResultsProxy()}
}

// token IDs
func (s MutableGetNativeTokenIDRegistryResults) Mapping() MapTokenIDToMutableBool {
	return MapTokenIDToMutableBool(s)
}

type ImmutableNativeTokenResults struct {
	Proxy wasmtypes.Proxy
}

// serialized foundry output
func (s ImmutableNativeTokenResults) FoundryOutputBin() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.Proxy.Root(ResultFoundryOutputBin))
}

type MutableNativeTokenResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableNativeTokenResults() MutableNativeTokenResults {
	return MutableNativeTokenResults{Proxy: wasmlib.NewResultsProxy()}
}

// serialized foundry output
func (s MutableNativeTokenResults) FoundryOutputBin() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.Proxy.Root(ResultFoundryOutputBin))
}

type ImmutableNftDataResults struct {
	Proxy wasmtypes.Proxy
}

// serialized NFT data
func (s ImmutableNftDataResults) NftData() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.Proxy.Root(ResultNftData))
}

type MutableNftDataResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableNftDataResults() MutableNftDataResults {
	return MutableNftDataResults{Proxy: wasmlib.NewResultsProxy()}
}

// serialized NFT data
func (s MutableNftDataResults) NftData() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.Proxy.Root(ResultNftData))
}

type ImmutableTotalAssetsResults struct {
	Proxy wasmtypes.Proxy
}

// balance per token ID, zero length indicates base token
func (s ImmutableTotalAssetsResults) Assets() MapTokenIDToImmutableBigInt {
	return MapTokenIDToImmutableBigInt(s)
}

type MutableTotalAssetsResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableTotalAssetsResults() MutableTotalAssetsResults {
	return MutableTotalAssetsResults{Proxy: wasmlib.NewResultsProxy()}
}

// balance per token ID, zero length indicates base token
func (s MutableTotalAssetsResults) Assets() MapTokenIDToMutableBigInt {
	return MapTokenIDToMutableBigInt(s)
}
