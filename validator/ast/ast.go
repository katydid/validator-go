//  Copyright 2024 Walter Schulze
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

package ast

import types "github.com/katydid/validator-go/validator/types"

// Grammar is the ast node representing the whole grammar.
type Grammar struct {
	TopPattern   *Pattern       `json:"TopPattern,omitempty"`
	PatternDecls []*PatternDecl `json:"PatternDecls,omitempty"`
	After        *Space         `json:"After,omitempty"`
}

func (m *Grammar) GetTopPattern() *Pattern {
	if m != nil {
		return m.TopPattern
	}
	return nil
}

func (m *Grammar) GetPatternDecls() []*PatternDecl {
	if m != nil {
		return m.PatternDecls
	}
	return nil
}

// PatternDecl is the ast node for the declaration of a pattern.
type PatternDecl struct {
	Hash    *Keyword `json:"Hash,omitempty"`
	Before  *Space   `json:"Before,omitempty"`
	Name    string   `json:"Name"`
	Eq      *Keyword `json:"Eq,omitempty"`
	Pattern *Pattern `json:"Pattern,omitempty"`
}

func (m *PatternDecl) GetHash() *Keyword {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PatternDecl) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PatternDecl) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

// Pattern is the ast node for the union of all possible patterns.
type Pattern struct {
	Empty      *Empty      `json:"Empty,omitempty"`
	TreeNode   *TreeNode   `json:"TreeNode,omitempty"`
	LeafNode   *LeafNode   `json:"LeafNode,omitempty"`
	Concat     *Concat     `json:"Concat,omitempty"`
	Or         *Or         `json:"Or,omitempty"`
	And        *And        `json:"And,omitempty"`
	ZeroOrMore *ZeroOrMore `json:"ZeroOrMore,omitempty"`
	Reference  *Reference  `json:"Reference,omitempty"`
	Not        *Not        `json:"Not,omitempty"`
	ZAny       *ZAny       `json:"ZAny,omitempty"`
	Contains   *Contains   `json:"Contains,omitempty"`
	Optional   *Optional   `json:"Optional,omitempty"`
	Interleave *Interleave `json:"Interleave,omitempty"`
}

func (m *Pattern) GetEmpty() *Empty {
	if m != nil {
		return m.Empty
	}
	return nil
}

func (m *Pattern) GetTreeNode() *TreeNode {
	if m != nil {
		return m.TreeNode
	}
	return nil
}

func (m *Pattern) GetLeafNode() *LeafNode {
	if m != nil {
		return m.LeafNode
	}
	return nil
}

func (m *Pattern) GetConcat() *Concat {
	if m != nil {
		return m.Concat
	}
	return nil
}

func (m *Pattern) GetOr() *Or {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *Pattern) GetAnd() *And {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *Pattern) GetZeroOrMore() *ZeroOrMore {
	if m != nil {
		return m.ZeroOrMore
	}
	return nil
}

func (m *Pattern) GetReference() *Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *Pattern) GetNot() *Not {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *Pattern) GetZAny() *ZAny {
	if m != nil {
		return m.ZAny
	}
	return nil
}

func (m *Pattern) GetContains() *Contains {
	if m != nil {
		return m.Contains
	}
	return nil
}

func (m *Pattern) GetOptional() *Optional {
	if m != nil {
		return m.Optional
	}
	return nil
}

func (m *Pattern) GetInterleave() *Interleave {
	if m != nil {
		return m.Interleave
	}
	return nil
}

func (this *Pattern) GetValue() interface{} {
	if this.Empty != nil {
		return this.Empty
	}
	if this.TreeNode != nil {
		return this.TreeNode
	}
	if this.LeafNode != nil {
		return this.LeafNode
	}
	if this.Concat != nil {
		return this.Concat
	}
	if this.Or != nil {
		return this.Or
	}
	if this.And != nil {
		return this.And
	}
	if this.ZeroOrMore != nil {
		return this.ZeroOrMore
	}
	if this.Reference != nil {
		return this.Reference
	}
	if this.Not != nil {
		return this.Not
	}
	if this.ZAny != nil {
		return this.ZAny
	}
	if this.Contains != nil {
		return this.Contains
	}
	if this.Optional != nil {
		return this.Optional
	}
	if this.Interleave != nil {
		return this.Interleave
	}
	return nil
}

// Empty is the ast node for the Empty pattern.
type Empty struct {
	Empty *Keyword `json:"Empty,omitempty"`
}

// TreeNode is the ast node for the TreeNode pattern.
type TreeNode struct {
	Name    *NameExpr `json:"Name,omitempty"`
	Colon   *Keyword  `json:"Colon,omitempty"`
	Pattern *Pattern  `json:"Pattern,omitempty"`
}

func (m *TreeNode) GetName() *NameExpr {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *TreeNode) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

// Contains is the ast node for the Contains pattern.
type Contains struct {
	Dot     *Keyword `json:"Dot,omitempty"`
	Pattern *Pattern `json:"Pattern,omitempty"`
}

func (m *Contains) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

// LeafNode is the ast node for the LeafNode pattern.
type LeafNode struct {
	Expr *Expr `json:"Expr,omitempty"`
}

func (m *LeafNode) GetExpr() *Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

// Concat is the ast node for the Concat pattern.
type Concat struct {
	OpenBracket  *Keyword `json:"OpenBracket,omitempty"`
	LeftPattern  *Pattern `json:"LeftPattern,omitempty"`
	Comma        *Keyword `json:"Comma,omitempty"`
	RightPattern *Pattern `json:"RightPattern,omitempty"`
	ExtraComma   *Keyword `json:"ExtraComma,omitempty"`
	CloseBracket *Keyword `json:"CloseBracket,omitempty"`
}

func (m *Concat) GetLeftPattern() *Pattern {
	if m != nil {
		return m.LeftPattern
	}
	return nil
}

func (m *Concat) GetRightPattern() *Pattern {
	if m != nil {
		return m.RightPattern
	}
	return nil
}

// Or is the ast node for the Or pattern.
type Or struct {
	OpenParen    *Keyword `json:"OpenParen,omitempty"`
	LeftPattern  *Pattern `json:"LeftPattern,omitempty"`
	Pipe         *Keyword `json:"Pipe,omitempty"`
	RightPattern *Pattern `json:"RightPattern,omitempty"`
	CloseParen   *Keyword `json:"CloseParen,omitempty"`
}

func (m *Or) GetLeftPattern() *Pattern {
	if m != nil {
		return m.LeftPattern
	}
	return nil
}

func (m *Or) GetRightPattern() *Pattern {
	if m != nil {
		return m.RightPattern
	}
	return nil
}

// And is the ast node for the And pattern.
type And struct {
	OpenParen    *Keyword `json:"OpenParen,omitempty"`
	LeftPattern  *Pattern `json:"LeftPattern,omitempty"`
	Ampersand    *Keyword `json:"Ampersand,omitempty"`
	RightPattern *Pattern `json:"RightPattern,omitempty"`
	CloseParen   *Keyword `json:"CloseParen,omitempty"`
}

func (m *And) GetLeftPattern() *Pattern {
	if m != nil {
		return m.LeftPattern
	}
	return nil
}

func (m *And) GetRightPattern() *Pattern {
	if m != nil {
		return m.RightPattern
	}
	return nil
}

// ZeroOrMore is the ast node for the ZeroOrMore pattern.
type ZeroOrMore struct {
	OpenParen  *Keyword `json:"OpenParen,omitempty"`
	Pattern    *Pattern `json:"Pattern,omitempty"`
	CloseParen *Keyword `json:"CloseParen,omitempty"`
	Star       *Keyword `json:"Star,omitempty"`
}

func (m *ZeroOrMore) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

// Reference is the ast node for the Reference pattern.
type Reference struct {
	At   *Keyword `json:"At,omitempty"`
	Name string   `json:"Name"`
}

func (m *Reference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Not is the ast node for the Not pattern.
type Not struct {
	Exclamation *Keyword `json:"Exclamation,omitempty"`
	OpenParen   *Keyword `json:"OpenParen,omitempty"`
	Pattern     *Pattern `json:"Pattern,omitempty"`
	CloseParen  *Keyword `json:"CloseParen,omitempty"`
}

func (m *Not) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

// ZAny is the ast node for the ZAny pattern.
type ZAny struct {
	Star *Keyword `json:"Star,omitempty"`
}

// Optional is the ast node for the Optional pattern.
type Optional struct {
	OpenParen    *Keyword `json:"OpenParen,omitempty"`
	Pattern      *Pattern `json:"Pattern,omitempty"`
	CloseParen   *Keyword `json:"CloseParen,omitempty"`
	QuestionMark *Keyword `json:"QuestionMark,omitempty"`
}

func (m *Optional) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

// Interleave is the ast node for the Interleave pattern.
type Interleave struct {
	OpenCurly      *Keyword `json:"OpenCurly,omitempty"`
	LeftPattern    *Pattern `json:"LeftPattern,omitempty"`
	SemiColon      *Keyword `json:"SemiColon,omitempty"`
	RightPattern   *Pattern `json:"RightPattern,omitempty"`
	ExtraSemiColon *Keyword `json:"ExtraSemiColon,omitempty"`
	CloseCurly     *Keyword `json:"CloseCurly,omitempty"`
}

func (m *Interleave) GetLeftPattern() *Pattern {
	if m != nil {
		return m.LeftPattern
	}
	return nil
}

func (m *Interleave) GetRightPattern() *Pattern {
	if m != nil {
		return m.RightPattern
	}
	return nil
}

// Expr is a union of all possible expression types, terminal, list, function and builtin function.
type Expr struct {
	RightArrow *Keyword  `json:"RightArrow,omitempty"`
	Comma      *Keyword  `json:"Comma,omitempty"`
	Terminal   *Terminal `json:"Terminal,omitempty"`
	List       *List     `json:"List,omitempty"`
	Function   *Function `json:"Function,omitempty"`
	BuiltIn    *BuiltIn  `json:"BuiltIn,omitempty"`
}

func (m *Expr) GetTerminal() *Terminal {
	if m != nil {
		return m.Terminal
	}
	return nil
}

func (m *Expr) GetList() *List {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *Expr) GetFunction() *Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *Expr) GetBuiltIn() *BuiltIn {
	if m != nil {
		return m.BuiltIn
	}
	return nil
}

// List is an expression that represents a typed list of expressions.
type List struct {
	Before     *Space     `json:"Before,omitempty"`
	Type       types.Type `json:"Type"`
	OpenCurly  *Keyword   `json:"OpenCurly,omitempty"`
	Elems      []*Expr    `json:"Elems,omitempty"`
	CloseCurly *Keyword   `json:"CloseCurly,omitempty"`
}

func (m *List) GetType() types.Type {
	if m != nil {
		return m.Type
	}
	return types.UNKNOWN
}

func (m *List) GetElems() []*Expr {
	if m != nil {
		return m.Elems
	}
	return nil
}

// Function is an expression that represents a function expression, which contains a function name and a list parameters.
type Function struct {
	Before     *Space   `json:"Before,omitempty"`
	Name       string   `json:"Name"`
	OpenParen  *Keyword `json:"OpenParen,omitempty"`
	Params     []*Expr  `json:"Params,omitempty"`
	CloseParen *Keyword `json:"CloseParen,omitempty"`
}

func (m *Function) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Function) GetParams() []*Expr {
	if m != nil {
		return m.Params
	}
	return nil
}

// BuiltIn is an expression that represents a builtin function.  This is represented by a symbol and an expression.
type BuiltIn struct {
	Symbol *Keyword `json:"Symbol,omitempty"`
	Expr   *Expr    `json:"Expr,omitempty"`
}

func (m *BuiltIn) GetSymbol() *Keyword {
	if m != nil {
		return m.Symbol
	}
	return nil
}

func (m *BuiltIn) GetExpr() *Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

// Terminal is an expression that represents a literal value or variable.
type Terminal struct {
	Before      *Space    `json:"Before,omitempty"`
	DoubleValue *float64  `json:"DoubleValue,omitempty"`
	IntValue    *int64    `json:"IntValue,omitempty"`
	UintValue   *uint64   `json:"UintValue,omitempty"`
	BoolValue   *bool     `json:"BoolValue,omitempty"`
	StringValue *string   `json:"StringValue,omitempty"`
	BytesValue  []byte    `json:"BytesValue,omitempty"`
	Variable    *Variable `json:"Variable,omitempty"`
}

func (m *Terminal) GetDoubleValue() float64 {
	if m != nil && m.DoubleValue != nil {
		return *m.DoubleValue
	}
	return 0
}

func (m *Terminal) GetIntValue() int64 {
	if m != nil && m.IntValue != nil {
		return *m.IntValue
	}
	return 0
}

func (m *Terminal) GetUintValue() uint64 {
	if m != nil && m.UintValue != nil {
		return *m.UintValue
	}
	return 0
}

func (m *Terminal) GetBoolValue() bool {
	if m != nil && m.BoolValue != nil {
		return *m.BoolValue
	}
	return false
}

func (m *Terminal) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Terminal) GetBytesValue() []byte {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

func (m *Terminal) GetVariable() *Variable {
	if m != nil {
		return m.Variable
	}
	return nil
}

// Variable is a terminal.
type Variable struct {
	Type types.Type `json:"Type"`
}

func (m *Variable) GetType() types.Type {
	if m != nil {
		return m.Type
	}
	return types.UNKNOWN
}

// Keyword represents any possible keyword.
type Keyword struct {
	Before *Space `json:"Before,omitempty"`
	Value  string `json:"Value"`
}

func (m *Keyword) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Space represents a comment or white space.
type Space struct {
	Space []string `json:"Space,omitempty"`
}

// NameExpr is a special type of expression for field names that contains a union of all the possible name expressions.
type NameExpr struct {
	Name          *Name          `json:"Name,omitempty"`
	AnyName       *AnyName       `json:"AnyName,omitempty"`
	AnyNameExcept *AnyNameExcept `json:"AnyNameExcept,omitempty"`
	NameChoice    *NameChoice    `json:"NameChoice,omitempty"`
}

func (m *NameExpr) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *NameExpr) GetAnyName() *AnyName {
	if m != nil {
		return m.AnyName
	}
	return nil
}

func (m *NameExpr) GetAnyNameExcept() *AnyNameExcept {
	if m != nil {
		return m.AnyNameExcept
	}
	return nil
}

func (m *NameExpr) GetNameChoice() *NameChoice {
	if m != nil {
		return m.NameChoice
	}
	return nil
}

func (this *NameExpr) GetValue() interface{} {
	if this.Name != nil {
		return this.Name
	}
	if this.AnyName != nil {
		return this.AnyName
	}
	if this.AnyNameExcept != nil {
		return this.AnyNameExcept
	}
	if this.NameChoice != nil {
		return this.NameChoice
	}
	return nil
}

// Name is a name expression and represents a union of all possible name type values.
type Name struct {
	Before      *Space   `json:"Before,omitempty"`
	DoubleValue *float64 `json:"DoubleValue,omitempty"`
	IntValue    *int64   `json:"IntValue,omitempty"`
	UintValue   *uint64  `json:"UintValue,omitempty"`
	BoolValue   *bool    `json:"BoolValue,omitempty"`
	StringValue *string  `json:"StringValue,omitempty"`
	BytesValue  []byte   `json:"BytesValue,omitempty"`
}

func (m *Name) GetDoubleValue() float64 {
	if m != nil && m.DoubleValue != nil {
		return *m.DoubleValue
	}
	return 0
}

func (m *Name) GetIntValue() int64 {
	if m != nil && m.IntValue != nil {
		return *m.IntValue
	}
	return 0
}

func (m *Name) GetUintValue() uint64 {
	if m != nil && m.UintValue != nil {
		return *m.UintValue
	}
	return 0
}

func (m *Name) GetBoolValue() bool {
	if m != nil && m.BoolValue != nil {
		return *m.BoolValue
	}
	return false
}

func (m *Name) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Name) GetBytesValue() []byte {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

// AnyName is a name expression that represents any name.
type AnyName struct {
	Underscore *Keyword `json:"Underscore,omitempty"`
}

// AnyNameExpr is a name expression that represents any name except the specified name expression.
type AnyNameExcept struct {
	Exclamation *Keyword  `json:"Exclamation,omitempty"`
	OpenParen   *Keyword  `json:"OpenParen,omitempty"`
	Except      *NameExpr `json:"Except,omitempty"`
	CloseParen  *Keyword  `json:"CloseParen,omitempty"`
}

func (m *AnyNameExcept) GetExcept() *NameExpr {
	if m != nil {
		return m.Except
	}
	return nil
}

// NameChoice is a name expression that represents a choice between two possible name expressions.
type NameChoice struct {
	OpenParen  *Keyword  `json:"OpenParen,omitempty"`
	Left       *NameExpr `json:"Left,omitempty"`
	Pipe       *Keyword  `json:"Pipe,omitempty"`
	Right      *NameExpr `json:"Right,omitempty"`
	CloseParen *Keyword  `json:"CloseParen,omitempty"`
}

func (m *NameChoice) GetLeft() *NameExpr {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *NameChoice) GetRight() *NameExpr {
	if m != nil {
		return m.Right
	}
	return nil
}
