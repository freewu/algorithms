package main

import (
	"fmt"
	"math"
)

/**
209. Minimum Size Subarray Sum
Given an array of positive integers nums and a positive integer target,
return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than
or equal to target. If there is no such subarray, return 0 instead.

Example 1:

	Input: target = 7, nums = [2,3,1,2,4,3]
	Output: 2
	Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Example 2:

	Input: target = 4, nums = [1,4,4]
	Output: 1

Example 3:

	Input: target = 11, nums = [1,1,1,1,1,1,1,1]
	Output: 0

Constraints:

	1 <= target <= 10^9
	1 <= nums.length <= 10^5
	1 <= nums[i] <= 10^5

Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).

# 解题思路
	滑动窗口
	在滑动窗口 [i,j]之间不断往后移动，
	如果总和小于 s，就扩大右边界 j，不断加入右边的值，直到 sum > s，
	之后再缩小 i 的左边界，不断缩小直到 sum < s，这时候右边界又可以往右移动
 */

func minSubArrayLen(target int, nums []int) int {
	left, sum, res := 0, 0, len(nums)+1
	for right, v := range nums { // 从头开始向尾部
		sum += v // 累加
		for sum >= target { // 如果超过目标值
			res = min(res, right - left + 1)  // 取个最小的长度
			sum -= nums[left] // 减掉开头的值
			left++ // 向尾部走一下
		}
	}
	if res == len(nums) + 1 { // 所有值累加都达不到目标值 返回 0
		return 0
	}
	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// best solution
func minSubArrayLenBest(target int, nums []int) int {
	// minLength := len(nums)
	answer := math.MaxInt32
	left := 0
	sum := 0

	for i:= 0; i < len(nums); i++ {
		sum += nums[i]

		for sum >= target {
			answer = int(math.Min(float64(answer), float64(i + 1 - left)))
			sum -= nums[left]
			left++
		}
	}
	if answer != math.MaxInt32 {
		return int(answer)
	} else {
		return 0
	}
}

func main() {
	fmt.Printf("minSubArrayLen(7,[]int{2,3,1,2,4,3}) = %v\n",minSubArrayLen(7,[]int{2,3,1,2,4,3})) // 2  [4,3]
	fmt.Printf("minSubArrayLen(4,[]int{1,4,4}) = %v\n",minSubArrayLen(4,[]int{1,4,4})) // 1 [4]
	fmt.Printf("minSubArrayLen(11,[]int{1,1,1,1,1,1,1,1}) = %v\n",minSubArrayLen(11,[]int{1,1,1,1,1,1,1,1})) // 0

	fmt.Printf("minSubArrayLenBest(7,[]int{2,3,1,2,4,3}) = %v\n",minSubArrayLenBest(7,[]int{2,3,1,2,4,3})) // 2  [4,3]
	fmt.Printf("minSubArrayLenBest(4,[]int{1,4,4}) = %v\n",minSubArrayLenBest(4,[]int{1,4,4})) // 1 [4]
	fmt.Printf("minSubArrayLenBest(11,[]int{1,1,1,1,1,1,1,1}) = %v\n",minSubArrayLenBest(11,[]int{1,1,1,1,1,1,1,1})) // 0
}