package main

import (
	_ "embed"
	"fmt"
	"os"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type monkey struct {
	items       []int
	operation   func(int) int
	testDivisor int
	trueMonkey  int
	falseMonkey int
}

var startingMonkeys = []monkey{
	{[]int{83, 62, 93}, func(i int) int { return i * 17 }, 2, 1, 6},
	{[]int{90, 55}, func(i int) int { return i + 1 }, 17, 6, 3},
	{[]int{91, 78, 80, 97, 79, 88}, func(i int) int { return i + 3 }, 19, 7, 5},
	{[]int{64, 80, 83, 89, 59}, func(i int) int { return i + 5 }, 3, 7, 2},
	{[]int{98, 92, 99, 51}, func(i int) int { return i * i }, 5, 0, 1},
	{[]int{68, 57, 95, 85, 98, 75, 98, 75}, func(i int) int { return i + 2 }, 13, 4, 0},
	{[]int{74}, func(i int) int { return i + 4 }, 7, 3, 2},
	{[]int{68, 64, 60, 68, 87, 80, 82}, func(i int) int { return i * 19 }, 11, 4, 5},
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
	rounds := 20
	inspections := make(map[int]int, 8)

	monkeys := append([]monkey{}, startingMonkeys...)

	for r := 0; r < rounds; r++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				inspections[i]++
				item = monkey.operation(item) / 3
				if item%monkey.testDivisor == 0 {
					monkeys[monkey.trueMonkey].items = append(
						monkeys[monkey.trueMonkey].items, item,
					)
				} else {
					monkeys[monkey.falseMonkey].items = append(
						monkeys[monkey.falseMonkey].items, item,
					)
				}
				monkeys[i].items = []int{}
			}
		}
	}

	keys := maps.Keys(inspections)
	slices.SortFunc(keys, func(a, b int) bool {
		return inspections[a] > inspections[b]
	})

	fmt.Println("level:", inspections[keys[0]]*inspections[keys[1]])

	return nil
}

func runB() error {
	rounds := 10000
	inspections := make(map[int]int, 8)

	monkeys := append([]monkey{}, startingMonkeys...)

	divisorProduct := 1
	for _, m := range monkeys {
		divisorProduct *= m.testDivisor
	}

	for r := 0; r < rounds; r++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				inspections[i]++
				item = monkey.operation(item) % divisorProduct

				if item%monkey.testDivisor == 0 {
					monkeys[monkey.trueMonkey].items = append(
						monkeys[monkey.trueMonkey].items, item,
					)
				} else {
					monkeys[monkey.falseMonkey].items = append(
						monkeys[monkey.falseMonkey].items, item,
					)
				}
			}

			monkeys[i].items = []int{}
		}
	}

	keys := maps.Keys(inspections)
	slices.SortFunc(keys, func(a, b int) bool {
		return inspections[a] > inspections[b]
	})

	fmt.Println("level:", inspections[keys[0]]*inspections[keys[1]])

	return nil
}
