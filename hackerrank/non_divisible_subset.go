package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the nonDivisibleSubset function below.
func nonDivisibleSubset(k int, S []int) int {

    /*

        if (n1 + n2) / k leaves a whole remainder then,
        (n1 mod k) + (n2 mod k) == k so we can break
        this down from a problem of division into one
        of addition.

    */

    // map S_i % k to the count of their occurance
    remcounts := make(map[int]int, 0)

    for _, v := range S {
        remcounts[v % k] += 1
    }


}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nk := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int(kTemp)

    STemp := strings.Split(readLine(reader), " ")

    var S []int

    for i := 0; i < int(n); i++ {
        SItemTemp, err := strconv.ParseInt(STemp[i], 10, 64)
        checkError(err)
        SItem := int(SItemTemp)
        S = append(S, SItem)
    }

    result := nonDivisibleSubset(k, S)

    fmt.Fprintf(writer, "%d\n", result)

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
