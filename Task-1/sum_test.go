package Task_1

import "testing"

func TestSum(t *testing.T) {
	num := Sum(1, 2)
	if num != 3 {
		t.Errorf("Sum fail")
	}
}
