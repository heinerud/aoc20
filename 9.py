from itertools import combinations


def sum_in_preamble(pre, x):
    for a, b in combinations(pre, 2):
        if a + b == x:
            return True

    return False


if __name__ == "__main__":
    with open("9.in") as f:
        lines = [x.strip() for x in f.readlines()]

    input = [int(x) for x in lines]

    for i in range(25, len(input)):
        preamble = input[i - 25 : i]
        if not sum_in_preamble(preamble, input[i]):
            faulty = input[i]
            break

    print(faulty)

    for i in range(len(input)):
        numbers = []
        for x in input[i:]:
            numbers.append(x)
            if sum(numbers) >= faulty:
                break

        if sum(numbers) == faulty:
            break

    print(min(numbers) + max(numbers))
