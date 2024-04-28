package main

// 77. Combinations
// Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].
// You may return the answer in any order.

// Example 1:
// Input: n = 4, k = 2
// Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
// Explanation: There are 4 choose 2 = 6 total combinations.
// Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.

// Example 2:
// Input: n = 1, k = 1
// Output: [[1]]
// Explanation: There is 1 choose 1 = 1 total combination.

// Constraints:
//     1 <= n <= 20
//     1 <= k <= n

// 解题思路:
//     给定两个整数 n 和 k，返回 1 … n 中所有可能的 k 个数的组合。
//     DFS

import "fmt"

func combine(n int, k int) [][]int {
    if n <= 0 || k <= 0 || k > n {
        return [][]int{}
    }
    res, c := [][]int{}, []int{}
    var generateCombinations func(n, k, start int, c []int, res *[][]int)
    generateCombinations = func(n, k, start int, c []int, res *[][]int) {
        if len(c) == k { // 如果生成了符合要求长度的值 写入到结果集中
            b := make([]int, len(c))
            copy(b, c)
            *res = append(*res, b)
            return
        }
        for i := start; i <= n-(k-len(c))+1; i++ {
            c = append(c, i)
            generateCombinations(n, k, i+1, c, res)
            c = c[:len(c)-1]
        }
        return
    }
    generateCombinations(n, k, 1, c, &res)
    return res
}

func combine1(n int, k int) [][]int {
    res, arr := [][]int{}, []int{}
    for i := 1; i <= k; i++ {
        arr = append(arr, i)
    }
    res = append(res, arr)
    for j := arr[k-1] + 1; j <= n; j++ {
        dst := make([]int, len(arr))
        copy(dst, arr)
        dst[k-1] = j
        res = append(res, dst)
    }
    for i := k - 2; i >= 0; i-- {
        for _, a := range res {
            for j := a[i] + 1; j < a[i+1]; j++ {
                dst := make([]int, k)
                copy(dst, a)
                dst[i] = j
                res = append(res, dst)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, k = 2
    // Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
    // Explanation: There are 4 choose 2 = 6 total combinations.
    // Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.
    fmt.Printf("combine(4,2) = %v\n",combine(4,2)) // [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
    // Example 2:
    // Input: n = 1, k = 1
    // Output: [[1]]
    // Explanation: There is 1 choose 1 = 1 total combination.
    fmt.Printf("combine(1,1) = %v\n",combine(1,1)) // [[1]]

    fmt.Printf("combine1(4,2) = %v\n",combine1(4,2)) // [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
    fmt.Printf("combine1(1,1) = %v\n",combine1(1,1)) // [[1]]
}
