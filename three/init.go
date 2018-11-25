package three

import "fmt"

// Package three
type Package struct{}

// Run starts and runs Package three stuff
func (pkg Package) Run() {

	fmt.Println("Pointers Package")

	fmt.Println(">>POINTERS")
	var p Pointers
	p.show()

	fmt.Println(">>STRUCTS")
	var s Structs
	s.vertex()
	s.accessValues()
	s.pointToStruct()
	s.structLiterals()

	fmt.Println(">>ARRAYS")
	var a Arrays
	a.helloWorld()
	a.slices()
	a.slicesOfNames()
	a.sliceLiteral()
	a.sliceDefaults()
}
