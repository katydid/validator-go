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

package auto_test

import (
	"testing"

	"github.com/katydid/validator-go/validator/auto"
	c "github.com/katydid/validator-go/validator/combinator"
	"github.com/katydid/validator-go/validator/parser"
)

func TestExplosionAndSameTree(t *testing.T) {
	var input = `(
		.A:.B:.DeepLevel:.DeeperLevel:.DeepestLevel:->contains($string,"el") &
		(
			.A:.B:.Rs:._:->eq("~a",$string) &
			(
				.A:.B:.NumI32:->contains($int,[]int{28,1,52}) &
				(
					.A:.B:.NumI64:>= 1 &
					(
						.A:.B:.NumU32:<= uint(4) &
						(
							.A:.B:.NumU64:== uint(4) &
							(
								.A:.B:.YesNo:== true &
								(
									.A:.B:.BS:== []byte{0x3, 0x2, 0x1, 0x0} &
									.A:.B:.Uuid: == []byte{0x3, 0x2, 0x1, 0x0}
								)
							)
						)
					)
				)
			)
		)
	)`
	g, err := parser.ParseGrammar(input)
	if err != nil {
		t.Fatal(err)
	}
	// This one causes a state explosion of over 14000 states.
	// Since we know field names can't repeat the simplification can be made for record (JSON and proto) like serialization formats, but not for XML.
	t.Logf("%v", g)
	// CompileRecord avoids the state explosion, whereas Compile does not do the Record simplifications, which results in the state space explosion.
	autoRecord, err := auto.CompileRecord(g)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("number of states %d", autoRecord.MetricNumberOfStates())
	if autoRecord.MetricNumberOfStates() > 1000 {
		t.Fatal("number of states exploded")
	}

	// even with the state explosion it should still be able to compile,
	// but it is slow, so we comment out this part of the test.
	//   autoNoRecord, err := auto.Compile(g)
	//   if err != nil {
	//     t.Fatal(err)
	//   }
	//   t.Logf("number of states %d", autoNoRecord.MetricNumberOfStates())
	//   if autoNoRecord.MetricNumberOfStates() < 1000 {
	// 	   t.Fatal("number of states was expected to explode")
	//   }
}

func TestPersonNumStates(t *testing.T) {
	g := c.G{"main": c.In("Person", c.InPath("Addresses",
		c.Any(),
		c.In("Number", c.Value(c.Eq(c.IntVar(), c.IntConst(456)))),
		c.Any(),
		c.In("Street", c.Value(c.Eq(c.StringVar(), c.StringConst("TheStreet")))),
		c.Any(),
	))}.Grammar()
	auto, err := auto.Compile(g)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("number of states %d", auto.MetricNumberOfStates())
	// If this does not compile then we are missing simplification rules.
	// This test was inspired by a bug, where
	// we weren't flattening the Or before constructing another Or
	// Now we have flattenByType function at the start of all binary operators,
	// for example:
	//   func (c *construct) NewOr(ps []*Pattern) (*Pattern, error) {
	//	   ps = flattenByType(ps, Or)
}
