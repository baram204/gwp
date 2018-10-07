package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {

	// 함수를 별도의 쓰레드에서 비동기적으로 실행
	// 즉 컴파일러는 go 루틴 다 건너 뛰고
	go say("비동기1")
	go say("비동기2")
	go say("비동기3")

	// 아래 say 실행 후 프로그램 종료시켜버림

	// 함수를 main() 쓰레드와 동기적으로 실행
	say("main() 쓰레드에서 동기 실행")

	// 3초 대기를 안 하면
	// main() 쓰레드와 다른 쓰레드에서 돌아가는 고루틴이 실행되기도 전에
	// main() 쓰레드가 끝나 버려서 화면에 고루틴의 실행 결과가 보여지지 않는다.
	time.Sleep(time.Second * 1)
}
