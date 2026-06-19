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

package intern

import (
	"fmt"
	"strings"

	"github.com/katydid/validator-go/validator/ast"
	"github.com/katydid/validator-go/validator/funcs"
)

type PatternType int

const Empty = PatternType(1)
const Node = PatternType(2)
const Concat = PatternType(4)
const Or = PatternType(5)
const And = PatternType(6)
const ZeroOrMore = PatternType(7)
const Reference = PatternType(8)
const Not = PatternType(9)
const ZAny = PatternType(10)
const Contains = PatternType(11)
const Optional = PatternType(12)
const Interleave = PatternType(13)
const Xor = PatternType(14)

type Pattern struct {
	Type     PatternType
	Func     funcs.Bool
	Patterns []*Pattern
	Ref      string
	hash     uint64
	nullable bool
	str      string
}

func newOpPattern(typ PatternType, ps ...*Pattern) *Pattern {
	p := &Pattern{
		Type:     typ,
		Patterns: ps,
	}
	switch typ {
	case Empty, ZAny:
		p.nullable = true
	case Or:
		p.nullable = anyNullable(ps)
	case Concat, And, Interleave:
		p.nullable = allNullable(ps)
	case ZeroOrMore, Optional:
		p.nullable = true
	case Not:
		p.nullable = !ps[0].nullable
	case Contains:
		p.nullable = ps[0].nullable
	case Xor:
		p.nullable = xorNullable(ps)
	default:
		panic(fmt.Sprintf("unsupported PatternType %v", typ))
	}
	p.hash = makeHash(p)
	p.str = p.String()
	return p
}

func newNodePattern(b funcs.Bool, child *Pattern) *Pattern {
	p := &Pattern{
		Type:     Node,
		Func:     b,
		Patterns: []*Pattern{child},
		nullable: false,
	}
	p.hash = makeHash(p)
	return p
}

func newRefPattern(name string, nullable bool) *Pattern {
	p := &Pattern{
		Type:     Reference,
		Ref:      name,
		nullable: nullable,
	}
	p.hash = makeHash(p)
	return p
}

func (p *Pattern) GoString() string {
	if p == nil {
		return "nil"
	}
	return fmt.Sprintf("&%#v", *p)
}

func writeStrings(sb *strings.Builder, patterns []*Pattern, sep string) {
	for i := range patterns {
		writeString(sb, patterns[i])
		if i != len(patterns)-1 {
			sb.WriteString(sep)
		}
	}
}

func writeString(sb *strings.Builder, p *Pattern) {
	if p.str != "" {
		sb.WriteString(p.str)
		return
	}
	switch p.Type {
	case Empty:
		sb.WriteString(ast.NewEmpty().String())
	case Node:
		if isEmpty(p.Patterns[0]) {
			sb.WriteString(p.Func.ToExpr().String())
		} else {
			sb.WriteString(p.Func.ToExpr().String())
			sb.WriteString(":")
			sb.WriteString(p.Patterns[0].String())
		}
	case Concat:
		sb.WriteString("[")
		writeStrings(sb, p.Patterns, ", ")
		sb.WriteString("]")
	case Or:
		sb.WriteString("(")
		writeStrings(sb, p.Patterns, " | ")
		sb.WriteString(")")
	case And:
		sb.WriteString("(")
		writeStrings(sb, p.Patterns, " & ")
		sb.WriteString(")")
	case ZeroOrMore:
		sb.WriteString("(")
		writeString(sb, p.Patterns[0])
		sb.WriteString(")*")
	case Reference:
		sb.WriteString("@")
		sb.WriteString(p.Ref)
	case Not:
		sb.WriteString("!(")
		writeString(sb, p.Patterns[0])
		sb.WriteString(")")
	case ZAny:
		sb.WriteString(ast.NewZAny().String())
	case Contains:
		sb.WriteString(".")
		writeString(sb, p.Patterns[0])
	case Optional:
		sb.WriteString("(")
		writeString(sb, p.Patterns[0])
		sb.WriteString(")?")
	case Interleave:
		sb.WriteString("{")
		writeStrings(sb, p.Patterns, "; ")
		sb.WriteString("}")
	case Xor:
		sb.WriteString("(")
		writeStrings(sb, p.Patterns, " ^ ")
		sb.WriteString(")")
	default:
		panic(fmt.Sprintf("unknown pattern: %d", p.Type))
	}
}

func (p *Pattern) String() string {
	if p == nil {
		return ""
	}
	if p.str != "" {
		return p.str
	}
	var sb strings.Builder
	writeString(&sb, p)
	s := sb.String()
	p.str = s
	return p.str
}

func (p *Pattern) Equal(pp *Pattern) bool {
	if p.hash != pp.hash {
		return false
	}
	if p.str != "" && pp.str != "" {
		return p.str == pp.str
	}
	if p.Type != pp.Type {
		return false
	}
	if p.Func == nil && pp.Func != nil {
		return false
	}
	if p.Func != nil && pp.Func == nil {
		return false
	}
	if p.Func != nil && pp.Func != nil {
		if !funcs.Equal(p.Func, pp.Func) {
			return false
		}
	}
	if len(p.Patterns) != len(pp.Patterns) {
		return false
	}
	for i := range p.Patterns {
		if !p.Patterns[i].Equal(pp.Patterns[i]) {
			return false
		}
	}
	if p.Ref != pp.Ref {
		return false
	}
	return true
}

func makeHash(p *Pattern) uint64 {
	h := uint64(17)
	h = 31*h + uint64(p.Type)
	if p.Func != nil {
		h = 31*h + p.Func.Hash()
	}
	for _, pattern := range p.Patterns {
		h = 31*h + pattern.hash
	}
	if len(p.Ref) > 0 {
		h = 31*h + deriveHashString(p.Ref)
	}
	if h == 0 {
		h = 1
	}
	return h
}

func (p *Pattern) Nullable() bool {
	return p.nullable
}
