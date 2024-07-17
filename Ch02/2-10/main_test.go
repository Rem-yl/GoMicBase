package main

import "testing"

func TestSum(t *testing.T) {
	res := Sum(2)
	if res != 3 {
		t.Errorf("Sum(2) = %d; want 3", res)
	}
}
