def game(input, turns):
    numbers = {}
    for i, x in enumerate(input):
        numbers[x] = i + 1

    speak = 0
    for turn in range(len(input) + 1, turns):
        spoken = speak
        if spoken in numbers:
            speak = turn - numbers[spoken]
        else:
            speak = 0

        numbers[spoken] = turn

    return speak


if __name__ == "__main__":
    input = [20, 9, 11, 0, 1, 2]

    print(game(input, 2020))
    print(game(input, 30000000))
