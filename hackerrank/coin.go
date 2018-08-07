package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

func ways(amount int, coins []int) int {
    kinds_of_coin := len(coins)
    /*

        memo[i] the number of ways to make change with the
        ith kind of coin. More more than the amount is needed
        for space as the memo is built from the bottom-up using
        zero (no money) as the base case.

    */
    memo := make([]int, amount + 1)  // O(amount + 1) space
    memo[0] = 1 // amount is zero

    /*

        Pick each coin one-by-one and update the memo's
        values after the index greater than or equal to
        the value of the picked coin.

    */
    var i, j int  // O(amount * kinds_of_coin) time
    for i = 0; i < kinds_of_coin; i++ {
        for j = coins[i]; j <= amount; j++ {
            memo[j] += memo[j - coins[i]]
        }
    }

    return memo[amount]
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nm := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nm[0], 10, 64)
    checkError(err)
    n := int(nTemp)

    mTemp, err := strconv.ParseInt(nm[1], 10, 64)
    checkError(err)
    m := int(mTemp)

    coinsTemp := strings.Split(readLine(reader), " ")

    var coins []int

    for i := 0; i < int(m); i++ {
        coinsItemTemp, err := strconv.ParseInt(coinsTemp[i], 10, 64)
        checkError(err)
        coinsItem := int(coinsItemTemp)
        coins = append(coins, coinsItem)
    }

    res := ways(n, coins)

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
