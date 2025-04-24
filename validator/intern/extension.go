//  Copyright 2017 Walter Schulze
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

import "errors"

// RegisterExtension registers a new user defined operator.
// It returns true, if it overwrote an operator that was already registered with the same name.
func RegisterExtension(name string, nullable func([]bool) bool) bool {
	_, ok := extensionsRegister[name]
	extensionsRegister[name] = newExtension(name, nullable)
	return ok
}

var errExtensionNotFound = errors.New("unknown extension")

func GetExtension(name string) (*extension, error) {
	ext, ok := extensionsRegister[name]
	if !ok {
		return nil, errExtensionNotFound
	}
	return ext, nil
}

var extensionsRegister extensionMap = make(extensionMap)

type extensionMap map[string]*extension

type extension struct {
	name     string
	nullable func([]bool) bool
}

func newExtension(name string, nullable func([]bool) bool) *extension {
	return &extension{
		name:     name,
		nullable: nullable,
	}
}
