package main

// LCR 093. 最长的斐波那契子序列的长度
// 如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：
//     n >= 3
//     对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2}

// 给定一个严格递增的正整数数组形成序列 arr ，找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回  0 。
// （回想一下，子序列是从原序列  arr 中派生出来的，它从 arr 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。
// 例如， [3, 5, 8] 是 [3, 4, 5, 6, 7, 8] 的一个子序列）


// 示例 1：
// 输入: arr = [1,2,3,4,5,6,7,8]
// 输出: 5
// 解释: 最长的斐波那契式子序列为 [1,2,3,5,8] 。

// 示例 2：
// 输入: arr = [1,3,7,11,12,14,18]
// 输出: 3
// 解释: 最长的斐波那契式子序列有 [1,11,12]、[3,11,14] 以及 [7,11,18] 。

// 提示：
//     3 <= arr.length <= 1000
//     1 <= arr[i] < arr[i + 1] <= 10^9


import "fmt"

func lenLongestFibSubseq(arr []int) int {
    res, set := 0, make(map[int]bool)
    for _, v := range arr {
        set[v] = true
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            a, b := arr[i], arr[j]
            c := a + b
            l := 2
            for set[c] {
                a = b
                b = c
                c = a + b
                l++
                res = max(res, l)
            }
        }
    }
    return res
}

func lenLongestFibSubseq1(arr []int) int {
    n := len(arr)
    res, set, dp := 0, make(map[int]int), make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for j := 0; j < n; j++ {
        k, v:= arr[j], j
        set[k] = v
        for i := 0; i < j; i++ {
            k, ok := set[arr[j] - arr[i]]
            if arr[j] - arr[i] < arr[i] && ok {
                dp[i][j] = dp[k][i] + 1
                res = max(res, dp[i][j])
            } else {
                dp[i][j] = 2
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [1,2,3,4,5,6,7,8]
    // Output: 5
    // Explanation: The longest subsequence that is fibonacci-like: [1,2,3,5,8].
    fmt.Println(lenLongestFibSubseq([]int{1,2,3,4,5,6,7,8})) // 5
    // Example 2:
    // Input: arr = [1,3,7,11,12,14,18]
    // Output: 3
    // Explanation: The longest subsequence that is fibonacci-like: [1,11,12], [3,11,14] or [7,11,18].
    fmt.Println(lenLongestFibSubseq([]int{1,3,7,11,12,14,18})) // 3

    fmt.Println(lenLongestFibSubseq1([]int{1,2,3,4,5,6,7,8})) // 5
    fmt.Println(lenLongestFibSubseq1([]int{1,3,7,11,12,14,18})) // 3
}