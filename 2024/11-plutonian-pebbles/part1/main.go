package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hasEvenNumberDigits(k int) bool {
	if (len(strconv.Itoa(k)) % 2) == 0 {
		return true
	} else {
		return false
	}
}

func splitDigits(k int) []int {
	digits := strconv.Itoa(k)
	midPoint := len(digits) / 2
	first, _ := strconv.Atoi(digits[:midPoint])
	second, _ := strconv.Atoi(digits[midPoint:])
	return []int{first, second}
}

func blink(stones []int) []int {
	// we know that the new stones will be at least the length of the existing stones
	// so we allocate the an array of that size upfront
	newStones := make([]int, 0, len(stones))
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if hasEvenNumberDigits(stone) {
			newStones = append(newStones, splitDigits(stone)...)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}
	return newStones
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

func main() {
	input := "814 1183689 0 1 766231 4091 93836 46"
	times := 75
	orginalStones := parseStones(input)
	finalStones := blinkSeveral(orginalStones, times)
	result := len(finalStones)
	fmt.Printf("Number of stones after blinking %d times: %d\n", times, result)
}
