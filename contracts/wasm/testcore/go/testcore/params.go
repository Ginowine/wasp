// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testcore

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type ImmutableCallOnChainParams struct {
	id int32
}

func (s ImmutableCallOnChainParams) HnameContract() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, idxMap[IdxParamHnameContract])
}

func (s ImmutableCallOnChainParams) HnameEP() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, idxMap[IdxParamHnameEP])
}

func (s ImmutableCallOnChainParams) IntValue() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamIntValue])
}

type MutableCallOnChainParams struct {
	id int32
}

func (s MutableCallOnChainParams) HnameContract() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, idxMap[IdxParamHnameContract])
}

func (s MutableCallOnChainParams) HnameEP() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, idxMap[IdxParamHnameEP])
}

func (s MutableCallOnChainParams) IntValue() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamIntValue])
}

type ImmutableCheckContextFromFullEPParams struct {
	id int32
}

func (s ImmutableCheckContextFromFullEPParams) AgentID() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamAgentID])
}

func (s ImmutableCheckContextFromFullEPParams) Caller() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamCaller])
}

func (s ImmutableCheckContextFromFullEPParams) ChainID() wasmlib.ScImmutableChainID {
	return wasmlib.NewScImmutableChainID(s.id, idxMap[IdxParamChainID])
}

func (s ImmutableCheckContextFromFullEPParams) ChainOwnerID() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamChainOwnerID])
}

func (s ImmutableCheckContextFromFullEPParams) ContractCreator() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamContractCreator])
}

type MutableCheckContextFromFullEPParams struct {
	id int32
}

func (s MutableCheckContextFromFullEPParams) AgentID() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamAgentID])
}

func (s MutableCheckContextFromFullEPParams) Caller() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamCaller])
}

func (s MutableCheckContextFromFullEPParams) ChainID() wasmlib.ScMutableChainID {
	return wasmlib.NewScMutableChainID(s.id, idxMap[IdxParamChainID])
}

func (s MutableCheckContextFromFullEPParams) ChainOwnerID() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamChainOwnerID])
}

func (s MutableCheckContextFromFullEPParams) ContractCreator() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamContractCreator])
}

type ImmutableInitParams struct {
	id int32
}

func (s ImmutableInitParams) Fail() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamFail])
}

type MutableInitParams struct {
	id int32
}

func (s MutableInitParams) Fail() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamFail])
}

type ImmutablePassTypesFullParams struct {
	id int32
}

func (s ImmutablePassTypesFullParams) Address() wasmlib.ScImmutableAddress {
	return wasmlib.NewScImmutableAddress(s.id, idxMap[IdxParamAddress])
}

func (s ImmutablePassTypesFullParams) AgentID() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamAgentID])
}

func (s ImmutablePassTypesFullParams) ChainID() wasmlib.ScImmutableChainID {
	return wasmlib.NewScImmutableChainID(s.id, idxMap[IdxParamChainID])
}

func (s ImmutablePassTypesFullParams) ContractID() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamContractID])
}

func (s ImmutablePassTypesFullParams) Hash() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, idxMap[IdxParamHash])
}

func (s ImmutablePassTypesFullParams) Hname() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, idxMap[IdxParamHname])
}

func (s ImmutablePassTypesFullParams) HnameZero() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, idxMap[IdxParamHnameZero])
}

func (s ImmutablePassTypesFullParams) Int64() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamInt64])
}

func (s ImmutablePassTypesFullParams) Int64Zero() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamInt64Zero])
}

func (s ImmutablePassTypesFullParams) String() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamString])
}

func (s ImmutablePassTypesFullParams) StringZero() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamStringZero])
}

type MutablePassTypesFullParams struct {
	id int32
}

func (s MutablePassTypesFullParams) Address() wasmlib.ScMutableAddress {
	return wasmlib.NewScMutableAddress(s.id, idxMap[IdxParamAddress])
}

func (s MutablePassTypesFullParams) AgentID() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamAgentID])
}

func (s MutablePassTypesFullParams) ChainID() wasmlib.ScMutableChainID {
	return wasmlib.NewScMutableChainID(s.id, idxMap[IdxParamChainID])
}

func (s MutablePassTypesFullParams) ContractID() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamContractID])
}

func (s MutablePassTypesFullParams) Hash() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, idxMap[IdxParamHash])
}

func (s MutablePassTypesFullParams) Hname() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, idxMap[IdxParamHname])
}

func (s MutablePassTypesFullParams) HnameZero() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, idxMap[IdxParamHnameZero])
}

func (s MutablePassTypesFullParams) Int64() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamInt64])
}

func (s MutablePassTypesFullParams) Int64Zero() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamInt64Zero])
}

func (s MutablePassTypesFullParams) String() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamString])
}

func (s MutablePassTypesFullParams) StringZero() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamStringZero])
}

type ImmutableRunRecursionParams struct {
	id int32
}

func (s ImmutableRunRecursionParams) IntValue() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamIntValue])
}

type MutableRunRecursionParams struct {
	id int32
}

func (s MutableRunRecursionParams) IntValue() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamIntValue])
}

type ImmutableSendToAddressParams struct {
	id int32
}

func (s ImmutableSendToAddressParams) Address() wasmlib.ScImmutableAddress {
	return wasmlib.NewScImmutableAddress(s.id, idxMap[IdxParamAddress])
}

type MutableSendToAddressParams struct {
	id int32
}

func (s MutableSendToAddressParams) Address() wasmlib.ScMutableAddress {
	return wasmlib.NewScMutableAddress(s.id, idxMap[IdxParamAddress])
}

type ImmutableSetIntParams struct {
	id int32
}

func (s ImmutableSetIntParams) IntValue() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamIntValue])
}

func (s ImmutableSetIntParams) Name() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamName])
}

type MutableSetIntParams struct {
	id int32
}

func (s MutableSetIntParams) IntValue() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamIntValue])
}

func (s MutableSetIntParams) Name() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamName])
}

type ImmutableSpawnParams struct {
	id int32
}

func (s ImmutableSpawnParams) ProgHash() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, idxMap[IdxParamProgHash])
}

type MutableSpawnParams struct {
	id int32
}

func (s MutableSpawnParams) ProgHash() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, idxMap[IdxParamProgHash])
}

type ImmutableTestEventLogGenericDataParams struct {
	id int32
}

func (s ImmutableTestEventLogGenericDataParams) Counter() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamCounter])
}

type MutableTestEventLogGenericDataParams struct {
	id int32
}

func (s MutableTestEventLogGenericDataParams) Counter() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamCounter])
}

type ImmutableWithdrawToChainParams struct {
	id int32
}

func (s ImmutableWithdrawToChainParams) ChainID() wasmlib.ScImmutableChainID {
	return wasmlib.NewScImmutableChainID(s.id, idxMap[IdxParamChainID])
}

type MutableWithdrawToChainParams struct {
	id int32
}

func (s MutableWithdrawToChainParams) ChainID() wasmlib.ScMutableChainID {
	return wasmlib.NewScMutableChainID(s.id, idxMap[IdxParamChainID])
}

type ImmutableCheckContextFromViewEPParams struct {
	id int32
}

func (s ImmutableCheckContextFromViewEPParams) AgentID() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamAgentID])
}

func (s ImmutableCheckContextFromViewEPParams) ChainID() wasmlib.ScImmutableChainID {
	return wasmlib.NewScImmutableChainID(s.id, idxMap[IdxParamChainID])
}

func (s ImmutableCheckContextFromViewEPParams) ChainOwnerID() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamChainOwnerID])
}

func (s ImmutableCheckContextFromViewEPParams) ContractCreator() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamContractCreator])
}

type MutableCheckContextFromViewEPParams struct {
	id int32
}

func (s MutableCheckContextFromViewEPParams) AgentID() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamAgentID])
}

func (s MutableCheckContextFromViewEPParams) ChainID() wasmlib.ScMutableChainID {
	return wasmlib.NewScMutableChainID(s.id, idxMap[IdxParamChainID])
}

func (s MutableCheckContextFromViewEPParams) ChainOwnerID() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamChainOwnerID])
}

func (s MutableCheckContextFromViewEPParams) ContractCreator() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamContractCreator])
}

type ImmutableFibonacciParams struct {
	id int32
}

func (s ImmutableFibonacciParams) IntValue() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamIntValue])
}

type MutableFibonacciParams struct {
	id int32
}

func (s MutableFibonacciParams) IntValue() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamIntValue])
}

type ImmutableGetIntParams struct {
	id int32
}

func (s ImmutableGetIntParams) Name() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamName])
}

type MutableGetIntParams struct {
	id int32
}

func (s MutableGetIntParams) Name() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamName])
}

type ImmutableGetStringValueParams struct {
	id int32
}

func (s ImmutableGetStringValueParams) VarName() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamVarName])
}

type MutableGetStringValueParams struct {
	id int32
}

func (s MutableGetStringValueParams) VarName() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamVarName])
}

type ImmutablePassTypesViewParams struct {
	id int32
}

func (s ImmutablePassTypesViewParams) Address() wasmlib.ScImmutableAddress {
	return wasmlib.NewScImmutableAddress(s.id, idxMap[IdxParamAddress])
}

func (s ImmutablePassTypesViewParams) AgentID() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamAgentID])
}

func (s ImmutablePassTypesViewParams) ChainID() wasmlib.ScImmutableChainID {
	return wasmlib.NewScImmutableChainID(s.id, idxMap[IdxParamChainID])
}

func (s ImmutablePassTypesViewParams) ContractID() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxParamContractID])
}

func (s ImmutablePassTypesViewParams) Hash() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, idxMap[IdxParamHash])
}

func (s ImmutablePassTypesViewParams) Hname() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, idxMap[IdxParamHname])
}

func (s ImmutablePassTypesViewParams) HnameZero() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, idxMap[IdxParamHnameZero])
}

func (s ImmutablePassTypesViewParams) Int64() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamInt64])
}

func (s ImmutablePassTypesViewParams) Int64Zero() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamInt64Zero])
}

func (s ImmutablePassTypesViewParams) String() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamString])
}

func (s ImmutablePassTypesViewParams) StringZero() wasmlib.ScImmutableString {
	return wasmlib.NewScImmutableString(s.id, idxMap[IdxParamStringZero])
}

type MutablePassTypesViewParams struct {
	id int32
}

func (s MutablePassTypesViewParams) Address() wasmlib.ScMutableAddress {
	return wasmlib.NewScMutableAddress(s.id, idxMap[IdxParamAddress])
}

func (s MutablePassTypesViewParams) AgentID() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamAgentID])
}

func (s MutablePassTypesViewParams) ChainID() wasmlib.ScMutableChainID {
	return wasmlib.NewScMutableChainID(s.id, idxMap[IdxParamChainID])
}

func (s MutablePassTypesViewParams) ContractID() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxParamContractID])
}

func (s MutablePassTypesViewParams) Hash() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, idxMap[IdxParamHash])
}

func (s MutablePassTypesViewParams) Hname() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, idxMap[IdxParamHname])
}

func (s MutablePassTypesViewParams) HnameZero() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, idxMap[IdxParamHnameZero])
}

func (s MutablePassTypesViewParams) Int64() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamInt64])
}

func (s MutablePassTypesViewParams) Int64Zero() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamInt64Zero])
}

func (s MutablePassTypesViewParams) String() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamString])
}

func (s MutablePassTypesViewParams) StringZero() wasmlib.ScMutableString {
	return wasmlib.NewScMutableString(s.id, idxMap[IdxParamStringZero])
}
