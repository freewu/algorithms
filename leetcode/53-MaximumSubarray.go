package main

import (
	"fmt"
	"math"
)

/*
53. Maximum Subarray
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.
More practice:
If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

Example 1:

	Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
	Output: 6
	Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:

	Input: nums = [1]
	Output: 1

Example 3:

	Input: nums = [5,4,-1,7,8]
	Output: 23

*/

func maxSubArray1(nums []int) int {
	var l = len(nums)
	var max = 0
	array := make([][]int, l)

	for i := 0; i < l; i++ {
		subArray := make([]int, l)
		subArray[0] = nums[i]
		for j := i + 1; j < l; j++ {
			subArray[i] += nums[j]
			if subArray[i] > max {
				max = subArray[i]
			}
		}
		array[i] = subArray
	}

	// 输出
	for i := range array {
		for j := range array[i] {
			fmt.Printf("%v ", array[i][j])
		}
		fmt.Println()
	}
	return max
}

func maxSubArray(nums []int) int {
	var max = float64(math.MinInt64)
	var sum float64 = 0
	for i := range nums {
		// 如果累加值都小于当前值，从当前值开始
		sum = math.Max((sum + float64(nums[i])), float64(nums[i]))
		max = math.Max(max, sum)
	}
	return int(max)
}

// 解法一 DP
func maxSubArrayDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 { // 如果值大于 0 则累加
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 解法二 模拟
func maxSubArray2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, res, p := nums[0], 0, 0
	for p < len(nums) {
		res += nums[p]
		if res > maxSum {
			maxSum = res
		}
		if res < 0 {
			res = 0
		}
		p++
	}
	return maxSum
}

// best solution
func maxSubArrayBest(nums []int) int {
	length := len(nums)
	fast, slow :=0,0
	result, sum := math.MinInt32,0
	for fast < length && slow < length {
		sum = sum + nums[fast]
		if sum > result {
			result = sum
		}
		for sum< 0 {
			sum = sum - nums[slow]
			slow++
		}
		fast++
	}
	return result
}

func main() {
	fmt.Printf("maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Printf("maxSubArray([]int{1}) = %v\n",maxSubArray([]int{1})) // 1
	fmt.Printf("maxSubArray([]int{5,4,-1,7,8}) = %v\n",maxSubArray([]int{5,4,-1,7,8})) // 23

	fmt.Printf("maxSubArray1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSubArray1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Printf("maxSubArray1([]int{1}) = %v\n",maxSubArray1([]int{1})) // 1
	fmt.Printf("maxSubArray1([]int{5,4,-1,7,8}) = %v\n",maxSubArray1([]int{5,4,-1,7,8})) // 23

	fmt.Printf("maxSubArrayDP([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSubArrayDP([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Printf("maxSubArrayDP([]int{1}) = %v\n",maxSubArrayDP([]int{1})) // 1
	fmt.Printf("maxSubArrayDP([]int{5,4,-1,7,8}) = %v\n",maxSubArrayDP([]int{5,4,-1,7,8})) // 23

	fmt.Printf("maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Printf("maxSubArray2([]int{1}) = %v\n",maxSubArray2([]int{1})) // 1
	fmt.Printf("maxSubArray2([]int{5,4,-1,7,8}) = %v\n",maxSubArray2([]int{5,4,-1,7,8})) // 23

	fmt.Printf("maxSubArrayBest([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSubArrayBest([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Printf("maxSubArrayBest([]int{1}) = %v\n",maxSubArrayBest([]int{1})) // 1
	fmt.Printf("maxSubArrayBest([]int{5,4,-1,7,8}) = %v\n",maxSubArrayBest([]int{5,4,-1,7,8})) // 23
}
