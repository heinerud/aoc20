package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Content struct {
	amount int
	color  string
}

func contains(rules map[string][]Content, bag string) (hits map[string]string) {
	hits = make(map[string]string)
	for k, v := range rules {
		for _, x := range v {
			if bag == x.color {
				hits[k] = k
				for k, _ := range contains(rules, k) {
					hits[k] = k
				}
			}
		}
	}

	return hits
}

func numBags(rules map[string][]Content, bag string) (sum int) {
	for _, b := range rules[bag] {
		sum += b.amount
		sum += b.amount * numBags(rules, b.color)
	}
	return sum
}

func main() {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	rules := make(map[string][]Content, len(lines))
	re := regexp.MustCompile(`(\d+) (\S+ \S+)`)
	for _, x := range lines {
		split := strings.Split(x, " bags contain ")
		matches := re.FindAllStringSubmatch(split[1], -1)
		contents := make([]Content, len(matches))
		for i, x := range matches {
			n, _ := strconv.Atoi(x[1])
			contents[i] = Content{n, x[2]}
		}
		rules[split[0]] = contents
	}

	fmt.Println(len(contains(rules, "shiny gold")))
	fmt.Println(numBags(rules, "shiny gold"))
}
