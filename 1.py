from functools import reduce
from itertools import combinations
from operator import mul


def sum2020(entries, num):
    for c in combinations(entries, num):
        if sum(c) == 2020:
            return c


if __name__ == "__main__":
    with open("1.in") as f:
        entries = [int(x) for x in f.readlines()]

    sum_2 = sum2020(entries, 2)
    print(sum_2, reduce(mul, sum_2))

    sum_3 = sum2020(entries, 3)
    print(sum_3, reduce(mul, sum_3))
