package main

import (
	"bufio"
	"fmt"
	"os"
)

func Add(a, b int) int { // 매개변수 a, b
	return a + b
}

func fibonacci(a int) int {
	if a == 1 || a == 2 {
		return 1
	} else {
		return fibonacci(a-1) + fibonacci(a-2)
	}
}

func main() {
	// 1. 함수 기초
	a, b := 3, 4
	c := Add(a, b) // 인수 a, b
	fmt.Println("Add result : ", c)

	// 2. 재귀 호출
	// fibonacci x번째 원소값을 확인할 때 재귀함수로 확인해보자.
	// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...
	stdin := bufio.NewReader(os.Stdin)
	var x int
	for true {
		n, err := fmt.Scanln(&x)
		if err != nil {
			fmt.Println(n, " number of valid input. Error occurred because of \"", err, "\"")
			stdin.ReadString('\n')
		} else {
			fmt.Println("Input fine! fibonacci input x is : ", x)
			break
		}
	}

	result := fibonacci(x)

	fmt.Println(x, "-th fibonacci value is : ", result)
}
