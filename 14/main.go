package main

import (
	_ "embed"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed input
var input string

var (
	air  = ""
	rock = "#"
	sand = "o"
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

	err = playAnimation()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runA() error {
	area := parseInput()
	ymax := len(area[0]) - 1
	xmax := len(area) - 1

	count := 0

loop:
	for {
		x, y := 500, 0

		for {
			if x >= xmax || y >= ymax {
				// we really fell off
				break loop
			}

			if below := area[x][y+1]; below == air {
				y++
				continue
			}

			if downleft := area[x-1][y+1]; downleft == air {
				x, y = x-1, y+1
				continue
			}

			if x+1 == len(area) {
				// we will off the right side
				break
			}

			if downright := area[x+1][y+1]; downright == air {
				x, y = x+1, y+1
				continue
			}

			break
		}

		area[x][y] = sand

		count++
	}

	// printArea(area)
	fmt.Println("count", count)

	return nil
}

func runB() error {
	area := parseInput()

	// just add in an arbitrarily large floor by extending x to 1000 and adding the 2 new ys
	for x := range area {
		area[x] = append(area[x], air, rock)
	}

	for x := len(area); x < 1000; x++ {
		area = append(area, make([]string, len(area[0])-1))
		area[x] = append(area[x], rock)
	}

	count := 0

	for {
		x, y := 500, 0

		for {
			if below := area[x][y+1]; below == air {
				y++
				continue
			}

			if downleft := area[x-1][y+1]; downleft == air {
				x, y = x-1, y+1
				continue
			}

			if downright := area[x+1][y+1]; downright == air {
				x, y = x+1, y+1
				continue
			}

			break
		}

		area[x][y] = sand

		count++

		if x == 500 && y == 0 {
			break
		}
	}

	// printArea(area)
	fmt.Println("count", count)

	return nil
}

func parseInput() [][]string {
	lines := strings.Split(input, "\n")
	coords := make([][][2]int, 0, len(lines))

	xmax, ymax := 0, 0

	for i, line := range lines {
		pairs := strings.Split(line, " -> ")
		coords = append(coords, make([][2]int, 0, len(pairs)))
		for _, pair := range pairs {
			split := strings.Split(pair, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])

			if x > xmax {
				xmax = x
			}

			if y > ymax {
				ymax = y
			}

			coords[i] = append(coords[i], [2]int{x, y})
		}
	}

	area := make([][]string, 0, xmax+1)
	for x := 0; x <= xmax; x++ {
		area = append(area, make([]string, ymax+1))
	}

	for _, path := range coords {
		start := path[0]
		for i := 1; i < len(path); i++ {
			end := path[i]

			if start[0] == end[0] {
				ys := []int{start[1], end[1]}
				if ys[0] > ys[1] {
					ys[0], ys[1] = ys[1], ys[0]
				}

				for y := ys[0]; y <= ys[1]; y++ {
					area[start[0]][y] = rock
				}
			} else {
				xs := []int{start[0], end[0]}
				if xs[0] > xs[1] {
					xs[0], xs[1] = xs[1], xs[0]
				}

				for x := xs[0]; x <= xs[1]; x++ {
					area[x][start[1]] = rock
				}
			}

			start = end
		}
	}

	return area
}

func printArea(area [][]string) {
	for y := 0; y < len(area[0]); y++ {
		for x := 300; x < 800 && x < len(area); x++ {
			if y == 0 && x == 500 && area[500][0] == air {
				fmt.Print("+")
				continue
			}

			if area[x][y] == "" {
				fmt.Print(".")
			} else {
				fmt.Print(area[x][y])
			}
		}
		fmt.Println()
	}
}

type animation struct {
	area [][]string

	image *image.RGBA
}

func (a *animation) Update() error {
	x, y := 500, 0

	for {
		if below := a.area[x][y+1]; below == air {
			y++
			continue
		}

		if downleft := a.area[x-1][y+1]; downleft == air {
			x, y = x-1, y+1
			continue
		}

		if downright := a.area[x+1][y+1]; downright == air {
			x, y = x+1, y+1
			continue
		}

		break
	}

	a.area[x][y] = sand

	for x := 300; x < 700; x++ {
		for y := range a.area[x] {
			r, g, b := 0, 0, 0

			if a.area[x][y] == sand {
				r, g, b = 0xc2, 0xb2, 0x80
			} else if a.area[x][y] == rock {
				r, g, b = 0x67, 0x67, 0x67
			}

			for _, xx := range []int{(x - 300) * 2, (x-300)*2 + 1} {
				for _, yy := range []int{y * 2, y*2 + 1} {
					a.image.Pix[yy*a.image.Stride+xx*4] = uint8(r)
					a.image.Pix[yy*a.image.Stride+xx*4+1] = uint8(g)
					a.image.Pix[yy*a.image.Stride+xx*4+2] = uint8(b)
					a.image.Pix[yy*a.image.Stride+xx*4+3] = 0xff
				}
			}
		}
	}

	if x == 500 && y == 0 {
		a.area = parseInput()
	}

	return nil
}

func (a *animation) Draw(screen *ebiten.Image) {
	screen.WritePixels(a.image.Pix)
}

func (a *animation) Layout(w, h int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(w) * s), int(float64(h) * s)
}

func playAnimation() error {
	area := parseInput()

	// just add in an arbitrarily large floor by extending x to 1000 and adding the 2 new ys
	for x := range area {
		area[x] = append(area[x], air, rock)
	}

	for x := len(area); x < 1000; x++ {
		area = append(area, make([]string, len(area[0])-1))
		area[x] = append(area[x], rock)
	}

	ebiten.SetWindowSize(800, len(area[0])*2)

	a := &animation{
		area:  area,
		image: image.NewRGBA(image.Rect(0, 0, 800, len(area[0])*2)),
	}

	return ebiten.RunGame(a)
}
