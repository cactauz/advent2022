package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
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

func parseInput() [][]int {
	lines := strings.Split(input, "\n")
	res := make([][]int, 0, len(lines))
	re1 := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

	for _, line := range lines {
		if line == "" {
			continue
		}

		results := re1.FindAllStringSubmatch(line, -1)

		parsed := make([]int, 0, 4)
		for _, s := range results[0][1:] {
			v, _ := strconv.Atoi(s)
			parsed = append(parsed, v)
		}

		res = append(res, parsed)
	}

	return res
}

type sensor struct {
	x, y int
	rang int
}

func runA() error {
	y := 2000000
	input := parseInput()

	sensors := make([]sensor, 0, len(input))
	beacons := make([][]int, 0, len(input))
	spans := make([][]int, 0, len(sensors))

	marked := make(map[int]bool)

	for _, in := range input {
		sensor := sensor{
			x:    in[0],
			y:    in[1],
			rang: dist(in[0], in[1], in[2], in[3]),
		}
		sensors = append(sensors, sensor)
		beacons = append(beacons, []int{in[2], in[3]})

		dy := int(math.Abs(float64(sensor.y - y)))
		if dy > sensor.rang {
			continue
		}

		min := sensor.x - (sensor.rang - dy)
		max := sensor.x + (sensor.rang - dy)
		spans = append(spans, []int{min, max})

		for i := min; i <= max; i++ {
			marked[i] = true
		}
	}

	for _, b := range beacons {
		if b[1] == y {
			delete(marked, b[0])
		}
	}
	fmt.Println("marked:", len(marked))

	return nil
}

func dist(xa, ya, xb, yb int) int {
	return int(math.Abs(float64(xa-xb)) + math.Abs(float64(ya-yb)))
}

func runB() error {
	input := parseInput()

	sensors := make([]sensor, 0, len(input))
	beacons := make([][]int, 0, len(input))

	for _, in := range input {
		sensor := sensor{
			x:    in[0],
			y:    in[1],
			rang: dist(in[0], in[1], in[2], in[3]),
		}
		sensors = append(sensors, sensor)
		beacons = append(beacons, []int{in[2], in[3]})
	}

	slices.SortFunc(sensors, func(a, b sensor) bool {
		return a.rang < b.rang
	})

	// search every point on the edge of the sensors until we find one outside
	// the range of all of them
	for _, sensor := range sensors {
		rang := sensor.rang + 1

		// north
		for x := sensor.x - rang; x <= sensor.x+rang; x++ {
			y := sensor.y - (rang - (sensor.x - x))

			if x < 0 || y < 0 || x > 4000000 || y > 4000000 {
				continue
			}

			matched := false
			for _, other := range sensors {
				if dist(x, y, other.x, other.y) <= other.rang {
					matched = true
					break
				}
			}

			if !matched {
				fmt.Println("x:", x, "y:", y)
				fmt.Println(x*4000000 + y)
				return nil
			}
		}

		// south
		for x := sensor.x - rang; x <= sensor.x+rang; x++ {
			y := sensor.y + (rang - (sensor.x - x))

			if x < 0 || y < 0 || x > 4000000 || y > 4000000 {
				continue
			}

			matched := false
			for _, other := range sensors {
				if dist(x, y, other.x, other.y) <= other.rang {
					matched = true
					break
				}
			}

			if !matched {
				fmt.Println("x:", x, "y:", y)
				fmt.Println(x*4000000 + y)
				return nil
			}
		}
	}

	return nil
}
