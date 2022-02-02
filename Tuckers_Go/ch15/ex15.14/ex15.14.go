package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello 월드"
	str2 := str1

	// str1 의 메모리 주소(&str1)을 unsafe.Pointer타입으로 변환.
	// *string -> unsafePointer -> *reflect.StringHeader
	// *string -//-> *reflect.StringHeader. 한번에 변환 불가능
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}
