package main

import "fmt"

type Data struct {
	name  string
	value int64
}

func main() {
	// 포인터 : '메모리 주소'를 값으로 갖는 타입이다.
	var myInt int64 = 10
	ptmyInt := &myInt

	fmt.Println(myInt, ptmyInt)

	// 포인터 변수의 초기화.
	// '인스턴스'란 메모리에 할당된 데이터의 실체다.
	var myInstance Data // myInstance 는 인스턴스다.
	// 또는 아래처럼 인스턴스 선언 및 초기화 하는 방법이 있다.
	myInstance2 := &Data{}
	myInstance3 := new(Data)

	fmt.Println(myInstance, myInstance2, myInstance3)

	// 이런 인스턴스들은 메모리에 올라가있는 데이터의 실체라고 했다.
	// 그럼, 이런 인스턴스들은 언제 메모리에서 해제되는 것일까?

	// Go 에는 가비지 컬렉터라는 것이 존재한다.
	// 이것이 일정 간격으로 메모리에서 쓸모없어진 데이터들을 청소한다.

	// 그럼, 쓸모없어진 데이터라는 것은 어떤 것일까 ? -> 아무도 찾지 않는 데이터이다.

	// 예를 들어서, 아래 함수가 호출된 다음 종료되었을 때, 포인터 변수 p 는 더이상 찾지 않는 데이터가 된다.
	// 따라서 해당 인스턴스의 메모리는 회수되게 된다.
	// func Testfunc(){
	// 	p := &Data{} --> 포인터변수 p 는 Data 데이터의 주소가 할당된 포인터 변수이고, 함수가 끝난 뒤 메모리에서 회수된다.
	// 	p.name = "dohyun"
	// 	fmt.Println(p.name)
	// }

	// 그러나, 이런 가비지 컬렉터는 사용하지 않는 메모리를 회수해주는 고마운 역할을 하지만,
	// 언제나 좋은 면만 있는 것은 아니다. 왜냐하면, 사용하지 않는 메모리들을 찾기 위해서 모든 메모리 영역을 검사해서
	// 쓸모없는 데이터를 지워주는 데 성능을 많이 쓰기 때문이다.

	// 따라서, 가비지 컬렉터를 사용하면 메모리 관리 측면에서 이점이 있을 수 있으나, 성능 면에서는 손해가 발생하는 것이다.

	// 아래 네 가지를 정리하자.
	// 1. 인스턴스는 메모리에 생성된 데이터의 실체이다. (e.g., &Data{})
	// 2. 포인터를 이용해서 인스턴스를 가리키게 할 수 있다. (e.g., p := &Data{} 에서 p 는 Data 타입의 주소를 가리키는 포인터 변수.)
	// 3. 함수 호출 시 포인터 인수를 통해서 인스턴스를 입력받고 그 값을 변경할 수 있게 된다.
	// 4. 쓸모 없어진 인스턴스는 가비지 컬렉터가 자동으로 지워준다.

	// <추가> 스택/힙 메모리
	// 스택 메모리 : 함수 내에서만 사용가능한 메모리 영역.
	// 힙 메모리 : 함수 외부로 공개되는 메모리 공간은 여기에서 할당한다.
	// Go 에서는 '탈출 검사'를 해서 어느 메모리에 할당할지를 결정한다.

	// 함수 외부로 공개되는 인스턴스의 경우, 함수가 종료되어도 사라지지 않는다.
	// 탈출 검사를 통해 u 변수의 인스턴스가 함수 외부로 공개되는 것을 분석해내서, u를 스택 메모리가 아닌 힙 메모리에 할당되게 한다.
	// 즉, Go 는 어떤 타입이나 메모리 할당에 의해서 스택 메모리를 사용할 지, 힙 메모리를 사용할 지 결정하지 않는다.
	// '메모리 공간이 함수 외부로 공개되는지 여부' 를 자동으로 검사해서 (=탈출 검사) 스택/힙 메모리 중 어디에 할당할지를 결정한다.

	// 마지막으로, Go 는 c, c++ 언어와 달리, 스택 메모리가 계속 증가되는 동적 메모리 풀이다. 일정 크기를 갖는 c, c++ 와 달라서
	// 메모리 효율성이 높고, 재귀 호출 때문에 스택 메모리가 고갈되는 문제도 발생하지 않는다.
	userPointer := NewUser("dohyun", 31)
	fmt.Println(userPointer.Name, userPointer.Age)
}

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u // 탈출 분석으로 u는 함수 내부에 있지만 u의 메모리가 사라지지 않는다.
}
