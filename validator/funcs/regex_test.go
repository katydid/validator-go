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

package funcs

import "testing"

func TestRegexMatch(t *testing.T) {
	r, err := Regex(StringConst(".*b.*"), StringConst("abc"))
	if err != nil {
		t.Fatal(err)
	}
	match, err := r.Eval()
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatalf("expected match")
	}
}

func TestRegexNotMatch(t *testing.T) {
	r, err := Regex(StringConst(".*b.*"), StringConst("adc"))
	if err != nil {
		t.Fatal(err)
	}
	match, err := r.Eval()
	if err != nil {
		t.Fatal(err)
	}
	if match {
		t.Fatalf("unexpected match")
	}
}
