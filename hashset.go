package intset

import (
	"fmt"
	"strings"
)

// HashSet is an integer set backed by a map.
type HashSet struct {
	data map[int]bool
}

// NewHashSet is the constructor for HashSet.
func NewHashSet(max int) *HashSet {
	return new(HashSet).init(max)
}

func (set *HashSet) init(max int) *HashSet {
	set.data = make(map[int]bool, max+1)
	return set
}

// Clear the set.
func (set *HashSet) Clear() *HashSet {
	set.data = make(map[int]bool, len(set.data))
	return set
}

// Size returns the number of integers in the set.
func (set *HashSet) Size() int {
	return len(set.data)
}

// Add one or more integers to the set.
func (set *HashSet) Add(ints ...int) *HashSet {
	for _, i := range ints {
		set.data[i] = true
	}
	return set
}

// Remove one or more integers from the set.
func (set *HashSet) Remove(ints ...int) *HashSet {
	for _, i := range ints {
		delete(set.data, i)
	}
	return set
}

// All returns a slice of all the integers in the set. It makes no guarantee
// that the integers are in the same order as they where inserted.
func (set *HashSet) All() []int {
	var all []int
	for i := range set.data {
		all = append(all, i)
	}
	return all
}

// Contains returns true if all ints are in the set, otherwise false.
func (set *HashSet) Contains(ints ...int) bool {
	for _, i := range ints {
		if _, found := set.data[i]; !found {
			return false
		}
	}
	return true
}

// Equal checks if two sets both contains all the same items.
func (set *HashSet) Equal(other *HashSet) bool {
	if set.Size() != other.Size() {
		return false
	}
	for k := range set.data {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

// SubsetOf checks if all items in set are also present in other set.
func (set *HashSet) SubsetOf(other *HashSet) bool {
	for i := range set.data {
		if !other.Contains(i) {
			return false
		}
	}
	return true
}

// SupersetOf checks if a set is a superset of another set.
func (set *HashSet) SupersetOf(other *HashSet) bool {
	return other.SubsetOf(set)
}

// Union returns a new set which is the union of two sets.
func (set *HashSet) Union(other *HashSet) *HashSet {
	result := NewHashSet(set.Size() + other.Size())
	for i := range set.data {
		result.Add(i)
	}
	for i := range other.data {
		result.Add(i)
	}
	return result
}

// Intersection returns a new set with integers common to both sets.
func (set *HashSet) Intersection(other *HashSet) *HashSet {
	result := &HashSet{}
	// always loop over the smallest set
	if set.Size() < other.Size() {
		result.init(set.Size())
		for i := range set.data {
			if other.Contains(i) {
				result.Add(i)
			}
		}
	} else {
		result.init(other.Size())
		for i := range other.data {
			if set.Contains(i) {
				result.Add(i)
			}
		}
	}
	return result
}

// Difference returns a new set with the integers in set which are not in other.
func (set *HashSet) Difference(other *HashSet) *HashSet {
	result := NewHashSet(set.Size())
	for i := range set.data {
		if !other.Contains(i) {
			result.Add(i)
		}
	}
	return result
}

// SymetricDifference returns a new set with the integers in current and other,
// but not in both.
func (set *HashSet) SymetricDifference(other *HashSet) *HashSet {
	a := set.Difference(other)
	b := other.Difference(set)
	return a.Union(b)
}

// Clone returns a new set which is a clone of current set.
func (set *HashSet) Clone() *HashSet {
	result := NewHashSet(set.Size())
	for i := range set.data {
		result.Add(i)
	}
	return result
}

// String implements the Stringer interface for HashSet.
func (set *HashSet) String() string {
	items := make([]string, 0, len(set.data))

	for k := range set.data {
		items = append(items, fmt.Sprintf("%v", k))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}
