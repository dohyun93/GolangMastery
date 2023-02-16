package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// step 1: 0~99사이의 랜덤한 숫자를 정한다.
// step 2: 사용자의 입력을 받는다.
// step 3: 입력과 랜덤 값을 비교한다.
// 3-1. 만약 입력한 숫자가 더 크다면 "입력하신 숫자가 더 큽니다."
// 3-2. 입력한 숫자가 더 작으면 "입력하신 숫자가 더 작습니다."를 출력. 다시 사용자 입력을 받아서 반복.
// step 4: 만약 숫자가 맞으면 "축하합니다. 숫자를 맞추셨습니다. 시도횟수: XX번"

func FindNumberGame() {
	var targetNumber int = rand.Intn(100)
	fmt.Println("### solution ### target is ", targetNumber)
	var stdin = bufio.NewReader(os.Stdin)
	InputNumber, cnt := 0, 0

	for {
		_, err := fmt.Scanln(&InputNumber)
		if err != nil {
			// 입력 버퍼를 비워준다.
			stdin.ReadString('\n')
		} else {
			cnt++
			if InputNumber < targetNumber {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else if InputNumber > targetNumber {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else {
				fmt.Printf("축하합니다. 숫자를 맞추셨습니다. 시도횟수: %d번\n", cnt)
				break
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	FindNumberGame()
}
