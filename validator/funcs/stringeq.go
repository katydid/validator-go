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

package funcs

func GetStringEqConst(f Bool) (string, bool) {
	ff, ok := f.(*stringEq)
	if !ok {
		return "", false
	}
	if _, aok := ff.V1.(aConst); aok {
		if _, bok := ff.V2.(aVariable); bok {
			c, err := ff.V1.Eval()
			if err != nil {
				return "", false
			}
			return c, true
		}
	} else if _, bok := ff.V2.(aConst); bok {
		if _, aok := ff.V1.(aVariable); aok {
			c, err := ff.V2.Eval()
			if err != nil {
				return "", false
			}
			return c, true
		}
	}
	return "", false
}
