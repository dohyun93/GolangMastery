package main

import (
	"fmt"
)

func main() {
	// if 초기문; 조건문 형태 많이 사용한다.
	if age := 31; age > 30 {
		fmt.Println("30대 네요.")
	} else {
		fmt.Println("어리네요.")
	}
}
