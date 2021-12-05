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

	ox, co2 := getShit(inputstrarr)
	fmt.Printf("Answer: %d\n", ox*co2)
}

func getShit(vals []string) (ox int, co2 int) {
	// This function should call getVal for oxygen number (most common bitwise)
	// And for the co2 number (least common bitwise)
	// Then return both

	oxBin := getVal(vals, true)
	co2bin := getVal(vals, false)

	fmt.Println(oxBin)
	oxInt, err := strconv.ParseInt(oxBin, 2, 32)
	if err != nil {
		panic(err)
	}
	ox = int(oxInt)

	fmt.Println(co2bin)
	co2Int, err := strconv.ParseInt(co2bin, 2, 32)
	if err != nil {
		panic(err)
	}
	co2 = int(co2Int)

	return ox, co2
}

// gelVal gets a value based on the most common bit selector, or least common if false.
func getVal(vals []string, mostCommon bool) string {
	for i := 0; i < len(vals[0]); i++ {
		if len(vals) == 1 {
			return vals[0]
		}

		// This loops over all values in the vals array.
		filtered1 := []string{}
		filtered0 := []string{}
		for a := 0; a < len(vals); a++ {
			switch string(vals[a][i]) {
			case "1":
				filtered1 = append(filtered1, vals[a])
			case "0":
				filtered0 = append(filtered0, vals[a])
			}
		}

		if mostCommon {
			if len(filtered1) > len(filtered0) {
				vals = filtered1
				continue
			}
			if len(filtered0) > len(filtered1) {
				vals = filtered0
				continue
			}
			// if equal, return 1?
			vals = filtered1
			continue
		} else {
			if len(filtered1) < len(filtered0) {
				vals = filtered1
				continue
			}
			if len(filtered0) < len(filtered1) {
				vals = filtered0
				continue
			}
			// if equal, return 0?
			vals = filtered0
			continue
		}
	}
	return vals[0]
}
