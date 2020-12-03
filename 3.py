#!/usr/bin/env python3

import sys


def traverse(slope, v_x, v_y):
    x = y = hits = 0
    while y < len(slope):
        row = slope[y]
        if row[x % len(row)] == "#":
            hits += 1
        x += v_x
        y += v_y

    return hits


if __name__ == "__main__":
    slope = [x.strip() for x in sys.stdin]
    print(traverse(slope, 3, 1))
    print(
        traverse(slope, 1, 1)
        * traverse(slope, 3, 1)
        * traverse(slope, 5, 1)
        * traverse(slope, 7, 1)
        * traverse(slope, 1, 2)
    )
