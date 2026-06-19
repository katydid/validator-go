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

type PatternsSet struct {
	list    [][]*intern.Pattern
	strings map[string]int
}

func NewPatternsSet() PatternsSet {
	return PatternsSet{
		list:    [][]*intern.Pattern{},
		strings: make(map[string]int),
	}
}

func (this *PatternsSet) Get(i int) []*intern.Pattern {
	return this.list[i]
}

func (this *PatternsSet) Len() int {
	return len(this.list)
}

func (this *PatternsSet) Index(x []*intern.Pattern) int {
	str := intern.StringPatterns(x)
	index, ok := this.strings[str]
	if !ok {
		return -1
	}
	return index
}

func (this *PatternsSet) Add(x []*intern.Pattern) int {
	index := this.Index(x)
	if index != -1 {
		// item has already been added, so just return that index
		return index
	}
	// add item to list
	this.list = append(this.list, x)
	index = len(this.list) - 1
	str := intern.StringPatterns(x)
	this.strings[str] = index
	return index
}
