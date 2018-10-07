package main

import (
	"fmt"
	"time"
)

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		// 채널에 0부터 4까지 다섯개의 숫자를 넣는다.
		c <- i
		fmt.Println("Threw  >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		// <-c 에 의해서 무언가를 꺼내기 위해서 채널은 꺼내고 나서 블락이 걸린다.
		num := <-c
		fmt.Println("Caught <<", num)
	}
}

func main() {

	c := make(chan int)
	go thrower(c)
	go catcher(c)

	// 슬립을 걸어주지 않으면, 메인 쓰레드와 상관 없이
	// 고루틴이 지들끼리 채널로 통신을 하려고 하다가
	// 메인쓰레드가 종료되어 버림
	time.Sleep(100 * time.Millisecond)
}
