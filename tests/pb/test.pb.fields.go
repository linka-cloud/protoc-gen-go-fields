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
