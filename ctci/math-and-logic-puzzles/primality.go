package main

import (
    "bufio"
    "fmt"
    "io"
    "math"
    "os"
    "strconv"
    "strings"
)

func primality(n int) string {

    /*

        One basic way to test if a number is prime or not is to
        attempt to divide the number by any number less that it,
        and if it is ever cleanly divisible by any number other
        than itself then we have a non-prime, otherwise prime.

    */
    // O(sqrt(n))
    k := int(math.Sqrt(float64(n))) + 1
    for i := 2; i <= k; i++ {
        if n % i == 0 {
            return "Not prime"
        }
    }

    return "Prime"
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    p := int(pTemp)

    for pItr := 0; pItr < int(p); pItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int(nTemp)

        result := primality(n)

        fmt.Fprintf(writer, "%s\n", result)
    }

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

