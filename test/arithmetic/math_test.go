package arithmetic

import (
	"github.com/thedarksociety/go-pherit/pkg/arithmetic"
	"testing"
)

func TestAdd(t *testing.T) {
	if arithmetic.Add(1, 2) != 3 {
		t.Error("1 + 2 did not equal 3")
	}
}

func TestSubtract(t *testing.T) {
	if arithmetic.Subtract(2, 1) != 1 {
		t.Error("2 - 1 did not equal 1")
	}
}

func TestMultiply(t *testing.T) {
	if arithmetic.Multiply(2, 3) != 6 {
		t.Error("2 * 3 did not equal 6")
	}
}

func TestDivide(t *testing.T) {
	if arithmetic.Divide(-4, 2) != -2 {
		t.Error("-4 divided by 2 did not equal -2")
	}
}

func TestRemainder(t *testing.T) {
	if arithmetic.Remainder(5, 2) != 1 {
		t.Error("the remainder of 5 / 2 did not equal 1")
	}
}

func TestNegate(t *testing.T) {
	if arithmetic.Negate(-2) != 2 {
		t.Error("-2 did not turn into a positive 2")
	}
}