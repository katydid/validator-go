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
	"io"

	"github.com/katydid/parser-go/parser"
	"github.com/katydid/validator-go/validator/ast"
)

// Interpret interprets the grammar given the parser and returns whether the parser is valid given the grammar.
// Interpret uses derivatives and simplification to recusively derive the resulting grammar.
// This resulting grammar's nullability then represents the result of the function.
// This implementation does not handle immediate recursion, see the HasRecursion function.
func Interpret(g *ast.Grammar, record bool, parser parser.Interface) (bool, error) {
	c := NewConstructor()
	if record {
		c = NewConstructorOptimizedForRecords()
	}
	main, err := c.AddGrammar(g)
	if err != nil {
		return false, err
	}
	return derivs(c, main, parser)
}

func derivs(c Construct, main *Pattern, parser parser.Interface) (bool, error) {
	finals, err := deriv(c, []*Pattern{main}, parser)
	if err != nil {
		return false, err
	}
	return finals[0].nullable, nil
}

func deriv(c Construct, patterns []*Pattern, tree parser.Interface) ([]*Pattern, error) {
	var resPatterns []*Pattern = patterns
	for {
		if !Escapable(resPatterns) {
			return resPatterns, nil
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		ifs := DeriveCalls(c, resPatterns)
		childPatterns, err := evalIfExprs(ifs, tree)
		if err != nil {
			return nil, err
		}
		if tree.IsLeaf() {
			//do nothing
		} else {
			tree.Down()
			z, _ := Zip(childPatterns)
			z.Patterns, err = deriv(c, z.Patterns, tree)
			if err != nil {
				return nil, err
			}
			childPatterns = z.Unzip()
			tree.Up()
		}
		nulls := nullables(childPatterns)
		resPatterns, err = DeriveReturns(c, resPatterns, nulls)
		if err != nil {
			return nil, err
		}
	}
	return resPatterns, nil
}

func DeriveCalls(construct Construct, patterns []*Pattern) []*IfExpr {
	res := []*IfExpr{}
	for _, pattern := range patterns {
		cs := derivCall(construct, pattern)
		res = append(res, cs...)
	}
	return res
}

func derivCall(c Construct, p *Pattern) []*IfExpr {
	switch p.Type {
	case Empty:
		return []*IfExpr{}
	case ZAny:
		return []*IfExpr{}
	case Node:
		return []*IfExpr{{p.Func, p.Patterns[0], c.NewNotZAny()}}
	case Concat:
		res := []*IfExpr{}
		for i := range p.Patterns {
			ps := derivCall(c, p.Patterns[i])
			res = append(res, ps...)
			if !p.Patterns[i].Nullable() {
				break
			}
		}
		return res
	case Or:
		return derivCallAll(c, p.Patterns)
	case And:
		return derivCallAll(c, p.Patterns)
	case Interleave:
		return derivCallAll(c, p.Patterns)
	case ZeroOrMore:
		return derivCall(c, p.Patterns[0])
	case Reference:
		return derivCall(c, c.Deref(p.Name))
	case Not:
		return derivCall(c, p.Patterns[0])
	case Contains:
		return derivCall(c, p.Patterns[0])
	case Optional:
		return derivCall(c, p.Patterns[0])
	case Extension:
		return derivCallAll(c, p.Patterns)
	}
	panic(fmt.Sprintf("unknown pattern typ %d", p.Type))
}

func derivCallAll(c Construct, ps []*Pattern) []*IfExpr {
	res := []*IfExpr{}
	for i := range ps {
		pss := derivCall(c, ps[i])
		res = append(res, pss...)
	}
	return res
}

func DeriveReturns(c Construct, originals []*Pattern, evaluated []bool) ([]*Pattern, error) {
	res := make([]*Pattern, len(originals))
	rest := evaluated
	var err error
	for i, original := range originals {
		res[i], rest, err = derivReturn(c, original, rest)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func derivReturn(c Construct, p *Pattern, nulls []bool) (*Pattern, []bool, error) {
	rest := nulls
	var dR = func(pattern *Pattern) (*Pattern, error) {
		var res *Pattern
		var err error
		res, rest, err = derivReturn(c, pattern, rest)
		return res, err
	}
	switch p.Type {
	case Empty:
		return c.NewNotZAny(), nulls, nil
	case ZAny:
		return c.NewZAny(), nulls, nil
	case Node:
		if nulls[0] {
			return c.NewEmpty(), nulls[1:], nil
		}
		return c.NewNotZAny(), nulls[1:], nil
	case Concat:
		orPatterns := make([]*Pattern, 0, len(p.Patterns))
		var err error
		for i := range p.Patterns {
			var ret *Pattern
			ret, err = dR(p.Patterns[i])
			if err != nil {
				return nil, nil, err
			}
			concatPat, err := c.NewConcat(append([]*Pattern{ret}, p.Patterns[i+1:]...))
			if err != nil {
				return nil, nil, err
			}
			orPatterns = append(orPatterns, concatPat)
			if !p.Patterns[i].nullable {
				break
			}
		}
		o, err := c.NewOr(orPatterns)
		return o, rest, err
	case Or:
		orPatterns, err := traverse(p.Patterns, dR)
		if err != nil {
			return nil, nil, err
		}
		o, err := c.NewOr(orPatterns)
		return o, rest, err
	case And:
		andPatterns, err := traverse(p.Patterns, dR)
		if err != nil {
			return nil, nil, err
		}
		a, err := c.NewAnd(andPatterns)
		return a, rest, err
	case Interleave:
		orPatterns := make([]*Pattern, len(p.Patterns))
		var err error
		for i := range p.Patterns {
			interleaves := make([]*Pattern, len(p.Patterns))
			copy(interleaves, p.Patterns)
			interleaves[i], err = dR(p.Patterns[i])
			if err != nil {
				return nil, nil, err
			}
			orPatterns[i], err = c.NewInterleave(interleaves)
			if err != nil {
				return nil, nil, err
			}
		}
		o, err := c.NewOr(orPatterns)
		return o, rest, err
	case ZeroOrMore:
		pp, err := dR(p.Patterns[0])
		if err != nil {
			return nil, nil, err
		}
		ppp, err := c.NewConcat([]*Pattern{pp, p})
		return ppp, rest, err
	case Reference:
		return derivReturn(c, c.Deref(p.Name), nulls)
	case Not:
		pp, err := dR(p.Patterns[0])
		if err != nil {
			return nil, nil, err
		}
		ppp, err := c.NewNot(pp)
		return ppp, rest, err
	case Contains:
		orPatterns := make([]*Pattern, 0, 3)
		orPatterns = append(orPatterns, p)
		ret, err := dR(p.Patterns[0])
		if err != nil {
			return nil, nil, err
		}
		cp, err := c.NewConcat([]*Pattern{ret, c.NewZAny()})
		if err != nil {
			return nil, nil, err
		}
		orPatterns = append(orPatterns, cp)
		if p.Patterns[0].nullable {
			orPatterns = append(orPatterns, c.NewZAny())
		}
		o, err := c.NewOr(orPatterns)
		return o, rest, err
	case Optional:
		res, err := dR(p.Patterns[0])
		return res, rest, err
	case Extension:
		ext, err := GetExtension(p.Name)
		if err != nil {
			return nil, nil, err
		}
		ret, err := ext.derive(c, dR, p.Patterns)
		if err != nil {
			return nil, nil, err
		}
		return ret, rest, nil
	}
	panic(fmt.Sprintf("unknown pattern typ %d", p.Type))
}
