package main

import "fmt"

//① return 뒤에 리턴 타입을 적어주는 방법
func divide1(dividend, divisor int) (int, int) {
	var quotient = (int)(dividend / divisor)
	var remainder = dividend % divisor
	return quotient, remainder
}

//② return뒤에 리턴할 변수를 선언하는 방법. ①과는 달리 함수 내부에서 `quotient`를 `var`로 선언하지 않고 바로 씁니다.
func divide2(dividend, divisor int) (quotient, remainder int) {
	quotient = (int)(dividend / divisor)
	remainder = dividend % divisor
	return //return이라고만 적으면 미리 return값으로 정해 놓은 quotient와 remainder를 return합니다.
}

func main() {
	//①로 한 번에 여러개의 결과를 return받는 부분
	var quotient, remainder int
	quotient, remainder = divide1(10, 3)
	fmt.Println("①의 결과:", quotient, remainder)

	//②로 한 번에 여러개의 결과를 return받는 부분
	quotient, remainder = divide2(10, 3)
	fmt.Println("②의 결과:", quotient, remainder)
}
