package main

import "fmt"

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func max_area(h []int) int {
    res  := 0
    i, j := 0, len(h) - 1

    for i < j {
        w := j - i

        if h[i] < h[j] {
            res = max(res, w * h[i])
            i += 1
        } else {
            res = max(res, w * h[j])
            j -= 1
        }
    }

    return res
}

func main() {
    h := [9]int{ 1, 8, 6, 2, 5, 4, 8, 3, 7 }
    fmt.Printf("max area is %d\n", max_area(h[:]))
}