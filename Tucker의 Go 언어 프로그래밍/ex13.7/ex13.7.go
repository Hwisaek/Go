package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 1
	C int8 // 1
	E int8 // 1
	B int  // 8
	D int  // 8
} // 19 => 24 : 메모리패딩

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
