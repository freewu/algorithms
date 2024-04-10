package main

// 740. Delete and Earn
// You are given an integer array nums. You want to maximize the number of points you get by performing the following operation any number of times:
//     Pick any nums[i] and delete it to earn nums[i] points. 
//     Afterwards, you must delete every element equal to nums[i] - 1 and every element equal to nums[i] + 1.

// Return the maximum number of points you can earn by applying the above operation some number of times.

// Example 1:
// Input: nums = [3,4,2]
// Output: 6
// Explanation: You can perform the following operations:
// - Delete 4 to earn 4 points. Consequently, 3 is also deleted. nums = [2].
// - Delete 2 to earn 2 points. nums = [].
// You earn a total of 6 points.

// Example 2:
// Input: nums = [2,2,3,3,3,4]
// Output: 9
// Explanation: You can perform the following operations:
// - Delete a 3 to earn 3 points. All 2's and 4's are also deleted. nums = [3,3].
// - Delete a 3 again to earn 3 points. nums = [3].
// - Delete a 3 once more to earn 3 points. nums = [].
// You earn a total of 9 points.
 
// Constraints:
//     1 <= nums.length <= 2 * 10^4
//     1 <= nums[i] <= 10^4

import "fmt"

// dp
func deleteAndEarn(nums []int) int {
    dp := make([]int,10003)
    for _,v := range nums {
        dp[v] +=v
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 3; i <= 10001; i++ {
        dp[i] = max(dp[i-1], max(dp[i-2], dp[i-3]) + dp[i])
    }
    return dp[10001]
}

func deleteAndEarn1(nums []int) int {
    mv, cnts := 0, make([]int, 10010)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 统计出 v 出现的次数 & 最大值
    for _, v := range nums {
        cnts[v]++
        mv = max(mv, v)
    }
    dp := make([][]int, mv + 1)
    for i, _ := range dp {
        dp[i] = make([]int, 2)
    }
    for i := 1; i <= mv; i++ {
        // 转成打家劫舍类的问题
        dp[i][1] = dp[i-1][0] + i * cnts[i]
        dp[i][0] = max(dp[i-1][1], dp[i-1][0])
    }
    return max(dp[mv][1], dp[mv][0])
}

func main() {
    // Explanation: You can perform the following operations:
    // - Delete 4 to earn 4 points. Consequently, 3 is also deleted. nums = [2].
    // - Delete 2 to earn 2 points. nums = [].
    // You earn a total of 6 points.
    fmt.Println(deleteAndEarn([]int{3,4,2})) // 6
    // Explanation: You can perform the following operations:
    // - Delete a 3 to earn 3 points. All 2's and 4's are also deleted. nums = [3,3].
    // - Delete a 3 again to earn 3 points. nums = [3].
    // - Delete a 3 once more to earn 3 points. nums = [].
    // You earn a total of 9 points.
    fmt.Println(deleteAndEarn([]int{2,2,3,3,3,4})) // 9

    fmt.Println(deleteAndEarn1([]int{3,4,2})) // 6
    fmt.Println(deleteAndEarn1([]int{2,2,3,3,3,4})) // 9
}