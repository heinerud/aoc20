package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func parseInput(line string) (lower int, upper int, letter string, password string) {
	splits := strings.Split(line, " ")
	limits := strings.Split(splits[0], "-")
	lower, err := strconv.Atoi(limits[0])
	check(err)
	upper, err = strconv.Atoi(limits[1])
	check(err)
	letter = string(splits[1][0])
	password = splits[2]
	return lower, upper, letter, password
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var valid1, valid2 int
	for scanner.Scan() {
		lower, upper, letter, password := parseInput(scanner.Text())
		matches := strings.Count(password, letter)
		if lower <= matches && matches <= upper {
			valid1 += 1
		}

		lowerUpper := string(password[lower-1]) + string(password[upper-1])
		if strings.Count(lowerUpper, letter) == 1 {
			valid2 += 1
		}
	}

	fmt.Println(valid1)
	fmt.Println(valid2)
}
