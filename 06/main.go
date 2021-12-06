package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
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
		feesh: []*Lanternfish{},
	}
	for _, val := range inputstrarr {
		n, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		world.feesh = append(world.feesh, &Lanternfish{internaltimer: n})
	}

	for i := 0; i < simulateNDays; i++ {
		now := time.Now()
		world.Tick()
		fmt.Printf("Step %d, %d fish, took %s\n", i, world.Size(), time.Since(now))
	}

	numberOfFish := world.Size()
	fmt.Printf("Answer: %d\n", numberOfFish)
}

type World struct {
	feesh []*Lanternfish

	day int
}

func (w *World) Size() int { return len(w.feesh) }

func (w *World) Tick() {
	newFeeshToAdd := []*Lanternfish{}
	for _, lf := range w.feesh {
		newFeesh := lf.Tick()
		if newFeesh != nil {
			newFeeshToAdd = append(newFeeshToAdd, newFeesh)
		}
	}
	w.feesh = append(w.feesh, newFeeshToAdd...)
	w.day++
}

type Lanternfish struct {
	internaltimer int
}

func (lf *Lanternfish) Tick() *Lanternfish {
	if lf.internaltimer == 0 {
		lf.internaltimer = eachFishReplicatesEveryNDays - 1
		return &Lanternfish{eachFishReplicatesEveryNDays + 1}
	}
	lf.internaltimer = lf.internaltimer - 1
	return nil
}
