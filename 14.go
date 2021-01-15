package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func applyMask1(val int, mask string) int {
	binary := strconv.FormatInt(int64(val), 2)
	padded := fmt.Sprintf("%0*s", len(mask), binary)
	masked := make([]byte, len(padded))
	for i := 0; i < len(masked); i++ {
		if mask[i] == 'X' {
			masked[i] = padded[i]
		} else {
			masked[i] = mask[i]
		}
	}
	result, _ := strconv.ParseInt(string(masked), 2, 64)
	return int(result)
}

func applyMask2(val int, mask string) []int {
	binary := strconv.FormatInt(int64(val), 2)
	padded := fmt.Sprintf("%0*s", len(mask), binary)
	masked := make([]byte, len(padded))
	floating := []int{}
	for i := 0; i < len(masked); i++ {
		if mask[i] == '0' {
			masked[i] = padded[i]
		} else if mask[i] == 'X' {
			floating = append(floating, i)
		} else {
			masked[i] = mask[i]
		}
	}

	addrs := []int{}
	for i := 0; i < int(math.Pow(2, float64(len(floating)))); i++ {
		binary := strconv.FormatInt(int64(i), 2)
		padded := fmt.Sprintf("%0*s", len(floating), binary)
		for i, _ := range padded {
			masked[floating[i]] = padded[i]
		}
		result, _ := strconv.ParseInt(string(masked), 2, 64)
		addrs = append(addrs, int(result))
	}
	return addrs
}

func part1(lines []string) int {
	maskRe := regexp.MustCompile(`^mask = ([01X]+)$`)
	memRe := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	mem := make(map[int]int)
	var mask string
	for _, l := range lines {
		if maskMatch := maskRe.FindStringSubmatch(l); len(maskMatch) > 0 {
			mask = maskMatch[1]
			continue
		}

		if memMatch := memRe.FindStringSubmatch(l); len(memMatch) > 0 {
			addr, _ := strconv.Atoi(memMatch[1])
			val, _ := strconv.Atoi(memMatch[2])
			mem[addr] = applyMask1(val, mask)
		}
	}

	var sum int
	for _, v := range mem {
		sum += v
	}
	return sum
}

func part2(lines []string) int {
	maskRe := regexp.MustCompile(`^mask = ([01X]+)$`)
	memRe := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	mem := make(map[int]int)
	var mask string
	for _, l := range lines {
		if maskMatch := maskRe.FindStringSubmatch(l); len(maskMatch) > 0 {
			mask = maskMatch[1]
			continue
		}

		if memMatch := memRe.FindStringSubmatch(l); len(memMatch) > 0 {
			addr, _ := strconv.Atoi(memMatch[1])
			val, _ := strconv.Atoi(memMatch[2])
			for _, a := range applyMask2(addr, mask) {
				mem[a] = val
			}
		}
	}

	var sum int
	for _, v := range mem {
		sum += v
	}
	return sum
}

func main() {
	bytes, err := ioutil.ReadFile("14.in")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
