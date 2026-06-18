// Copyright 2026 Walter Schulze
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auto

import (
	"testing"

	"github.com/katydid/parser-go-json/json"
	"github.com/katydid/parser-go/parse"
	"github.com/katydid/parser-go/parse/debug"
	"github.com/katydid/validator-go/validator/parser"
)

// This is a test of a translated jsonschema, that was failing with auto, but not with intern or mem.
func TestDependent(t *testing.T) {
	translated :=
		`(
		((.ExtendedAddress:*&.StreetAddress:*)|(!(ExtendedAddress):*)*)
		&{
			((CountryName::$string|ExtendedAddress::$string|StreetAddress::$string))*
			;(!((CountryName|ExtendedAddress|StreetAddress)):*)*
		}
	)`

	failInput := `{"ExtendedAddress":"a"}`

	g, err := parser.ParseGrammar(translated)
	if err != nil {
		t.Fatal(err)
	}
	a, err := Compile(g, WithRecordOpts())
	if err != nil {
		t.Fatal(err)
	}
	var p parse.ParserWithInit = json.NewParser()
	p = debug.NewLogger(p, debug.NewLineLogger())

	p.Init([]byte(failInput))
	m, err := a.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
	if m {
		t.Error("expected fail")
	}

	passInput := `{"ExtendedAddress":"a", "StreetAddress": "b"}`
	p.Init([]byte(passInput))
	m, err = a.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
	if !m {
		t.Error("expected pass")
	}
}
