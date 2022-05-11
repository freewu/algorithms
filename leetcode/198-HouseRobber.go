package main

import "fmt"

/**
198. House Robber
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed,
the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected
and it will automatically contact the police if two adjacent houses were broken into on the same night.
Given an integer array nums representing the amount of money of each house,
return the maximum amount of money you can rob tonight without alerting the police.

Example 1:

	Input: nums = [1,2,3,1]
	Output: 4
	Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
	Total amount you can rob = 1 + 3 = 4.

Example 2:

	Input: nums = [2,7,9,3,1]
	Output: 12
	Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
	Total amount you can rob = 2 + 9 + 1 = 12.

Constraints:

	1 <= nums.length <= 100
	0 <= nums[i] <= 400
 */

// 解法一 DP
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// dp[i] 代表抢 nums[0...i] 房子的最大价值
	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}

// 解法二 DP 优化辅助空间，把迭代的值保存在 2 个变量中
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	curMax, preMax := 0, 0
	for i := 0; i < n; i++ {
		tmp := curMax
		curMax = max(curMax, nums[i]+preMax)
		preMax = tmp
	}
	return curMax
}

// 不能连续，意味
// 解法三 模拟
func rob2(nums []int) int {
	// a 对于偶数位上的最大值的记录
	// b 对于奇数位上的最大值的记录
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		if i % 2 == 0 {
			a = max(a + nums[i], b)
		} else {
			b = max(a, b + nums[i])
		}
	}
	return max(a, b)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("rob([]int{ 1,2,3,1 }) = %v\n",rob([]int{ 1,2,3,1 })) // 4   1 + 3
	fmt.Printf("rob([]int{ 2,7,9,3,1 }) = %v\n",rob([]int{ 2,7,9,3,1 })) // 12  2 + 9 + 1 = 12.
	fmt.Printf("rob([]int{ 2,7,9,3,1,10,3,1 }) = %v\n",rob([]int{ 2,7,9,3,1,10,3,1 })) // 22  2 + 9 + 10 + １

	fmt.Printf("rob1([]int{ 1,2,3,1 }) = %v\n",rob1([]int{ 1,2,3,1 })) // 4   1 + 3
	fmt.Printf("rob1([]int{ 2,7,9,3,1 }) = %v\n",rob1([]int{ 2,7,9,3,1 })) // 12  2 + 9 + 1 = 12.
	fmt.Printf("rob([]int{ 2,7,9,3,1,10,3,1 }) = %v\n",rob1([]int{ 2,7,9,3,1,10,3,1 })) // 22  2 + 9 + 10 + １

	fmt.Printf("rob2([]int{ 1,2,3,1 }) = %v\n",rob2([]int{ 1,2,3,1 })) // 4   1 + 3
	fmt.Printf("rob3([]int{ 2,7,9,3,1 }) = %v\n",rob2([]int{ 2,7,9,3,1 })) // 12  2 + 9 + 1 = 12.
	fmt.Printf("rob([]int{ 2,7,9,3,1,10,3,1 }) = %v\n",rob2([]int{ 2,7,9,3,1,10,3,1 })) // 22  2 + 9 + 10 + １
}
