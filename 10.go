package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	result := 1
	for i := n; i > 1; i-- {
		result *= i
	}
	return result
}

func numCombinations(n, r int) int {
	if n < r {
		return 0
	}

	numerator := factorial(n)
	denominator := factorial(r) * factorial(n-r)
	return numerator / denominator
}

func main() {
	bytes, err := ioutil.ReadFile("10.in")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	input := []int{0}
	for _, x := range lines {
		xi, _ := strconv.Atoi(x)
		input = append(input, xi)
	}
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	// Part 1
	jumps := []int{}
	var ones, threes int
	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]
		if diff == 1 {
			ones++
		}
		if diff == 3 {
			threes++
		}
		jumps = append(jumps, diff)
	}
	fmt.Println(ones * threes)

	// Part 2
	holes := []int{}
	hole := 0
	for _, x := range jumps {
		if x == 1 {
			hole++
		} else if hole != 0 {
			holes = append(holes, hole)
			hole = 0
		}
	}

	arrangements := 1
	for _, x := range holes {
		arrangements *= numCombinations(x, 2) + 1
	}
	fmt.Println(arrangements)
}
