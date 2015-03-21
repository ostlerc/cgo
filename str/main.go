package str

/*
#cgo CPPFLAGS: -I./cpp
#cgo CXXFLAGS: -std=c++11 -Wall
#include <test.h>
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func Greet() {
	fmt.Println("HI!")
	v := C.HelloStr()
	defer C.del(v)
	fmt.Println(C.GoString(v.data))
}

func Test() {
	fmt.Println(string(Data()))
}

// A no copy cgo unsafe to slice
// The memory of the slice depends on the internal unsafe data
func Data() []byte {
	v := C.HelloStr()
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(v.data)),
		Len:  int(v.length),
		Cap:  int(v.length),
	}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}
