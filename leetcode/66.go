package main

import "fmt"

func plusOne(digits []int) []int {
    curr_pos := len(digits) - 1
    last_val := 0

    for curr_pos >= 0 {
        last_val = digits[curr_pos] + 1

        if last_val > 9 { // then last_val = 10
            digits[curr_pos] = 0
            curr_pos -= 1
        } else {
            digits[curr_pos] = last_val
            break
        }
    }

    if last_val > 9 {
        digits = append([]int{1}, digits...)
    }

    return digits
}

func main() {
    digits := []int{9, 9, 9}
    fmt.Printf("%v", plusOne(digits))
}