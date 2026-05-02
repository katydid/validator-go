//  Copyright 2013 Walter Schulze
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

package funcs

import (
	"math/big"

	"github.com/katydid/parser-go/cast"
	"github.com/katydid/parser-go/parse"
)

type aVariable interface {
	isVariable()
}

// Setter is an interface that represents a variable of which the value must be set.
type Setter interface {
	SetValue(parse.Token)
}

func (this *varDouble) Eval() (float64, error) {
	if this.Value == nil {
		return 0, ErrNotDoubleConst{}
	}
	kind, v, err := this.Value.Token()
	if err != nil {
		return 0, err
	}
	switch kind {
	case parse.Int64Kind:
		return float64(cast.ToInt64(v)), nil
	case parse.Float64Kind:
		return cast.ToFloat64(v), nil
	}
	return 0, ErrNotDoubleConst{}
}

func (this *varInt) Eval() (int64, error) {
	if this.Value == nil {
		return 0, ErrNotIntConst{}
	}
	kind, v, err := this.Value.Token()
	if err != nil {
		return 0, err
	}
	switch kind {
	case parse.Int64Kind:
		return cast.ToInt64(v), nil
	case parse.Float64Kind:
		return int64(cast.ToFloat64(v)), nil
	}
	return 0, ErrNotIntConst{}
}

func (this *varUint) Eval() (uint64, error) {
	if this.Value == nil {
		return 0, ErrNotUintConst{}
	}
	kind, v, err := this.Value.Token()
	if err != nil {
		return 0, err
	}
	if kind != parse.DecimalKind {
		return 0, ErrNotUintConst{}
	}
	s := cast.ToString(v)
	bigfloat, _, err := big.ParseFloat(s, 10, 200, big.ToNearestAway)
	if err != nil {
		return 0, ErrNotUintConst{}
	}
	u, acc := bigfloat.Uint64()
	if acc != big.Exact {
		return 0, ErrNotUintConst{}
	}
	return u, nil
}

func (this *varBool) Eval() (bool, error) {
	if this.Value == nil {
		return false, ErrNotBoolConst{}
	}
	kind, _, err := this.Value.Token()
	if err != nil {
		return false, err
	}
	if kind == parse.TrueKind {
		return true, nil
	} else if kind == parse.FalseKind {
		return false, nil
	}
	return false, ErrNotBoolConst{}
}

func (this *varString) Eval() (string, error) {
	if this.Value == nil {
		return "", ErrNotStringConst{}
	}
	kind, v, err := this.Value.Token()
	if err != nil {
		return "", err
	}
	if kind != parse.StringKind {
		return "", ErrNotStringConst{}
	}
	return cast.ToString(v), nil
}

func (this *varBytes) Eval() ([]byte, error) {
	if this.Value == nil {
		return nil, ErrNotBytesConst{}
	}
	kind, v, err := this.Value.Token()
	if err != nil {
		return nil, err
	}
	if kind != parse.StringKind {
		return nil, ErrNotBytesConst{}
	}
	return v, nil
}
