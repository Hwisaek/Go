package main

import "fmt"

// 상수도 선언과 동시에 초기화하면 타입을 지정하지 않아도 됩니다.
const Pi1 float32 = 3.14
const Pi2 = 3.14

// 괄호로 묶으면 상수 키워드를 한 번만 명시합니다.
const (
	Big_const   = 1 << 100
	Small_const = Big_const >> 99
)

// 오버플로우가 발생합니다.
// var Big_var = 1 << 100

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("needInt(Small_const):", needInt(Small_const))     // Small_const 상수 선언시 타입을 지정하지 않았기 때문에 int형으로 자동으로 변환됩니다
	fmt.Println("needFloat(Small_const):", needFloat(Small_const)) // Small_const 상수 선언시 타입을 지정하지 않았기 때문에 float64형으로 자동으로 변환됩니다
	fmt.Println("needFloat(Big_const):", needFloat(Big_const))

}
