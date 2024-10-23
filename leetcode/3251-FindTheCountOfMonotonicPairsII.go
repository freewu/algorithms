package main

// 3251. Find the Count of Monotonic Pairs II
// You are given an array of positive integers nums of length n.

// We call a pair of non-negative integer arrays (arr1, arr2) monotonic if:
//     The lengths of both arrays are n.
//     arr1 is monotonically non-decreasing, in other words, arr1[0] <= arr1[1] <= ... <= arr1[n - 1].
//     arr2 is monotonically non-increasing, in other words, arr2[0] >= arr2[1] >= ... >= arr2[n - 1].
//     arr1[i] + arr2[i] == nums[i] for all 0 <= i <= n - 1.

// Return the count of monotonic pairs.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [2,3,2]
// Output: 4
// Explanation:
// The good pairs are:
// ([0, 1, 1], [2, 2, 1])
// ([0, 1, 2], [2, 2, 0])
// ([0, 2, 2], [2, 1, 0])
// ([1, 2, 2], [1, 1, 0])

// Example 2:
// Input: nums = [5,5,5,5]
// Output: 126

// Constraints:
//     1 <= n == nums.length <= 2000
//     1 <= nums[i] <= 1000

import "fmt"

func countOfPairs(nums []int) int {
    res, n, mod := 0, len(nums), 1_000_000_007
    dp, pref := make([][1001]int, n), make([]int, 1002)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        t := make([]int, 1002)
        for j := 0; j <= nums[i]; j++ {
            if i == n - 1 {
                dp[i][j] = 1
            } else {
                if j > nums[i+1] {
                    t[j+1] = t[j] + dp[i][j]
                    continue
                }
                mx := max(j, nums[i+1] - (nums[i]-j))
                dp[i][j] = (dp[i][j] + pref[nums[i+1]+1] - pref[mx]) % mod
            }
            t[j+1] = t[j] + dp[i][j];
        }
        pref = t
    }
    for i := 0; i <= nums[0]; i++  {
        res = (res + dp[0][i]) % mod
    }
    return res
}

func countOfPairs1(nums []int) int {
    n, m, mod := len(nums), nums[len(nums) - 1], 1_000_000_007
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        m -= max(nums[i] - nums[i-1], 0)
        if m < 0 {
            return 0
        }
    }
    pow := func(x, n int) int {
        res := 1
        for ; n > 0; n /= 2 {
            if n % 2 > 0 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    mx := 3001 // MAX_N + MAX_M = 2000 + 1000 = 3000
    f, invF := make([]int, mx), make([]int, mx)    // f[i] = i!, // invF[i] = i!^-1
    init := func () {
        f[0] = 1
        for i := 1; i < mx; i++ {
            f[i] = f[i-1] * i % mod
        }
        invF[mx-1] = pow(f[mx-1], mod-2)
        for i := mx - 1; i > 0; i-- {
            invF[i-1] = invF[i] * i % mod
        }
    }
    init()
    comb := func (n, m int) int { return f[n] * invF[m] % mod * invF[n-m] % mod }
    return comb(m + n, n)
}

func main() {
    // Example 1:
    // Input: nums = [2,3,2]
    // Output: 4
    // Explanation:
    // The good pairs are:
    // ([0, 1, 1], [2, 2, 1])
    // ([0, 1, 2], [2, 2, 0])
    // ([0, 2, 2], [2, 1, 0])
    // ([1, 2, 2], [1, 1, 0])
    fmt.Println(countOfPairs([]int{2,3,2})) // 4
    // Example 2:
    // Input: nums = [5,5,5,5]
    // Output: 126
    fmt.Println(countOfPairs([]int{5,5,5,5})) // 126

    fmt.Println(countOfPairs1([]int{2,3,2})) // 4
    fmt.Println(countOfPairs1([]int{5,5,5,5})) // 126
}