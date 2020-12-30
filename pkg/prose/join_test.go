package prose

import "testing"

func TestTwoElements(t *testing.T) {
	list := []string{"orange", "apple"}
	if JoinWithCommas(list) != "orange and apple" {
		t.Error("did not match expected value")
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"orange", "apple", "cherry"}
	if JoinWithCommas(list) != "orange, apple, and cherry" {
		t.Error("did not match the expected value")
	}
}
