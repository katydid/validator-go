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

package types

type Type int32

const (
	UNKNOWN       Type = 0
	SINGLE_DOUBLE Type = 101
	SINGLE_INT    Type = 103
	SINGLE_UINT   Type = 104
	SINGLE_BOOL   Type = 108
	SINGLE_STRING Type = 109
	SINGLE_BYTES  Type = 112
	LIST_DOUBLE   Type = 201
	LIST_INT      Type = 203
	LIST_UINT     Type = 204
	LIST_BOOL     Type = 208
	LIST_STRING   Type = 209
	LIST_BYTES    Type = 212
)

var Type_name = map[int32]string{
	0:   "UNKNOWN",
	101: "SINGLE_DOUBLE",
	103: "SINGLE_INT",
	104: "SINGLE_UINT",
	108: "SINGLE_BOOL",
	109: "SINGLE_STRING",
	112: "SINGLE_BYTES",
	201: "LIST_DOUBLE",
	203: "LIST_INT",
	204: "LIST_UINT",
	208: "LIST_BOOL",
	209: "LIST_STRING",
	212: "LIST_BYTES",
}

var Type_value = map[string]int32{
	"UNKNOWN":       0,
	"SINGLE_DOUBLE": 101,
	"SINGLE_INT":    103,
	"SINGLE_UINT":   104,
	"SINGLE_BOOL":   108,
	"SINGLE_STRING": 109,
	"SINGLE_BYTES":  112,
	"LIST_DOUBLE":   201,
	"LIST_INT":      203,
	"LIST_UINT":     204,
	"LIST_BOOL":     208,
	"LIST_STRING":   209,
	"LIST_BYTES":    212,
}

func (t Type) String() string {
	return Type_name[int32(t)]
}
