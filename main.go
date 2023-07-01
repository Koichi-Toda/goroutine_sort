package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func randomList(size int, n int) []int {
	l := make([]int, size)
	// rand.Seed has been deprecated since Go 1.20
	for i := range l {
		l[i] = rand.Intn(n)
	}
	return l
}

func getMax(l []int) int {
	max := 0
	for _, v := range l {
		if max < v {
			max = v
		}
	}
	return max
}

func divideList(l []int) ([]int, int) {
	divl := make([]int, len(l))
	i, j := 0, len(l)-1
	max := getMax(l)

	for _, v := range l {
		if v <= max/2 {
			divl[i] = v
			i++
		} else {
			divl[j] = v
			j--
		}
	}

	return divl, i
}

func divideListEnhanced(l []int) ([]int, int) {
	i, j := 0, len(l)-1
	pivot := getMax(l) / 2
	for {
		for l[i] <= pivot {
			i++
		}
		for l[j] > pivot {
			j--
		}
		if i >= j {
			return l, i
		}
		// swap
		l[i], l[j] = l[j], l[i]

		i++
		j--
	}
}

func sortList(l []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(l)
}

func main() {
	var wg sync.WaitGroup
	size := 10000000
	list := randomList(size, 100)
	list2 := make([]int, len(list))
	copy(list2, list)

	now := time.Now()
	l, sep := divideListEnhanced(list2)
	wg.Add(2)
	go sortList(l[:sep], &wg)
	go sortList(l[sep:], &wg)
	wg.Wait()
	elapsed1 := time.Since(now).Milliseconds()
	fmt.Printf("elapsed time(sort using goroutine): %vms\n", elapsed1)

	now = time.Now()
	sort.Ints(list)
	elapsed2 := time.Since(now).Milliseconds()
	fmt.Printf("elapsed time(basic sort): %vms\n", elapsed2)
}
