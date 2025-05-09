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

package mem

import (
	"io"

	"github.com/katydid/parser-go/parser"
	"github.com/katydid/validator-go/validator/intern"
)

func deriv(mem *Mem, patterns int, tree parser.Interface) (int, error) {
	for {
		if !mem.states.Get(patterns).IsEscapable() {
			return patterns, nil
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				return 0, err
			}
		}
		callTree, err := mem.GetCall(patterns)
		if err != nil {
			return 0, err
		}
		childPatterns, zipIndex, err := mem.eval(callTree, tree)
		if err != nil {
			return 0, err
		}
		if !tree.IsLeaf() {
			tree.Down()
			childPatterns, err = deriv(mem, childPatterns, tree)
			if err != nil {
				return 0, err
			}
			tree.Up()
		}
		nullIndex := mem.states.Get(childPatterns).NullIndex
		patterns, err = mem.GetReturn(patterns, zipIndex, nullIndex)
		if err != nil {
			return 0, err
		}
	}
	return patterns, nil
}

func (this *Mem) eval(ifs *intern.IfExprs, label parser.Value) (int, int, error) {
	state, err := ifs.Eval(this.states, label)
	if err != nil {
		return 0, 0, err
	}
	p := this.states.Get(state)
	return p.ZippedPatternsIndex, p.ZippedIndexesIndex, nil
}
