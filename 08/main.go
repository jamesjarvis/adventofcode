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

/*
Mate idk, but I'm going to follow this mapping:
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg

0: abc efg    6
1:   c  f  -- 2
2: a cde g    5
3: a cd fg    5
4:  bcd f  -- 4
5: ab d fg    5
6: ab defg    6
7: a c  f  -- 3
8: abcdefg -- 7
9: abcd fg    6


Aside from 1,4 which we already know, everything includes aaaa so we can ignore that
// To do that, take the different character from 1 and 7 and delete.
Aside from 1,4,7 which we already know, everything contains gggg, so we can ignore that
// To do that, idk yet.

2 is the only one that does not contain ff.

*/

/*
 dddd
e    a
e    a
 ffff
g    b
g    b
 cccc

acedgfb: 8
cdfbe: 5
gcdfa: 2
fbcad: 3
dab: 7
cefabd: 9
cdfgeb: 6
eafb: 4
cagedb: 0
ab: 1

*/

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

		segmentValuesplus := [][]string{}
		for _, s := range segmentValues {
			segmentValuesplus = append(segmentValuesplus, strings.Split(s, ""))
		}

		testValuesPlust := [][]string{}
		for _, s := range testValues {
			testValuesPlust = append(testValuesPlust, strings.Split(s, ""))
		}

		p := &Problem{
			segmentValues: segmentValuesplus,
			testValues:    testValuesPlust,
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
		p.Komput()

		for _, test := range p.testValues {
			val := p.WhatIsSegmentValue(test)
			fmt.Printf("%s = %d\n", test, val)
			if _, ok := digits[val]; ok {
				count++
			}
		}
	}
	return count
}

type Problem struct {
	segmentValues [][]string
	testValues    [][]string

	// foundValues will be an array where the index is the value and sthe value is the combination to achieve it?
	foundValues [10][]string
}

func (p *Problem) Komput() {
	length6digs := [][]string{}
	for _, segment := range p.segmentValues {
		// simple cases out of the way first, 1,4,7,8
		val := p.WhatIsSegmentValue(segment)
		if val == -1 {
			if len(segment) == 6 {
				length6digs = append(length6digs, segment)
			}
			continue
		}
		p.foundValues[val] = segment
	}

	// find digits of length 6 now...
	for {
		if len(length6digs) == 0 {
			// yay we found them all.
			break
		}

		// if we have a diff of 1 between segment 1 and this segment, this segment is 6.
		for i, val := range length6digs {
			if diff(p.foundValues[1], val) == 1 {
				p.foundValues[6] = val
				length6digs = append(length6digs[:i], length6digs[i+1:]...)
				break
			}
		}

		// if we have a diff of 1 between segment 4 and this segment, this segment is 0.
		for i, val := range length6digs {
			if diff(p.foundValues[4], val) == 1 {
				p.foundValues[0] = val
				length6digs = append(length6digs[:i], length6digs[i+1:]...)
				break
			}
		}

		// the last one must be 9...
		if len(length6digs) == 1 {
			p.foundValues[9] = length6digs[0]
			length6digs = nil
		}
	}
}

// diff returns the number of digits from a but not in b.
// eg: ab, acd would return 1 as the second does not contain "b".
func diff(sub, super []string) (count int) {
	for i := 0; i < len(sub); i++ {
		char := sub[i]
		for a := 0; a < len(super); a++ {

		}
		if !contains(super, char) {
			count++
		}
	}
	return count
}

// contains returns true if s contains char.
func contains(s []string, char string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == char {
			return true
		}
	}
	return false
}

// returns segment value, or -1 if not known yet.
func (p *Problem) WhatIsSegmentValue(testValue []string) int {
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

	for i, val := range p.foundValues {
		if len(val) == len(testValue) && diff(val, testValue) == 0 {
			return i
		}
	}

	// shi idk bro
	return -1
}
