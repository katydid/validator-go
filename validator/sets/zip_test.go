//  Copyright 2017 Walter Schulze
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

import (
	"reflect"
	"sort"
	"testing"

	"katydid.org.za/go/validator-go/validator/ast"
)

func Zip(patterns []*ast.Pattern) ([]*ast.Pattern, []int) {

	ps := make([]*ast.Pattern, len(patterns))
	orighashes := make([]uint64, len(patterns))
	hashes := make([]uint64, len(patterns))
	for i, p := range patterns {
		ps[i] = patterns[i]
		h := p.Hash()
		hashes[i] = h
		orighashes[i] = h
	}

	// sort
	sortable := &sortable{hashes, ps}
	sort.Sort(sortable)

	// remove zany and not zany
	u := 0
	for i := 0; i < len(ps); i++ {
		if ps[i].ZAny != nil {
			continue
		}
		if ps[i].Not != nil && ps[i].Not.Pattern.ZAny != nil {
			continue
		}
		ps[u] = ps[i]
		hashes[u] = hashes[i]
		u++
	}
	ps = ps[:u]
	hashes = hashes[:u]

	// remove duplicates
	if len(ps) > 0 {
		u = 0
		for i := 1; i < len(ps); i++ {
			if hashes[i] == hashes[u] &&
				ps[i].Equal(ps[u]) {
				continue
			}
			u++
			if u != i {
				ps[u] = ps[i]
				hashes[u] = hashes[i]
			}
		}
		ps = ps[:u+1]
		hashes = hashes[:u+1]
	}

	// calculate indexes by doing a reverse lookup using the original hashes and moved hashes.
	revhashes := make(map[uint64][]int)
	for i, h := range hashes {
		revhashes[h] = append(revhashes[h], i)
	}
	indexes := make([]int, len(patterns))
	for i := range patterns {
		if patterns[i].ZAny != nil {
			indexes[i] = -1
		} else if patterns[i].Not != nil && patterns[i].Not.Pattern.ZAny != nil {
			indexes[i] = -2
		} else {
			hashindexes, ok := revhashes[orighashes[i]]
			if !ok {
				panic("wtf")
			}
			if len(hashindexes) == 1 {
				indexes[i] = hashindexes[0]
			} else {
				for _, index := range hashindexes {
					if ps[index].Equal(patterns[i]) {
						indexes[i] = index
					}
				}
			}
		}
	}

	return ps, indexes
}

func Unzip(patterns []*ast.Pattern, indexes []int) []*ast.Pattern {
	res := make([]*ast.Pattern, len(indexes))
	for i, index := range indexes {
		if index < 0 {
			index += 1
			index = index * -1
			res[i] = zipIgnoreSet[index]
		} else {
			res[i] = patterns[index]
		}
	}
	return res
}

func TestZip0(t *testing.T) {
	want := []*ast.Pattern{}
	zips, zipi := Zip(want)
	if len(zips) != 0 {
		t.Errorf("wanted zero in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipZAny(t *testing.T) {
	want := []*ast.Pattern{ast.NewZAny(), ast.NewZAny()}
	zips, zipi := Zip(want)
	if len(zips) != 0 {
		t.Errorf("wanted zero in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipNotZAny(t *testing.T) {
	want := []*ast.Pattern{ast.NewNot(ast.NewZAny()), ast.NewNot(ast.NewZAny()), ast.NewNot(ast.NewZAny())}
	zips, zipi := Zip(want)
	if len(zips) != 0 {
		t.Errorf("wanted zero in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipNotAndZAny(t *testing.T) {
	want := []*ast.Pattern{ast.NewZAny(), ast.NewNot(ast.NewZAny()), ast.NewZAny(), ast.NewZAny(), ast.NewZAny(), ast.NewZAny(), ast.NewNot(ast.NewZAny()), ast.NewNot(ast.NewZAny()), ast.NewNot(ast.NewZAny())}
	zips, zipi := Zip(want)
	if len(zips) != 0 {
		t.Errorf("wanted zero in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipA(t *testing.T) {
	a := ast.NewLeafNode(ast.NewStringConst("a"))
	want := []*ast.Pattern{a, a, a, a}
	zips, zipi := Zip(want)
	if len(zips) != 1 {
		t.Errorf("wanted one in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipAB(t *testing.T) {
	a := ast.NewLeafNode(ast.NewStringConst("a"))
	b := ast.NewLeafNode(ast.NewStringConst("b"))
	want := []*ast.Pattern{a, b, a, b, b}
	zips, zipi := Zip(want)
	if len(zips) != 2 {
		t.Errorf("wanted two in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipABNotAndZAny(t *testing.T) {
	a := ast.NewLeafNode(ast.NewStringConst("a"))
	b := ast.NewLeafNode(ast.NewStringConst("b"))
	want := []*ast.Pattern{a, b, a, ast.NewZAny(), b, ast.NewZAny(), b, b, ast.NewNot(ast.NewZAny())}
	zips, zipi := Zip(want)
	if len(zips) != 2 {
		t.Errorf("wanted two in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipANoZip(t *testing.T) {
	a := ast.NewLeafNode(ast.NewStringConst("a"))
	want := []*ast.Pattern{a}
	zips, zipi := Zip(want)
	if len(zips) != 1 {
		t.Errorf("wanted one in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipABNoZip(t *testing.T) {
	a := ast.NewLeafNode(ast.NewStringConst("a"))
	b := ast.NewLeafNode(ast.NewStringConst("b"))
	want := []*ast.Pattern{a, b}
	zips, zipi := Zip(want)
	if len(zips) != 2 {
		t.Errorf("wanted two in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestZipABCNoZip(t *testing.T) {
	a := ast.NewLeafNode(ast.NewStringConst("a"))
	b := ast.NewLeafNode(ast.NewStringConst("b"))
	c := ast.NewLeafNode(ast.NewStringConst("c"))
	want := []*ast.Pattern{a, b, c}
	zips, zipi := Zip(want)
	if len(zips) != 3 {
		t.Errorf("wanted three in my zipped set, but got %d", len(zips))
	}
	got := Unzip(zips, zipi)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %s got %s", want, got)
	}
}
