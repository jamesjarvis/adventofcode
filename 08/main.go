package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var digitsUnderTest = map[int]struct{}{
	1: {},
	4: {},
	7: {},
	8: {},
}

func main() {
	if input == "" {
		panic("input cannot be empty")
	}

	problems := []*Problem{}

	inputstrarr := strings.Split(input, "\n")
	for _, line := range inputstrarr {
		lineSplit := strings.Split(line, " | ")
		segmentValues := strings.Split(lineSplit[0], " ")
		testValues := strings.Split(lineSplit[1], " ")

		p := &Problem{
			segmentValues: segmentValues,
			testValues:    testValues,
		}
		// fmt.Printf("%+v\n", p)
		problems = append(problems, p)
	}
	world := &World{
		problems: problems,
	}

	HowManyTimesDoDigitsAppear := world.HowManyTimesDoDigitsAppear(digitsUnderTest)
	fmt.Printf("Answer: %d\n", HowManyTimesDoDigitsAppear)
}

type World struct {
	problems []*Problem
}

func (w *World) HowManyTimesDoDigitsAppear(digits map[int]struct{}) (count int) {
	for _, p := range w.problems {
		for _, test := range p.testValues {
			val := p.WhatIsSegmentValue(test)
			// fmt.Printf("%s = %d\n", test, val)
			if _, ok := digits[val]; ok {
				count++
			}
		}
	}
	return count
}

type Problem struct {
	segmentValues []string
	testValues    []string
}

func (p *Problem) WhatIsSegmentValue(testValue string) int {
	// A simple test for getting values 1,4,7,8 is to check length
	// as the number of segments is unique for these values.
	switch len(testValue) {
	case 2:
		return 1
	case 4:
		return 4
	case 3:
		return 7
	case 7:
		return 8
	}

	return 0
}
