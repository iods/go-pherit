package prose

import "testing"

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		// t.Error("did not match expected value")
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"orange", "apple", "cherry"}
	want := "orange, apple, and cherry"
	got := JoinWithCommas(list)
	if got != want {
		// t.Error("did not match expected value")
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
	}
}
