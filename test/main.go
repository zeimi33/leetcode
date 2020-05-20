package main

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lhello
#include "../test/test.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export goadd
func goadd(a unsafe.Pointer) {
	fmt.Println(a)
	//	v := reflect.ValueOf(pot)
	d := (*stu)(a)
	fmt.Println(d)
	fmt.Println(*d)
}

type stu struct {
	age  int
	name string
}

func main() {
	u := stu{}
	u.age = 10
	u.name = "666"
	fmt.Println("地址", (unsafe.Pointer)(&u))
	a := (unsafe.Pointer)(&u)
	fmt.Println(a)
	C.hello(a)
	fmt.Println(a)
}
