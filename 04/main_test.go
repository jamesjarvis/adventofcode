package main

import "testing"

func TestBoard(t *testing.T) {
	board := &Board{
		boardArr:    [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
		seenNumbers: map[int]struct{}{},
	}

	t.Run("horizontal first row", func(t *testing.T) {
		board.seenNumbers = map[int]struct{}{
			1: struct{}{},
			2: struct{}{},
			3: struct{}{},
			4: struct{}{},
			5: struct{}{},
		}
		if !board.checkBoard() {
			t.Error("Should be true")
		}
	})

	t.Run("vertical first row", func(t *testing.T) {
		board.seenNumbers = map[int]struct{}{
			1:  struct{}{},
			6:  struct{}{},
			11: struct{}{},
			16: struct{}{},
			21: struct{}{},
		}
		if !board.checkBoard() {
			t.Error("Should be true")
		}
	})

	t.Run("vertical random row", func(t *testing.T) {
		board.seenNumbers = map[int]struct{}{
			3:  struct{}{},
			8:  struct{}{},
			13: struct{}{},
			18: struct{}{},
			23: struct{}{},
		}
		if !board.checkBoard() {
			t.Error("Should be true")
		}
	})

	t.Run("horizontal random row", func(t *testing.T) {
		board.seenNumbers = map[int]struct{}{
			11: struct{}{},
			12: struct{}{},
			13: struct{}{},
			14: struct{}{},
			15: struct{}{},
		}
		if !board.checkBoard() {
			t.Error("Should be true")
		}
	})

	t.Run("horizontal last row", func(t *testing.T) {
		board.seenNumbers = map[int]struct{}{
			21: struct{}{},
			22: struct{}{},
			23: struct{}{},
			24: struct{}{},
			25: struct{}{},
		}
		if !board.checkBoard() {
			t.Error("Should be true")
		}
	})

	t.Run("vertical last row", func(t *testing.T) {
		board.seenNumbers = map[int]struct{}{
			5:  struct{}{},
			10: struct{}{},
			15: struct{}{},
			20: struct{}{},
			25: struct{}{},
		}
		if !board.checkBoard() {
			t.Error("Should be true")
		}
	})
}
