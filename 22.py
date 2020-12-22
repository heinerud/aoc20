from collections import deque


def score(deck):
    score = 0
    for i, x in enumerate(reversed(deck)):
        score += (i + 1) * x
    return score


def part1(p1, p2):
    while p1 and p2:
        c1 = p1.popleft()
        c2 = p2.popleft()
        if c1 > c2:
            p1.append(c1)
            p1.append(c2)
        else:
            p2.append(c2)
            p2.append(c1)

    winner = p1 if p1 else p2
    return score(winner)


def game(p1, p2):
    rounds = set()
    while p1 and p2:
        if (str(p1), str(p2)) in rounds:
            return 1
        rounds.add((str(p1), str(p2)))

        c1 = p1.popleft()
        c2 = p2.popleft()
        if len(p1) >= c1 and len(p2) >= c2:
            p1c = p1.copy()
            p2c = p2.copy()
            while len(p1c) > c1:
                p1c.pop()
            while len(p2c) > c2:
                p2c.pop()
            winner = game(p1c, p2c)
            if winner == 1:
                p1.append(c1)
                p1.append(c2)
            else:
                p2.append(c2)
                p2.append(c1)
        elif c1 > c2:
            p1.append(c1)
            p1.append(c2)
        else:
            p2.append(c2)
            p2.append(c1)

    if p1:
        return 1
    if p2:
        return 2


def part2(p1, p2):
    winner = p1 if game(p1, p2) == 1 else p2
    return score(winner)


def main():
    with open("22.in") as f:
        input = [l.strip() for l in f.readlines()]

    players = []
    for x in input:
        if "Player" in x:
            deck = deque()
        elif not x:
            players.append(deck)
        else:
            deck.append(int(x))

    players.append(deck)
    p1 = players[0]
    p2 = players[1]

    print(part1(p1.copy(), p2.copy()))
    print(part2(p1.copy(), p2.copy()))


if __name__ == "__main__":
    main()
