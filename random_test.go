package random

import (
	"testing"
)

func Test(t *testing.T) {
	a := make([]int, 5)
	rndInd := Weighted(a)
	rndInd()
}