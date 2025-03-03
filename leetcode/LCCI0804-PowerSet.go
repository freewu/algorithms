package main

// 面试题 08.04. Power Set LCCI
// Write a method to return all subsets of a set. The elements in a set are pairwise distinct.

// Note: The result set should not contain duplicated subsets.

// Example:
// Input:  nums = [1,2,3]
// Output: 
// [
//     [3],
//     [1],
//     [2],
//     [1,2,3],
//     [1,3],
//     [2,3],
//     [1,2],
//     []
// ]

import "fmt"
import "slices"

func subsets(nums []int) [][]int {
    n := len(nums)
    res, path := make([][]int, 0, 1 << n), make([]int, 0, n) // 预分配空间
    var dfs func(i int)
    dfs = func(i int) {
        if i == n { // 子集构造完毕
            res = append(res, slices.Clone(path)) // 复制 path
            return
        }
        // 不选 nums[i]
        dfs(i + 1)
        // 选 nums[i]
        path = append(path, nums[i])
        dfs(i + 1)
        path = path[:len(path) - 1] // 恢复现场
    }
    dfs(0)
    return res
}

func subsets1(nums []int) [][]int {
    n := len(nums)
    res, path := make([][]int, 0, 1 << n), make([]int, 0, n) // 预分配空间
    var dfs func(i int)
    dfs = func(i int) {
        res = append(res, slices.Clone(path)) // 复制 path
        for j := i; j < n; j++ { // 枚举选择的数字
            path = append(path, nums[j])
            dfs(j + 1)
            path = path[:len(path)-1] // 恢复现场
        }
    }
    dfs(0)
    return res
}

func subsets2(nums []int) [][]int {
    res := make([][]int, 1<<len(nums))
    for i := range res { // 枚举全集 U 的所有子集 i
        for j, v := range nums {
            if i >> j & 1 == 1 { // j 在集合 i 中
                res[i] = append(res[i], v)
            }
        }
    }
    return res
}

func main() {
    // Example:
    // Input:  nums = [1,2,3]
    // Output: 
    // [
    //     [3],
    //     [1],
    //     [2],
    //     [1,2,3],
    //     [1,3],
    //     [2,3],
    //     [1,2],
    //     []
    // ]
    fmt.Println(subsets([]int{1,2,3})) // [[] [3] [2] [2 3] [1] [1 3] [1 2] [1 2 3]]
    fmt.Println(subsets([]int{1,2})) // [[] [2] [1] [1 2]]
    fmt.Println(subsets([]int{1,1,1})) // [[] [1] [1] [1 1] [1] [1 1] [1 1] [1 1 1]]

    fmt.Println(subsets1([]int{1,2,3})) // [[] [3] [2] [2 3] [1] [1 3] [1 2] [1 2 3]]
    fmt.Println(subsets1([]int{1,2})) // [[] [2] [1] [1 2]]
    fmt.Println(subsets1([]int{1,1,1})) // [[] [1] [1] [1 1] [1] [1 1] [1 1] [1 1 1]]

    fmt.Println(subsets2([]int{1,2,3})) // [[] [3] [2] [2 3] [1] [1 3] [1 2] [1 2 3]]
    fmt.Println(subsets2([]int{1,2})) // [[] [2] [1] [1 2]]
    fmt.Println(subsets2([]int{1,1,1})) // [[] [1] [1] [1 1] [1] [1 1] [1 1] [1 1 1]]
}