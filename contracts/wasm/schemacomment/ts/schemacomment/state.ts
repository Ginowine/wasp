// COPYRIGHT OF A TEST SCHEMA DEFINITION 1
// COPYRIGHT OF A TEST SCHEMA DEFINITION 2

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ImmutableSchemaCommentState extends wasmtypes.ScProxy {
	// comment for TestState1
	testState1(): sc.ArrayOfImmutableInt64 {
		return new sc.ArrayOfImmutableInt64(this.proxy.root(sc.StateTestState1));
	}

	// comment for TestState2
	testState2(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.StateTestState2));
	}
}

export class MutableSchemaCommentState extends wasmtypes.ScProxy {
	asImmutable(): sc.ImmutableSchemaCommentState {
		return new sc.ImmutableSchemaCommentState(this.proxy);
	}

	// comment for TestState1
	testState1(): sc.ArrayOfMutableInt64 {
		return new sc.ArrayOfMutableInt64(this.proxy.root(sc.StateTestState1));
	}

	// comment for TestState2
	testState2(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.StateTestState2));
	}
}