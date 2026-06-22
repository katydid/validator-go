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
	"errors"
	"io"

	"github.com/katydid/parser-go/parse"
	"github.com/katydid/validator-go/validator/intern"
)

func derivEnter(mem *Mem, patterns int, tree parse.Parser) (int, int, error) {
	childPatterns, zipIndex, ok := mem.GetFieldNameCall(patterns, tree)
	if ok {
		return childPatterns, zipIndex, nil
	}
	return deriveEnterNoFieldNameCall(mem, patterns, tree)
}

func deriveEnterNoFieldNameCall(mem *Mem, patterns int, tree parse.Parser) (int, int, error) {
	callTree, err := mem.GetCall(patterns)
	if err != nil {
		return 0, 0, err
	}
	return mem.eval(callTree, tree)
}

func derivLeave(mem *Mem, patterns int, childPatterns int, zipIndex int) (int, error) {
	nullIndex := mem.states.Get(childPatterns).NullIndex
	return mem.GetReturn(patterns, zipIndex, nullIndex)
}

func derivValue(mem *Mem, patterns int, tree parse.Parser) (int, error) {
	childPatterns, zipIndex, err := derivEnter(mem, patterns, tree)
	if err != nil {
		return 0, err
	}
	return derivLeave(mem, patterns, childPatterns, zipIndex)
}

var errUnknownHint = errors.New("unknonw hint")

var errUnknownFieldHint = errors.New("unknonw field hint")

func deriv(mem *Mem, patterns int, tree parse.Parser) (int, error) {
	for {
		if !mem.states.Get(patterns).IsEscapable() {
			tree.Skip()
			return patterns, nil
		}
		hint, err := tree.Next()
		if err != nil {
			if err == io.EOF {
				return patterns, nil
			}
			return 0, err
		}
		switch hint {
		case parse.EnterHint:
			// derive children, until a LeaveHint and then derive the Next
			patterns, err = deriv(mem, patterns, tree)
			if err != nil {
				return 0, err
			}
		case parse.LeaveHint:
			return patterns, nil
		case parse.FieldHint:
			childPatterns, zipIndex, err := derivEnter(mem, patterns, tree)
			if err != nil {
				return 0, err
			}
			childHint, err := tree.Next()
			switch childHint {
			case parse.EnterHint:
				childPatterns, err = deriv(mem, childPatterns, tree)
				if err != nil {
					return 0, err
				}
			case parse.ValueHint:
				childPatterns, err = derivValue(mem, childPatterns, tree)
				if err != nil {
					return 0, err
				}
			default:
				return 0, errUnknownFieldHint
			}
			patterns, err = derivLeave(mem, patterns, childPatterns, zipIndex)
			if err != nil {
				return 0, err
			}
		case parse.ValueHint:
			// derive value and then derive the Next
			patterns, err = derivValue(mem, patterns, tree)
			if err != nil {
				return 0, err
			}

		case parse.UnknownHint:
			return 0, errUnknownHint
		}
	}
}

func (this *Mem) eval(ifs *intern.IfExprs, label parse.Token) (int, int, error) {
	state, err := ifs.Eval(this.states, label)
	if err != nil {
		return 0, 0, err
	}
	p := this.states.Get(state)
	return p.ZippedPatternsIndex, p.ZippedIndexesIndex, nil
}
