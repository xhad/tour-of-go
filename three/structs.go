package three

import "fmt"

// Structs is struct
type Structs struct{}

// Vertex is math
type Vertex struct {
	X int
	Y int
}

func (s Structs) vertex() {
	fmt.Println(Vertex{1, 2})
}

func (s Structs) accessValues() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func (s Structs) pointToStruct() {
	v := Vertex{1, 2}
	p := &v // pointer to v
	p.X = 1e9
	fmt.Println(v)
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func (s Structs) structLiterals() {
	fmt.Println(v1, p, v2, v3)
}
