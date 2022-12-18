package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input
var input string

var sample = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5
`

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
	input := parseInput()

	space := make([][][]int, 23)
	for x := range space {
		ys := make([][]int, 23)
		for y := range ys {
			ys[y] = make([]int, 23)
		}
		space[x] = ys
	}

	for _, coord := range input {
		x, y, z := coord[0], coord[1], coord[2]
		space[x][y][z] = 1
	}

	count := 0

	for x := range space {
		for y := range space[x] {
			for z := range space[x][y] {
				if space[x][y][z] == 0 {
					continue
				}

				for _, d := range []int{-1, 1} {
					if x+d < 0 || space[x+d][y][z] == 0 {
						count++
					}

					if y+d < 0 || space[x][y+d][z] == 0 {
						count++
					}

					if z+d < 0 || space[x][y][z+d] == 0 {
						count++
					}
				}

			}
		}
	}

	fmt.Println("count:", count)

	return nil
}

func runB() error {
	input := parseInput()

	space := make([][][]int, 23)
	for x := range space {
		ys := make([][]int, 23)
		for y := range ys {
			ys[y] = make([]int, 23)
		}
		space[x] = ys
	}

	for _, coord := range input {
		x, y, z := coord[0], coord[1], coord[2]
		space[x][y][z] = 1
	}

	markOutside(0, 0, 0, space)

	count := 0

	for x := range space {
		for y := range space[x] {
			for z := range space[x][y] {
				if space[x][y][z] <= 0 {
					continue
				}

				for _, d := range []int{-1, 1} {
					if x+d < 0 || space[x+d][y][z] == -1 {
						count++
					}

					if y+d < 0 || space[x][y+d][z] == -1 {
						count++
					}

					if z+d < 0 || space[x][y][z+d] == -1 {
						count++
					}
				}

			}
		}
	}

	fmt.Println("count:", count)

	return nil
}

func markOutside(x, y, z int, space [][][]int) {
	space[x][y][z] = -1

	for _, d := range []int{-1, 1} {
		if x+d >= 0 && x+d < 23 && space[x+d][y][z] == 0 {
			markOutside(x+d, y, z, space)
		}

		if y+d >= 0 && y+d < 23 && space[x][y+d][z] == 0 {
			markOutside(x, y+d, z, space)
		}

		if z+d >= 0 && z+d < 23 && space[x][y][z+d] == 0 {
			markOutside(x, y, z+d, space)
		}
	}
}

func parseInput() [][]int {
	lines := strings.Split(input, "\n")
	res := make([][]int, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, ",")
		res = append(res, []int{mustParse(split[0]), mustParse(split[1]), mustParse(split[2])})
	}

	return res
}

func mustParse(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return i
}
