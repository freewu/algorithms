package main

// 3952. Maximum Total Value of Covered Indices
// You are given an integer array nums of length n and a binary string s of length n, 
// where s[i] == '1' means index i initially contains a token and s[i] == '0' means it does not.

// You may perform the following operation any number of times:
//     1. Choose a token currently located at index i, where i > 0, such that this token has not been moved before.
//     2. Move this token from index i to index i - 1.

// An index is considered covered if it contains a token after all moves.

// Return an integer denoting the maximum total value of nums at the covered indices after optimally performing the operations.

// Example 1:
// Input: nums = [9,2,6,1], s = "0101"
// Output: 15
// Explanation:
// Initially, indices 1 and 3 contain tokens.
// Move the token from index 3 to index 2.
// Move the token from index 1 to index 0.
// The covered indices are [0, 2], so the total value is nums[0] + nums[2] = 9 + 6 = 15.

// Example 2:
// Input: nums = [5,1,4], s = "001"
// Output: 4
// Explanation:
// Initially, only index 2 contains a token.
// It is optimal to leave the token at index 2.
// The covered index is [2], so the total value is nums[2] = 4.

// Example 3:
// Input: nums = [9,3,5], s = "011"
// Output: 14
// Explanation:
// Initially, indices 1 and 2 contain tokens.
// Move the token from index 1 to index 0.
// The covered indices are [0, 2], so the total value is nums[0] + nums[2] = 9 + 5 = 14.

// Constraints:
//     1 <= n == nums.length == s.length <= 10^5
//     1 <= nums[i] <= 10^5
//     ​​​​​​​s[i] is either '0' or '1'

import "fmt"

func maxTotal(nums []int, s string) int64 {
    sum, curr := 0, 0
    for i, v := range nums {
        if s[i] == '0' {
            curr = v
        } else if curr < v {
            sum += v
        } else {
            sum += curr
            curr = v
        }
    }
    return int64(sum)
}

func maxTotal1(nums []int, s string) int64 {
    dp := make([][2]int64,2) // 0 移动， 1不移动
    for i, v :=range nums{
        if i == 0 {
            if s[i] == '1'{
                dp[0][0] = -1 << 61 // i % 2
                dp[0][1] =  int64(v)
            }
            continue
        }
        if s[i] == '0' {
            dp[i%2][0] = max(dp[(i+1) % 2][0], dp[(i+1) % 2][1])
            dp[i%2][1] = max(dp[(i+1) % 2][0], dp[(i+1) % 2][1])
        } else {
            if s[i-1] == '0' {
                dp[i%2][0] = dp[(i+1) % 2][0] + int64(nums[i-1])
                dp[i%2][1] = dp[(i+1) % 2][0] + int64(v)
            } else {
                dp[i%2][0] = dp[(i+1) % 2][0] + int64(nums[i-1])
                dp[i%2][1] = max(dp[(i+1) % 2][0],dp[(i+1) % 2][1]) + int64(v)
            }
        }
    }
    return max(dp[(len(nums)-1) % 2][0], dp[(len(nums)-1) % 2][1])
}

func main() {
    // Example 1:
    // Input: nums = [9,2,6,1], s = "0101"
    // Output: 15
    // Explanation:
    // Initially, indices 1 and 3 contain tokens.
    // Move the token from index 3 to index 2.
    // Move the token from index 1 to index 0.
    // The covered indices are [0, 2], so the total value is nums[0] + nums[2] = 9 + 6 = 15.
    fmt.Println(maxTotal([]int{9,2,6,1}, "0101")) // 15
    // Example 2:
    // Input: nums = [5,1,4], s = "001"
    // Output: 4
    // Explanation:
    // Initially, only index 2 contains a token.
    // It is optimal to leave the token at index 2.
    // The covered index is [2], so the total value is nums[2] = 4.
    fmt.Println(maxTotal([]int{5,1,4}, "001")) // 4
    // Example 3:
    // Input: nums = [9,3,5], s = "011"
    // Output: 14
    // Explanation:
    // Initially, indices 1 and 2 contain tokens.
    // Move the token from index 1 to index 0.
    // The covered indices are [0, 2], so the total value is nums[0] + nums[2] = 9 + 5 = 14.
    fmt.Println(maxTotal([]int{9,3,5}, "011")) // 14

    fmt.Println(maxTotal([]int{1,2,3,4,5,6,7,8,9}, "000000000")) // 0
    fmt.Println(maxTotal([]int{1,2,3,4,5,6,7,8,9}, "111111111")) // 45
    fmt.Println(maxTotal([]int{9,8,7,6,5,4,3,2,1}, "000000000")) // 0
    fmt.Println(maxTotal([]int{9,8,7,6,5,4,3,2,1}, "111111111")) // 45

    fmt.Println(maxTotal1([]int{5,1,4}, "001")) // 4
    fmt.Println(maxTotal1([]int{9,3,5}, "011")) // 14
    fmt.Println(maxTotal1([]int{1,2,3,4,5,6,7,8,9}, "000000000")) // 0
    fmt.Println(maxTotal1([]int{1,2,3,4,5,6,7,8,9}, "111111111")) // 45
    fmt.Println(maxTotal1([]int{9,8,7,6,5,4,3,2,1}, "000000000")) // 0
    fmt.Println(maxTotal1([]int{9,8,7,6,5,4,3,2,1}, "111111111")) // 45
}