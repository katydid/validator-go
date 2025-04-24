//  Copyright 2017 Walter Schulze
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

package intern

import (
	"fmt"

	"github.com/katydid/validator-go/validator/ast"
	"github.com/katydid/validator-go/validator/sets"
)

// Nullable returns whether the input Pattern p also matches the empty string.
// This is a naive implementation and it does not handle left recursion.
func Nullable(refs ast.RefLookup, p *ast.Pattern) (bool, error) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty:
		return true, nil
	case *ast.TreeNode:
		return false, nil
	case *ast.LeafNode:
		return false, nil
	case *ast.Concat:
		l, err := Nullable(refs, v.GetLeftPattern())
		if err != nil {
			return false, err
		}
		r, err := Nullable(refs, v.GetRightPattern())
		if err != nil {
			return false, err
		}
		return l && r, nil
	case *ast.Or:
		l, err := Nullable(refs, v.GetLeftPattern())
		if err != nil {
			return false, err
		}
		r, err := Nullable(refs, v.GetRightPattern())
		if err != nil {
			return false, err
		}
		return l || r, nil
	case *ast.And:
		l, err := Nullable(refs, v.GetLeftPattern())
		if err != nil {
			return false, err
		}
		r, err := Nullable(refs, v.GetRightPattern())
		if err != nil {
			return false, err
		}
		return l && r, nil
	case *ast.ZeroOrMore:
		return true, nil
	case *ast.Reference:
		return Nullable(refs, refs[v.GetName()])
	case *ast.Not:
		n, err := Nullable(refs, v.GetPattern())
		if err != nil {
			return false, err
		}
		return !n, nil
	case *ast.ZAny:
		return true, nil
	case *ast.Contains:
		return Nullable(refs, v.GetPattern())
	case *ast.Optional:
		return true, nil
	case *ast.Interleave:
		l, err := Nullable(refs, v.GetLeftPattern())
		if err != nil {
			return false, err
		}
		r, err := Nullable(refs, v.GetRightPattern())
		if err != nil {
			return false, err
		}
		return l && r, nil
	case *ast.Extension:
		children := getExtensionsFromAST(p, true)
		nullables, err := traverse(children, func(child *ast.Pattern) (bool, error) { return Nullable(refs, child) })
		if err != nil {
			return false, err
		}
		ext, err := GetExtension(v.Name)
		if err != nil {
			return false, err
		}
		return ext.nullable(nullables), nil
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func anyNullable(ps []*Pattern) bool {
	for _, any := range ps {
		if any.nullable {
			return true
		}
	}
	return false
}

func allNullable(ps []*Pattern) bool {
	for _, all := range ps {
		if !all.nullable {
			return false
		}
	}
	return true
}

func newNullableBits(patterns []*Pattern) sets.Bits {
	nulls := sets.NewBits(len(patterns))
	for i, p := range patterns {
		nulls.Set(i, p.nullable)
	}
	return nulls
}

func nullables(ps []*Pattern) []bool {
	nulls := make([]bool, len(ps))
	for i, p := range ps {
		nulls[i] = p.nullable
	}
	return nulls
}
