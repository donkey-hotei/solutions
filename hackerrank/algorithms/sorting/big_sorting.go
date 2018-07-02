package main

import (
    "bufio"
    "fmt"
    "io"
    "math/big"
    "os"
    "strconv"
    "strings"
)

func bigSorting(a []string) []string {
    n := len(a)

    if n <= 1 {
        return a
    }

    // divide the array into two halves
    l := a[:n / 2]
    r := a[n / 2:]

    // and proceed to recursively split
    l = bigSorting(l)
    r = bigSorting(r)

    // before finally merging
    return merge(l, r)
}


func merge(left, right []string) []string {
    l_len := len(left)
    r_len := len(right)
    n := l_len + r_len

    array := make([]string, n)

    l, r := 0, 0
    for k := 0; k < n; k++ {
        if l >= l_len {
            array[k] = right[r]
            r++
            continue
        } else if r >= r_len {
            array[k] = left[l]
            l++
            continue
        }

        l_val := big.NewInt(0)
        r_val := big.NewInt(0)
        l_val.SetString(left[l], 10)
        r_val.SetString(right[r], 10)

        if (*l_val).Uint64() >= (*r_val).Uint64() {
            array[k] = right[r]
            r++
        } else {
            array[k] = left[l]
            l++
        }
    }

    return array
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var unsorted []string

    for i := 0; i < int(n); i++ {
        unsortedItem := readLine(reader)
        unsorted = append(unsorted, unsortedItem)
    }

    result := bigSorting(unsorted)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%s", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
