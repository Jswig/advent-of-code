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

func TestBlink(t *testing.T) {
	stones := []int{0, 1, 10, 99, 999}
	got := blink(stones)
	want := []int{1, 2024, 1, 0, 9, 9, 2021976}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestBlinkSeveral(t *testing.T) {
	stones := []int{125, 17}
	got := blinkSeveral(stones, 6)
	want := []int{
		2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48,
		2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2,
	}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
