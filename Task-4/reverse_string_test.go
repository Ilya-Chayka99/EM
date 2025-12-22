package Task_4

import (
	"testing"
	"testing/quick"
)

func TestReverseString_Quick(t *testing.T) {
	f := func(s string) bool {
		return ReverseString(ReverseString(s)) == s
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
