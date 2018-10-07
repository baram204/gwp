package main

import (
	"fmt"
	"time"
)

func callerA(c chan string) {
	c <- "Hello World!"
	//close(c)
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
	//close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)

	for i := 0; i < 5; i++ {
		/*
		앞서 봤듯이 고루틴이 실행될 때까지 좀 지연이 있기 때문에
		채널에서 메시지 꺼내기도 전에 루프 자체가 끝나버리면 아래처럼 첫번째에 기본값이 나타난다.

		both channel deadlock! Default
		Hello World! from A
		Hola Mundo! from B
		both channel deadlock! Default
		both channel deadlock! Default

		그것을 방지하기 고루틴이 충분히 실행되도록 시간을 지연해준다.
		*/

		time.Sleep(1 * time.Microsecond)
		select {
		case msg := <-a:
			fmt.Printf("%s from A\n", msg)
		case msg := <-b:
			fmt.Printf("%s from B\n", msg)
		// a, b 모든 채널이 수신이 안되는 차단 상태이며, 고루틴도 모두 슬립할 때
		// 계속 채널을 대기하면 데드락이 걸리니까. 대신
		// 호출될 기본 케이스를 넣어서 보여준다.
		default:
			fmt.Printf("both channel deadlock! Default\n")
		}

	}
}

//func main() {
//	a, b := make(chan string), make(chan string)
//	go callerA(a)
//	go callerB(b)
//	var msg string
//	openA, openB := true, true
//	for openA || openB {
//		select {
//		case msg, openA = <-a:
//			if openA {
//				fmt.Printf("%s from A\n", msg)
//			}
//		case msg, openB = <-b:
//			if openB {
//				fmt.Printf("%s from B\n", msg)
//			}
//		}
//	}
//}

//func main() {
//	a, b := make(chan string), make(chan string)
//	go callerA(a)
//	go callerB(b)
//	msg1, msg2 := "A", "B"
//	for {
//		time.Sleep(1 * time.Microsecond)
//
//		select {
//		case msg1 = <-a:
//			fmt.Printf("%s from A\n", msg1)
//		case msg2 = <-b:
//			fmt.Printf("%s from B\n", msg2)
//		// default:
//		// 	fmt.Println("Default")
//		}
//		if msg1 == "" && msg2 == "" {
//			break
//		}
//
//	}
//}
