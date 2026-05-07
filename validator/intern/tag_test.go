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

func TestTag(t *testing.T) {
	p := json.NewJSONSchemaParser()

	grammarParser := parser.NewParser()
	g, err := grammarParser.ParseGrammar("tag(object):*")
	if err != nil {
		t.Fatal(err)
	}

	p.Init([]byte(`{"abc": "def"}`))
	match, err := intern.Interpret(g, true, p)
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatal("expected match")
	}

	p.Init([]byte(`["abc", "def"]`))
	match, err = intern.Interpret(g, true, p)
	if err != nil {
		t.Fatal(err)
	}
	if match {
		t.Fatal("expected not match")
	}

	g, err = grammarParser.ParseGrammar("tag(array):*")
	if err != nil {
		t.Fatal(err)
	}

	p.Init([]byte(`{"abc": "def"}`))
	match, err = intern.Interpret(g, true, p)
	if err != nil {
		t.Fatal(err)
	}
	if match {
		t.Fatal("expected not match")
	}

	p.Init([]byte(`["abc", "def"]`))
	match, err = intern.Interpret(g, true, p)
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatal("expected match")
	}

	g, err = grammarParser.ParseGrammar("tag(object):a::$string")
	if err != nil {
		t.Fatal(err)
	}
	p.Init([]byte(`{"a": "b"}`))
	match, err = intern.Interpret(g, true, p)
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatal("expected match")
	}
	p.Init([]byte(`{"a": []}`))
	match, err = intern.Interpret(g, true, p)
	if err != nil {
		t.Fatal(err)
	}
	if match {
		t.Fatal("expected not match")
	}

	g, err = grammarParser.ParseGrammar("tag(object):a:!((tag(object)|tag(array))):<empty>")
	if err != nil {
		t.Fatal(err)
	}
	p.Init([]byte(`{"a": []}`))
	match, err = intern.Interpret(g, true, p)
	if err != nil {
		t.Fatal(err)
	}
	if match {
		t.Fatal("expected not match")
	}

	g, err = grammarParser.ParseGrammar("tag(object):a:tag(array):<empty>")
	if err != nil {
		t.Fatal(err)
	}
	p.Init([]byte(`{"a": []}`))
	match, err = intern.Interpret(g, true, p)
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatal("expected match")
	}
}
