package one

import (
	"fmt"
)

// Package of all stuff
type Package struct{}

// Run runs all the one package methods
func (o Package) Run() {
	fmt.Println("#Package One#")
	fmt.Println(">>NUMBERS")
	var number Number
	fmt.Println("A random number is", number.Random())
	fmt.Printf("\nThe square root of %d is %g", 3, number.SquareRoot(3))
	fmt.Printf("\nThe value of pi is %g", number.Pi())

	var a, b int = 1, -9
	fmt.Printf("\n%d plus %d is %d", a, b, number.Add(a, b))

	fmt.Println(">>STRINGS")
	var strings Strings
	var str1, str2 string
	str1, str2 = "tester", "wonderful"
	fmt.Printf("\n %s and %s\n", str1, str2)
	fmt.Println(strings.Swap(str1, str2))

	fmt.Println(">>VALUES")
	var values Values
	fmt.Println(values.Split(10))
	values.Zero()

	var variables Variables
	fmt.Println(">>VARIABLES")
	variables.Simple()
	variables.Initializers()
	variables.Declarations()
	variables.Constants()
	variables.NumericConst()

	fmt.Println(">>TYPES")
	var types Types
	types.Examples()
	types.Conversions()
	types.Inference()
}
