def main():
    with open("21.in") as f:
        input = [l.strip() for l in f.readlines()]

    food = []
    for l in input:
        ingredients, allergens = l.split("(contains ")
        ingredients = ingredients.split()
        allergens = allergens.rstrip(")").split(", ")
        food.append((set(ingredients), set(allergens)))

    allergens = {}
    for ings, alls in food:
        for a in alls:
            try:
                allergens[a] &= ings
            except KeyError:
                allergens[a] = set(ings)

    known = set()
    while len(known) < len(allergens):
        for suspects in allergens.values():
            if len(suspects - known) == 1:
                suspects -= known
                known |= suspects

    print(sum(len(ings - known) for ings, _ in food))
    print(",".join(next(iter(i)) for _, i in sorted(allergens.items())))


if __name__ == "__main__":
    main()
