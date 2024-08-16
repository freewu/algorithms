package main

// LCR 104. 组合总和 Ⅳ
// 给定一个由 不同 正整数组成的数组 nums ，和一个目标整数 target 。
// 请从 nums 中找出并返回总和为 target 的元素组合的个数。
// 数组中的数字可以在一次排列中出现任意次，但是顺序不同的序列被视作不同的组合。

// 题目数据保证答案符合 32 位整数范围。

// 示例 1：
// 输入：nums = [1,2,3], target = 4
// 输出：7
// 解释：
// 所有可能的组合为：
// (1, 1, 1, 1)
// (1, 1, 2)
// (1, 2, 1)
// (1, 3)
// (2, 1, 1)
// (2, 2)
// (3, 1)
// 请注意，顺序不同的序列被视作不同的组合。

// 示例 2：
// 输入：nums = [9], target = 3
// 输出：0
 
// 提示：
//     1 <= nums.length <= 200
//     1 <= nums[i] <= 1000
//     nums 中的所有元素 互不相同
//     1 <= target <= 1000

// 进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？

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
    fmt.Println(combinationSum4([]int{1,2,3}, 4)) // 7
    // Example 2:
    // Input: nums = [9], target = 3
    // Output: 0
    fmt.Println(combinationSum4([]int{9}, 3)) // 0
}