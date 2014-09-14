package random

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
)

func Test(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	a := []int {10, 10, 20}
	gen := Weighted(r, a)
	hist := make(map[int]int)
	trials := 10000
	for i := 0; i < trials; i++ {
		//fmt.Println(r.Float32())
		hist[gen()]++
	}
	for k, v := range hist {
		fmt.Printf("%v : %v\n", k, float64(v)/float64(trials))
	}
}