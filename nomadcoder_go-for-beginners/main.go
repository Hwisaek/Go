package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	length, uppercase = len(name), strings.ToUpper(name)
	return
}

func main() {
	totalLength, _ := lenAndUpper("hwisaek")
	fmt.Println(totalLength)
}
