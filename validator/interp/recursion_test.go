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

package interp

import (
	"testing"

	"github.com/katydid/validator-go/validator/ast"
	. "github.com/katydid/validator-go/validator/combinator"
)

func TestRecursionPositive(t *testing.T) {
	a := In("A", Any()) // .A:*
	b := In("B", Any()) // .B:*
	positives := []*ast.Grammar{
		// @main
		{TopPattern: ast.NewReference("main")},
		// #main = @main
		{PatternDecls: []*ast.PatternDecl{ast.NewPatternDecl("main", ast.NewReference("main"))}},
		// #main = (.A:* | @main)
		{PatternDecls: []*ast.PatternDecl{ast.NewPatternDecl("main", ast.NewOr(a, ast.NewReference("main")))}},
		// 	#main = (A:* | @ref)
		// 	#ref = @main
		{PatternDecls: []*ast.PatternDecl{
			ast.NewPatternDecl("main", ast.NewOr(a, ast.NewReference("ref"))),
			ast.NewPatternDecl("ref", ast.NewReference("main")),
		}},
		// #main = (A:* | @ref)
		// #ref = (B:* | @main | <empty>)
		{PatternDecls: []*ast.PatternDecl{
			ast.NewPatternDecl("main", ast.NewOr(a, ast.NewReference("ref"))),
			ast.NewPatternDecl("ref", ast.NewOr(b, ast.NewReference("main"), ast.NewEmpty())),
		}},
		// #main = A:@ref
		// #ref = @ref
		{PatternDecls: []*ast.PatternDecl{
			ast.NewPatternDecl("main", ast.NewOr(a, ast.NewReference("ref"))),
			ast.NewPatternDecl("ref", ast.NewReference("ref")),
		}},
	}
	for _, g := range positives {
		if !HasRecursion(g) {
			t.Errorf("expected recursion for %v", g)
		}
	}
}

func TestRecursionNegative(t *testing.T) {
	negatives := []*ast.Grammar{
		// .A:@main
		{TopPattern: In("A", ast.NewReference("main"))},
		// #main = .A:@main
		{PatternDecls: []*ast.PatternDecl{ast.NewPatternDecl("main", In("A", ast.NewReference("main")))}},
		// #main = (.A:* | .B:@main)
		{PatternDecls: []*ast.PatternDecl{ast.NewPatternDecl("main", ast.NewOr(In("A", Any()), In("B", ast.NewReference("main"))))}},
		// #main = (.A:* | @ref)
		// #ref = .C:@main
		{PatternDecls: []*ast.PatternDecl{
			ast.NewPatternDecl("main", ast.NewOr(In("A", Any()), ast.NewReference("ref"))),
			ast.NewPatternDecl("ref", In("C", ast.NewReference("main"))),
		}},
		// #main = (.A:* | @ref)
		// #ref = (.B:* | .D:@main | <empty>)
		{PatternDecls: []*ast.PatternDecl{
			ast.NewPatternDecl("main", ast.NewOr(In("A", Any()), ast.NewReference("ref"))),
			ast.NewPatternDecl("ref", ast.NewOr(In("B", Any(), In("D", ast.NewReference("main"), ast.NewEmpty())))),
		}},
	}
	for _, g := range negatives {
		if HasRecursion(g) {
			t.Errorf("unexpected recursion for %v", g)
		}
	}
}
