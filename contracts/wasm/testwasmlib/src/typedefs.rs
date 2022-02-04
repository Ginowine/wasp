// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use wasmlib::*;
use crate::*;

#[derive(Clone)]
pub struct ArrayOfImmutableString {
	pub(crate) proxy: Proxy,
}

impl ArrayOfImmutableString {
    pub fn length(&self) -> u32 {
        self.proxy.length()
    }

    pub fn get_string(&self, index: u32) -> ScImmutableString {
        ScImmutableString::new(self.proxy.index(index))
    }
}

pub type ImmutableStringArray = ArrayOfImmutableString;

#[derive(Clone)]
pub struct ArrayOfMutableString {
	pub(crate) proxy: Proxy,
}

impl ArrayOfMutableString {
	pub fn append_string(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.append())
	}

	pub fn clear(&self) {
        self.proxy.clear_array();
    }

    pub fn length(&self) -> u32 {
        self.proxy.length()
    }

    pub fn get_string(&self, index: u32) -> ScMutableString {
        ScMutableString::new(self.proxy.index(index))
    }
}

pub type MutableStringArray = ArrayOfMutableString;

#[derive(Clone)]
pub struct MapStringToImmutableString {
	pub(crate) proxy: Proxy,
}

impl MapStringToImmutableString {
    pub fn get_string(&self, key: &str) -> ScImmutableString {
        ScImmutableString::new(self.proxy.key(&string_to_bytes(key)))
    }
}

pub type ImmutableStringMap = MapStringToImmutableString;

#[derive(Clone)]
pub struct MapStringToMutableString {
	pub(crate) proxy: Proxy,
}

impl MapStringToMutableString {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_string(&self, key: &str) -> ScMutableString {
        ScMutableString::new(self.proxy.key(&string_to_bytes(key)))
    }
}

pub type MutableStringMap = MapStringToMutableString;
