package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("Working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	//retChan := make(chan string)
	retChan := make(chan string,1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retChan <- ret
		fmt.Println("service exited.")
	}()
	return retChan
}

func TestAsyncService(t *testing.T) {
	retChan := AsyncService()
	otherTask()
	fmt.Println(<-retChan)
	time.Sleep(time.Second * 1)
}

func TestAsyncServiceBufferChannel(t *testing.T)  {
	retChan := AsyncService()
	otherTask()
	fmt.Println(<-retChan)
}