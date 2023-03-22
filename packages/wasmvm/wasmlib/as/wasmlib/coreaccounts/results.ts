// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

import * as wasmtypes from '../wasmtypes';
import * as sc from './index';

export class ImmutableFoundryCreateNewResults extends wasmtypes.ScProxy {
    // serial number of the newly created foundry
    foundrySN(): wasmtypes.ScImmutableUint32 {
        return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ResultFoundrySN));
    }
}

export class MutableFoundryCreateNewResults extends wasmtypes.ScProxy {
    // serial number of the newly created foundry
    foundrySN(): wasmtypes.ScMutableUint32 {
        return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ResultFoundrySN));
    }
}

export class MapUint32ToImmutableBytes extends wasmtypes.ScProxy {

    getBytes(key: u32): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.key(wasmtypes.uint32ToBytes(key)));
    }
}

export class ImmutableAccountFoundriesResults extends wasmtypes.ScProxy {
    // foundry serial numbers owned by the given account
    foundries(): sc.MapUint32ToImmutableBytes {
        return new sc.MapUint32ToImmutableBytes(this.proxy);
    }
}

export class MapUint32ToMutableBytes extends wasmtypes.ScProxy {

    clear(): void {
        this.proxy.clearMap();
    }

    getBytes(key: u32): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.key(wasmtypes.uint32ToBytes(key)));
    }
}

export class MutableAccountFoundriesResults extends wasmtypes.ScProxy {
    // foundry serial numbers owned by the given account
    foundries(): sc.MapUint32ToMutableBytes {
        return new sc.MapUint32ToMutableBytes(this.proxy);
    }
}

export class ImmutableAccountNFTAmountResults extends wasmtypes.ScProxy {
    // amount of NFTs owned by the account
    amount(): wasmtypes.ScImmutableUint32 {
        return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ResultAmount));
    }
}

export class MutableAccountNFTAmountResults extends wasmtypes.ScProxy {
    // amount of NFTs owned by the account
    amount(): wasmtypes.ScMutableUint32 {
        return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ResultAmount));
    }
}

export class ImmutableAccountNFTAmountInCollectionResults extends wasmtypes.ScProxy {
    // amount of NFTs in collection owned by the account
    amount(): wasmtypes.ScImmutableUint32 {
        return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ResultAmount));
    }
}

export class MutableAccountNFTAmountInCollectionResults extends wasmtypes.ScProxy {
    // amount of NFTs in collection owned by the account
    amount(): wasmtypes.ScMutableUint32 {
        return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ResultAmount));
    }
}

export class ArrayOfImmutableNftID extends wasmtypes.ScProxy {

    length(): u32 {
        return this.proxy.length();
    }

    getNftID(index: u32): wasmtypes.ScImmutableNftID {
        return new wasmtypes.ScImmutableNftID(this.proxy.index(index));
    }
}

export class ImmutableAccountNFTsResults extends wasmtypes.ScProxy {
    // NFT IDs owned by the account
    nftIDs(): sc.ArrayOfImmutableNftID {
        return new sc.ArrayOfImmutableNftID(this.proxy.root(sc.ResultNftIDs));
    }
}

export class ArrayOfMutableNftID extends wasmtypes.ScProxy {

    appendNftID(): wasmtypes.ScMutableNftID {
        return new wasmtypes.ScMutableNftID(this.proxy.append());
    }

    clear(): void {
        this.proxy.clearArray();
    }

    length(): u32 {
        return this.proxy.length();
    }

    getNftID(index: u32): wasmtypes.ScMutableNftID {
        return new wasmtypes.ScMutableNftID(this.proxy.index(index));
    }
}

export class MutableAccountNFTsResults extends wasmtypes.ScProxy {
    // NFT IDs owned by the account
    nftIDs(): sc.ArrayOfMutableNftID {
        return new sc.ArrayOfMutableNftID(this.proxy.root(sc.ResultNftIDs));
    }
}

export class ImmutableAccountNFTsInCollectionResults extends wasmtypes.ScProxy {
    // NFT IDs in collection owned by the account
    nftIDs(): sc.ArrayOfImmutableNftID {
        return new sc.ArrayOfImmutableNftID(this.proxy.root(sc.ResultNftIDs));
    }
}

export class MutableAccountNFTsInCollectionResults extends wasmtypes.ScProxy {
    // NFT IDs in collection owned by the account
    nftIDs(): sc.ArrayOfMutableNftID {
        return new sc.ArrayOfMutableNftID(this.proxy.root(sc.ResultNftIDs));
    }
}

export class MapAgentIDToImmutableBytes extends wasmtypes.ScProxy {

    getBytes(key: wasmtypes.ScAgentID): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.key(wasmtypes.agentIDToBytes(key)));
    }
}

export class ImmutableAccountsResults extends wasmtypes.ScProxy {
    // agent IDs
    allAccounts(): sc.MapAgentIDToImmutableBytes {
        return new sc.MapAgentIDToImmutableBytes(this.proxy);
    }
}

export class MapAgentIDToMutableBytes extends wasmtypes.ScProxy {

    clear(): void {
        this.proxy.clearMap();
    }

    getBytes(key: wasmtypes.ScAgentID): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.key(wasmtypes.agentIDToBytes(key)));
    }
}

export class MutableAccountsResults extends wasmtypes.ScProxy {
    // agent IDs
    allAccounts(): sc.MapAgentIDToMutableBytes {
        return new sc.MapAgentIDToMutableBytes(this.proxy);
    }
}

export class MapTokenIDToImmutableBigInt extends wasmtypes.ScProxy {

    getBigInt(key: wasmtypes.ScTokenID): wasmtypes.ScImmutableBigInt {
        return new wasmtypes.ScImmutableBigInt(this.proxy.key(wasmtypes.tokenIDToBytes(key)));
    }
}

export class ImmutableBalanceResults extends wasmtypes.ScProxy {
    // balance per token ID, zero length indicates base token
    balances(): sc.MapTokenIDToImmutableBigInt {
        return new sc.MapTokenIDToImmutableBigInt(this.proxy);
    }
}

export class MapTokenIDToMutableBigInt extends wasmtypes.ScProxy {

    clear(): void {
        this.proxy.clearMap();
    }

    getBigInt(key: wasmtypes.ScTokenID): wasmtypes.ScMutableBigInt {
        return new wasmtypes.ScMutableBigInt(this.proxy.key(wasmtypes.tokenIDToBytes(key)));
    }
}

export class MutableBalanceResults extends wasmtypes.ScProxy {
    // balance per token ID, zero length indicates base token
    balances(): sc.MapTokenIDToMutableBigInt {
        return new sc.MapTokenIDToMutableBigInt(this.proxy);
    }
}

export class ImmutableBalanceBaseTokenResults extends wasmtypes.ScProxy {
    // amount of base tokens in the account
    balance(): wasmtypes.ScImmutableUint64 {
        return new wasmtypes.ScImmutableUint64(this.proxy.root(sc.ResultBalance));
    }
}

export class MutableBalanceBaseTokenResults extends wasmtypes.ScProxy {
    // amount of base tokens in the account
    balance(): wasmtypes.ScMutableUint64 {
        return new wasmtypes.ScMutableUint64(this.proxy.root(sc.ResultBalance));
    }
}

export class ImmutableBalanceNativeTokenResults extends wasmtypes.ScProxy {
    // amount of native tokens in the account
    tokens(): wasmtypes.ScImmutableBigInt {
        return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ResultTokens));
    }
}

export class MutableBalanceNativeTokenResults extends wasmtypes.ScProxy {
    // amount of native tokens in the account
    tokens(): wasmtypes.ScMutableBigInt {
        return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ResultTokens));
    }
}

export class ImmutableFoundryOutputResults extends wasmtypes.ScProxy {
    // serialized foundry output
    foundryOutputBin(): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ResultFoundryOutputBin));
    }
}

export class MutableFoundryOutputResults extends wasmtypes.ScProxy {
    // serialized foundry output
    foundryOutputBin(): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ResultFoundryOutputBin));
    }
}

export class ImmutableGetAccountNonceResults extends wasmtypes.ScProxy {
    // account nonce
    accountNonce(): wasmtypes.ScImmutableUint64 {
        return new wasmtypes.ScImmutableUint64(this.proxy.root(sc.ResultAccountNonce));
    }
}

export class MutableGetAccountNonceResults extends wasmtypes.ScProxy {
    // account nonce
    accountNonce(): wasmtypes.ScMutableUint64 {
        return new wasmtypes.ScMutableUint64(this.proxy.root(sc.ResultAccountNonce));
    }
}

export class MapTokenIDToImmutableBytes extends wasmtypes.ScProxy {

    getBytes(key: wasmtypes.ScTokenID): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.key(wasmtypes.tokenIDToBytes(key)));
    }
}

export class ImmutableGetNativeTokenIDRegistryResults extends wasmtypes.ScProxy {
    // token IDs
    mapping(): sc.MapTokenIDToImmutableBytes {
        return new sc.MapTokenIDToImmutableBytes(this.proxy);
    }
}

export class MapTokenIDToMutableBytes extends wasmtypes.ScProxy {

    clear(): void {
        this.proxy.clearMap();
    }

    getBytes(key: wasmtypes.ScTokenID): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.key(wasmtypes.tokenIDToBytes(key)));
    }
}

export class MutableGetNativeTokenIDRegistryResults extends wasmtypes.ScProxy {
    // token IDs
    mapping(): sc.MapTokenIDToMutableBytes {
        return new sc.MapTokenIDToMutableBytes(this.proxy);
    }
}

export class ImmutableNftDataResults extends wasmtypes.ScProxy {
    // serialized NFT data
    nftData(): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ResultNftData));
    }
}

export class MutableNftDataResults extends wasmtypes.ScProxy {
    // serialized NFT data
    nftData(): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ResultNftData));
    }
}

export class ImmutableTotalAssetsResults extends wasmtypes.ScProxy {
    // balance per token ID, zero length indicates base token
    assets(): sc.MapTokenIDToImmutableBigInt {
        return new sc.MapTokenIDToImmutableBigInt(this.proxy);
    }
}

export class MutableTotalAssetsResults extends wasmtypes.ScProxy {
    // balance per token ID, zero length indicates base token
    assets(): sc.MapTokenIDToMutableBigInt {
        return new sc.MapTokenIDToMutableBigInt(this.proxy);
    }
}
