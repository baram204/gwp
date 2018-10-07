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
	wg.Wait()             // 일단 웨잇을 걸어 놓고 차감을 기다림
}
