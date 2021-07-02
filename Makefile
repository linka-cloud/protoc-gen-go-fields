
.PHONY: install
install:
	@go install .

.PHONY: gen-tests
gen-tests:
	@protoc -I. --debug_out="tests:." tests/pb/test.proto

PROTO_OPTS = paths=source_relative

.PHONY: gen-example
gen-example: install
	@protoc -I. --go_out=$(PROTO_OPTS):. --go-fields_out=$(PROTO_OPTS):. tests/pb/test.proto
