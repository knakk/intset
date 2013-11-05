package intset

import (
	"math/rand"
	"testing"
)

func BenchmarkHashSetAdd(b *testing.B) {
	set := NewHashSet(1000)
	for i := 0; i < b.N; i++ {
		set.Add(rand.Intn(1000))
	}
}

func BenchmarkSliceSetAdd(b *testing.B) {
	set := NewSliceSet(1000)
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

func BenchmarkHashSetUnion(b *testing.B) {
	setA := NewHashSet(100).Add(1, 3, 7, 88)
	setB := NewHashSet(100).Add(33, 44, 7, 1)
	setC := NewHashSet(100).Add(13, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Union(setB).Union(setC)
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

func BenchmarkHashSetIntersection(b *testing.B) {
	setA := NewHashSet(100).Add(1, 3, 7, 88)
	setB := NewHashSet(100).Add(33, 44, 7, 1)
	setC := NewHashSet(100).Add(13, 3, 7, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Intersection(setB).Intersection(setC)
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

func BenchmarkHashSetSymetricDifference(b *testing.B) {
	setA := NewHashSet(100).Add(1, 3, 7, 88)
	setB := NewHashSet(100).Add(33, 44, 7, 1)
	setC := NewHashSet(100).Add(13, 3, 27, 89)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.SymetricDifference(setB).SymetricDifference(setC)
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
