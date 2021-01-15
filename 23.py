from collections import deque
from itertools import cycle


class Ring:
    def __init__(self, list):
        self._data = list
        self.pos = 0

    def current(self):
        return self._data[self.pos]

    def next(self, n=3):
        n_next = deque()
        next_pos = self.pos
        for i in range(n):
            next_pos = (next_pos + 1) % len(self._data)
            n_next.append(self._data[next_pos])

        rest = deque()
        while len(n_next) + len(rest) < len(self._data) - 1:
            next_pos = (next_pos + 1) % len(self._data)
            rest.append(self._data[next_pos])

        target = self.current() - 1
        while target not in rest:
            if target < min(rest):
                target = max(rest)
                break
            target -= 1

        while n_next or rest:
            self.go()
            self._data[self.pos] = rest.popleft()
            if self.current() == target:
                while n_next:
                    self.go()
                    self._data[self.pos] = n_next.popleft()

        self.go()
        self.go()

    def go(self):
        self.pos = (self.pos + 1) % len(self._data)


def main():
    input = "389125467"
    # input = "318946572"

    cups = []
    for x in input:
        cups.append(int(x))
    ring = Ring(cups)
    for i in range(100):
        print(i)
        ring.next()

    res = deque(ring._data)
    while res[0] != 1:
        res.rotate()
    res.popleft()
    print("".join((str(x) for x in res)))

    cups = []
    for x in input:
        cups.append(int(x))

    print(cups)
    for i in range(max(cups) + 1, 1000001):
        cups.append(1)
    print(len(cups))
    ring = Ring(cups)
    for i in range(100):
        print(i)
        ring.next()

    res = deque(ring._data)
    while res[0] != 1:
        res.rotate()
    res.popleft()
    print("".join((str(x) for x in res)))


if __name__ == "__main__":
    main()
