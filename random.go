package random

import (
	"math/rand"
	"sort"
	//"fmt"
)

func Weighted(r *rand.Rand, w []int) func() int {
	n := len(w)
	acc := make([]int, n)
	sum := 0
	for i, val := range w {
		sum += val
		acc[i] = sum
	}
	gen := func() int {
		randVal := 1 + r.Intn(sum)
		return sort.Search(n,
			func(i int) bool {
				return acc[i] >= randVal
			})
	}
	return gen
}
