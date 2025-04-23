//  Copyright 2025 Walter Schulze
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
	"github.com/katydid/validator-go/validator/intern"
)

// type PatternsSet = *sets.IndexedSet[[]*intern.Pattern]
type PatternsSet = Patterns

func NewPatternsSet() PatternsSet {
	return NewPatterns()
}

// Patterns represents an indexed list of list of Patterns.
// It reverse maps a list of Patterns into a single int.
type Patterns [][]*intern.Pattern

func NewPatterns() Patterns {
	return Patterns([][]*intern.Pattern{})
}

func (this Patterns) Get(i int) []*intern.Pattern {
	return this[i]
}

func (this Patterns) Len() int {
	return len(this)
}

func (this Patterns) Index(patterns []*intern.Pattern) int {
	for i, ps := range this {
		// TODO maybe we should rather use hashes to do this more efficiently?
		if intern.EqualPatterns(ps, patterns) {
			return i
		}
	}
	return -1
}

func (this *Patterns) Add(patterns []*intern.Pattern) int {
	index := this.Index(patterns)
	if index != -1 {
		return index
	}
	*this = append(*this, patterns)
	return len(*this) - 1
}
