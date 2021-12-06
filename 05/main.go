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
	board := Board{
		b: make(map[Coordinate]int),
	}
	var err error
	for _, inputline := range inputstrarr {
		lineSplit := strings.Split(inputline, " -> ")
		line := Line{}
		coord1 := strings.Split(lineSplit[0], ",")
		line.Start.X, err = strconv.Atoi(coord1[0])
		if err != nil {
			panic(err)
		}
		line.Start.Y, err = strconv.Atoi(coord1[1])
		if err != nil {
			panic(err)
		}

		coord2 := strings.Split(lineSplit[1], ",")
		line.End.X, err = strconv.Atoi(coord2[0])
		if err != nil {
			panic(err)
		}
		line.End.Y, err = strconv.Atoi(coord2[1])
		if err != nil {
			panic(err)
		}

		board.Ingest(line)
	}

	numIntersectingPoints := board.NumPointsOverlapping()
	fmt.Printf("Answer: %d\n", numIntersectingPoints)
}

type Line struct {
	Start Coordinate
	End   Coordinate
}

type Coordinate struct {
	X int
	Y int
}

func (l Line) IsHorizontal() bool { return l.Start.X == l.End.X }

func (l Line) IsVertical() bool { return l.Start.Y == l.End.Y }

type Board struct {
	b map[Coordinate]int
}

func (b *Board) Ingest(l Line) {
	// Since we are only ever moving at 45 degree angles, this should be fine...

	var dx, dy int
	if l.Start.X < l.End.X {
		dx = 1
	} else if l.Start.X > l.End.X {
		dx = -1
	}
	if l.Start.Y < l.End.Y {
		dy = 1
	} else if l.Start.Y > l.End.Y {
		dy = -1
	}

	currentCoordinate := l.Start

	for {
		b.b[currentCoordinate]++

		if currentCoordinate == l.End {
			break
		}

		currentCoordinate.X += dx
		currentCoordinate.Y += dy
	}
}

func (b *Board) NumPointsOverlapping() (numOverlapping int) {
	for _, count := range b.b {
		// fmt.Printf("coord %+v, count %d\n", coord, count)
		if count > 1 {
			numOverlapping++
		}
	}
	return numOverlapping
}
