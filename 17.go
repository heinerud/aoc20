package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type cube3 struct {
	x int
	y int
	z int
}

func neighbors3(c cube3) []cube3 {
	neighs := []cube3{}
	for dz := -1; dz < 2; dz++ {
		for dy := -1; dy < 2; dy++ {
			for dx := -1; dx < 2; dx++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				neighs = append(neighs, cube3{c.x + dx, c.y + dy, c.z + dz})
			}
		}
	}
	return neighs
}

func iterate3(world map[cube3]bool) {
	inactiveNeighborsToActive := make(map[cube3]int)
	deactivate := make(map[cube3]bool)
	for x, _ := range world {
		activeNeighbors := 0
		for _, n := range neighbors3(x) {
			if _, ok := world[n]; ok {
				activeNeighbors++
			} else {
				inactiveNeighborsToActive[n]++
			}
		}
		if activeNeighbors != 2 && activeNeighbors != 3 {
			deactivate[x] = true
		}
	}
	for k, _ := range deactivate {
		delete(world, k)
	}
	for k, v := range inactiveNeighborsToActive {
		if v == 3 {
			world[k] = true
		}
	}
}

type cube4 struct {
	x int
	y int
	z int
	w int
}

func neighbors4(c cube4) []cube4 {
	neighs := []cube4{}
	for dw := -1; dw < 2; dw++ {
		for dz := -1; dz < 2; dz++ {
			for dy := -1; dy < 2; dy++ {
				for dx := -1; dx < 2; dx++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					neighs = append(neighs, cube4{c.x + dx, c.y + dy, c.z + dz, c.w + dw})
				}
			}
		}
	}
	return neighs
}

func iterate4(world map[cube4]bool) {
	inactiveNeighborsToActive := make(map[cube4]int)
	deactivate := make(map[cube4]bool)
	for x, _ := range world {
		activeNeighbors := 0
		for _, n := range neighbors4(x) {
			if _, ok := world[n]; ok {
				activeNeighbors++
			} else {
				inactiveNeighborsToActive[n]++
			}
		}
		if activeNeighbors != 2 && activeNeighbors != 3 {
			deactivate[x] = true
		}
	}
	for k, _ := range deactivate {
		delete(world, k)
	}
	for k, v := range inactiveNeighborsToActive {
		if v == 3 {
			world[k] = true
		}
	}
}

func part1(input []string) int {
	cubes := make(map[cube3]bool)
	for y, row := range input {
		for x, c := range row {
			if c == '#' {
				cubes[cube3{x, -y, 0}] = true
			}
		}
	}

	for i := 0; i < 6; i++ {
		iterate3(cubes)
	}
	return len(cubes)
}

func part2(input []string) int {
	cubes := make(map[cube4]bool)
	for y, row := range input {
		for x, c := range row {
			if c == '#' {
				cubes[cube4{x, -y, 0, 0}] = true
			}
		}
	}

	for i := 0; i < 6; i++ {
		iterate4(cubes)
	}
	return len(cubes)
}

func main() {
	bytes, err := ioutil.ReadFile("17.in")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
