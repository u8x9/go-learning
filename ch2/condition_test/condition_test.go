package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			t.Log(i, "零")
		case 2, 4:
			t.Log(i, "偶数")
		case 1, 3:
			t.Log(i, "奇数")
		default:
			t.Log(i, "其它")
		}
	}
}
func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i == 0:
			t.Log(i, "零")
		case i%2 == 0:
			t.Log(i, "偶数")
		case i%2 == 1:
			t.Log(i, "奇数")
		default:
			t.Log(i, "其它")
		}
	}
}
