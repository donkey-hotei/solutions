#!/usr/bin/python

memo = {0: 0, 1: 1, 2: 2, 3: 4}


def climb(n):
    if n not in memo:
        resu = climb(n - 1) + climb(n - 2) + climb(n - 3)
        memo[n] = resu

    return memo[n]


if __name__ == "__main__":
    import sys

    s = int(sys.stdin.readline().strip())
    for _ in range(s):
        n = int(sys.stdin.readline().strip())
        print("{}".format(climb(n)))
