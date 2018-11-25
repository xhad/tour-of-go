package two

import "fmt"

// defer defers the execution of
// a statement until the
// surrounding function returns

// Defer returns hello world
type Defer struct{}

func (d Defer) hello() {
	defer fmt.Println("world!")
	fmt.Println("hello")
}

func (d Defer) stack() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
