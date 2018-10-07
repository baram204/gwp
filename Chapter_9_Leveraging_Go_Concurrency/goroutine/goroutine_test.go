package main

import (
	"testing"
)

//import "time"

// Test cases

// normal run
//func TestPrint1(t *testing.T) {
//	print1()
//}
//
// // run with goroutines
//func TestGoPrint1(t *testing.T) {
//	goPrint1()
//	// 슬립이 없으면 Testprint1 이 끌났을 때 이미 TestGoPrint1 은 끝이 나서 Pass 가 되어버린다.
//	time.Sleep(1 * time.Millisecond)
//}
//

//func TestPrint1(t *testing.T) {
//	print2()
//}

// // run with goroutines and some work
//func TestGoPrint2(t *testing.T) {
//	goPrint2()
//	time.Sleep(2 * time.Millisecond)
//}
//
// // Benchmark cases
//
// normal run
func BenchmarkPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print1()
	}
}

// run with goroutines
func BenchmarkGoPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint1()
	}
}

// run with some work
func BenchmarkPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print2()
	}
}

// run with goroutines and some work
func BenchmarkGoPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint2()
	}
}
