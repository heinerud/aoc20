if __name__ == "__main__":
    with open("16.in") as f:
        input = [l.strip() for l in f.readlines()]

    your_index = input.index("your ticket:")
    nearby_index = input.index("nearby tickets:")

    rules = {}
    for l in input[:your_index]:
        if not l:
            continue
        k, v = l.split(":")
        rules[k] = set()
        for x in v.split("or"):
            lower, upper = x.split("-")
            for x in range(int(lower), int(upper) + 1):
                rules[k].add(x)

    your_ticket = [int(x) for x in input[your_index + 1].split(",")]

    nearby_tickets = []
    for l in input[nearby_index + 1 :]:
        nearby_tickets.append([int(x) for x in l.split(",")])

    # Part 1
    valid_values = set()
    for v in rules.values():
        valid_values |= set(v)

    scanning_error_rate = 0
    valid_nearby_tickets = []
    for t in nearby_tickets:
        valid = True
        for v in t:
            if v not in valid_values:
                valid = False
                scanning_error_rate += v

        if valid:
            valid_nearby_tickets.append(t)

    print(scanning_error_rate)

    # Part 2
    cols = []
    for i in range(len(your_ticket)):
        col = [x[i] for x in valid_nearby_tickets]
        matches = set()
        for k, v in rules.items():
            if all(c in v for c in col):
                matches.add(k)

        cols.append(matches)

    known = {}
    while any(cols):
        for i, c in enumerate(cols):
            if len(c) != 1:
                continue
            known[i] = next(iter(c))
            for cc in cols:
                cc.discard(known[i])

    prod_departure = 1
    for i, x in enumerate(your_ticket):
        if "departure" in known[i]:
            prod_departure *= x

    print(prod_departure)
