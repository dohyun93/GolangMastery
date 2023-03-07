package main

import (
	"fmt"
)

func Sum(nums ...int) int {
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
	fmt.Println(Sum(10, 20, 30, 40))
}
