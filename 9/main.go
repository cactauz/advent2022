package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
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

type pos struct {
	x, y int
}

func runA() error {
	scanner := bufio.NewScanner(strings.NewReader(input))
	visited := map[pos]struct{}{}
	var head, tail pos
	visited[tail] = struct{}{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir := line[0]
		n, _ := strconv.Atoi(line[1])

		for i := 0; i < n; i++ {
			switch dir {
			case "U":
				head.y++
			case "D":
				head.y--
			case "R":
				head.x++
			case "L":
				head.x--
			}
			tail = moveTail(head, tail)
			visited[tail] = struct{}{}
		}
	}

	fmt.Println("total:", len(visited))
	return nil
}

func runB() error {
	scanner := bufio.NewScanner(strings.NewReader(input))
	visited := map[pos]struct{}{}
	knots := make([]pos, 10)
	visited[knots[9]] = struct{}{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir := line[0]
		n, _ := strconv.Atoi(line[1])

		for i := 0; i < n; i++ {
			switch dir {
			case "U":
				knots[0].y++
			case "D":
				knots[0].y--
			case "R":
				knots[0].x++
			case "L":
				knots[0].x--
			}

			for j, knot := range knots[1:] {
				knots[j+1] = moveTail(knots[j], knot)
			}

			visited[knots[9]] = struct{}{}
		}
	}

	fmt.Println("total:", len(visited))
	return nil
}

func touching(a, b pos) bool {
	return math.Abs(float64(a.x-b.x)) <= 1 &&
		math.Abs(float64(a.y-b.y)) <= 1
}

func moveTail(h, t pos) pos {
	if touching(h, t) {
		return t
	}

	var new pos

	xd := h.x - t.x
	if xd > 0 {
		new.x = t.x + 1
	} else if xd < 0 {
		new.x = t.x - 1
	} else {
		new.x = t.x
	}

	yd := h.y - t.y
	if yd > 0 {
		new.y = t.y + 1
	} else if yd < 0 {
		new.y = t.y - 1
	} else {
		new.y = t.y
	}

	return new
}
