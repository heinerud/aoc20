#!/usr/bin/env python3

import sys


def run(program):
    acc = 0
    p = 0
    visited = set()
    while p < len(program):
        if p in visited:
            return acc, False

        visited.add(p)
        instruction, n = program[p]
        if instruction == "acc":
            acc += n
            p += 1
        elif instruction == "jmp":
            p += n
        else:
            p += 1

    return acc, True


def swap(x):
    if x == "jmp":
        return "nop"
    elif x == "nop":
        return "jmp"
    else:
        raise ValueError("Illegal instruction", x)


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        lines = [x.strip() for x in f.readlines()]

    program = [(x, int(n)) for x, n in [l.split() for l in lines]]

    print(run(program))

    for i in range(len(program)):
        instruction, n = program[i]
        try:
            program[i] = (swap(instruction), n)
        except ValueError:
            continue

        acc, ok = run(program)
        if ok:
            print(acc, "error in row", i + 1)
            break

        program[i] = (instruction, n)
