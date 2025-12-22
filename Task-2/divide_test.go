package Task_2

import "testing"

func TestDivide(t *testing.T) {
	_, err := Divide(10, 0)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
