package main

import (
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

// converts a string "12   45\n67.. 89\n" to a pair of lists {12, 67}, {67, 89}
func buildLists(lists string) ([]int, []int) {
	pairs := strings.Split(lists, "\n")
	// remove the last element which is just the empty string ""
	pairs = pairs[:len(pairs)-1]

	length := len(pairs)
	first := make([]int, length)
	second := make([]int, length)
	for i, pair := range pairs {
		// a line is of the form "1234   5678"
		elems := strings.Split(pair, "   ")
		first[i] = toInt(elems[0])
		second[i] = toInt(elems[1])
	}
	return first, second
}

func distanceBetweenLists(first []int, second []int) int {
	slices.Sort(first)
	slices.Sort(second)

	distance := 0
	for i := range len(first) {
		val1 := first[i]
		val2 := second[i]
		if val1 > val2 {
			distance += (val1 - val2)
		} else {
			distance += (val2 - val1)
		}
	}
	return distance
}
