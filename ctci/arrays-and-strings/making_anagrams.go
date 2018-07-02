package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/* Determines number of deletions needed to make a and b anagrams */
func makeAnagram(a, b string) int32 {
    count := int32(0)
    freqs := make(map[rune]int32)

    for _, c := range a {
        freqs[c]++
    }

    for _, c := range b {
        freqs[c]--
    }

    for _, v := range freqs {
        count += int32(math.Abs(float64(v)))
    }

    return count
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    a := readLine(reader)

    b := readLine(reader)

    res := makeAnagram(a, b)

    fmt.Fprintf(writer, "%d\n", res)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

