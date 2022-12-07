package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "embed"
)

type Dir struct {
	Parent  *Dir
	Subdirs map[string]*Dir
	Files   map[string]int
}

func NewDir(parent *Dir) *Dir {
	return &Dir{
		Parent:  parent,
		Subdirs: make(map[string]*Dir),
		Files:   make(map[string]int),
	}
}

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
	root := parseInput()
	fmt.Println("total:", searchA(root))
	return nil
}

func searchA(dir *Dir) int {
	var total int

	for _, subdir := range dir.Subdirs {
		size := dirSize(subdir)
		if size < 100000 {
			total += size
		}

		total += searchA(subdir)
	}

	return total
}

func runB() error {
	root := parseInput()
	total := dirSize(root)
	free := 70000000 - total
	target := 30000000 - free

	fmt.Println("size of best:", searchB(target, root))
	return nil
}

func searchB(target int, dir *Dir) int {
	best := 99999999999999

	for _, subdir := range dir.Subdirs {
		size := dirSize(subdir)
		if size > target && size < best {
			best = size
		}

		subBest := searchB(target, subdir)
		if subBest < best {
			best = subBest
		}
	}

	return best
}

func parseInput() *Dir {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan() // skip $ cd /

	root := NewDir(nil)
	current := root

	// hacky but we can ignore ls, assume everything without a $ is
	// an item in the current dir, and that all dirs exist for cd bc
	// they were created when ls'd
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		if split[0] == "$" {
			if split[1] == "ls" {
				continue
			}

			// not ls so cd
			if split[2] == ".." {
				current = current.Parent
				continue
			}

			if sub, ok := current.Subdirs[split[2]]; ok {
				current = sub
				continue
			}

			panic("idk how to cd?????")
		}

		if split[0] == "dir" {
			current.Subdirs[split[1]] = NewDir(current)
			continue
		}

		// we must be a file
		size, err := strconv.Atoi(split[0])
		if err != nil {
			panic(fmt.Sprintf("cant parse size: %s", err))
		}

		current.Files[split[1]] = size
	}

	return root
}

func dirSize(d *Dir) int {
	var total int

	for _, size := range d.Files {
		total += size
	}

	for _, dir := range d.Subdirs {
		total += dirSize(dir)
	}

	return total
}
