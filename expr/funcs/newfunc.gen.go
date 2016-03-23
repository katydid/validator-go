// Code generated by funcs-gen.
// DO NOT EDIT!

package funcs

func NewDoubleFunc(uniq string, values ...interface{}) (Double, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Double), nil
}

func NewIntFunc(uniq string, values ...interface{}) (Int, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Int), nil
}

func NewUintFunc(uniq string, values ...interface{}) (Uint, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Uint), nil
}

func NewBoolFunc(uniq string, values ...interface{}) (Bool, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Bool), nil
}

func NewStringFunc(uniq string, values ...interface{}) (String, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(String), nil
}

func NewBytesFunc(uniq string, values ...interface{}) (Bytes, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Bytes), nil
}

func NewDoublesFunc(uniq string, values ...interface{}) (Doubles, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Doubles), nil
}

func NewIntsFunc(uniq string, values ...interface{}) (Ints, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Ints), nil
}

func NewUintsFunc(uniq string, values ...interface{}) (Uints, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Uints), nil
}

func NewBoolsFunc(uniq string, values ...interface{}) (Bools, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Bools), nil
}

func NewStringsFunc(uniq string, values ...interface{}) (Strings, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Strings), nil
}

func NewListOfBytesFunc(uniq string, values ...interface{}) (ListOfBytes, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(ListOfBytes), nil
}