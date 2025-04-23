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

func isEmpty(p *Pattern) bool {
	return p.Type == Empty
}

func isNotZAny(p *Pattern) bool {
	return p.Type == Not && p.Patterns[0].Type == ZAny
}

func isZAny(p *Pattern) bool {
	return p.Type == ZAny
}

func notEmptySet(p *Pattern) bool {
	return !(p.Type == Not && p.Patterns[0].Type == ZAny)
}

func notEmpty(p *Pattern) bool {
	return !(p.Type == Empty)
}

func isNullable(p *Pattern) bool {
	return p.nullable
}
