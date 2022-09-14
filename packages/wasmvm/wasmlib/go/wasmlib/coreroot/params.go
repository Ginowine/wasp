// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead
package coreroot

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableDeployContractParams struct {
	proxy wasmtypes.Proxy
}

// default 'N/A'
func (s ImmutableDeployContractParams) Description() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamDescription))
}

func (s ImmutableDeployContractParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

// TODO variable init params for deployed contract
func (s ImmutableDeployContractParams) ProgramHash() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamProgramHash))
}

type MutableDeployContractParams struct {
	proxy wasmtypes.Proxy
}

// default 'N/A'
func (s MutableDeployContractParams) Description() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamDescription))
}

func (s MutableDeployContractParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

// TODO variable init params for deployed contract
func (s MutableDeployContractParams) ProgramHash() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamProgramHash))
}

type ImmutableGrantDeployPermissionParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGrantDeployPermissionParams) Deployer() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamDeployer))
}

type MutableGrantDeployPermissionParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGrantDeployPermissionParams) Deployer() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamDeployer))
}

type ImmutableRequireDeployPermissionsParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableRequireDeployPermissionsParams) DeployPermissionsEnabled() wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(s.proxy.Root(ParamDeployPermissionsEnabled))
}

type MutableRequireDeployPermissionsParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableRequireDeployPermissionsParams) DeployPermissionsEnabled() wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(s.proxy.Root(ParamDeployPermissionsEnabled))
}

type ImmutableRevokeDeployPermissionParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableRevokeDeployPermissionParams) Deployer() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamDeployer))
}

type MutableRevokeDeployPermissionParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableRevokeDeployPermissionParams) Deployer() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamDeployer))
}

type ImmutableSubscribeBlockContextParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableSubscribeBlockContextParams) CloseFunc() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamCloseFunc))
}

func (s ImmutableSubscribeBlockContextParams) OpenFunc() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamOpenFunc))
}

type MutableSubscribeBlockContextParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableSubscribeBlockContextParams) CloseFunc() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamCloseFunc))
}

func (s MutableSubscribeBlockContextParams) OpenFunc() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamOpenFunc))
}

type ImmutableFindContractParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableFindContractParams) Hname() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamHname))
}

type MutableFindContractParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableFindContractParams) Hname() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamHname))
}
