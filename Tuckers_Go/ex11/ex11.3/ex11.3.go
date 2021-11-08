package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("숫자를 입력하세요")
		var number int
		_, err := fmt.Scanln(&number)

		if err != nil {
			fmt.Println("숫자로 입력하세요")

			// 키보드 버퍼를 지웁니다.
			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다.", number)
		if number%2 == 0 {
			break // 짝수면 멈춤
		}

	}
}
