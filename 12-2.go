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

func main() {
	bytes, err := ioutil.ReadFile("12.in")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	ship := 0 + 0i
	waypoint := 10 + 1i
	for _, x := range lines {
		action := x[0]
		value, _ := strconv.ParseFloat(x[1:], 64)
		v := complex(value, 0)
		switch action {
		case 'N':
			waypoint += v * (0 + 1i)
		case 'S':
			waypoint += v * (0 - 1i)
		case 'E':
			waypoint += v * (1 + 0i)
		case 'W':
			waypoint += v * (-1 + 0i)
		case 'F':
			ship += v * waypoint
		case 'L':
			waypoint = rotate(waypoint, value)
		case 'R':
			waypoint = rotate(waypoint, -value)
		}
	}

	fmt.Println(math.Abs(real(ship)) + math.Abs(imag(ship)))
}
