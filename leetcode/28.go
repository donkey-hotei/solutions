package main

import "fmt"

func strStr(haystack string, needle string) int {
    if haystack == needle {
        return 0
    }

    for i := 0; i < len(haystack) - len(needle) + 1; i++ {
        if haystack[i:len(needle) + i] == needle {
            return i
        }
    }

    return -1
}


func main() {
    fmt.Printf("%d\n", strStr("hello", "ll"))
}