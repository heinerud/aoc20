import re


def eval_plus(x):
    while m := re.search("\d+ \+ \d+", x):
        x = x.replace(m.group(), str(eval(m.group())), 1)

    return x


def eval_all(x):
    while m := re.search("\d+ [\+\-\*\/] \d+", x):
        x = x.replace(m.group(), str(eval(m.group())), 1)

    return x


def eval_expr(x, evaluate):
    while m := re.search("\([^\(\)]+\)", x):
        m = m.group(0)
        x = x.replace(m, eval_expr(m[1:-1], evaluate), 1)

    return evaluate(x)


def part1(input):
    total = 0
    for x in input:
        x = eval_expr(x, lambda x: eval_all(x))
        total += int(x)
    return total


def part2(input):
    total = 0
    for x in input:
        x = eval_expr(x, lambda x: eval_all(eval_plus(x)))
        total += int(x)
    return total


def main():
    with open("18.in") as f:
        input = [l.strip() for l in f.readlines()]

    print(part1(input))
    print(part2(input))


if __name__ == "__main__":
    main()
