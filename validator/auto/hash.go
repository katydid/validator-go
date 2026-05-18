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
	"github.com/katydid/parser-go/parse/debug"
	"github.com/katydid/validator-go/validator/ast"
	"github.com/katydid/validator-go/validator/intern"
)

type stringCollector struct {
	ss []string
}

func (c *stringCollector) Visit(node any) any {
	term, ok := node.(*ast.Terminal)
	if ok {
		if term.StringValue != nil {
			c.ss = append(c.ss, *term.StringValue)
		}
	}
	return c
}

func getStringsFromPattern(p *intern.Pattern) []string {
	if p.Func == nil {
		return nil
	}
	expr := p.Func.ToExpr()
	collector := &stringCollector{[]string{}}
	expr.Walk(collector)
	return collector.ss
}

func getStringsFromPatterns(ps []*intern.Pattern) []string {
	strings := []string{}
	for i := range ps {
		strings = append(strings, getStringsFromPattern(ps[i])...)
	}
	return strings
}

type callResult struct {
	child      int
	stackIndex int
}

func (this *compiler) calcHashCalls(state int) error {
	ps := this.patterns.Get(state)
	names := getStringsFromPatterns(ps)
	for len(this.hashedCalls) <= state {
		this.hashedCalls = append(this.hashedCalls, map[string]callResult{})
	}
	for _, name := range names {
		child, stackIndex, err := this.calls[state].eval(debug.NewStringValue(name))
		if err != nil {
			return err
		}
		this.hashedCalls[state][name] = callResult{child: child, stackIndex: stackIndex}
	}
	return nil
}
