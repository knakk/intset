// Package intset implements an unsorted integer set type, with some common set
// operations.
package intset

// IntSet represents an unordered collection of unique integers backed by a map.
type IntSet map[int]bool

// New returns a new IntSet.
func New() IntSet {
	return make(IntSet)
}

// NewFromSlice creates an IntSet from a slice of integers.
func NewFromSlice(s []int) IntSet {
	set := New()
	for _, i := range s {
		set.Add(i)
	}
	return set
}

// Add adds an integer to the set. It returns false if allready in the set,
// otherwise true.
func (set IntSet) Add(i int) bool {
	_, found := set[i]
	set[i] = true
	return !found
}

// Contains checks if a given int is part of the set.
func (set IntSet) Contains(i int) bool {
	_, found := set[i]
	return found
}

// Remove deletes an integer from the set.
func (set IntSet) Remove(i int) {
	delete(set, i)
}

// Size returns the size of the set.
func (set IntSet) Size() int {
	return len(set)
}

// Merge updates set to include values in other. If you do not want to modify
// an existing set, use the Union function.
func (set IntSet) Merge(other IntSet) {
	for i := range other {
		set.Add(i)
	}
}

// Union returns a new IntSet which is the union of two sets.
func (set IntSet) Union(other IntSet) IntSet {
	s := New()
	for i := range set {
		s.Add(i)
	}
	for i := range other {
		s.Add(i)
	}
	return s
}

// Equal checks if two sets are equal; if both contains all the same items.
func (set IntSet) Equal(other IntSet) bool {
	if set.Size() != other.Size() {
		return false
	}
	for k := range set {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

// SubsetOf checks if all items in set are also present in other set.
func (set IntSet) SubsetOf(other IntSet) bool {
	for i := range set {
		if !other.Contains(i) {
			return false
		}
	}
	return true
}

// SupersetOf checks if a set is a superset of another set, i.e the does a
// reverse SubsetOf.
func (set IntSet) SupersetOf(other IntSet) bool {
	return other.SubsetOf(set)
}

// Intersect returns a new set with integers common to both sets.
func (set IntSet) Intersect(other IntSet) IntSet {
	s := New()
	// always loop over the smallest set
	if set.Size() < other.Size() {
		for i := range set {
			if other.Contains(i) {
				s.Add(i)
			}
		}
	} else {
		for i := range other {
			if set.Contains(i) {
				s.Add(i)
			}
		}
	}
	return s
}

// Diff returns a new set with the integers in set which are not in other.
func (set IntSet) Diff(other IntSet) IntSet {
	s := New()
	for i := range set {
		if !other.Contains(i) {
			s.Add(i)
		}
	}
	return s
}

// SymDiff returns a new set with the integers in current and other, but not
// in both.
func (set IntSet) SymDiff(other IntSet) IntSet {
	a := set.Diff(other)
	b := other.Diff(set)
	return a.Union(b)
}
