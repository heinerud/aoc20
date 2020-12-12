import math

directions = {"N": 0 + 1j, "S": 0 - 1j, "E": 1 + 0j, "W": -1 + 0j}


class Ship:
    def __init__(self, pos, dir):
        self.pos = pos
        self.dir = dir

    def move(self, dir, value):
        if dir == "F":
            d = self.dir
        else:
            d = directions[dir]
        self.pos += d * value

    def rotate(self, angle):
        x = math.cos(angle) * self.dir.real - math.sin(angle) * self.dir.imag
        y = math.sin(angle) * self.dir.real + math.cos(angle) * self.dir.imag
        self.dir = complex(x, y)


if __name__ == "__main__":
    with open("12.in") as f:
        input = [x.strip() for x in f.readlines()]

    s = Ship(0 + 0j, 1 + 0j)
    for x in input:
        action = x[0]
        value = int(x[1:])
        if action in "NSEWF":
            s.move(action, value)
        elif action in "LR":
            if action == "R":
                value *= -1
            value *= math.pi / 180
            s.rotate(value)

    print(s.pos)
    print(abs(s.pos.real) + abs(s.pos.imag))
