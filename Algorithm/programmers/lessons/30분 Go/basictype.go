package main

import (
	"fmt"
	"math/cmplx"
)

var (
	i      int
	f      float64
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const format = "%T(%v)\n"
	fmt.Printf(format, MaxInt, MaxInt)
	fmt.Printf(format, z, z)

	// int에서 float64로 묵시적 타입 변환을 할 수 없습니다
	// f = i

	// 다른 타입을 저장하려면 변환할타입(변수) 와 같이, 형 변환을 해줘야 합니다.
	f = float64(i)
}
