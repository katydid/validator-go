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
	"math"
	"math/big"

	"katydid.org.za/go/parser-go/cast"
	"katydid.org.za/go/parser-go/parse"
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
		var i64 int64
		// TODO: Consider not supporting Int64Kind here
		cast.ToInt64Ptr(v, &i64)
		return float64(i64), nil
	case parse.Float64Kind:
		var u64 uint64
		cast.ToFloat64BitsPtr(v, &u64)
		return math.Float64frombits(u64), nil
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
	case parse.NanosecondsKind:
		var i64 int64
		cast.ToInt64Ptr(v, &i64)
		return i64, nil
	case parse.Int64Kind:
		var i64 int64
		cast.ToInt64Ptr(v, &i64)
		return i64, nil
	case parse.Float64Kind:
		// TODO: Consider not supporting Float64Kind here
		var u uint64
		cast.ToFloat64BitsPtr(v, &u)
		return int64(math.Float64frombits(u)), nil
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
	switch kind {
	case parse.DecimalKind:
		var s string
		cast.ToStringPtr(v, &s)
		bigfloat, _, err := big.ParseFloat(s, 10, 200, big.ToNearestAway)
		if err != nil {
			return 0, ErrNotUintConst{}
		}
		u, acc := bigfloat.Uint64()
		if acc != big.Exact {
			return 0, ErrNotUintConst{}
		}
		return u, nil
	case parse.Int64Kind:
		var i int64
		cast.ToInt64Ptr(v, &i)
		if i < 0 {
			return 0, ErrNotUintConst{}
		}
		return uint64(i), nil
	}
	// TODO consider support Float64Kind and NanosecondsKind
	return 0, ErrNotUintConst{}
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
	switch kind {
	case parse.StringKind:
		var s string
		cast.ToStringPtr(v, &s)
		return s, nil
	}
	return "", ErrNotStringConst{}
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

func (this *varTag) Eval() (string, error) {
	if this.Value == nil {
		return "", ErrNotTagConst{}
	}
	kind, v, err := this.Value.Token()
	if err != nil {
		return "", err
	}
	switch kind {
	case parse.TagKind:
		var s string
		cast.ToStringPtr(v, &s)
		return s, nil
	}
	return "", ErrNotTagConst{}
}
