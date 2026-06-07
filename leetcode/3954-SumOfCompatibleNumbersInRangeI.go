package main

// 3954. Sum of Compatible Numbers in Range I
// You are given two integers n and k.

// A positive integer x is called compatible if it satisfies both of the following conditions:
//     1. abs(n - x) <= k
//     2. (n & x) == 0

// Return the sum of all compatible integers x.

// Note:
//     1. Here, & denotes the bitwise AND operator.
//     2. The absolute difference between integers i and j is defined as abs(i - j).

// Example 1:
// Input: n = 2, k = 3
// Output: 10
// Explanation:
// The compatible integers are:
// x = 1, since abs(2 - 1) = 1 and 2 & 1 = 0.
// x = 4, since abs(2 - 4) = 2 and 2 & 4 = 0.
// x = 5, since abs(2 - 5) = 3 and 2 & 5 = 0.
// Thus, the answer is 1 + 4 + 5 = 10.

// Example 2:
// Input: n = 5, k = 1
// Output: 0
// Explanation:
// There are no compatible integers in the range [4, 6]. Thus, the answer is 0.

// Constraints:
//     1 <= n <= 100
//     1 <= k <= 100

import "fmt"

func sumOfGoodIntegers(n int, k int) int {
    res, start := 0, n - k
    if start < 1 {
        start = 1
    }
    for i := start; i <= n + k; i++ {
        if (n & i) == 0 {
            res += i
        }
    }
    return res
}

func sumOfGoodIntegers1(n int, k int) int {
    res := 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 1; i <= 200; i++ {
        if abs(n - i) <= k && (n & i) == 0 {
            res += i
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2, k = 3
    // Output: 10
    // Explanation:
    // The compatible integers are:
    // x = 1, since abs(2 - 1) = 1 and 2 & 1 = 0.
    // x = 4, since abs(2 - 4) = 2 and 2 & 4 = 0.
    // x = 5, since abs(2 - 5) = 3 and 2 & 5 = 0.
    // Thus, the answer is 1 + 4 + 5 = 10.
    fmt.Println(sumOfGoodIntegers(2, 3)) // 10
    // Example 2:
    // Input: n = 5, k = 1
    // Output: 0
    // Explanation:
    // There are no compatible integers in the range [4, 6]. Thus, the answer is 0.
    fmt.Println(sumOfGoodIntegers(5, 1)) // 0

    fmt.Println(sumOfGoodIntegers(1, 1)) // 2
    fmt.Println(sumOfGoodIntegers(1, 100)) // 2550
    fmt.Println(sumOfGoodIntegers(100, 1)) // 0
    fmt.Println(sumOfGoodIntegers(100, 100)) // 2480

    fmt.Println(sumOfGoodIntegers1(2, 3)) // 10
    fmt.Println(sumOfGoodIntegers1(5, 1)) // 0
    fmt.Println(sumOfGoodIntegers1(1, 1)) // 2
    fmt.Println(sumOfGoodIntegers1(1, 100)) // 2550
    fmt.Println(sumOfGoodIntegers1(100, 1)) // 0
    fmt.Println(sumOfGoodIntegers1(100, 100)) // 2480
}