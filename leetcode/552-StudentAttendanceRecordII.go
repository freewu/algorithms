package main

// 552. Student Attendance Record II
// An attendance record for a student can be represented as a string where each character signifies whether the student was absent, late, or present on that day. 
// The record only contains the following three characters:
//     'A': Absent.
//     'L': Late.
//     'P': Present.

// Any student is eligible for an attendance award if they meet both of the following criteria:
//     The student was absent ('A') for strictly fewer than 2 days total.
//     The student was never late ('L') for 3 or more consecutive days.

// Given an integer n, return the number of possible attendance records of length n that make a student eligible for an attendance award. 
// The answer may be very large, so return it modulo 10^9 + 7.

// Example 1:
// Input: n = 2
// Output: 8
// Explanation: There are 8 records with length 2 that are eligible for an award:
// "PP", "AP", "PA", "LP", "PL", "AL", "LA", "LL"
// Only "AA" is not eligible because there are 2 absences (there need to be fewer than 2).

// Example 2:
// Input: n = 1
// Output: 3

// Example 3:
// Input: n = 10101
// Output: 183236316
 
// Constraints:
//     1 <= n <= 10^5

import "fmt"

// 假设没有缺勤，可被奖励的出勤记录必然满足P.../LP.../LLP...，且三种前缀解不存在重复。因此不存在缺勤的解共有:
//     f(n) = f(n-1) + f(n-2) + f(n-3)
// 再把 A 往每个位置插入，可行解的数量为:
//     \sum_{i=0}^{n-1}{f(i) * f(n-i-1)}
func checkRecord(n int) int {
    if n == 1 {
        return 3
    }
    d, mod := make([]int, n + 1), 1_000_000_007
    d[0], d[1], d[2] = 1, 2, 4
    for i := 3; i <= n; i++ {
        d[i] = (d[i - 1] + d[i -2 ] + d[i - 3]) % mod
    }
    res := d[n]
    for i := 0; i < n; i++ {
        res = (res + d[i] * d[n - i - 1]) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: 8
    // Explanation: There are 8 records with length 2 that are eligible for an award:
    // "PP", "AP", "PA", "LP", "PL", "AL", "LA", "LL"
    // Only "AA" is not eligible because there are 2 absences (there need to be fewer than 2).
    fmt.Println(checkRecord(2)) // 8 "PP", "AP", "PA", "LP", "PL", "AL", "LA", "LL"
    // Example 2:
    // Input: n = 1
    // Output: 3
    fmt.Println(checkRecord(1)) // 3 "A", "P", "L"
    // Example 3:
    // Input: n = 10101
    // Output: 183236316
    fmt.Println(checkRecord(10101)) // 183236316
}