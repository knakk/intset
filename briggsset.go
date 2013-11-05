package intset

import (
	"fmt"
	"strings"
)

// BriggsSet is an integer set implementeation based on Briggs/Torczon paper
// "An Efficient Representation for Sparse Sets" from 1993
type BriggsSet struct {
	dense  []int
	sparse []int
	size   int
}

// NewBriggsSet is the constructor for BriggsSet.
func NewBriggsSet(max int) *BriggsSet {
	return new(BriggsSet).init(max)
}

func (set *BriggsSet) init(max int) *BriggsSet {
	set.sparse = make([]int, max+1)
	set.dense = make([]int, max+1)
	return set
}

// Clear the set.
func (set *BriggsSet) Clear() *BriggsSet {
	set.size = 0
	return set
}

// Size returns the number of integers in the set.
func (set *BriggsSet) Size() int {
	return set.size
}

// Add one or more integers to the set.
func (set *BriggsSet) Add(ints ...int) *BriggsSet {
	for _, i := range ints {
		if !set.Contains(i) {
			set.dense[set.size] = i
			set.sparse[i] = set.size
			set.size++
		}
	}
	return set
}

// Remove one or more integers from the set.
func (set *BriggsSet) Remove(ints ...int) *BriggsSet {
	for _, i := range ints {
		if set.Contains(i) {
			j := set.dense[set.size-1]
			set.dense[set.sparse[i]] = j
			set.sparse[j] = set.sparse[i]
			set.size--
		}
	}
	return set
}

// All returns a slice of all the integers in the set. It makes no guarantee
// that the integers are in the same order as they where inserted.
func (set *BriggsSet) All() []int {
	var all []int
	for i := 0; i < set.size; i++ {
		all = append(all, set.dense[i])
	}
	return all
}

// Contains returns true if all ints are in the set, otherwise false.
func (set *BriggsSet) Contains(ints ...int) bool {
	if set.size <= 0 {
		return false
	}
	for _, i := range ints {
		if set.dense[set.sparse[i]] != i {
			return false
		}
	}
	return true
}

// Equal checks if two sets both contains all the same items.
func (set *BriggsSet) Equal(other *BriggsSet) bool {
	if set.Size() != other.Size() {
		return false
	}
	for i := 0; i < set.size; i++ {
		if !other.Contains(set.dense[i]) {
			return false
		}
	}
	return true
}

// SubsetOf checks if all items in set are also present in other set.
func (set *BriggsSet) SubsetOf(other *BriggsSet) bool {
	for i := 0; i < set.size; i++ {
		if !other.Contains(set.dense[i]) {
			return false
		}
	}
	return true
}

// SupersetOf checks if a set is a superset of another set.
func (set *BriggsSet) SupersetOf(other *BriggsSet) bool {
	return other.SubsetOf(set)
}

// Union returns a new set which is the union of two sets.
func (set *BriggsSet) Union(other *BriggsSet) *BriggsSet {
	result := NewBriggsSet(max(len(set.dense), len(other.dense)))
	for i := 0; i < set.size; i++ {
		result.Add(set.dense[i])
	}
	for i := 0; i < other.size; i++ {
		result.Add(other.dense[i])
	}
	return result
}

// Intersection returns a new set with integers common to both sets.
func (set *BriggsSet) Intersection(other *BriggsSet) *BriggsSet {
	result := &BriggsSet{}
	// always loop over the smallest set
	if set.Size() < other.Size() {
		result.init(max(len(set.dense), len(other.dense)))
		for i := 0; i < set.size; i++ {
			if other.Contains(set.dense[i]) {
				result.Add(set.dense[i])
			}
		}
	} else {
		result.init(max(len(set.dense), len(other.dense)))
		for i := 0; i < other.size; i++ {
			if set.Contains(other.dense[i]) {
				result.Add(other.dense[i])
			}
		}
	}
	return result
}

// Difference returns a new set with the integers in set which are not in other.
func (set *BriggsSet) Difference(other *BriggsSet) *BriggsSet {
	result := NewBriggsSet(max(len(set.dense), len(other.dense)))
	for i := 0; i < set.size; i++ {
		if !other.Contains(set.dense[i]) {
			result.Add(set.dense[i])
		}
	}
	return result
}

// SymetricDifference returns a new set with the integers in current and other,
// but not in both.
func (set *BriggsSet) SymetricDifference(other *BriggsSet) *BriggsSet {
	a := set.Difference(other)
	b := other.Difference(set)
	return a.Union(b)
}

// Clone returns a new set which is a clone of current set.
func (set *BriggsSet) Clone() *BriggsSet {
	result := NewBriggsSet(len(set.dense))
	for i := 0; i < set.size; i++ {
		result.Add(set.dense[i]) // TODO use copy(a, b)
	}
	return result
}

// String implements the Stringer interface for BriggsSet.
func (set *BriggsSet) String() string {
	items := make([]string, 0, len(set.dense))

	for _, i := range set.dense {
		items = append(items, fmt.Sprintf("%v", i))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}
