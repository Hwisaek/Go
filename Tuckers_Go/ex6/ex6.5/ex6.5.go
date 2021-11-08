package main

import "fmt"

func main() {

	var x int8 = 127

	fmt.Printf("%d < %d+1: %v\n", x, x, x < x+1)
	fmt.Printf("x\t= %4d, %08b\n", x, x)
	fmt.Printf("x+1\t= %4d, %08b\n", x+1, x+1)

	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Println(a + b)
}
