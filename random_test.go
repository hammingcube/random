package random

import (
	"testing"
)

func Test(t *testing.T) {
	a := []int {1, 2, 3, 4, 5}
	rndInd := Weighted(a)
	rndInd()
}