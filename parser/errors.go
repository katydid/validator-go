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

package parser

import (
	"fmt"
)

//ErrNotDouble is an error that represents a type error.
var ErrNotDouble = fmt.Errorf("value is not a double")

//ErrNotInt is an error that represents a type error.
var ErrNotInt = fmt.Errorf("value is not a int")

//ErrNotUint is an error that represents a type error.
var ErrNotUint = fmt.Errorf("value is not a uint")

//ErrNotBool is an error that represents a type error.
var ErrNotBool = fmt.Errorf("value is not a bool")

//ErrNotString is an error that represents a type error.
var ErrNotString = fmt.Errorf("value is not a string")

//ErrNotBytes is an error that represents a type error.
var ErrNotBytes = fmt.Errorf("value is not a bytes")
