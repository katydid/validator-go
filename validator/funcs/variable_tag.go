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

package funcs

import (
	"github.com/katydid/parser-go/parse"
	"github.com/katydid/validator-go/validator/ast"
)

type varTag struct {
	Value parse.Token
	hash  uint64
}

var _ Setter = &varTag{}
var _ aVariable = &varTag{}

type ErrNotTagConst struct{}

func (this ErrNotTagConst) Error() string {
	return "$tag is not a const"
}

func (this *varTag) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if _, ok := that.(*varTag); ok {
		return 0
	}
	return this.ToExpr().Compare(that.ToExpr())
}

func (this *varTag) Hash() uint64 {
	return this.hash
}

func (this *varTag) HasVariable() bool { return true }

func (this *varTag) isVariable() {}

func (this *varTag) SetValue(v parse.Token) {
	this.Value = v
}

func (this *varTag) ToExpr() *ast.Expr {
	return ast.NewTagVar()
}

// TagVar returns a variable of type Tag
func TagVar() *varTag {
	h := uint64(17)
	h = 31*h + 2486848561
	return &varTag{hash: h}
}
