package main

// 숫자를 출력한다
func printNumbers1() {
	//time.Sleep(1 *time.Millisecond)
	for i := 0; i < 10; i++ {
		//fmt.Printf("%d ", i)
	}
}

// 문자열을 출력한다.
func printLetters1() {
	//time.Sleep(1 *time.Millisecond)
	// A(0), B(1), C(2) .... J(10)
	for i := 'A'; i < 'A'+10; i++ {
		//fmt.Printf("%c ", i)
	}
}

func print1() {
	printNumbers2()
	printLetters2()
}

//
func goPrint1() {
	go printNumbers2()
	go printLetters2()
}

// 숫자를 출력한다
func printNumbers2() {
	//time.Sleep(1 *time.Millisecond)
	for i := 0; i < 100; i++ {
		//fmt.Printf("%d ", i)
	}
}

// 문자열을 출력한다.
func printLetters2() {
	//time.Sleep(1 *time.Millisecond)
	// A(0), B(1), C(2) .... J(10)
	for i := 'A'; i < 'A'+100; i++ {
		//fmt.Printf("%c ", i)
	}
}

func print2() {
	printNumbers2()
	printLetters2()
}

//
func goPrint2() {
	go printNumbers2()
	go printLetters2()
}
func main() {
	// 함수를 동기적으로 실행
	print2()

	// 함수를 비동기적으로 실행
	goPrint2()

	// 3초 대기
	//time.Sleep(time.Second * 3)
}
