package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed inputBoards.txt
var inputBoards string

func main() {
	if input == "" {
		panic("input cannot be empty")
	}
	if inputBoards == "" {
		panic("inputBoards cannot be empty")
	}
	inputstrarr := strings.Split(input, ",")
	inputarr := make([]int, len(inputstrarr))
	for i, v := range inputstrarr {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		inputarr[i] = n
	}

	fmt.Println(inputarr)

	boards := [][][]int{}

	// string: " 3 15  0  2 22"
	// int: 1
	// float/real: 2.3
	// bool: true/false
	// arrays: [type, type]

	boardsstrarr := strings.Split(inputBoards, "\n")

	board := [][]int{}
	for _, line := range boardsstrarr {
		if line == "" {
			boards = append(boards, board)
			board = [][]int{}
			continue
		}
		vals := strings.Split(line, " ")

		boardline := []int{}
		for _, val := range vals {
			if val == "" {
				continue
			}
			n, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			boardline = append(boardline, n)
		}
		board = append(board, boardline)
	}

	fmt.Println(boards)

	// ox, co2 := getShit(inputstrarr)
	// fmt.Printf("Answer: %d\n", ox*co2)
}
