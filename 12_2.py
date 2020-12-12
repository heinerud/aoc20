import math

directions = {"N": 0 + 1j, "S": 0 - 1j, "E": 1 + 0j, "W": -1 + 0j}


class Waypoint:
    def __init__(self, pos):
        self.pos = pos

    def move(self, dir, value):
        self.pos += directions[dir] * value

    def rotate(self, angle):
        x = math.cos(angle) * self.pos.real - math.sin(angle) * self.pos.imag
        y = math.sin(angle) * self.pos.real + math.cos(angle) * self.pos.imag
        self.pos = complex(x, y)


class Ship:
    def __init__(self, pos):
        self.pos = pos

    def move(self, waypoint, times):
        for i in range(times):
            self.pos += waypoint.pos


if __name__ == "__main__":
    with open("12.in") as f:
        input = [x.strip() for x in f.readlines()]

    s = Ship(0 + 0j)
    w = Waypoint(10 + 1j)
    for x in input:
        action = x[0]
        value = int(x[1:])
        if action in "NSEW":
            w.move(action, value)
        elif action in "LR":
            if action == "R":
                value *= -1
            value *= math.pi / 180
            w.rotate(value)
        elif action == "F":
            s.move(w, value)

    print(s.pos)
    print(abs(s.pos.real) + abs(s.pos.imag))
