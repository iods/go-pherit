package prose

import (
	"fmt"
	"github.com/thedarksociety/go-pherit/pkg/prose"
	"testing"
)

func TestFirstLarger(t *testing.T) {
	want := 2
	got := prose.Larger(2, 1)
	if got != want {
		t.Error(errorMessage(2, 1, got, want))
	}
}

func TestSecondLarger(t *testing.T) {
	want := 8
	got := prose.Larger(4, 8)
	if got != want {
		t.Error(errorMessage(4, 8, got, want))
	}
}

func errorMessage(a int, b int, got int, want int)  string {
	return fmt.Sprintf("Larger(%d, %d) = %d, want %d", a, b, got, want)
}