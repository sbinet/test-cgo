package main

import (
	"fmt"

	"github.com/sbinet/test-cgo/pkg"
)

func main() {
	fmt.Printf(">>>\n")
	pkg.Print()
	fmt.Printf("<<<\n")
}
