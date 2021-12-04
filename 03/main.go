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

	gamma, epsilon := getShit(inputstrarr)
	fmt.Printf("Answer: %d\n", gamma*epsilon)
}

func getShit(vals []string) (gamma int, epsilon int) {
	// This loops over the length of each binary
	gammaBin := ""
	epsilonBin := ""
	for i := 0; i < len(vals[0]); i++ {
		// This loops over all values in the vals array.
		count1 := 0
		count0 := 0
		for a := 0; a < len(vals); a++ {
			switch string(vals[a][i]) {
			case "1":
				count1++
			case "0":
				count0++
			}
		}

		if count0 > count1 {
			gammaBin = gammaBin + "0"
			epsilonBin = epsilonBin + "1"
		} else {
			gammaBin = gammaBin + "1"
			epsilonBin = epsilonBin + "0"
		}
	}

	fmt.Println(gammaBin)
	gammaInt, err := strconv.ParseInt(gammaBin, 2, 32)
	if err != nil {
		panic(err)
	}
	gamma = int(gammaInt)

	fmt.Println(epsilonBin)
	epsilonInt, err := strconv.ParseInt(epsilonBin, 2, 32)
	if err != nil {
		panic(err)
	}
	epsilon = int(epsilonInt)

	return gamma, epsilon
}
