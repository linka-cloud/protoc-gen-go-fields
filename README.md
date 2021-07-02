# protoc-gen-go-fields

*This project is currently in **alpha**. The API should be considered unstable and likely to change*

**protoc-gen-go-fields** is a protoc plugin to generate golang structs containing fields protobuf names. 

NB: Its original purpose was only to test [protoc-gen-star](github.com/lyft/protoc-gen-star).


## Overview

Protobuf definition:

```proto
syntax = "proto3";

package go.fields.test;

option go_package = "go.linka.cloud/protoc-gen-go-fields/tests/pb/test";


message Test {
  message Inner {
    string one = 1;
    string two = 2;
  }
  string a_string = 1;
  int64 an_int = 2;
  Sub sub_message = 3;
  Inner inner_message = 4;

  oneof oneof {
    Sub sub_oneof = 5;
    Inner inner_oneof = 6;
  }
}

message Sub {
  string field_one = 1;
  int32 field_two = 2;
}

```

Generated Fields structs:
```go
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

```
