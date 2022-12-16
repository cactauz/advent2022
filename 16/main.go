package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
	"gonum.org/v1/gonum/stat/combin"
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
	nodes := parseInput()

	path := getNodePaths(nodes)

	fmt.Println(search(nodes, path, []string{}, "AA", 0, 0))

	return nil
}

var maxT = 30

func search(nodes map[string]node, paths func(a, b string) []string, visited []string, current string, t int, indent int) ([]string, int) {
	visited = append(visited, current)

	nodeflow, max := 0, 0
	p := []string{}
	if maxT > t {
		nodeflow = (maxT - t) * nodes[current].flow
		max = nodeflow
		p = []string{current}
	}

	// get all paths and calcuate value
	for n, node := range nodes {
		if slices.Index(visited, n) != -1 || n == current {
			continue
		}

		if node.flow == 0 {
			continue
		}

		path := paths(current, n)
		if len(path) == 0 {
			// no path
			continue
		}
		path = path[1:] // path starts with ourself

		destT := (t + len(path) + 1)
		if maxT-destT < 1 {
			continue
		}

		next, total := search(nodes, paths, visited, n, destT, indent+1)

		if total+nodeflow > max {
			max = total + nodeflow
			p = append([]string{current}, next...)
		}
	}

	return p, max
}

type outcome struct {
	MyDest  string
	MyDestT int

	EleDest  string
	EleDestT int

	MyPath  []string
	ElePath []string

	Total int
}

func genPerms(items []string) [][]string {
	if len(items) < 2 {
		return [][]string{}
	}
	if len(items) == 2 {
		return [][]string{items}
	}

	perms := combin.Permutations(len(items), 2)
	res := make([][]string, 0, len(perms))
	for _, perm := range perms {
		res = append(res, []string{items[perm[0]], items[perm[1]]})
	}

	return res
}

func genPairs(items []string) [][]string {
	if len(items) < 2 {
		return [][]string{}
	}
	if len(items) == 2 {
		return [][]string{items}
	}

	perms := combin.Combinations(len(items), 2)
	res := make([][]string, 0, len(perms))
	for _, perm := range perms {
		res = append(res, []string{items[perm[0]], items[perm[1]]})
	}

	return res
}

func runB() error {
	nodes := parseInput()
	path := getNodePaths(nodes)

	dests := make([]string, 0, len(nodes))
	for name, node := range nodes {
		if node.flow > 0 {
			dests = append(dests, name)
		}
	}

	outcomes := []*outcome{}

	for _, pair := range genPerms(dests) {
		outcomes = append(outcomes, &outcome{
			MyDest:   pair[0],
			EleDest:  pair[1],
			MyDestT:  len(path("AA", pair[0])),
			EleDestT: len(path("AA", pair[1])),
		})
	}

	now := time.Now()

	// js, _ := json.Marshal(outcomes)
	// fmt.Println(string(js))

	for t := 1; t < 26; t++ {
		fmt.Println(t, len(outcomes))

		if t > 15 && len(outcomes) > 2000000 {
			// discard the worst half i guess idk
			slices.SortFunc(outcomes, func(a, b *outcome) bool {
				return a.Total > b.Total
			})
			outcomes = outcomes[:len(outcomes)/2]
		}

		// js, _ := json.Marshal(outcomes)
		// fmt.Println(string(js))

		newOutcomes := make([]*outcome, 0, len(outcomes))
		for _, out := range outcomes {
			out.MyDestT--
			out.EleDestT--

			if out.EleDestT != 0 && out.MyDestT != 0 {
				newOutcomes = append(newOutcomes, out)
				continue
			}

			options := make([]string, 0, len(dests))
			for _, dest := range dests {
				if dest == out.MyDest {
					// already here
					continue
				}

				if dest == out.EleDest {
					continue
				}

				if slices.Index(out.MyPath, dest) != -1 {
					continue
				}

				if slices.Index(out.ElePath, dest) != -1 {
					continue
				}

				options = append(options, dest)
			}

			if out.MyDestT == 0 && out.EleDestT == 0 {
				for _, pair := range genPairs(options) {
					newOutcomes = append(newOutcomes, &outcome{
						MyDest:   pair[0],
						EleDest:  pair[1],
						MyDestT:  len(path(out.MyDest, pair[0])),
						EleDestT: len(path(out.EleDest, pair[1])),
						MyPath:   append(out.MyPath[:], out.MyDest),
						ElePath:  append(out.ElePath[:], out.EleDest),
						Total:    out.Total + (26-t)*(nodes[out.MyDest].flow+nodes[out.EleDest].flow),
					})
				}

				continue
			}

			if out.MyDestT == 0 {
				out.Total += (26 - t) * nodes[out.MyDest].flow

				if len(options) == 0 {
					// just return to the list
					newOutcomes = append(newOutcomes, out)
				} else {
					for _, opt := range options {
						newOutcomes = append(newOutcomes, &outcome{
							MyDest:   opt,
							MyDestT:  len(path(out.MyDest, opt)),
							EleDest:  out.EleDest,
							EleDestT: out.EleDestT,
							MyPath:   append(out.MyPath[:], out.MyDest),
							ElePath:  out.ElePath[:],
							Total:    out.Total,
						})
					}
				}

				continue
			}

			if out.EleDestT == 0 {
				out.Total += (26 - t) * nodes[out.EleDest].flow

				if len(options) == 0 {
					// just return to the list
					newOutcomes = append(newOutcomes, out)
				} else {
					for _, opt := range options {
						newOutcomes = append(newOutcomes, &outcome{
							MyDest:   out.MyDest,
							MyDestT:  out.MyDestT,
							EleDest:  opt,
							EleDestT: len(path(out.EleDest, opt)),
							MyPath:   out.MyPath[:],
							ElePath:  append(out.ElePath[:], out.EleDest),
							Total:    out.Total,
						})
					}
				}
			}
		}

		outcomes = newOutcomes
	}

	max := 0
	var res *outcome
	for _, out := range outcomes {
		if out.Total > max {
			max = out.Total
			res = out
		}
	}

	fmt.Println("max:", max)
	js, _ := json.Marshal(res)
	fmt.Println(string(js))
	fmt.Println("finished with", len(outcomes), "outcomes in", time.Since(now))

	return nil
}

type node struct {
	neighbors []string
	flow      int
}

func parseInput() map[string]node {
	lines := strings.Split(input, "\n")
	res := make(map[string]node, len(lines))
	re1 := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.*)`)

	for _, line := range lines {
		if line == "" {
			continue
		}

		results := re1.FindAllStringSubmatch(line, -1)

		neighbors := strings.Split(results[0][3], ", ")
		flow, _ := strconv.Atoi(results[0][2])

		res[results[0][1]] = node{
			neighbors: neighbors,
			flow:      flow,
		}
	}

	return res
}

func getNodePaths(nodes map[string]node) func(a, b string) []string {
	// floyd-warshall algorithm
	dists := make(map[string]map[string]int, len(nodes))
	next := make(map[string]map[string]string, len(nodes))

	i := 0
	for n, node := range nodes {
		dists[n] = make(map[string]int, len(nodes))

		for i := range dists[n] {
			dists[n][i] = 99999
		}
		next[n] = make(map[string]string, len(nodes))

		dists[n][n] = 0
		next[n][n] = n

		for _, neighbor := range node.neighbors {
			dists[n][neighbor] = 1
			next[n][neighbor] = neighbor
		}

		i++
	}

	dist := func(a, b string) int {
		d, ok := dists[a][b]
		if !ok {
			return 99999
		}

		return d
	}

	for k := range nodes {
		for i := range nodes {
			for j := range nodes {
				if dist(i, j) > dist(i, k)+dist(k, j) {
					dists[i][j] = dist(i, k) + dist(k, j)
					next[i][j] = next[i][k]
				}
			}
		}
	}

	return func(a, b string) []string {
		if _, ok := next[a][b]; !ok {
			return []string{}
		}

		path := []string{a}
		for a != b {
			a = next[a][b]
			path = append(path, a)
		}

		return path
	}
}
