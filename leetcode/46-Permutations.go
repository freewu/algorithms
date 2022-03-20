package main

import "fmt"

/**
46. Permutations
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Constraints:

	1 <= nums.length <= 6
	-10 <= nums[i] <= 10
	All the integers of nums are unique.

Example 1:

	Input: nums = [1,2,3]
	Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:

	Input: nums = [0,1]
	Output: [[0,1],[1,0]]

Example 3:

	Input: nums = [1]
	Output: [[1]]

解题思路:
	给定一个没有重复数字的序列，返回其所有可能的全排列。
 */

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	generatePermutation(nums, 0, p, &res, &used)
	return res
}

func generatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}

func main() {
	fmt.Printf("permute([]int{1,2,3}) = %v\n",permute([]int{1,2,3}))
}