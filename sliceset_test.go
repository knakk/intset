package intset

import (
	"math/rand"
	"testing"

	"github.com/knakk/specs"
)

func TestSliceSetAdd(t *testing.T) {
	specs := specs.New(t)

	set := NewSliceSet(100).Add(1, 2, 5, 2)

	specs.Expect(set.Contains(1, 2, 5), true)
	specs.Expect(set.Size(), 3)
}

func TestSliceSetRemove(t *testing.T) {
	specs := specs.New(t)

	set := NewSliceSet(100).Add(3, 1)
	specs.Expect(set.Contains(3, 1), true)
	set.Remove(1, 3)
	specs.Expect(set.Contains(1, 3), false)
}

func TestSliceSetContains(t *testing.T) {
	specs := specs.New(t)

	set := NewSliceSet(100).Add(1, 2, 3).Remove(2)

	specs.Expect(set.Contains(2), false)
	specs.Expect(set.Contains(1, 3), true)
	specs.Expect(set.Contains(1, 2, 3), false)
}

func TestSliceSetClear(t *testing.T) {
	specs := specs.New(t)

	set := NewSliceSet(100).Add(1, 2, 3, 4, 5).Clear()

	specs.Expect(set.Contains(1, 2, 3, 4, 5), false)
	specs.Expect(set.Size(), 0)
}

func TestSliceSetSize(t *testing.T) {
	specs := specs.New(t)

	set := NewSliceSet(100).Add(1, 2, 3)
	specs.Expect(set.Size(), 3)
	set.Remove(2)
	specs.Expect(set.Size(), 2)
}

func TestSliceSetAll(t *testing.T) {
	specs := specs.New(t)

	set := NewSliceSet(100).Add(99, 1, 5)
	all := set.All()

	specs.Expect(len(all), 3)
	specs.Expect(all[0], 1)
	specs.Expect(all[1], 5)
	specs.Expect(all[2], 99)
}

func TestSliceSetEqual(t *testing.T) {
	specs := specs.New(t)

	setA := NewSliceSet(100).Add(1, 2)
	setB := NewSliceSet(100).Add(2, 1, 1)
	setC := NewSliceSet(100).Add(1, 3)

	specs.Expect(setA.Equal(setB), true)
	specs.Expect(setA.Equal(setC), false)
}

func TestSliceSetSubsetOf(t *testing.T) {
	specs := specs.New(t)

	setA := NewSliceSet(100).Add(1, 2)
	setB := NewSliceSet(100).Add(1, 2, 3)
	setC := NewSliceSet(100).Add(3, 4, 5)

	specs.Expect(setA.SubsetOf(setB), true)
	specs.Expect(setA.SubsetOf(setC), false)
}

func TestSliceSetSupersetOf(t *testing.T) {
	specs := specs.New(t)

	setA := NewSliceSet(100).Add(1, 2)
	setB := NewSliceSet(100).Add(1, 2, 3)
	setC := NewSliceSet(100).Add(3, 4, 5)

	specs.Expect(setB.SupersetOf(setA), true)
	specs.Expect(setC.SupersetOf(setA), false)
}

func TestSliceSetUnion(t *testing.T) {
	specs := specs.New(t)

	setA := NewSliceSet(100).Add(1, 2)
	setB := NewSliceSet(100).Add(3, 4)
	setC := NewSliceSet(100).Add(1, 99)

	specs.Expect(setA.Union(setB).Equal(NewSliceSet(100).Add(1, 2, 3, 4)), true)
	specs.Expect(setA.Union(setC).Equal(NewSliceSet(100).Add(1, 2, 99)), true)
}

func TestSliceSetIntersection(t *testing.T) {
	specs := specs.New(t)

	setA := NewSliceSet(100).Add(1, 2)
	setB := NewSliceSet(100).Add(1, 2, 3)
	setC := NewSliceSet(100).Add(3, 4, 5)

	specs.Expect(setA.Intersection(setB).Equal(setA), true)
	specs.Expect(setB.Intersection(setC).Equal(NewSliceSet(100).Add(3)), true)
}

func TestSliceSetSymetricDifference(t *testing.T) {
	specs := specs.New(t)

	setA := NewSliceSet(100).Add(1, 2, 4)
	setB := NewSliceSet(100).Add(1, 2, 3)
	setC := NewSliceSet(100).Add(3, 4, 5)

	specs.Expect(setA.SymetricDifference(setB).Equal(NewSliceSet(100).Add(3, 4)), true)
	specs.Expect(setB.SymetricDifference(setC).Equal(NewSliceSet(100).Add(1, 2, 4, 5)), true)
}

func TestSliceSetClone(t *testing.T) {
	specs := specs.New(t)

	setA := NewSliceSet(100).Add(9, 3, 1)
	setB := setA.Clone()

	specs.Expect(setB.Equal(setA), true)
}

// Benchmarks

func BenchmarkSliceSetAdd(b *testing.B) {
	set := NewSliceSet(1000)
	for i := 0; i < b.N; i++ {
		set.Add(rand.Intn(1000))
	}
}

func BenchmarkSliceSetRemove(b *testing.B) {
	set := NewSliceSet(1000)
	for i := 0; i < 500; i++ {
		set.Add(rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Remove(rand.Intn(1000))
	}
}

func BenchmarkSliceSetContains(b *testing.B) {
	set := NewSliceSet(1000)
	for i := 0; i < 500; i++ {
		set.Add(rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(rand.Intn(1000))
	}
}

func BenchmarkSliceSetClear(b *testing.B) {
	set := NewSliceSet(1000)
	for i := 0; i < 500; i++ {
		set.Add(rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Clear()
	}
}

func BenchmarkSliceSetEqual(b *testing.B) {
	setA := NewSliceSet(100).Add(1, 3, 7, 88)
	setB := NewSliceSet(100).Add(88, 3, 7, 1)
	setC := NewSliceSet(100).Add(1, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.Equal(setB)
		setA.Equal(setC)
	}
}

func BenchmarkSliceSetBigEqual(b *testing.B) {
	setA := NewSliceSet(10000).Add(1, 3, 700, 8888)
	setB := NewSliceSet(10000).Add(8888, 3, 700, 1)
	setC := NewSliceSet(10000).Add(1, 3, 700, 8889)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.Equal(setB)
		setA.Equal(setC)
	}
}

func BenchmarkSliceSetBigSubsetOf(b *testing.B) {
	setA := NewSliceSet(10000).Add(3, 700, 8888)
	setB := NewSliceSet(10000).Add(8888, 3, 700, 1)
	setC := NewSliceSet(10000).Add(1, 3, 700, 8889)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.SubsetOf(setB)
		setA.SubsetOf(setC)
	}
}

func BenchmarkSliceSetUnion(b *testing.B) {
	setA := NewSliceSet(100).Add(1, 3, 7, 88)
	setB := NewSliceSet(100).Add(33, 44, 7, 1)
	setC := NewSliceSet(100).Add(13, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Union(setB).Union(setC)
	}
}

func BenchmarkSliceSetIntersection(b *testing.B) {
	setA := NewSliceSet(100).Add(1, 3, 7, 88)
	setB := NewSliceSet(100).Add(33, 44, 7, 1)
	setC := NewSliceSet(100).Add(13, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Intersection(setB).Intersection(setC)
	}
}

func BenchmarkSliceSetSymetricDifference(b *testing.B) {
	setA := NewSliceSet(100).Add(1, 3, 7, 88)
	setB := NewSliceSet(100).Add(33, 44, 7, 1)
	setC := NewSliceSet(100).Add(13, 3, 27, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.SymetricDifference(setB).SymetricDifference(setC)
	}
}
