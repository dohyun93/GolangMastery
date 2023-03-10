package main

import (
	"fmt"
)

func Sum(nums ...int) int {
	// 가변 인수를 표시할 때 자료형 앞에 ... 를 붙인다.
	// 함수 내에서는 슬라이스 처럼 동작한다.
	sum := 0
	fmt.Printf("nums 타입 : %T\n", nums)

	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println(Sum(100, 200))
	fmt.Println(Sum(10, 20, 40))
}
