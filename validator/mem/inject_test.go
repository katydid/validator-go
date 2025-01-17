//  Copyright 2016 Walter Schulze
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

package mem_test

import (
	"testing"

	jsonparser "github.com/katydid/parser-go-json/json"
	"github.com/katydid/validator-go/validator/ast"
	. "github.com/katydid/validator-go/validator/combinator"
	"github.com/katydid/validator-go/validator/funcs"
	"github.com/katydid/validator-go/validator/mem"
)

func newInjectable() *injectableInt {
	return &injectableInt{}
}

type injectableInt struct {
	context *funcs.Context
}

func (this *injectableInt) Eval() (int64, error) {
	v := this.context.Value.(int64)
	return v, nil
}

func (this *injectableInt) Compare(that funcs.Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if _, ok := that.(*injectableInt); ok {
		return 0
	}
	return this.ToExpr().Compare(that.ToExpr())
}

func (this *injectableInt) String() string {
	return "inject()"
}

func (this *injectableInt) ToExpr() *ast.Expr {
	return ast.NewFunction("inject")
}

func (this *injectableInt) Hash() uint64 {
	return 17
}

func (this *injectableInt) SetContext(context *funcs.Context) {
	this.context = context
}

func (this *injectableInt) HasVariable() bool {
	return true
}

func init() {
	funcs.Register("inject", newInjectable)
}

var injectNumber = G(map[string]*ast.Pattern{"main": In("Num", Value(Eq(IntVar(), ast.NewFunction("inject"))))})

type Number struct {
	Num int64
}

func testInject(t *testing.T, m *mem.Mem) bool {
	parser := jsonparser.NewJsonParser()
	if err := parser.Init([]byte(`{"Num": 456}`)); err != nil {
		t.Fatal(err)
	}
	res, err := m.Validate(parser)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func TestInject(t *testing.T) {
	grammar := injectNumber.Grammar()
	t.Logf("parsed Grammar: %s\n", grammar)
	m, err := mem.New(grammar)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("trying to set context...\n")
	c := &funcs.Context{Value: int64(0)}
	m.SetContext(c)
	t.Logf("hopefully context was set\n")
	c.Value = int64(456)
	if !testInject(t, m) {
		t.Fatalf("expected match")
	}
	c.Value = int64(123)
	if testInject(t, m) {
		t.Fatalf("expected non match")
	}
}
