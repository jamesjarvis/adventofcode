package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type command struct {
	bearing  string
	distance int
}

func main() {
	if input == "" {
		panic("input cannot be empty")
	}
	inputstrarr := strings.Split(input, "\n")
	inputarr := make([]command, len(inputstrarr))
	for i, v := range inputstrarr {
		commandarr := strings.Split(v, " ")
		n, err := strconv.Atoi(commandarr[1])
		if err != nil {
			panic(err)
		}
		c := command{
			bearing:  commandarr[0],
			distance: n,
		}
		inputarr[i] = c
	}

	horizontal, depth := getPosFromCommands(inputarr)
	fmt.Printf("Answer: %d\n", horizontal*depth)
}

func getPosFromCommands(vals []command) (horizontalPos int, depth int) {
	aim := 0
	for _, c := range vals {
		switch c.bearing {
		case "forward":
			horizontalPos += c.distance
			depth += aim * c.distance
		case "down":
			aim += c.distance
		case "up":
			aim -= c.distance
		default:
			fmt.Println("what da fuk", c.bearing)
		}
	}
	return horizontalPos, depth
}
