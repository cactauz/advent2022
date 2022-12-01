package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var input = "./1/input"

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
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var max, current int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if current > max {
				max = current
			}
			current = 0
			continue
		}

		amount, err := strconv.Atoi(line)
		if err != nil {
			return err
		}

		current += amount
	}

	fmt.Println("max:", max)
	return scanner.Err()
}

func runB() error {
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var current int
	var totals []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			totals = append(totals, current)
			current = 0
			continue
		}

		amount, err := strconv.Atoi(line)
		if err != nil {
			return err
		}

		current += amount
	}

	sort.Ints(totals)
	fmt.Println("sum of top 3:", sum(totals[len(totals)-3:]))
	return scanner.Err()
}

func sum(ints []int) int {
	var total int
	for _, i := range ints {
		total += i
	}
	return total
}
