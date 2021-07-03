// Copyright 2021 Linka Cloud  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

var TestFields = struct {
	AString      string
	AnInt        string
	SubMessage   string
	InnerMessage string
	SubOneof     string
	InnerOneof   string
}{
	AString:      "a_string",
	AnInt:        "an_int",
	SubMessage:   "sub_message",
	InnerMessage: "inner_message",
	SubOneof:     "sub_oneof",
	InnerOneof:   "inner_oneof",
}

var SubFields = struct {
	FieldOne string
	FieldTwo string
}{
	FieldOne: "field_one",
	FieldTwo: "field_two",
}

var Test_InnerFields = struct {
	One string
	Two string
}{
	One: "one",
	Two: "two",
}
