import itertools


class Memory:
    def __init__(self, turn):
        self.first = turn
        self.last = turn

    @property
    def age(self):
        return self.last - self.first

    def add(self, turn):
        self.first, self.last = self.last, turn


if __name__ == "__main__":
    input = [20, 9, 11, 0, 1, 2]

    numbers = {}
    for i, x in enumerate(input):
        numbers[x] = Memory(i + 1)

    s = input[-1]
    for turn in itertools.count(len(input) + 1):
        s = numbers[s].age
        try:
            numbers[s].add(turn)
        except KeyError:
            numbers[s] = Memory(turn)

        if turn == 2020:
            print(s)

        if turn == 30000000:
            print(s)
            break
