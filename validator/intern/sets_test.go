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

import (
	"testing"

	"github.com/katydid/validator-go/validator/sets"
)

func TestSetsAddIndex(t *testing.T) {
	s := NewSetOfPatterns()
	zanys := []*Pattern{newZAny()}
	want := s.Add(zanys)
	got := s.Index(zanys)
	if got != want {
		t.Fatalf("got %d != want %d", got, want)
	}
	if !s.Get(want).IsAccept() {
		t.Fatal("not accept")
	}
	notzanys := []*Pattern{newNotZAny(), newNotZAny()}
	if s.Index(notzanys) != -1 {
		t.Fatal("not found")
	}
	state := s.Add(notzanys)
	if s.Get(state).NullIndex != 2 {
		t.Fatalf("nullindex != 2, but %d", s.Get(state).NullIndex)
	}
	if s.setOfNullables.Index(sets.NewBits(2)) != s.Get(state).NullIndex {
		t.Fatal("wrong nullindex")
	}
	if s.Get(state).IsAccept() {
		t.Fatal("accept")
	}
}

func TestSetsLookup(t *testing.T) {
	s := NewSetOfPatterns()
	zanys := []*Pattern{newZAny()}
	want1 := s.Add(zanys)
	got1 := s.Index(zanys)
	if got1 != want1 {
		t.Fatalf("got %d != want %d", got1, want1)
	}
	notzanys := []*Pattern{newNotZAny(), newNotZAny()}
	want2 := s.Add(notzanys)
	got2 := s.Index(notzanys)
	if got2 != want2 {
		t.Fatalf("got %d != want %d", got2, want2)
	}
	got1 = s.Add(zanys)
	if got1 != want1 {
		t.Fatalf("got %d != want %d", got1, want1)
	}
	got2 = s.Index(notzanys)
	if got2 != want2 {
		t.Fatalf("got %d != want %d", got2, want2)
	}
}
