import itertools


def part1(input):
    rule_lines = []
    messages = []
    for x in input:
        if ":" in x:
            rule_lines.append(x)
        elif x:
            messages.append(x)

    rules = {}
    for l in rule_lines:
        if '"' in l:
            i, r = l.split(":")
            i = int(i)
            r = r.strip()
            r = r.replace('"', "")
            rules[i] = {r}

    while len(rules) < len(rule_lines):
        for l in rule_lines:
            n, r = l.split(":")
            n = int(n)

            if n in rules:
                continue

            r = r.strip()

            if not all(int(y) in rules for y in r.split() if y.isdigit()):
                continue

            rsplit = r.split("|")
            all_options = set()
            for s in rsplit:
                options = {""}
                for num in [int(x) for x in s.split()]:
                    sub_options = set()
                    for p in itertools.product(options, rules[num]):
                        sub_options.add("".join(p))
                    options = sub_options
                all_options |= options

            rules[n] = all_options

    matches = 0
    for m in messages:
        if m in rules[0]:
            matches += 1

    print(matches)


def part2(input):
    rule_lines = []
    messages = []
    rule_8 = "42 | 42 8"
    rule_11 = "42 31 | 42 11 31"
    for x in input:
        if ":" in x:
            if x.startswith("8:"):
                x = f"8: {rule_8}"
            elif x.startswith("11:"):
                x = f"11: {rule_11}"
            elif x.startswith("0:"):
                rule_0 = x
            rule_lines.append(x)
        elif x:
            messages.append(x)

    rules = {}
    for l in rule_lines:
        if '"' in l:
            i, r = l.split(":")
            i = int(i)
            r = r.strip()
            r = r.replace('"', "")
            rules[i] = {r}

    deadlock = True
    while not deadlock:
        for l in rule_lines:
            n, r = l.split(":")
            n = int(n)

            if n in rules:
                continue

            r = r.strip()

            if not all(int(y) in rules for y in r.split() if y.isdigit()):
                continue

            deadlock = False
            rsplit = r.split("|")
            all_options = set()
            for s in rsplit:
                options = {""}
                for num in [int(x) for x in s.split()]:
                    sub_options = set()
                    for p in itertools.product(options, rules[num]):
                        sub_options.add("".join(p))
                    options = sub_options
                all_options |= options

            rules[n] = all_options

    print(rule_8)
    print(rule_11)
    print(rule_0)

    matches = 0
    for m in messages:
        matches += 1

    print(matches)


def main():
    with open("19.in") as f:
        input = [l.strip() for l in f.readlines()]

    part1(input)
    part2(input)


if __name__ == "__main__":
    main()
