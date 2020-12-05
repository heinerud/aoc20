#!/usr/bin/env python3

import sys


def seat(s):
    s = s.replace("F", "0")
    s = s.replace("B", "1")
    s = s.replace("L", "0")
    s = s.replace("R", "1")
    row = s[:7]
    col = s[7:]
    return int(row, 2), int(col, 2)


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        lines = [x.strip() for x in f.readlines()]

    seats = [seat(x) for x in lines]
    ids = [r * 8 + c for r, c in seats]
    print(max(ids))

    all_seats = range(min(ids), max(ids))
    print(set(all_seats) - set(ids))
