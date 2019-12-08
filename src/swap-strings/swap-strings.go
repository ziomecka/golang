package main

import (
	"fmt"
)

func swapStrings(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swapStrings("hello", "world")
	fmt.Println("Print", a, b)
}
