package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1, len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Log(m2, len(m2))

	m3 := make(map[int]int, 10)
	t.Log(m3, len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	if i, ok := m1[99]; ok {
		t.Log(i)
	} else {
		t.Log("Not exists")
	}
}

func TestMapTravel(t *testing.T) {
	m := map[int]int{1: 1, 2: 4, 3: 9, 4: 16}
	for k, v := range m {
		t.Logf("%d * %d = %d\n", k, k, v)
	}
}
