package main

// 3339. Find the Number of K-Even Arrays
// You are given three integers n, m, and k.

// An array arr is called k-even if there are exactly k indices such that, for each of these indices i (0 <= i < n - 1):
//     (arr[i] * arr[i + 1]) - arr[i] - arr[i + 1] is even.

// Return the number of possible k-even arrays of size n where all elements are in the range [1, m].

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 3, m = 4, k = 2
// Output: 8
// Explanation:
// The 8 possible 2-even arrays are:
// [2, 2, 2]
// [2, 2, 4]
// [2, 4, 2]
// [2, 4, 4]
// [4, 2, 2]
// [4, 2, 4]
// [4, 4, 2]
// [4, 4, 4]

// Example 2:
// Input: n = 5, m = 1, k = 0
// Output: 1
// Explanation:
// The only 0-even array is [1, 1, 1, 1, 1].

// Example 3:
// Input: n = 7, m = 7, k = 5
// Output: 5832

// Constraints:
//     1 <= n <= 750
//     0 <= k <= n - 1
//     1 <= m <= 1000

import "fmt"

func countOfArrays(n int, m int, k int) int {
    facts := make([][][2]int, n)
    for i := range facts {
        facts[i] = make([][2]int, k + 1)
        for j := range facts[i] {
            facts[i][j] = [2]int{ -1, -1 }
        }
    }
    const mod = 1_000_000_007
    cnt0, cnt1 := m / 2, m - (m / 2)
    var dfs func(i, j, k int) int
    dfs = func(i, j, k int) int {
        if j < 0 { return 0 }
        if i >= n {
            if j == 0 { return 1 }
            return 0
        }
        if facts[i][j][k] != -1 { return facts[i][j][k] }
        a, b := cnt1 * dfs(i + 1, j, 1) % mod, cnt0 * dfs(i + 1, j- (k & 1^1), 0) % mod
        facts[i][j][k] = (a + b) % mod
        return facts[i][j][k]
    }
    return dfs(0, k, 1)
}

func countOfArrays1(n int, m int, k int) int {
    const mod = 1_000_000_007
    cnt0, cnt1 := m / 2, m - (m / 2)
    facts := make([][2]int, k + 1)
    facts[0][1] = 1
    for i := 0; i < n; i++ {
        temp := make([][2]int, k + 1)
        for j := 0; j <= k; j++ {
            if j > 0 {
                temp[j][0] = (cnt0 * (facts[j][1] + facts[j-1][0]))  % mod
            } else {
                temp[j][0] = (cnt0 * facts[j][1])  % mod
            }
            temp[j][1] = (cnt1 * (facts[j][0] + facts[j][1]) % mod) % mod
        }
        facts = temp
    }
    return (facts[k][0] + facts[k][1]) % mod
}

func main() {
    // Example 1:
    // Input: n = 3, m = 4, k = 2
    // Output: 8
    // Explanation:
    // The 8 possible 2-even arrays are:
    // [2, 2, 2]
    // [2, 2, 4]
    // [2, 4, 2]
    // [2, 4, 4]
    // [4, 2, 2]
    // [4, 2, 4]
    // [4, 4, 2]
    // [4, 4, 4]
    fmt.Println(countOfArrays(3, 4, 2)) // 2
    // Example 2:
    // Input: n = 5, m = 1, k = 0
    // Output: 1
    // Explanation:
    // The only 0-even array is [1, 1, 1, 1, 1].
    fmt.Println(countOfArrays(5, 1, 0)) // 1
    // Example 3:
    // Input: n = 7, m = 7, k = 5
    // Output: 5832
    fmt.Println(countOfArrays(7, 7, 5)) // 5832

    fmt.Println(countOfArrays(1, 1, 0)) // 1
    fmt.Println(countOfArrays(750, 1000, 749)) // 333434247
    fmt.Println(countOfArrays(750, 1, 749)) // 0
    fmt.Println(countOfArrays(1, 1000, 0)) // 1000

    fmt.Println(countOfArrays1(3, 4, 2)) // 2
    fmt.Println(countOfArrays1(5, 1, 0)) // 1
    fmt.Println(countOfArrays1(7, 7, 5)) // 5832
    fmt.Println(countOfArrays1(1, 1, 0)) // 1
    fmt.Println(countOfArrays1(750, 1000, 749)) // 333434247
    fmt.Println(countOfArrays1(750, 1, 749)) // 0
    fmt.Println(countOfArrays1(1, 1000, 0)) // 1000
}