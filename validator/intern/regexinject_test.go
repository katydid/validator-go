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
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/katydid/parser-go-json/json"
	"github.com/katydid/validator-go/validator/ast"
	"github.com/katydid/validator-go/validator/funcs"
	"github.com/katydid/validator-go/validator/intern"
)

// Test whether we can inject our own regex engine for the name expression.
func TestInjectRegexName(t *testing.T) {
	runmatch := func() bool {
		p := json.NewParser()
		p.Init([]byte(`{"abc": "def"}`))
		pattern := ast.NewTreeNode(ast.NewRegexName("a.*"), ast.NewZAny())
		g := ast.NewGrammar(map[string]*ast.Pattern{"main": pattern})
		match, err := intern.Interpret(g, true, p)
		if err != nil {
			t.Fatal(err)
		}
		return match
	}
	if !runmatch() {
		t.Fatalf("expected initial match")
	}

	// try to override regex engine
	funcs.Register("regex", WrongRegex)
	defer funcs.Register("regex", funcs.Regex)
	if runmatch() {
		t.Fatalf("expected not match with wrong regex engine")
	}

	funcs.Register("regex", funcs.Regex)
	if !runmatch() {
		t.Fatalf("expected match again")
	}
}

// Test whether we can inject our own regex engine for leaves.
func TestInjectRegexLeaf(t *testing.T) {
	runmatch := func() bool {
		p := json.NewParser()
		p.Init([]byte(`{"abc": "def"}`))
		pattern := ast.NewTreeNode(ast.NewStringName("abc"), ast.NewLeafNode(ast.NewFunction("regex", ast.NewStringConst("d.*"), ast.NewStringVar())))
		g := ast.NewGrammar(map[string]*ast.Pattern{"main": pattern})
		match, err := intern.Interpret(g, true, p)
		if err != nil {
			t.Fatal(err)
		}
		return match
	}
	if !runmatch() {
		t.Fatalf("expected initial match")
	}

	// try to override regex engine
	funcs.Register("regex", WrongRegex)
	defer funcs.Register("regex", funcs.Regex)
	if runmatch() {
		t.Fatalf("expected not match with wrong regex engine")
	}

	funcs.Register("regex", funcs.Regex)
	if !runmatch() {
		t.Fatalf("expected match again")
	}
}

// WrongRegex always returns the opposite of the match of the normal regex func.
func WrongRegex(pattern funcs.ConstString, input funcs.String) (funcs.Bool, error) {
	if pattern.HasVariable() {
		return nil, fmt.Errorf("regex requires a constant expression as its first parameter, but it has a variable parameter")
	}
	p, err := pattern.Eval()
	if err != nil {
		return nil, err
	}
	r, err := regexp.Compile(p)
	if err != nil {
		return nil, err
	}
	return funcs.TrimBool(&wrongRegex{
		pattern:     p,
		r:           r,
		S:           input,
		hash:        funcs.Hash("regex", pattern, input),
		hasVariable: input.HasVariable(),
	}), nil
}

type wrongRegex struct {
	pattern     string
	r           *regexp.Regexp
	S           funcs.String
	hash        uint64
	hasVariable bool
}

func (this *wrongRegex) HasVariable() bool {
	return this.hasVariable
}

func (this *wrongRegex) ToExpr() *ast.Expr {
	return ast.NewFunction("regex", ast.NewStringConst(this.pattern), this.S.ToExpr())
}

func (this *wrongRegex) Eval() (bool, error) {
	s, err := this.S.Eval()
	if err != nil {
		return false, err
	}
	// Always return the opposite of the match
	return !this.r.MatchString(s), nil
}

func (this *wrongRegex) Compare(that funcs.Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*wrongRegex); ok {
		if c := strings.Compare(this.pattern, other.pattern); c != 0 {
			return c
		}
		if c := this.S.Compare(other.S); c != 0 {
			return c
		}
		return 0
	}
	return this.ToExpr().Compare(that.ToExpr())
}

func (this *wrongRegex) Hash() uint64 {
	return this.hash
}
