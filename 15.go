package main

import (
	"fmt"
)

type Memory struct {
	first int
	last  int
}

func (m Memory) age() int {
	return m.last - m.first
}

func (m Memory) add(turn int) Memory {
	return Memory{m.last, turn}
}

func main() {
	input := []int{20, 9, 11, 0, 1, 2}
	numbers := make(map[int]Memory, len(input))

	for i, x := range input {
		numbers[x] = Memory{i + 1, i + 1}
	}

	s := input[len(input)-1]
	for i := len(input) + 1; i <= 30000000; i++ {
		s = numbers[s].age()
		if _, ok := numbers[s]; !ok {
			numbers[s] = Memory{i, i}
		} else {
			numbers[s] = numbers[s].add(i)
		}

		if i == 2020 {
			fmt.Println(s)
		}
	}
	fmt.Println(s)

}
