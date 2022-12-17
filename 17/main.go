package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"
)

//go:embed input
var input string

var blocks = [][][]bool{
	{
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false},
		{true, true, true, true},
	},
	{
		{false, false, false, false},
		{false, true, false, false},
		{true, true, true, false},
		{false, true, false, false},
	},
	{
		{false, false, false, false},
		{false, false, true, false},
		{false, false, true, false},
		{true, true, true, false},
	},
	{
		{true, false, false, false},
		{true, false, false, false},
		{true, false, false, false},
		{true, false, false, false},
	},
	{
		{false, false, false, false},
		{false, false, false, false},
		{true, true, false, false},
		{true, true, false, false},
	},
}

func tryMove(board [][]bool, block [][]bool, r, c, rd, cd int) bool {
	for i, row := range block {
		for j, col := range row {
			if !col {
				continue
			}

			if c+j+cd < 0 || c+j+cd >= len(board[0]) {
				return false
			}

			if r-i+rd < 0 {
				return false
			}

			// board rows go down so bottom row = 0
			if board[r-i+rd][c+j+cd] {
				return false
			}
		}
	}

	return true
}

func findTop(board [][]bool) int {
	for row := len(board) - 1; row >= 0; row-- {
		for _, col := range board[row] {
			if col {
				return row
			}
		}
	}

	return -1
}

func main() {
	err := runA()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = runB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runA() error {
	board := make([][]bool, 0, 10000)

	reader := strings.NewReader(input)

	blockidx := 0
	count := 0

	for {
		start := findTop(board) + 7 // 3 spaces + 4 block height

		if start >= len(board) {
			n := start - len(board) + 1
			for i := 0; i < n; i++ {
				board = append(board, make([]bool, 7))
			}
		}

		block := blocks[blockidx]
		r, c := start, 2

		for {
			jet, _, err := reader.ReadRune()
			if err == io.EOF {
				reader = strings.NewReader(input)
				continue
			}

			var dir int
			if jet == '>' {
				dir = 1
			} else if jet == '<' {
				dir = -1
			} else {
				continue
			}

			if tryMove(board, block, r, c, 0, dir) {
				c += dir
			}

			if tryMove(board, block, r, c, -1, 0) {
				r -= 1
			} else {
				for i := 0; i < 4; i++ {
					for j := 0; j < 4; j++ {
						if c+j >= len(board[r-i]) {
							continue
						}

						board[r-i][c+j] = board[r-i][c+j] || block[i][j]
					}
				}
				break
			}

		}

		count++

		if count == 2022 {
			fmt.Println("2022 height:", findTop(board)+1)
			return nil
		}

		blockidx = (blockidx + 1) % len(blocks)
	}
}

func runB() error {
	board := make([][]bool, 0, 200000)

	reader := strings.NewReader(input)

	blockidx := 0
	count := 0
	offset := 0

	// ok, after 3447 we go into a cycle where the count increases by 1725 at each reset of the reader
	// and the height increases by 2685. (1 trillion - 3447) / 1725 is 579710142 with a remainder of 1603.
	// so simulate (3447 + 1603) count and then add (579710142 * 2685) to the height.
	// just figured this out empirically by observing the state when the reader was reset, idk if i got lucky
	for {
		start := findTop(board) + 7 // 3 spaces + 4 block height

		if start-offset >= len(board) {
			n := start - len(board) + 1
			for i := 0; i < n; i++ {
				board = append(board, make([]bool, 7))
			}
		}

		block := blocks[blockidx]
		r, c := start, 2

		for {
			jet, _, err := reader.ReadRune()
			if err == io.EOF {
				reader = strings.NewReader(input)
				continue
			}

			var dir int
			if jet == '>' {
				dir = 1
			} else if jet == '<' {
				dir = -1
			} else {
				continue
			}

			if tryMove(board, block, r, c, 0, dir) {
				c += dir
			}

			if tryMove(board, block, r, c, -1, 0) {
				r -= 1
			} else {
				for i := 0; i < 4; i++ {
					for j := 0; j < 4; j++ {
						if c+j >= len(board[r-i]) {
							continue
						}

						board[r-i][c+j] = board[r-i][c+j] || block[i][j]
					}
				}
				break
			}

		}

		count++

		if count == 5050 { // 3447 + 1603
			fmt.Println("height:", offset+findTop(board)+1+(579710142*2685))
			return nil
		}

		blockidx = (blockidx + 1) % len(blocks)
	}
}

func printBoard(board [][]bool) {
	for row := len(board) - 1; row >= 0; row-- {
		fmt.Print("|")
		for _, col := range board[row] {
			if col {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|")
		fmt.Println()
	}

	fmt.Print("+", strings.Repeat("-", len(board[0])), "+", "\n")
}
