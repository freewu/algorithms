package main

/**
90. Subsets II
Given an integer array nums that may contain duplicates, return all possible subsets (the power set).
The solution set must not contain duplicate subsets. Return the solution in any order.

Constraints:

	1 <= nums.length <= 10
	-10 <= nums[i] <= 10

Example 1:

	Input: nums = [1,2,2]
	Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

Example 2:

	Input: nums = [0]
	Output: [[],[0]]

 */

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	sort.Ints(nums) // 这里是去重的关键逻辑
	for k := 0; k <= len(nums); k++ {
		generateSubsetsWithDup(nums, k, 0, c, &res)
	}
	return res
}

func generateSubsetsWithDup(nums []int, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// i will at most be n - (k - c.size()) + 1
	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		fmt.Printf("i = %v start = %v c = %v\n", i, start, c)
		if i > start && nums[i] == nums[i-1] { // 这里是去重的关键逻辑,本次不取重复数字，下次循环可能会取重复数字
			continue
		}
		c = append(c, nums[i])
		generateSubsetsWithDup(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}

// best solution
func subsetsWithDupBest(nums []int) [][]int {
	c := Counter(nums)
	res := [][]int{{}}
	for num, freq := range c {
		res2 := append([][]int{}, res...)
		for _, r := range res {
			for f := 1; f <= freq; f++ {
				r2 := append([]int{}, r...)
				for i := 0; i < f; i++ {
					r2 = append(r2, num)
				}
				res2 = append(res2, r2)
			}
		}
		res = res2
	}
	return res
}

func Counter(nums []int) map[int]int {
	mapped := map[int]int{}
	for _, num := range nums {
		mapped[num]++
	}
	return mapped
}

func main() {
	fmt.Printf("subsetsWithDup([]int{1,2,3}) = %v\n",subsetsWithDup([]int{1,2,3}))
	fmt.Printf("subsetsWithDup([]int{0}) = %v\n",subsetsWithDup([]int{0}))

	fmt.Printf("subsetsWithDupBest([]int{1,2,3}) = %v\n",subsetsWithDupBest([]int{1,2,3}))
	fmt.Printf("subsetsWithDupBest([]int{0}) = %v\n",subsetsWithDupBest([]int{0}))
}
