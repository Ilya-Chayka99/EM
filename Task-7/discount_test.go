package Task_7

import (
	"testing"
	"time"
)

func TestDiscount(t *testing.T) {
	timeNow = func() time.Time {
		return time.Date(2025, 12, 11, 10, 0, 0, 0, time.UTC)
	}
	defer func() {
		timeNow = time.Now
	}()

	if discount() != 10 {
		t.Fatalf("discount should be 10, got %d", discount())
	}

}
