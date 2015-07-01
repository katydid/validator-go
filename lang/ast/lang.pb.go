// Code generated by protoc-gen-gogo.
// source: lang.proto
// DO NOT EDIT!

/*
	Package lang is a generated protocol buffer package.

	It is generated from these files:
		lang.proto

	It has these top-level messages:
		Grammar
		PatternDecl
		NameExpr
		Name
		AnyName
		AnyNameExcept
		NameChoice
		Pattern
		Empty
		EmptySet
		TreeNode
		LeafNode
		Concat
		Or
		And
		ZeroOrMore
		Reference
		Not
*/
package lang

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"
import expr "github.com/katydid/katydid/expr/ast"

import fmt "fmt"
import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Grammar struct {
	PatternDecls []*PatternDecl `protobuf:"bytes,1,rep" json:"PatternDecls,omitempty"`
	After        *expr.Space    `protobuf:"bytes,2,opt" json:"After,omitempty"`
}

func (m *Grammar) Reset()      { *m = Grammar{} }
func (*Grammar) ProtoMessage() {}

func (m *Grammar) GetPatternDecls() []*PatternDecl {
	if m != nil {
		return m.PatternDecls
	}
	return nil
}

func (m *Grammar) GetAfter() *expr.Space {
	if m != nil {
		return m.After
	}
	return nil
}

type PatternDecl struct {
	Before  *expr.Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	Name    string        `protobuf:"bytes,2,opt" json:"Name"`
	Eq      *expr.Keyword `protobuf:"bytes,3,opt" json:"Eq,omitempty"`
	Pattern *Pattern      `protobuf:"bytes,4,opt" json:"Pattern,omitempty"`
}

func (m *PatternDecl) Reset()      { *m = PatternDecl{} }
func (*PatternDecl) ProtoMessage() {}

func (m *PatternDecl) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *PatternDecl) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PatternDecl) GetEq() *expr.Keyword {
	if m != nil {
		return m.Eq
	}
	return nil
}

func (m *PatternDecl) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

type NameExpr struct {
	Name          *Name          `protobuf:"bytes,1,opt" json:"Name,omitempty"`
	AnyName       *AnyName       `protobuf:"bytes,2,opt" json:"AnyName,omitempty"`
	AnyNameExcept *AnyNameExcept `protobuf:"bytes,3,opt" json:"AnyNameExcept,omitempty"`
	NameChoice    *NameChoice    `protobuf:"bytes,4,opt" json:"NameChoice,omitempty"`
}

func (m *NameExpr) Reset()      { *m = NameExpr{} }
func (*NameExpr) ProtoMessage() {}

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

type Name struct {
	Before     *expr.Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	OpenParen  *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	BeforeName *expr.Space   `protobuf:"bytes,3,opt" json:"BeforeName,omitempty"`
	Name       string        `protobuf:"bytes,4,opt" json:"Name"`
	CloseParen *expr.Keyword `protobuf:"bytes,5,opt" json:"CloseParen,omitempty"`
}

func (m *Name) Reset()      { *m = Name{} }
func (*Name) ProtoMessage() {}

func (m *Name) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Name) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *Name) GetBeforeName() *expr.Space {
	if m != nil {
		return m.BeforeName
	}
	return nil
}

func (m *Name) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Name) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type AnyName struct {
	Before *expr.Space `protobuf:"bytes,1,opt" json:"Before,omitempty"`
}

func (m *AnyName) Reset()      { *m = AnyName{} }
func (*AnyName) ProtoMessage() {}

func (m *AnyName) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

type AnyNameExcept struct {
	Before     *expr.Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	OpenParen  *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	Except     *NameExpr     `protobuf:"bytes,3,opt" json:"Except,omitempty"`
	CloseParen *expr.Keyword `protobuf:"bytes,4,opt" json:"CloseParen,omitempty"`
}

func (m *AnyNameExcept) Reset()      { *m = AnyNameExcept{} }
func (*AnyNameExcept) ProtoMessage() {}

func (m *AnyNameExcept) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *AnyNameExcept) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *AnyNameExcept) GetExcept() *NameExpr {
	if m != nil {
		return m.Except
	}
	return nil
}

func (m *AnyNameExcept) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type NameChoice struct {
	Before     *expr.Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	OpenParen  *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	Left       *NameExpr     `protobuf:"bytes,3,opt" json:"Left,omitempty"`
	Comma      *expr.Keyword `protobuf:"bytes,4,opt" json:"Comma,omitempty"`
	Right      *NameExpr     `protobuf:"bytes,5,opt" json:"Right,omitempty"`
	CloseParen *expr.Keyword `protobuf:"bytes,6,opt" json:"CloseParen,omitempty"`
}

func (m *NameChoice) Reset()      { *m = NameChoice{} }
func (*NameChoice) ProtoMessage() {}

func (m *NameChoice) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *NameChoice) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *NameChoice) GetLeft() *NameExpr {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *NameChoice) GetComma() *expr.Keyword {
	if m != nil {
		return m.Comma
	}
	return nil
}

func (m *NameChoice) GetRight() *NameExpr {
	if m != nil {
		return m.Right
	}
	return nil
}

func (m *NameChoice) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type Pattern struct {
	Empty      *Empty      `protobuf:"bytes,1,opt" json:"Empty,omitempty"`
	EmptySet   *EmptySet   `protobuf:"bytes,2,opt" json:"EmptySet,omitempty"`
	TreeNode   *TreeNode   `protobuf:"bytes,3,opt" json:"TreeNode,omitempty"`
	LeafNode   *LeafNode   `protobuf:"bytes,4,opt" json:"LeafNode,omitempty"`
	Concat     *Concat     `protobuf:"bytes,5,opt" json:"Concat,omitempty"`
	Or         *Or         `protobuf:"bytes,6,opt" json:"Or,omitempty"`
	And        *And        `protobuf:"bytes,7,opt" json:"And,omitempty"`
	ZeroOrMore *ZeroOrMore `protobuf:"bytes,8,opt" json:"ZeroOrMore,omitempty"`
	Reference  *Reference  `protobuf:"bytes,9,opt" json:"Reference,omitempty"`
	Not        *Not        `protobuf:"bytes,10,opt" json:"Not,omitempty"`
}

func (m *Pattern) Reset()      { *m = Pattern{} }
func (*Pattern) ProtoMessage() {}

func (m *Pattern) GetEmpty() *Empty {
	if m != nil {
		return m.Empty
	}
	return nil
}

func (m *Pattern) GetEmptySet() *EmptySet {
	if m != nil {
		return m.EmptySet
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

type Empty struct {
	Before *expr.Space `protobuf:"bytes,1,opt" json:"Before,omitempty"`
}

func (m *Empty) Reset()      { *m = Empty{} }
func (*Empty) ProtoMessage() {}

func (m *Empty) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

type EmptySet struct {
	Before *expr.Space `protobuf:"bytes,1,opt" json:"Before,omitempty"`
}

func (m *EmptySet) Reset()      { *m = EmptySet{} }
func (*EmptySet) ProtoMessage() {}

func (m *EmptySet) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

type TreeNode struct {
	Before     *expr.Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	OpenParen  *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	Name       *NameExpr     `protobuf:"bytes,4,opt" json:"Name,omitempty"`
	Comma      *expr.Keyword `protobuf:"bytes,5,opt" json:"Comma,omitempty"`
	Pattern    *Pattern      `protobuf:"bytes,6,opt" json:"Pattern,omitempty"`
	CloseParen *expr.Keyword `protobuf:"bytes,7,opt" json:"CloseParen,omitempty"`
}

func (m *TreeNode) Reset()      { *m = TreeNode{} }
func (*TreeNode) ProtoMessage() {}

func (m *TreeNode) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *TreeNode) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *TreeNode) GetName() *NameExpr {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *TreeNode) GetComma() *expr.Keyword {
	if m != nil {
		return m.Comma
	}
	return nil
}

func (m *TreeNode) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *TreeNode) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type LeafNode struct {
	Before     *expr.Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	OpenParen  *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	Expr       *expr.Expr    `protobuf:"bytes,3,opt" json:"Expr,omitempty"`
	CloseParen *expr.Keyword `protobuf:"bytes,4,opt" json:"CloseParen,omitempty"`
}

func (m *LeafNode) Reset()      { *m = LeafNode{} }
func (*LeafNode) ProtoMessage() {}

func (m *LeafNode) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *LeafNode) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *LeafNode) GetExpr() *expr.Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

func (m *LeafNode) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type Concat struct {
	Before       *expr.Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	OpenParen    *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	LeftPattern  *Pattern      `protobuf:"bytes,3,opt" json:"LeftPattern,omitempty"`
	Comma        *expr.Keyword `protobuf:"bytes,4,opt" json:"Comma,omitempty"`
	RightPattern *Pattern      `protobuf:"bytes,5,opt" json:"RightPattern,omitempty"`
	CloseParen   *expr.Keyword `protobuf:"bytes,6,opt" json:"CloseParen,omitempty"`
}

func (m *Concat) Reset()      { *m = Concat{} }
func (*Concat) ProtoMessage() {}

func (m *Concat) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Concat) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *Concat) GetLeftPattern() *Pattern {
	if m != nil {
		return m.LeftPattern
	}
	return nil
}

func (m *Concat) GetComma() *expr.Keyword {
	if m != nil {
		return m.Comma
	}
	return nil
}

func (m *Concat) GetRightPattern() *Pattern {
	if m != nil {
		return m.RightPattern
	}
	return nil
}

func (m *Concat) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type Or struct {
	Before       *expr.Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	OpenParen    *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	LeftPattern  *Pattern      `protobuf:"bytes,3,opt" json:"LeftPattern,omitempty"`
	Comma        *expr.Keyword `protobuf:"bytes,4,opt" json:"Comma,omitempty"`
	RightPattern *Pattern      `protobuf:"bytes,5,opt" json:"RightPattern,omitempty"`
	CloseParen   *expr.Keyword `protobuf:"bytes,6,opt" json:"CloseParen,omitempty"`
}

func (m *Or) Reset()      { *m = Or{} }
func (*Or) ProtoMessage() {}

func (m *Or) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Or) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *Or) GetLeftPattern() *Pattern {
	if m != nil {
		return m.LeftPattern
	}
	return nil
}

func (m *Or) GetComma() *expr.Keyword {
	if m != nil {
		return m.Comma
	}
	return nil
}

func (m *Or) GetRightPattern() *Pattern {
	if m != nil {
		return m.RightPattern
	}
	return nil
}

func (m *Or) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type And struct {
	Before       *expr.Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	OpenParen    *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	LeftPattern  *Pattern      `protobuf:"bytes,3,opt" json:"LeftPattern,omitempty"`
	Comma        *expr.Keyword `protobuf:"bytes,4,opt" json:"Comma,omitempty"`
	RightPattern *Pattern      `protobuf:"bytes,5,opt" json:"RightPattern,omitempty"`
	CloseParen   *expr.Keyword `protobuf:"bytes,6,opt" json:"CloseParen,omitempty"`
}

func (m *And) Reset()      { *m = And{} }
func (*And) ProtoMessage() {}

func (m *And) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *And) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *And) GetLeftPattern() *Pattern {
	if m != nil {
		return m.LeftPattern
	}
	return nil
}

func (m *And) GetComma() *expr.Keyword {
	if m != nil {
		return m.Comma
	}
	return nil
}

func (m *And) GetRightPattern() *Pattern {
	if m != nil {
		return m.RightPattern
	}
	return nil
}

func (m *And) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type ZeroOrMore struct {
	Before     *expr.Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	OpenParen  *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	Pattern    *Pattern      `protobuf:"bytes,3,opt" json:"Pattern,omitempty"`
	CloseParen *expr.Keyword `protobuf:"bytes,4,opt" json:"CloseParen,omitempty"`
}

func (m *ZeroOrMore) Reset()      { *m = ZeroOrMore{} }
func (*ZeroOrMore) ProtoMessage() {}

func (m *ZeroOrMore) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *ZeroOrMore) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *ZeroOrMore) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *ZeroOrMore) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type Reference struct {
	Before     *expr.Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	OpenParen  *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	BeforeName *expr.Space   `protobuf:"bytes,3,opt" json:"BeforeName,omitempty"`
	Name       string        `protobuf:"bytes,4,opt" json:"Name"`
	CloseParen *expr.Keyword `protobuf:"bytes,5,opt" json:"CloseParen,omitempty"`
}

func (m *Reference) Reset()      { *m = Reference{} }
func (*Reference) ProtoMessage() {}

func (m *Reference) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Reference) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *Reference) GetBeforeName() *expr.Space {
	if m != nil {
		return m.BeforeName
	}
	return nil
}

func (m *Reference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Reference) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type Not struct {
	Before     *expr.Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	OpenParen  *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	Pattern    *Pattern      `protobuf:"bytes,3,opt" json:"Pattern,omitempty"`
	CloseParen *expr.Keyword `protobuf:"bytes,4,opt" json:"CloseParen,omitempty"`
}

func (m *Not) Reset()      { *m = Not{} }
func (*Not) ProtoMessage() {}

func (m *Not) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Not) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *Not) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *Not) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

func init() {
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

func (this *NameExpr) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *Name:
		this.Name = vt
	case *AnyName:
		this.AnyName = vt
	case *AnyNameExcept:
		this.AnyNameExcept = vt
	case *NameChoice:
		this.NameChoice = vt
	default:
		return false
	}
	return true
}
func (this *Pattern) GetValue() interface{} {
	if this.Empty != nil {
		return this.Empty
	}
	if this.EmptySet != nil {
		return this.EmptySet
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
	return nil
}

func (this *Pattern) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *Empty:
		this.Empty = vt
	case *EmptySet:
		this.EmptySet = vt
	case *TreeNode:
		this.TreeNode = vt
	case *LeafNode:
		this.LeafNode = vt
	case *Concat:
		this.Concat = vt
	case *Or:
		this.Or = vt
	case *And:
		this.And = vt
	case *ZeroOrMore:
		this.ZeroOrMore = vt
	case *Reference:
		this.Reference = vt
	case *Not:
		this.Not = vt
	default:
		return false
	}
	return true
}
func (this *Grammar) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.Grammar{` +
		`PatternDecls:` + fmt.Sprintf("%#v", this.PatternDecls),
		`After:` + fmt.Sprintf("%#v", this.After) + `}`}, ", ")
	return s
}
func (this *PatternDecl) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.PatternDecl{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`Name:` + fmt.Sprintf("%#v", this.Name),
		`Eq:` + fmt.Sprintf("%#v", this.Eq),
		`Pattern:` + fmt.Sprintf("%#v", this.Pattern) + `}`}, ", ")
	return s
}
func (this *NameExpr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.NameExpr{` +
		`Name:` + fmt.Sprintf("%#v", this.Name),
		`AnyName:` + fmt.Sprintf("%#v", this.AnyName),
		`AnyNameExcept:` + fmt.Sprintf("%#v", this.AnyNameExcept),
		`NameChoice:` + fmt.Sprintf("%#v", this.NameChoice) + `}`}, ", ")
	return s
}
func (this *Name) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.Name{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`BeforeName:` + fmt.Sprintf("%#v", this.BeforeName),
		`Name:` + fmt.Sprintf("%#v", this.Name),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *AnyName) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.AnyName{` +
		`Before:` + fmt.Sprintf("%#v", this.Before) + `}`}, ", ")
	return s
}
func (this *AnyNameExcept) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.AnyNameExcept{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`Except:` + fmt.Sprintf("%#v", this.Except),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *NameChoice) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.NameChoice{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`Left:` + fmt.Sprintf("%#v", this.Left),
		`Comma:` + fmt.Sprintf("%#v", this.Comma),
		`Right:` + fmt.Sprintf("%#v", this.Right),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *Pattern) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.Pattern{` +
		`Empty:` + fmt.Sprintf("%#v", this.Empty),
		`EmptySet:` + fmt.Sprintf("%#v", this.EmptySet),
		`TreeNode:` + fmt.Sprintf("%#v", this.TreeNode),
		`LeafNode:` + fmt.Sprintf("%#v", this.LeafNode),
		`Concat:` + fmt.Sprintf("%#v", this.Concat),
		`Or:` + fmt.Sprintf("%#v", this.Or),
		`And:` + fmt.Sprintf("%#v", this.And),
		`ZeroOrMore:` + fmt.Sprintf("%#v", this.ZeroOrMore),
		`Reference:` + fmt.Sprintf("%#v", this.Reference),
		`Not:` + fmt.Sprintf("%#v", this.Not) + `}`}, ", ")
	return s
}
func (this *Empty) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.Empty{` +
		`Before:` + fmt.Sprintf("%#v", this.Before) + `}`}, ", ")
	return s
}
func (this *EmptySet) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.EmptySet{` +
		`Before:` + fmt.Sprintf("%#v", this.Before) + `}`}, ", ")
	return s
}
func (this *TreeNode) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.TreeNode{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`Name:` + fmt.Sprintf("%#v", this.Name),
		`Comma:` + fmt.Sprintf("%#v", this.Comma),
		`Pattern:` + fmt.Sprintf("%#v", this.Pattern),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *LeafNode) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.LeafNode{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`Expr:` + fmt.Sprintf("%#v", this.Expr),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *Concat) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.Concat{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`LeftPattern:` + fmt.Sprintf("%#v", this.LeftPattern),
		`Comma:` + fmt.Sprintf("%#v", this.Comma),
		`RightPattern:` + fmt.Sprintf("%#v", this.RightPattern),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *Or) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.Or{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`LeftPattern:` + fmt.Sprintf("%#v", this.LeftPattern),
		`Comma:` + fmt.Sprintf("%#v", this.Comma),
		`RightPattern:` + fmt.Sprintf("%#v", this.RightPattern),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *And) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.And{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`LeftPattern:` + fmt.Sprintf("%#v", this.LeftPattern),
		`Comma:` + fmt.Sprintf("%#v", this.Comma),
		`RightPattern:` + fmt.Sprintf("%#v", this.RightPattern),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *ZeroOrMore) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.ZeroOrMore{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`Pattern:` + fmt.Sprintf("%#v", this.Pattern),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *Reference) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.Reference{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`BeforeName:` + fmt.Sprintf("%#v", this.BeforeName),
		`Name:` + fmt.Sprintf("%#v", this.Name),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *Not) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&lang.Not{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`Pattern:` + fmt.Sprintf("%#v", this.Pattern),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func valueToGoStringLang(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringLang(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
