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
	var last int = vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] <= last {
			continue
		}
		count++
		last = vals[i]
	}
	return count
}
