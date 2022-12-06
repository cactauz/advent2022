package main

import (
	"fmt"
	"os"

	_ "embed"
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
	for i := 0; i < len(input); i++ {
		if !hasRepeats(input[i : i+4]) {
			fmt.Println("index:", i+4)
			return nil
		}
	}

	fmt.Println("not found????")

	return nil
}

func runB() error {
	for i := 0; i < len(input); i++ {
		if !hasRepeats(input[i : i+14]) {
			fmt.Println("index:", i+14)
			return nil
		}
	}

	fmt.Println("not found????")

	return nil
}

func hasRepeats(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return true
			}
		}
	}

	return false
}
