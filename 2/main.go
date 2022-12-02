package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

var opponentShapes = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}

var myShapes = map[string]string{
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

var outcomes = map[string]map[string]string{
	"rock": {
		"rock":     "draw",
		"paper":    "loss",
		"scissors": "win",
	},
	"paper": {
		"rock":     "win",
		"paper":    "draw",
		"scissors": "loss",
	},
	"scissors": {
		"rock":     "loss",
		"paper":    "win",
		"scissors": "draw",
	},
}

var choiceScores = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var outcomeScores = map[string]int{
	"win":  6,
	"draw": 3,
	"loss": 0,
}

func runA() error {
	f, err := os.Open("./2/input")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total int

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		if len(split) != 2 {
			return fmt.Errorf("unepected length %d", len(split))
		}

		mine, theirs := myShapes[split[1]], opponentShapes[split[0]]

		outcome := outcomes[mine][theirs]
		score := outcomeScores[outcome] + choiceScores[mine]

		total += score
	}

	fmt.Println("total:", total)

	return scanner.Err()
}

var decoder = map[string]map[string]string{
	"X": { // lose
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	},
	"Y": { // draw
		"rock":     "rock",
		"paper":    "paper",
		"scissors": "scissors",
	},
	"Z": { // win
		"rock":     "paper",
		"paper":    "scissors",
		"scissors": "rock",
	},
}

func runB() error {
	f, err := os.Open("./2/input")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total int

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		if len(split) != 2 {
			return fmt.Errorf("unepected length %d", len(split))
		}

		theirs := opponentShapes[split[0]]
		mine := decoder[split[1]][theirs]

		outcome := outcomes[mine][theirs]
		score := outcomeScores[outcome] + choiceScores[mine]

		total += score
	}

	fmt.Println("total:", total)

	return scanner.Err()
}
