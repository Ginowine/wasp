package rstemplates

var libRs = map[string]string{
	// *******************************
	"lib.rs": `
#![allow(dead_code)]
#![allow(unused_imports)]

use $package::*;
use wasmlib::*;

use crate::consts::*;
$#if events useEvents
$#if params useParams
$#if results useResults
use crate::state::*;
$#if structs useStructs
$#if typedefs useTypedefs

mod consts;
mod contract;
$#if events modEvents
$#if params modParams
$#if results modResults
mod state;
$#if structs modStructs
$#if typedefs modTypeDefs
mod $package;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
$#each func libExportFunc
}
$#each func libThunk
`,
	// *******************************
	"libExportFunc": `
    exports.add_$kind($KIND$+_$FUNC_NAME,$func_pad $kind$+_$func_name$+_thunk);
`,
	// *******************************
	"libThunk": `

pub struct $FuncName$+Context {
$#if func PackageEvents
$#if param ImmutableFuncNameParams
$#if result MutableFuncNameResults
$#if func MutablePackageState
$#if view ImmutablePackageState
}

fn $kind$+_$func_name$+_thunk(ctx: &Sc$Kind$+Context) {
	ctx.log("$package.$kind$FuncName");
	let f = $FuncName$+Context {
$#if func PackageEventsInit
$#if param ImmutableFuncNameParamsInit
$#if result MutableFuncNameResultsInit
$#if func MutablePackageStateInit
$#if view ImmutablePackageStateInit
	};
$#emit accessCheck
$#each mandatory requireMandatory
	$kind$+_$func_name(ctx, &f);
$#if result sendResults
	ctx.log("$package.$kind$FuncName ok");
}
`,
	// *******************************
	"sendResults": `
	ctx.results(&f.results.proxy.kv_store);
`,
	// *******************************
	"PackageEvents": `
$#if events PackageEventsExist
`,
	// *******************************
	"PackageEventsExist": `
	events:  $Package$+Events,
`,
	// *******************************
	"PackageEventsInit": `
$#if events PackageEventsInitExist
`,
	// *******************************
	"PackageEventsInitExist": `
		events:  $Package$+Events {},
`,
	// *******************************
	"ImmutableFuncNameParams": `
	params: Immutable$FuncName$+Params,
`,
	// *******************************
	"ImmutableFuncNameParamsInit": `
		params: Immutable$FuncName$+Params { proxy: params_proxy() },
`,
	// *******************************
	"MutableFuncNameResults": `
	results: Mutable$FuncName$+Results,
`,
	// *******************************
	"MutableFuncNameResultsInit": `
		results: Mutable$FuncName$+Results { proxy: results_proxy() },
`,
	// *******************************
	"MutablePackageState": `
	state: Mutable$Package$+State,
`,
	// *******************************
	"MutablePackageStateInit": `
		state: Mutable$Package$+State { proxy: state_proxy() },
`,
	// *******************************
	"ImmutablePackageState": `
	state: Immutable$Package$+State,
`,
	// *******************************
	"ImmutablePackageStateInit": `
		state: Immutable$Package$+State { proxy: state_proxy() },
`,
	// *******************************
	"requireMandatory": `
	ctx.require(f.params.$fld_name().exists(), "missing mandatory $fldName");
`,
	// *******************************
	"accessCheck": `
$#set accessFinalize accessOther
$#emit caseAccess$funcAccess
$#emit $accessFinalize
`,
	// *******************************
	"caseAccess": `
$#set accessFinalize accessDone
`,
	// *******************************
	"caseAccessself": `
$#if funcAccessComment accessComment
	ctx.require(ctx.caller() == ctx.account_id(), "no permission");

$#set accessFinalize accessDone
`,
	// *******************************
	"caseAccesschain": `
$#if funcAccessComment accessComment
	ctx.require(ctx.caller() == ctx.chain_owner_id(), "no permission");

$#set accessFinalize accessDone
`,
	// *******************************
	"caseAccesscreator": `
$#if funcAccessComment accessComment
	ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

$#set accessFinalize accessDone
`,
	// *******************************
	"accessOther": `
$#if funcAccessComment accessComment
	let access = f.state.$func_access();
	ctx.require(access.exists(), "access not set: $funcAccess");
	ctx.require(ctx.caller() == access.value(), "no permission");

`,
	// *******************************
	"accessDone": `
`,
	// *******************************
	"accessComment": `

	$funcAccessComment
`,
}
