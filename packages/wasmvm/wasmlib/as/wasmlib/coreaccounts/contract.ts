// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

import * as wasmlib from '../index';
import * as sc from './index';

export class DepositCall {
    func:   wasmlib.ScFunc;

    public constructor(ctx: wasmlib.ScFuncClientContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncDeposit);
    }
}

export class FoundryCreateNewCall {
    func:    wasmlib.ScFunc;
    params:  sc.MutableFoundryCreateNewParams = new sc.MutableFoundryCreateNewParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableFoundryCreateNewResults = new sc.ImmutableFoundryCreateNewResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScFuncClientContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncFoundryCreateNew);
    }
}

export class NativeTokenCreateCall {
    func:   wasmlib.ScFunc;
    params: sc.MutableNativeTokenCreateParams = new sc.MutableNativeTokenCreateParams(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScFuncClientContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncNativeTokenCreate);
    }
}

export class NativeTokenDestroyCall {
    func:   wasmlib.ScFunc;
    params: sc.MutableNativeTokenDestroyParams = new sc.MutableNativeTokenDestroyParams(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScFuncClientContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncNativeTokenDestroy);
    }
}

export class NativeTokenModifySupplyCall {
    func:   wasmlib.ScFunc;
    params: sc.MutableNativeTokenModifySupplyParams = new sc.MutableNativeTokenModifySupplyParams(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScFuncClientContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncNativeTokenModifySupply);
    }
}

export class TransferAccountToChainCall {
    func:   wasmlib.ScFunc;
    params: sc.MutableTransferAccountToChainParams = new sc.MutableTransferAccountToChainParams(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScFuncClientContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncTransferAccountToChain);
    }
}

export class TransferAllowanceToCall {
    func:   wasmlib.ScFunc;
    params: sc.MutableTransferAllowanceToParams = new sc.MutableTransferAllowanceToParams(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScFuncClientContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncTransferAllowanceTo);
    }
}

export class WithdrawCall {
    func:   wasmlib.ScFunc;

    public constructor(ctx: wasmlib.ScFuncClientContext) {
        this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncWithdraw);
    }
}

export class AccountFoundriesCall {
    func:    wasmlib.ScView;
    params:  sc.MutableAccountFoundriesParams = new sc.MutableAccountFoundriesParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableAccountFoundriesResults = new sc.ImmutableAccountFoundriesResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewAccountFoundries);
    }
}

export class AccountNFTAmountCall {
    func:    wasmlib.ScView;
    params:  sc.MutableAccountNFTAmountParams = new sc.MutableAccountNFTAmountParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableAccountNFTAmountResults = new sc.ImmutableAccountNFTAmountResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewAccountNFTAmount);
    }
}

export class AccountNFTAmountInCollectionCall {
    func:    wasmlib.ScView;
    params:  sc.MutableAccountNFTAmountInCollectionParams = new sc.MutableAccountNFTAmountInCollectionParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableAccountNFTAmountInCollectionResults = new sc.ImmutableAccountNFTAmountInCollectionResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewAccountNFTAmountInCollection);
    }
}

export class AccountNFTsCall {
    func:    wasmlib.ScView;
    params:  sc.MutableAccountNFTsParams = new sc.MutableAccountNFTsParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableAccountNFTsResults = new sc.ImmutableAccountNFTsResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewAccountNFTs);
    }
}

export class AccountNFTsInCollectionCall {
    func:    wasmlib.ScView;
    params:  sc.MutableAccountNFTsInCollectionParams = new sc.MutableAccountNFTsInCollectionParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableAccountNFTsInCollectionResults = new sc.ImmutableAccountNFTsInCollectionResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewAccountNFTsInCollection);
    }
}

export class AccountsCall {
    func:    wasmlib.ScView;
    results: sc.ImmutableAccountsResults = new sc.ImmutableAccountsResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewAccounts);
    }
}

export class BalanceCall {
    func:    wasmlib.ScView;
    params:  sc.MutableBalanceParams = new sc.MutableBalanceParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableBalanceResults = new sc.ImmutableBalanceResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewBalance);
    }
}

export class BalanceBaseTokenCall {
    func:    wasmlib.ScView;
    params:  sc.MutableBalanceBaseTokenParams = new sc.MutableBalanceBaseTokenParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableBalanceBaseTokenResults = new sc.ImmutableBalanceBaseTokenResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewBalanceBaseToken);
    }
}

export class BalanceNativeTokenCall {
    func:    wasmlib.ScView;
    params:  sc.MutableBalanceNativeTokenParams = new sc.MutableBalanceNativeTokenParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableBalanceNativeTokenResults = new sc.ImmutableBalanceNativeTokenResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewBalanceNativeToken);
    }
}

export class GetAccountNonceCall {
    func:    wasmlib.ScView;
    params:  sc.MutableGetAccountNonceParams = new sc.MutableGetAccountNonceParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableGetAccountNonceResults = new sc.ImmutableGetAccountNonceResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetAccountNonce);
    }
}

export class GetNativeTokenIDRegistryCall {
    func:    wasmlib.ScView;
    results: sc.ImmutableGetNativeTokenIDRegistryResults = new sc.ImmutableGetNativeTokenIDRegistryResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetNativeTokenIDRegistry);
    }
}

export class NativeTokenCall {
    func:    wasmlib.ScView;
    params:  sc.MutableNativeTokenParams = new sc.MutableNativeTokenParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableNativeTokenResults = new sc.ImmutableNativeTokenResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewNativeToken);
    }
}

export class NftDataCall {
    func:    wasmlib.ScView;
    params:  sc.MutableNftDataParams = new sc.MutableNftDataParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableNftDataResults = new sc.ImmutableNftDataResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewNftData);
    }
}

export class TotalAssetsCall {
    func:    wasmlib.ScView;
    results: sc.ImmutableTotalAssetsResults = new sc.ImmutableTotalAssetsResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewClientContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewTotalAssets);
    }
}

export class ScFuncs {
    // A no-op that has the side effect of crediting any transferred tokens to the sender's account.
    static deposit(ctx: wasmlib.ScFuncClientContext): DepositCall {
        return new DepositCall(ctx);
    }

    // Creates a new foundry with the specified token scheme, and assigns the foundry to the sender.
    static foundryCreateNew(ctx: wasmlib.ScFuncClientContext): FoundryCreateNewCall {
        const f = new FoundryCreateNewCall(ctx);
        f.params = new sc.MutableFoundryCreateNewParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableFoundryCreateNewResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Creates a new foundry and registers it as a ERC20 and IRC30 token
    static nativeTokenCreate(ctx: wasmlib.ScFuncClientContext): NativeTokenCreateCall {
        const f = new NativeTokenCreateCall(ctx);
        f.params = new sc.MutableNativeTokenCreateParams(wasmlib.newCallParamsProxy(f.func));
        return f;
    }

    // Destroys a given foundry output on L1, reimbursing the storage deposit to the caller.
    // The foundry must be owned by the caller.
    static nativeTokenDestroy(ctx: wasmlib.ScFuncClientContext): NativeTokenDestroyCall {
        const f = new NativeTokenDestroyCall(ctx);
        f.params = new sc.MutableNativeTokenDestroyParams(wasmlib.newCallParamsProxy(f.func));
        return f;
    }

    // Mints or destroys tokens for the given foundry, which must be owned by the caller.
    static nativeTokenModifySupply(ctx: wasmlib.ScFuncClientContext): NativeTokenModifySupplyCall {
        const f = new NativeTokenModifySupplyCall(ctx);
        f.params = new sc.MutableNativeTokenModifySupplyParams(wasmlib.newCallParamsProxy(f.func));
        return f;
    }

    // Transfers the specified allowance from the sender SC's L2 account on
    // the target chain to the sender SC's L2 account on the origin chain.
    static transferAccountToChain(ctx: wasmlib.ScFuncClientContext): TransferAccountToChainCall {
        const f = new TransferAccountToChainCall(ctx);
        f.params = new sc.MutableTransferAccountToChainParams(wasmlib.newCallParamsProxy(f.func));
        return f;
    }

    // Transfers the specified allowance from the sender's L2 account
    // to the given L2 account on the chain.
    static transferAllowanceTo(ctx: wasmlib.ScFuncClientContext): TransferAllowanceToCall {
        const f = new TransferAllowanceToCall(ctx);
        f.params = new sc.MutableTransferAllowanceToParams(wasmlib.newCallParamsProxy(f.func));
        return f;
    }

    // Moves tokens from the caller's on-chain account to the caller's L1 address.
    // The number of tokens to be withdrawn must be specified via the allowance of the request.
    static withdraw(ctx: wasmlib.ScFuncClientContext): WithdrawCall {
        return new WithdrawCall(ctx);
    }

    // Returns a set of all foundries owned by the given account.
    static accountFoundries(ctx: wasmlib.ScViewClientContext): AccountFoundriesCall {
        const f = new AccountFoundriesCall(ctx);
        f.params = new sc.MutableAccountFoundriesParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableAccountFoundriesResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns the amount of NFTs owned by the given account.
    static accountNFTAmount(ctx: wasmlib.ScViewClientContext): AccountNFTAmountCall {
        const f = new AccountNFTAmountCall(ctx);
        f.params = new sc.MutableAccountNFTAmountParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableAccountNFTAmountResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns the amount of NFTs in the specified collection owned by the given account.
    static accountNFTAmountInCollection(ctx: wasmlib.ScViewClientContext): AccountNFTAmountInCollectionCall {
        const f = new AccountNFTAmountInCollectionCall(ctx);
        f.params = new sc.MutableAccountNFTAmountInCollectionParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableAccountNFTAmountInCollectionResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns the NFT IDs for all NFTs owned by the given account.
    static accountNFTs(ctx: wasmlib.ScViewClientContext): AccountNFTsCall {
        const f = new AccountNFTsCall(ctx);
        f.params = new sc.MutableAccountNFTsParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableAccountNFTsResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns the NFT IDs for all NFTs in the specified collection owned by the given account.
    static accountNFTsInCollection(ctx: wasmlib.ScViewClientContext): AccountNFTsInCollectionCall {
        const f = new AccountNFTsInCollectionCall(ctx);
        f.params = new sc.MutableAccountNFTsInCollectionParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableAccountNFTsInCollectionResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns a set of all agent IDs that own assets on the chain.
    static accounts(ctx: wasmlib.ScViewClientContext): AccountsCall {
        const f = new AccountsCall(ctx);
        f.results = new sc.ImmutableAccountsResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns the fungible tokens owned by the given Agent ID on the chain.
    static balance(ctx: wasmlib.ScViewClientContext): BalanceCall {
        const f = new BalanceCall(ctx);
        f.params = new sc.MutableBalanceParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableBalanceResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns the amount of base tokens owned by an agent on the chain
    static balanceBaseToken(ctx: wasmlib.ScViewClientContext): BalanceBaseTokenCall {
        const f = new BalanceBaseTokenCall(ctx);
        f.params = new sc.MutableBalanceBaseTokenParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableBalanceBaseTokenResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns the amount of specific native tokens owned by an agent on the chain
    static balanceNativeToken(ctx: wasmlib.ScViewClientContext): BalanceNativeTokenCall {
        const f = new BalanceNativeTokenCall(ctx);
        f.params = new sc.MutableBalanceNativeTokenParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableBalanceNativeTokenResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns the current account nonce for an Agent.
    // The account nonce is used to issue unique off-ledger requests.
    static getAccountNonce(ctx: wasmlib.ScViewClientContext): GetAccountNonceCall {
        const f = new GetAccountNonceCall(ctx);
        f.params = new sc.MutableGetAccountNonceParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableGetAccountNonceResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns a set of all native tokenIDs that are owned by the chain.
    static getNativeTokenIDRegistry(ctx: wasmlib.ScViewClientContext): GetNativeTokenIDRegistryCall {
        const f = new GetNativeTokenIDRegistryCall(ctx);
        f.results = new sc.ImmutableGetNativeTokenIDRegistryResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns specified foundry output in serialized form.
    static nativeToken(ctx: wasmlib.ScViewClientContext): NativeTokenCall {
        const f = new NativeTokenCall(ctx);
        f.params = new sc.MutableNativeTokenParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableNativeTokenResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns the data for a given NFT that is on the chain.
    static nftData(ctx: wasmlib.ScViewClientContext): NftDataCall {
        const f = new NftDataCall(ctx);
        f.params = new sc.MutableNftDataParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableNftDataResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    // Returns the balances of all fungible tokens controlled by the chain.
    static totalAssets(ctx: wasmlib.ScViewClientContext): TotalAssetsCall {
        const f = new TotalAssetsCall(ctx);
        f.results = new sc.ImmutableTotalAssetsResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }
}
