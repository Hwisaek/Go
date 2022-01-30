package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		// 에러가 발생하면 에러 출력
		fmt.Println(err)

		// '\n' 문자가 나타날 때 까지 버퍼에서 문자를 읽음
		_, _ = stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)

	if err != nil {
		// 에러가 발생하면 에러 출력
		fmt.Println(err)

		// '\n' 문자가 나타날 때 까지 버퍼에서 문자를 읽음
		_, _ = stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
}
