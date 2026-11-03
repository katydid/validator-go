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

package intern_test

import (
	"testing"

	"katydid.org.za/go/parser-go/hedge"
	"katydid.org.za/go/validator-go/validator/ast"
	"katydid.org.za/go/validator-go/validator/intern"
)

func TestBytes(t *testing.T) {
	p := hedge.NewParser(hedge.Hedge{hedge.NewBytesNode([]byte("abc"))})
	g := ast.NewGrammar(map[string]*ast.Pattern{"main": ast.NewLeafNode(ast.NewType(ast.NewBytesVar()))})
	m, err := intern.Interpret(g, false, p)
	if err != nil {
		t.Fatal(err)
	}
	if !m {
		t.Fatal("expected match")
	}
}
