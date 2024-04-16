package main

// 377. Combination Sum IV
// Given an array of distinct integers nums and a target integer target, 
// return the number of possible combinations that add up to target.

// The test cases are generated so that the answer can fit in a 32-bit integer.

// Example 1:
// Input: nums = [1,2,3], target = 4
// Output: 7
// Explanation:
// The possible combination ways are:
// (1, 1, 1, 1)
// (1, 1, 2)
// (1, 2, 1)
// (1, 3)
// (2, 1, 1)
// (2, 2)
// (3, 1)
// Note that different sequences are counted as different combinations.

// Example 2:
// Input: nums = [9], target = 3
// Output: 0
 
// Constraints:
//     1 <= nums.length <= 200
//     1 <= nums[i] <= 1000
//     All the elements of nums are unique.
//     1 <= target <= 1000
    
// Follow up: What if negative numbers are allowed in the given array? How does it change the problem? What limitation we need to add to the question to allow negative numbers?

import "fmt"

// dp
// dp[i] 为总和为 target = i 的组合总数。最终答案存在 dp[target] 中。状态转移方程为：
//           |  1 , i = 0
//  dp[i] =  | 
//           | ∑dp[i−j],i != 0
func combinationSum4(nums []int, target int) int {
    dp := make([]int, target+1)
    dp[0] = 1 // 1 , i = 0
    for i := 1; i <= target; i++ {
        for _, num := range nums {
            if i - num >= 0 {
                dp[i] += dp[i - num] // ∑dp[i−j],i != 0
            }
        }
    }
    return dp[target]
}

func main() {
    // The possible combination ways are:
    // (1, 1, 1, 1)
    // (1, 1, 2)
    // (1, 2, 1)
    // (1, 3)
    // (2, 1, 1)
    // (2, 2)
    // (3, 1)
    // Note that different sequences are counted as different combinations.
    fmt.Println(combinationSum4([]int{1,2,3}, 4)) // 7
    fmt.Println(combinationSum4([]int{9}, 3)) // 0
}