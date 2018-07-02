#!/usr/bin/python


def two_sum(a, n):
    """ Return two indices in arr s.t. the sum
        of their elements is closest to n.
    """
    if len(a) <= 1:
        return [0, 0]

    d = {}

    for i in range(len(a)):
        if a[i] in d:
            return [d[a[i]] + 1, i + 1]
        else:
            d[n - a[i]] = i


if __name__ == "__main__":
    import sys

    t = int(sys.stdin.readline().strip())

    for _ in range(t):
        money = int(sys.stdin.readline().strip())
        n = int(sys.stdin.readline().strip())
        arr = list(map(int, sys.stdin.readline().split()))

        result = two_sum(arr, money)
        print("{} {}".format(result[0], result[1]))
