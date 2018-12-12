package main

// import "github.com/stretchr/testify/assert"
import "fmt"

func searchInsert(nums []int, target int) int {
    index := 0

    for i, n := range nums {
        if target == n {
            return i
        } else if target < n && i == 0 {
            return 0
        } else if target < n {
            return (i + 1) - 1
        }
        index = i
    }

    return index + 1
}

func main() {
    // assert.Equal(searchInsert([]int{1, 3, 5, 6}, 5), 2)
    // assert.Equal(searchInsert([]int{1, 3, 5, 6}, 2), 1)
    // assert.Equal(searchinsert([]int{1, 3, 5, 6}, 7), 4)
    fmt.Printf("%d\n", searchInsert([]int{1, 3, 5, 6}, 2))
    fmt.Printf("%d\n", searchInsert([]int{1, 3, 5, 6}, 7))
}
