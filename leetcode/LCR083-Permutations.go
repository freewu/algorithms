package main

// LCR 083. 全排列
// 给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。

// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// 示例 2：
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]

// 示例 3：
// 输入：nums = [1]
// 输出：[[1]]

// 提示：
//     1 <= nums.length <= 6
//     -10 <= nums[i] <= 10
//     nums 中的所有整数 互不相同

// 解题思路:
//     给定一个没有重复数字的序列，返回其所有可能的全排列。

import "fmt"

func permute(nums []int) [][]int {
    used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
    if len(nums) == 0 {
        return res
    }
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
                used[i] = true
                p = append(p, nums[i])
                generatePermutation(nums, index + 1, p)
                p = p[:len(p)-1]
                used[i] = false
            }
        }
    }
    generatePermutation(nums, 0, p)
    return res
}

func main() {
    fmt.Printf("permute([]int{1,2,3}) = %v\n",permute([]int{1,2,3})) // [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
    fmt.Printf("permute([]int{0,1}) = %v\n",permute([]int{0,1})) // [[0,1],[1,0]]
    fmt.Printf("permute([]int{1}) = %v\n",permute([]int{1})) // [[1]]
}