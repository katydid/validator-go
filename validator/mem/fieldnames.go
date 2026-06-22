// Copyright 2026 Walter Schulze
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mem

import (
	"github.com/katydid/parser-go/cast"
	"github.com/katydid/parser-go/parse"
	"github.com/katydid/validator-go/validator/intern"
)

func (this *Mem) GetFieldNameCall(state int, tree parse.Parser) (int, int, bool) {
	if this.fieldNameCalls == nil {
		// optimization was disabled
		return 0, 0, false
	}
	if state >= len(this.fieldNameCalls) {
		for i := len(this.fieldNameCalls); i <= state; i++ {
			names := intern.GetFieldNames(this.states.Get(i).Patterns)
			nameMap := make(map[string]*callResult)
			for _, name := range names {
				nameMap[name] = nil
			}
			this.fieldNameCalls = append(this.fieldNameCalls, nameMap)
		}
	}
	k, v, err := tree.Token()
	if err != nil {
		return 0, 0, false
	}
	if k != parse.StringKind {
		return 0, 0, false
	}
	name := cast.ToString(v)
	r, ok := this.fieldNameCalls[state][name]
	if !ok {
		return 0, 0, false
	}
	if r == nil {
		r1, r2, err := deriveEnterNoFieldNameCall(this, state, tree)
		if err != nil {
			return 0, 0, false
		}
		r = &callResult{child: r1, stackIndex: r2}
		this.fieldNameCalls[state][name] = r
	}
	return r.child, r.stackIndex, true
}
