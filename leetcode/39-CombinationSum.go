package main

/**
39. Combination Sum
Given a set of candidate numbers (candidates) (without duplicates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sums to target.
The same repeated number may be chosen from candidates unlimited number of times.

Note:
	All numbers (including target) will be positive integers.
	The solution set must not contain duplicate combinations.

Example 1:

	Input: candidates = [2,3,6,7], target = 7,
	A solution set is:
	[
	  [7],
	  [2,2,3]
	]

Example 2:

	Input: candidates = [2,3,5], target = 8,
	A solution set is:
	[
	  [2,2,2,2],
	  [2,3,3],
	  [3,5]
	]

解题思路:
	递归组合
	抛弃不匹配的组合
	num[i] > target 直接跳出
 */

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findcombinationSum(candidates, target, 0, c, &res)
	return res
}

func findcombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	if target <= 0 {
		if target == 0 { // 如果 target = 0 说明 数组c的里值 可以累加成 target 值
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		// 如果 target 为 负数了 说明 数组c的里值 不能组合出 target 值
		return
	}
	for i := index; i < len(nums); i++ {
		if nums[i] > target { // 这里可以剪枝优化 // 如果值都大于目标值就没有处理的必要了
			break
		}
		c = append(c, nums[i])
		// 递归处理，值可以取多次
		// target - nums[i] 下次递归需要的值 num1 + ... + xxx = target
		findcombinationSum(nums, target - nums[i], i, c, res) // 注意这里迭代的时候 index 依旧不变，因为一个元素可以取多次
		c = c[:len(c)-1]
	}
}

// best solution
func combinationSumBest(candidates []int, target int) [][]int {
	combinations := [][]int{}
	var dfs func(int, []int, int)
	dfs = func(i int, current []int, total int) {
		if i >= len(candidates) || total > target {
			return
		}
		if total == target {
			combinations = append(combinations, append([]int{}, current...))
			return
		}
		current = append(current, candidates[i])
		dfs(i, current, total + candidates[i])
		if len(current) > 0 {
			current = current[0:len(current)-1] // pop
		}
		dfs(i+1, current, total)
	}
	dfs(0, []int{}, 0)
	return combinations
}

func main() {
	fmt.Printf("combinationSum([]int{2,3,6,7},7) = %v\n",combinationSum([]int{2,3,6,7},7))
	fmt.Printf("combinationSum([]int{2,3,5},8) = %v\n",combinationSum([]int{2,3,5},8))
	fmt.Printf("combinationSumBest([]int{2,3,6,7},7) = %v\n",combinationSumBest([]int{2,3,6,7},7))
	fmt.Printf("combinationSumBest([]int{2,3,5},8) = %v\n",combinationSumBest([]int{2,3,5},8))
}