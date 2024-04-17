// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use crate::*;

pub const SC_NAME        : &str = "accounts";
pub const SC_DESCRIPTION : &str = "Chain account ledger contract";
pub const HSC_NAME       : ScHname = ScHname(0x3c4b5e02);

pub(crate) const PARAM_AGENT_ID         : &str = "a";
pub(crate) const PARAM_COLLECTION       : &str = "C";
pub(crate) const PARAM_DESTROY_TOKENS   : &str = "y";
pub(crate) const PARAM_FOUNDRY_SN       : &str = "s";
pub(crate) const PARAM_GAS_RESERVE      : &str = "g";
pub(crate) const PARAM_NFT_ID           : &str = "z";
pub(crate) const PARAM_SUPPLY_DELTA_ABS : &str = "d";
pub(crate) const PARAM_TOKEN_DECIMALS   : &str = "td";
pub(crate) const PARAM_TOKEN_ID         : &str = "N";
pub(crate) const PARAM_TOKEN_NAME       : &str = "tn";
pub(crate) const PARAM_TOKEN_SCHEME     : &str = "t";
pub(crate) const PARAM_TOKEN_SYMBOL     : &str = "ts";

pub(crate) const RESULT_ACCOUNT_NONCE      : &str = "n";
pub(crate) const RESULT_ALL_ACCOUNTS       : &str = "this";
pub(crate) const RESULT_AMOUNT             : &str = "A";
pub(crate) const RESULT_ASSETS             : &str = "this";
pub(crate) const RESULT_BALANCE            : &str = "B";
pub(crate) const RESULT_BALANCES           : &str = "this";
pub(crate) const RESULT_FOUNDRIES          : &str = "this";
pub(crate) const RESULT_FOUNDRY_OUTPUT_BIN : &str = "b";
pub(crate) const RESULT_FOUNDRY_SN         : &str = "s";
pub(crate) const RESULT_MAPPING            : &str = "this";
pub(crate) const RESULT_NFT_DATA           : &str = "e";
pub(crate) const RESULT_NFT_I_DS           : &str = "i";
pub(crate) const RESULT_TOKENS             : &str = "B";

pub(crate) const FUNC_DEPOSIT                          : &str = "deposit";
pub(crate) const FUNC_FOUNDRY_CREATE_NEW               : &str = "foundryCreateNew";
pub(crate) const FUNC_NATIVE_TOKEN_CREATE              : &str = "nativeTokenCreate";
pub(crate) const FUNC_NATIVE_TOKEN_DESTROY             : &str = "nativeTokenDestroy";
pub(crate) const FUNC_NATIVE_TOKEN_MODIFY_SUPPLY       : &str = "nativeTokenModifySupply";
pub(crate) const FUNC_TRANSFER_ACCOUNT_TO_CHAIN        : &str = "transferAccountToChain";
pub(crate) const FUNC_TRANSFER_ALLOWANCE_TO            : &str = "transferAllowanceTo";
pub(crate) const FUNC_WITHDRAW                         : &str = "withdraw";
pub(crate) const VIEW_ACCOUNT_FOUNDRIES                : &str = "accountFoundries";
pub(crate) const VIEW_ACCOUNT_NFT_AMOUNT               : &str = "accountNFTAmount";
pub(crate) const VIEW_ACCOUNT_NFT_AMOUNT_IN_COLLECTION : &str = "accountNFTAmountInCollection";
pub(crate) const VIEW_ACCOUNT_NF_TS                    : &str = "accountNFTs";
pub(crate) const VIEW_ACCOUNT_NF_TS_IN_COLLECTION      : &str = "accountNFTsInCollection";
pub(crate) const VIEW_ACCOUNTS                         : &str = "accounts";
pub(crate) const VIEW_BALANCE                          : &str = "balance";
pub(crate) const VIEW_BALANCE_BASE_TOKEN               : &str = "balanceBaseToken";
pub(crate) const VIEW_BALANCE_NATIVE_TOKEN             : &str = "balanceNativeToken";
pub(crate) const VIEW_GET_ACCOUNT_NONCE                : &str = "getAccountNonce";
pub(crate) const VIEW_GET_NATIVE_TOKEN_ID_REGISTRY     : &str = "getNativeTokenIDRegistry";
pub(crate) const VIEW_NATIVE_TOKEN                     : &str = "nativeToken";
pub(crate) const VIEW_NFT_DATA                         : &str = "nftData";
pub(crate) const VIEW_TOTAL_ASSETS                     : &str = "totalAssets";

pub(crate) const HFUNC_DEPOSIT                          : ScHname = ScHname(0xbdc9102d);
pub(crate) const HFUNC_FOUNDRY_CREATE_NEW               : ScHname = ScHname(0x41822f5f);
pub(crate) const HFUNC_NATIVE_TOKEN_CREATE              : ScHname = ScHname(0x0c2d1791);
pub(crate) const HFUNC_NATIVE_TOKEN_DESTROY             : ScHname = ScHname(0xf0b0ab00);
pub(crate) const HFUNC_NATIVE_TOKEN_MODIFY_SUPPLY       : ScHname = ScHname(0x24c2eab6);
pub(crate) const HFUNC_TRANSFER_ACCOUNT_TO_CHAIN        : ScHname = ScHname(0x07005c45);
pub(crate) const HFUNC_TRANSFER_ALLOWANCE_TO            : ScHname = ScHname(0x23f4e3a1);
pub(crate) const HFUNC_WITHDRAW                         : ScHname = ScHname(0x9dcc0f41);
pub(crate) const HVIEW_ACCOUNT_FOUNDRIES                : ScHname = ScHname(0xdc3a0c38);
pub(crate) const HVIEW_ACCOUNT_NFT_AMOUNT               : ScHname = ScHname(0xabefd5b5);
pub(crate) const HVIEW_ACCOUNT_NFT_AMOUNT_IN_COLLECTION : ScHname = ScHname(0xd7028e1b);
pub(crate) const HVIEW_ACCOUNT_NF_TS                    : ScHname = ScHname(0x27422359);
pub(crate) const HVIEW_ACCOUNT_NF_TS_IN_COLLECTION      : ScHname = ScHname(0xa37fb50f);
pub(crate) const HVIEW_ACCOUNTS                         : ScHname = ScHname(0x3c4b5e02);
pub(crate) const HVIEW_BALANCE                          : ScHname = ScHname(0x84168cb4);
pub(crate) const HVIEW_BALANCE_BASE_TOKEN               : ScHname = ScHname(0x4c8ccd0f);
pub(crate) const HVIEW_BALANCE_NATIVE_TOKEN             : ScHname = ScHname(0x1fea3104);
pub(crate) const HVIEW_GET_ACCOUNT_NONCE                : ScHname = ScHname(0x529d7df9);
pub(crate) const HVIEW_GET_NATIVE_TOKEN_ID_REGISTRY     : ScHname = ScHname(0x2ad8a59f);
pub(crate) const HVIEW_NATIVE_TOKEN                     : ScHname = ScHname(0x28e34b65);
pub(crate) const HVIEW_NFT_DATA                         : ScHname = ScHname(0x83c5c4da);
pub(crate) const HVIEW_TOTAL_ASSETS                     : ScHname = ScHname(0xfab0f8d2);
