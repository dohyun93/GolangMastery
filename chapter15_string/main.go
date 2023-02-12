package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// stdin.ReadString('\n') // 입력 스트림 버퍼 비우기.

	string1 := "hello\tworld\n"
	string2 := `GO is "awesome"!\nGo is simple and \t'powerful'` // back quote로 묶은 문자열 내부는 그대로 출력된다.

	fmt.Println(string1)
	fmt.Println(string2)

	// ===============================================================
	// 백 쿼트를 사용하면 여러 줄의 문자열이나 특수문자를 그대로 출력할 수 있다.
	string3 := `안녕하세요.
	저는 권도현 입니다.
	당신은 누구십니까?`

	fmt.Println(string3)

	// Go는 UTF-8 문자코드를 표준 문자코드로 사용한다.
	// UTF-8은 다국어 문자를 지원하고 문자열 크기를 절약할 목적으로 Go 창시자인 롭 파이크와 켄 톰슨이 고안한 문자코드다.
	// UTF-8은 자주 사용되는 영문자, 숫자, 일부 특수문자를 1바이트로 표현하고, 그 외 다른 문자들은 2~3바이트로 표현한다.
	// 이를 통해 UTF-8 은 자주 사용되는 문자/숫자들을 1바이트로 표현해서 UTF-16보다 크기를 절약할 수 있고, ANSI 코드와 1:1 대응이 되어,
	// 바로 ANSI 로 변환된다는 장점이 있다. Go 는 UTF-8 을 표준 문자코드로 사용하기 때문에, 한글이나 한자를 별다른 변환없이 사용할 수 있다.

	// ===============================================================
	// rune
	// rune 은 int32 의 alias로, 문자 하나를 표현하는데 사용한다.
	// UTF-8 은 1~3바이트 크기이기 때문에, 3바이트가 필요하나, Go 의 기본타입 중 3바이트가 없으므로, 4바이트인 int32 가 사용된다.

	var mychar rune = '령'
	fmt.Printf("%T\n", mychar) // mychar type
	fmt.Println(mychar)        // int32 출력
	fmt.Printf("%c\n", mychar) // 문자로 출력

	// ===============================================================
	// len() 으로 문자열 크기 (=메모리 크기)알아내기
	strKr := "김가령"
	strEn := "Abc"

	fmt.Println("size of strKr: ", len(strKr))
	fmt.Println("size of strEn: ", len(strEn))
	// UTF-8에서 한글은 3바이트를 차지하기 때문에, 3글자여서 9 바이트가 나왔고, 알파벳은 글자당 1바이트라 3바이트가 나온다.

	// ===============================================================
	// []rune 타입 변환으로 글자 수 알아내기
	// string, []rune 은 상호 타입 변환이 가능하다.

	myString := "Hello 세상"
	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	fmt.Println(myString)
	fmt.Println(string(runes))
	// rune으로 정수 슬라이스를 string 으로 변환시키면 그 정수에 매핑되는 글자로 변환이 된다.

	// 반대로 문자열을 rune 슬라이스로 바꿔보자.
	fmt.Println([]rune(myString))

	// []rune 과 문자열에 내장함수 len() 을 썼을 때 차이를 알아보자.
	fmt.Println("(slice) runes and it's len: ", runes, len(runes))           // len(슬라이스) : 길이반환
	fmt.Println("(string) myString and it's len: ", myString, len(myString)) // len(문자열) : 크기반환

	// ===============================================================
	// 문자열 순회.
	// 1. range
	// 2. index
	// 3. []rune 변환 후 한 글자씩 순회

	stringIterate := "Hello 월드!"
	// 1. index
	for i := 0; i < len(stringIterate); i++ { // len(문자열)이라 전체 길이보다 큰 값이 조건문에 들어가지만, 상관없이 마지막 요소까지만 순회한다.
		fmt.Printf("%d번째 요소 타입, 정수, 문자 값: %T, %d, %c\n", i, stringIterate[i], stringIterate[i], stringIterate[i])
	}
	// stringIterate[i] 처럼 인덱스로 접근하면 '문자열의 각 요소'의 타입은 uint8, 즉 byte로 읽기 때문에, 3바이트인 한글은 깨져 표시된다.
	// 다른 방법을 보자.

	// 2. []rune 으로 변환 후 순회
	stringIterateRuneSlice := []rune(stringIterate)
	for i := 0; i < len(stringIterateRuneSlice); i++ {
		fmt.Printf("%d번째 요소 타입, 정수, 문자 값: %T, %d, %c\n", i, stringIterateRuneSlice[i], stringIterateRuneSlice[i], stringIterateRuneSlice[i])
	}
	// int32, 즉, 4바이트 크기의 요소들로 이루어진 슬라이스의 각 요소를 꺼내서 문자를 출력한다.
	// 따라서, 3바이트인 한글도 정상적으로 출력된다.
	// 그러나, string -> []rune 으로 변환되는 과정에서 별도의 배열을 할당하므로, 불필요한 메모리를 사용하게 된다.

	// 3. range 키워드를 이용해 한 글자씩 순회하기.
	for idx, val := range stringIterate {
		fmt.Printf("%d번째 요소 타입, 정수, 문자 값: %T, %d, %c\n", idx, val, val, val)
	}
	// 인덱스로 접근하지 않고 그대로 문자열을 사용했지만, range 를 사용하니 int32 타입으로 값을 읽어온다.
	// 추가 메모리 할당 없이 한 글자씩 순회할 수 있다.

	// ===============================================================
	// 문자열 합치기
	str1 := "Hello "
	str2 := "world!"
	if str1+str2 == "Hello world!" {
		fmt.Println("They are same strings")
	}
	// strings.Join([]string, delimeter) 사용으로 문자열을 이어붙일 수 있다.
	strArray := []string{str1, str2}
	myString = strings.Join(strArray, "")

	// 그 외 문자열 동일비교 및 문자열 대소비교도 가능하다.
	// 동일비교 : ==, !=
	// 대소비교 : <, >, >=, <= 로 앞에서부터 한 글자씩 UTF-8의 값을 비교한다.

	valorigin := "Xi story"
	valcopy := valorigin

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&valorigin))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&valcopy))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
	// 결론은, 문자열을 복사한다고해서 문자열 변수가 가리키는 문자열 데이터가 복사되는 것이 아니라,
	// 16바이트 값만 복사될 뿐 문자열 데이터는 복사되지 않는다.

	// 아래 StringHeader 구조체가 곧 String 타입의 구조이며, string 복사는 이 구조체의 복사와 같으므로, 16바이트 값에 대한 복사만 이루어진다.
	// type StringHeader struct {
	// 	   Data uintptr // 8 byte
	// 	   Len  int     // 8 byte (64-bit pc)
	// }

	// ===============================================================
	// 문자열은 불변이다? No. 전체변경은 가능하지만 일부 변경은 불가하다.
	// 문자열을 이루는 문자 일부만 변경할 수 없다는 뜻이다.
	var str string = "Hello world"
	//str[2] = 'a' -> 일부만 변경할 수 없다.

	// []byte <-> string 을 통한 일부 변경가능?
	byteArray := []byte(str)

	strH1 := (*reflect.StringHeader)(unsafe.Pointer(&str))
	strH2 := (*reflect.StringHeader)(unsafe.Pointer(&byteArray))

	fmt.Printf("origin string : \t%x\n", strH1.Data)
	fmt.Printf("copied slice of byte : \t%x\n", strH2.Data)
	// 두 변수 str, byteArray 가 가리키는 주소가 다르다. 따라서, byteArray의 일부 데이터를 바꾸더라도 str 의 문자열값에 영향을주지 않는다.

	// ===============================================================
	// 문자열 합산
	StringAdd := "A"
	stringHead := (*reflect.StringHeader)(unsafe.Pointer(&StringAdd))
	fmt.Printf("0 Added string's address: %x\n", stringHead.Data)

	StringAdd += " B"
	fmt.Printf("0 Added string's address: %x\n", stringHead.Data)

	// Go 언어는 기존 문자열 공간을 건드리지 않고, 새로운 메모리 공간을 만들어서 두 문자열을 합치기 때문에, string 합 연산 이후 주소값이 변경된다.
	// 따라서 문자열 불변 원칙이 준수된다. 하지만 string 합 연산을 빈번하게 하면 메모리가 낭비되기 때문에, strings 패키지의 Builder를 이용해서 메모리 낭비를 줄일 수 있다.

	// 아래 AggregateRune 메소드를 통해 메모리 낭비를 줄일 수 있다.
	myStringBuilder := "abcde"
	fmt.Println("origin string: ", myStringBuilder)
	fmt.Println("aggregated string : ", AggregateRune(myStringBuilder))
}

func AggregateRune(myString string) string {
	var builder strings.Builder
	for _, val := range myString {
		if val >= 'a' && val <= 'z' {
			// string += "a" 같은 것 대신 아래와 같이 builder 를 활용한다.
			builder.WriteRune('A' + (val - 'a'))
		} else {
			builder.WriteRune(val)
		}
	}
	return builder.String()
}
