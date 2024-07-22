package main

// LCR 079. 子集
// 给定一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// 示例 2：
// 输入：nums = [0]
// 输出：[[],[0]]

// 提示：
//     1 <= nums.length <= 10
//     -10 <= nums[i] <= 10
//     nums 中的所有元素 互不相同

import "fmt"
import "sort"

// dfs 
func subsets(nums []int) [][]int {
    c, res := []int{}, [][]int{}
    var dfs func (nums []int, k, start int, c []int, res *[][]int)
    dfs = func (nums []int, k, start int, c []int, res *[][]int) {
        if len(c) == k {
            b := make([]int, len(c))
            copy(b, c)
            *res = append(*res, b)
            return
        }
        // i will at most be n - (k - c.size()) + 1
        for i := start; i < len(nums)-(k-len(c))+1; i++ {
            c = append(c, nums[i])
            dfs(nums, k, i+1, c, res)
            c = c[:len(c)-1]
        }
    }
    for k := 0; k <= len(nums); k++ {
        dfs(nums, k, 0, c, &res)
    }
    return res
}

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

// 位运算
func subsets2(nums []int) [][]int {
    if len(nums) == 0 {
        return nil
    }
    res, sum := [][]int{}, 1 << uint(len(nums))
    for i := 0; i < sum; i++ {
        stack := []int{}
        tmp := i // i 从 000...000 到 111...111
        for j := len(nums) - 1; j >= 0; j-- { // 遍历 i 的每一位
            if tmp & 1 == 1 {
                stack = append([]int{nums[j]}, stack...)
            }
            tmp >>= 1
        }
        res = append(res, stack)
    }
    return res
}

func main() {
    fmt.Printf("subsets([]int{1,2,3}) = %v\n",subsets([]int{1,2,3})) // [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
    fmt.Printf("subsets([]int{0}) = %v\n",subsets([]int{0})) // [[],[0]]

    fmt.Printf("subsets1([]int{1,2,3}) = %v\n",subsets1([]int{1,2,3})) // [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
    fmt.Printf("subsets1([]int{0}) = %v\n",subsets1([]int{0})) // [[],[0]]

    fmt.Printf("subsets2([]int{1,2,3}) = %v\n",subsets2([]int{1,2,3})) // [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
    fmt.Printf("subsets2([]int{0}) = %v\n",subsets2([]int{0})) // [[],[0]]
}
 