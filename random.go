package random

import (
	"math/rand"
	"fmt"
)

func Weighted(w []int) func() int {
	fmt.Println("Hi from the random package!")
	i := rand.Intn(len(w))
	return func() int {
		return i
	}
}
