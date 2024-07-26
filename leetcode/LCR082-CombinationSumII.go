package main

// LCR 082. 组合总和 II
// 给定一个可能有重复数字的整数数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。 

// 示例 1:
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 输出:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]

// 示例 2:
// 输入: candidates = [2,5,2,1,2], target = 5,
// 输出:
// [
// [1,2,2],
// [5]
// ]

// 提示:
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
