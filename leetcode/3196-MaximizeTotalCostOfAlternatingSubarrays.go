package main

// 3196. Maximize Total Cost of Alternating Subarrays
// You are given an integer array nums with length n.

// The cost of a subarray nums[l..r], where 0 <= l <= r < n, is defined as:
//     cost(l, r) = nums[l] - nums[l + 1] + ... + nums[r] * (−1)r − l

// Your task is to split nums into subarrays such that the total cost of the subarrays is maximized, ensuring each element belongs to exactly one subarray.

// Formally, if nums is split into k subarrays, where k > 1, at indices i1, i2, ..., ik − 1, where 0 <= i1 < i2 < ... < ik - 1 < n - 1, then the total cost will be:
//     cost(0, i1) + cost(i1 + 1, i2) + ... + cost(ik − 1 + 1, n − 1)

// Return an integer denoting the maximum total cost of the subarrays after splitting the array optimally.

// Note: If nums is not split into subarrays, i.e. k = 1, the total cost is simply cost(0, n - 1).

// Example 1:
// Input: nums = [1,-2,3,4]
// Output: 10
// Explanation:
// One way to maximize the total cost is by splitting [1, -2, 3, 4] into subarrays [1, -2, 3] and [4]. The total cost will be (1 + 2 + 3) + 4 = 10.

// Example 2:
// Input: nums = [1,-1,1,-1]
// Output: 4
// Explanation:
// One way to maximize the total cost is by splitting [1, -1, 1, -1] into subarrays [1, -1] and [1, -1]. The total cost will be (1 + 1) + (1 + 1) = 4.

// Example 3:
// Input: nums = [0]
// Output: 0
// Explanation:
// We cannot split the array further, so the answer is 0.

// Example 4:
// Input: nums = [1,-1]
// Output: 2
// Explanation:
// Selecting the whole array gives a total cost of 1 + 1 = 2, which is the maximum.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"

func maximumTotalCost(nums []int) int64 {
    dp := make([][2][2]int64, len(nums) + 1)
    for index := range dp {
        for i := 0; i < 2; i++ {
            for j := 0; j < 2; j++ {
                dp[index][i][j] = int64(-1e15)
            }
        }
    }
    var dfs func(index int, isStart int, sign int) int64
    dfs = func(index int, isStart int, sign int) int64 {
        if index == len(nums) { return 0 }
        if dp[index][isStart][sign] != int64(-1e15) { return dp[index][isStart][sign] }
        res := int64(-1e15)
        if isStart == 0 {
            res = max(res, int64(nums[index]) + dfs(index + 1, 1 - isStart, 1 - sign)) // take
        } else {
            if sign == 0 {
                res = max(res, int64(nums[index]) + dfs(index + 1, 1 - isStart, 1 - sign)) // take
                res = max(res, dfs(index, 1 - isStart, 1 - sign)) // not take
            } else {
                res = max(res, int64(-nums[index]) + dfs(index + 1, 1 - isStart, 1 - sign)) // take
                res = max(res, dfs(index, 1 - isStart, 1 - sign)) // not take
            }
        }
        dp[index][isStart][sign] = res
        return dp[index][isStart][sign]
    }
    return dfs(0, 0, 0)
}

func maximumTotalCost1(nums []int) int64 {
    facts := make([]int, len(nums) + 1)
    facts[1] = nums[0]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(nums); i++ {
        facts[i+1] = max(facts[i] + nums[i], facts[i-1] + nums[i-1] - nums[i])
    }  
    return int64(facts[len(nums)])
}

func main() {
    // Example 1:
    // Input: nums = [1,-2,3,4]
    // Output: 10
    // Explanation:
    // One way to maximize the total cost is by splitting [1, -2, 3, 4] into subarrays [1, -2, 3] and [4]. The total cost will be (1 + 2 + 3) + 4 = 10.
    fmt.Println(maximumTotalCost([]int{1,-2,3,4})) // 10
    // Example 2:
    // Input: nums = [1,-1,1,-1]
    // Output: 4
    // Explanation:
    // One way to maximize the total cost is by splitting [1, -1, 1, -1] into subarrays [1, -1] and [1, -1]. The total cost will be (1 + 1) + (1 + 1) = 4.
    fmt.Println(maximumTotalCost([]int{1,-1,1,-1})) // 4
    // Example 3:
    // Input: nums = [0]
    // Output: 0
    // Explanation:
    // We cannot split the array further, so the answer is 0.
    fmt.Println(maximumTotalCost([]int{0})) // 0
    // Example 4:
    // Input: nums = [1,-1]
    // Output: 2
    // Explanation:
    // Selecting the whole array gives a total cost of 1 + 1 = 2, which is the maximum.
    fmt.Println(maximumTotalCost([]int{1,-1})) // 2

    fmt.Println(maximumTotalCost1([]int{1,-2,3,4})) // 10
    fmt.Println(maximumTotalCost1([]int{1,-1,1,-1})) // 4
    fmt.Println(maximumTotalCost1([]int{0})) // 0
    fmt.Println(maximumTotalCost1([]int{1,-1})) // 2
}