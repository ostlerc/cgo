package main

import (
	"fmt"

	"github.com/ostlerc/cgo/ex1"
	"github.com/ostlerc/cgo/str"
)

func main() {
	ex1.Greet()
	str.Greet()
	str.Test()
	fmt.Println("BYE")
}
