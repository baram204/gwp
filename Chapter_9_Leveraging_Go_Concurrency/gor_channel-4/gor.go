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
	// 출력을 위해 대기(락) 걸린 w1 채널에 값을 넣어주면 블락이 풀리며 프로그램이 종료된다.
	w<- true
}

// 문자열을 출력한다.
func printLetters1(w chan bool) {
	time.Sleep(1 *time.Millisecond)
	// A(0), B(1), C(2) .... J(10)
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c ", i)
	}

	// 메인 스레드 쪽에서 블락을 걸어주지 않으면 개별 고루틴에서 전달된 채널에 무슨짓을 해도
	// 메인 스레드는 그냥 종료되어 버린다.
	w <- false
}

func main() {

	w1, w2 := make(chan bool), make(chan bool)
	go printNumbers1(w1)
	go printLetters1(w2)

	// w1 채널에서 무언가를 꺼내기 위해서, 넣어지는 것을 대기(블록)이 걸린다.
	// printNumber1 내부에서 w1 에 아무것도안 넣어주면 채널은 기다리고 있는데
	// 모든 고루틴은 다 끝나서 잠이 든 상태이므로 데드락이 난다.
	<-w1
	<-w2

}
