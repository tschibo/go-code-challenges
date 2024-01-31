package conway

import (
	"strings"
)

func string2board(input string) [5][5]bool {
	var board [5][5]bool
	lines := strings.Split(input, "\n")
	for row, line := range lines {
		for col, field := range line {
			if field == '.' {
				board[row][col] = false
			} else if field == '*' {
				board[row][col] = true
			} else {
				panic("not detected")
			}
		}
	}
	return board
}

func board2string(board [5][5]bool) string {
	var output string
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if board[row][col] {
				output += "*"
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return output
}

func cycle(board0 [5][5]bool) [5][5]bool {
	var board1 [5][5]bool
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			nb := countNeighbours(board0, row, col)
			cell(&board1, &board0, row, col, nb)
		}
	}
	return board1
}

func countNeighbours(board0 [5][5]bool, row int, col int) int {
	var nb int
	for dr := -1; dr < 2; dr++ {
		for dc := -1; dc < 2; dc++ {
			if col+dc < 0 || col+dc > 4 || row+dr < 0 || row+dr > 4 {
				// stay in board
				continue
			} else if dc == 0 && dr == 0 {
				// don't count self as neighbour
				continue
			}
			if board0[row+dr][col+dc] {
				nb++
			}
		}
	}
	return nb
}

func cell(board1 *[5][5]bool, board0 *[5][5]bool, row int, col int, nb int) {
	isAlive := board0[row][col]
	if isAlive {
		if nb < 2 || nb > 3 {
			// dying
		} else {
			// surviving
			board1[row][col] = true
		}
	} else if nb == 3 {
		// birth
		board1[row][col] = true
	}
}

func ConwayConv(input string) string {
	board0 := string2board(input)
	board1 := cycle(board0)
	output := board2string(board1)
	return output
}
