package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 3, 4}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	// 按位置清零，xx &^ xx ,右边为1，为0，右边为0，不变
	a := 7 // 0111
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
