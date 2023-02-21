package main

import "fmt"

func main() {
	// 메서드를 선언하려면 '리시버'를 func 키워드와 함수 이름 사이에 중괄호로 명시해야 한다.
	// 리시버가 없으면, 함수다.
	a := &account{balance: 100}
	a.withDrawMethod(30)
	fmt.Println(a.balance)
}

type Rabbit struct {
	Width  int
	Height int
}

func (r Rabbit) info() int {
	return r.Width * r.Height
}

type account struct {
	balance int
}

func (a *account) withDrawMethod(amount int) {
	// 함수가 아닌 메서드로 표현해보자.
	a.balance -= amount
	// 1. withDrawMethod 메서드는 *account 타입에 속한다.
	// 2. *account 는 리시버 타입이다.
}
