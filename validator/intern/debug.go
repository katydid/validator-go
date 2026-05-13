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

package intern

import (
	"fmt"
	"strings"
)

func printState(prefix string, state []*Pattern) {
	fmt.Printf("%s: %s\n", prefix, patternsString(state))
}

func patternsString(ps []*Pattern) string {
	ss := make([]string, len(ps))
	for i := range ps {
		ss[i] = ps[i].String()
	}
	return strings.Join(ss, ",")
}
