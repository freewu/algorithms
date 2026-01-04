package main

// 3792. Sum of Increasing Product Blocks
// You are given an integer n.

// A sequence is formed as follows:
//     1. The 1st block contains 1.
//     2. The 2nd block contains 2 * 3.
//     3. The ith block is the product of the next i consecutive integers.

// Let F(n) be the sum of the first n blocks.

// Return an integer denoting F(n) modulo 10^9 + 7.

// Example 1:
// Input: n = 3
// Output: 127
// Explanation:​​​​​​​
// Block 1: 1
// Block 2: 2 * 3 = 6
// Block 3: 4 * 5 * 6 = 120
// F(3) = 1 + 6 + 120 = 127

// Example 2:
// Input: n = 7
// Output: 6997165
// Explanation:
// Block 1: 1
// Block 2: 2 * 3 = 6
// Block 3: 4 * 5 * 6 = 120
// Block 4: 7 * 8 * 9 * 10 = 5040
// Block 5: 11 * 12 * 13 * 14 * 15 = 360360
// Block 6: 16 * 17 * 18 * 19 * 20 * 21 = 39070080
// Block 7: 22 * 23 * 24 * 25 * 26 * 27 * 28 = 5967561600
// F(7) = 6006997207 % (109 + 7) = 6997165

// Constraints:
//     1 <= n <= 1000

import "fmt"

func sumOfBlocks(n int) int {
    res, begin, mod := 0, 1, 1_000_000_007
    for i := 1; i <= n; i++ {
        m := 1
        for j := begin; j < begin + i; j++ {
            m = m * j % mod
        }
        res = (res + m) % mod
        begin += i
    }
    return res % mod
}

var dp []int

func init() {
    const mod = 1_000_000_007
    nth := func (n int) (int, int) {
        start := n * (n - 1) / 2 + 1
        return start, start + n - 1
    }
    dp = make([]int, 1001)
    for i := 1; i <= 1000; i++ {
        start, end := nth(i)
        prod := 1
        for x := start; x <= end; x++ {
            prod = prod * x % mod
        }
        dp[i] = (dp[i - 1] + prod) % mod
    }
}

func sumOfBlocks1(n int) int {
    return dp[n]
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: 127
    // Explanation:​​​​​​​
    // Block 1: 1
    // Block 2: 2 * 3 = 6
    // Block 3: 4 * 5 * 6 = 120
    // F(3) = 1 + 6 + 120 = 127
    fmt.Println(sumOfBlocks(3)) // 127
    // Example 2:
    // Input: n = 7
    // Output: 6997165
    // Explanation:
    // Block 1: 1
    // Block 2: 2 * 3 = 6
    // Block 3: 4 * 5 * 6 = 120
    // Block 4: 7 * 8 * 9 * 10 = 5040
    // Block 5: 11 * 12 * 13 * 14 * 15 = 360360
    // Block 6: 16 * 17 * 18 * 19 * 20 * 21 = 39070080
    // Block 7: 22 * 23 * 24 * 25 * 26 * 27 * 28 = 5967561600
    // F(7) = 6006997207 % (109 + 7) = 6997165
    fmt.Println(sumOfBlocks(7)) // 6997165

    fmt.Println(sumOfBlocks(1)) // 1
    fmt.Println(sumOfBlocks(8)) // 103897425
    fmt.Println(sumOfBlocks(99)) // 518735420
    fmt.Println(sumOfBlocks(100)) // 523969427
    fmt.Println(sumOfBlocks(101)) // 91198711
    fmt.Println(sumOfBlocks(110)) // 469386446
    fmt.Println(sumOfBlocks(999)) // 30330820
    fmt.Println(sumOfBlocks(1000)) // 587219233

    fmt.Println(sumOfBlocks1(3)) // 127
    fmt.Println(sumOfBlocks1(7)) // 6997165
    fmt.Println(sumOfBlocks1(1)) // 1
    fmt.Println(sumOfBlocks1(8)) // 103897425
    fmt.Println(sumOfBlocks1(99)) // 518735420
    fmt.Println(sumOfBlocks1(100)) // 523969427
    fmt.Println(sumOfBlocks1(101)) // 91198711
    fmt.Println(sumOfBlocks1(110)) // 469386446
    fmt.Println(sumOfBlocks1(999)) // 30330820
    fmt.Println(sumOfBlocks1(1000)) // 587219233
}