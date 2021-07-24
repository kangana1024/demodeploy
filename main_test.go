package main

import "testing"

func TestAdd(t *testing.T) {
	want := 5
	got := Add(2, 3)

	if got != want {
		t.Fatalf("Want %d got %d", want, got)
	}

	if 1 != 1 {
		t.Fatal("1 != 1")
	}
}
