package main

// LCR 102. 目标和
// 给定一个正整数数组 nums 和一个整数 target 。
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//     例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。

// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

// 示例 1：
// 输入：nums = [1,1,1,1,1], target = 3
// 输出：5
// 解释：一共有 5 种方法让最终目标和为 3 。
// -1 + 1 + 1 + 1 + 1 = 3
// +1 - 1 + 1 + 1 + 1 = 3
// +1 + 1 - 1 + 1 + 1 = 3
// +1 + 1 + 1 - 1 + 1 = 3
// +1 + 1 + 1 + 1 - 1 = 3

// 示例 2：
// 输入：nums = [1], target = 1
// 输出：1

// 提示：
//     1 <= nums.length <= 20
//     0 <= nums[i] <= 1000
//     0 <= sum(nums[i]) <= 1000
//     -1000 <= target <= 1000

import "fmt"

// dp
func findTargetSumWays(nums []int, target int) int {
    type key struct {
        current, index int
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