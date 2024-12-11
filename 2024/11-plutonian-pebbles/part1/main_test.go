package main

import (
	"slices"
	"testing"
)

func TestParseStones(t *testing.T) {
	text := "0 1 10 99 999"
	got := parseStones(text)
	want := []int{0, 1, 10, 99, 999}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFormatStones(t *testing.T) {
	stones := []int{0, 1, 10, 99, 999}
	got := formatStones(stones)
	want := "0 1 10 99 999"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestBlink(t *testing.T) {
}

func TestBlinkSeveral(t *testing.T) {

}
