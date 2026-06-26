//  Copyright 2026 Walter Schulze
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

func IsDisjoints(fs []Bool) bool {
	m := make(map[string]struct{})
	for _, f := range fs {
		switch ft := f.(type) {
		case *stringEq:
			c, ok := getConstString(ft.V1, ft.V2)
			if !ok {
				return false
			}
			if _, ok := m[c]; ok {
				return false
			}
			m[c] = struct{}{}
		default:
			return false
		}
	}
	return true
}

// IsDisjoint returns true if only one pattern can be true at the same time.
// We cannot be perfect, so this does return false negatives, but never false postives.
func IsDisjoint(f1, f2 Bool) bool {
	switch ft1 := f1.(type) {
	case *stringEq:
		switch ft2 := f2.(type) {
		case *stringEq:
			c1, ok1 := getConstString(ft1.V1, ft1.V2)
			c2, ok2 := getConstString(ft1.V1, ft2.V2)
			return ok1 && ok2 && c1 != c2
		}
	}
	return false
}

// getConstString returns the const if the other input is variable.
func getConstString(a, b interface{}) (string, bool) {
	if c, aok := a.(ConstString); aok {
		if _, bok := b.(aVariable); bok {
			s, err := c.Eval()
			if err != nil {
				return "", false
			}
			return s, true
		}
	} else if c, bok := b.(ConstString); bok {
		if _, aok := a.(aVariable); aok {
			s, err := c.Eval()
			if err != nil {
				return "", false
			}
			return s, true
		}
	}
	return "", false
}
