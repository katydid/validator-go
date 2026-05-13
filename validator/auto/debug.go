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

package auto

import (
	"fmt"
	"strings"

	"github.com/katydid/validator-go/validator/intern"
	"github.com/katydid/validator-go/validator/sets"
)

type debugger struct {
	patterns  PatternsSet
	zis       sets.Ints
	stackElms sets.Pairs
	nullables sets.BitsSet
}

func newDebugger(c *compiler) *debugger {
	return &debugger{c.patterns, c.zis, c.stackElms, c.nullables}
}

func (d *debugger) printState(prefix string, state int) {
	if d == nil {
		return
	}
	ps := d.patterns.Get(state)
	fmt.Printf("%s: %s\n", prefix, patternsString(ps))
}

func (d *debugger) printStackElm(stackElm int) {
	s := d.stackElms[stackElm]
	d.printState("stackElem", s.First)
}

func (d *debugger) printNulls(stackElm int, nullIndex int) {
	nullable := sets.UnzipBits(d.nullables[nullIndex], d.zis[d.stackElms[stackElm].Second])
	fmt.Printf("nulls: %v\n", nullable)
}

func patternsString(ps []*intern.Pattern) string {
	ss := make([]string, len(ps))
	for i := range ps {
		ss[i] = ps[i].String()
	}
	return strings.Join(ss, ",")
}
