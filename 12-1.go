package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Ship struct {
	pos complex128
	dir complex128
}

func (s *Ship) move(heading complex128, distance int) {
	s.pos += heading * complex(float64(distance), 0)
}

func (s *Ship) rotate(degrees int) {
	r := float64(degrees) * math.Pi / 180
	x := math.Cos(r)*real(s.dir) - math.Sin(r)*imag(s.dir)
	y := math.Sin(r)*real(s.dir) + math.Cos(r)*imag(s.dir)
	s.dir = complex(x, y)
}

func main() {
	bytes, err := ioutil.ReadFile("12.in")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	s := Ship{dir: 1 + 0i}
	for _, x := range lines {
		action := x[0]
		value, _ := strconv.Atoi(x[1:])
		switch action {
		case 'N':
			s.move(1i, value)
		case 'S':
			s.move(-1i, value)
		case 'E':
			s.move(1+0i, value)
		case 'W':
			s.move(-1+0i, value)
		case 'F':
			s.move(s.dir, value)
		case 'L':
			s.rotate(value)
		case 'R':
			s.rotate(-value)
		}
	}

	fmt.Println(math.Abs(real(s.pos)) + math.Abs(imag(s.pos)))
}
