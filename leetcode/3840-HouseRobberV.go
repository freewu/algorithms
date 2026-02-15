package main

// 3840. House Robber V
// You are a professional robber planning to rob houses along a street. 
// Each house has a certain amount of money stashed and is protected by a security system with a color code.

// You are given two integer arrays nums and colors, both of length n, 
// where nums[i] is the amount of money in the ith house and colors[i] is the color code of that house.

// You cannot rob two adjacent houses if they share the same color code.

// Return the maximum amount of money you can rob.

// Example 1:
// Input: nums = [1,4,3,5], colors = [1,1,2,2]
// Output: 9
// Explanation:
// Choose houses i = 1 with nums[1] = 4 and i = 3 with nums[3] = 5 because they are non-adjacent.
// Thus, the total amount robbed is 4 + 5 = 9.

// Example 2:
// Input: nums = [3,1,2,4], colors = [2,3,2,2]
// Output: 8
// Explanation:
// Choose houses i = 0 with nums[0] = 3, i = 1 with nums[1] = 1, and i = 3 with nums[3] = 4.
// This selection is valid because houses i = 0 and i = 1 have different colors, and house i = 3 is non-adjacent to i = 1.
// Thus, the total amount robbed is 3 + 1 + 4 = 8.

// Example 3:
// Input: nums = [10,1,3,9], colors = [1,1,1,2]
// Output: 22
// Explanation:
// Choose houses i = 0 with nums[0] = 10, i = 2 with nums[2] = 3, and i = 3 with nums[3] = 9.
// This selection is valid because houses i = 0 and i = 2 are non-adjacent, and houses i = 2 and i = 3 have different colors.
// Thus, the total amount robbed is 10 + 3 + 9 = 22.
 
// Constraints:
//     1 <= n == nums.length == colors.length <= 10^5
//     1 <= nums[i], colors[i] <= 10^5

import "fmt"

func rob(nums []int, colors []int) int64 {
    n := len(nums)
    dp := make([]int, n + 1)
    dp[1] = nums[0]
    for i := 1; i < n; i++ {
        if colors[i] != colors[i-1] {
            dp[i+1] = dp[i] + nums[i]
        } else {
            dp[i+1] = max(dp[i-1] + nums[i], dp[i]) // 选或不选
        }
    }
    return int64(dp[n])
}

func rob1(nums, colors []int) int64 {
    n := len(nums)
    f0, f1 := 0, nums[0]
    for i := 1; i < n; i++ {
        if colors[i] != colors[i-1] {
            f0 = f1
            f1 += nums[i]
        } else {
            f0, f1 = f1, max(f0 + nums[i], f1)
        }
    }
    return int64(f1)
}


func main() {
    // Example 1:
    // Input: nums = [1,4,3,5], colors = [1,1,2,2]
    // Output: 9
    // Explanation:
    // Choose houses i = 1 with nums[1] = 4 and i = 3 with nums[3] = 5 because they are non-adjacent.
    // Thus, the total amount robbed is 4 + 5 = 9.
    fmt.Println(rob([]int{1,4,3,5}, []int{1,1,2,2})) // 9
    // Example 2:
    // Input: nums = [3,1,2,4], colors = [2,3,2,2]
    // Output: 8
    // Explanation:
    // Choose houses i = 0 with nums[0] = 3, i = 1 with nums[1] = 1, and i = 3 with nums[3] = 4.
    // This selection is valid because houses i = 0 and i = 1 have different colors, and house i = 3 is non-adjacent to i = 1.
    // Thus, the total amount robbed is 3 + 1 + 4 = 8.
    fmt.Println(rob([]int{3,1,2,4}, []int{2,3,2,2})) // 8
    // Example 3:
    // Input: nums = [10,1,3,9], colors = [1,1,1,2]
    // Output: 22
    // Explanation:
    // Choose houses i = 0 with nums[0] = 10, i = 2 with nums[2] = 3, and i = 3 with nums[3] = 9.
    // This selection is valid because houses i = 0 and i = 2 are non-adjacent, and houses i = 2 and i = 3 have different colors.
    // Thus, the total amount robbed is 10 + 3 + 9 = 22. 
    fmt.Println(rob([]int{10,1,3,9}, []int{1,1,1,2})) // 22

    fmt.Println(rob([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(rob([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 45
    fmt.Println(rob([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(rob([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(rob1([]int{1,4,3,5}, []int{1,1,2,2})) // 9
    fmt.Println(rob1([]int{3,1,2,4}, []int{2,3,2,2})) // 8
    fmt.Println(rob1([]int{10,1,3,9}, []int{1,1,1,2})) // 22
    fmt.Println(rob1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(rob1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 45
    fmt.Println(rob1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(rob1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 45
}