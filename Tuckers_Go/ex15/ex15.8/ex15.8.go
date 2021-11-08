package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str := "Hello 월드"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입: %T, 값: %d, 문자값: %c, 크기: %v\n", arr[i], arr[i], arr[i], unsafe.Sizeof(arr[i]))
	}
}
