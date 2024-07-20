package main

// 1186. Maximum Subarray Sum with One Deletion
// Given an array of integers, return the maximum sum for a non-empty subarray (contiguous elements) with at most one element deletion. 
// In other words, you want to choose a subarray and optionally delete one element from it so 
// that there is still at least one element left and the sum of the remaining elements is maximum possible.

// Note that the subarray needs to be non-empty after deleting one element.

// Example 1:
// Input: arr = [1,-2,0,3]
// Output: 4
// Explanation: Because we can choose [1, -2, 0, 3] and drop -2, thus the subarray [1, 0, 3] becomes the maximum value.

// Example 2:
// Input: arr = [1,-2,-2,3]
// Output: 3
// Explanation: We just choose [3] and it's the maximum sum.

// Example 3:
// Input: arr = [-1,-1,-1,-1]
// Output: -1
// Explanation: The final subarray needs to be non-empty. You can't choose [-1] and delete -1 from it, then get an empty subarray to make the sum equals to 0.

// Constraints:
//     1 <= arr.length <= 10^5
//     -10^4 <= arr[i] <= 10^4

import "fmt"

func maximumSum(nums []int) int {
    res, n := nums[0], len(nums)
    dp := make([][2]int, n) // dp[i][0] stores max sum ending at i without deletion, // dp[i][1] stores max sum ending at i with one deletion
    dp[0][0], dp[0][1] = nums[0], nums[0]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        dp[i][0] = max(nums[i], dp[i-1][0] + nums[i])
        dp[i][1] = max(dp[i-1][0], dp[i-1][1]+ nums[i])
        res = max(res, max(dp[i][0], dp[i][1]))
    }
    return res
}

func maximumSum1(arr []int) int {
    dp0, dp1, res := arr[0], 0, arr[0]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(arr); i++ {
        dp0, dp1 = max(dp0, 0) + arr[i], max(dp1 + arr[i], dp0)
        res = max(res, max(dp0, dp1))
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [1,-2,0,3]
    // Output: 4
    // Explanation: Because we can choose [1, -2, 0, 3] and drop -2, thus the subarray [1, 0, 3] becomes the maximum value.
    fmt.Println(maximumSum([]int{1,-2,0,3})) // 4
    // Example 2:
    // Input: arr = [1,-2,-2,3]
    // Output: 3
    // Explanation: We just choose [3] and it's the maximum sum.
    fmt.Println(maximumSum([]int{1,-2,-2,3})) // 3
    // Example 3:
    // Input: arr = [-1,-1,-1,-1]
    // Output: -1
    // Explanation: The final subarray needs to be non-empty. You can't choose [-1] and delete -1 from it, then get an empty subarray to make the sum equals to 0.
    fmt.Println(maximumSum([]int{-1,-1,-1,-1})) // -1

    fmt.Println(maximumSum1([]int{1,-2,0,3})) // 4
    fmt.Println(maximumSum1([]int{1,-2,-2,3})) // 3
    fmt.Println(maximumSum1([]int{-1,-1,-1,-1})) // -1
}