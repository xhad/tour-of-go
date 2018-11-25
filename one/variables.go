package one

import "fmt"

// Variables stuff
type Variables struct{}

// Simple returns variables stuff
func (vbl Variables) Simple() {
	var c, python, java bool
	var i int
	fmt.Println(i, c, python, java)
}

// Initializers prints some stuff
func (vbl Variables) Initializers() {
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

// Declarations prints declarations
func (vbl Variables) Declarations() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

// Constants can be a character, string, boolean or number
func (vbl Variables) Constants() {
	const Pi = 3.14
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("go rules?", Truth)
}

// NumericConst are number constants
func (vbl Variables) NumericConst() {
	const (
		// Create a huge numer by shifting 1 bit left 100 places
		// In other words, the binary number that is 1 followed by 100 zeroes
		Big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2
		Small = Big >> 99
	)

	fmt.Println(Small*10 + 1)
	fmt.Println(Small * 0.1)
	fmt.Println(Big * 0.1)

}
