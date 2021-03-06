package main

import (
	"fmt"
	"time"
)

// 숫자를 출력한다
func printNumbers1(w chan bool) {
	time.Sleep(1 *time.Millisecond)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

// 문자열을 출력한다.
func printLetters1(w chan bool) {
	time.Sleep(1 *time.Millisecond)
	// A(0), B(1), C(2) .... J(10)
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c ", i)
	}
}

func main() {

	w1, w2 := make(chan bool), make(chan bool)
	go printNumbers1(w1)
	go printLetters1(w2)

	// w1 채널에서 무언가를 꺼내기 위해서, 넣어지는 것을 대기(락)이 걸린다.
	// printNumber1 내부에서 w1 에 아무것도안 넣어주면 채널은 기다리고 있는데
	// 모든 고루틴은 다 끝나서 잠이 든 상태이므로 데드락이 난다.
	<-w1

}
