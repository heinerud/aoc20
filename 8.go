package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Op struct {
	instruction string
	value       int
}

func run(program []Op) (acc int, e error) {
	var p int
	visited := make(map[int]int)
	for p < len(program) {
		if _, visited := visited[p]; visited {
			return acc, errors.New("error")
		}
		visited[p] = p

		switch op := program[p]; op.instruction {
		case "acc":
			acc += op.value
			p++
		case "jmp":
			p += op.value
		default:
			p++
		}
	}

	return acc, e
}

func main() {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	program := make([]Op, len(lines))
	for i, l := range lines {
		split := strings.Fields(l)
		value, _ := strconv.Atoi(split[1])
		program[i] = Op{split[0], value}
	}

	fmt.Println(run(program))

	swap := map[string]string{"jmp": "nop", "nop": "jmp"}
	for i := 0; i < len(program); i++ {
		op := program[i]
		if swapOp, ok := swap[op.instruction]; ok {
			program[i].instruction = swapOp
		} else {
			continue
		}

		if acc, e := run(program); e == nil {
			fmt.Println(acc, "error in row ", i+1)
			break
		}

		program[i].instruction = op.instruction
	}
}
