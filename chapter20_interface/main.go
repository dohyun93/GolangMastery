package main

import "fmt"

func main() {
	// 인터페이스는 '상호작용'할 수 있는 접점이다.
	// 인터페이스를 이용하면 추상화된 객체로 상호작용할 수 있다.

	// 2. 인터페이스 사용 예시
	// 2-1. Student 구조체를 만들어 인스턴스를 생성했다. String() 은 매개변수가 없고 string 반환을 하는 Student 타입에 속하는 메서드이다.
	student := &Student{"도현", 31}
	// 2-2. Stringer 인터페이스 변수를 만들었다. 이 인터페이스는 String() 이라는 매개변수가 없고 string 반환을 하는 메서드를 갖는다.
	var stringer Stringer
	// 2-3. String() string 메서드를 포함하는 인터페이스에 String() string 메서드를 포함하고 있는 인스턴스를 대입할 수 있다.
	stringer = student
	//fmt.Println(student)

	fmt.Printf("%s\n", stringer.String())

	// [핵심] - 인터페이스의 메서드를 포함하는 타입은 인터페이스에 할당될 수 있다.
	/// 그럼, 인터페이스는 왜  쓸까 ????

	// 구체화된 객체가 아닌 인터페이스만 가지고 메서드를 호출할 수  있기 때문이다. -> 큰 코드 수정이 필요 없다.

	//
}

/////////////// 1. 인터페이스 선언은 아래처럼 할 수 있다.

// 인터페이스도 구조체처럼 '타입'중 하나이기 때문에, type을 써줘야 한다.
// 즉, 인터페이스도 변수 선언이 가능하고, 변수의 값으로 사용할 수 있다.

// 다만, 아래 인터페이스에서 알 수 있듯이, 메서드 집합을 써줄 때 세 가지 유의사항이 있다.
// 1. 메서드는 반드시 메서드 명이 있다.
// 2. 매개변수와 반환이 다르더라도 '이름'이 같은 메서드가 존재할 수 없다.
// 3. 인터페이스에서는 메서드 구현을 포함하지 않는다.
type DuckInterface interface {
	// 인터페이스 밑에 메서드들을 넣을 수 있다.
	Fly()
	Walk(distance int) int
	// _(x int) -> 메서드는 반드시 이름이 있어야 한다.
}

//////////// 2. 인터페이스 사용 예시
type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕, 나는 %d 살 %s 라고 해 \n", s.Age, s.Name)
}
