package main

import "sort"

func threesum(nums []int) [][]int {
    sort.Ints(nums)

    n := len(nums)
    triplets := make(map[[]int]bool, 0)

    for i := 0; i < n - 2; i++ {
        l := i + 1;
        r := n - 1;

        for l < r {
            s := nums[i] + nums[l] + nums[r]  

            if s == 0 {
                triplet := []int{nums[i], nums[l], nums[r]}
                triplets = append(triplets, triplet)

                l += 1
                r -= 1
            } else if s < 0 {
                l += 1
            } else {
                r -= 1
            }
        }
    }

    return triplets
}

func main() {
    a := [6]int{-1, 0, 1, 2, -1, -4}
    threesum(a[:])
}