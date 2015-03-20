package ex1

/*
#cgo CPPFLAGS: -I./cpp
#cgo CXXFLAGS: -std=c++11 -Wall
#include <test.h>
*/
import "C"
import (
	"fmt"
)

func Greet() {
	fmt.Println("HI!")
	C.Say_hello()
}
