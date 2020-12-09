package main

import (
	"./itertools"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func sumInPreamble(preamble []int, sum int) bool {
	c := make(chan []int)
	go itertools.Combinations(preamble, 2, c)
	for x := range c {
		if x[0]+x[1] == sum {
			return true
		}
	}
	return false
}

func main() {
	bytes, err := ioutil.ReadFile("9.in")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	input := make([]int, len(lines))
	for i, x := range lines {
		xi, _ := strconv.Atoi(x)
		input[i] = xi
	}

	// Part 1
	var faulty int
	for i := 25; i < len(input); i++ {
		preamble := input[i-25 : i]
		if !sumInPreamble(preamble, input[i]) {
			faulty = input[i]
		}
	}
	fmt.Println(faulty)

	// Part 2
	var min, max, sum int
	for i := 0; i < len(input); i++ {
		sum = 0
		max = 0
		min = math.MaxInt32
		for _, x := range input[i:] {
			if x < min {
				min = x
			}
			if x > max {
				max = x
			}
			sum += x
			if sum >= faulty {
				break
			}
		}
		if sum == faulty {
			break
		}
	}
	fmt.Println(min + max)
}
