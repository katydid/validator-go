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

package sets

// Ints represents an indexed list of list of integers.
// It reverse maps a list of ints into a single int.
type Ints [][]int

func NewInts() Ints {
	return Ints([][]int{})
}

func (this Ints) Get(i int) []int {
	return this[i]
}

func (this Ints) Index(is []int) int {
	for i, iis := range this {
		if len(is) == len(iis) {
			eq := true
			for j := range iis {
				if iis[j] != is[j] {
					eq = false
					break
				}
			}
			if eq {
				return i
			}
		}
	}
	return -1
}

func (this *Ints) Add(is []int) int {
	index := this.Index(is)
	if index != -1 {
		return index
	}
	*this = append(*this, is)
	return len(*this) - 1
}

// BitsSet represents an indexed list of Bits.
// It reverse maps a Bits into a single int.
type BitsSet []Bits

func NewBitsSet() BitsSet {
	return BitsSet([]Bits{})
}

func (this BitsSet) Get(i int) Bits {
	return this[i]
}

func (this BitsSet) Index(bs Bits) int {
	for i, ibs := range this {
		if ibs.Equal(bs) {
			return i
		}
	}
	return -1
}

func (this *BitsSet) Add(bs Bits) int {
	index := this.Index(bs)
	if index != -1 {
		return index
	}
	*this = append(*this, bs)
	return len(*this) - 1
}

type Pair struct {
	First  int
	Second int
}

// Pairs represents an indexed list of Pair pairs.
// It reverse maps a list of Pairs into a single int.
type Pairs []Pair

func NewPairs() Pairs {
	return Pairs([]Pair{})
}

func (this Pairs) Index(se Pair) int {
	for i, ise := range this {
		if ise == se {
			return i
		}
	}
	return -1
}

func (this *Pairs) Add(se Pair) int {
	index := this.Index(se)
	if index != -1 {
		return index
	}
	*this = append(*this, se)
	return len(*this) - 1
}

type IndexedSet[A any] struct {
	hashFunc  func(A) uint64
	equalFunc func(A, A) bool
	list      []A
	hashes    map[uint64][]int
}

func NewIndexedSet[A any](hash func(A) uint64, eq func(A, A) bool) *IndexedSet[A] {
	return &IndexedSet[A]{
		hashFunc:  hash,
		equalFunc: eq,
		list:      []A{},
		hashes:    make(map[uint64][]int),
	}
}

func (this *IndexedSet[A]) Get(i int) A {
	return this.list[i]
}

func (this *IndexedSet[A]) Len() int {
	return len(this.list)
}

func (this *IndexedSet[A]) Index(x A) int {
	h := this.hashFunc(x)
	// find items that match the same hash
	hashedIndexesOfItems := this.hashes[h]
	// loop through indexes of items with hash conflicts
	for _, index := range hashedIndexesOfItems {
		// if item is equal, then we found a duplicate item and we can return the index
		if this.equalFunc(x, this.list[index]) {
			return index
		}
	}
	return -1
}

func (this *IndexedSet[A]) Add(x A) int {
	index := this.Index(x)
	if index != -1 {
		// item has already been added, so just return that index
		return index
	}
	// add item to list
	this.list = append(this.list, x)

	// add index of item to hash map for faster lookup
	index = len(this.list) - 1
	h := this.hashFunc(x)
	this.hashes[h] = append(this.hashes[h], index)
	return index
}
