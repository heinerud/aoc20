package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

type Cell struct {
	pos      Point
	seat     bool
	occupied bool
}

func (c Cell) neighbors() []Point {
	n := Point{c.pos.x, c.pos.y - 1}
	e := Point{c.pos.x + 1, c.pos.y}
	s := Point{c.pos.x, c.pos.y + 1}
	w := Point{c.pos.x - 1, c.pos.y}
	ne := Point{c.pos.x + 1, c.pos.y - 1}
	se := Point{c.pos.x + 1, c.pos.y + 1}
	sw := Point{c.pos.x - 1, c.pos.y + 1}
	nw := Point{c.pos.x - 1, c.pos.y - 1}
	return []Point{n, e, s, w, ne, se, sw, nw}
}

type Grid struct {
	grid   [][]Cell
	width  int
	height int
}

func (g Grid) numOccupied() int {
	occupied := 0
	for _, row := range g.grid {
		for _, x := range row {
			if x.occupied {
				occupied++
			}
		}
	}
	return occupied
}

func (g Grid) changes1() []Cell {
	changes := []Cell{}
	for _, row := range g.grid {
		for _, c := range row {
			if !c.seat {
				continue
			}

			occupied := 0
			for _, n := range c.neighbors() {
				if n.x < 0 || n.x >= g.width || n.y < 0 || n.y >= g.height {
					continue
				}

				if g.grid[n.y][n.x].occupied {
					occupied++
				}
			}
			if !c.occupied && occupied == 0 {
				changes = append(changes, Cell{c.pos, c.seat, true})
			}
			if c.occupied && occupied > 3 {
				changes = append(changes, Cell{c.pos, c.seat, false})
			}
		}
	}
	return changes
}

func (g Grid) rayNeighbor(c Cell, vx, vy int) int {
	n := Point{c.pos.x + vx, c.pos.y + vy}
	if n.x < 0 || n.y < 0 || n.x >= g.width || n.y >= g.height {
		return 0
	}
	next := g.grid[n.y][n.x]
	if !next.seat {
		return g.rayNeighbor(next, vx, vy)
	}
	if !next.occupied {
		return 0
	}
	return 1
}

func (g Grid) neighbors2(c Cell) int {
	occupied := 0
	for _, v := range g.grid[0][0].neighbors() {
		occupied += g.rayNeighbor(c, v.x, v.y)
	}
	return occupied
}

func (g Grid) changes2() []Cell {
	changes := []Cell{}
	for _, row := range g.grid {
		for _, c := range row {
			if !c.seat {
				continue
			}

			occupied := g.neighbors2(c)
			if !c.occupied && occupied == 0 {
				changes = append(changes, Cell{c.pos, c.seat, true})
			}
			if c.occupied && occupied > 4 {
				changes = append(changes, Cell{c.pos, c.seat, false})
			}
		}
	}
	return changes
}

func parse(lines []string) Grid {
	height := len(lines)
	width := len(lines[0])
	rows := make([][]Cell, height)
	for y := 0; y < len(lines); y++ {
		row := make([]Cell, width)
		for x, c := range lines[y] {
			row[x] = Cell{seat: c == 'L', pos: Point{x, y}}
		}
		rows[y] = row
	}
	return Grid{rows, width, height}
}

func main() {
	bytes, err := ioutil.ReadFile("11.in")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	// Part 1
	grid := parse(lines)
	for true {
		changes := grid.changes1()
		if len(changes) == 0 {
			fmt.Println(grid.numOccupied())
			break
		}
		for _, c := range changes {
			grid.grid[c.pos.y][c.pos.x] = c
		}
	}

	// Part2
	grid2 := parse(lines)
	for true {
		changes := grid2.changes2()
		if len(changes) == 0 {
			fmt.Println(grid2.numOccupied())
			break
		}
		for _, c := range changes {
			grid2.grid[c.pos.y][c.pos.x] = c
		}
	}

}
