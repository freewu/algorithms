package main

// LCR 081. 组合总和
// 给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。
// candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是不同的。 
// 对于给定的输入，保证和为 target 的唯一组合数少于 150 个。

// 示例 1：
// 输入: candidates = [2,3,6,7], target = 7
// 输出: [[7],[2,2,3]]

// 示例 2：
// 输入: candidates = [2,3,5], target = 8
// 输出: [[2,2,2,2],[2,3,3],[3,5]]

// 示例 3：
// 输入: candidates = [2], target = 1
// 输出: []

// 示例 4：
// 输入: candidates = [1], target = 1
// 输出: [[1]]

// 示例 5：
// 输入: candidates = [1], target = 2
// 输出: [[1,1]]

// 提示：
//     1 <= candidates.length <= 30
//     1 <= candidates[i] <= 200
//     candidate 中的每个元素都是独一无二的。
//     1 <= target <= 500

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