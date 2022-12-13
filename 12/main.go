package main

import (
	_ "embed"
	"fmt"
	"os"
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

func parseInput() [][]int {
	lines := strings.Split(input, "\n")
	graph := make([][]int, 0, len(lines[0]))

	for y, line := range lines {
		for x, r := range line {
			if y == 0 {
				graph = append(graph, make([]int, len(lines)))
			}

			n := int(r - 'a')
			if r == 'E' {
				n = 'z' - 'a'
			}
			if r == 'S' {
				n = 0
			}
			graph[x][y] = n
		}
	}

	return graph
}

var start, end = p{0, 20}, p{43, 20}

type p struct {
	x, y int
}

func runA() error {
	graph := parseInput()
	visited := make(map[p]int)
	unvisited := make(map[p]int)

	for x, xs := range graph {
		for y := range xs {
			unvisited[p{x, y}] = 1e6
		}
	}

	distance := func(from, to p) int {
		diff := graph[to.x][to.y] - graph[from.x][from.y]
		if diff <= 1 {
			return 1
		} else {
			return 1e6
		}
	}

	neighbors := func(point p) []p {
		return []p{
			{point.x + 1, point.y},
			{point.x - 1, point.y},
			{point.x, point.y + 1},
			{point.x, point.y - 1},
		}
	}

	currentp, currentdist := start, 0
	for len(unvisited) > 0 {
		visited[currentp] = currentdist
		delete(unvisited, currentp)

		for _, n := range neighbors(currentp) {
			cur, ok := unvisited[n]
			if !ok {
				continue
			}

			dist := distance(currentp, n) + currentdist
			if dist < cur {
				unvisited[n] = dist
			}
		}

		next := min(unvisited)
		dist := unvisited[next]
		if next == end {
			fmt.Println("dist:", dist)
			break
		}

		currentp, currentdist = next, dist
	}

	return nil
}

func runB() error {
	graph := parseInput()
	visited := make(map[p]int)
	unvisited := make(map[p]int)

	for x, xs := range graph {
		for y := range xs {
			unvisited[p{x, y}] = 1e6
		}
	}

	distance := func(from, to p) int {
		diff := graph[from.x][from.y] - graph[to.x][to.y]
		if diff <= 1 {
			return 1
		} else {
			return 1e6
		}
	}

	neighbors := func(point p) []p {
		return []p{
			{point.x + 1, point.y},
			{point.x - 1, point.y},
			{point.x, point.y + 1},
			{point.x, point.y - 1},
		}
	}

	currentp, currentdist := end, 0
	for len(unvisited) > 0 {
		visited[currentp] = currentdist
		delete(unvisited, currentp)

		for _, n := range neighbors(currentp) {
			cur, ok := unvisited[n]
			if !ok {
				continue
			}

			dist := distance(currentp, n) + currentdist
			if dist < cur {
				unvisited[n] = dist
			}
		}

		next := min(unvisited)
		dist := unvisited[next]
		// if we reached a 0 elev we're done
		if graph[next.x][next.y] == 0 {
			fmt.Println("dist:", dist)
			break
		}

		currentp, currentdist = next, dist
	}

	return nil
}

func min(points map[p]int) p {
	// this is really bad hehe
	minp, mind := p{}, int(1e6)

	for p, d := range points {
		if d < mind {
			minp, mind = p, d
		}
	}

	return minp
}
