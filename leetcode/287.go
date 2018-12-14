package main

func findDuplicate(nums []int) int {
    dupe := nums[0]

    for nums[dupe] != dupe {
        nums[0], nums[dupe] = nums[dupe], nums[0]
        dupe = nums[0]
    }

    return dupe
}