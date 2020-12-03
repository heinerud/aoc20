package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	entries := []int{}
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		entries = append(entries, number)
	}

	var found1, found2 bool
	for _, x := range entries {
		for _, y := range entries {
			if !found1 && x+y == 2020 {
				fmt.Println("Task 1:", x, y, x*y)
				found1 = true
			}
			if !found2 {
				for _, z := range entries {
					if x+y+z == 2020 {
						fmt.Println("Task 2:", x, y, z, x*y*z)
						found2 = true
					}
				}
			}
		}
	}
}
