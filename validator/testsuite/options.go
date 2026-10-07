//  Copyright 2026 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package testsuite

import (
	"os"
	"path/filepath"
)

type options struct {
	path   string
	must   bool
	codecs map[string]struct{}
}

var testpath string
var benchpath string

func init() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = "../../../../../../"
	}
	testpath = filepath.Join(gopath, "src/katydid.org.za/go/validator-testsuite/validator/tests")
	benchpath = filepath.Join(gopath, "src/katydid.org.za/go/validator-testsuite/validator/benches")
}

func newDefaultTestOptions() *options {
	return &options{
		path: testpath,
		must: os.Getenv("TESTSUITE") == "MUST",
		codecs: map[string]struct{}{
			"hedge":     {},
			"json":      {},
			"xml":       {},
			"goreflect": {},
		},
	}
}

func newTestOptions(opts []Option) *options {
	o := newDefaultTestOptions()
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func newDefaultBenchOptions() *options {
	return &options{
		path: benchpath,
		must: os.Getenv("TESTSUITE") == "MUST",
		codecs: map[string]struct{}{
			"hedge":     {},
			"json":      {},
			"xml":       {},
			"goreflect": {},
		},
	}
}

func newBenchOptions(opts []Option) *options {
	o := newDefaultBenchOptions()
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type Option func(o *options)

func WithPath(path string) Option {
	return func(o *options) {
		o.path = path
	}
}

func Must() Option {
	return func(o *options) {
		o.must = true
	}
}

func OnlyHedgeCodec() Option {
	return func(o *options) {
		o.codecs = map[string]struct{}{
			"hedge": {},
		}
	}
}
