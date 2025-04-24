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

import (
	"testing"

	"github.com/katydid/parser-go-json/json"
)

func TestExtensionXOR(t *testing.T) {
	nullable := func(bs []bool) bool {
		allTrue := areAll(func(b bool) bool { return b == true }, bs)
		allFalse := areAll(func(b bool) bool { return b == false }, bs)
		return !(allTrue || allFalse)
	}
	derive := func(c Construct, d func(*Pattern) (*Pattern, error), ps []*Pattern) (*Pattern, error) {
		if len(ps) == 0 {
			return nil, nil
		}
		if len(ps) == 1 {
			return d(ps[0])
		}
		dps, err := traverse(ps, d)
		if err != nil {
			return nil, err
		}
		return c.NewExtension("xor", dps)
	}
	RegisterExtension("xor", nullable, derive)

	t.Run("alwaysTrue:$xor(!(*),*)", func(t *testing.T) {
		input := `{"a": 1}`
		j := json.NewParser()
		if err := j.Init([]byte(input)); err != nil {
			t.Fatal(err)
		}

		c := NewConstructorOptimizedForRecords()
		main, err := c.NewExtension("xor", []*Pattern{c.NewNotZAny(), c.NewZAny()})
		if err != nil {
			t.Fatal(err)
		}
		match, err := derivs(c, main, j)
		if err != nil {
			t.Fatal(err)
		}
		if !match {
			t.Fatal("expected alwaysTrue xor to match")
		}
	})

	t.Run("alwaysFalse:$xor(*,*)", func(t *testing.T) {
		input := `{"a": 1}`
		j := json.NewParser()
		if err := j.Init([]byte(input)); err != nil {
			t.Fatal(err)
		}

		c := NewConstructorOptimizedForRecords()
		main, err := c.NewExtension("xor", []*Pattern{c.NewZAny(), c.NewZAny()})
		if err != nil {
			t.Fatal(err)
		}
		match, err := derivs(c, main, j)
		if err != nil {
			t.Fatal(err)
		}
		if match {
			t.Fatal("expected alwaysFalse xor to not match")
		}
	})
}
