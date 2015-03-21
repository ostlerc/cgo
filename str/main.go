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

type CMemory interface {
	String() string
	Bytes() []byte
	Free()
}

type cstr struct {
	str C.struct_mystring
}

func new(s C.struct_mystring) CMemory {
	return &cstr{str: s}
}

func (c *cstr) String() string {
	return C.GoString(c.str.data)
}

// A no copy data conversion
// this memory will not be gc'd
func (c *cstr) Bytes() []byte {
	hdr := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(c.str.data)),
		Len:  int(c.str.length),
		Cap:  int(c.str.length),
	}
	return *(*[]byte)(unsafe.Pointer(hdr))
}

func (c *cstr) Free() {
	C.del(c.str)
}

func Greet() {
	v := new(C.HelloStr())
	defer v.Free()
	fmt.Println(v.String())
}

func Test() {
	v := new(C.HelloStr())
	defer v.Free()
	fmt.Println(string(v.Bytes()))
}
