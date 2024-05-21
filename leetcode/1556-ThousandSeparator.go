package main

// 1556. Thousand Separator
// Given an integer n, add a dot (".") as the thousands separator and return it in string format.

// Example 1:
// Input: n = 987
// Output: "987"

// Example 2:
// Input: n = 1234
// Output: "1.234"

// Constraints:
//     0 <= n <= 2^31 - 1

import "fmt"

func thousandSeparator(n int) string {
    if n < 10 { return string('0' + byte(n)); }
    res, i := "", 0
    for ; n > 0; n /= 10 {
        i++
        res = string('0' + byte(n%10)) + res
        if i % 3 == 0 && n > 9 { // 逢 3 加个 .
            res = "." + res
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 987
    // Output: "987"
    fmt.Println(thousandSeparator(987)) // 987
    // Example 2:
    // Input: n = 1234
    // Output: "1.234"
    fmt.Println(thousandSeparator(1234)) // 1.234
}