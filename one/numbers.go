package one

import (
	"math"
	"math/rand"
)

// Number struct
type Number struct{}

// Random prints a random number
func (c Number) Random() int {
	return rand.Intn(10)
}

// SquareRoot returns the square root of the number
func (c Number) SquareRoot(num float64) float64 {
	return math.Sqrt(num)
}

// Pi prints the value of pi
func (c Number) Pi() float64 {
	return math.Pi
}

// Add adds two numbers
func (c Number) Add(a, b int) int {
	var sum int
	sum = a + b
	return sum
}
