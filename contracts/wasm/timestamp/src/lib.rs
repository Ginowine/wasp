// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use timestamp::*;
use wasmlib::*;

use crate::consts::*;
use crate::results::*;

mod consts;
mod contract;
mod results;

mod timestamp;

const EXPORT_MAP: ScExportMap = ScExportMap {
    names: &[
    	FUNC_NOW,
    	VIEW_GET_TIMESTAMP,
	],
    funcs: &[
    	func_now_thunk,
	],
    views: &[
    	view_get_timestamp_thunk,
	],
};

#[no_mangle]
fn on_call(index: i32) {
	ScExports::call(index, &EXPORT_MAP);
}

#[no_mangle]
fn on_load() {
    ScExports::export(&EXPORT_MAP);
}

pub struct NowContext {
}

fn func_now_thunk(ctx: &ScFuncContext) {
	ctx.log("timestamp.funcNow");
	let f = NowContext {
	};
	func_now(ctx, &f);
	ctx.log("timestamp.funcNow ok");
}

pub struct GetTimestampContext {
	results: MutableGetTimestampResults,
}

fn view_get_timestamp_thunk(ctx: &ScViewContext) {
	ctx.log("timestamp.viewGetTimestamp");
	let f = GetTimestampContext {
		results: MutableGetTimestampResults { proxy: results_proxy() },
	};
	view_get_timestamp(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("timestamp.viewGetTimestamp ok");
}
