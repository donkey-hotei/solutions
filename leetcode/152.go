package main

import "fmt"

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    p_min, p_max, p := nums[0], nums[0], nums[0]

    for _, n := range nums[1:] {
        fmt.Printf("n=%d\n", n)
        if n < 0 {
            p_min, p_max = p_max, p_min
        }

        fmt.Printf("p_min=%d, p_max=%d\n", p_min, p_max)

        p_min = min(n * p_min, n)
        p_max = max(n * p_max, n)

        p = max(p, p_max)
    }

    return p
}


func main() {
    a := []int{-2, 0, -1}
    fmt.Printf("%d\n", maxProduct(a[:]))
}