package intset

import (
	"fmt"
	"strings"
)

// SliceSet is an integer set backed by a slice.
type SliceSet struct {
	data  []bool
	count int
}

// NewSliceSet is the constructor for SliceSet.
func NewSliceSet(max int) *SliceSet {
	return new(SliceSet).init(max)
}

func (set *SliceSet) init(max int) *SliceSet {
	set.data = make([]bool, max+1)
	return set
}

// Clear the set.
func (set *SliceSet) Clear() *SliceSet {
	set.data = make([]bool, len(set.data))
	set.count = 0
	return set
}

// Size returns the number of integers in the set.
func (set *SliceSet) Size() int {
	return set.count
}

// Add one or more integers to the set.
func (set *SliceSet) Add(ints ...int) *SliceSet {
	for _, i := range ints {
		if !set.data[i] {
			set.count++
			set.data[i] = true
		}
	}
	return set
}

// Remove one or more integers from the set.
func (set *SliceSet) Remove(ints ...int) *SliceSet {
	for _, i := range ints {
		if set.data[i] == true {
			set.count--
		}
		set.data[i] = false

	}
	return set
}

// All returns a slice of all the integers in the set, in ascending order.
func (set *SliceSet) All() []int {
	var all []int
	for i, b := range set.data {
		if b {
			all = append(all, i)
		}
	}
	return all
}

// Contains returns true if all ints are in the set, otherwise false.
func (set *SliceSet) Contains(ints ...int) bool {
	for _, i := range ints {
		if !set.data[i] {
			return false
		}
	}
	return true
}

// Equal checks if two sets both contains all the same items.
func (set *SliceSet) Equal(other *SliceSet) bool {
	if set.Size() != other.Size() {
		return false
	}
	for i, b := range set.data {
		if b && !other.Contains(i) {
			return false
		}
	}
	return true
}

// SubsetOf checks if all items in set are also present in other set.
func (set *SliceSet) SubsetOf(other *SliceSet) bool {
	for i, b := range set.data {
		if b && !other.Contains(i) {
			return false
		}
	}
	return true
}

// SupersetOf checks if a set is a superset of another set.
func (set *SliceSet) SupersetOf(other *SliceSet) bool {
	return other.SubsetOf(set)
}

// Union returns a new set which is the union of two sets.
func (set *SliceSet) Union(other *SliceSet) *SliceSet {
	result := NewSliceSet(max(len(set.data), len(other.data)))
	for i, b := range set.data {
		if b {
			result.Add(i)
		}
	}
	for i, b := range other.data {
		if b {
			result.Add(i)
		}
	}
	return result
}

// Intersection returns a new set with integers common to both sets.
func (set *SliceSet) Intersection(other *SliceSet) *SliceSet {
	result := &SliceSet{}
	// always loop over the smallest set
	if len(set.data) < len(other.data) {
		result.init(max(len(set.data), len(other.data)))
		for i, b := range set.data {
			if b && other.Contains(i) {
				result.Add(i)
			}
		}
	} else {
		result.init(max(len(set.data), len(other.data)))
		for i, b := range other.data {
			if b && set.Contains(i) {
				result.Add(i)
			}
		}
	}
	return result
}

// Difference returns a new set with the integers in set which are not in other.
func (set *SliceSet) Difference(other *SliceSet) *SliceSet {
	result := NewSliceSet(len(set.data))
	for i, b := range set.data {
		if b && !other.Contains(i) {
			result.Add(i)
		}
	}
	return result
}

// SymetricDifference returns a new set with the integers in current and other,
// but not in both.
func (set *SliceSet) SymetricDifference(other *SliceSet) *SliceSet {
	a := set.Difference(other)
	b := other.Difference(set)
	return a.Union(b)
}

// Clone returns a new set which is a clone of current set.
func (set *SliceSet) Clone() *SliceSet {
	result := NewSliceSet(len(set.data))
	for i, b := range set.data {
		if b {
			result.Add(i)
		}
	}
	return result
}

// String implements the Stringer interface for SliceSet.
func (set *SliceSet) String() string {
	items := make([]string, 0, len(set.data))

	for i, b := range set.data {
		if b {
			items = append(items, fmt.Sprintf("%v", i))
		}
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
