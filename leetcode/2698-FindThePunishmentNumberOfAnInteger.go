package main

// 2698. Find the Punishment Number of an Integer
// Given a positive integer n, return the punishment number of n.

// The punishment number of n is defined as the sum of the squares of all integers i such that:
//     1. 1 <= i <= n
//     2. The decimal representation of i * i can be partitioned into contiguous substrings 
//        such that the sum of the integer values of these substrings equals i.

// Example 1:
// Input: n = 10
// Output: 182
// Explanation: There are exactly 3 integers i that satisfy the conditions in the statement:
// - 1 since 1 * 1 = 1
// - 9 since 9 * 9 = 81 and 81 can be partitioned into 8 + 1.
// - 10 since 10 * 10 = 100 and 100 can be partitioned into 10 + 0.
// Hence, the punishment number of 10 is 1 + 81 + 100 = 182

// Example 2:
// Input: n = 37
// Output: 1478
// Explanation: There are exactly 4 integers i that satisfy the conditions in the statement:
// - 1 since 1 * 1 = 1. 
// - 9 since 9 * 9 = 81 and 81 can be partitioned into 8 + 1. 
// - 10 since 10 * 10 = 100 and 100 can be partitioned into 10 + 0. 
// - 36 since 36 * 36 = 1296 and 1296 can be partitioned into 1 + 29 + 6.
// Hence, the punishment number of 37 is 1 + 81 + 100 + 1296 = 1478

// Constraints:
//     1 <= n <= 1000

import "fmt"

func punishmentNumber(n int) int {
    res := 0
    var backtracking func(ts, p, i int) bool 
    backtracking = func(ts, p, i int) bool {
        if p <= 0 {
            if ts == i { return true }
            return false
        }
        for j := 10; j < 1000000; j *= 10 {
            if backtracking(ts + p % j, p / j, i) { return true }
        }
        return false
    }
    for i := 1; i <= n; i++ {
        if s := i * i; backtracking(0, s, i) { 
            res += s
        }
    }
    return res
}

func punishmentNumber1(n int) int {
    var canPartition func(num, target int) bool
    canPartition = func(num, target int) bool {
        if target < 0 || num < target { return false } // Invalid partition found
        if num == target { return true } // Valid partition found
        // Recursively check all partitions for a valid partition
        return  canPartition(num / 10,   target - (num % 10))  ||
                canPartition(num / 100,  target - (num % 100)) ||
                canPartition(num / 1000, target - (num % 1000))
    }
    res := 0
    for i := 1; i <= n; i++ {
        square := i * i
        if canPartition(square, i) {
            res += square
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 10
    // Output: 182
    // Explanation: There are exactly 3 integers i that satisfy the conditions in the statement:
    // - 1 since 1 * 1 = 1
    // - 9 since 9 * 9 = 81 and 81 can be partitioned into 8 + 1.
    // - 10 since 10 * 10 = 100 and 100 can be partitioned into 10 + 0.
    // Hence, the punishment number of 10 is 1 + 81 + 100 = 182
    fmt.Println(punishmentNumber(10)) // 182
    // Example 2:
    // Input: n = 37
    // Output: 1478
    // Explanation: There are exactly 4 integers i that satisfy the conditions in the statement:
    // - 1 since 1 * 1 = 1. 
    // - 9 since 9 * 9 = 81 and 81 can be partitioned into 8 + 1. 
    // - 10 since 10 * 10 = 100 and 100 can be partitioned into 10 + 0. 
    // - 36 since 36 * 36 = 1296 and 1296 can be partitioned into 1 + 29 + 6.
    // Hence, the punishment number of 37 is 1 + 81 + 100 + 1296 = 1478
    fmt.Println(punishmentNumber(37)) // 1478

    fmt.Println(punishmentNumber(1)) // 1
    fmt.Println(punishmentNumber(2)) // 1
    fmt.Println(punishmentNumber(8)) // 1
    fmt.Println(punishmentNumber(16)) // 182
    fmt.Println(punishmentNumber(512)) // 772866
    fmt.Println(punishmentNumber(999)) // 9804657
    fmt.Println(punishmentNumber(1000)) // 10804657

    fmt.Println(punishmentNumber1(10)) // 182
    fmt.Println(punishmentNumber1(37)) // 1478
    fmt.Println(punishmentNumber1(1)) // 1
    fmt.Println(punishmentNumber1(2)) // 1
    fmt.Println(punishmentNumber1(8)) // 1
    fmt.Println(punishmentNumber1(16)) // 182
    fmt.Println(punishmentNumber1(512)) // 772866
    fmt.Println(punishmentNumber1(999)) // 9804657
    fmt.Println(punishmentNumber1(1000)) // 10804657
}