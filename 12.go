package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func rotate(v complex128, degrees float64) complex128 {
	r := degrees * math.Pi / 180
	x := math.Cos(r)*real(v) - math.Sin(r)*imag(v)
	y := math.Sin(r)*real(v) + math.Cos(r)*imag(v)
	return complex(x, y)
}

func part1(lines []string) {
	p := 0 + 0i
	v := 1 + 0i
	for _, x := range lines {
		action := x[0]
		value, _ := strconv.ParseFloat(x[1:], 64)
		switch action {
		case 'N':
			p += complex(0, value)
		case 'S':
			p += complex(0, -value)
		case 'E':
			p += complex(value, 0)
		case 'W':
			p += complex(-value, 0)
		case 'F':
			p += v * complex(value, 0)
		case 'L':
			v = rotate(v, value)
		case 'R':
			v = rotate(v, -value)
		}
	}
	fmt.Println(math.Abs(real(p)) + math.Abs(imag(p)))
}

func part2(lines []string) {
	p := 0 + 0i
	v := 10 + 1i
	for _, x := range lines {
		action := x[0]
		value, _ := strconv.ParseFloat(x[1:], 64)
		switch action {
		case 'N':
			v += complex(0, value)
		case 'S':
			v += complex(0, -value)
		case 'E':
			v += complex(value, 0)
		case 'W':
			v += complex(-value, 0)
		case 'F':
			p += v * complex(value, 0)
		case 'L':
			v = rotate(v, value)
		case 'R':
			v = rotate(v, -value)
		}
	}
	fmt.Println(math.Abs(real(p)) + math.Abs(imag(p)))
}

func main() {
	bytes, err := ioutil.ReadFile("12.in")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	part1(lines)
	part2(lines)
}
