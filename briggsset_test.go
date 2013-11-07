package intset

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/knakk/specs"
)

func TestBriggsSetAdd(t *testing.T) {
	specs := specs.New(t)

	set := NewBriggsSet(100).Add(1, 2, 5, 2)

	specs.Expect(set.Contains(1, 2, 5), true)
	specs.Expect(set.Size(), 3)
}

func TestBriggsSetRemove(t *testing.T) {
	specs := specs.New(t)

	set := NewBriggsSet(100).Add(3, 1)
	specs.Expect(set.Contains(3, 1), true)
	set.Remove(1, 3)
	specs.Expect(set.Contains(1, 3), false)
}

func TestBriggsSetContains(t *testing.T) {
	specs := specs.New(t)

	set := NewBriggsSet(100).Add(1, 2, 3).Remove(2)

	specs.Expect(set.Contains(2), false)
	specs.Expect(set.Contains(1, 3), true)
	specs.Expect(set.Contains(1, 2, 3), false)
}

func TestBriggsSetClear(t *testing.T) {
	specs := specs.New(t)

	set := NewBriggsSet(100).Add(1, 2, 3, 4, 5).Clear()

	specs.Expect(set.Contains(1, 2, 3, 4, 5), false)
	specs.Expect(set.Size(), 0)
}

func TestBriggsSetSize(t *testing.T) {
	specs := specs.New(t)

	set := NewBriggsSet(100).Add(1, 2, 3)
	specs.Expect(set.Size(), 3)
	set.Remove(2)
	specs.Expect(set.Size(), 2)
}

func TestBriggsSetAll(t *testing.T) {
	specs := specs.New(t)

	set := NewBriggsSet(100).Add(99, 1, 5)
	all := set.All()

	specs.Expect(len(all), 3)
	sort.Ints(all) // Hashet.All() doesn't guarantee order of ints
	specs.Expect(all[0], 1)
	specs.Expect(all[1], 5)
	specs.Expect(all[2], 99)
}

func TestBriggsSetEqual(t *testing.T) {
	specs := specs.New(t)

	setA := NewBriggsSet(100).Add(1, 2)
	setB := NewBriggsSet(100).Add(2, 1, 1)
	setC := NewBriggsSet(100).Add(1, 3)

	specs.Expect(setA.Equal(setB), true)
	specs.Expect(setA.Equal(setC), false)
}

func TestBriggsSetSubsetOf(t *testing.T) {
	specs := specs.New(t)

	setA := NewBriggsSet(100).Add(1, 2)
	setB := NewBriggsSet(100).Add(1, 2, 3)
	setC := NewBriggsSet(100).Add(3, 4, 5)

	specs.Expect(setA.SubsetOf(setB), true)
	specs.Expect(setA.SubsetOf(setC), false)
}

func TestBriggsSetSupersetOf(t *testing.T) {
	specs := specs.New(t)

	setA := NewBriggsSet(100).Add(1, 2)
	setB := NewBriggsSet(100).Add(1, 2, 3)
	setC := NewBriggsSet(100).Add(3, 4, 5)

	specs.Expect(setB.SupersetOf(setA), true)
	specs.Expect(setC.SupersetOf(setA), false)
}

func TestBriggsSetUnion(t *testing.T) {
	specs := specs.New(t)

	setA := NewBriggsSet(100).Add(1, 2)
	setB := NewBriggsSet(100).Add(3, 4)
	setC := NewBriggsSet(100).Add(1, 99)

	specs.Expect(setA.Union(setB).Equal(NewBriggsSet(100).Add(1, 2, 3, 4)), true)
	specs.Expect(setA.Union(setC).Equal(NewBriggsSet(100).Add(1, 2, 99)), true)
}

func TestBriggsSetIntersection(t *testing.T) {
	specs := specs.New(t)

	setA := NewBriggsSet(100).Add(1, 2)
	setB := NewBriggsSet(100).Add(1, 2, 3)
	setC := NewBriggsSet(100).Add(3, 4, 5)

	specs.Expect(setA.Intersection(setB).Equal(setA), true)
	specs.Expect(setB.Intersection(setC).Equal(NewBriggsSet(100).Add(3)), true)
}

func TestBriggsSetSymetricDifference(t *testing.T) {
	specs := specs.New(t)

	setA := NewBriggsSet(100).Add(1, 2, 4)
	setB := NewBriggsSet(100).Add(1, 2, 3)
	setC := NewBriggsSet(100).Add(3, 4, 5)

	specs.Expect(setA.SymetricDifference(setB).Equal(NewBriggsSet(100).Add(3, 4)), true)
	specs.Expect(setB.SymetricDifference(setC).Equal(NewBriggsSet(100).Add(1, 2, 4, 5)), true)
}

func TestBriggsSetClone(t *testing.T) {
	specs := specs.New(t)

	setA := NewBriggsSet(100).Add(9, 3, 1)
	setB := setA.Clone()

	specs.Expect(setB.Equal(setA), true)
}

// Benchmarks

func BenchmarkBriggsSetAdd(b *testing.B) {
	set := NewBriggsSet(1000)
	for i := 0; i < b.N; i++ {
		set.Add(rand.Intn(1000))
	}
}

func BenchmarkBriggsSetRemove(b *testing.B) {
	set := NewBriggsSet(1000)
	for i := 0; i < 500; i++ {
		set.Add(rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Remove(rand.Intn(1000))
	}
}

func BenchmarkBriggsSetContains(b *testing.B) {
	set := NewBriggsSet(1000)
	for i := 0; i < 500; i++ {
		set.Add(rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(rand.Intn(1000))
	}
}

func BenchmarkBriggsSetClear(b *testing.B) {
	set := NewBriggsSet(1000)
	for i := 0; i < 500; i++ {
		set.Add(rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Clear()
	}
}

func BenchmarkBriggsSetEqual(b *testing.B) {
	setA := NewBriggsSet(100).Add(1, 3, 7, 88)
	setB := NewBriggsSet(100).Add(88, 3, 7, 1)
	setC := NewBriggsSet(100).Add(1, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.Equal(setB)
		setA.Equal(setC)
	}
}

func BenchmarkBriggsSetBigEqual(b *testing.B) {
	setA := NewBriggsSet(10000).Add(1, 3, 700, 8888)
	setB := NewBriggsSet(10000).Add(8888, 3, 700, 1)
	setC := NewBriggsSet(10000).Add(1, 3, 700, 8889)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.Equal(setB)
		setA.Equal(setC)
	}
}

func BenchmarkBriggsSetBigSubsetOf(b *testing.B) {
	setA := NewBriggsSet(10000).Add(3, 700, 8888)
	setB := NewBriggsSet(10000).Add(8888, 3, 700, 1)
	setC := NewBriggsSet(10000).Add(1, 3, 700, 8889)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.SubsetOf(setB)
		setA.SubsetOf(setC)
	}
}

func BenchmarkBriggsSetUnion(b *testing.B) {
	setA := NewBriggsSet(100).Add(1, 3, 7, 88)
	setB := NewBriggsSet(100).Add(33, 44, 7, 1)
	setC := NewBriggsSet(100).Add(13, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Union(setB).Union(setC)
	}
}

func BenchmarkBriggsSetBigUnion(b *testing.B) {
	setA := NewBriggsSet(10000)
	setB := NewBriggsSet(10000)
	for i := 0; i < 5000; i++ {
		setA.Add(rand.Intn(10000))
		setB.Add(rand.Intn(10000))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Union(setB)
	}
}

func BenchmarkBriggsSetIntersection(b *testing.B) {
	setA := NewBriggsSet(100).Add(1, 3, 7, 88)
	setB := NewBriggsSet(100).Add(33, 44, 7, 1)
	setC := NewBriggsSet(100).Add(13, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Intersection(setB).Intersection(setC)
	}
}

func BenchmarkBriggsSetSymetricDifference(b *testing.B) {
	setA := NewBriggsSet(100).Add(1, 3, 7, 88)
	setB := NewBriggsSet(100).Add(33, 44, 7, 1)
	setC := NewBriggsSet(100).Add(13, 3, 27, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.SymetricDifference(setB).SymetricDifference(setC)
	}
}

func BenchmarkBriggsSetBigSymetricDifference(b *testing.B) {
	setA := NewBriggsSet(10000)
	setB := NewBriggsSet(10000)
	setC := NewBriggsSet(10000)
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
