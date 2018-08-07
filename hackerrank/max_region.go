package main

import (
    "fmt"
)

func dfs(grid [][]int, row, col, size int) int {
    n := len(grid)
    m := len(grid[0])

    if row >= n || row < 0 || col >= m || col < 0 {
        return size
    }

    if grid[row][col] == 0 {
        return size
    }

    grid[row][col] = 0
    size++

    size = dfs(grid, row - 1, col, size)     // N
    size = dfs(grid, row -1, col - 1, size)  // NW
    size = dfs(grid, row, col -1, size)      // W
    size = dfs(grid, row + 1, col - 1, size) // SW
    size = dfs(grid, row + 1, col, size)     // S
    size = dfs(grid, row + 1, col + 1, size) // SE
    size = dfs(grid, row, col + 1, size)     // E
    size = dfs(grid, row - 1, col + 1, size) // NE

    return size
}


func printGrid(grid [][]int) {
    for _, row := range grid {
        fmt.Printf("%v\n", row)
    }
    println("")
}

func maxRegion(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])

    max_size := 0

    // O(n * m)
    for row := 0; row < n; row++ {
        for col := 0; col < m; col++ {
            if grid[row][col] == 1 {
                size := dfs(grid, row, col, 0)

                if size > max_size {
                    max_size = size
                }
            }
        }
    }

    return max_size
}

func main() {
    grid := [][]int{
        { 1, 1, 0, 0, 0, 1},
        { 0, 1, 1, 0, 0, 0},
        { 0, 0, 0, 1, 0, 0},
    }

    c := maxRegion(grid)
    fmt.Printf("Largest region has %d cells.", c)
}