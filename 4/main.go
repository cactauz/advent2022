package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var count int
	for scanner.Scan() {
		line := scanner.Text()
		first, second := parseAssignments(line)

		if first.start <= second.start && first.end >= second.end {
			count++
			continue
		}

		if second.start <= first.start && second.end >= first.end {
			count++
		}
	}

	fmt.Println("count:", count)

	return scanner.Err()
}

func runB() error {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var count int
	for scanner.Scan() {
		line := scanner.Text()
		first, second := parseAssignments(line)

		if isBetween(first.start, second.start, second.end) {
			count++
			continue
		}

		if isBetween(first.end, second.start, second.end) {
			count++
			continue
		}

		if isBetween(second.start, first.start, first.end) {
			count++
			continue
		}

		if isBetween(second.end, first.start, first.end) {
			count++
		}
	}

	fmt.Println("count:", count)

	return scanner.Err()
}

type assignment struct {
	start, end int
}

func isBetween(x, start, end int) bool {
	return x >= start && x <= end
}

func parseAssignments(str string) (assignment, assignment) {
	assignments := strings.Split(str, ",")
	if len(assignments) != 2 {
		panic("unexpected assignments")
	}

	res := make([]assignment, 0, len(assignments))
	for _, a := range assignments {
		split := strings.Split(a, "-")
		if len(split) != 2 {
			panic("invalid assignment")
		}

		start, errA := strconv.Atoi(split[0])
		end, errB := strconv.Atoi(split[1])
		if errA != nil || errB != nil {
			panic("invalid assignment")
		}

		res = append(res, assignment{start, end})
	}

	return res[0], res[1]
}
