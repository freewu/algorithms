package main

// 40. Combination Sum II
// Given a collection of candidate numbers (candidates) and a target number (target),
// find all unique combinations in candidates where the candidate numbers sums to target.
// Each number in candidates may only be used once in the combination.

// Note:
//     All numbers (including target) will be positive integers.
//     The solution set must not contain duplicate combinations.

// Example 1:
// Input: candidates = [10,1,2,7,6,1,5], target = 8,
// A solution set is:
// [
//     [1, 7],
//     [1, 2, 5],
//     [2, 6],
//     [1, 1, 6]
// ]

// Example 2:
// Input: candidates = [2,5,2,1,2], target = 5,
// A solution set is:
// [
//     [1,2,2],
//     [5]
// ]

// Constraints:
//     1 <= candidates.length <= 100
//     1 <= candidates[i] <= 50
//     1 <= target <= 30

// 解题思路:
//     递归组合
//     抛弃不匹配的组合
//     num[i] > target 直接跳出

import "fmt"
import "sort"

func combinationSum2(candidates []int, target int) [][]int {
    if len(candidates) == 0 {
        return [][]int{}
    }
    c, res := []int{}, [][]int{}
    sort.Ints(candidates) // 先排序
    var helper func(nums []int, target, index int, c []int) 
    helper = func(nums []int, target, index int, c []int) {
        if target == 0 {
            b := make([]int, len(c))
            copy(b, c)
            res = append(res, b)
            return
        }
        for i := index; i < len(nums); i++ {
            if i > index && nums[i] == nums[i-1] { // 这里是去重的关键逻辑,本次不取重复数字，下次循环可能会取重复数字
                continue
            }
            if target >= nums[i] {
                c = append(c, nums[i])
                helper(nums, target-nums[i], i+1, c)
                c = c[:len(c)-1]
            }
        }
    }
    helper(candidates, target, 0, c)
    return res
}

// best solution
func combinationSum2Best(candidates []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(candidates)
    var search func(candidates []int, idx int, target int, nums []int)
    search = func(candidates []int, idx int, target int, nums []int) {
        if target == 0 {
            res = append(res, nums)
            return
        }
        for i := idx; i < len(candidates); i++ {
            if i > idx && candidates[i] == candidates[i-1] {
                continue
            }
            candidate := candidates[i]
            if candidate > target {
                continue
            }
            newNums := []int{}
            for _, num := range nums {
                newNums = append(newNums, num)
            }
            newNums = append(newNums, candidate)
            search(candidates, i + 1, target - candidate, newNums)
        }
    }
    search(candidates, 0, target, []int{})
    return res
}

func main() {
    fmt.Printf("combinationSum2([]int{10,1,2,7,6,1,5},8) = %v\n",combinationSum2([]int{10,1,2,7,6,1,5},8))
    fmt.Printf("combinationSum2([]int{2,5,2,1,2},5) = %v\n",combinationSum2([]int{2,5,2,1,2},5))
    fmt.Printf("combinationSum2Best([]int{10,1,2,7,6,1,5},8) = %v\n",combinationSum2Best([]int{10,1,2,7,6,1,5},8))
    fmt.Printf("combinationSum2Best([]int{2,5,2,1,2},5) = %v\n",combinationSum2Best([]int{2,5,2,1,2},5))
}
