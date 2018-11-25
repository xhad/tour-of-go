package one

import "fmt"

// Values returns magic shit
type Values struct{}

// Split returns this magically return
func (v Values) Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Zero returns zero values
func (v Values) Zero() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
