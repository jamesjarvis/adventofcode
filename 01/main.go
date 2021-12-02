package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	if input == "" {
		panic("input cannot be empty")
	}
	inputstrarr := strings.Split(input, "\n")
	inputarr := make([]int, len(inputstrarr))
	for i, v := range inputstrarr {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		inputarr[i] = n
	}

	count := countLargerThanPrevious(inputarr)
	fmt.Printf("Answer: %d\n", count)
}

func countLargerThanPrevious(vals []int) int {
	var count int
	var topWindow int = vals[0] + vals[1] + vals[2]

	// top window = A,B,C
	//
	// On each check
	// W2 = W1 - v[i-3] + v[i]
	// if W2 > W1 count++
	// W1 = W2

	for i := 3; i < len(vals); i++ {
		bottomWindow := topWindow - vals[i-3] + vals[i]
		if bottomWindow > topWindow {
			count++
		}
		topWindow = bottomWindow
	}
	return count
}
