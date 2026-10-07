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

package main

import (
	"fmt"
	"strings"

	"katydid.org.za/go/parser-go/parse"
	"katydid.org.za/go/validator-go/validator/ast"
	"katydid.org.za/go/validator-go/validator/auto"
	"katydid.org.za/go/validator-go/validator/intern"
	"katydid.org.za/go/validator-go/validator/mem"
	"katydid.org.za/go/validator-go/validator/testsuite"
)

func main() {
	opts := []testsuite.Option{
		testsuite.Must(),
		testsuite.OnlyHedgeCodec(),
		testsuite.WithPath("../../../../src/katydid.org.za/go/validator-testsuite/validator/tests"),
	}
	exists, err := testsuite.TestSuiteExists(opts...)
	if !exists {
		if err != nil {
			panic(err)
		}
		panic("could not find testsuite")
	}
	tests, err := testsuite.ReadTestSuite(opts...)
	if err != nil {
		panic(err)
	}
	for _, testCase := range tests {
		testintern(testCase.Grammar, testCase.Parser, testCase.Expected, "", testCase.Record)
	}
	tests, err = testsuite.ReadTestSuite(opts...)
	if err != nil {
		panic(err)
	}
	for _, testCase := range tests {
		testauto(testCase.Name, testCase.Grammar, testCase.Parser, testCase.Expected, "", testCase.Record)
	}
	tests, err = testsuite.ReadTestSuite(opts...)
	if err != nil {
		panic(err)
	}
	for _, testCase := range tests {
		testmem(testCase.Grammar, testCase.Parser, testCase.Expected, "", testCase.Record)
	}
}

func testmem(g *ast.Grammar, p parse.Parser, expected bool, desc string, record bool) {
	if intern.HasRecursion(g) {
		// "was not designed to handle left recursion"
		return
	}
	var m *mem.Mem
	var err error
	if record {
		m, err = mem.New(g, mem.WithRecordSimplificationRules(), mem.WithFieldNameTable())
	} else {
		m, err = mem.New(g)
	}
	if err != nil {
		panic(err)
	}
	match, err := m.Validate(p)
	if err != nil {
		panic(err)
	}
	if match != expected {
		panic(fmt.Errorf("Expected %v on given \n%s\n on \n%s", expected, g.String(), desc))
	}
}

func testauto(name string, g *ast.Grammar, p parse.Parser, expected bool, desc string, record bool) {
	if intern.HasRecursion(g) {
		// "was not designed to handle left recursion"
		return
	}
	if strings.HasPrefix(name, "GoBigOr") {
		// This one seems to be working, maybe we can delete this?
		// too big to fail: the number of Ors creates a state space explosion
		return
	}
	if strings.HasPrefix(name, "BananaLarge") {
		// This one does not seem to exist anymore.
		// too big to fail: this test was specifically created to test that nested if expressions don't result in exponential explosions but since the auto implementation does compile everything it will explode.
		return
	}
	var a *auto.Auto
	var err error
	if record {
		a, err = auto.Compile(g, auto.WithRecordSimplificationRules(), auto.WithMaxBitSetSize(20), auto.WithFieldNameTable())
	} else {
		a, err = auto.Compile(g)
	}
	if err != nil {
		panic(err)
	}
	match, err := a.Validate(p)
	if err != nil {
		panic(err)
	}
	if match != expected {
		panic(fmt.Errorf("Expected %v on given \n%s\n on \n%s", expected, g.String(), desc))
	}
}

func testintern(g *ast.Grammar, p parse.Parser, expected bool, desc string, record bool) {
	if intern.HasRecursion(g) {
		// "was not designed to handle left recursion"
		return
	}
	match, err := intern.Interpret(g, record, p)
	if err != nil {
		panic(err)
	}
	if match != expected {
		panic(fmt.Errorf("Expected %v on given \n%s\n", expected, g.String()))
	}
}
