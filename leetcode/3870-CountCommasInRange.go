package main

// 3870. Count Commas in Range
// You are given an integer n.

// Return the total number of commas used when writing all integers from [1, n] (inclusive) in standard number formatting.

// In standard formatting:
//     1. A comma is inserted after every three digits from the right.
//     2. Numbers with fewer than 4 digits contain no commas.

// Example 1:
// Input: n = 1002
// Output: 3
// Explanation:
// The numbers "1,000", "1,001", and "1,002" each contain one comma, giving a total of 3.

// Example 2:
// Input: n = 998
// Output: 0
// Explanation:
// All numbers from 1 to 998 have fewer than four digits. Therefore, no commas are used.

// Constraints:
//     1 <= n <= 10^5

import "fmt"

func countCommas(n int) int {
    // 小于1000的数没有逗号，直接返回0
    if n < 1000 { return 0 }
    // 大于等于1000的数，从1000到n的每个数贡献1个逗号
    // 总数 = n - 999
    return n - 999
}

func main() {
    // Example 1:
    // Input: n = 1002
    // Output: 3
    // Explanation:
    // The numbers "1,000", "1,001", and "1,002" each contain one comma, giving a total of 3.
    fmt.Println(countCommas(1002)) // 3
    // Example 2:
    // Input: n = 998
    // Output: 0
    // Explanation:
    // All numbers from 1 to 998 have fewer than four digits. Therefore, no commas are used.
    fmt.Println(countCommas(998)) // 0

    fmt.Println(countCommas(1)) // 0
    fmt.Println(countCommas(2)) // 0
    fmt.Println(countCommas(8)) // 0
    fmt.Println(countCommas(64)) // 0
    fmt.Println(countCommas(99)) // 0
    fmt.Println(countCommas(100)) // 0
    fmt.Println(countCommas(1024)) // 25
    fmt.Println(countCommas(99_999)) // 99000
    fmt.Println(countCommas(100_000)) // 399001
}