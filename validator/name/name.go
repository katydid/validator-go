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

// Package name has functions for a parsed name expression which include compilation and evaluation.
package name

import (
	"fmt"

	"github.com/katydid/parser-go/parser"
	"github.com/katydid/validator-go/validator/ast"
	"github.com/katydid/validator-go/validator/compose"
	"github.com/katydid/validator-go/validator/funcs"
)

// EvalName evaluates a name expression given a name value.
func EvalName(nameExpr *ast.NameExpr, name parser.Value) bool {
	f := NameToFunc(nameExpr)
	b, err := compose.NewBoolFunc(f)
	if err != nil {
		panic(err)
	}
	v, err := b.Eval(name)
	if err != nil {
		panic(err)
	}
	return v
}

// NameToFunc compiles a parsed name expression into a function.
func NameToFunc(n *ast.NameExpr) funcs.Bool {
	typ := n.GetValue()
	switch v := typ.(type) {
	case *ast.Name:
		if v.DoubleValue != nil {
			return funcs.DoubleEq(funcs.DoubleVar(), funcs.DoubleConst(v.GetDoubleValue()))
		} else if v.IntValue != nil {
			return funcs.IntEq(funcs.IntVar(), funcs.IntConst(v.GetIntValue()))
		} else if v.UintValue != nil {
			return funcs.UintEq(funcs.UintVar(), funcs.UintConst(v.GetUintValue()))
		} else if v.BoolValue != nil {
			return funcs.BoolEq(funcs.BoolVar(), funcs.BoolConst(v.GetBoolValue()))
		} else if v.StringValue != nil {
			return funcs.StringEq(funcs.StringVar(), funcs.StringConst(v.GetStringValue()))
		} else if v.BytesValue != nil {
			return funcs.BytesEq(funcs.BytesVar(), funcs.BytesConst(v.GetBytesValue()))
		}
		panic(fmt.Sprintf("unknown name expr name %#v", v))
	case *ast.AnyName:
		return funcs.BoolConst(true)
	case *ast.AnyNameExcept:
		return funcs.Not(NameToFunc(v.GetExcept()))
	case *ast.NameChoice:
		return funcs.Or(NameToFunc(v.GetLeft()), NameToFunc(v.GetRight()))
	}
	panic(fmt.Sprintf("unknown name expr typ %T", typ))
}
