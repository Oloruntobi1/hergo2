package main

import "testing"



func TestSumTwoNumbers(t *testing.T) {

	got := SumTwoNumbers(3, 4)
	want := 7

	if got != want {
		t.Errorf("Wanted %d got %d", want, got)
	}

}