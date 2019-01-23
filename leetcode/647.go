package main

func grow(i, j int, s string) int {
    c := 0

    for i >= 0 && j < len(s) {
        if s[i] != s[j] {
            break
        }
        j += 1
        i -= 1
        c += 1
    }

    return c
}

func countSubstrings(s string) int {
    count := 0

    for i := 0; i < len(s); i++ {
        for j := i; j <= i + 1; j++ {
            count += grow(i, j, s)
        }
    }

    return count
}