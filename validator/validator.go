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

// Package validator contains the validation language and the functions necessary for running it.
// See katydid.github.io for the language documentation.
package validator

import (
	"github.com/katydid/validator-go/parser"
	"github.com/katydid/validator-go/validator/ast"
	"github.com/katydid/validator-go/validator/mem"
	validatorparser "github.com/katydid/validator-go/validator/parser"
)

// Parse parses the validator string into an ast (abstract syntax tree)
func Parse(validator string) (*ast.Grammar, error) {
	return validatorparser.ParseGrammar(validator)
}

// Prepare creates a memoizing object given the grammar.
// The memoizing object is used to memorize any previous states created from previous validations.
// This results in a more efficient execution each time the memoizing object is used to validate a parser.
func Prepare(g *ast.Grammar) (*mem.Mem, error) {
	return mem.New(g)
}

// Validate validates the parser with the given memoizing object, containing the grammar, for efficiency.
func Validate(m *mem.Mem, p parser.Interface) (bool, error) {
	return m.Validate(p)
}
