package main;

import (
    "fmt"
)

func length_of_longest_substring(s string) int {
    seen := make(map[byte]bool)
    i, j, max, n := 0, 0, 0, len(s) - 1

    for j <= n {
        _, p := seen[s[j]]

        if !p {
            seen[s[j]] = true
            j += 1

            if j - i > max {
                max = j - i
            }

        } else {
            delete(seen, s[i])
            i += 1
        }
    }

    return max
}

func main() {
    s := "abcabcbb"
    fmt.Printf("s=%s, len=%d\n", s, length_of_longest_substring(s))
    s = "xxquizz"
    fmt.Printf("s=%s, len=%d\n", s, length_of_longest_substring(s))
    s = "bbbbbbb"
    fmt.Printf("s=%s, len=%d\n", s, length_of_longest_substring(s))
    s = "au"
    fmt.Printf("s=%s, len=%d\n", s, length_of_longest_substring(s))
}