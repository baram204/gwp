package main

import "fmt"

//func printNumbers(w chan bool) {
func printNumbers() {
	for i := 0; i < 4; i++ {
		//time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	fmt.Print("숫자출력 포문")
	// 채널에 값이 넣어지기 전 까지 고루틴은 지들끼리 돈다.
	//w <- true
}

func printLetters() {
	//func printLetters(w chan bool) {
	for i := 'A'; i < 'A'+4; i++ {
		//time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	fmt.Print("문자출력 포문")

	//w <- true
}

func main() {
	//w1, w2 := make(chan bool), make(chan bool)
	go printNumbers()
	//go printNumbers(w1)
	go printLetters()
	//go printLetters(w2)
	//fmt.Printf("%t", <-w1)
	//fmt.Printf("%t", <-w2)

	// 각 채널은 상대방이 전송할 준비가 될 때까지 기다린다.
	//<-w1
	//<-w2
}
