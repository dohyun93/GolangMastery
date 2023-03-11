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

// 여러 타입의 인수들을 받아서 출력하는 함수를 만들어보자.
func Print(args ...interface{}) {
	for _, val := range args {
		switch T := val.(type) {
		case bool:
			value := val.(bool)
			fmt.Println("boolean data print!: ", value)
		case int64:
			value := val.(int64)
			fmt.Println("int64 data print!: ", value)
		default:
			fmt.Printf("type %T is not supported.\n", T)
			// fmt.Printf 로 %T 를 이용해 타입출력도 가능하지만, reflect.TypeOf(val) 도 된다.
		}
	}
}

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println(Sum(100, 200))
	fmt.Println(Sum(10, 20, 40))

	var myInt64 int64 = 10
	fmt.Println()
	Print(true, 34, myInt64)
}
