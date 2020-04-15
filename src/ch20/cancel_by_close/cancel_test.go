package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCanceled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		//fmt.Println(<-cancelChan)
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCanceled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled.")
		}(i, cancelChan)
	}
	//cancel_1(cancelChan)
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}
