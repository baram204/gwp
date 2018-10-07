package main

import (
	"testing"
	"time"
)

func TestPrint1(t *testing.T){
	print2() // TestPrint1() 쓰레드와 동기화된 숫자/문자 출력 만복문을 실행
}

func TestGoPrint1(t *testing.T){
	goPrint2() // TestGoPrint1() 쓰레드와 비동기화된 각각의 고루틴화된 숫자/문자 출력 만복문을 실행

	// 고루틴이 완료되기 전에 현재 테스트가 종료되지 않도록 ,
	// 현재 테스트 스콥을 지속지킨다.
	time.Sleep(15 *time.Millisecond)
}
