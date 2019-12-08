package main

import (
	"fmt"
)

func namedReturnValues(a, b int) (x, y int) {
	x = a / b
	y = b - x
	return
}
func main() {
	fmt.Println(namedReturnValues(10, 3))
}
