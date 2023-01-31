package main

import "fmt"

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
}
