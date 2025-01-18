package main

// 3287. Find the Maximum Sequence Value of Array
// You are given an integer array nums and a positive integer k.

// The value of a sequence seq of size 2 * x is defined as:
//     (seq[0] OR seq[1] OR ... OR seq[x - 1]) XOR (seq[x] OR seq[x + 1] OR ... OR seq[2 * x - 1]).

// Return the maximum value of any subsequence of nums having size 2 * k.

// Example 1:
// Input: nums = [2,6,7], k = 1
// Output: 5
// Explanation:
// The subsequence [2, 7] has the maximum value of 2 XOR 7 = 5.

// Example 2:
// Input: nums = [4,2,5,6,7], k = 2
// Output: 2
// Explanation:
// The subsequence [4, 5, 6, 7] has the maximum value of (4 OR 5) XOR (6 OR 7) = 2.

// Constraints:
//     2 <= nums.length <= 400
//     1 <= nums[i] < 27
//     1 <= k <= nums.length / 2

import "fmt"

func maxValue(nums []int, k int) int {
    reverse := func(arr []int) {
        for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    findORs := func(nums []int, k int) []map[int]bool {
        dp, prev := make([]map[int]bool, 0), make([]map[int]bool, k + 1)
        for i := 0; i <= k; i++ {
            prev[i] = make(map[int]bool)
        }
        prev[0][0] = true
        for i := 0; i < len(nums); i++ {
            for j := min(k - 1, i + 1); j >= 0; j-- {
                for x := range prev[j] {
                    prev[j + 1][x | nums[i]] = true
                }
            }
            current := make(map[int]bool)
            for key := range prev[k] {
                current[key] = true
            }
            dp = append(dp, current)
        }
        return dp
    }
    arr1 := findORs(nums, k)
    reverse(nums)
    arr2 := findORs(nums, k)
    res := 0
    for i := k - 1; i < len(nums) - k; i++ {
        for a := range arr1[i] {
            for b := range arr2[len(nums) - i - 2] {
                res = max(res, a ^ b)
            }
        }
    }
    return res
}

func maxValue1(nums []int, k int) int {
    n := len(nums)
    dp := make([][][]bool, n + 1)
    for i := range dp {
        dp[i] = make([][]bool, k + 1)
        for j := range dp[i] {
            dp[i][j] = make([]bool, 1 << 7)
        }
    }
    dp[n][0][0] = true
    for i := n - 1; i >= 0; i-- {
        copy(dp[i][0], dp[i+1][0])
        for j := 1; j <= k; j++ {
            copy(dp[i][j], dp[i+1][j])
            for mask := 0; mask < 1 << 7; mask++ {
                if dp[i + 1][j - 1][mask] {
                    dp[i][j][mask | nums[i]] = true
                }
            }
        }
    }
    pre := make([][][]bool, n + 1)
    for i := range pre {
        pre[i] = make([][]bool, k + 1)
        for j := range pre[i] {
            pre[i][j] = make([]bool, 1 << 7)
        }
    }
    pre[0][0][0] = true
    for i := 1; i <= n; i++ {
        copy(pre[i][0], pre[i - 1][0])
        for j := 1; j <= k; j++ {
            copy(pre[i][j], pre[i - 1][j])
            for mask := 0; mask < 1 << 7; mask++ {
                if pre[i - 1][j - 1][mask] {
                    pre[i][j][mask | nums[i - 1]] = true
                }
            }
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := 0
    for i := 0; i < n-1; i++ {
        for left := 0; left < 1 << 7; left++ {
            if pre[i + 1][k][left] {
                for right := 0; right < 1 << 7; right++ {
                    if dp[i + 1][k][right] {
                        res = max(res, left ^ right)
                    }
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,6,7], k = 1
    // Output: 5
    // Explanation:
    // The subsequence [2, 7] has the maximum value of 2 XOR 7 = 5.
    fmt.Println(maxValue([]int{2,6,7}, 1)) // 5
    // Example 2:
    // Input: nums = [4,2,5,6,7], k = 2
    // Output: 2
    // Explanation:
    // The subsequence [4, 5, 6, 7] has the maximum value of (4 OR 5) XOR (6 OR 7) = 2.
    fmt.Println(maxValue([]int{4,2,5,6,7}, 2)) // 2

    fmt.Println(maxValue([]int{1,2,3,4,5,6,7,8,9}, 2)) // 15
    fmt.Println(maxValue([]int{9,8,7,6,5,4,3,2,1}, 2)) // 15

    fmt.Println(maxValue1([]int{2,6,7}, 1)) // 5
    fmt.Println(maxValue1([]int{4,2,5,6,7}, 2)) // 2
    fmt.Println(maxValue1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 15
    fmt.Println(maxValue1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 15
}