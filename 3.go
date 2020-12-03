package main

import (
	"bufio"
	"fmt"
	"os"
)

func traverse(slope []string, vx int, vy int) int {
	var x, y, hits int
	for y < len(slope) {
		row := slope[y]
		if row[x%len(row)] == '#' {
			hits += 1
		}
		x += vx
		y += vy
	}
	return hits
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	slope := []string{}
	for scanner.Scan() {
		slope = append(slope, scanner.Text())
	}

	fmt.Println(traverse(slope, 3, 1))

	res_2 := traverse(slope, 1, 1)
	res_2 *= traverse(slope, 3, 1)
	res_2 *= traverse(slope, 5, 1)
	res_2 *= traverse(slope, 7, 1)
	res_2 *= traverse(slope, 1, 2)
	fmt.Println(res_2)
}
