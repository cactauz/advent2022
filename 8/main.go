package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input
var input string

// row -> column
var trees [][]int

var nrows, ncols int

func init() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()

		row := make([]int, 0, len(line))
		for _, r := range line {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			row = append(row, v)
		}

		trees = append(trees, row)
	}

	nrows = len(trees)
	ncols = len(trees[0])
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
	visibility := make([][]bool, 0, len(trees))

	for range trees {
		visibility = append(visibility, make([]bool, ncols))
	}

	for i := 0; i < nrows; i++ {
		max := -1
		for j := 0; j < ncols; j++ {
			if trees[i][j] > max {
				visibility[i][j] = true
				max = trees[i][j]
			}

			if max == 9 {
				break
			}
		}

		max = -1
		for j := ncols - 1; j >= 0; j-- {
			if trees[i][j] > max {
				visibility[i][j] = true
				max = trees[i][j]
			}

			if max == 9 {
				break
			}
		}
	}

	for j := 0; j < ncols; j++ {
		max := -1
		for i := 0; i < nrows; i++ {
			if trees[i][j] > max {
				visibility[i][j] = true
				max = trees[i][j]
			}

			if max == 9 {
				break
			}
		}

		max = -1
		for i := nrows - 1; i >= 0; i-- {
			if trees[i][j] > max {
				visibility[i][j] = true
				max = trees[i][j]
			}

			if max == 9 {
				break
			}
		}
	}

	count := 0
	for _, row := range visibility {
		for _, visible := range row {
			if visible {
				count++
			}
		}
	}

	fmt.Println("count:", count)

	return nil
}

func runB() error {
	var max int
	for row := range trees {
		for col := range trees[row] {
			score := score(row, col)
			if score > max {
				max = score
			}
		}
	}

	fmt.Println("max score:", max)

	return nil
}

func score(row, col int) int {
	height := trees[row][col]
	var score, count int

	for i := row + 1; i < nrows; i++ {
		count++

		if trees[i][col] >= height {
			break
		}
	}
	score = count

	count = 0
	for i := row - 1; i >= 0; i-- {
		count++

		if trees[i][col] >= height {
			break
		}
	}
	score *= count

	count = 0
	for j := col - 1; j >= 0; j-- {
		count++

		if trees[row][j] >= height {
			break
		}
	}
	score *= count

	count = 0
	for j := col + 1; j < ncols; j++ {
		count++

		if trees[row][j] >= height {
			break
		}
	}
	score *= count

	return score
}
