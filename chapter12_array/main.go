package main

import "fmt"

func main() {
	// 1. 배열은 상수 길이의 선언을 할 수 있으며 초기화도 가능하다.
	var arrayname [5]int64 = [5]int64{1, 2, 3, 4, 5}
	for idx := 0; idx < len(arrayname); idx++ {
		fmt.Println(idx, "-th element: ", arrayname[idx])
	}

	// 2. ... 로 선언/초기화 하기.
	array2 := [...]int64{1, 2, 3, 4, 5} // 길이가 5인 int64 형 배열 생성.
	array2[2] = 700
	for idx := 0; idx < len(array2); idx++ {
		fmt.Println(idx, "-th element: ", array2[idx])
	}

	// 3. range 로 배열 순회하기
	var t [10]int = [10]int{1, 2, 3, 4, 5}
	for idx, val := range t {
		fmt.Println("t[", idx, "] is : ", val)
	}

	// 4. 배열 복사
	from := [5]int{1, 2, 3, 4, 5}
	to := [5]int{500, 400, 300, 200, 100}

	from = to // to 배열의 값을 from 메모리 공간에 복사한다. 타입이 같아야 에러가 발생하지 않는다.

	for idx, val := range from {
		fmt.Println("from[", idx, "] : ", val)
	}

	// 5. 다중 배열
	Map := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}, // 끝도 , 를 해줘야한다.
	}

	for r_idx, arr := range Map {
		for c_idx, val := range arr {
			fmt.Println("Map[", r_idx, "][", c_idx, "] : ", Map[r_idx][c_idx])
			fmt.Println("Map[", r_idx, "][", c_idx, "] : ", val)
		}
	}
}
