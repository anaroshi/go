package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	stdin = bufio.NewReader(os.Stdin)
	money = 1000
)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5)
	cnt := 1

	for {
		fmt.Printf("1~5사이의 숫자를 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("1~5사이의 숫자만 입력하세요!")
		}

		if n < 1 || n > 5 {
			fmt.Println("1~5사이의 숫자만 입력하세요:")
		} else if n == r {
			fmt.Println("입력하신 숫자가 다릅니다. 아쉽지만 가진 돈에 100원이 차감됩니다. ")
			money -= 100
			fmt.Printf("현재까지 적립된 금액은 %d입니다.\n", money)
			fmt.Println("시도한 횟수 : ", cnt)
		} else {
			fmt.Println("숫자를 맞췄습니다. 축하합니다. 상금으로 500원을 적립해드립니다.")
			money += 500
			fmt.Printf("현재까지 적립된 금액은 %d입니다.\n", money)
			fmt.Println("시도한 횟수 : ", cnt)
		}

		if money <= 0 {
			fmt.Println("가진돈이 부족합니다.")
			break
		} else if money > 5000 {
			fmt.Println("최대 상금을 취득하셨습니다. 축하드립니다.")
			break
		}

		cnt++

	}

}
