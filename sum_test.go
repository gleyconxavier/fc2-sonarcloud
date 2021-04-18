package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be (5)")
	}

}

func TestSub(t *testing.T) {
	result := sub(42, 84)

	if result != -42 {
		t.Error("The result must be (-42)")
	}
}
