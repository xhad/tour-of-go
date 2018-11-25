package two

import (
	"fmt"
	"math"
)

// For loops and stuff
type For struct{}

// Loop example
func (f For) Loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	val := 5
	for j := 0; j < 30; j++ {
		val -= j
	}
	fmt.Println(val)

	bal := 1
	for bal < 1000 {
		bal += bal
	}
	fmt.Println(bal)
}

// While in go
func (f For) While() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

/* infinite loop

for {}

*/

func (f For) exercise1(x float64) float64 {
	z := float64(1)
	var flag float64
	for i := 10; i > 0; i-- {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Value is: %f Previous value is: %f\n", z, flag)

		if math.Round(flag) == math.Round(z) {
			fmt.Printf("The final value is: %f\n", z)
			i = 0
		}
		flag = z
	}

	return z
}
