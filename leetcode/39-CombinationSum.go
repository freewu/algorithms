package main

// 39. Combination Sum
// Given an array of distinct integers candidates and a target integer target, 
// return a list of all unique combinations of candidates where the chosen numbers sum to target. 
// You may return the combinations in any order.

// The same number may be chosen from candidates an unlimited number of times.
// Two combinations are unique if the frequency of at least one of the chosen numbers is different.

// The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

// Example 1:
// Input: candidates = [2,3,6,7], target = 7
// Output: [[2,2,3],[7]]
// Explanation:
// 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
// 7 is a candidate, and 7 = 7.
// These are the only two combinations.

// Example 2:
// Input: candidates = [2,3,5], target = 8
// Output: [[2,2,2,2],[2,3,3],[3,5]]

// Example 3:
// Input: candidates = [2], target = 1
// Output: []
 
// Constraints:
//     1 <= candidates.length <= 30
//     2 <= candidates[i] <= 40
//     All elements of candidates are distinct.
//     1 <= target <= 40

// # 解题思路:
//     递归组合
//     抛弃不匹配的组合
//     num[i] > target 直接跳出

import "fmt"
import "sort"

// dfs
func combinationSum(candidates []int, target int) [][]int {
    if len(candidates) == 0 {
        return [][]int{}
    }
    c, res := []int{}, [][]int{}
    sort.Ints(candidates)
    var dfs func (nums []int, target, index int, c []int, res *[][]int)
    dfs = func (nums []int, target, index int, c []int, res *[][]int) {
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
            dfs(nums, target - nums[i], i, c, res) // 注意这里迭代的时候 index 依旧不变，因为一个元素可以取多次
            c = c[:len(c)-1]
        }
    }
    dfs(candidates, target, 0, c, &res)
    return res
}


// best solution
func combinationSum1(candidates []int, target int) [][]int {
    res := [][]int{}
    var dfs func(int, []int, int)
    dfs = func(i int, current []int, total int) {
        if i >= len(candidates) || total > target {
            return
        }
        if total == target {
            res = append(res, append([]int{}, current...))
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
    return res
}

func combinationSum2(candidates []int, target int) [][]int {
    comb, res := []int{}, [][]int{}
    var dfs func(target, index int)
    dfs = func(target, index int) {
        if index == len(candidates) {
            return
        }
        if target == 0 {
            res = append(res, append([]int(nil), comb...))
            return
        }
        dfs(target, index + 1) // 直接跳过
        if target - candidates[index] >= 0 { // 选择当前数
            comb = append(comb, candidates[index])
            dfs(target-candidates[index], index)
            comb = comb[:len(comb) - 1]
        }
    }
    dfs(target, 0)
    return res
}

func main() {
    // Explanation:
    // 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
    // 7 is a candidate, and 7 = 7.
    // These are the only two combinations.
    fmt.Printf("combinationSum([]int{2,3,6,7},7) = %v\n",combinationSum([]int{2,3,6,7},7)) // [[2,2,3],[7]]
    fmt.Printf("combinationSum([]int{2,3,5},8) = %v\n",combinationSum([]int{2,3,5},8)) // [[2,2,2,2],[2,3,3],[3,5]]
    fmt.Printf("combinationSum([]int{2},1) = %v\n",combinationSum([]int{2},1)) // []

    fmt.Printf("combinationSum1([]int{2,3,6,7},7) = %v\n",combinationSum1([]int{2,3,6,7},7)) // [[2,2,3],[7]]
    fmt.Printf("combinationSum1([]int{2,3,5},8) = %v\n",combinationSum1([]int{2,3,5},8)) // [[2,2,2,2],[2,3,3],[3,5]]
    fmt.Printf("combinationSum1([]int{2},1) = %v\n",combinationSum1([]int{2},1)) // []
    
    fmt.Printf("combinationSum2([]int{2,3,6,7},7) = %v\n",combinationSum2([]int{2,3,6,7},7)) // [[2,2,3],[7]]
    fmt.Printf("combinationSum2([]int{2,3,5},8) = %v\n",combinationSum2([]int{2,3,5},8)) // [[2,2,2,2],[2,3,3],[3,5]]
    fmt.Printf("combinationSum2([]int{2},1) = %v\n",combinationSum2([]int{2},1)) // []
}