package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int) int {
    sums := make([]int, 0)
    // skipping the first and last (row, column)
    for row := 1; row < len(arr) - 1; row++ {
        for col := 1; col < len(arr) - 1; col++ {
            s := 0
            s += arr[row - 1][col - 1]
            s += arr[row - 1][col]
            s += arr[row - 1][col + 1]
            s += arr[row][col]
            s += arr[row + 1][col - 1]
            s += arr[row + 1][col]
            s += arr[row + 1][col + 1]
            println(s)
            sums = append(sums, s)
        }
    }
    max := sums[0]
    for _, sum:= range sums {
        if sum > max {
            max = sum
        }
    }
    return max
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    var arr [][]int
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(readLine(reader), " ")

        var arrRow []int
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(6) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    result := hourglassSum(arr)

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


