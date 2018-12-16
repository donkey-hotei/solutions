package main

import "fmt"

func findDuplicate(nums []int) int {
    r, t := nums[0], nums[0]
    for {
        r = nums[nums[r]]
        t = nums[t]

        if r == t {
            break
        }
    }

    p := nums[0]
    q := t

    for p != q {
        p = nums[p]
        q = nums[q]
    }

    return p
}

func main() {
    a := []int{2,5,9,6,9,3,8,9,7,1};
    fmt.Printf("%d\n", findDuplicate(a))
}