package main

import (
	"fmt"
	"math"
)

func main() {
	var i int
	var o float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, o, b, s)
	i=30203010
	fmt.Printf("%x %v %v %p\n", i, i, i, i)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	p := 42.10230303 // change me!
	fmt.Printf("p is of type %T\n", p)

	w := 42+.3i  // change me!
	fmt.Printf("w is of type %T\n", w)
}

