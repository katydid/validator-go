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

package intern

func EqualPatterns(ps1, ps2 []*Pattern) bool {
	return deriveEquals(ps1, ps2)
}

func HashPatterns(patterns []*Pattern) uint64 {
	h := uint64(17)
	for _, pattern := range patterns {
		h = 31*h + pattern.hash
	}
	return h
}

func Escapable(patterns []*Pattern) bool {
	for _, p := range patterns {
		if isZAny(p) {
			continue
		}
		if isNotZAny(p) {
			continue
		}
		return true
	}
	return false
}
