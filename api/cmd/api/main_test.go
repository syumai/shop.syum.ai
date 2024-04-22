package main

import "testing"

func Test_Dummy(t *testing.T) {
	one := 1
	two := 2
	got := one + two
	want := 3
	if want != got {
		t.Fatalf("want: %d, got: %d", want, got)
	}
}
