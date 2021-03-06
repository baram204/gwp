package main

import (
	"fmt"
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
		select {
		// a 채널에서 메시지 꺼내면 비어버린다. 고루틴은 슬립
		case msg := <-a:
			fmt.Printf("%s from A\n", msg)
		// b 채널에서 메시지 꺼내면 또 비어버린다 고루틴은 슬립
		case msg := <-b:
			fmt.Printf("%s from B\n", msg)
		}
		// a 채널 b 채널 모두 꺼낼 것이 없는데 case 를 돌아도 채널을 꺼내려고 대기중이므로 데드락 걸림

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
