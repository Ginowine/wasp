// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use wasmlib::*;
use crate::*;

pub struct DonateCall {
	pub func: ScFunc,
	pub params: MutableDonateParams,
}

pub struct WithdrawCall {
	pub func: ScFunc,
	pub params: MutableWithdrawParams,
}

pub struct DonationCall {
	pub func: ScView,
	pub params: MutableDonationParams,
	pub results: ImmutableDonationResults,
}

pub struct DonationInfoCall {
	pub func: ScView,
	pub results: ImmutableDonationInfoResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn donate(_ctx: &dyn ScFuncCallContext) -> DonateCall {
        let mut f = DonateCall {
            func: ScFunc::new(HSC_NAME, HFUNC_DONATE),
            params: MutableDonateParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn withdraw(_ctx: &dyn ScFuncCallContext) -> WithdrawCall {
        let mut f = WithdrawCall {
            func: ScFunc::new(HSC_NAME, HFUNC_WITHDRAW),
            params: MutableWithdrawParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn donation(_ctx: &dyn ScViewCallContext) -> DonationCall {
        let mut f = DonationCall {
            func: ScView::new(HSC_NAME, HVIEW_DONATION),
            params: MutableDonationParams { proxy: Proxy::nil() },
            results: ImmutableDonationResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn donation_info(_ctx: &dyn ScViewCallContext) -> DonationInfoCall {
        let mut f = DonationInfoCall {
            func: ScView::new(HSC_NAME, HVIEW_DONATION_INFO),
            results: ImmutableDonationInfoResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }
}
