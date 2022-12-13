package main

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
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

	index := 1
	total := 0
	for scanner.Scan() {
		var left, right []interface{}

		json.Unmarshal([]byte(scanner.Text()), &left)

		scanner.Scan()
		json.Unmarshal([]byte(scanner.Text()), &right)

		res := compare(left, right)
		if res == nil || *res {
			total += index
		}

		index++
		scanner.Scan() // consume blank line
	}

	fmt.Println("total:", total)

	return nil
}

// If both values are integers, the lower integer should come first. If the left integer is lower than the right integer, the inputs are in the right order. If the left integer is higher than the right integer, the inputs are not in the right order. Otherwise, the inputs are the same integer; continue checking the next part of the input.
// If both values are lists, compare the first value of each list, then the second value, and so on. If the left list runs out of items first, the inputs are in the right order. If the right list runs out of items first, the inputs are not in the right order. If the lists are the same length and no comparison makes a decision about the order, continue checking the next part of the input.
// If exactly one value is an integer, convert the integer to a list which contains that integer as its only value, then retry the comparison. For example, if comparing [0,0,0] and 2, convert the right value to [2] (a list containing 2); the result is then found by instead comparing [0,0,0] and [2].

func compare(left, right []interface{}) *bool {
	for i := range left {
		if i >= len(right) {
			return boolPtr(false)
		}

		if li, ok := left[i].(float64); ok {
			if ri, ok := right[i].(float64); ok {
				// both ints
				if li == ri {
					continue
				}

				if li < ri {
					return boolPtr(true)
				} else {
					return boolPtr(false)
				}
			}

			// only left is int
			left[i] = []interface{}{li}
			cmp := compare(left[i].([]interface{}), right[i].([]interface{}))
			if cmp != nil {
				return cmp
			}
			continue
		}

		if ri, ok := right[i].(float64); ok {
			// only right is int
			right[i] = []interface{}{ri}
		}

		// now comparing lists
		cmp := compare(left[i].([]interface{}), right[i].([]interface{}))
		if cmp != nil {
			return cmp
		}
	}

	if len(right) > len(left) {
		return boolPtr(true)
	}

	return nil
}

func boolPtr(b bool) *bool {
	return &b
}

func runB() error {
	lines := strings.Split(input, "\n")
	lines = append([]string{`[[2]]`, `[[6]]`}, lines...)

	packets := make([][]interface{}, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		var inf []interface{}
		json.Unmarshal([]byte(line), &inf)

		packets = append(packets, inf)
	}
	a, b := packets[0], packets[1]

	slices.SortFunc(packets, func(a, b []interface{}) bool {
		res := compare(a, b)
		return res == nil || *res
	})

	idxa := slices.IndexFunc(packets, func(inf []interface{}) bool {
		return compare(inf, a) == nil
	}) + 1

	idxb := slices.IndexFunc(packets, func(inf []interface{}) bool {
		return compare(inf, b) == nil
	}) + 1

	fmt.Println("code:", idxa*idxb)

	return nil
}
