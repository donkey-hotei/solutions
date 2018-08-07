package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func quickSort(arr []int32) []int32 {
    pivot := arr[0]

    lo := make([]int32, 0, len(arr))
    eq := make([]int32, 0, len(arr))
    hi := make([]int32, 0, len(arr))

    for i := 0; i < len(arr); i++ {
        if arr[i] > pivot {
            hi = append(hi, arr[i])
        } else if arr[i] == pivot {
            eq = append(eq, arr[i])
        } else {
            lo = append(lo, arr[i])
        }
    }

    result := make([]int32, 0, len(arr))
    result = append(result, lo...)
    result = append(result, eq...)
    result = append(result, hi...)

    return result
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

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := quickSort(arr)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, " ")
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

