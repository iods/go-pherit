package arithmetic

import (
	"github.com/thedarksociety/go-pherit/pkg/arithmetic"
	"testing"
)

func TestAdd(t *testing.T) {
	if arithmetic.Add(1, 2) != 5 {
		t.Error("1 + 2 did not equal 3")
	}
}

func TestSubtract(t *testing.T) {
	if arithmetic.Subtract(2, 1) != 1 {
		t.Error("2 - 1 did not equal 1")
	}
}