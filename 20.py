import math


def flip(m):
    """Flip a matrix along vertical axis"""
    return [x[::-1] for x in m]


def rot(m):
    """Rotate a matrix 90 degrees counter clockwise"""
    rotated = []
    for i in range(len(m[0]) - 1, -1, -1):
        rot = [m[j][i] for j in range(len(m))]
        rotated.append("".join(rot))

    return rotated


def fits(edge1, edge2):
    for x1, x2 in zip(edge1, edge2):
        if x1 != x2:
            return False
    return True


def match(t1, t2):
    for f in [rot, rot, rot, rot, flip, rot, rot, rot]:
        t2 = f(t2)
        t1_right = [x[-1] for x in t1]
        t2_left = [x[0] for x in t2]
        if fits(t1_right, t2_left):
            return t2, "right"
        if fits(t1[0], t2[-1]):
            return t2, "top"
        t2_right = [x[-1] for x in t2]
        t1_left = [x[0] for x in t1]
        if fits(t2_right, t1_left):
            return t2, "left"
        if fits(t1[-1], t2[0]):
            return t2, "bottom"
    return None, None


def remove_border(tile):
    tile = tile[1:-1]
    tile = [x[1:-1] for x in tile]
    return tile


def pattern_match(p1, p2):
    for r1, r2 in zip(p1, p2):
        for x1, x2 in zip(r1, r2):
            if x1 == "#" and x2 != x1:
                return False
    return True


def part1(tiles):
    res = 1
    for k, v in tiles.items():
        fitting = []
        for kk, vv in tiles.items():
            if kk == k:
                continue
            vv, pos = match(v, vv)
            if vv:
                fitting.append((vv, pos))

        if len(fitting) == 2:
            res *= k

    print(res)


def part2(tiles):
    size = int(math.sqrt(len(tiles)))
    puzzle = [[None for _ in range(size)] for _ in range(size)]
    # Find a corner piece
    for k, v in tiles.items():
        adjacent = []
        for kk, vv in tiles.items():
            if kk == k:
                continue
            vv, pos = match(v, vv)
            if vv:
                adjacent.append((kk, vv, pos))

        if len(adjacent) == 2:
            pos_list = [pos for _, _, pos in adjacent]
            if "left" in pos_list and "bottom" in pos_list:
                rotations = 1
            elif "top" in pos_list and "left" in pos_list:
                rotations = 2
            elif "top" in pos_list and "right" in pos_list:
                rotations = 3
            else:
                rotations = 0

            for _ in range(rotations):
                v = rot(v)

            puzzle[0][0] = v
            del tiles[k]
            break

    row = 0
    col = 0
    # Find the rest of the pieces one by one in a snake pattern
    while tiles:
        adjacent = []
        for k, v in tiles.items():
            v, pos = match(puzzle[row][col], v)
            if v:
                adjacent.append((k, v, pos))

        pos_list = [pos for _, _, pos in adjacent]
        if "right" in pos_list:
            col += 1
            next = adjacent[pos_list.index("right")]
        elif "left" in pos_list:
            col -= 1
            next = adjacent[pos_list.index("left")]
        else:
            row += 1
            next = adjacent[pos_list.index("bottom")]

        id = next[0]
        tile = next[1]
        puzzle[row][col] = tile
        del tiles[id]

    # Remove borders from tiles
    puzzle2 = []
    for r in [[remove_border(x) for x in r] for r in puzzle]:
        for i in range(len(r[0])):
            puzzle2.append("".join([x[i] for x in r]))
    puzzle = puzzle2

    # Find monsters
    monster = ["                  # ", "#    ##    ##    ###", " #  #  #  #  #  #   "]
    num_monsters = 0
    for f in [rot, rot, rot, rot, flip, rot, rot, rot]:
        puzzle = f(puzzle)
        for r in range(len(puzzle) - len(monster) + 1):
            for c in range(len(puzzle[0]) - len(monster[0]) + 1):
                slice = [
                    row[c : c + len(monster[0]) + 1]
                    for row in puzzle[r : r + len(monster)]
                ]
                if pattern_match(monster, slice):
                    num_monsters += 1
        if num_monsters:
            break

    total_waves = sum(r.count("#") for r in puzzle)
    monster_waves = sum(r.count("#") for r in monster)
    print(total_waves - monster_waves * num_monsters)


def main():
    with open("20.in") as f:
        input = [l.strip() for l in f.readlines()]

    tiles = {}
    for l in input:
        if "Tile" in l:
            _, id = l.split()
            id = int(id[:-1])
            rows = []

        elif not l:
            tiles[id] = rows
        else:
            rows.append(l)
    tiles[id] = rows

    part1(tiles)
    part2(tiles)


if __name__ == "__main__":
    main()
