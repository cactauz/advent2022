package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

//go:embed input
var input string

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
	scanner := bufio.NewScanner(strings.NewReader(input))
	checks := []int{20, 60, 100, 140, 180, 220}

	clock := 0
	adding, num := false, 0
	strength := 0
	x := 1

	for clock < 220 {
		clock++

		if slices.Index(checks, clock) != -1 {
			strength += (x * clock)
		}

		if adding {
			x += num
			adding = false
		} else {
			scanner.Scan()
			line := scanner.Text()
			command := strings.Split(line, " ")

			switch command[0] {
			case "noop":
				// nothin
			case "addx":
				adding = true
				num, _ = strconv.Atoi(command[1])
			}
		}

	}

	fmt.Println("total:", strength)

	return nil
}

func runB() error {
	scanner := bufio.NewScanner(strings.NewReader(input))

	clock := 0
	adding, num := false, 0
	x := 1

	for clock < 240 {
		pos := clock % 40
		clock++

		if pos == 0 {
			fmt.Println()
		}

		if math.Abs(float64(pos-x)) <= 1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		if adding {
			x += num
			adding = false
		} else {
			scanner.Scan()
			line := scanner.Text()
			command := strings.Split(line, " ")

			switch command[0] {
			case "noop":
				// nothin
			case "addx":
				adding = true
				num, _ = strconv.Atoi(command[1])
			}
		}

	}

	return nil
}
