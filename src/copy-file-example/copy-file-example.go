package main

import (
	"copyfile"
	"fmt"
)

var (
	srcFile  string = "sourceFile.txt"
	destFile string = "destFile.txt"
)

func main() {
	result, err := copyfile.Copy(srcFile, destFile)
	fmt.Printf("Copying result %q, %q", result, err)
}
