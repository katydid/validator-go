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
	"strings"

	"github.com/katydid/parser-go/parser"
	"github.com/katydid/validator-go/validator/compose"
	"github.com/katydid/validator-go/validator/funcs"
)

type IfExprs struct {
	prev funcs.Bool
	f    funcs.Bool
	ifs  []*IfExpr

	Cond compose.Bool
	Thn  *IfExprs
	Els  *IfExprs

	ps  []*Pattern
	res int
}

func NewIfExprs(ifs []*IfExpr) *IfExprs {
	return &IfExprs{prev: funcs.BoolConst(true), ifs: ifs}
}

func CompileIfExprs(set *SetOfPatterns, ifs []*IfExpr) (*IfExprs, error) {
	ifExprs := NewIfExprs(ifs)
	if err := ifExprs.compileAll(set); err != nil {
		return nil, err
	}
	return ifExprs, nil
}

func (node *IfExprs) GetRes(set *SetOfPatterns) (*Patterns, error) {
	if len(node.ifs) > 0 {
		if err := node.compileRes(set); err != nil {
			return nil, err
		}
		return set.Get(node.res), nil
	}
	return nil, nil
}

func (node *IfExprs) GetAllRes(set *SetOfPatterns) ([]int, error) {
	if len(node.ifs) > 0 {
		if err := node.compileRes(set); err != nil {
			return nil, err
		}
		return []int{node.res}, nil
	}
	thens, err := node.Thn.GetAllRes(set)
	if err != nil {
		return nil, err
	}
	elses, err := node.Els.GetAllRes(set)
	if err != nil {
		return nil, err
	}
	ress := []int{}
	ress = append(ress, thens...)
	ress = append(ress, elses...)
	return ress, nil
}

func (node *IfExprs) compileAll(set *SetOfPatterns) error {
	if err := node.compileRes(set); err != nil {
		return err
	}
	if err := node.compileCond(); err != nil {
		return err
	}
	if node.Cond == nil {
		if err := node.compileRes(set); err != nil {
			return err
		}
		return nil
	}
	if err := node.compileThn(); err != nil {
		return err
	}
	if node.Thn != nil {
		if err := node.Thn.compileAll(set); err != nil {
			return err
		}
	}
	if err := node.compileEls(); err != nil {
		return err
	}
	if node.Els != nil {
		if err := node.Els.compileAll(set); err != nil {
			return err
		}
	}
	return nil
}

func (ifExprs *IfExprs) Eval(set *SetOfPatterns, label parser.Value) (int, error) {
	return ifExprs.eval(set, label)
}

func (this *IfExprs) String() string {
	if this == nil {
		return "...\n"
	}
	if len(this.ifs) == 0 {
		ss := make([]string, len(this.ps))
		for i := range this.ps {
			ss[i] = this.ps[i].String()
		}
		return fmt.Sprintf("%s\n", strings.Join(ss, ","))
	}
	if this.f == nil {
		return "...\n"
	}
	return "if (" + this.f.ToExpr().String() + ") then {\n" + this.Thn.String() + "} else {\n" + this.Els.String() + "}"
}

var falseConst = funcs.BoolConst(false)
var trueConst = funcs.BoolConst(true)

func (node *IfExprs) eval(set *SetOfPatterns, label parser.Value) (int, error) {
	if len(node.ifs) == 0 {
		if err := node.compileRes(set); err != nil {
			return 0, err
		}
		return node.res, nil
	}
	if err := node.compileCond(); err != nil {
		return 0, err
	}
	// everything simplified to always true or always false
	if node.Cond == nil {
		if err := node.compileRes(set); err != nil {
			return 0, err
		}
		return node.res, nil
	}
	cond, err := node.Cond.Eval(label)
	if err != nil {
		return 0, err
	}
	if cond {
		if err := node.compileThn(); err != nil {
			return 0, err
		}
		return node.Thn.eval(set, label)
	}
	if err := node.compileEls(); err != nil {
		return 0, err
	}
	return node.Els.eval(set, label)
}

func (node *IfExprs) compileRes(set *SetOfPatterns) error {
	if len(node.ifs) == 0 {
		if node.ps != nil {
			node.res = set.Add(node.ps)
			node.ps = nil
		}
	}
	return nil
}

func (node *IfExprs) compileCond() error {
	if node.f != nil {
		return nil
	}
	if len(node.ifs) == 0 {
		return nil
	}
	node.f = node.ifs[0].Cond
	if funcs.Equal(node.prev, node.f) {
		node.f = trueConst
	} else if funcs.IsFalse(funcs.And(node.prev, node.f)) {
		node.f = falseConst
	} else if funcs.IsFalse(funcs.And(node.prev, funcs.Not(node.f))) {
		node.f = trueConst
	}
	if funcs.IsTrue(node.f) {
		node.ps = append(node.ps, node.ifs[0].Thn)
		node.ifs = node.ifs[1:]
		node.f = nil
		return node.compileCond()
	}
	if funcs.IsFalse(node.f) {
		node.ps = append(node.ps, node.ifs[0].Els)
		node.ifs = node.ifs[1:]
		node.f = nil
		return node.compileCond()
	}
	c, err := compose.NewBoolFunc(node.f)
	if err != nil {
		return err
	}
	node.Cond = c
	return nil
}

func (node *IfExprs) compileThn() error {
	if node.Thn != nil {
		return nil
	}
	if len(node.ifs) == 0 {
		return nil
	}
	node.Thn = &IfExprs{}
	node.Thn.ps = make([]*Pattern, 0, len(node.ps)+1)
	node.Thn.ps = append(node.Thn.ps, node.ps...)
	node.Thn.ps = append(node.Thn.ps, node.ifs[0].Thn)
	node.Thn.prev = funcs.And(node.prev, node.f)
	node.Thn.ifs = node.ifs[1:]
	return nil
}

func (node *IfExprs) compileEls() error {
	if node.Els != nil {
		return nil
	}
	if len(node.ifs) == 0 {
		return nil
	}
	node.Els = &IfExprs{}
	node.Els.ps = make([]*Pattern, 0, len(node.ps)+1)
	node.Els.ps = append(node.Els.ps, node.ps...)
	node.Els.ps = append(node.Els.ps, node.ifs[0].Els)
	node.Els.prev = funcs.And(node.prev, funcs.Not(node.f))
	node.Els.ifs = node.ifs[1:]
	return nil
}

type IfExpr struct {
	Cond funcs.Bool
	Thn  *Pattern
	Els  *Pattern
}

func NewIfExpr(c funcs.Bool, thn, els *Pattern) *IfExpr {
	return &IfExpr{c, thn, els}
}

func (this *IfExpr) eval(label parser.Value) (*Pattern, error) {
	f, err := compose.NewBoolFunc(this.Cond)
	if err != nil {
		return nil, err
	}
	cond, err := f.Eval(label)
	if err != nil {
		return nil, err
	}
	if cond {
		return this.Thn, nil
	}
	return this.Els, nil
}

func evalIfExprs(ifs []*IfExpr, label parser.Value) ([]*Pattern, error) {
	patterns := make([]*Pattern, len(ifs))
	for i, ifExpr := range ifs {
		c, err := ifExpr.eval(label)
		if err != nil {
			return nil, err
		}
		patterns[i] = c
	}
	return patterns, nil
}
