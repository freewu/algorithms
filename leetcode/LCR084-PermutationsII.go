package main

// LCR 084. 全排列 II 
// 给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。

// 示例 1：
// 输入：nums = [1,1,2]
// 输出：
// [[1,1,2],
//  [1,2,1],
//  [2,1,1]]

// 示例 2：
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// 提示：
//     1 <= nums.length <= 8
//     -10 <= nums[i] <= 10

import "fmt"
import "sort"

func permuteUnique(nums []int) [][]int {
    used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
    if len(nums) == 0 {
        return res
    }
    sort.Ints(nums)
    var generatePermutation func(nums []int, index int, p []int)
    generatePermutation = func(nums []int, index int, p []int) {
        if index == len(nums) {
            temp := make([]int, len(p))
            copy(temp, p)
            res = append(res, temp)
            return
        }
        for i := 0; i < len(nums); i++ {
            if !used[i] {
                if i > 0 && nums[i] == nums[i-1] && !used[i-1] { // 去重判断
                    continue
                }
                used[i] = true
                p = append(p, nums[i])
                generatePermutation(nums, index+1, p)
                p = p[:len(p)-1]
                used[i] = false
            }
        }
    }
    generatePermutation(nums, 0, p)
    return res
}

func permuteUnique1(nums []int) [][]int {
    res := [][]int{}
    sort.Ints(nums)

    var solve func(comb []int, used []bool)
    solve = func(comb []int, used []bool) {
        if len(comb) == len(nums) {
            a := append([]int{}, comb...)
            res = append(res, a)
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
    return res
}

func main() {
    fmt.Printf("permuteUnique([]int{1,2,3}) = %v\n", permuteUnique([]int{1,2,3})) // [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
    fmt.Printf("permuteUnique([]int{1,1,2}) = %v\n", permuteUnique([]int{1,1,2})) // [[1,1,2], [1,2,1], [2,1,1]]
    fmt.Printf("permuteUnique1([]int{1,2,3}) = %v\n", permuteUnique1([]int{1,2,3})) // [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
    fmt.Printf("permuteUnique1([]int{1,1,2}) = %v\n", permuteUnique1([]int{1,1,2})) //[[1,1,2], [1,2,1], [2,1,1]]
}