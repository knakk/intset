package intset

import (
	"math/rand"
	"testing"

	"github.com/knakk/specs"
)

func TestBitSetAdd(t *testing.T) {
	specs := specs.New(t)

	set := NewBitSet(10).Add(1, 2, 5, 2)

	specs.Expect(set.Contains(1, 2, 5), true)
	specs.Expect(set.Size(), 3)
}

func TestBitSetRemove(t *testing.T) {
	specs := specs.New(t)

	set := NewBitSet(10).Add(3, 1)
	specs.Expect(set.Contains(3, 1), true)
	set.Remove(1, 3)
	specs.Expect(set.Contains(1, 3), false)
}

func TestBitSetContains(t *testing.T) {
	specs := specs.New(t)

	set := NewBitSet(10).Add(1, 2, 3).Remove(2)

	specs.Expect(set.Contains(2), false)
	specs.Expect(set.Contains(1, 3), true)
	specs.Expect(set.Contains(1, 2, 3), false)
}

func TestBitSetClear(t *testing.T) {
	specs := specs.New(t)

	set := NewBitSet(10).Add(1, 2, 3, 4, 5).Clear()

	specs.Expect(set.Contains(1, 2, 3, 4, 5), false)
	specs.Expect(set.Size(), 0)
}

func TestBitSetSize(t *testing.T) {
	specs := specs.New(t)

	set := NewBitSet(10).Add(11, 2, 3)
	specs.Expect(set.Size(), 3)
	set.Remove(2)
	specs.Expect(set.Size(), 2)
}

func TestBitSetAll(t *testing.T) {
	specs := specs.New(t)

	set := NewBitSet(10).Add(99, 1, 5)
	all := set.All()

	specs.Expect(len(all), 3)
	specs.Expect(all[0], 1)
	specs.Expect(all[1], 5)
	specs.Expect(all[2], 99)
}

func TestBitSetEqual(t *testing.T) {
	specs := specs.New(t)

	setA := NewBitSet(10).Add(1, 2)
	setB := NewBitSet(10).Add(2, 1, 1)
	setC := NewBitSet(10).Add(1, 3)

	specs.Expect(setA.Equal(setB), true)
	specs.Expect(setA.Equal(setC), false)
}

func TestBitSetSubsetOf(t *testing.T) {
	specs := specs.New(t)

	setA := NewBitSet(10).Add(1, 2)
	setB := NewBitSet(10).Add(1, 2, 3)
	setC := NewBitSet(10).Add(3, 4, 5)

	specs.Expect(setA.SubsetOf(setB), true)
	specs.Expect(setA.SubsetOf(setC), false)
}

func TestBitSetSupersetOf(t *testing.T) {
	specs := specs.New(t)

	setA := NewBitSet(10).Add(1, 2)
	setB := NewBitSet(10).Add(1, 2, 3)
	setC := NewBitSet(10).Add(3, 4, 5)

	specs.Expect(setB.SupersetOf(setA), true)
	specs.Expect(setC.SupersetOf(setA), false)
}

func TestBitSetUnion(t *testing.T) {
	specs := specs.New(t)

	setA := NewBitSet(10).Add(1, 2)
	setB := NewBitSet(10).Add(3, 4)
	setC := NewBitSet(10).Add(1, 99)

	specs.Expect(setA.Union(setB).Equal(NewBitSet(10).Add(1, 2, 3, 4)), true)
	specs.Expect(setA.Union(setC).Equal(NewBitSet(10).Add(1, 2, 99)), true)
}

func TestBitSetIntersection(t *testing.T) {
	specs := specs.New(t)

	setA := NewBitSet(10).Add(1, 2)
	setB := NewBitSet(10).Add(1, 2, 3)
	setC := NewBitSet(10).Add(3, 4, 5)

	specs.Expect(setA.Intersection(setB).Equal(setA), true)
	specs.Expect(setB.Intersection(setC).Equal(NewBitSet(10).Add(3)), true)
}

func TestBitSetSymetricDifference(t *testing.T) {
	specs := specs.New(t)

	setA := NewBitSet(10).Add(1, 2, 4)
	setB := NewBitSet(10).Add(1, 2, 3)
	setC := NewBitSet(10).Add(3, 4, 5)

	specs.Expect(setA.SymetricDifference(setB).Equal(NewBitSet(10).Add(3, 4)), true)
	specs.Expect(setB.SymetricDifference(setC).Equal(NewBitSet(10).Add(1, 2, 4, 5)), true)
}

func TestBitSetClone(t *testing.T) {
	specs := specs.New(t)

	setA := NewBitSet(10).Add(9, 3, 1)
	setB := setA.Clone()

	specs.Expect(setB.Equal(setA), true)
}

// Benchmarks

func BenchmarkBitSetAdd(b *testing.B) {
	set := NewBitSet(1000)
	for i := 0; i < b.N; i++ {
		set.Add(rand.Intn(1000))
	}
}

func BenchmarkBitSetRemove(b *testing.B) {
	set := NewBitSet(1000)
	for i := 0; i < 500; i++ {
		set.Add(rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Remove(rand.Intn(1000))
	}
}

func BenchmarkBitSetContains(b *testing.B) {
	set := NewBitSet(1000)
	for i := 0; i < 500; i++ {
		set.Add(rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(rand.Intn(1000))
	}
}

func BenchmarkBitSetClear(b *testing.B) {
	set := NewBitSet(1000)
	for i := 0; i < 500; i++ {
		set.Add(rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Clear()
	}
}

func BenchmarkBitSetEqual(b *testing.B) {
	setA := NewBitSet(100).Add(1, 3, 7, 88)
	setB := NewBitSet(100).Add(88, 3, 7, 1)
	setC := NewBitSet(100).Add(1, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.Equal(setB)
		setA.Equal(setC)
	}
}

func BenchmarkBitSetBigEqual(b *testing.B) {
	setA := NewBitSet(10000).Add(1, 3, 700, 8888)
	setB := NewBitSet(10000).Add(8888, 3, 700, 1)
	setC := NewBitSet(10000).Add(1, 3, 700, 8889)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.Equal(setB)
		setA.Equal(setC)
	}
}

func BenchmarkBitSetBigSubsetOf(b *testing.B) {
	setA := NewBitSet(10000).Add(3, 700, 8888)
	setB := NewBitSet(10000).Add(8888, 3, 700, 1)
	setC := NewBitSet(10000).Add(1, 3, 700, 8889)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.SubsetOf(setB)
		setA.SubsetOf(setC)
	}
}

func BenchmarkBitSetUnion(b *testing.B) {
	setA := NewBitSet(100).Add(1, 3, 7, 88)
	setB := NewBitSet(100).Add(33, 44, 7, 1)
	setC := NewBitSet(100).Add(13, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Union(setB).Union(setC)
	}
}

func BenchmarkBitSetBigUnion(b *testing.B) {
	setA := NewBitSet(10000)
	setB := NewBitSet(10000)
	for i := 0; i < 5000; i++ {
		setA.Add(rand.Intn(10000))
		setB.Add(rand.Intn(10000))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Union(setB)
	}
}

func BenchmarkBitSetIntersection(b *testing.B) {
	setA := NewBitSet(100).Add(1, 3, 7, 88)
	setB := NewBitSet(100).Add(33, 44, 7, 1)
	setC := NewBitSet(100).Add(13, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Intersection(setB).Intersection(setC)
	}
}

func BenchmarkBitSetSymetricDifference(b *testing.B) {
	setA := NewBitSet(100).Add(1, 3, 7, 88)
	setB := NewBitSet(100).Add(33, 44, 7, 1)
	setC := NewBitSet(100).Add(13, 3, 27, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.SymetricDifference(setB).SymetricDifference(setC)
	}
}

func BenchmarkBitSetBigSymetricDifference(b *testing.B) {
	setA := NewBitSet(10000)
	setB := NewBitSet(10000)
	setC := NewBitSet(10000)
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
