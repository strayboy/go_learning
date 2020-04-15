package interface__test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (gopher *GoProgrammer) WriteHelloWorld() string {
	return "fmt.PrintLn(\"Hello World\")"
}

func TestClient(t *testing.T) {
	//var p Programmer
	//p = new(GoProgrammer)
	p := new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
