package one

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Types is for types
type Types struct{}

var (
	// ToBe is a bool
	ToBe bool
	// MaxInt is an int
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

// Examples shows some examples of types
func (t Types) Examples() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// Conversions prints some conversions
func (t Types) Conversions() {
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
}

// Inference returns stuff
func (t Types) Inference() {
	var i int
	j := i            // j is in int
	h := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("The types are infered: %d %d %d %f %g", i, j, h, f, g)
}
