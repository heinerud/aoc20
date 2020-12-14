package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func nextDeparture(time int, buses []int) (int, int) {
	var bus int
	tMin := math.MaxInt64
	for _, b := range buses {
		t := b * (time/b + 1)
		if t < tMin {
			tMin = t
			bus = b
		}
	}

	return bus, tMin
}

type Entry struct {
	offset int
	bus    int
}

func matchSeries(series []Entry) int {
	sort.Slice(series, func(i, j int) bool {
		return series[i].offset > series[j].offset
	})
	t := 0
	step := 1
	for _, x := range series {
		for (t+x.offset)%x.bus != 0 {
			t += step
		}
		step *= x.bus
	}
	return t
}

func main() {
	bytes, err := ioutil.ReadFile("13.in")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	// Part 1
	time, _ := strconv.Atoi(lines[0])
	buses := []int{}
	for _, x := range strings.Split(lines[1], ",") {
		if i, e := strconv.Atoi(x); e == nil {
			buses = append(buses, i)
		}
	}
	bus, departure := nextDeparture(time, buses)
	fmt.Println((departure - time) * bus)

	// Part 2
	series := []Entry{}
	for offset, s := range strings.Split(lines[1], ",") {
		if x, e := strconv.Atoi(s); e == nil {
			series = append(series, Entry{offset, x})
		}
	}
	fmt.Println(matchSeries(series))
}
