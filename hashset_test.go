package intset

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/knakk/specs"
)

func TestHashSetAdd(t *testing.T) {
	specs := specs.New(t)

	set := NewHashSet(10).Add(1, 2, 5, 2)

	specs.Expect(set.Contains(1, 2, 5), true)
	specs.Expect(set.Size(), 3)
}

func TestHashSetRemove(t *testing.T) {
	specs := specs.New(t)

	set := NewHashSet(10).Add(3, 1)
	specs.Expect(set.Contains(3, 1), true)
	set.Remove(1, 3)
	specs.Expect(set.Contains(1, 3), false)
}

func TestHashSetContains(t *testing.T) {
	specs := specs.New(t)

	set := NewHashSet(10).Add(1, 2, 3).Remove(2)

	specs.Expect(set.Contains(2), false)
	specs.Expect(set.Contains(1, 3), true)
	specs.Expect(set.Contains(1, 2, 3), false)
}

func TestHashSetClear(t *testing.T) {
	specs := specs.New(t)

	set := NewHashSet(10).Add(1, 2, 3, 4, 5).Clear()

	specs.Expect(set.Contains(1, 2, 3, 4, 5), false)
	specs.Expect(set.Size(), 0)
}

func TestHashSetSize(t *testing.T) {
	specs := specs.New(t)

	set := NewHashSet(10).Add(1, 2, 3)
	specs.Expect(set.Size(), 3)
	set.Remove(2)
	specs.Expect(set.Size(), 2)
}

func TestHashSetAll(t *testing.T) {
	specs := specs.New(t)

	set := NewHashSet(10).Add(99, 1, 5)
	all := set.All()

	specs.Expect(len(all), 3)
	sort.Ints(all) // Hashet.All() doesn't guarantee order of ints
	specs.Expect(all[0], 1)
	specs.Expect(all[1], 5)
	specs.Expect(all[2], 99)
}

func TestHashSetEqual(t *testing.T) {
	specs := specs.New(t)

	setA := NewHashSet(10).Add(1, 2)
	setB := NewHashSet(10).Add(2, 1, 1)
	setC := NewHashSet(10).Add(1, 3)

	specs.Expect(setA.Equal(setB), true)
	specs.Expect(setA.Equal(setC), false)
}

func TestHashSetSubsetOf(t *testing.T) {
	specs := specs.New(t)

	setA := NewHashSet(10).Add(1, 2)
	setB := NewHashSet(10).Add(1, 2, 3)
	setC := NewHashSet(10).Add(3, 4, 5)

	specs.Expect(setA.SubsetOf(setB), true)
	specs.Expect(setA.SubsetOf(setC), false)
}

func TestHashSetSupersetOf(t *testing.T) {
	specs := specs.New(t)

	setA := NewHashSet(10).Add(1, 2)
	setB := NewHashSet(10).Add(1, 2, 3)
	setC := NewHashSet(10).Add(3, 4, 5)

	specs.Expect(setB.SupersetOf(setA), true)
	specs.Expect(setC.SupersetOf(setA), false)
}

func TestHashSetUnion(t *testing.T) {
	specs := specs.New(t)

	setA := NewHashSet(10).Add(1, 2)
	setB := NewHashSet(10).Add(3, 4)
	setC := NewHashSet(10).Add(1, 99)

	specs.Expect(setA.Union(setB).Equal(NewHashSet(10).Add(1, 2, 3, 4)), true)
	specs.Expect(setA.Union(setC).Equal(NewHashSet(10).Add(1, 2, 99)), true)
}

func TestHashSetIntersection(t *testing.T) {
	specs := specs.New(t)

	setA := NewHashSet(10).Add(1, 2)
	setB := NewHashSet(10).Add(1, 2, 3)
	setC := NewHashSet(10).Add(3, 4, 5)

	specs.Expect(setA.Intersection(setB).Equal(setA), true)
	specs.Expect(setB.Intersection(setC).Equal(NewHashSet(10).Add(3)), true)
}

func TestHashSetSymetricDifference(t *testing.T) {
	specs := specs.New(t)

	setA := NewHashSet(10).Add(1, 2, 4)
	setB := NewHashSet(10).Add(1, 2, 3)
	setC := NewHashSet(10).Add(3, 4, 5)

	specs.Expect(setA.SymetricDifference(setB).Equal(NewHashSet(10).Add(3, 4)), true)
	specs.Expect(setB.SymetricDifference(setC).Equal(NewHashSet(10).Add(1, 2, 4, 5)), true)
}

func TestHashSetClone(t *testing.T) {
	specs := specs.New(t)

	setA := NewHashSet(10).Add(9, 3, 1)
	setB := setA.Clone()

	specs.Expect(setB.Equal(setA), true)
}

// Benchmarks

func BenchmarkHashSetAdd(b *testing.B) {
	set := NewHashSet(1000)
	for i := 0; i < b.N; i++ {
		set.Add(rand.Intn(1000))
	}
}

func BenchmarkHashSetRemove(b *testing.B) {
	set := NewHashSet(1000)
	for i := 0; i < 500; i++ {
		set.Add(rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Remove(rand.Intn(1000))
	}
}

func BenchmarkHashSetContains(b *testing.B) {
	set := NewHashSet(1000)
	for i := 0; i < 500; i++ {
		set.Add(rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(rand.Intn(1000))
	}
}

func BenchmarkHashSetClear(b *testing.B) {
	set := NewHashSet(1000)
	for i := 0; i < 500; i++ {
		set.Add(rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Clear()
	}
}

func BenchmarkHashSetEqual(b *testing.B) {
	setA := NewHashSet(100).Add(1, 3, 7, 88)
	setB := NewHashSet(100).Add(88, 3, 7, 1)
	setC := NewHashSet(100).Add(1, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.Equal(setB)
		setA.Equal(setC)
	}
}

func BenchmarkHashSetBigEqual(b *testing.B) {
	setA := NewHashSet(10000).Add(1, 3, 700, 8888)
	setB := NewHashSet(10000).Add(8888, 3, 700, 1)
	setC := NewHashSet(10000).Add(1, 3, 700, 8889)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.Equal(setB)
		setA.Equal(setC)
	}
}

func BenchmarkHashSetBigSubsetOf(b *testing.B) {
	setA := NewHashSet(10000).Add(3, 700, 8888)
	setB := NewHashSet(10000).Add(8888, 3, 700, 1)
	setC := NewHashSet(10000).Add(1, 3, 700, 8889)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.SubsetOf(setB)
		setA.SubsetOf(setC)
	}
}

func BenchmarkHashSetUnion(b *testing.B) {
	setA := NewHashSet(100).Add(1, 3, 7, 88)
	setB := NewHashSet(100).Add(33, 44, 7, 1)
	setC := NewHashSet(100).Add(13, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Union(setB).Union(setC)
	}
}

func BenchmarkHashSetBigUnion(b *testing.B) {
	setA := NewHashSet(10000)
	setB := NewHashSet(10000)
	for i := 0; i < 5000; i++ {
		setA.Add(rand.Intn(10000))
		setB.Add(rand.Intn(10000))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Union(setB)
	}
}

func BenchmarkHashSetIntersection(b *testing.B) {
	setA := NewHashSet(100).Add(1, 3, 7, 88)
	setB := NewHashSet(100).Add(33, 44, 7, 1)
	setC := NewHashSet(100).Add(13, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Intersection(setB).Intersection(setC)
	}
}

func BenchmarkHashSetSymetricDifference(b *testing.B) {
	setA := NewHashSet(100).Add(1, 3, 7, 88)
	setB := NewHashSet(100).Add(33, 44, 7, 1)
	setC := NewHashSet(100).Add(13, 3, 27, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.SymetricDifference(setB).SymetricDifference(setC)
	}
}

func BenchmarkHashSetBigSymetricDifference(b *testing.B) {
	setA := NewHashSet(10000)
	setB := NewHashSet(10000)
	setC := NewHashSet(10000)
	for i := 0; i < 5000; i++ {
		setA.Add(rand.Intn(10000))
		setB.Add(rand.Intn(10000))
		setC.Add(rand.Intn(10000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.SymetricDifference(setB).SymetricDifference(setC)
	}
}
