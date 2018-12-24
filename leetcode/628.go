package main

import "fmt"

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func maximumProduct(nums []int) int {
    if len(nums) < 3 {
        return 0
    }

    lo1 := int(^uint(0) >> 1)
    lo2 := int(^uint(0) >> 1)
    hi1 := -lo1 - 1
    hi2 := -lo1 - 1
    hi3 := -lo1 - 1

    for _, n := range nums {
        if n <= lo1 {
            lo2 = lo1
            lo1 = n
        } else if n <= lo2 {
            lo2 = n
        }

        if hi3 <= n {
            hi1 = hi2
            hi2 = hi3
            hi3 = n
        } else if hi2 <= n {
            hi1 = hi2
            hi2 = n
        } else if hi1 <= n {
            hi1 = n
        }
    }

    return max(lo1 * lo2 * hi3, hi1 * hi2 * hi3)
}

func main() {
    a := []int{1, 2, 3, 4}

    fmt.Printf("%d\n", maximumProduct(a[:]))
}