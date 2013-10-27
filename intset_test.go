package intset

import (
	"testing"

	"github.com/knakk/specs"
)

func TestIntSet(t *testing.T) {
	s := specs.New(t)

	set := New()
	s.Expect(set.Contains(1), false)
	set.Add(1)
	s.Expect(set.Contains(1), true)
	s.Expect(set.Size(), 1)
	set.Add(2)
	set.Add(3)
	s.Expect(set.Add(3), false)
	s.Expect(set.Size(), 3)
	set.Remove(2)
	s.Expect(set.Contains(2), false)
	s.Expect(set.Size(), 2)

	set2 := NewFromSlice([]int{1, 99})
	s.Expect(set2.Contains(1), true)
	s.Expect(set2.Contains(99), true)
	s.Expect(set2.Size(), 2)

	set.Merge(set2)
	s.Expect(set.Contains(99), true)
	s.Expect(set.Size(), 3)
}

func TestIntSetOperations(t *testing.T) {
	s := specs.New(t)

	s1 := NewFromSlice([]int{1, 2, 3})
	s2 := NewFromSlice([]int{1, 2, 3, 4})
	s3 := NewFromSlice([]int{3, 1, 2})
	s4 := NewFromSlice([]int{8, 9})
	s5 := s1.Union(s4)
	s6 := NewFromSlice([]int{1, 2, 3, 8, 9})
	s7 := s6.Clone()

	tests := []specs.Spec{
		{s1.Equal(s2), false},
		{s1.Equal(s3), true},
		{s5.Equal(s6), true},
		{s1.SubsetOf(s2), true},
		{s1.SubsetOf(s6), true},
		{s3.SubsetOf(s4), false},
		{s2.SupersetOf(s1), true},
		{s6.SupersetOf(s4), true},
		{s1.Intersect(s2).Equal(s3), true},
		{s4.Intersect(s6).Equal(s4), true},
		{s2.Diff(s1).Equal(NewFromSlice([]int{4})), true},
		{s1.SymDiff(s4).Equal(s6), true},
		{s7.Equal(s6), true},
	}
	s.ExpectAll(tests)

	sl := s7.ToSlice()
	for _, i := range sl {
		s.Expect(s7.Contains(i), true)
	}

}
