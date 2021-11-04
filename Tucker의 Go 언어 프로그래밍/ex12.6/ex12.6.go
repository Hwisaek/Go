package main

import "fmt"

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
	}

	b := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
	}

	b = a

	b[0][1] = 3

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
