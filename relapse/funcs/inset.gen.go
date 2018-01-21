// Code generated by funcs-gen. DO NOT EDIT.
package funcs

import (
	"strings"
)

type inSetInt struct {
	Elem        Int
	List        ConstInts
	set         map[int64]struct{}
	hash        uint64
	hasVariable bool
}

func (this *inSetInt) Eval() (bool, error) {
	v, err := this.Elem.Eval()
	if err != nil {
		return false, err
	}
	_, ok := this.set[v]
	return ok, nil
}

func (this *inSetInt) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*inSetInt); ok {
		if c := this.Elem.Compare(other.Elem); c != 0 {
			return c
		}
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *inSetInt) String() string {
	return "contains(" + sjoin(this.Elem, this.List) + ")"
}

func (this *inSetInt) HasVariable() bool {
	return this.hasVariable
}

func (this *inSetInt) Hash() uint64 {
	return this.hash
}

func init() {
	Register("contains", ContainsInt)
}

//ContainsInt returns a function that checks whether the element is contained in the list.
func ContainsInt(element Int, list ConstInts) (Bool, error) {
	if list.HasVariable() {
		return nil, ErrContainsListNotConst{}
	}
	l, err := list.Eval()
	if err != nil {
		return nil, err
	}
	set := make(map[int64]struct{})
	for i := range l {
		set[l[i]] = struct{}{}
	}
	return TrimBool(&inSetInt{
		Elem:        element,
		List:        list,
		set:         set,
		hash:        hashWithId(73679, element, list),
		hasVariable: element.HasVariable() || list.HasVariable(),
	}), nil
}

type inSetUint struct {
	Elem        Uint
	List        ConstUints
	set         map[uint64]struct{}
	hash        uint64
	hasVariable bool
}

func (this *inSetUint) Eval() (bool, error) {
	v, err := this.Elem.Eval()
	if err != nil {
		return false, err
	}
	_, ok := this.set[v]
	return ok, nil
}

func (this *inSetUint) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*inSetUint); ok {
		if c := this.Elem.Compare(other.Elem); c != 0 {
			return c
		}
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *inSetUint) String() string {
	return "contains(" + sjoin(this.Elem, this.List) + ")"
}

func (this *inSetUint) HasVariable() bool {
	return this.hasVariable
}

func (this *inSetUint) Hash() uint64 {
	return this.hash
}

func init() {
	Register("contains", ContainsUint)
}

//ContainsUint returns a function that checks whether the element is contained in the list.
func ContainsUint(element Uint, list ConstUints) (Bool, error) {
	if list.HasVariable() {
		return nil, ErrContainsListNotConst{}
	}
	l, err := list.Eval()
	if err != nil {
		return nil, err
	}
	set := make(map[uint64]struct{})
	for i := range l {
		set[l[i]] = struct{}{}
	}
	return TrimBool(&inSetUint{
		Elem:        element,
		List:        list,
		set:         set,
		hash:        hashWithId(2636666, element, list),
		hasVariable: element.HasVariable() || list.HasVariable(),
	}), nil
}

type inSetString struct {
	Elem        String
	List        ConstStrings
	set         map[string]struct{}
	hash        uint64
	hasVariable bool
}

func (this *inSetString) Eval() (bool, error) {
	v, err := this.Elem.Eval()
	if err != nil {
		return false, err
	}
	_, ok := this.set[v]
	return ok, nil
}

func (this *inSetString) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*inSetString); ok {
		if c := this.Elem.Compare(other.Elem); c != 0 {
			return c
		}
		if c := this.List.Compare(other.List); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *inSetString) String() string {
	return "contains(" + sjoin(this.Elem, this.List) + ")"
}

func (this *inSetString) HasVariable() bool {
	return this.hasVariable
}

func (this *inSetString) Hash() uint64 {
	return this.hash
}

func init() {
	Register("contains", ContainsString)
}

//ContainsString returns a function that checks whether the element is contained in the list.
func ContainsString(element String, list ConstStrings) (Bool, error) {
	if list.HasVariable() {
		return nil, ErrContainsListNotConst{}
	}
	l, err := list.Eval()
	if err != nil {
		return nil, err
	}
	set := make(map[string]struct{})
	for i := range l {
		set[l[i]] = struct{}{}
	}
	return TrimBool(&inSetString{
		Elem:        element,
		List:        list,
		set:         set,
		hash:        hashWithId(2486848561, element, list),
		hasVariable: element.HasVariable() || list.HasVariable(),
	}), nil
}
