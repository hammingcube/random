package random

import (
	"math/rand"
	"fmt"
)

var constantFunc func() int = func() int {
			return 1
		}

func Weighted(w []int) func() int {
	if len(w) < 1 {
		return constantFunc
	}
	
	acc := make([]int, len(w))
	sum := 0
	for i, val := range(w) {
		sum += val
		acc[i] = sum
	}
	fmt.Println(acc)

	i := rand.Intn(len(w))
	return func() int {
		return i
	}
}
