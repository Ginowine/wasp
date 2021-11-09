// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class ImmutableTestCoreState extends wasmlib.ScMapID {
    counter(): wasmlib.ScImmutableInt64 {
		return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxStateCounter]);
	}

    hnameEP(): wasmlib.ScImmutableHname {
		return new wasmlib.ScImmutableHname(this.mapID, sc.idxMap[sc.IdxStateHnameEP]);
	}

    ints(): sc.MapStringToImmutableInt64 {
		let mapID = wasmlib.getObjectID(this.mapID, sc.idxMap[sc.IdxStateInts], wasmlib.TYPE_MAP);
		return new sc.MapStringToImmutableInt64(mapID);
	}

    mintedColor(): wasmlib.ScImmutableColor {
		return new wasmlib.ScImmutableColor(this.mapID, sc.idxMap[sc.IdxStateMintedColor]);
	}

    mintedSupply(): wasmlib.ScImmutableInt64 {
		return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxStateMintedSupply]);
	}
}

export class MutableTestCoreState extends wasmlib.ScMapID {
    counter(): wasmlib.ScMutableInt64 {
		return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxStateCounter]);
	}

    hnameEP(): wasmlib.ScMutableHname {
		return new wasmlib.ScMutableHname(this.mapID, sc.idxMap[sc.IdxStateHnameEP]);
	}

    ints(): sc.MapStringToMutableInt64 {
		let mapID = wasmlib.getObjectID(this.mapID, sc.idxMap[sc.IdxStateInts], wasmlib.TYPE_MAP);
		return new sc.MapStringToMutableInt64(mapID);
	}

    mintedColor(): wasmlib.ScMutableColor {
		return new wasmlib.ScMutableColor(this.mapID, sc.idxMap[sc.IdxStateMintedColor]);
	}

    mintedSupply(): wasmlib.ScMutableInt64 {
		return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxStateMintedSupply]);
	}
}
