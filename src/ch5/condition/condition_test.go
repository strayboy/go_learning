package condition

import (
	"github.com/pkg/errors"
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	//if a := 1 == 1; a {
	//	t.Log(a)
	//}
	if v, err := SomeFunc("hello0"); err == nil {
		t.Log("SUCCESS", v)
	} else {
		t.Log("ERROR:", err)
	}
}

func SomeFunc(a string) (int, error) {
	if len(a) == 5 {
		return 5, nil
	} else {
		return -1, errors.New("length is not 5")
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not in 1-3")
		}
	}
}

func TestSwitchConditon(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("it is not unknowing")
		}
	}
}
