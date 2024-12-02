package main

import (
	"fmt"
	"os"
)

func readContents(filename string) string {
	data, err := os.ReadFile(filename)
	check(err)
	content := string(data)
	return content
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./historian_hysteria input.txt")
		os.Exit(1)
	}
	fileName := os.Args[1]

	content := readContents(fileName)
	first, second := buildLists(content)
	distance := distanceBetweenLists(first, second)

	fmt.Printf("Total distance between lists: %d\n", distance)
}
