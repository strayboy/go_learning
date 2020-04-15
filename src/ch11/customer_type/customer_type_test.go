package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int
type OutConv func(op int) int

func timeSpent(inner IntConv) OutConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestParamsReturnAsFunc(t *testing.T) {
	tssf := timeSpent(slowFun)
	ret := tssf(1000)
	t.Log(ret)
}