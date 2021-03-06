package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed inputBoards.txt
var inputBoards string

// Board contains the board 2d array and also a map of seen numbers.
type Board struct {
	boardArr    [][]int
	seenNumbers map[int]struct{}
}

// checkBoard checks whether the board has won.
// The board is a 5x5 grid so that makes things a bit easier?
func (b *Board) checkBoard() bool {
	fmt.Printf("%s\n", b.seenNumbers)

	nHoriz := 0
	nVer := 0

	const size = 5

	// horiz: i == horiz, j == ver
	// ver: i == ver, j == horiz
	for i := 0; i < len(b.boardArr); i++ {
		nHoriz = 0
		nVer = 0
		for j := 0; j < len(b.boardArr[i]); j++ {
			if _, ok := b.seenNumbers[b.boardArr[i][j]]; ok {
				nHoriz++
			}
			if _, ok := b.seenNumbers[b.boardArr[j][i]]; ok {
				nVer++
			}
		}
		if nHoriz == size || nVer == size {
			return true
		}
	}

	return false
}

// markBoard just marks the board number as seen.
func (b *Board) markBoard(n int) {
	b.seenNumbers[n] = struct{}{}
}

// markBoard just marks the board number as seen.
func (b *Board) sumOfUnmarked() (sum int) {
	for i := 0; i < len(b.boardArr); i++ {
		for j := 0; j < len(b.boardArr[i]); j++ {
			if _, ok := b.seenNumbers[b.boardArr[i][j]]; !ok {
				sum += b.boardArr[i][j]
			}
		}
	}
	return sum
}

func main() {
	if input == "" {
		panic("input cannot be empty")
	}
	if inputBoards == "" {
		panic("inputBoards cannot be empty")
	}
	inputstrarr := strings.Split(input, ",")
	inputarr := make([]int, len(inputstrarr))
	for i, v := range inputstrarr {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		inputarr[i] = n
	}

	fmt.Println(inputarr)

	boards := []*Board{}

	boardsstrarr := strings.Split(inputBoards, "\n")

	board := [][]int{}
	for _, line := range boardsstrarr {
		if line == "" {
			boards = append(boards, &Board{
				boardArr:    board,
				seenNumbers: map[int]struct{}{},
			})
			board = [][]int{}
			continue
		}
		vals := strings.Split(line, " ")

		boardline := []int{}
		for _, val := range vals {
			if val == "" {
				continue
			}
			n, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			boardline = append(boardline, n)
		}
		board = append(board, boardline)
	}

	answer := ifMusicBeTheFoodOfLovePlayOn(boards, inputarr)
	fmt.Printf("Winner: %d\n", answer)
}

// ifMusicBeTheFoodOfLovePlayOn returns the sum of all unmarked numbers on the winning board times the winning number.
func ifMusicBeTheFoodOfLovePlayOn(boards []*Board, numbers []int) int {
	var lastWinningBoard *Board
	var lastWinningNumber int
	for _, call := range numbers {
		fmt.Printf("new call %d\n", call)
		for bi := 0; bi < len(boards); bi++ {
			board := boards[bi]
			board.markBoard(call)
			if board.checkBoard() {
				lastWinningBoard = board
				lastWinningNumber = call
				fmt.Printf("reached here, bi=%d, len=%d\n", bi, len(boards))
				if bi <= len(boards) {
					fmt.Println("reached even here")
					boards = append(boards[:bi], boards[bi+1:]...)
					bi = bi - 1
				}
				fmt.Printf("Found winning board! Removed from list %v\n", board.boardArr)
				fmt.Printf("new reached here here, bi=%d, len=%d\n", bi, len(boards))
				fmt.Printf("bi: %d, length of boards remaining: %d\n", bi, len(boards))
				if len(boards) == 0 {
					return lastWinningNumber * lastWinningBoard.sumOfUnmarked()
				}
			}
		}
	}
	return lastWinningNumber * lastWinningBoard.sumOfUnmarked()
}
