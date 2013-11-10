package intset

import (
	"fmt"
	"math/big"
	"strings"
)

const wordSize = 1 << (^uintptr(0)>>32&1 + ^uintptr(0)>>16&1 + ^uintptr(0)>>8&1 + 3)

// BitSet is an integer set backed by a bitset implemented using big.Int.
type BitSet struct {
	data *big.Int
}

// NewBitSet is the constructor for BitSet. Max is ignored in this set
// implementation.
func NewBitSet(max int) *BitSet {
	return new(BitSet).init(max)
}

func (set *BitSet) init(max int) *BitSet {
	set.data = big.NewInt(0)
	return set
}

// Clear the set.
func (set *BitSet) Clear() *BitSet {
	set.data = big.NewInt(0)
	return set
}

// Size returns the number of integers in the set.
// Alogrithm stolen from github.com/kisielk/bigset
func (set *BitSet) Size() int {
	var l int
	zero := big.NewInt(0)
	v := new(big.Int).Set(set.data)
	for l = 0; v.Cmp(zero) != 0; l++ {
		vMinusOne := new(big.Int).Sub(v, big.NewInt(1))
		v.And(v, vMinusOne)
	}
	return l
}

// Add one or more integers to the set.
func (set *BitSet) Add(ints ...int) *BitSet {
	for _, i := range ints {
		set.data.SetBit(set.data, i, 1)
	}
	return set
}

// Remove one or more integers from the set.
func (set *BitSet) Remove(ints ...int) *BitSet {
	for _, i := range ints {
		set.data.SetBit(set.data, i, 0)
	}
	return set
}

// All returns a slice of all the integers in the set in ascending order.
func (set *BitSet) All() []int {
	var all []int
	for i, w := range set.data.Bits() {
		for j, z := 0, w; z != 0; j, z = j+1, z>>1 {
			if z&1 != 0 {
				all = append(all, i*wordSize+j)
			}
		}
	}
	return all
}

// Contains returns true if all ints are in the set, otherwise false.
func (set *BitSet) Contains(ints ...int) bool {
	for _, i := range ints {
		if set.data.Bit(i) != uint(1) {
			return false
		}
	}
	return true
}

// Equal checks if two sets both contains all the same items.
func (set *BitSet) Equal(other *BitSet) bool {
	return set.data.Cmp(other.data) == 0
}

// SubsetOf checks if all items in set are also present in other set.
func (set *BitSet) SubsetOf(other *BitSet) bool {
	for _, i := range set.All() {
		if !other.Contains(i) {
			return false
		}
	}
	return true
}

// SupersetOf checks if a set is a superset of another set.
func (set *BitSet) SupersetOf(other *BitSet) bool {
	return other.SubsetOf(set)
}

// Union returns a new set which is the union of two sets.
func (set *BitSet) Union(other *BitSet) *BitSet {
	result := NewBitSet(0)
	result.data.Or(set.data, other.data)
	return result
}

// Intersection returns a new set with integers common to both sets.
func (set *BitSet) Intersection(other *BitSet) *BitSet {
	result := NewBitSet(0)
	result.data.And(set.data, other.data)
	return result
}

// Difference returns a new set with the integers in set which are not in other.
func (set *BitSet) Difference(other *BitSet) *BitSet {
	result := NewBitSet(0)
	result.data.And(set.data, new(big.Int).Not(other.data))
	return result
}

// SymetricDifference returns a new set with the integers in current and other,
// but not in both.
func (set *BitSet) SymetricDifference(other *BitSet) *BitSet {
	result := NewBitSet(0)
	result.data.Xor(set.data, other.data)
	return result
}

// Clone returns a new set which is a clone of current set.
func (set *BitSet) Clone() *BitSet {
	result := NewBitSet(0)
	result.data.Set(set.data)
	return result
}

// String implements the Stringer interface for BitSet.
func (set *BitSet) String() string {
	items := make([]string, 0, set.Size())

	for _, i := range set.All() {
		items = append(items, fmt.Sprintf("%v", i))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}
