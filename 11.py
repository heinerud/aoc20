def neighbors(
    r,
    c,
    *,
    rmin=float("-inf"),
    rmax=float("inf"),
    cmin=float("-inf"),
    cmax=float("inf"),
):
    west = c > cmin
    east = c < cmax
    north = r > rmin
    south = r < rmax

    if west:
        yield (r, c - 1)
    if east:
        yield (r, c + 1)
    if north:
        yield (r - 1, c)
    if south:
        yield (r + 1, c)
    if north and west:
        yield (r - 1, c - 1)
    if north and east:
        yield (r - 1, c + 1)
    if south and west:
        yield (r + 1, c - 1)
    if south and east:
        yield (r + 1, c + 1)


if __name__ == "__main__":
    with open("11.in") as f:
        input = [x.strip() for x in f.readlines()]

    # Part 1
    lines = [list(x) for x in input]
    height = len(lines)
    width = len(lines[0])
    while True:
        no_changes = True
        new_lines = [list(line) for line in lines]

        for r in range(height):
            for c in range(width):
                if lines[r][c] == ".":
                    continue

                occupied = 0
                for dr, dc in neighbors(
                    r, c, rmin=0, rmax=height - 1, cmin=0, cmax=width - 1
                ):
                    if lines[dr][dc] == "#":
                        occupied += 1

                if lines[r][c] == "L" and occupied == 0:
                    new_lines[r][c] = "#"
                    no_changes = False
                if lines[r][c] == "#" and occupied > 3:
                    new_lines[r][c] = "L"
                    no_changes = False

        if no_changes:
            break

        lines = new_lines

    print(sum(sum(x == "#" for x in row) for row in lines))

    # Part 2
    lines = [list(x) for x in input]
    while True:
        no_changes = True
        new_lines = [list(l) for l in lines]

        for r in range(height):
            for c in range(width):
                if lines[r][c] == ".":
                    continue

                occupied = 0
                for vr, vc in neighbors(0, 0):
                    dr = vr
                    dc = vc

                    while 0 <= r + dr < height and 0 <= c + dc < width:
                        if lines[r + dr][c + dc] in "L#":
                            if lines[r + dr][c + dc] == "#":
                                occupied += 1
                            break

                        dr += vr
                        dc += vc

                if lines[r][c] == "L" and occupied == 0:
                    new_lines[r][c] = "#"
                    no_changes = False
                if lines[r][c] == "#" and occupied > 4:
                    new_lines[r][c] = "L"
                    no_changes = False

        if no_changes:
            break

        lines = new_lines

    print(sum(sum(c == "#" for c in row) for row in lines))
