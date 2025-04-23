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

package intern

import "github.com/katydid/validator-go/validator/sets"

// SetOfPatterns represents an indexed list of list of Patterns.
// It reverse maps a list of Patterns into a single int.
type SetOfPatterns struct {
	list            []*Patterns
	hashes          map[uint64][]int
	setOfNullables  sets.BitsSet
	setOfZipIndexes sets.Ints
}

func NewSetOfPatterns() *SetOfPatterns {
	return &SetOfPatterns{
		list:            []*Patterns{},
		hashes:          make(map[uint64][]int),
		setOfNullables:  sets.NewBitsSet(),
		setOfZipIndexes: sets.NewInts(),
	}
}

func (this *SetOfPatterns) Get(i int) *Patterns {
	return this.list[i]
}

func (this *SetOfPatterns) Index(patterns []*Pattern) int {
	h := HashPatterns(patterns)
	pss := this.hashes[h]
	for _, index := range pss {
		ps := this.list[index]
		if deriveEquals(ps.Patterns, patterns) {
			return index
		}
	}
	return -1
}

func (this *SetOfPatterns) Add(ps []*Pattern) int {
	index := this.Index(ps)
	if index != -1 {
		return index
	}
	index = len(this.list)
	patterns := &Patterns{}
	this.list = append(this.list, patterns)
	h := HashPatterns(ps)
	this.hashes[h] = append(this.hashes[h], index)

	patterns.Patterns = ps

	nulls := newNullableBits(ps)
	patterns.NullIndex = this.setOfNullables.Add(nulls)

	patterns.escapable = Escapable(ps)
	patterns.accept = len(ps) == 1 && ps[0].nullable
	zipped, _ := Zip(ps)
	patterns.ZippedPatternsIndex = this.Add(zipped.Patterns)
	patterns.ZippedIndexesIndex = this.setOfZipIndexes.Add(zipped.Indexes)
	return index
}

func (this *SetOfPatterns) Nullable(nullIndex, zipIndex int) []bool {
	return sets.UnzipBits(this.setOfNullables[nullIndex], this.setOfZipIndexes[zipIndex])
}

type Patterns struct {
	Patterns []*Pattern

	hash      uint64
	escapable bool
	accept    bool
	nulls     sets.Bits
	zipped    *ZippedPatterns

	NullIndex           int
	ZippedPatternsIndex int
	ZippedIndexesIndex  int
}

func NewPatterns(ps []*Pattern, zipped *ZippedPatterns) *Patterns {
	return &Patterns{
		Patterns:  ps,
		hash:      HashPatterns(ps),
		escapable: Escapable(ps),
		accept:    len(ps) == 1 && ps[0].nullable,
		nulls:     newNullableBits(ps),
		zipped:    zipped,
	}
}

func (this *Patterns) IsAccept() bool {
	return this.accept
}

func (this *Patterns) IsEscapable() bool {
	return this.escapable
}
