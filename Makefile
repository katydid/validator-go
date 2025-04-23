# Copyright 2013 Walter Schulze
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: nuke dep regenerate gofmt build test

all: nuke dep regenerate build test vet

dep:
	go install -v github.com/goccmack/gocc
	go install -v github.com/awalterschulze/goderive

test:
	go clean -testcache
	TESTSUITE=MUST go test ./...

build:
	go build ./...

install:
	go install ./...

bench:
	TESTSUITE=MUST go test -test.v -test.run=XXX -test.bench=. ./...

vet:
	go vet ./gen/...
	go vet ./validator/...

regenerate:
	goderive ./...
	(cd validator && make regenerate)
	(cd validator/funcs && go test -test.run=GenFuncList | grep "func\ " >../../list_of_functions.txt)

clean:
	go clean ./...
	(cd validator && make clean)

nuke: clean
	(cd validator && make nuke)
	rm list_of_functions.txt || true
	go clean -i ./...

gofmt:
	gofmt -l -s -w .

errcheck:
	go get github.com/kisielk/errcheck
	errcheck -ignore 'fmt:[FS]?[Pp]rint*' ./...

diff:
	git diff --exit-code .
