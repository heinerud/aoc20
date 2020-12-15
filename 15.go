package main

import (
	"fmt"
)

func game(input []int, turns int) int {
	numbers := make(map[int]int, len(input))
	for i, x := range input {
		numbers[x] = i + 1
	}

	speak := 0
	for turn := len(input) + 1; turn < turns; turn++ {
		spoken := speak
		if _, ok := numbers[spoken]; ok {
			speak = turn - numbers[spoken]
		} else {
			speak = 0
		}

		numbers[spoken] = turn
	}

	return speak
}

func main() {
	input := []int{20, 9, 11, 0, 1, 2}
	fmt.Println(game(input, 2020))
	fmt.Println(game(input, 30000000))
}
