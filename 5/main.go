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

	stacks := parseStacks(scanner)
	// read empty line before instructions
	scanner.Scan()

	for scanner.Scan() {
		count, from, to := parseInstruction(scanner.Text())

		for i := count; i > 0; i-- {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	fmt.Print("top items: ")
	for i := 1; i <= 9; i++ {
		fmt.Print(stacks[i][len(stacks[i])-1])
	}
	fmt.Println()

	return nil
}

func runB() error {
	scanner := bufio.NewScanner(strings.NewReader(input))

	stacks := parseStacks(scanner)
	// read empty line before instructions
	scanner.Scan()

	for scanner.Scan() {
		count, from, to := parseInstruction(scanner.Text())

		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-count:]...)
		stacks[from] = stacks[from][:len(stacks[from])-count]
	}

	fmt.Print("top items: ")
	for i := 1; i <= 9; i++ {
		fmt.Print(stacks[i][len(stacks[i])-1])
	}
	fmt.Println()

	return nil
}

// stack number -> stack items 0 = bottom
func parseStacks(scanner *bufio.Scanner) map[int][]string {
	stacks := make(map[int][]string, 9)

	for scanner.Scan() {
		line := scanner.Text()
		if line[1] == '1' {
			break
		}

		// parse row
		for i := 0; i < 9; i++ {
			r := line[i*4+1]
			if r == ' ' {
				continue
			}
			// push to front so lowest box is 0
			stacks[i+1] = append([]string{string(r)}, stacks[i+1]...)
		}
	}

	return stacks
}

func parseInstruction(instruction string) (count, from, to int) {
	split := strings.Split(instruction, " ")
	count, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}

	from, err = strconv.Atoi(split[3])
	if err != nil {
		panic(err)
	}

	to, err = strconv.Atoi(split[5])
	if err != nil {
		panic(err)
	}

	return count, from, to
}
