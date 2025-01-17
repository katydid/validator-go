//  Copyright 2015 Walter Schulze
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

package compose_test

import (
	"testing"

	"github.com/katydid/parser-go/parser/debug"
	"github.com/katydid/validator-go/validator/ast"
	. "github.com/katydid/validator-go/validator/combinator"
	"github.com/katydid/validator-go/validator/compose"
)

func TestNoEqualError(t *testing.T) {
	e := Eq(ast.NewFunction("elem", ast.NewStringList(StringVar()), IntConst(3)), StringConst("0123456789"))
	b, err := compose.NewBool(e)
	if err != nil {
		t.Fatal(err)
	}
	f, err := compose.NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	if v, err := f.Eval(debug.NewStringValue("0")); err != nil {
		t.Fatal(err)
	} else if v {
		t.Fatalf("expected false")
	}
}
