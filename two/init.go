package two

import "fmt"

// Package two
type Package struct{}

// Run the stuff
func (p Package) Run() {
	fmt.Println("\n#Package Two#")

	fmt.Println(">>LOOP")
	var f For
	f.Loop()
	f.While()
	f.exercise1(2)

	fmt.Println(">>IF")
	var i If
	fmt.Println(i.sqrt(-10))
	fmt.Println(i.sqrt(10))
	fmt.Println(
		i.pow(3, 2, 10),
		i.pow(3, 6, 10),
	)

	fmt.Println(">>ELSE")
	var e Else
	fmt.Println(
		e.pow2(3, 2, 10),
		e.pow2(3, 3, 20),
	)

	fmt.Println(">>SWITCH")
	var s Switch
	s.printOs()
	s.whensSaturday()
	s.greeting()

	fmt.Println(">>DEFER")
	var d Defer
	d.hello()
	d.stack()
}
