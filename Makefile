# Copyright 2021 Linka Cloud  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
