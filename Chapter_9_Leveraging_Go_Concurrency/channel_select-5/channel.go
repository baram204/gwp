package main

import "fmt"

func callerA(c chan string) {
	c <- "Hello World!"
	// 각 채널을 닫아(수신기에 입력이 없다)버리면 수신하기 위해 대기하는 채널이 없으므로 데드락이 걸리지 않는다.
	close(c)
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	var msg string
	openA, openB := true, true
	// 루프를 두 번 돌면 채널이 모두 닫혔기 때문에 for 문 자체가 종료된다.
	for openA || openB {
		select {
		// 루프 돌 때마다 채널이 닫힘(수신할 것이 더 없음)상태가 된다.
		case msg, openA = <-a:
			if openA {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, openB = <-b:
			if openB {
				fmt.Printf("%s from B\n", msg)
			}
		}
	}
}

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
