// Code generated by funcs-gen. DO NOT EDIT.
package funcs

import (
	"strings"
)

type elemDoubles struct {
	List        Doubles
	Index       Int
	hash        uint64
	hasVariable bool
}

func (this *elemDoubles) Eval() (float64, error) {
	list, err := this.List.Eval()
	if err != nil {
		return 0, err
	}
	index64, err := this.Index.Eval()
	if err != nil {
		return 0, err
	}
	index := int(index64)
	if len(list) == 0 {
		return 0, NewRangeCheckErr(index, len(list))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return 0, NewRangeCheckErr(index, len(list))
	}
	return list[index], nil
}

func (this *elemDoubles) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*elemDoubles); ok {
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		if c := this.Index.Compare(other.Index); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *elemDoubles) HasVariable() bool {
	return this.hasVariable
}

func (this *elemDoubles) String() string {
	return "elem(" + sjoin(this.List, this.Index) + ")"
}

func (this *elemDoubles) Hash() uint64 {
	return this.hash
}

func init() {
	Register("elem", ElemDoubles)
}

// ElemDoubles returns a function that returns the n'th element of the list.
func ElemDoubles(list Doubles, n Int) Double {
	return TrimDouble(&elemDoubles{
		List:        list,
		Index:       n,
		hash:        hashWithId(63639164578, n, list),
		hasVariable: n.HasVariable() || list.HasVariable(),
	})
}

type elemInts struct {
	List        Ints
	Index       Int
	hash        uint64
	hasVariable bool
}

func (this *elemInts) Eval() (int64, error) {
	list, err := this.List.Eval()
	if err != nil {
		return 0, err
	}
	index64, err := this.Index.Eval()
	if err != nil {
		return 0, err
	}
	index := int(index64)
	if len(list) == 0 {
		return 0, NewRangeCheckErr(index, len(list))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return 0, NewRangeCheckErr(index, len(list))
	}
	return list[index], nil
}

func (this *elemInts) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*elemInts); ok {
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		if c := this.Index.Compare(other.Index); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *elemInts) HasVariable() bool {
	return this.hasVariable
}

func (this *elemInts) String() string {
	return "elem(" + sjoin(this.List, this.Index) + ")"
}

func (this *elemInts) Hash() uint64 {
	return this.hash
}

func init() {
	Register("elem", ElemInts)
}

// ElemInts returns a function that returns the n'th element of the list.
func ElemInts(list Ints, n Int) Int {
	return TrimInt(&elemInts{
		List:        list,
		Index:       n,
		hash:        hashWithId(2284164, n, list),
		hasVariable: n.HasVariable() || list.HasVariable(),
	})
}

type elemUints struct {
	List        Uints
	Index       Int
	hash        uint64
	hasVariable bool
}

func (this *elemUints) Eval() (uint64, error) {
	list, err := this.List.Eval()
	if err != nil {
		return 0, err
	}
	index64, err := this.Index.Eval()
	if err != nil {
		return 0, err
	}
	index := int(index64)
	if len(list) == 0 {
		return 0, NewRangeCheckErr(index, len(list))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return 0, NewRangeCheckErr(index, len(list))
	}
	return list[index], nil
}

func (this *elemUints) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*elemUints); ok {
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		if c := this.Index.Compare(other.Index); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *elemUints) HasVariable() bool {
	return this.hasVariable
}

func (this *elemUints) String() string {
	return "elem(" + sjoin(this.List, this.Index) + ")"
}

func (this *elemUints) Hash() uint64 {
	return this.hash
}

func init() {
	Register("elem", ElemUints)
}

// ElemUints returns a function that returns the n'th element of the list.
func ElemUints(list Uints, n Int) Uint {
	return TrimUint(&elemUints{
		List:        list,
		Index:       n,
		hash:        hashWithId(81736761, n, list),
		hasVariable: n.HasVariable() || list.HasVariable(),
	})
}

type elemBools struct {
	List        Bools
	Index       Int
	hash        uint64
	hasVariable bool
}

func (this *elemBools) Eval() (bool, error) {
	list, err := this.List.Eval()
	if err != nil {
		return false, err
	}
	index64, err := this.Index.Eval()
	if err != nil {
		return false, err
	}
	index := int(index64)
	if len(list) == 0 {
		return false, NewRangeCheckErr(index, len(list))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return false, NewRangeCheckErr(index, len(list))
	}
	return list[index], nil
}

func (this *elemBools) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*elemBools); ok {
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		if c := this.Index.Compare(other.Index); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *elemBools) HasVariable() bool {
	return this.hasVariable
}

func (this *elemBools) String() string {
	return "elem(" + sjoin(this.List, this.Index) + ")"
}

func (this *elemBools) Hash() uint64 {
	return this.hash
}

func init() {
	Register("elem", ElemBools)
}

// ElemBools returns a function that returns the n'th element of the list.
func ElemBools(list Bools, n Int) Bool {
	return TrimBool(&elemBools{
		List:        list,
		Index:       n,
		hash:        hashWithId(64369321, n, list),
		hasVariable: n.HasVariable() || list.HasVariable(),
	})
}

type elemStrings struct {
	List        Strings
	Index       Int
	hash        uint64
	hasVariable bool
}

func (this *elemStrings) Eval() (string, error) {
	list, err := this.List.Eval()
	if err != nil {
		return "", err
	}
	index64, err := this.Index.Eval()
	if err != nil {
		return "", err
	}
	index := int(index64)
	if len(list) == 0 {
		return "", NewRangeCheckErr(index, len(list))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return "", NewRangeCheckErr(index, len(list))
	}
	return list[index], nil
}

func (this *elemStrings) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*elemStrings); ok {
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		if c := this.Index.Compare(other.Index); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *elemStrings) HasVariable() bool {
	return this.hasVariable
}

func (this *elemStrings) String() string {
	return "elem(" + sjoin(this.List, this.Index) + ")"
}

func (this *elemStrings) Hash() uint64 {
	return this.hash
}

func init() {
	Register("elem", ElemStrings)
}

// ElemStrings returns a function that returns the n'th element of the list.
func ElemStrings(list Strings, n Int) String {
	return TrimString(&elemStrings{
		List:        list,
		Index:       n,
		hash:        hashWithId(77092305506, n, list),
		hasVariable: n.HasVariable() || list.HasVariable(),
	})
}

type elemListOfBytes struct {
	List        ListOfBytes
	Index       Int
	hash        uint64
	hasVariable bool
}

func (this *elemListOfBytes) Eval() ([]byte, error) {
	list, err := this.List.Eval()
	if err != nil {
		return nil, err
	}
	index64, err := this.Index.Eval()
	if err != nil {
		return nil, err
	}
	index := int(index64)
	if len(list) == 0 {
		return nil, NewRangeCheckErr(index, len(list))
	}
	if index < 0 {
		index = index % len(list)
	}
	if len(list) <= index {
		return nil, NewRangeCheckErr(index, len(list))
	}
	return list[index], nil
}

func (this *elemListOfBytes) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*elemListOfBytes); ok {
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		if c := this.Index.Compare(other.Index); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *elemListOfBytes) HasVariable() bool {
	return this.hasVariable
}

func (this *elemListOfBytes) String() string {
	return "elem(" + sjoin(this.List, this.Index) + ")"
}

func (this *elemListOfBytes) Hash() uint64 {
	return this.hash
}

func init() {
	Register("elem", ElemListOfBytes)
}

// ElemListOfBytes returns a function that returns the n'th element of the list.
func ElemListOfBytes(list ListOfBytes, n Int) Bytes {
	return TrimBytes(&elemListOfBytes{
		List:        list,
		Index:       n,
		hash:        hashWithId(65169257167589942, n, list),
		hasVariable: n.HasVariable() || list.HasVariable(),
	})
}
