package main

// 254. Factor Combinations
// Numbers can be regarded as the product of their factors.
//     For example, 8 = 2 x 2 x 2 = 2 x 4.

// Given an integer n, return all possible combinations of its factors. You may return the answer in any order.
// Note that the factors should be in the range [2, n - 1].

// Example 1:
// Input: n = 1
// Output: []

// Example 2:
// Input: n = 12
// Output: [[2,6],[3,4],[2,2,3]]

// Example 3:
// Input: n = 37
// Output: []
 
// Constraints:
//     1 <= n <= 10^7

import "fmt"
import "math"

func getFactors(n int) [][]int {
    res, factor := [][]int{}, []int{} // 先遍历求出所有因子 结果是因子一个组和 问题变为深度优先遍历数组
    if n == 1 {
        return res
    }
    for i := n - 1; i > 1; i-- { // 因子从大到小排序 正好遍历 index 递增
        if n % i == 0 {
            factor = append(factor, i)
        }
    }
    path, m := []int{}, len(factor)
    var dfs func(index, target int)
    dfs = func(index, target int) {
        if target == 1 {
            res = append(res, append([]int{}, path...))
            return
        }
        for i := index; i < m; i++ {
            if target % factor[i] == 0 { // 因子还能被剩余数整除
                path = append(path, factor[i])
                dfs(i, target / factor[i]) // 这里i不需要+1,因为允许重复数
                path = path[:len(path)-1]
            }
        }
    }
    dfs(0, n)
    return res
}

func getFactors1(n int) [][]int {
    res := [][]int{}
    var backTrack func(n int, temp []int, cur int, ans *[][]int) 
    backTrack = func(n int, temp []int, cur int, ans *[][]int) {
        if cur == n {
            if len(temp) > 1 {
                *ans = append(*ans, temp)
            }
            return
        }
        i := 2
        if len(temp) > 0 {
            i = temp[len(temp)-1]
        }
        for i <= int(math.Pow(float64(n/cur), 0.5)) {
            if n % (cur*i) == 0 {
                t := make([]int, len(temp))
                copy(t, temp)
                t = append(t, i)
                backTrack(n, t, cur*i, ans)
            }
            i++
        }
        t := make([]int, len(temp))
        copy(t, temp)
        t = append(t, n/cur)
        backTrack(n, t, n, ans)
    }
    backTrack(n, []int{}, 1, &res)
    return res
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: []
    fmt.Println(getFactors(1)) // []
    // Example 2:
    // Input: n = 12
    // Output: [[2,6],[3,4],[2,2,3]]
    fmt.Println(getFactors(12)) // [[2,6],[3,4],[2,2,3]]
    // Example 3:
    // Input: n = 37
    // Output: []
    fmt.Println(getFactors(37)) //  []
    fmt.Println(getFactors(4)) // [[2,2]]

    fmt.Println(getFactors1(1)) // []
    fmt.Println(getFactors1(12)) // [[2,6],[3,4],[2,2,3]]
    fmt.Println(getFactors1(37)) //  []
    fmt.Println(getFactors1(4)) // [[2,2]]
}