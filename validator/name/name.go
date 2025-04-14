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
	f, err := NameToFunc(nameExpr)
	if err != nil {
		panic(err)
	}
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
func NameToFunc(n *ast.NameExpr) (funcs.Bool, error) {
	typ := n.GetValue()
	switch v := typ.(type) {
	case *ast.Name:
		if v.DoubleValue != nil {
			return funcs.DoubleEq(funcs.DoubleVar(), funcs.DoubleConst(v.GetDoubleValue())), nil
		} else if v.IntValue != nil {
			return funcs.IntEq(funcs.IntVar(), funcs.IntConst(v.GetIntValue())), nil
		} else if v.UintValue != nil {
			return funcs.UintEq(funcs.UintVar(), funcs.UintConst(v.GetUintValue())), nil
		} else if v.BoolValue != nil {
			return funcs.BoolEq(funcs.BoolVar(), funcs.BoolConst(v.GetBoolValue())), nil
		} else if v.StringValue != nil {
			return funcs.StringEq(funcs.StringVar(), funcs.StringConst(v.GetStringValue())), nil
		} else if v.BytesValue != nil {
			return funcs.BytesEq(funcs.BytesVar(), funcs.BytesConst(v.GetBytesValue())), nil
		}
		panic(fmt.Sprintf("unknown name expr name %#v", v))
	case *ast.AnyName:
		return funcs.BoolConst(true), nil
	case *ast.AnyNameExcept:
		n, err := NameToFunc(v.GetExcept())
		if err != nil {
			return nil, err
		}
		return funcs.Not(n), nil
	case *ast.NameChoice:
		l, err := NameToFunc(v.GetLeft())
		if err != nil {
			return nil, err
		}
		r, err := NameToFunc(v.GetRight())
		if err != nil {
			return nil, err
		}
		return funcs.Or(l, r), nil
	case *ast.RegexName:
		return funcs.Regex(funcs.StringConst(v.Pattern), funcs.StringVar())
	}
	panic(fmt.Sprintf("unknown name expr typ %T", typ))
}
