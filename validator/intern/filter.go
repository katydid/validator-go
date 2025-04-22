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

func removeAllNotZAny(ps []*Pattern) []*Pattern {
	return filter(notEmptySet, ps)
}

func removeAllZAny(ps []*Pattern) []*Pattern {
	return filter(func(p *Pattern) bool {
		return p.Type != ZAny
	}, ps)
}

func removeNotZAnyExceptOne(ps []*Pattern) []*Pattern {
	pps := filter(notEmptySet, ps)
	if len(pps) == 0 {
		return ps[:1]
	}
	return pps
}

func removeEmptyExceptOne(ps []*Pattern) []*Pattern {
	pps := filter(notEmpty, ps)
	if len(pps) == 0 {
		return ps[:1]
	}
	return pps
}

func removeZAnyExceptOne(ps []*Pattern) []*Pattern {
	pps := filter(func(p *Pattern) bool {
		return p.Type != ZAny
	}, ps)
	if len(pps) == 0 {
		return ps[:1]
	}
	return pps
}
