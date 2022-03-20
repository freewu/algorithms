package main

import "fmt"

/**
45. Jump Game II
Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Your goal is to reach the last index in the minimum number of jumps.
You can assume that you can always reach the last index.

Constraints:

	1 <= nums.length <= 1000
	0 <= nums[i] <= 10^5

Example 1:

	Input: nums = [2,3,1,1,4]
	Output: 2
	Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:

	Input: nums = [2,3,0,1,4]
	Output: 2

解题思路:
	给定一个非负整数数组，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。
	你的目标是使用最少的跳跃次数到达数组的最后一个位置。

 */
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	needChoose, canReach, step := 0, 0, 0
	for i, x := range nums {
		if i+x > canReach {
			canReach = i + x
			if canReach >= len(nums)-1 {
				return step + 1
			}
		}
		if i == needChoose {
			needChoose = canReach
			step++
		}
	}
	return step
}

// best solution
func jumpBest(nums []int) int {
	l := len(nums)
	c := 0

	if l == 1 {
		return c
	}
	for i := 0; i < l; {
		c++
		m, mj := 0, 0
		for j := i + 1; j < i + 1 + nums[i] && j < l; j++ {
			val := j - i + nums[j]
			if m <= val {
				m, mj = val, j
			}
			if j == l - 1 {
				return c
			}
		}
		i = mj
		if mj >= l - 1 {
			break
		} else if mj + nums[mj] >= l - 1{
			c++
			break
		}
	}
	return c
}

func main() {
	fmt.Printf("jump([]int{2,3,1,1,4}) = %v\n",jump([]int{2,3,1,1,4}))
	fmt.Printf("jump([]int{2,3,0,1,4}) = %v\n",jump([]int{2,3,0,1,4}))
	fmt.Printf("jumpBest([]int{2,3,1,1,4}) = %v\n",jumpBest([]int{2,3,1,1,4}))
	fmt.Printf("jumpBest([]int{2,3,0,1,4}) = %v\n",jumpBest([]int{2,3,0,1,4}))
}