package main

import (
	"fmt"
	"strconv"
	"strings"
)

func blink(stones []int) []int {
	return stones
}

func blinkSeveral(stones []int, times int) []int {
	for range times {
		stones = blink(stones)
	}
	return stones
}

func parseStones(text string) []int {
	stones := strings.Split(text, " ")
	parsedStones := make([]int, len(stones))
	for i, stone := range stones {
		parsedStones[i], _ = strconv.Atoi(stone)
	}
	return parsedStones
}

func formatStones(stones []int) string {
	unparsedStones := make([]string, len(stones))
	for i, stone := range stones {
		unparsedStones[i] = strconv.Itoa(stone)
	}
	return strings.Join(unparsedStones, " ")
}

func main() {
	input := "814 1183689 0 1 766231 4091 93836 46"
	times := 25
	orginalStones := parseStones(input)
	finalStones := blinkSeveral(orginalStones, times)
	result := formatStones(finalStones)
	fmt.Printf("Stones after blinking %d times:\n%v", times, result)
}
