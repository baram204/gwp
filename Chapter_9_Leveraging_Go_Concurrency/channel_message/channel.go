package main

import (
	"fmt"
	"time"
)

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		// 채널에 0부터 4까지 정수를 넣는다.
		c <- i
		fmt.Println("Threw  >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		// 채널에서 5번 값을 꺼낸다.
		num := <-c
		fmt.Println("Caught <<", num)
	}
}

func main() {
	// 버버 3개의 채널을 만든다.
	c := make(chan int, 3)
	go thrower(c)
	go catcher(c)
	time.Sleep(100 * time.Millisecond)
}
