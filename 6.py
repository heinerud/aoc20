#!/usr/bin/env python3

import sys


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        lines = [x.strip() for x in f.readlines()]

    entries = []
    entry = []
    for l in lines:
        if not l:
            entries.append(entry)
            entry = []
        else:
            entry.append(l)

    if entry:
        entries.append(entry)

    sum_anyone = 0
    for e in entries:
        answers = "".join(e)
        anyone = set(answers)
        sum_anyone += len(anyone)
    print(sum_anyone)

    sum_everyone = 0
    for e in entries:
        sets = [set(x) for x in e]
        everyone = set.intersection(*sets)
        sum_everyone += len(everyone)
    print(sum_everyone)
