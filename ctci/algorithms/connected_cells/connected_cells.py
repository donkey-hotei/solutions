#!/usr/bin/python

#
# DFS: Connected Cells in a Grid
#
# Write program to return the size of the largest
# connected region in the grid.
#


def get_region_size(grid, row, column):
    if not(row in range(len(grid)) and column in range(len(grid[0]))):
        return 0
    if (grid[row][column] == 0):
        return 0

    grid[row][column] = 0

    size = 1
    size += get_region_size(grid, row + 1, column)      # south
    size += get_region_size(grid, row - 1, column)      # north
    size += get_region_size(grid, row, column + 1)      # east
    size += get_region_size(grid, row, column - 1)      # west
    size += get_region_size(grid, row - 1, column - 1)  # north-west
    size += get_region_size(grid, row + 1, column + 1)  # south-east
    size += get_region_size(grid, row - 1, column + 1)  # north-east
    size += get_region_size(grid, row + 1, column - 1)  # south-west

    return size


def find_largest_region(grid):
    max_size = 0

    for row in range(len(grid)):
        for column in range(len(grid[row])):
            if grid[row][column] == 1:
                size = get_region_size(grid, row, column)
                if size >= max_size:
                    max_size = size

    return max_size


if __name__ == "__main__":
    import sys

    n = int(sys.stdin.readline())
    m = int(sys.stdin.readline())

    grid = [
        list(map(int, sys.stdin.readline().rstrip().split()))
        for _ in range(n)
    ]

    print(find_largest_region(grid))
