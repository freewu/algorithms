package main

// 2501. Longest Square Streak in an Array
// You are given an integer array nums. 
// A subsequence of nums is called a square streak if:
//     The length of the subsequence is at least 2, and
//     after sorting the subsequence, each element (except the first element) is the square of the previous number.

// Return the length of the longest square streak in nums, or return -1 if there is no square streak.

// A subsequence is an array that can be derived from another array by deleting some
// or no elements without changing the order of the remaining elements.

// Example 1:
// Input: nums = [4,3,6,16,8,2]
// Output: 3
// Explanation: Choose the subsequence [4,16,2]. After sorting it, it becomes [2,4,16].
// - 4 = 2 * 2.
// - 16 = 4 * 4.
// Therefore, [4,16,2] is a square streak.
// It can be shown that every subsequence of length 4 is not a square streak.

// Example 2:
// Input: nums = [2,3,5,6,7]
// Output: -1
// Explanation: There is no square streak in nums so return -1.

// Constraints:
//     2 <= nums.length <= 10^5
//     2 <= nums[i] <= 10^5

import "fmt"
import "sort"
import "math"

// 哈希加排序
func longestSquareStreak(nums []int) int {
    sort.Ints(nums)
    res, n, mp := 0, len(nums), make(map[int]bool)
    for _, v := range nums {
        mp[v] = true
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        start := nums[i]
        t := 0
        for start * start <= nums[n - 1] { // 内层循环判断平方数是否还存在，条件是小于等于数组的最大值
            start *= start
            if _, ok := mp[start]; ok { 
                t++
                continue // 如果存在就继续平方
            }
            break // 不存在就要立即中止内部循环
        }
        res = max(t + 1, res)
    }
    if res == 1 { return -1 } // 如果不存在 方波长度至少为 2
    return res
}

func longestSquareStreak1(nums []int) int {
    sort.Ints(nums)
    res, dp := -1, make([]int, 100001)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        dp[v] = max(dp[v], 1)
        x := int(math.Sqrt(float64(v)))
        if x * x != v { continue }
        dp[v] = max(dp[v], dp[x] + 1)
        if dp[v] > 1 {
            res = max(res, dp[v])
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,3,6,16,8,2]
    // Output: 3
    // Explanation: Choose the subsequence [4,16,2]. After sorting it, it becomes [2,4,16].
    // - 4 = 2 * 2.
    // - 16 = 4 * 4.
    // Therefore, [4,16,2] is a square streak.
    // It can be shown that every subsequence of length 4 is not a square streak.
    fmt.Println(longestSquareStreak([]int{4,3,6,16,8,2})) // 3
    // Example 2:
    // Input: nums = [2,3,5,6,7]
    // Output: -1
    // Explanation: There is no square streak in nums so return -1.
    fmt.Println(longestSquareStreak([]int{2,3,5,6,7})) // -1

    fmt.Println(longestSquareStreak1([]int{4,3,6,16,8,2})) // 3
    fmt.Println(longestSquareStreak1([]int{2,3,5,6,7})) // -1
}