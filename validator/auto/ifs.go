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

package auto

import (
	"github.com/katydid/parser-go/parser"
	"github.com/katydid/validator-go/validator/compose"
	"github.com/katydid/validator-go/validator/funcs"
	"github.com/katydid/validator-go/validator/intern"
)

type ifExprs struct {
	cond     funcs.Bool
	composed compose.Bool
	then     *ifExprs
	els      *ifExprs
	ret      []*intern.Pattern
}

// compileIfExprs combines several if expressions into one nested if expression with a list of return values.
// While combining these if expressions, duplicate and impossible (always false) conditions are removed for efficiency.
func compileIfExprs(ifs []*intern.IfExpr) *ifExprs {
	if len(ifs) == 0 {
		return &ifExprs{
			ret: []*intern.Pattern{},
		}
	}
	root := &ifExprs{}
	if ifs[0].Els == nil || ifs[0].Thn.Equal(ifs[0].Els) {
		root.ret = []*intern.Pattern{ifs[0].Thn}
	} else {
		root.cond = ifs[0].Cond
		root.then = &ifExprs{ret: []*intern.Pattern{ifs[0].Thn}}
		root.els = &ifExprs{ret: []*intern.Pattern{ifs[0].Els}}
	}
	for _, ifexpr := range ifs[1:] {
		if ifexpr.Cond == nil {
			root.addReturn(ifexpr.Thn)
		} else {
			root.addIfExpr(ifexpr.Cond, ifexpr.Thn, ifexpr.Els)
		}
	}
	return root
}

func (this *ifExprs) eval(label parser.Value) ([]*intern.Pattern, error) {
	if this.ret != nil {
		return this.ret, nil
	}
	if this.composed == nil {
		composed, err := compose.NewBoolFunc(this.cond)
		if err != nil {
			return nil, err
		}
		this.composed = composed
	}
	cond, err := this.composed.Eval(label)
	if err != nil {
		return nil, err
	}
	if cond {
		return this.then.eval(label)
	}
	return this.els.eval(label)
}

// addReturn finds the leafs and appends a return to each.
func (this *ifExprs) addReturn(ret *intern.Pattern) {
	if this.ret != nil {
		this.ret = append(this.ret, ret)
		return
	}
	this.then.addReturn(ret)
	this.els.addReturn(ret)
	return
}

func (this *ifExprs) addIfExpr(cond funcs.Bool, then, els *intern.Pattern) {
	// efficienctly append the then and else return to two copies of the current returns.
	if this.ret != nil {
		this.cond = cond
		thenterms := make([]*intern.Pattern, len(this.ret)+1)
		copy(thenterms, this.ret)
		thenterms[len(thenterms)-1] = then
		this.then = &ifExprs{ret: thenterms}
		this.els = &ifExprs{ret: append(this.ret, els)}
		this.ret = nil
		return
	}
	// remove duplicate condition
	if funcs.Equal(this.cond, cond) {
		this.then.addReturn(then)
		this.els.addReturn(els)
		return
	}
	// remove impossible (always false) then condition
	if funcs.IsFalse(funcs.And(this.cond, cond)) {
		this.then.addReturn(els)
		this.els.addIfExpr(cond, then, els)
		return
	}
	// remove impossible (always false) else condition
	if funcs.IsFalse(funcs.And(this.cond, funcs.Not(cond))) {
		this.then.addIfExpr(cond, then, els)
		this.els.addReturn(then)
		return
	}
	this.then.addIfExpr(cond, then, els)
	this.els.addIfExpr(cond, then, els)
	return
}
