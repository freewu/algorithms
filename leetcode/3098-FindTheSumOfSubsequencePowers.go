package main

// 3098. Find the Sum of Subsequence Powers
// You are given an integer array nums of length n, and a positive integer k.
// The power of a subsequence is defined as the minimum absolute difference between any two elements in the subsequence.
// Return the sum of powers of all subsequences of nums which have length equal to k.
// Since the answer may be large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [1,2,3,4], k = 3
// Output: 4
// Explanation:
// There are 4 subsequences in nums which have length 3: [1,2,3], [1,3,4], [1,2,4], and [2,3,4]. The sum of powers is |2 - 3| + |3 - 4| + |2 - 1| + |3 - 4| = 4.

// Example 2:
// Input: nums = [2,2], k = 2
// Output: 0
// Explanation:
// The only subsequence in nums which has length 2 is [2,2]. The sum of powers is |2 - 2| = 0.

// Example 3:
// Input: nums = [4,3,-1], k = 2
// Output: 10
// Explanation:
// There are 3 subsequences in nums which have length 2: [4,3], [4,-1], and [3,-1]. The sum of powers is |4 - 3| + |4 - (-1)| + |3 - (-1)| = 10.

// Constraints:
//     2 <= n == nums.length <= 50
//     -10^8 <= nums[i] <= 10^8
//     2 <= k <= n

import "fmt"
import "sort"

func sumOfPowers(nums []int, k int) int {
    res, n, mod := 0, len(nums), 1_000_000_007
    sort.Ints(nums)
    dp := make(map[int]int)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(lastIdx, k, minDiff int, dp map[int]int, nums []int) int
    dfs = func(lastIdx, k, minDiff int, dp map[int]int, nums []int) int {
        if k == 0 {
            return minDiff
        }
        key := ((minDiff << 12) + (lastIdx << 6) + k)
        if val, ok := dp[key]; ok {
            return val
        }
        res := 0
        for i := lastIdx + 1; i <= len(nums)-k; i++ {
            res = (res + dfs(i, k-1, min(minDiff, nums[i]-nums[lastIdx]), dp, nums)) % mod
        }
        dp[key] = res
        return res
    }
    for i := 0; i <= n-k; i++ {
        res = (res + dfs(i, k-1, nums[n-1]-nums[0], dp, nums)) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4], k = 3
    // Output: 4
    // Explanation:
    // There are 4 subsequences in nums which have length 3: [1,2,3], [1,3,4], [1,2,4], and [2,3,4]. The sum of powers is |2 - 3| + |3 - 4| + |2 - 1| + |3 - 4| = 4.
    fmt.Println(sumOfPowers([]int{1,2,3,4}, 3)) // 4
    // Example 2:
    // Input: nums = [2,2], k = 2
    // Output: 0
    // Explanation:
    // The only subsequence in nums which has length 2 is [2,2]. The sum of powers is |2 - 2| = 0.
    fmt.Println(sumOfPowers([]int{0,0}, 2)) // 0
    // Example 3:
    // Input: nums = [4,3,-1], k = 2
    // Output: 10
    // Explanation:
    // There are 3 subsequences in nums which have length 2: [4,3], [4,-1], and [3,-1]. The sum of powers is |4 - 3| + |4 - (-1)| + |3 - (-1)| = 10.
    fmt.Println(sumOfPowers([]int{4,3,-1}, 2)) // 10

    fmt.Println(sumOfPowers([]int{-24, -921, 119, -291, -65, -628, 372, 274, 962, -592, -10, 67, -977, 85, -294, 349, -119, -846, -959, -79, -877, 833, 857, 44, 826, -295, -855, 554, -999, 759, -653, -423, -599, -928}, 19)) // 990202285
}