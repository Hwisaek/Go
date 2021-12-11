package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length, uppercase = len(name), strings.ToUpper(name)
	return
}

func main() {
	totalLength, _ := lenAndUpper("hwisaek")
	fmt.Println(totalLength)
}
