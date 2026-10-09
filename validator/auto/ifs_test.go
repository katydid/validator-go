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
	"katydid.org.za/go/parser-go/parse"
	"katydid.org.za/go/validator-go/validator/compose"
	"katydid.org.za/go/validator-go/validator/intern"
)

func (this *ifExprs) eval(label parse.Token) ([]*intern.Pattern, error) {
	if this.ret != nil {
		return this.ret, nil
	}
	composed, err := compose.NewBoolFunc(this.cond)
	if err != nil {
		return nil, err
	}
	cond, err := composed.Eval(label)
	if err != nil {
		return nil, err
	}
	if cond {
		return this.then.eval(label)
	}
	return this.els.eval(label)
}
