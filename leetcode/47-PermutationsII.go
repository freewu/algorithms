package main

/**
47. Permutations II
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example 1:

	Input: nums = [1,1,2]
	Output: [[1,1,2],[1,2,1],[2,1,1]]

Example 2:

	Input: nums = [1,2,3]
	Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 */

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	sort.Ints(nums)
	generatePermutation2(nums, 0, p, &res, &used)
	return res
}

func generatePermutation2(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] { // 去重判断
				continue
			}
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation2(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}

// best solution
func permuteUniqueBest(nums []int) [][]int {
	var result [][]int
	var solve func(comb []int, used []bool)
	sort.Ints(nums)
	solve = func(comb []int, used []bool) {
		if len(comb) == len(nums){
			a := append([]int{}, comb...)
			result = append(result, a)
			return
		}
		for i, v := range nums {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1]{
				continue
			}
			used[i] = true
			solve(append(comb, v), used)
			used[i] = false
		}
	}
	solve([]int{},make([]bool, len(nums)))
	return result
}

func main() {
	fmt.Printf("permuteUnique([]int{1,2,3}) = %v\n",permuteUnique([]int{1,2,3}))
	fmt.Printf("permuteUnique([]int{1,1,2}) = %v\n",permuteUnique([]int{1,1,2}))
	fmt.Printf("permuteUniqueBest([]int{1,2,3}) = %v\n",permuteUniqueBest([]int{1,2,3}))
	fmt.Printf("permuteUniqueBest([]int{1,1,2}) = %v\n",permuteUniqueBest([]int{1,1,2}))
}