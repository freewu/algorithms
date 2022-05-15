package main

import "fmt"

/**
213. House Robber II
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed. All houses at this place are arranged in a circle.
That means the first house is the neighbor of the last one.
Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house,
return the maximum amount of money you can rob tonight without alerting the police.

Example 1:

	Input: nums = [2,3,2]
	Output: 3
	Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.

Example 2:

	Input: nums = [1,2,3,1]
	Output: 4
	Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
	Total amount you can rob = 1 + 3 = 4.

Example 3:

	Input: nums = [1,2,3]
	Output: 3

Constraints:

	1 <= nums.length <= 100
	0 <= nums[i] <= 1000

# 解题思路 #
	198 题的加强版
	不过这次是在一个环形的街道中，即最后一个元素和第一个元素是邻居
	由于首尾是相邻的，所以在取了第一个房子以后就不能取第 n 个房子，
	那么就在 [0,n - 1] 的区间内找出总价值最多的解，然后再 [1,n] 的区间内找出总价值最多的解，两者取最大值即可。
 */

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	// 由于首尾是相邻的，所以需要对比 [0，n-1]、[1，n] 这两个区间的最大值
	return max(_rob(nums, 0, n-2), _rob(nums, 1, n-1))
}

func _rob(nums []int, start, end int) int {
	preMax := nums[start]
	curMax := max(preMax, nums[start+1])
	for i := start + 2; i <= end; i++ {
		tmp := curMax
		curMax = max(curMax, nums[i]+preMax)
		preMax = tmp

		fmt.Printf("i = %v,curMax = %v, preMax = %v \n",i,curMax,preMax)
	}
	return curMax
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// best solution
func robBest(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var p1, p2, max int
	for i := 0; i < len(nums) - 1; i += 1 {
		tmp := p1
		if p2 + nums[i] > p1 {
			tmp = p2 + nums[i]
		}
		p2 = p1
		p1 = tmp
	}
	max = p1
	p1, p2 = 0, 0
	for i := 1; i < len(nums); i += 1 {
		tmp := p1
		if p2 + nums[i] > p1 {
			tmp = p2 + nums[i]
		}
		p2 = p1
		p1 = tmp
	}
	if p1 > max {
		max = p1
	}
	return max
}

func main() {
	fmt.Printf("rob([]int{ 2,3,2 }) = %v\n",rob([]int{ 2,3,2 })) // 3
	fmt.Printf("rob([]int{ 1,2,3,1 }) = %v\n",rob([]int{ 1,2,3,1 })) // 4  1 + 3
	fmt.Printf("rob([]int{ 1,2,3 }) = %v\n",rob([]int{ 1,2,3 })) // 3

	fmt.Printf("robBest([]int{ 2,3,2 }) = %v\n",robBest([]int{ 2,3,2 })) // 3
	fmt.Printf("robBest([]int{ 1,2,3,1 }) = %v\n",robBest([]int{ 1,2,3,1 })) // 4
	fmt.Printf("robBest([]int{ 1,2,3 }) = %v\n",robBest([]int{ 1,2,3 })) // 3
}