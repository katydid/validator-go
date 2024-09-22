// Code generated by goderive DO NOT EDIT.

package funcs

import (
	"math"
)

// deriveHashDouble returns the hash of the object.
func deriveHashDouble(object float64) uint64 {
	return math.Float64bits(object)
}

// deriveHashInt returns the hash of the object.
func deriveHashInt(object int64) uint64 {
	return uint64(object)
}

// deriveHashUint returns the hash of the object.
func deriveHashUint(object uint64) uint64 {
	return object
}

// deriveHashBool returns the hash of the object.
func deriveHashBool(object bool) uint64 {
	if object {
		return 1
	}
	return 0
}

// deriveHashString returns the hash of the object.
func deriveHashString(object string) uint64 {
	var h uint64
	for _, c := range object {
		h = 31*h + uint64(c)
	}
	return h
}

// deriveHashBytes returns the hash of the object.
func deriveHashBytes(object []byte) uint64 {
	if object == nil {
		return 0
	}
	h := uint64(17)
	for i := 0; i < len(object); i++ {
		h = 31*h + uint64(object[i])
	}
	return h
}

// deriveHashDoubles returns the hash of the object.
func deriveHashDoubles(object []float64) uint64 {
	if object == nil {
		return 0
	}
	h := uint64(17)
	for i := 0; i < len(object); i++ {
		h = 31*h + math.Float64bits(object[i])
	}
	return h
}

// deriveHashInts returns the hash of the object.
func deriveHashInts(object []int64) uint64 {
	if object == nil {
		return 0
	}
	h := uint64(17)
	for i := 0; i < len(object); i++ {
		h = 31*h + uint64(object[i])
	}
	return h
}

// deriveHashUints returns the hash of the object.
func deriveHashUints(object []uint64) uint64 {
	if object == nil {
		return 0
	}
	h := uint64(17)
	for i := 0; i < len(object); i++ {
		h = 31*h + object[i]
	}
	return h
}

// deriveHashBools returns the hash of the object.
func deriveHashBools(object []bool) uint64 {
	if object == nil {
		return 0
	}
	h := uint64(17)
	for i := 0; i < len(object); i++ {
		h = 31*h + deriveHashBool(object[i])
	}
	return h
}

// deriveHashStrings returns the hash of the object.
func deriveHashStrings(object []string) uint64 {
	if object == nil {
		return 0
	}
	h := uint64(17)
	for i := 0; i < len(object); i++ {
		h = 31*h + deriveHashString(object[i])
	}
	return h
}

// deriveHashListOfBytes returns the hash of the object.
func deriveHashListOfBytes(object [][]byte) uint64 {
	if object == nil {
		return 0
	}
	h := uint64(17)
	for i := 0; i < len(object); i++ {
		h = 31*h + deriveHashBytes(object[i])
	}
	return h
}
