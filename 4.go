package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func properties(lines []string) []map[string]string {
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

	properties := make([]map[string]string, len(entries))
	for i, props := range entries {
		propMap := make(map[string]string, len(props))
		for _, x := range props {
			split := strings.Split(x, ":")
			propMap[split[0]] = split[1]
		}
		properties[i] = propMap
	}

	return properties
}

func check1(props map[string]string) bool {
	_, byr := props["byr"]
	_, ecl := props["ecl"]
	_, eyr := props["eyr"]
	_, hcl := props["hcl"]
	_, hgt := props["hgt"]
	_, iyr := props["iyr"]
	_, pid := props["pid"]

	return byr && ecl && eyr && hcl && hgt && iyr && pid
}

func check2(props map[string]string) bool {
	if !check1(props) {
		return false
	}

	byr, _ := strconv.Atoi(props["byr"])
	if len(props["byr"]) != 4 || byr < 1920 || byr > 2002 {
		return false
	}

	iyr, _ := strconv.Atoi(props["iyr"])
	if len(props["iyr"]) != 4 || iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr, _ := strconv.Atoi(props["eyr"])
	if len(props["eyr"]) != 4 || eyr < 2020 || eyr > 2030 {
		return false
	}

	replacer := strings.NewReplacer("cm", "", "in", "")
	hgt, _ := strconv.Atoi(replacer.Replace(props["hgt"]))
	if strings.HasSuffix(props["hgt"], "in") {
		if hgt < 59 || hgt > 76 {
			return false
		}
	} else if strings.HasSuffix(props["hgt"], "cm") {
		if hgt < 150 || hgt > 193 {
			return false
		}
	} else {
		return false
	}

	const hairClrPattern = "^#[A-Za-z0-9]{6}$"
	if ok, _ := regexp.MatchString(hairClrPattern, props["hcl"]); !ok {
		return false
	}

	const eyeClrPattern = "^(amb|blu|brn|gry|grn|hzl|oth)$"
	if ok, _ := regexp.MatchString(eyeClrPattern, props["ecl"]); !ok {
		return false
	}

	const countryCodePattern = "^[0-9]{9}$"
	if ok, _ := regexp.MatchString(countryCodePattern, props["pid"]); !ok {
		return false
	}

	return true
}

func main() {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(bytes), "\n")
	entries := properties(lines)

	var validPassports1, validPassports2 int
	for _, x := range entries {
		if check1(x) {
			validPassports1++
		}
		if check2(x) {
			validPassports2++
		}
	}

	fmt.Println(validPassports1)
	fmt.Println(validPassports2)
}
