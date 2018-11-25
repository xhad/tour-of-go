package main

import (
	"github.com/xhad/tour-of-go/one"
	"github.com/xhad/tour-of-go/three"
	"github.com/xhad/tour-of-go/two"
)

func main() {
	var o one.Package
	o.Run()

	var t two.Package
	t.Run()

	var th three.Package
	th.Run()
}
