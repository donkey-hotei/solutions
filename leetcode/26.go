package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    } else if len(nums) == 1 {
        return 1
    }

    uniq := 1
    curr := nums[0]

    for _, n := range nums[1:] {
        if n != curr {
            nums[uniq] = n
            uniq += 1
            curr = n
        }
    }

    return uniq
}

func main() {
    a := []int{1, 1, 2}

    fmt.Printf("%d", removeDuplicates(a[:]))
    fmt.Printf("%v", a)
}