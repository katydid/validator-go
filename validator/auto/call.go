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
	"github.com/katydid/parser-go/cast"
	"github.com/katydid/parser-go/parse"
	"github.com/katydid/validator-go/validator/compose"
	"github.com/katydid/validator-go/validator/funcs"
	"github.com/katydid/validator-go/validator/intern"
	"github.com/katydid/validator-go/validator/sets"
	"github.com/katydid/validator-go/validator/std"
)

// callNode represents a node in the call tree.
// The call tree is basically a nested if then else tree, which results in child states and stack elements.
type callNode struct {
	cond       compose.Bool
	thens      map[string]*callNode
	then       *callNode
	els        *callNode
	child      int
	stackIndex int
}

func (this *compiler) newRetNode(parentPatterns int, ps []*intern.Pattern) *callNode {
	zipped, _ := intern.Zip(ps)
	zippedPatterns, zipper := zipped.Patterns, zipped.Indexes
	zipperIndex := this.zis.Add(zipper)
	stackIndex := this.stackElms.Add(sets.Pair{First: parentPatterns, Second: zipperIndex})
	zippedPatternIndex := this.patterns.Add(zippedPatterns)
	return &callNode{
		child:      zippedPatternIndex,
		stackIndex: stackIndex,
	}
}

func (this *compiler) newCallTree(parentPatterns int, node *ifExprs) (*callNode, error) {
	if node.ret != nil {
		return this.newRetNode(parentPatterns, node.ret), nil
	}
	stringThens := make(map[string]*ifExprs)
	stringElse := getThenStrings(node, stringThens)
	if len(stringThens) > 3 {
		// We assume it is quicker to check in a hashmap once, than do 3 equal comparisons, but this heuristic could be tuned
		thens := make(map[string]*callNode)
		for name := range stringThens {
			thens[name] = this.newRetNode(parentPatterns, stringThens[name].ret)
		}
		els, err := this.newCallTree(parentPatterns, stringElse)
		if err != nil {
			return nil, err
		}
		return &callNode{
			thens: thens,
			els:   els,
		}, nil
	}

	then, err := this.newCallTree(parentPatterns, node.then)
	if err != nil {
		return nil, err
	}
	els, err := this.newCallTree(parentPatterns, node.els)
	if err != nil {
		return nil, err
	}
	f, err := compose.NewBoolFunc(node.cond)
	if err != nil {
		return nil, err
	}
	return &callNode{
		cond: f,
		then: then,
		els:  els,
	}, nil
}

func getThenStrings(node *ifExprs, acc map[string]*ifExprs) *ifExprs {
	if node.cond == nil {
		return node
	}
	name, ok := funcs.GetStringEqConst(node.cond)
	if !ok {
		return node
	}
	if node.then.ret == nil {
		return node
	}
	acc[name] = node.then
	return getThenStrings(node.els, acc)
}

// eval evaluates the call tree returning the child state and stack element given the label value of the parser element.
func (this *callNode) eval(label parse.Token) (int, int, error) {
	if this.cond != nil {
		cond, err := this.cond.Eval(label)
		if err != nil {
			return 0, 0, err
		}
		if cond {
			return this.then.eval(label)
		}
		return this.els.eval(label)
	} else if this.thens != nil {
		k, v, err := label.Token()
		if err != nil {
			return this.els.eval(label)
		}
		if k != parse.StringKind {
			return this.els.eval(label)
		}
		name := cast.ToString(v)
		res, ok := this.thens[name]
		if !ok {
			return this.els.eval(label)
		}
		return res.child, res.stackIndex, nil
	}
	return this.child, this.stackIndex, nil
}

func (this *callNode) getLeafs() []*callNode {
	if this.cond != nil {
		then := this.then.getLeafs()
		els := this.els.getLeafs()
		return append(then, els...)
	} else if this.thens != nil {
		els := this.els.getLeafs()
		res := make([]*callNode, 0, len(this.thens)+len(els))
		names := std.SortedKeys(this.thens)
		for _, name := range names {
			res = append(res, this.thens[name])
		}
		res = append(res, els...)
		return res
	}
	return []*callNode{this}
}
