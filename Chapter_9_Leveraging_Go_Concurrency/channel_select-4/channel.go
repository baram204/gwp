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
	msg1, msg2 := "A", "B"
	oc1, oc2 := false, false
	for {
		// 데드락이 발생을 막고 싶다면 강제로 메인 프로세스를 끝내버리면 될까?
		fmt.Printf("for 문 시작 - A : %v, B : %v \n", oc1, oc2)
		fmt.Printf("for 문 시작 - A : %v, B : %v \n", msg1, msg2)

		// 처음엔 닫힘 상태였지만
		
		time.Sleep(1 * time.Microsecond)
		select {
		case msg1, oc1 = <-a:
			// 수신 후에는 열림으로 변한다
			fmt.Printf("A : %v \n", oc1)
			fmt.Printf("%s from A\n", msg1)
		case msg2, oc2 = <-b:
			fmt.Printf("B : %v \n", oc2)
			fmt.Printf("%s from B\n", msg2)
		// default:
		// 	fmt.Println("Default")
		}

		fmt.Println("--------------------")
		time.Sleep(1 * time.Microsecond)

		// 두 번 돌면 모든 채널이 열림(수신할 것이 남아있다) 상태인데
		// 실제로 모든 고루틴은 슬립 상태이므로 데드락이 발생한다.

		// 따라서 루프가 한 번 더 돌면서 데드락을 발생하기 전에
		// 메시지를 검사해서 for 문을 빠져나간다.
		//if msg1 == "Hello World!" && msg2 == "Hola Mundo!" {
		//	break
		//}

		// 메시지를 검사하는 것이 아닌, 상태를 검사할 수도 있을 것이다.
		if oc1  && oc2 {
			break
		}
		// 하지만 채널이 모두 열려있다면 종료시킨다는 건 좀 이상하지 않을까?

	}
}
