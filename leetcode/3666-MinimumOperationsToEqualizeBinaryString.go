package main

// 3666. Minimum Operations to Equalize Binary String
// You are given a binary string s, and an integer k.

// In one operation, you must choose exactly k different indices and flip each '0' to '1' and each '1' to '0'.

// Return the minimum number of operations required to make all characters in the string equal to '1'. If it is not possible, return -1.

// Example 1:
// Input: s = "110", k = 1
// Output: 1
// Explanation:
// There is one '0' in s.
// Since k = 1, we can flip it directly in one operation.

// Example 2:
// Input: s = "0101", k = 3
// Output: 2
// Explanation:
// One optimal set of operations choosing k = 3 indices in each operation is:
// Operation 1: Flip indices [0, 1, 3]. s changes from "0101" to "1000".
// Operation 2: Flip indices [1, 2, 3]. s changes from "1000" to "1111".
// Thus, the minimum number of operations is 2.

// Example 3:
// Input: s = "101", k = 2
// Output: -1
// Explanation:
// Since k = 2 and s has only one '0', it is impossible to flip exactly k indices to make all '1'. Hence, the answer is -1.

// Constraints:
//     1 <= s.length <= 10^​​​​​​​5
//     s[i] is either '0' or '1'.
//     1 <= k <= s.length

import "fmt"
import "strings"

func minOperations(s string, k int) int {
    res, n, zero := 1 << 31, len(s), strings.Count(s, "0")
    if zero == 0 {
        return 0
    }
    if n == k {
        if zero == n {
            return 1
        }
        return -1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 情况一：操作次数 m 是偶数
    if zero % 2 == 0 { // z 必须是偶数
        m := max((zero + k - 1) / k, (zero + n - k - 1) / (n - k)) // 下界
        res = m + m % 2 // 把 m 往上调整为偶数
    }
    // 情况二：操作次数 m 是奇数
    if zero % 2 == k % 2 { // z 和 k 的奇偶性必须相同
        m := max((zero + k - 1) / k, (n - zero + n - k - 1) / (n - k)) // 下界
        res = min(res, m | 1) // 把 m 往上调整为奇数
    }
    if res < 1 << 31 {
        return res
    }
    return -1
}

func main() {
    // Example 1:
    // Input: s = "110", k = 1
    // Output: 1
    // Explanation:
    // There is one '0' in s.
    // Since k = 1, we can flip it directly in one operation.
    fmt.Println(minOperations("110", 1)) // 1
    // Example 2:
    // Input: s = "0101", k = 3
    // Output: 2
    // Explanation:
    // One optimal set of operations choosing k = 3 indices in each operation is:
    // Operation 1: Flip indices [0, 1, 3]. s changes from "0101" to "1000".
    // Operation 2: Flip indices [1, 2, 3]. s changes from "1000" to "1111".
    // Thus, the minimum number of operations is 2.
    fmt.Println(minOperations("0101", 3)) // 2
    // Example 3:
    // Input: s = "101", k = 2
    // Output: -1
    // Explanation:
    // Since k = 2 and s has only one '0', it is impossible to flip exactly k indices to make all '1'. Hence, the answer is -1.
    fmt.Println(minOperations("101", 2)) // -1

    fmt.Println(minOperations("0000000000", 2)) // 5
    fmt.Println(minOperations("1111111111", 2)) // 0
    fmt.Println(minOperations("0000011111", 2)) // -1
    fmt.Println(minOperations("1111100000", 2)) // -1
    fmt.Println(minOperations("0101010101", 2)) // -1
    fmt.Println(minOperations("1010101010", 2)) // -1
}