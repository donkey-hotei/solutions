#!/usr/bin/python

def permutations(array, path):
    if len(array) == 0:
        yield path

    for i in range(len(array)):
        yield from permutations(array[:i] + array[i + 1:], path + [array[i]]


if __name__ == "__main__":
    print(list(permutations([1, 1, 3], [])))