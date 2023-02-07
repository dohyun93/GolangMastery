package main

import "fmt"

func main() {
	// 반복문에는 for 가 있다.
	// for 초기문; 조건문; 후처리
	// 로 표현할 수 있고, 아래 조합들로 구현될 수 있다.
	// for 초기문; 조건문;
	// for ; 조건문; 후처리
	// for ; 조건문 ;  -> for 조건문 과 동일.

	// 또한 무한루프도 만들 수 있다.
	// for
	// for true

	// 이중 포문을 플래그 변수없이 손쉽게 빠져나갈 수 있는 방법은 '레이블'을 이용하는 것이다.
	var a, b int
OUTERFOR:
	for a = 1; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OUTERFOR
			}
			fmt.Println("### Not found solution. ###", a, b)
		}
	}
	fmt.Println("Found solution. a and b are: ", a, b)

	// 그러나 레이블이 편할 수는 있지만, 혼동을 불러일으킬 수 있고 자칫 잘못하면 버그를 발생시킬 수 있으니,
	// 클린 코드를 지향하려면 중첩된 내부 로직을 '함수'로 묶어서 중첩을 줄이고, 플래그 변수나 레이블 사용을 최소화 해야한다.

	// 위 소스를 리팩토링하면 아래와 같다.
	var b_val int
	var found bool
	for val := 1; val <= 9; val++ {
		if b_val, found = find45(val); found { // 중첩되지 않도록 for문을 함수에서 처리하였다.
			break
		}
	}
	if b_val != 0 {
		fmt.Println("Answer (b): ", b_val)
	} else {
		fmt.Println("could not found the answer.")
	}

}

func find45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}
