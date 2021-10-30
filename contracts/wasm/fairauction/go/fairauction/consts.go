// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	ScName        = "fairauction"
	ScDescription = "Decentralized auction to securely sell tokens to the highest bidder"
	HScName       = wasmlib.ScHname(0x1b5c43b1)
)

const (
	ParamColor       = wasmlib.Key("color")
	ParamDescription = wasmlib.Key("description")
	ParamDuration    = wasmlib.Key("duration")
	ParamMinimumBid  = wasmlib.Key("minimumBid")
	ParamOwnerMargin = wasmlib.Key("ownerMargin")
)

const (
	ResultBidders       = wasmlib.Key("bidders")
	ResultColor         = wasmlib.Key("color")
	ResultCreator       = wasmlib.Key("creator")
	ResultDeposit       = wasmlib.Key("deposit")
	ResultDescription   = wasmlib.Key("description")
	ResultDuration      = wasmlib.Key("duration")
	ResultHighestBid    = wasmlib.Key("highestBid")
	ResultHighestBidder = wasmlib.Key("highestBidder")
	ResultMinimumBid    = wasmlib.Key("minimumBid")
	ResultNumTokens     = wasmlib.Key("numTokens")
	ResultOwnerMargin   = wasmlib.Key("ownerMargin")
	ResultWhenStarted   = wasmlib.Key("whenStarted")
)

const (
	StateAuctions    = wasmlib.Key("auctions")
	StateBidderList  = wasmlib.Key("bidderList")
	StateBids        = wasmlib.Key("bids")
	StateOwnerMargin = wasmlib.Key("ownerMargin")
)

const (
	FuncFinalizeAuction = "finalizeAuction"
	FuncPlaceBid        = "placeBid"
	FuncSetOwnerMargin  = "setOwnerMargin"
	FuncStartAuction    = "startAuction"
	ViewGetInfo         = "getInfo"
)

const (
	HFuncFinalizeAuction = wasmlib.ScHname(0x8d534ddc)
	HFuncPlaceBid        = wasmlib.ScHname(0x9bd72fa9)
	HFuncSetOwnerMargin  = wasmlib.ScHname(0x1774461a)
	HFuncStartAuction    = wasmlib.ScHname(0xd5b7bacb)
	HViewGetInfo         = wasmlib.ScHname(0xcfedba5f)
)
