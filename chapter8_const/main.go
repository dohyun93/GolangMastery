package main

import (
	"fmt"
)

const (
	firstOrder = iota
	secondOrder
	thirdOrder
)

const (
	shiftone = 1 << iota // 1 << 0 = 1
	shifttwo = 2 << iota // 2 << 1 = 4
)

func main() {
	// 상수는 변수와 달리 var 이 아니라 const 라는 키워드로 선언할 수 있다.
	const myconst int = 31
	fmt.Println(myconst)

	// 변수는 값/이름/주소/타입을 갖지만, 상수는 값/이름/타입만 갖고 주소는 갖지 않는다.
	fmt.Println(firstOrder, secondOrder, thirdOrder)

	// iota 는 scope 내에서 유효하며, 0부터 1씩 증가한다.
	fmt.Println(shiftone, shifttwo)

	// 상수는 타입없이 선언할 수 있는데, 변수에 할당할 때 지정되므로 다양한 타입에 사용될 상수값에 타입없이 선언하여 사용할 수 있다.

	// @ 상수와 리터럴 @
	// 아래에서 30, 20, "my string" 은 리터럴이라고 하며, 고정된 값, 값 자체로 쓰인 문구라고 볼 수 있다.
	// var a int = 30
	// var b string = "my string"
	// a = 20

	// 상수는 리터럴과 같이 취급되어서, 컴파일 할 때 상수는 리터럴로 변환된다.
	// 그래서 아래 코드는 컴파일 타임때 상수 표현식은 결과값 리터럴로 변환되기 때문에 상수 표현식 계산에 cpu자원을 사용하지않는다.
	// [작성코드]
	// const PI = 3.14
	// var aa int = PI * 100

	// [컴파일 타임]
	// var aa int = 314

	// 위에서 살펴본 상수의 메모리 주소값에 접근할 수 없는 이유 역시,
	// 컴파일 타임에 상수를 리터럴로 변환하여 실행 파일에 값 형태로 쓰이기 때문임. 그래서 동적 할당 메모리 영역을 사용하지 않는다.

	// 프로그램을 로드할 때 실행파일이 올라가는 영역을 '코드 영역' 이라고 하고, 프로그램 실행을 위해 실행 중 메모리를 할당해서 사용하는 영역을
	// '동적 할당 메모리 영역'이라고 한다. 상수는 컴파일 시, 리터럴로 '코드 영역'에 포함되기 때문에 이 동적 할당 메모리 영역을 사용하지 않는다.
}
