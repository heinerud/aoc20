#!/usr/bin/env python3

import sys


def id(s):
    for x in [("F", "0"), ("B", "1"), ("L", "0"), ("R", "1")]:
        s = s.replace(*x)
    return int(s, 2)


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        lines = [x.strip() for x in f.readlines()]

    ids = [id(x) for x in lines]
    print(max(ids))

    all_seats = range(min(ids), max(ids) + 1)
    print(set(all_seats) - set(ids))
