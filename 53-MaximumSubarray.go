package main

import (
	"fmt"
	"math"
)

/*
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.

More practice:
If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
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
	fmt.Println(maxSubArray1)
	var sum float64 = 0
	for i := range nums {
		// 如果累加值都小于当前值，从当前值开始
		sum = math.Max((sum + float64(nums[i])), float64(nums[i]))
		max = math.Max(max, sum)
	}
	return int(max)
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
}
