package three

import (
	"fmt"
	"strings"
)

// Arrays does arrays
type Arrays struct{}

func (arr Arrays) helloWorld() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func (arr Arrays) slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	s := primes[1:4]
	fmt.Println(s)
}

func (arr Arrays) slicesOfNames() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func (arr Arrays) sliceLiteral() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func (arr Arrays) sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func (arr Arrays) printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func (arr Arrays) lengthCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	arr.printSlice(s)
	// slice the slice to give it zero length
	s = s[:0]
	arr.printSlice(s)

	// Extend its length
	s = s[:4]
	arr.printSlice(s)

	// drop the first two values
	s = s[2:]
	arr.printSlice(s)
}

func (arr Arrays) nilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func (arr Arrays) printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func (arr Arrays) useMake() {
	a := make([]int, 5)
	arr.printSlice2("a", a)

	b := make([]int, 0, 5)
	arr.printSlice2("b", b)

	c := b[:2]
	arr.printSlice2("c", c)

	d := c[2:5]
	arr.printSlice2("d", d)

}

func (arr Arrays) sliceOfSlice() {
	// create a tic tack toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// the player takes a turn.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func (arr Arrays) appendSlice() {
	var s []int
	arr.printSlice(s)

	// appned works on nil slices.
	s = append(s, 0)
	arr.printSlice(s)

	// the slice grows as needed
	s = append(s, 1)
	arr.printSlice(s)

	// we can add more than one element at a time
	s = append(s, 2, 3, 4)
	arr.printSlice(s)
}
