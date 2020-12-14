def apply_mask(val, mask):
    val = bin(val).lstrip("0b").zfill(len(mask))
    masked_val = []
    for x, m in zip(val, mask):
        if m == "X":
            masked_val.append(x)
        else:
            masked_val.append(m)
    return int("".join(masked_val), 2)


def apply_mask2(val, mask):
    val = bin(val).lstrip("0b").zfill(len(mask))
    masked_val = []
    floating = []
    for i, (x, m) in enumerate(zip(val, mask)):
        if m == "0":
            masked_val.append(x)
            continue

        if m == "X":
            floating.append(i)

        masked_val.append(m)

    for i in range(2 ** len(floating)):
        x = bin(i).lstrip("0b").zfill(len(floating))
        for i, b in zip(floating, x):
            masked_val[i] = b

        yield int("".join(masked_val), 2)


if __name__ == "__main__":
    with open("14.in") as f:
        input = [l.strip() for l in f.readlines()]

    # Part 1
    mem = {}
    mask = ""
    for l in input:
        op, val = l.split(" = ")
        if op == "mask":
            mask = val
            continue

        addr = int(op.lstrip("mem[").rstrip("]"))
        mem[addr] = apply_mask(int(val), mask)

    print(sum(mem.values()))

    # Part 2
    mem = {}
    mask = ""
    for l in input:
        op, val = l.split(" = ")
        if op == "mask":
            mask = val
            continue

        addr = int(op.lstrip("mem[").rstrip("]"))
        for a in apply_mask2(int(addr), mask):
            mem[a] = int(val)

    print(sum(mem.values()))
