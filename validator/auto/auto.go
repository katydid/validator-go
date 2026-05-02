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

// Package auto compiles a parsed validator grammar into a visual pushdown automaton and executes it.
// Compilation into a VPA may result in an exponential explosion, since fully converting a grammar to VPA is O(2^n^2).
// Rather use the mem package.  It gives comparable speed and has no exponential behaviour.
// This package is just here to provide a benchmark against the mem package.
package auto

import (
	"errors"
	"io"

	"github.com/katydid/parser-go/parse"
)

// Auto is the structure that represents the automaton.
type Auto struct {
	calls           []*callNode
	returns         []map[int]int
	escapables      []bool
	start           int
	stateToNullable []int
	accept          []bool
}

// Validate executes an automaton with the given parser and returns whether the parser is valid given the automaton's original grammar.
func (auto *Auto) Validate(p parse.Parser) (bool, error) {
	final, err := deriv(auto, auto.start, p)
	if err != nil {
		return false, err
	}
	return auto.accept[final], nil
}

func (auto *Auto) MetricNumberOfStates() int {
	return len(auto.accept)
}

func derivEnter(auto *Auto, current int, tree parse.Parser) (int, int, error) {
	callTree := auto.calls[current]
	return callTree.eval(tree)
}

func derivLeave(auto *Auto, childState int, stackElm int) int {
	nullIndex := auto.stateToNullable[childState]
	return auto.returns[stackElm][nullIndex]
}

func derivValue(auto *Auto, patterns int, tree parse.Parser) (int, error) {
	childState, stackElm, err := derivEnter(auto, patterns, tree)
	if err != nil {
		return 0, err
	}
	return derivLeave(auto, childState, stackElm), nil
}

var errUnknownHint = errors.New("unknonw hint")

var errUnknownFieldHint = errors.New("unknonw field hint")

func deriv(auto *Auto, current int, tree parse.Parser) (int, error) {
	for {
		if !auto.escapables[current] {
			tree.Skip()
			return current, nil
		}
		hint, err := tree.Next()
		if err != nil {
			if err == io.EOF {
				return current, nil
			}
			return 0, err
		}
		switch hint {
		case parse.EnterHint:
			// derive children, until a LeaveHint and then derive the Next
			current, err = deriv(auto, current, tree)
			if err != nil {
				return 0, err
			}
		case parse.LeaveHint:
			return current, nil
		case parse.FieldHint:
			childState, stackElm, err := derivEnter(auto, current, tree)
			if err != nil {
				return 0, err
			}
			childHint, err := tree.Next()
			switch childHint {
			case parse.EnterHint:
				childState, err = deriv(auto, childState, tree)
				if err != nil {
					return 0, err
				}
			case parse.ValueHint:
				childState, err = derivValue(auto, childState, tree)
				if err != nil {
					return 0, err
				}
			default:
				return 0, errUnknownFieldHint
			}
			current = derivLeave(auto, childState, stackElm)
		case parse.ValueHint:
			// derive value and then derive the Next
			current, err = derivValue(auto, current, tree)
			if err != nil {
				return 0, err
			}

		case parse.UnknownHint:
			return 0, errUnknownHint
		}
	}
}
