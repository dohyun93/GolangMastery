package main

import "fmt"

type Dohyun struct {
	name   string
	age    int64
	height int64
}

type Human struct {
	Dohyun   // 필드명 생략
	weight   int64
	footsize int64
	age      int64
}

func main() {
	dohyun := Dohyun{name: "dohyun", age: 31, height: 187}
	fmt.Println(dohyun)
	human := Human{Dohyun{name: "humanDefaultName", age: 31, height: 187}, 80, 300, 100}
	fmt.Println(human)

	fmt.Println("nested same field (Dohyun): ", human.Dohyun.age)
	fmt.Println("nested same field (Human): ", human.age)

	// 구조체 복사 - 값 복사.
	garyung := dohyun
	fmt.Println(garyung)

	// 구조체 변수를 배치할 때 순서는 크기가 작은 것부터 위에서부터 선언하는 것이 더 낫다. (메모리 패딩)
	// 메모리 정렬을 참고하자.
}
