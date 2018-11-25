package three

import "fmt"

// Pointers does pointers
type Pointers struct{}

func (pnts Pointers) show() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // reads i through pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)

	p = &j       // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j)
}
