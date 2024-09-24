package main

// 1340. Jump Game V
// Given an array of integers arr and an integer d. 
// In one step you can jump from index i to index:
//     i + x where: i + x < arr.length and 0 < x <= d.
//     i - x where: i - x >= 0 and 0 < x <= d.

// In addition, you can only jump from index i to index j if arr[i] > arr[j] and arr[i] > arr[k] for all indices k between i 
// and j (More formally min(i, j) < k < max(i, j)).

// You can choose any index of the array and start jumping. Return the maximum number of indices you can visit.

// Notice that you can not jump outside of the array at any time.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/01/23/meta-chart.jpeg" />
// Input: arr = [6,4,14,6,8,13,9,7,10,6,12], d = 2
// Output: 4
// Explanation: You can start at index 10. You can jump 10 --> 8 --> 6 --> 7 as shown.
// Note that if you start at index 6 you can only jump to index 7. You cannot jump to index 5 because 13 > 9. You cannot jump to index 4 because index 5 is between index 4 and 6 and 13 > 9.
// Similarly You cannot jump from index 3 to index 2 or index 1.

// Example 2:
// Input: arr = [3,3,3,3,3], d = 3
// Output: 1
// Explanation: You can start at any index. You always cannot jump to any index.

// Example 3:
// Input: arr = [7,6,5,4,3,2,1], d = 1
// Output: 7
// Explanation: Start at index 0. You can visit all the indicies. 

// Constraints:
//     1 <= arr.length <= 1000
//     1 <= arr[i] <= 10^5
//     1 <= d <= arr.length

import "fmt"
import "slices"

func maxJumps(arr []int, d int) int {
    n, mp, res := len(arr), make(map[int]int), 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var solve func(idx int) int
    solve = func(idx int) int {
        if v, ok := mp[idx]; ok {  return v  }
        res := 1
        for j := idx + 1; j <= min(idx + d, n - 1) && arr[j] < arr[idx]; j++ {
            res = max(res, 1 + solve(j))
        }
        for j := idx - 1; j >= max(idx - d, 0) && arr[j] < arr[idx]; j-- {
            res = max(res, 1 + solve(j))
        }
        mp[idx] = res
        return mp[idx]
    }
    for i := 0; i < n; i++ {
        res = max(res, solve(i))
    }
    return res
}

func maxJumps1(arr []int, d int) int {
    if len(arr) == 1 { return 1 }
    res, n := 0, len(arr)
    memo := make([]int, n)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(i int) int
    dfs = func(i int) int {
        p := &memo[i]
        if *p != 0 {
            return *p
        }
        val := 1
        for j := i - 1; j >= 0 && j >= i-d; j-- { // left
            if arr[j] < arr[i] {
                val = max(val, dfs(j)+1)
            } else {
                break
            }
        }
        for j := i + 1; j < n && j <= i+d; j++ { // right
            if arr[j] < arr[i] {
                val = max(val, dfs(j)+1)
            } else {
                break
            }
        }
        *p = val
        return val
    }
    for i := 0; i < len(arr); i++ {
        res = max(res, dfs(i))
    }
    return res
}

func maxJumps2(arr []int, d int) int {
    n := len(arr)
    dp, sortedIndices:= make([]int, n), make([]int, n)
    for i := range arr {
        sortedIndices[i] = i
    }
    slices.SortFunc(sortedIndices, func(i, j int) int {
        return arr[i] - arr[j]
    })
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, i := range sortedIndices {
        dp[i] = 1
        // go forward
        for j := i + 1; j <= min(i+d, n-1); j++ {
            if arr[i] <= arr[j] {
                break
            }
            dp[i] = max(dp[i], 1+dp[j])
        }
        // go prev
        for j := i - 1; j >= max(i-d, 0); j-- {
            if arr[i] <= arr[j] {
                break
            }
            dp[i] = max(dp[i], 1+dp[j])
        }
    }
    return slices.Max(dp)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/01/23/meta-chart.jpeg" />
    // Input: arr = [6,4,14,6,8,13,9,7,10,6,12], d = 2
    // Output: 4
    // Explanation: You can start at index 10. You can jump 10 --> 8 --> 6 --> 7 as shown.
    // Note that if you start at index 6 you can only jump to index 7. You cannot jump to index 5 because 13 > 9. You cannot jump to index 4 because index 5 is between index 4 and 6 and 13 > 9.
    // Similarly You cannot jump from index 3 to index 2 or index 1.
    fmt.Println(maxJumps([]int{6,4,14,6,8,13,9,7,10,6,12}, 2)) // 4
    // Example 2:
    // Input: arr = [3,3,3,3,3], d = 3
    // Output: 1
    // Explanation: You can start at any index. You always cannot jump to any index.
    fmt.Println(maxJumps([]int{3,3,3,3,3}, 3)) // 1
    // Example 3:
    // Input: arr = [7,6,5,4,3,2,1], d = 1
    // Output: 7
    // Explanation: Start at index 0. You can visit all the indicies. 
    fmt.Println(maxJumps([]int{7,6,5,4,3,2,1}, 1)) // 7

    fmt.Println(maxJumps1([]int{6,4,14,6,8,13,9,7,10,6,12}, 2)) // 4
    fmt.Println(maxJumps1([]int{3,3,3,3,3}, 3)) // 1
    fmt.Println(maxJumps1([]int{7,6,5,4,3,2,1}, 1)) // 7

    fmt.Println(maxJumps2([]int{6,4,14,6,8,13,9,7,10,6,12}, 2)) // 4
    fmt.Println(maxJumps2([]int{3,3,3,3,3}, 3)) // 1
    fmt.Println(maxJumps2([]int{7,6,5,4,3,2,1}, 1)) // 7
}