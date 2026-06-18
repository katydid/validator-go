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

package intern_test

import (
	"testing"

	"github.com/katydid/parser-go-json/json"
	"github.com/katydid/validator-go/validator/intern"
	"github.com/katydid/validator-go/validator/parser"
)

func TestXor(t *testing.T) {
	p := json.NewParser()

	grammarParser := parser.NewParser()
	g, err := grammarParser.ParseGrammar("abc:(==1 ^ ==2)")
	if err != nil {
		t.Fatal(err)
	}

	p.Init([]byte(`{"abc": 1}`))
	match, err := intern.Interpret(g, true, p)
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatal("expected match")
	}

	g, err = grammarParser.ParseGrammar("abc:(>1 ^ <5)")
	if err != nil {
		t.Fatal(err)
	}

	p.Init([]byte(`{"abc": 0}`))
	match, err = intern.Interpret(g, true, p)
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatal("expected match")
	}

	p.Init([]byte(`{"abc": 3}`))
	match, err = intern.Interpret(g, true, p)
	if err != nil {
		t.Fatal(err)
	}
	if match {
		t.Fatal("unexpected match")
	}

	p.Init([]byte(`{"abc": 7}`))
	match, err = intern.Interpret(g, true, p)
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatal("expected match")
	}
}
