package main

import (
	"fmt"
	"sync"
)

func main() {
	// WaitGroup 생성. 2개의 Go루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done() //끝나면 .Done() 호출
		for i := 0; i < 4; i++ {
			//time.Sleep(1 * time.Microsecond)
			fmt.Printf("%d ", i)
		}
		fmt.Print("숫자출력 포문")
		// 채널에 값이 넣어지기 전 까지 고루틴은 지들끼리 돈다.
		//w <- true
	}()

	// 익명함수에 파라미터 전달
	go func(msg int) {
		defer wait.Done() //끝나면 .Done() 호출
		for i := msg; i < 10; i++ {
			fmt.Printf("%d ", i)
		}
	}(0)

	wait.Wait() //Go루틴 모두 끝날 때까지 대기
}
