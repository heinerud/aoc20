package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func seat(s string) (row int, col int) {
	s = strings.ReplaceAll(s, "F", "0")
	s = strings.ReplaceAll(s, "B", "1")
	s = strings.ReplaceAll(s, "L", "0")
	s = strings.ReplaceAll(s, "R", "1")
	row64, _ := strconv.ParseInt(s[:7], 2, 0)
	col64, _ := strconv.ParseInt(s[7:], 2, 0)
	return int(row64), int(col64)
}

func main() {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	ids := make(map[int]int, len(lines))
	var min, max int
	min = math.MaxInt32
	for _, x := range lines {
		row, col := seat(x)
		id := row*8 + col
		ids[id] = id

		if id > max {
			max = id
		}

		if id < min {
			min = id
		}
	}
	fmt.Println(max)

	for i := min; i <= max; i++ {
		_, seen := ids[i]
		if !seen {
			fmt.Println(i)
			break
		}
	}
}
