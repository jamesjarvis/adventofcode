package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const simulateNDays = 256
const eachFishReplicatesEveryNDays = 7

func main() {
	if input == "" {
		panic("input cannot be empty")
	}
	inputstrarr := strings.Split(input, ",")
	world := &World{
		feesh: [eachFishReplicatesEveryNDays + 2]int{},
	}
	for _, val := range inputstrarr {
		n, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		world.feesh[n] = world.feesh[n] + 1
	}

	for i := 0; i < simulateNDays; i++ {
		// now := time.Now()
		world.Tick()
		// fmt.Printf("Step %d, %d fish, took %s\n", i, world.Size(), time.Since(now))
	}

	numberOfFish := world.Size()
	fmt.Printf("Answer: %d\n", numberOfFish)
}

// Represent fish lifetimes as an array of number of fish of length eachFishReplicatesEveryNDays + 2,
// where the index of each element contains the count of fish at that age.

type World struct {
	feesh [eachFishReplicatesEveryNDays + 2]int
}

func (w *World) Size() int {
	count := 0
	for _, numFish := range w.feesh {
		count += numFish
	}
	return count
}

func (w *World) Tick() {
	// fmt.Println("Number of fish before:", w.Size())
	resetFish := w.feesh[0]
	// fmt.Println("Fish to reset:", resetFish)
	for day := 0; day < eachFishReplicatesEveryNDays-1; day++ {
		w.feesh[day] = w.feesh[day+1]
	}
	w.feesh[eachFishReplicatesEveryNDays-1] = w.feesh[eachFishReplicatesEveryNDays] + resetFish
	w.feesh[eachFishReplicatesEveryNDays] = w.feesh[eachFishReplicatesEveryNDays+1]
	w.feesh[eachFishReplicatesEveryNDays+1] = resetFish
	// fmt.Println("Number of fish after:", w.Size())
}
