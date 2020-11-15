package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s []int
	t.Log(len(s), cap(s))
	s = append(s, 1)
	t.Log(len(s), cap(s))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 1)
	t.Log(s2)
	t.Log(len(s2), cap(s2))
}
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 1; i <= 10; i++ {
		s = append(s, i)
		t.Log(i, len(s), cap(s))
	}
}
func TestSliceShareMemory(t *testing.T) {
	months := []string{
		"Jan", "Feb", "Mar", "Apr", "May",
		"Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	q2 := months[3:6]
	t.Log(q2, len(q2), cap(q2))
	summer := months[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknow"
	t.Log(q2)
	t.Log(months)
}
func TestSliceComparing(t *testing.T) {
	// a := []int{1, 2, 3, 4}
	// b := []int{1, 2, 3, 4}
	// if a == b { //slice can only be compared to nil
	// 	t.Log("equal")
	// }
}
