package main

import (
	"fmt"
)

func main() {
	// switch 문
	// switch 비교값/대상 case 조건/값
	// 불리안/대상과 조건/값이 일치하는 경우에 실행하도록 하는 구문.

	// 1. 대상과 값 비교
	var myAge = 31
	switch myAge {
	case 10:
		fmt.Println("10 age.")
	case 15:
		fmt.Println("15 age.")
	case 31:
		fmt.Println("31 age. It's my age.")
	default:
		fmt.Println("I don't know age.")
	}

	// 2. 불리안과 조건
	// switch의 불리안을 만족시키는 케이스가 실행된다.
	switch true {
	case 10 < myAge && myAge < 30:
		fmt.Println("11~29 age.")
	case 30 <= myAge && myAge < 40:
		fmt.Println("30~39 age.")
	default:
		fmt.Println("I don't know the age variation.")
	}

	// 3. 초기문과 대상, 또는 초기문과 불리안
	// (3-1) 초기문과 대상
	switch yourAge := 40; yourAge {
	case 30:
		fmt.Println("your age is 30.")
	case 40:
		fmt.Println("your age is 40.")
	default:
		fmt.Println("I don't know your age.")
	}

	// (3-2) 초기문과 불리안
	switch yourAge := 50; true {
	case 30 >= yourAge:
		fmt.Println("Your age is equal or less than 30.")
	case 40 <= yourAge && yourAge <= 50:
		fmt.Println("Your age is between 40 and 50.")
	}

	// 4. break, fallthrough
	// 일반적으로 다른 언어에서는 케이스문 안에 break 를 넣어야 그 다음 케이스가 실행되지 않지만,
	// 고는 break 를 넣지않아도 케이스 실행 후 스위치문을 빠져나간다.
	// 그래서 break 를 사용하던 안하던 조건 또는 값만 맞으면 해당 케이스만 실행되기때문에 이를 고려해 개발하면 된다.

	// 다른 언어와 다르게 fallthrough 라는게 존재하는데, 말그대로 밑으로 내려간다는 의미다.
	// fallthrough 가 있다면 값/조건이 맞지 않아도 그 다음 케이스까지 실행한다.

	switch myValue := 31; myValue {
	case 10:
		fmt.Println("my value is 10.")
	case 31:
		fmt.Println("my value is 31.")
		fallthrough // 조건/값에 상관없이 다음 케이스를 실행한다.
	case 40:
		fmt.Println("my value is 40.")
		fallthrough
	case 50:
		fmt.Println("my value is 50.")
	default:
		fmt.Println("default.")
	}
}
