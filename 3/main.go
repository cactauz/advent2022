package main

import (
	"bufio"
	_ "embed"
	"strings"

	"fmt"
	"os"
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

// a = 97 A = 65
func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - 96
	}

	return int(r) - 64 + 26
}

func findDuplicate(a, b string) rune {
	for _, r := range a {
		if strings.ContainsRune(b, r) {
			return r
		}
	}

	panic("no dupe found")
}

func findDuplicateMany(strs ...string) rune {
	if len(strs) == 0 {
		panic("index out of range")
	}

loop:
	for _, r := range strs[0] {
		for _, other := range strs[1:] {
			if !strings.ContainsRune(other, r) {
				continue loop
			}
		}

		return r
	}

	panic("no dupe found")
}

func runA() error {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var total int

	for scanner.Scan() {
		line := scanner.Text()
		first, second := line[:len(line)/2], line[len(line)/2:]
		dupe := findDuplicate(first, second)

		total += priority(dupe)
	}

	fmt.Println("total:", total)

	return scanner.Err()
}

func runB() error {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var total int

loop:
	for {
		lines := make([]string, 0, 3)

		for i := 0; i < 3; i++ {
			if !scanner.Scan() {
				break loop
			}
			lines = append(lines, scanner.Text())
		}

		dupe := findDuplicateMany(lines...)
		total += priority(dupe)
	}

	fmt.Println("total:", total)

	return scanner.Err()
}
