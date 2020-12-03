#!/usr/bin/env python3

import sys


def parse_input(rows):
    for x in rows:
        limits, letter, password = x.split()
        lower, upper = limits.split("-")
        yield int(lower), int(upper), letter[0], password


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        input_lines = f.readlines()

    valid_1 = valid_2 = 0
    for lower, upper, letter, password in parse_input(input_lines):
        if lower <= password.count(letter) <= upper:
            valid_1 += 1

        lower_upper = password[lower - 1] + password[upper - 1]
        if lower_upper.count(letter) == 1:
            valid_2 += 1

    print(valid_1)
    print(valid_2)
