import itertools


def neighbors(*dimensions):
    for diff in itertools.product((-1, 0, 1), repeat=len(dimensions)):
        if all(x == 0 for x in diff):
            continue
        yield tuple(dim + d for dim, d in zip(dimensions, diff))


def iterate(world):
    inactive_neighs_to_active = {}
    deactivate = set()
    for x in world:
        active_neighbors = 0
        inactive_neighbors = set()
        for n in neighbors(*x):
            if n in world:
                active_neighbors += 1
            else:
                try:
                    inactive_neighs_to_active[n] += 1
                except KeyError:
                    inactive_neighs_to_active[n] = 1

        if active_neighbors not in (2, 3):
            deactivate.add(x)

    world -= deactivate

    for k, v in inactive_neighs_to_active.items():
        if v == 3:
            world.add(k)


def render3(world):
    minx = min(x[0] for x in world)
    maxx = max(x[0] for x in world)
    miny = min(x[1] for x in world)
    maxy = max(x[1] for x in world)
    minz = min(x[2] for x in world)
    maxz = max(x[2] for x in world)
    slices = []
    for z in range(minz, maxz + 1):
        slice = []
        for y in reversed(range(miny, maxy + 1)):
            row = ""
            for x in range(minx, maxx + 1):
                if (x, y, z) in world:
                    row += "#"
                else:
                    row += "."
            slice.append(row)
        slices.append(slice)

    for i in range(len(slices[0])):
        print(" ".join(s[i] for s in slices))


def main():
    with open("17.in") as f:
        input = [l.strip() for l in f.readlines()]

    # Part 1
    cubes = set()
    for y, row in enumerate(input):
        for x, c in enumerate(row):
            if c == "#":
                cubes.add((x, -y, 0))

    for i in range(6):
        iterate(cubes)
        # render3(cubes)
        # print()

    print(len(cubes))

    # Part 2
    cubes = set()
    for y, row in enumerate(input):
        for x, c in enumerate(row):
            if c == "#":
                cubes.add((x, -y, 0, 0))

    for i in range(6):
        iterate(cubes)

    print(len(cubes))


if __name__ == "__main__":
    main()
