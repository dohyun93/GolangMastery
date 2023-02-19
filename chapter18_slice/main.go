package main

import (
	"fmt"
	"sort"
)

func main() {
	// 이번 시간에는 go 언어에서 제공하는 '동적 배열'인 슬라이스를 정복해본다.
	// var myArray [10]int 는 10개의 int 만 담을 수 있는 배열이다. 만약 이보다 더 많은 요소를 저장하고자 한다면,
	// 새로운 메모리를 할당받아서, 요소 하나씩 복사를 해야 한다.

	// 슬라이스를 쓰면 이런 불편함이 없다.

	// 1. 슬라이스는 초기화 하지 않고 요소접근을 하면 패닉이 발생한다.
	var mySlice1 []int

	if len(mySlice1) == 0 {
		fmt.Println("mySlice1 은 길이가 0인 슬라이스 입니다.")
	}
	// mySlice1[1] = 100 //**  panic  **
	// panic: runtime error: index out of range [1] with length 0

	// goroutine 1 [running]:
	// main.main()
	// 	    /Users/dohyun/Desktop/Go/src/GolangMastery/chapter18_slice/main.go:18 +0x85
	// exit status 2
	fmt.Println(mySlice1)

	// //////////////////////// 슬라이스 초기화 //////////////////////////////////
	// 2. {}를 이용해 초기화하기.
	var slice2 []int = []int{1, 2, 3}
	fmt.Println(slice2)

	var slice3 []int = []int{1, 5: 2, 10: 3} // x: y 는 [x]=y라는 것이고, 그 앞은 0으로 채운다는 뜻이다.
	fmt.Println(slice3)                      // [1 0 0 0 0 2 0 0 0 0 3]

	// 참고로, 아래 두 선언은 배열과 슬라이스의 것임을 기억하자.
	var array = [...]int{1, 2, 3} // 배열
	var slice = []int{1, 2, 3}    // 슬라이스
	fmt.Println(array, slice)

	// 3. make() 를 이용한 초기화
	var slice4 = make([]int, 3) // 길이 3인 int 기본값 0으로 만들어진 슬라이스.
	fmt.Println(slice4)

	// //////////////////// slice 순회 : range //////////////////////////////
	var mySlice2 = make([]int, 10)
	for _, val := range mySlice2 {
		fmt.Println(val)
	}

	/////////////////////// slice 에 요소 추가하기 - append() /////////////////
	// 길이가 한 번 정해지면 늘릴 수 없는 배열과 다르게, 슬라이스는 요소를 추가할 수 있다.
	// append(slice, element) 를 해주면, 슬라이스 맨 뒤에 요소를 추가해서 만든 '새로운 슬라이스'를 결과로 반환한다.
	mySlice3 := []int{1, 2, 3}
	mySlice3Added := append(mySlice3, 100)
	fmt.Println(mySlice3, mySlice3Added)

	// 값을 여러개 추가할 수 있다.
	mySlice3Added2 := append(mySlice3, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(mySlice3, mySlice3Added2)

	// 슬라이스는 이처럼 배열과 []로 요소접근을 하고, range 로 순회할 수 있다는 비슷한 점이 있지만,
	// 아래와 같이 중요한 차이점이 있다.

	/////////////////// 슬라이스 동작 원리 ////////////////////
	// type SliceHeader struct{
	// 		Data uintptr --> 실제 배열을 가리키는 포인터. 슬라이스가 실제 배열을 가리키는 포인터를 가지고 있어서, 쉽게 크기가 다른 배열을 가리키도록 변경가능. 슬라이스 변수 대입 시 배열에 비해 사용되는 메모리나 속도에 이점.
	// 		Len int --> 배열의 요소수
	// 		Cap int --> 배열의 길이
	// }

	// make() 를 통해 슬라이스를 만들 경우, 아래와 같은 차이가 있다.
	mySlice4 := make([]int, 3)    // Len, cap = 3. 각 요소들은 모두 기본값 0으로 초기화됨.
	mySlice5 := make([]int, 3, 5) // Len=3, cap=5. 배열 길이는 5, 요소 개수는 3이다.
	fmt.Println(mySlice4, mySlice5)

	////////////////////// 슬라이스와 배열의 동작 차이 ////////////////
	mySlice6 := []int{1, 2, 3, 4, 5}
	myArray6 := [5]int{1, 2, 3, 4, 5}
	changeArray(myArray6)
	changeSlice(mySlice6)
	fmt.Println(mySlice6) // [1 2 3 4 5] -> [1 2 100 4 5]
	fmt.Println(myArray6) // [1 2 3 4 5] -> [1 2 3 4 5]

	// 왜 슬라이스는 바뀌었고, 배열은 바뀌지 않았을까?
	// Go 언어는 모든 값의 대입은 '복사'로 일어난다.
	// 함수에 인수로 전달하거나, 다른 변수에 대입할 때 모두 값을 '복사'한다.

	// 포인터도 마찬가지다. 포인터는 포인터의 값인 '메모리 주소'가 복사된다.

	// myArray 를 함수에 넘기면 값의 복사가 일어나서 원래 배열과 별도의 배열이 생성되며 값이 복사되고, 변경되는 것은 복사된 배열이므로, 기존 배열의 값에 영향을 미치지 않는다.

	// 반면, 슬라이스의 구조를 떠올려보자. 3개 필드로 구성된 구조체인 슬라이스는 함수호출 시 동일한 Data, len, cap 을 갖는 객체가 생성되고,
	// 복사된 슬라이스 구조체의 Data 는 기존 슬라이스의 Data 가 가리키는 배열과 동일한 것을 가리킨다. 따라서, 함수 내에서 요소를 변경하는 경우 원본 슬라이스가 가리키는 배열의 값을 변경하는것과 동일하기 때문에, 슬라이스만 값이 변경되었다.₩₩₩₩₩₩₩₩₩₩₩₩₩₩₩₩

	// /////////////////// append() 를 사용했을 때 예상치 못한 문제 1 //////////////////////////
	// -> 같은 배열을 가리키더라도, slice 마다 len이 다를 수 있고, 같은 배열을 가리키더라도 데이터 접근이 불가할 수 있다.
	mySlice7 := make([]int, 3, 5)      // len : 3 , cap : 5
	mySlice8 := append(mySlice7, 4, 5) // mySlice7 : len,cap : 3,5 / mySlice8 : len,cap : 5,5

	fmt.Println(len(mySlice7), cap(mySlice7), mySlice7) // 3 5 [0 0 0]
	fmt.Println(len(mySlice8), cap(mySlice8), mySlice8) // 5 5 [0 0 0 4 5]
	// 같은 배열을 가지고 있지만, 슬라이스 구조체가 갖는 각 필드가 상이해서 가리키는 배열은 같지만, 데이터 접근에 차이가 발생한다.
	// fmt.Println(mySlice7[3]) len:3인 슬라이스에 네 번째 요소를 접근하려고하면 에러발생한다.

	mySlice7 = append(mySlice7, 100)
	fmt.Println(len(mySlice7), cap(mySlice7), mySlice7) // 4 5 [0 0 0 100]
	fmt.Println(len(mySlice8), cap(mySlice8), mySlice8) // 5 5 [0 0 0 100 5]
	// mySlice7의 뒤에 100을 추가하면 같은 배열을 가리키고 있던 slice8의 값도 바뀐다.
	// mySlice7의 len 은 1증가한다.

	// /////////////////// append() 를 사용했을 때 예상치 못한 문제 2 ///////////////////////
	// cap을 넘어서는 append 를 했을 때 기존 배열이 아닌 새로운 배열을 가리키게 된다.
	slice9 := []int{1, 2, 3}
	slice10 := append(slice9, 4, 5)
	fmt.Println(len(slice9), cap(slice9), slice9)    // 3 3 [1 2 3]
	fmt.Println(len(slice10), cap(slice10), slice10) // 5 6 [1 2 3 4 5]

	// slice9, slice10 가 서로 다른 배열을 가리키는지 [1]을 바꿔보자.
	slice9[1] = 999
	fmt.Println(len(slice9), cap(slice9), slice9)    // 4 5 [1 999 3]
	fmt.Println(len(slice10), cap(slice10), slice10) // 5 5 [1 2 3 4 5]
	// slice9만 값이 바뀐것을 보아 서로 다른 배열을 바라보고 있는 것을 확인했다.

	//////////////////////////////////// 슬라이싱 /////////////////////////
	// slice를 배열의 슬라이싱 결과로 만든다면 cap 은 어떻게 될까?
	array10 := [5]int{1, 2, 3, 4, 5}
	slice11 := array10[1:2]
	fmt.Println(array10, len(array10))               // [1 2 3 4 5] 5
	fmt.Println(slice11, len(slice11), cap(slice11)) // cap 은 [1]부터 끝까지 배열의 크기로 정해지므로, [1]~[4] 총 4가 cap이 된다. [2] 1 4
	// Data 포인터는 array10[1] 을 가리킨다.
	// 따라서, array10[1] = 999 로 바꾸면, slice11 도 바뀐다.
	array10[1] = 999
	fmt.Println(array10, len(array10))               // [1 999 3 4 5] 5
	fmt.Println(slice11, len(slice11), cap(slice11)) // [999] 1 4

	// 슬라이싱은 배열뿐 아니라 슬라이스도 할 수 있다.
	slice12 := []int{1, 2, 3, 4, 5}
	slice13 := slice12[1:2]
	fmt.Println(slice12, len(slice12), cap(slice12)) // [1 2 3 4 5] 5 5
	fmt.Println(slice13, len(slice13), cap(slice13)) // [2] 1 4
	// slice13 은 slice12[1] 을 가리키는 포인터를 구조체 필드로 갖는다.

	//////////////////// cap 크기 조절하기 ////////////////////////
	// 인덱스 3개로 슬라이싱해서 cap 크기 조절할 수 있다. !!
	slice14 := []int{1, 2, 3, 4, 5}
	slice15 := slice14[1:2:3]                        // [2], len : 1, cap : 최대인덱스(3)-시작인덱스(1) = 2.
	fmt.Println(slice14, len(slice14), cap(slice14)) // [1 2 3 4 5] 5 5
	fmt.Println(slice15, len(slice15), cap(slice15)) // [2] 1 2

	//////////////////// slice 유용한 기능 ////////////////////////
	// 1. 슬라이스 복제
	slice16 := []int{1, 2, 3, 4, 5}
	slice17 := append([]int{}, slice16...) // 배열이나 슬라이스 뒤에 ...를 넣어주면 '모든 요소값'을 넣어준 것과 동일하다.
	// 즉, 이는 slice17 := append([]int{}, slice16[0], slice16[1], ..., slice16[4]) 와 같다.

	// append 말고, copy() 를 사용해도 슬라이스 복제가 가능하다.
	// func copy (dst, src []Type) int -> 반환값은 실제로 복사된 요소개수다.
	slice17 = []int{1, 2, 3, 4, 5}
	slice18 := make([]int, 3, 10) // len: 3, cap : 10
	slice19 := make([]int, 10)    // len: 10, cap: 10

	copied_cnt1 := copy(slice18, slice17)
	copied_cnt2 := copy(slice19, slice17)

	fmt.Println(copied_cnt1, slice18, len(slice18), cap(slice18)) // 3 [1 2 3] 3 10 -> 5개 담기지 않고, len 만큼 담긴다. (cap은 변하지 않는다.)
	fmt.Println(copied_cnt2, slice19, len(slice19), cap(slice19)) // 5 [1 2 3 4 5 0 0 0 0 0] 10 10

	// 따라서, copy 를 사용해서 값을 모두 다 복사하려면 아래처럼 할 수 있다.
	slice20 := make([]int, len(slice17))             // (1)복사할 대상의 길이만큼 만들고,
	copy(slice20, slice17)                           // (2) 복사한다.
	fmt.Println(slice20, len(slice20), cap(slice20)) // [1 2 3 4 5] 5 5

	/////////////////////////////////////////////////////////
	// 2. 슬라이스 요소 삭제
	slice21 := []int{1, 2, 3, 4, 5}
	delIdx := 2
	slice22 := append(slice21[:delIdx], slice21[delIdx+1:]...)
	fmt.Println(slice22, len(slice22), cap(slice22)) // [1 2 4 5] 4 5

	// 3. 슬라이스 요소 추가
	slice23 := []int{1, 2, 3, 4, 5, 6, 7}
	slice23 = append(slice23, 0)
	idxInsert := 2
	copy(slice23[idxInsert+1:], slice23[idxInsert:]) // slice23[idxInsert+1] 기준으로 넣기 때문에 5회 넣어진다.
	slice23[idxInsert] = 100

	fmt.Println(slice23, len(slice23), cap(slice23)) // [1 2 100 3 4 5 6 7] 8 14

	///////////////////////////////////////////////////////////
	// 4. slice 정렬 - int 슬라이스
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s) // Int sort 의 약자인듯. Float64s 같은 함수도 있음.
	fmt.Println(s)

	// slice 정렬 - 구조체 슬라이스
	// Len(), Less(), swap() 메서드가 필요하다.
	students := Students{
		{"도현", 31},
		{"가령", 31},
		{"피그", 29},
	}
	sort.Sort(students)
	// Less() 메서드로 각 요소의 Age 값을 비교한다.
	// 아래 메서드들의 리스너인 Students 타입은 Len, Less, Swap 메서드를 가지고 있지 않기 때문에,
	// sort.Sort 의 인수로 사용될 수 없다. 따라서 아래 메서드들을 '구현'해서 커스텀 구조체 타입도 sort.Sort 를 통해 정렬될 수 있도록 개발한 것이다.
	fmt.Println(students)
}

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int           { return len(s) }              // Len() : Student 슬라이스 길이 반환 메서드.
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age } // Less() : Student 각 객체의 나이 비교.
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func changeArray(myArray [5]int) {
	myArray[2] = 100
}

func changeSlice(mySlice []int) {
	mySlice[2] = 100
}
