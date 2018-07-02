package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// O(n^2)
func insertionSort(n int32, arr []int32) {
    // O(n)
    for i := int32(1); i < n; i++ {
        insertionSortAux(i + 1, arr)
    }
}

func insertionSortAux(n int32, arr []int32) {
    tmp := arr[n - 1]

    // O(n)
    for i := n - 2; i >= 0; i-- {
        if arr[i] < tmp {
            arr[i + 1] = tmp
            printArray(arr)
            return
        } else {
           arr[i + 1] = arr[i]
        }
    }

    arr[0] = tmp
    printArray(arr)
}

func printArray(arr []int32) {
    for i := 0; i < len(arr); i++ {
        fmt.Printf("%d ", arr[i])
    }
    fmt.Println("")
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    // first line of input contains size of array
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

    insertionSort2(n, arr)
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
