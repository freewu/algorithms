package main

// 873. Length of Longest Fibonacci Subsequence
// A sequence x1, x2, ..., xn is Fibonacci-like if:
//     n >= 3
//     xi + xi+1 == xi+2 for all i + 2 <= n

// Given a strictly increasing array arr of positive integers forming a sequence, 
// return the length of the longest Fibonacci-like subsequence of arr. 
// If one does not exist, return 0.

// A subsequence is derived from another sequence arr by deleting any number of elements (including none) from arr, 
// without changing the order of the remaining elements. 
// For example, [3, 5, 8] is a subsequence of [3, 4, 5, 6, 7, 8].

// Example 1:
// Input: arr = [1,2,3,4,5,6,7,8]
// Output: 5
// Explanation: The longest subsequence that is fibonacci-like: [1,2,3,5,8].

// Example 2:
// Input: arr = [1,3,7,11,12,14,18]
// Output: 3
// Explanation: The longest subsequence that is fibonacci-like: [1,11,12], [3,11,14] or [7,11,18].

// Constraints:
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

func lenLongestFibSubseq2(arr []int) int {
    res, n := 0, len(arr)
    mp := make(map[int]int, n)
    for i, v := range arr {
        mp[v] = i
    }
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for r := 2; r < n; r++ {
        for m := r - 1; m > 0 && arr[m] * 2 > arr[r]; m-- {
            if l, ok := mp[arr[r]-arr[m]]; ok {
                dp[m][r] = max(3, dp[l][m] + 1)
                res = max(res, dp[m][r])
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
    
    fmt.Println(lenLongestFibSubseq([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(lenLongestFibSubseq([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(lenLongestFibSubseq1([]int{1,2,3,4,5,6,7,8})) // 5
    fmt.Println(lenLongestFibSubseq1([]int{1,3,7,11,12,14,18})) // 3
    fmt.Println(lenLongestFibSubseq1([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(lenLongestFibSubseq1([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(lenLongestFibSubseq2([]int{1,2,3,4,5,6,7,8})) // 5
    fmt.Println(lenLongestFibSubseq2([]int{1,3,7,11,12,14,18})) // 3
    fmt.Println(lenLongestFibSubseq1([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(lenLongestFibSubseq1([]int{9,8,7,6,5,4,3,2,1})) // 0
}