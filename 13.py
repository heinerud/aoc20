import math
import itertools


def next_departure(time, buses):
    bus = None
    t_min = float("inf")
    for b in buses:
        t = b * math.ceil(time / b)
        if t < t_min:
            t_min = t
            bus = b

    return bus, t_min


def match_series(series):
    series.sort(key=lambda x: x[0], reverse=True)
    t = 0
    step = 1
    for offset, bus in series:
        while (t + offset) % bus:
            t += step
        step *= bus

    return t


if __name__ == "__main__":
    with open("13.in") as f:
        input = [x.strip() for x in f.readlines()]

    # Part 1
    time = int(input[0])
    buses = []
    for x in input[1].split(","):
        if x == "x":
            continue
        buses.append(int(x))

    bus, departure = next_departure(time, buses)
    print((departure - time) * bus)

    # Part 2
    series = []
    for i, x in enumerate(input[1].split(",")):
        if x == "x":
            continue
        else:
            x = int(x)
        series.append((i, x))

    print(match_series(series))
