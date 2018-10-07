package main

import (
	"testing"
)

func BenchmarkPrint1(b *testing.B){
	for i:=0; i<b.N; i++{
		print1() // TestPrint1() 쓰레드와 동기화된 숫자/문자 출력 만복문을 실행
	}
}

func BenchmarkGoPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint1() // TestGoPrint1() 쓰레드와 비동기화된 각각의 고루틴화된 숫자/문자 출력 만복문을 실행
	}
}

func BenchmarkPrint2(b *testing.B){
	for i:=0; i<b.N; i++{
		print2() // TestPrint1() 쓰레드와 동기화된 숫자/문자 출력 만복문을 실행
	}
}

func BenchmarkGoPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint2() // TestGoPrint1() 쓰레드와 비동기화된 각각의 고루틴화된 숫자/문자 출력 만복문을 실행
	}
}

/*



goos: windows
goarch: amd64

1 cpu

BenchmarkPrint1   	20000000	        74.1 ns/op
BenchmarkGoPrint1 	 1000000	      1003 ns/op
BenchmarkPrint2   	20000000	        70.0 ns/op
BenchmarkGoPrint2 	 1000000	      1055 ns/op
PASS

2 cpu - 100 배 작업이면 750 일텐데 519로 줄었다.

BenchmarkPrint1-2     	20000000	        75.3 ns/op
BenchmarkGoPrint1-2   	 3000000	       519 ns/op
BenchmarkPrint2-2     	20000000	        68.1 ns/op
BenchmarkGoPrint2-2   	 3000000	       515 ns/op

3 cpu - 절반 수준으로 줄어들었다.

BenchmarkPrint1-3     	20000000	        73.7 ns/op
BenchmarkGoPrint1-3   	 5000000	       313 ns/op
BenchmarkPrint2-3     	20000000	        71.3 ns/op
BenchmarkGoPrint2-3   	 5000000	       318 ns/op

4 cpu - cpu 스케쥴링에 소모되는 비용이 많아서 3 cpu 보다 느리다.

BenchmarkPrint1-4     	20000000	        74.0 ns/op
BenchmarkGoPrint1-4   	 5000000	       352 ns/op
BenchmarkPrint2-4     	20000000	        69.4 ns/op
BenchmarkGoPrint2-4   	 5000000	       356 ns/op

5 cpu - 점점 느려진다.

BenchmarkPrint1-5     	20000000	        75.4 ns/op
BenchmarkGoPrint1-5   	 5000000	       384 ns/op
BenchmarkPrint2-5     	20000000	        70.0 ns/op
BenchmarkGoPrint2-5   	 5000000	       387 ns/op

8 cpu - 8 cpu 나 썼는데 결과는 2 cpu 정도다.

BenchmarkPrint1-8     	20000000	        74.5 ns/op
BenchmarkGoPrint1-8   	 3000000	       417 ns/op
BenchmarkPrint2-8     	20000000	        70.2 ns/op
BenchmarkGoPrint2-8   	 3000000	       418 ns/op
*/