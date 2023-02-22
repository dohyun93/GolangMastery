package main

import "fmt"

func main() {
	// 메서드를 선언하려면 '리시버'를 func 키워드와 함수 이름 사이에 중괄호로 명시해야 한다.
	// 리시버가 없으면, 함수다.
	a := &account{balance: 100}
	a.withDrawMethod(30)
	fmt.Println(a.balance)

	// 메서드는 함수 앞에 매개변수 하나를 두었을 뿐, 도대체 무슨 이점이 있어서 사용하는 걸까?
	// '소속'이다.
	// 리스너 타입에 함수가 속한다는 개념이 된다.
	// 리스너는 데이터를, 메서드는 기능을 담당하게 된다. 따라서, 응집도는 높아지게 된다.

	// 한 곳을 수정하면 다수 부분의 수정이 일어나야 하는 것을 '산탄총 수술문제'라고 하는데, 이는 응집도가 높을수록 문제의 심각성을 낮출 수 있다.
	// 그래서, 메서드를 사용하는 이유는, 데이터와 기능을 묶어 응집도를 높이기 위함이다.

	// 절차지향 -> 객체지향.
	// 객체 간 상호작용에 대한 코딩을 통해 프로그램을 개발한다는 사상이다.

	// Go 는 상속과 클래스가 없고, 메서드와 인터페이스만 지원한다.
	// 객체 간의 상호관계 중심으로 프로그래밍 할 수 있기 때문에 Go를 객체지향 언어라고 볼 수 있다.

	////////////////////////////////////////////////////////////////
	// 포인터 메서드 vs 값 타입 메서드
	// 리시버를 값 타입과 포인터로 정의할 수 있다.
	var mainA *account = &account{100, "dohyun", "kwon"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance) // 포인터 메서드 호출. mainA 와 a1 은 같은 인스턴스를 가리킨다.

	mainA.withdrawValue(30)    // 값 타입 메서드 호출.
	fmt.Println(mainA.balance) // -> 30빼진 값이 아님.

	var mainB account = mainA.withdrawReturnValue(20) // mainA 의 값이 a3 로 모두 복사된다.
	fmt.Println(mainB.balance)                        // 70인 객체 값을 받아서 20을 뺀 뒤, 그 값을 리턴하므로 50.

	mainB.withdrawPointer(30) // 50인 객체 포인터 메서드 호출하여 30 뺀다.
	fmt.Println(mainB.balance)

	// 포인터 메서드를 호출하면, 포인터가 가리키고 있는 메모리 주소값이 복사된다.
	// 반면, 값 타입 메서드를 호출하면 리시버 타입의 모든 값이 복사된다. (리시버 타입이 구조체면 구조체의 모든 데이터가 복사된다.)

}

type account struct {
	balance   int
	firstName string
	lastName  string
}

// 1. 포인터 메서드
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 2. 값 타입 메서드
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// 3. 변경된 값을 반환하는 '값 타입 메서드'
func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

type Rabbit struct {
	Width  int
	Height int
}

func (r Rabbit) info() int {
	return r.Width * r.Height
}

func (a *account) withDrawMethod(amount int) {
	// 함수가 아닌 메서드로 표현해보자.
	a.balance -= amount
	// 1. withDrawMethod 메서드는 *account 타입에 속한다.
	// 2. *account 는 리시버 타입이다.
}
