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
	"github.com/katydid/validator-go/validator/ast"
	"github.com/katydid/validator-go/validator/funcs"
	"github.com/katydid/validator-go/validator/intern"
	"github.com/katydid/validator-go/validator/sets"
)

func newCompiler(g *ast.Grammar, record bool) (*compiler, error) {
	construct := intern.NewConstructor()
	if record {
		construct = intern.NewConstructorOptimizedForRecords()
	}
	main, err := construct.AddGrammar(g)
	if err != nil {
		return nil, err
	}
	m := &compiler{
		construct: construct,
		patterns:  NewPatternsSet(),
		zis:       sets.NewInts(),
		stackElms: sets.NewPairs(),
		nullables: sets.NewBitsSet(),

		calls:           []*callNode{},
		returns:         []map[int]int{},
		escapables:      []bool{},
		stateToNullable: []int{},
		accept:          []bool{},
	}
	start := m.patterns.Add([]*intern.Pattern{main})
	// TOOD: What to do with context?
	m.start = start
	return m, nil
}

type compiler struct {
	context *funcs.Context

	construct intern.Construct
	patterns  PatternsSet
	zis       sets.Ints
	stackElms sets.Pairs
	nullables sets.BitsSet

	start           int
	calls           []*callNode
	returns         []map[int]int
	escapables      []bool
	stateToNullable []int
	accept          []bool
}

func (this *compiler) calcEscapables(upto int) {
	for i := len(this.escapables); i <= upto; i++ {
		patterns := this.patterns.Get(i)
		this.escapables = append(this.escapables, intern.Escapable(patterns))
	}
}

func (this *compiler) escapable(patterns int) bool {
	this.calcEscapables(patterns)
	return this.escapables[patterns]
}

func (this *compiler) calcAccepts(upto int) {
	for i := len(this.accept); i <= upto; i++ {
		patterns := this.patterns.Get(i)
		if len(patterns) != 1 {
			this.accept = append(this.accept, false)
		} else {
			this.accept = append(this.accept, patterns[0].Nullable())
		}
	}
}

func (this *compiler) getAccept(patterns int) bool {
	this.calcAccepts(patterns)
	return this.accept[patterns]
}

func (this *compiler) calcCallTrees(upto int) error {
	for i := len(this.calls); i <= upto; i++ {
		listOfIfExpr := intern.DeriveCalls(this.construct, this.patterns.Get(i))
		compiledIfExprs := compileIfExprs(listOfIfExpr)
		memCallTree, err := this.newCallTree(i, compiledIfExprs)
		if err != nil {
			return err
		}
		this.calls = append(this.calls, memCallTree)
	}
	return nil
}

func (this *compiler) getCallTree(patterns int) (*callNode, error) {
	if err := this.calcCallTrees(patterns); err != nil {
		return nil, err
	}
	return this.calls[patterns], nil
}

func nullables(patterns []*intern.Pattern) sets.Bits {
	nulls := sets.NewBits(len(patterns))
	for i, p := range patterns {
		nulls.Set(i, p.Nullable())
	}
	return nulls
}

func (this *compiler) calcNullables(upto int) {
	for i := len(this.stateToNullable); i <= upto; i++ {
		childPatterns := this.patterns.Get(i)
		nullable := nullables(childPatterns)
		nullIndex := this.nullables.Add(nullable)
		this.stateToNullable = append(this.stateToNullable, nullIndex)
	}
}

func (this *compiler) getNullable(s int) int {
	this.calcNullables(s)
	return this.stateToNullable[s]
}

func (this *compiler) getReturn(stackIndex int, nullIndex int) int {
	if len(this.returns) <= stackIndex {
		for i := len(this.returns); i <= stackIndex; i++ {
			this.returns = append(this.returns, make(map[int]int))
		}
	}
	if ret, ok := this.returns[stackIndex][nullIndex]; ok {
		return ret
	}
	stackElm := this.stackElms[stackIndex]
	zullable := this.nullables[nullIndex]
	childrenZipper := stackElm.Second
	nullable := sets.UnzipBits(zullable, this.zis[childrenZipper])
	parentPatterns := stackElm.First
	currentPatterns := this.patterns.Get(parentPatterns)
	newPatterns, err := intern.DeriveReturns(this.construct, currentPatterns, nullable)
	if err != nil {
		// TODO return error
		panic(err)
	}
	res := this.patterns.Add(newPatterns)
	this.returns[stackIndex][nullIndex] = res
	return res
}
