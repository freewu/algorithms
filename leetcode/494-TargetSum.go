package main

// 494. Target Sum
// You are given an integer array nums and an integer target.
// You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.
//     For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1".

// Return the number of different expressions that you can build, which evaluates to target.

// Example 1:
// Input: nums = [1,1,1,1,1], target = 3
// Output: 5
// Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
// -1 + 1 + 1 + 1 + 1 = 3
// +1 - 1 + 1 + 1 + 1 = 3
// +1 + 1 - 1 + 1 + 1 = 3
// +1 + 1 + 1 - 1 + 1 = 3
// +1 + 1 + 1 + 1 - 1 = 3

// Example 2:
// Input: nums = [1], target = 1
// Output: 1
 
// Constraints:
//     1 <= nums.length <= 20
//     0 <= nums[i] <= 1000
//     0 <= sum(nums[i]) <= 1000
//     -1000 <= target <= 1000

import "fmt"

// dp
func findTargetSumWays(nums []int, target int) int {
    type key struct {
        current int
        index   int
    }
    footprint := make(map[key]int)
    var dp func(nums []int, l, index int, footprint map[key]int, current, target int) int 
    dp = func(nums []int, l, index int, footprint map[key]int, current, target int) int {
        if index == l-1 {
            if current == target {
                return 1
            }
            return 0
        }
        k := key{
            current: current,
            index:   index,
        }
        if w, ok := footprint[k]; ok {
            return w
        }
        ways := dp(nums, l, index+1, footprint, current+nums[index+1], target) + dp(nums, l, index+1, footprint, current-nums[index+1], target)
        footprint[k] = ways
        return ways
    }
    return dp(nums, len(nums), 0, footprint, nums[0], target) + dp(nums, len(nums), 0, footprint, -nums[0], target)
}

func findTargetSumWays1(nums []int, target int) int {
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
    }
    if (sum-target)%2 == 1 || sum < target {
        return 0
    }
    r := (sum - target) / 2
    dp := make([]int, r+1)
    dp[0] = 1
    for i := 1; i < len(dp); i++ {
        dp[i] = 0
    }
    for i := 0; i < len(nums); i++ {
        for j := r; j >= nums[i]; j-- {
            dp[j] = dp[j] + dp[j-nums[i]]
        }
    }
    return dp[r]
}

func main() {
    // Example 1:
    // Input: nums = [1,1,1,1,1], target = 3
    // Output: 5
    // Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
    // -1 + 1 + 1 + 1 + 1 = 3
    // +1 - 1 + 1 + 1 + 1 = 3
    // +1 + 1 - 1 + 1 + 1 = 3
    // +1 + 1 + 1 - 1 + 1 = 3
    // +1 + 1 + 1 + 1 - 1 = 3
    fmt.Println(findTargetSumWays([]int{1,1,1,1,1},3))
    // Example 2:
    // Input: nums = [1], target = 1
    // Output: 1
    fmt.Println(findTargetSumWays([]int{1},1))

    fmt.Println(findTargetSumWays1([]int{1,1,1,1,1},3))
    fmt.Println(findTargetSumWays1([]int{1},1))
}