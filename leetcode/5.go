package main

import "fmt"

func grow(i, j int, s string) (int, int) {
    for i >= 0 && j < len(s) {
        if s[i] != s[j] {
            break
        }

        i -= 1
        j += 1
    }

    return i + 1, j
}

func longest_palindrome(s string) string {
    max_s, max_e := 0, 0

    for i := 0; i < len(s); i++ {
        for j := i; j <= i + 1; j++ {
            s, e := grow(i, j, s)

            if e - s > max_e - max_s {
                max_s = s
                max_e = e
            }
        }
    }

    return s[max_s:max_e]
}


func main() {
    s := "babad"
    fmt.Printf("longest substring palindrome of %s is %s\n", s, longest_palindrome(s))
    s = "cbbd"
    fmt.Printf("longest substring palindrome of %s is %s\n", s, longest_palindrome(s))
}