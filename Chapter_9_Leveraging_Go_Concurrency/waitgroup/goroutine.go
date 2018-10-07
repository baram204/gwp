package main

import "fmt"
import "time"
import "sync"

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// 웨잇그룹 2 개
	wg.Add(2)
	go printNumbers2(&wg) // done 해서 1차감
	go printLetters2(&wg) // done 해서 1차감

	// 원래라면 고루틴이 한참 실행되고 있을 때 main() 쓰레드는 종료되어버린다.
	// 하지만 wg 를 통해서 특정 카운트가 0이 되기 전 까지 쓰레드 종료가 미뤄진다.
	// 즉 고루틴끼리 다 제각각 노는데, 고루틴 쪽에서 메인에 있는 wg 에 영향을 미쳐서 메인을 대기시킴
	wg.Wait()
}
