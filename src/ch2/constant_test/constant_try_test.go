package constant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTryOffset(t *testing.T) {
	a := 7 // 0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	b := 1 // 0001
	t.Log(b&Readable == Readable, b&Writable == Writable, b&Executable == Executable)
}
