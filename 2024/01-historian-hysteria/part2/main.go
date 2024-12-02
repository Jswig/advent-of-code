package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// reads the contents of a text files into a string
func readContents(filename string) string {
	data, err := os.ReadFile(filename)
	check(err)
	content := string(data)
	return content
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

// computes the similarity score between two lists
func similarity(first []int, second []int) int {
	occurences := make(map[int]int)
	// compute all occurences in the second list
	for _, num := range second {
		_, present := occurences[num]
		if present {
			occurences[num] += 1
		} else {
			occurences[num] = 1
		}
	}

	score := 0
	for _, num := range first {
		count, present := occurences[num]
		if present {
			score += (num * count)
		}
	}

	return score
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s input.txt\n", os.Args[0])
		os.Exit(1)
	}
	fileName := os.Args[1]

	content := readContents(fileName)
	first, second := buildLists(content)
	similarity := similarity(first, second)

	fmt.Printf("Similarity between lists: %d\n", similarity)
}
