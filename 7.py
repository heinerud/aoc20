import re


def contains(rules, bag):
    hits = set()
    for k, v in rules.items():
        if bag in [b for _, b in v]:
            hits.add(k)
            hits |= contains(rules, k)

    return hits


def num_bags(rules, bag):
    total = 0
    for n, inner in rules[bag]:
        total += n
        total += n * num_bags(rules, inner)

    return total


if __name__ == "__main__":
    with open("7.in") as f:
        lines = [x.strip() for x in f.readlines()]

    rules = {}
    for l in lines:
        bag, contents = l.split(" bags contain ")
        contents = re.findall("(\d+) (\S+ \S+)", contents)
        contents = [(int(n), b) for n, b in contents]
        rules[bag] = contents

    print(len(contains(rules, "shiny gold")))
    print(num_bags(rules, "shiny gold"))
