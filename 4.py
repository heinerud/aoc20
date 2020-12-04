#!/usr/bin/env python3

import sys


def passport(*, byr, iyr, eyr, hgt, hcl, ecl, pid, cid=None):
    if len(byr) != 4 or not 1920 <= int(byr) <= 2002:
        return False

    if len(iyr) != 4 or not 2010 <= int(iyr) <= 2020:
        return False

    if len(eyr) != 4 or not 2020 <= int(eyr) <= 2030:
        return False

    if hgt.endswith("cm"):
        if not 150 <= int(hgt[:-2]) <= 193:
            return False
    elif hgt.endswith("in"):
        if not 59 <= int(hgt[:-2]) <= 76:
            return False
    else:
        return False

    if not hcl.startswith("#") or len(hcl[1:]) != 6 or not hcl[1:].isalnum():
        return False

    if not ecl in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]:
        return False

    if not len(pid) == 9 or not pid.isdigit():
        return False

    return True


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        lines = [x.strip() for x in f.readlines()]

    entries = []
    entry = []
    for l in lines:
        if not l:
            entries.append(entry)
            entry = []
        else:
            entry.extend(l.split())

    if entry:
        entries.append(entry)

    passports = []
    for e in entries:
        try:
            p = passport(**{k: v for k, v in [x.split(":") for x in e]})
        except TypeError:
            pass
        else:
            passports.append(p)

    print(len(passports))
    print(passports.count(True))
