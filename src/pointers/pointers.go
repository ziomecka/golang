package main

import (
	"fmt"
)

func main() {
	a := 0x00
	b := 0x0A
	// var pa *int
	// pa = &a
	pa := &a
	pb := &b

	*pa = 100
	*pb = 200
	fmt.Println("Memory address of a is", &a, a)
	fmt.Println("Memory address of b is", &b, b)
	fmt.Println("Pointer", pa, *pa)
	fmt.Println("Pointer b", pb, *pb)
}
