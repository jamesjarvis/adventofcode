package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var minPos, maxPos int

func main() {
	if input == "" {
		panic("input cannot be empty")
	}

	inputstrarr := strings.Split(input, ",")
	world := &World{
		crabs:                make([]int, 0, len(inputstrarr)),
		minPos:               1000000000,
		amountOfFuelRequired: 1000000000,
	}
	for _, val := range inputstrarr {
		n, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		if n < world.minPos {
			world.minPos = n
		}
		if n > world.maxPos {
			world.maxPos = n
		}
		world.crabs = append(world.crabs, n)
	}

	idealFuelRequired := world.IdealFuelRequired()
	fmt.Printf("Answer: %d\n", idealFuelRequired)
}

// Naive solution is to find the min and max values of the array, compute all fuel requirements for travelling
// to all points between and choose the best one?

type World struct {
	minPos, maxPos int

	crabs []int

	idealPos             int
	amountOfFuelRequired int
}

func (w *World) IdealFuelRequired() (fuel int) {
	for pos := w.minPos; pos <= w.maxPos; pos++ {
		fuelRequired := w.FuelRequired(pos)
		if fuelRequired < w.amountOfFuelRequired {
			w.idealPos = pos
			w.amountOfFuelRequired = fuelRequired
		}
	}
	return w.amountOfFuelRequired
}

// Now fuel required goes 1,3,6,10,15
func (w *World) FuelRequired(pos int) (fuel int) {
	for _, c := range w.crabs {
		if c < pos {
			fuel += triangularNumber(pos - c)
		} else {
			fuel += triangularNumber(c - pos)
		}
	}
	return fuel
}

// triangularNumber returns the amount of fuel required to reach a positional difference
// based on the pattern 1,3,6,10,15....
func triangularNumber(n int) int {
	return (n * (n + 1)) / 2
}
