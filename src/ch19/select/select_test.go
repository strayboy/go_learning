package select__test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}
func AsyncService() chan string {
	//retChan := make(chan string)
	retChan := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("service returned result.")
		retChan <- ret
		fmt.Println("service exited.")
	}()
	return retChan
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		//time.After(time.Millisecond * 100) 是一个channel
		//t.Log(<-time.After(time.Millisecond * 100))

		t.Error("time out")
	}
}
