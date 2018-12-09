package main

import "fmt"

func removeElements(nums []int, val int) int {
    remove := false
    last_n := 0

    for _, n := range nums {
        if n == val {
            remove = true
        } else if remove {
            nums[last_n] = n
            last_n += 1
        } else {
            last_n += 1
        }
    }

    return last_n
}

func main() {
    a := []int{3, 2, 2, 3}
    fmt.Printf("%d\n", removeElements(a[:], 3))
    fmt.Printf("%v\n", a)
}