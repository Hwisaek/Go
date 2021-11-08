package main

import (
	"fmt"
	"strings"
)

// 문자열은 더하기마다 새로운 문자열 생성

// 소문자 -> 대문자 변환
func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "Hello Wolrd"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}
