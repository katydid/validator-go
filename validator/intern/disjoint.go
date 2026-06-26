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

package intern

import "github.com/katydid/validator-go/validator/funcs"

// returns true if only one pattern can be true at the same time.
// We cannot be perfect, so this does return false negatives, but never false postives.
func IsDisjoint(ps []*Pattern) bool {
	fs := make([]funcs.Bool, len(ps))
	for i := range ps {
		fs[i] = ps[i].Func
	}
	return funcs.IsDisjoints(fs)
}
