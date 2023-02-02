package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("fmt.Print() does not print \\n")
	fmt.Println()
	fmt.Println("fmt.Println() does print \\n")
	fmt.Printf("fmt.Printf() represents data for corresponding format.\n")

	a := "string value"
	flag := true
	number := 1000
	fmt.Printf("%v\n", a)    // %v : 데이터 타입에 맞게 출력한다.
	fmt.Printf("%T\n", a)    // %T : 데이터 타입을 출력한다.
	fmt.Printf("%t\n", flag) // %t : 불리언을 true/false 로 출력한다.

	fmt.Printf("%b\n", number) // %b : 2진수로 출력한다. %o는 8진수.
	fmt.Printf("%c\n", number) // %c : 유니코드 문자를 출력한다. (정수 타입만 가능.)

	// %x 16진수
	// %X 16진수 (대문자)
	// %e 지수 형태로 실숫값을 출력.
	// %g, %G: 값이 큰 실숫값은 지수 형태(%e)로 출력하고, 작은 실숫값은 %f 형태로 실수 그대로 출력한다.

	// 입력 - Scanln 함수는 입력 받을 때 공백 (엔터제외)을 기준으로 구별하여 데이터를 입력받는다.
	var ai, bi int
	n, err := fmt.Scanln(&ai, &bi)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, ai, bi)
	}

	// 키보드 입력과 Scan 함수의 동작 원리
	// 표준입력장치 (Pc에서는 키보드)로 입력하면 '표준 입력 스트림'이라는 메모리 공간에 임시 저장된다.
	// 이는 FIFO로, 하나씩 꺼내어 값을 확인하는데, 만약 정수가 올 것을 예상하였지만 다른 타입이 온 경우
	// 나머지 데이터들은 그대로 스트림에 남아있다.

	// 이 경우 예상치 못한 동작을 발생시킬 수 있어, 표준 입력 스트림을 비워주어야 한다.
	stdin := bufio.NewReader(os.Stdin) // 표준 입력을 읽는 객체

	var Int1, Int2 int64

	n, err = fmt.Scanln(&Int1, &Int2)
	if err != nil {
		fmt.Println(err)       // "expected integer"
		stdin.ReadString('\n') // <- 여기서 입력 스트림을 비워주는 기능을 한다.
	} else {
		fmt.Println(n, Int1, Int2)
	}

	n, err = fmt.Scanln(&Int1, &Int2)
	if err != nil {
		fmt.Println(err)
		//	stdin.ReadString('\n')
	} else {
		fmt.Println(n, Int1, Int2)
	}

}
