package main

import (
	"slices"
	"testing"
)

func TestBuildLists(t *testing.T) {
	input := "1   2\n3   4\n"

	wantFirst := []int{1, 3}
	wantSecond := []int{2, 4}

	gotFirst, gotSecond := buildLists(input)

	if !slices.Equal(gotFirst, wantFirst) {
		t.Errorf("got %v for the first list want %v", gotFirst, wantFirst)
	}
	if !slices.Equal(gotSecond, wantSecond) {
		t.Errorf("got %v for the second list want %v", gotSecond, wantSecond)
	}
}

func TestDistance(t *testing.T) {
	first := []int{3, 4, 2, 1, 3, 3}
	second := []int{4, 3, 5, 3, 9, 3}

	got := distance(first, second)
	want := 11

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
