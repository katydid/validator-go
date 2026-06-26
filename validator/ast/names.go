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

package ast

func (n *NameExpr) GetConstNames() map[string]struct{} {
	if n.Name != nil {
		return n.Name.GetConstNames()
	}
	if n.AnyNameExcept != nil {
		return n.AnyNameExcept.GetConstNames()
	}
	if n.NameChoice != nil {
		return n.NameChoice.GetConstNames()
	}
	if n.NameConj != nil {
		return n.NameConj.GetConstNames()
	}
	return nil
}

func (n *Name) GetConstNames() map[string]struct{} {
	if n.StringValue != nil {
		return map[string]struct{}{*n.StringValue: {}}
	}
	if n.TagValue != nil {
		return map[string]struct{}{*n.TagValue: {}}
	}
	return nil
}

func (n *AnyNameExcept) GetConstNames() map[string]struct{} {
	if n.Except != nil {
		return n.Except.GetConstNames()
	}
	return nil
}

func (n *NameChoice) GetConstNames() map[string]struct{} {
	res := make(map[string]struct{})
	if n.Left != nil {
		for name := range n.Left.GetConstNames() {
			res[name] = struct{}{}
		}
	}
	if n.Right != nil {
		for name := range n.Right.GetConstNames() {
			res[name] = struct{}{}
		}
	}
	return res
}

func (n *NameConj) GetConstNames() map[string]struct{} {
	res := make(map[string]struct{})
	if n.Left != nil {
		for name := range n.Left.GetConstNames() {
			res[name] = struct{}{}
		}
	}
	if n.Right != nil {
		for name := range n.Right.GetConstNames() {
			res[name] = struct{}{}
		}
	}
	return res
}
