package main

import (
	"sort"
	"sync"
	"testing"
)

func BenchmarkMakeRandomList(b *testing.B) {
	size := 1000000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randomList(size, 100)
	}
}

func BenchmarkSortUsingGorutine(b *testing.B) {
	size := 1000000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		list := randomList(size, 100)

		l, sep := divideList(list)
		wg.Add(2)
		go sortList(l[:sep], &wg)
		go sortList(l[sep:], &wg)
		wg.Wait()
	}
}

func BenchmarkSortUsingGorutineEnhanced(b *testing.B) {
	size := 1000000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		list := randomList(size, 100)

		l, sep := divideListEnhanced(list)
		wg.Add(2)
		go sortList(l[:sep], &wg)
		go sortList(l[sep:], &wg)
		wg.Wait()
	}
}

func BenchmarkSortBasic(b *testing.B) {
	size := 1000000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list := randomList(size, 100)
		sort.Ints(list)
	}
}
