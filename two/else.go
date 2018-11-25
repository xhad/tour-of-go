package two

import (
	"fmt"
	"math"
)

// Else is stuff
type Else struct{}

func (e Else) pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, return lim
	return lim
}
