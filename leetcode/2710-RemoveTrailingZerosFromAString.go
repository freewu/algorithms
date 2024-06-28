package main

// 2710. Remove Trailing Zeros From a String
// Given a positive integer num represented as a string, return the integer num without trailing zeros as a string.

// Example 1:
// Input: num = "51230100"
// Output: "512301"
// Explanation: Integer "51230100" has 2 trailing zeros, we remove them and return integer "512301".

// Example 2:
// Input: num = "123"
// Output: "123"
// Explanation: Integer "123" has no trailing zeros, we return integer "123".

// Constraints:
//     1 <= num.length <= 1000
//     num consists of only digits.
//     num doesn't have any leading zeros.

import "fmt"
import "strings"

func removeTrailingZeros(num string) string {
    n, bs := len(num), []byte(num)
    for i := n - 1; i >= 0; i-- {
        if bs[i] == '0' {
            n--
        } else {
            break
        }
    }
    if n == 0 {
        return ""
    }
    return string(bs[:n])
}

// lib
func removeTrailingZeros1(num string) string {
    return strings.TrimRight(num, "0")
}

func removeTrailingZeros2(num string) string {
    i := len(num) - 1
    for i > 0 && num[i] == '0' {
        i--
    }
    return num[:i + 1]
}

func main() {
    // Example 1:
    // Input: num = "51230100"
    // Output: "512301"
    // Explanation: Integer "51230100" has 2 trailing zeros, we remove them and return integer "512301".
    fmt.Println(removeTrailingZeros("51230100")) // "512301"
    // Example 2:
    // Input: num = "123"
    // Output: "123"
    // Explanation: Integer "123" has no trailing zeros, we return integer "123".
    fmt.Println(removeTrailingZeros("123")) // "123"
    fmt.Println(removeTrailingZeros("0000")) // ""

    fmt.Println(removeTrailingZeros1("51230100")) // "512301"
    fmt.Println(removeTrailingZeros1("123")) // "123"
    fmt.Println(removeTrailingZeros1("0000")) // ""

    fmt.Println(removeTrailingZeros2("51230100")) // "512301"
    fmt.Println(removeTrailingZeros2("123")) // "123"
    fmt.Println(removeTrailingZeros2("0000")) // ""
}