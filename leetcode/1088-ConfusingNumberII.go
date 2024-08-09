package main

// 1088. Confusing Number II
// A confusing number is a number that when rotated 180 degrees becomes a different number with each digit valid.

// We can rotate digits of a number by 180 degrees to form new digits.
//     When 0, 1, 6, 8, and 9 are rotated 180 degrees, they become 0, 1, 9, 8, and 6 respectively.
//     When 2, 3, 4, 5, and 7 are rotated 180 degrees, they become invalid.

// Note that after rotating a number, we can ignore leading zeros.
//     For example, after rotating 8000, we have 0008 which is considered as just 8.

// Given an integer n, return the number of confusing numbers in the inclusive range [1, n].

// Example 1:
// Input: n = 20
// Output: 6
// Explanation: The confusing numbers are [6,9,10,16,18,19].
// 6 converts to 9.
// 9 converts to 6.
// 10 converts to 01 which is just 1.
// 16 converts to 91.
// 18 converts to 81.
// 19 converts to 61.

// Example 2:
// Input: n = 100
// Output: 19
// Explanation: The confusing numbers are [6,9,10,16,18,19,60,61,66,68,80,81,86,89,90,91,98,99,100].

// Constraints:
//     1 <= n <= 10^9

import "fmt"

func confusingNumberII(n int) int {
    tas := map[int]int{6: 9, 9: 6, 0: 0, 1: 1, 8: 8}
    res, k, m := 0, 1, []int{0, 1, 6, 8, 9}
    arr := make([]int, 0, 1000)
    arr = append(arr, 0)
    check := func (x int) bool {
        old, nw := x, 0
        for x > 0 {
            t := x % 10
            t = tas[t]
            nw = nw * 10 + t
            x /= 10
        }
        if nw == old {
            return false
        }
        return true
    }
    for k <= n {
        lenD := len(arr)
        for i := 0; i < lenD; i++ {
            for j := 1; j < len(m); j++ {
                now := k * m[j] + arr[i]
                if now <= n && check(now) {
                    res++
                }
                arr = append(arr, now)
            }
        }
        k *= 10
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 20
    // Output: 6
    // Explanation: The confusing numbers are [6,9,10,16,18,19].
    // 6 converts to 9.
    // 9 converts to 6.
    // 10 converts to 01 which is just 1.
    // 16 converts to 91.
    // 18 converts to 81.
    // 19 converts to 61.
    fmt.Println(confusingNumberII(20)) // 6
    // Example 2:
    // Input: n = 100
    // Output: 19
    // Explanation: The confusing numbers are [6,9,10,16,18,19,60,61,66,68,80,81,86,89,90,91,98,99,100].
    fmt.Println(confusingNumberII(100)) // 19
}