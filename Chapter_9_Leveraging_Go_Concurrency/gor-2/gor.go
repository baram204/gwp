package main

import "fmt"

// 숫자를 출력한다
func printNumbers1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

// 문자열을 출력한다.
func printLetters1() {
	// A(0), B(1), C(2) .... J(10)
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c ", i)
	}
}

func print1() {
	printNumbers1()
	printLetters1()
}

//
func goPrint1() {
	go printNumbers1()
	go printLetters1()
}

func main() {
	// 함수를 동기적으로 실행
	print1()

	// 함수를 비동기적으로 실행
	goPrint1()

	// 3초 대기
	//time.Sleep(time.Second * 3)
}
