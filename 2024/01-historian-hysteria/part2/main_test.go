package main

import "testing"

func TestSimilarity(t *testing.T) {
	first := []int{3, 4, 2, 1, 3, 3}
	second := []int{4, 3, 5, 3, 9, 3}

	got := similarity(first, second)
	want := 31

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
