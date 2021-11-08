package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var str string = "Hello World"
	var slice []byte = []byte(str)

	slice[2] = 'a'

	fmt.Println(str)
	fmt.Printf("%s\n", slice)

	stringHEader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t %x\n", stringHEader.Data)
	fmt.Printf("slice:\t %x\n", sliceHeader.Data)
}
