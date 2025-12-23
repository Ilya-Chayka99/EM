package Task_9

import (
	"errors"
	"testing"
)

func TestExistence(t *testing.T) {
	err := existence(0)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
