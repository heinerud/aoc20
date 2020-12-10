from itertools import combinations
import math


def num_combinations(n, r=2):
    if n < r:
        return 0

    numerator = math.factorial(n)
    denominator = math.factorial(r) * math.factorial(n - r)
    return int(numerator / denominator)


if __name__ == "__main__":
    with open("10.in") as f:
        lines = [x.strip() for x in f.readlines()]

    input = [0]
    input.extend(sorted([int(x) for x in lines]))
    input = sorted(input)
    input.append(input[-1] + 3)

    # Part 1
    jumps = [b - a for a, b in zip(input[:-1], input[1:])]
    print(jumps.count(1) * jumps.count(3))

    # Part 2
    holes = []
    hole = 0
    for x in jumps:
        if x == 1:
            hole += 1
        elif hole:
            holes.append(hole)
            hole = 0

    arrangements = 1
    for x in holes:
        arrangements *= num_combinations(x) + 1

    print(arrangements)
