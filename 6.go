package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(bytes), "\n")

	var entries [][]string
	var entry []string
	for _, l := range lines {
		if l == "" {
			entries = append(entries, entry)
			entry = []string{}
		} else {
			entry = append(entry, strings.Fields(l)...)
		}
	}

	if len(entry) > 0 {
		entries = append(entries, entry)
	}

	var sumAnyone int
	for _, e := range entries {
		anyone := make(map[rune]int)
		for _, c := range strings.Join(e, "") {
			anyone[c]++
		}
		sumAnyone += len(anyone)
	}
	fmt.Println(sumAnyone)

	var sumEveryone int
	for _, e := range entries {
		for _, c := range e[0] {
			ok := true
			for _, g := range e {
				if !strings.Contains(g, string(c)) {
					ok = false
				}
			}
			if ok {
				sumEveryone++
			}
		}
	}
	fmt.Println(sumEveryone)
}
