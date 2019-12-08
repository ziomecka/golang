package main

import (
	"fmt"
	mathRand "math/rand"
	"time"
)

func main() {
	mathRand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("My favorite number is", mathRand.Intn(10))
}
