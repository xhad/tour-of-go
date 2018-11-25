package two

import (
	"fmt"
	"math"
)

// If struct
type If struct{}

func (i If) sqrt(x float64) string {
	if x < 0 {
		return i.sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func (i If) pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
