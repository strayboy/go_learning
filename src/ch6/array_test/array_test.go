package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr1, arr2)
	arr[2] = 5
	t.Log(arr[0], arr[1], arr[2])
}

func TestArrayTravel(t *testing.T) {
	//arr := [...]int{1, 2, 3, 4, 5}
	//for i := 0; i < len(arr); i++ {
	//	t.Log(arr[i])
	//}
	//for i, val := range arr {
	//	t.Log(i, val)
	//}

	//for _, val := range arr {
	//	t.Log(val)
	//}
	arr2 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	for _, v := range arr2 {
		for _, r := range v {
			t.Log(r)
		}
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arrSec := arr[:3]
	a := arr[3:]
	t.Log(arrSec, a)
}
