package main

import (
	"fmt"
	"sort"
)

/**
15. 3Sum
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.

Note:
The solution set must not contain duplicate triplets.

Example:
Given array nums = [-1, 0, 1, 2, -1, -4],
A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

 */


// 解法一 最优解，双指针 + 排序
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

// 解法二
func threeSum1(nums []int) [][]int {
	res := [][]int{}
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res
}

// best solution
func threeSumBest(nums []int) [][]int {
	var ans [][]int
	if len(nums) < 3 {
		return ans
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		for k, j := i+1, len(nums)-1; k < j; {
			n := nums[i] + nums[k] + nums[j]
			if n == 0 {
				ans = append(ans, []int{nums[i], nums[k], nums[j]})
				p := k
				for p < j && nums[p] == nums[k] {
					p++
				}
				k = p
			} else if n > 0 {
				j--
			} else {
				k++
			}
		}
	}
	return ans
}

func main() {
	fmt.Printf("threeSum([]int{-1, 0, 1, 2, -1, -4}) = %v\n",threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf("threeSum1([]int{-1, 0, 1, 2, -1, -4}) = %v\n",threeSum1([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf("threeSumBest([]int{-1, 0, 1, 2, -1, -4}) = %v\n",threeSumBest([]int{-1, 0, 1, 2, -1, -4}))
}
