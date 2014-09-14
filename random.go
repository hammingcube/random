package random

import (
	"math/rand"
)

func Discrete(weights []int) (index int) {
	index = rand.Intn(len(weights))
	return
}
