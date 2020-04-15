package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	t.Log(m["a"])
	t.Logf("len m=%d", len(m))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	// 不能用cap访问map
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T)  {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k,v := range m{
		t.Log(k,v)
	}
}