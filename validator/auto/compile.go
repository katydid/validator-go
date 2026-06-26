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
	"errors"

	"github.com/katydid/validator-go/validator/ast"
	"github.com/katydid/validator-go/validator/sets"
)

// A maximum number is set for bitsets as an option, when it is exceeded the number of options to consider is too large and we recommend rather using the memoized version of katydid.
var ErrTooBig = errors.New("a state explosion has occured")

type options struct {
	maxBitSetSize  int
	record         bool
	fieldNameTable bool
}

type Option func(o *options)

func WithMaxBitSetSize(m int) Option {
	return func(o *options) {
		o.maxBitSetSize = m
	}
}

// compiles a parsed validator grammar and optimizes it for the case where the input structures are records, by shrinking space with extra simplification rules.
// Do not use with XML, but JSON, Reflect and Protobufs are safe.
func WithRecordSimplificationRules() Option {
	return func(o *options) {
		o.record = true
	}
}

// creates a lookup strings found in field name expressions.
func WithFieldNameTable() Option {
	return func(o *options) {
		o.fieldNameTable = true
	}
}

func newOptions(opts ...Option) *options {
	o := &options{
		fieldNameTable: false,
		maxBitSetSize:  64,
		record:         false,
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Compile compiles a parsed validator grammar ast into a visual pushdown automaton.
func Compile(g *ast.Grammar, options ...Option) (*Auto, error) {
	return compileAuto(g, options...)
}

func compileAuto(g *ast.Grammar, opts ...Option) (*Auto, error) {
	o := newOptions(opts...)
	c, err := newCompiler(g, o.record, o.maxBitSetSize)
	if err != nil {
		return nil, err
	}
	if err := c.compile(); err != nil {
		return nil, err
	}
	a := &Auto{
		calls:           c.calls,
		returns:         c.returns,
		escapables:      c.escapables,
		start:           c.start,
		stateToNullable: c.stateToNullable,
		accept:          c.accept,
		hashedCalls:     c.hashedCalls,
	}
	if !o.fieldNameTable {
		a.hashedCalls = nil
	}
	return a, nil
}

// Compile memoizes the full state space, all possible things that can be memoized.
func (c *compiler) compile() error {
	changed := true
	visited := make(map[int]bool)
	for changed {
		changed = false
		len := c.patterns.Len()
		for state := range len {
			if visited[state] {
				continue
			}
			visited[state] = true
			if err := compile(c, state); err != nil {
				return err
			}
			changed = true
		}
	}
	for state := range c.patterns.Len() {
		if err := c.calcHashCalls(state); err != nil {
			return err
		}
	}
	return nil
}

func compile(c *compiler, patterns int) error {
	c.getNullable(patterns)
	c.getAccept(patterns)
	c.escapable(patterns)

	callTree, err := c.getCallTree(patterns)
	if err != nil {
		return err
	}
	allPossibleCalls := callTree.getLeafs()
	for _, call := range allPossibleCalls {

		numOfChildPatterns := len(c.patterns.Get(call.child))
		if numOfChildPatterns > c.maxBitSetSize {
			return ErrTooBig
		}

		maxPossibleNullables := sets.NewBits(numOfChildPatterns)
		for i := 0; i < numOfChildPatterns; i++ {
			maxPossibleNullables.Set(i, true)
		}

		possibleNullables := sets.NewBits(numOfChildPatterns)
		for {
			nullIndex := c.nullables.Add(possibleNullables)
			_, err := c.getReturn(call.stackIndex, nullIndex)
			if err != nil {
				return err
			}
			if possibleNullables.Equal(maxPossibleNullables) {
				break
			}
			possibleNullables = possibleNullables.Inc()
		}
	}
	return nil
}
