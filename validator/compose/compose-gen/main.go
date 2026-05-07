//  Copyright 2013 Walter Schulze
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

// Command compose-gen generates some of the code in the compose package.
package main

import (
	"github.com/katydid/validator-go/gen"
)

const composeStr = `
func compose{{.SingleName}}(expr *ast.Expr) (funcs.{{.FuncSingleName}}, error) {
	f, err := prep(expr, types.{{.SingleType}})
	if err != nil {
		return nil, err
	}
	if expr.Terminal != nil {
		if expr.GetTerminal().Variable != nil {
			return funcs.{{.SingleName}}Var(), nil
		} else {
			return funcs.{{.FuncSingleName}}Const(expr.GetTerminal().Get{{.SingleValue}}Value()), nil
		}
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return f.New{{.FuncSingleName}}(values...)
}

func compose{{.ListName}}(expr *ast.Expr) (funcs.{{.FuncListName}}, error) {
	f, err := prep(expr, types.{{.ListType}})
	if err != nil {
		return nil, err
	}
	if expr.List != nil {
		vs, err := newValues(expr.GetList().GetElems())
		if err != nil {
			return nil, err
		}
		bs := make([]funcs.{{.FuncSingleName}}, len(vs))
		var ok bool
		for i := range vs {
			bs[i], ok = vs[i].(funcs.{{.FuncSingleName}})
			if !ok {
				return nil, &errExpected{types.{{.SingleType}}.String(), expr.String()}
			}
		}
		return funcs.NewListOf{{.FuncSingleName}}(bs), nil
	}
	values, err := newValues(expr.GetFunction().GetParams())
	if err != nil {
		return nil, err
	}
	return f.New{{.FuncListName}}(values...)
}
`

type composer struct {
	SingleName     string
	FuncSingleName string
	SingleType     string
	SingleValue    string
	ListName       string
	FuncListName   string
	ListType       string
}

func main() {
	gen := gen.NewPackage("compose")
	gen(composeStr, "compose.gen.go", []interface{}{
		&composer{"Double", "Double", "SINGLE_DOUBLE", "Double", "Doubles", "Doubles", "LIST_DOUBLE"},
		&composer{"Int", "Int", "SINGLE_INT", "Int", "Ints", "Ints", "LIST_INT"},
		&composer{"Uint", "Uint", "SINGLE_UINT", "Uint", "Uints", "Uints", "LIST_UINT"},
		&composer{"Bool", "Bool", "SINGLE_BOOL", "Bool", "Bools", "Bools", "LIST_BOOL"},
		&composer{"String", "String", "SINGLE_STRING", "String", "Strings", "Strings", "LIST_STRING"},
		&composer{"Bytes", "Bytes", "SINGLE_BYTES", "Bytes", "ListOfBytes", "ListOfBytes", "LIST_BYTES"},
		&composer{"Tag", "String", "SINGLE_TAG", "Tag", "Tags", "Strings", "LIST_TAG"},
	},
		`"github.com/katydid/validator-go/validator/ast"`,
		`"github.com/katydid/validator-go/validator/funcs"`,
		`"github.com/katydid/validator-go/validator/types"`)
}
