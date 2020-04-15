package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	//var a = 1
	//var b = 1
	//var (
	//	a int = 1
	//	b = 1
	//)
	a := 1
	b := 1
	//fmt.Print(a)
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	//fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	t.Log(a, b)
	//tmp := a
	//a = b
	//b = tmp
	a, b = b, a
	t.Log(a, b)
}
