package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	rand.Seed(t.UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(100), ", ") // 0이상 100 미만의 랜덤 값
	}
}
