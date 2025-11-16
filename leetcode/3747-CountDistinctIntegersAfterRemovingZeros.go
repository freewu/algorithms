package main

// 3747. Count Distinct Integers After Removing Zeros
// You are given a positive integer n.

// For every integer x from 1 to n, we write down the integer obtained by removing all zeros from the decimal representation of x.

// Return an integer denoting the number of distinct integers written down.

// Example 1:
// Input: n = 10
// Output: 9
// Explanation:
// The integers we wrote down are 1, 2, 3, 4, 5, 6, 7, 8, 9, 1. There are 9 distinct integers (1, 2, 3, 4, 5, 6, 7, 8, 9).

// Example 2:
// Input: n = 3
// Output: 3
// Explanation:
// The integers we wrote down are 1, 2, 3. There are 3 distinct integers (1, 2, 3).

// Constraints:
//     1 <= n <= 10^15

import "fmt"
import "strconv"
import "math"

func countDistinct(n int64) int64 {
    s := strconv.FormatInt(n, 10)
    m := len(s)

    // 计算长度小于 m 的不含 0 的整数个数
    // 9 + 9^9 + ... + 9^(m-1) = (9^m - 9) / 8
    pow9 := int64(math.Pow(9, float64(m)))
    res := (pow9 - 9) / 8

    // 计算长度恰好等于 m 的不含 0 的整数个数
    for i, d := range s {
        if d == '0' { // 只能填 0，不合法，跳出循环
            break
        }
        // 这一位填 1 到 d-1，后面的数位可以随便填 1 到 9
        v := d - '1'
        if i == m-1 {
            v++ // 最后一位可以等于 d
        }
        pow9 /= 9
        res += int64(v) * pow9
        // 然后，这一位填 d，继续遍历
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 10
    // Output: 9
    // Explanation:
    // The integers we wrote down are 1, 2, 3, 4, 5, 6, 7, 8, 9, 1. There are 9 distinct integers (1, 2, 3, 4, 5, 6, 7, 8, 9).
    fmt.Println(countDistinct(10)) // 9
    // Example 2:
    // Input: n = 3
    // Output: 3
    // Explanation:
    // The integers we wrote down are 1, 2, 3. There are 3 distinct integers (1, 2, 3).
    fmt.Println(countDistinct(3)) // 3

    fmt.Println(countDistinct(1)) // 1
    fmt.Println(countDistinct(8)) // 8
    fmt.Println(countDistinct(64)) // 58
    fmt.Println(countDistinct(99)) // 90
    fmt.Println(countDistinct(100)) // 90
    fmt.Println(countDistinct(1024)) // 819
    fmt.Println(countDistinct(1_000_000_007)) // 435848049
    fmt.Println(countDistinct(999_999_999_999_999)) // 231627523606479
    fmt.Println(countDistinct(1_000_000_000_000_000)) // 231627523606479
}