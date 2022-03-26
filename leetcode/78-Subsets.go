package main

/**
78. Subsets
Given an integer array nums of unique elements, return all possible subsets (the power set).
The solution set must not contain duplicate subsets. Return the solution in any order.

Constraints:

	1 <= nums.length <= 10
	-10 <= nums[i] <= 10
	All the numbers of nums are unique.

Example 1:

	Input: nums = [1,2,3]
	Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:

	Input: nums = [0]
	Output: [[],[0]]

解题思路:
	DFS 暴力枚举
 */

import (
	"fmt"
	"sort"
)

// 解法一 dfs 递归
func subsets(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	for k := 0; k <= len(nums); k++ {
		generateSubsets(nums, k, 0, c, &res)
	}
	return res
}

func generateSubsets(nums []int, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// i will at most be n - (k - c.size()) + 1
	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		c = append(c, nums[i])
		generateSubsets(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}

// 解法二
func subsets1(nums []int) [][]int {
	res := make([][]int, 1)
	sort.Ints(nums)
	for i := range nums {
		for _, org := range res {
			clone := make([]int, len(org), len(org)+1)
			copy(clone, org)
			clone = append(clone, nums[i])
			res = append(res, clone)
		}
	}
	return res
}

// 解法三：位运算的方法 best solution
func subsets2(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	sum := 1 << uint(len(nums))
	for i := 0; i < sum; i++ {
		stack := []int{}
		tmp := i // i 从 000...000 到 111...111
		fmt.Printf("tmp = %v\n",tmp)
		for j := len(nums) - 1; j >= 0; j-- { // 遍历 i 的每一位
			if tmp & 1 == 1 {
				stack = append([]int{nums[j]}, stack...)
				fmt.Printf("stack = %v\n",stack)
			}
			tmp >>= 1
		}
		res = append(res, stack)
	}
	return res
}

func main() {
	fmt.Printf("subsets([]int{1,2,3}) = %v\n",subsets([]int{1,2,3}))
	fmt.Printf("subsets([]int{0}) = %v\n",subsets([]int{0}))

	fmt.Printf("subsets1([]int{1,2,3}) = %v\n",subsets1([]int{1,2,3}))
	fmt.Printf("subsets1([]int{0}) = %v\n",subsets1([]int{0}))

	fmt.Printf("subsets2([]int{1,2,3}) = %v\n",subsets2([]int{1,2,3}))
	fmt.Printf("subsets2([]int{0}) = %v\n",subsets2([]int{0}))
}
