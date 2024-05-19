package main

// 3068. Find the Maximum Sum of Node Values
// There exists an undirected tree with n nodes numbered 0 to n - 1. 
// You are given a 0-indexed 2D integer array edges of length n - 1, where edges[i] = [ui, vi] indicates that there is an edge between nodes ui and vi in the tree. 
// You are also given a positive integer k, and a 0-indexed array of non-negative integers nums of length n, where nums[i] represents the value of the node numbered i.

// Alice wants the sum of values of tree nodes to be maximum, 
// for which Alice can perform the following operation any number of times (including zero) on the tree:
//     Choose any edge [u, v] connecting the nodes u and v, and update their values as follows:
//         nums[u] = nums[u] XOR k
//         nums[v] = nums[v] XOR k

// Return the maximum possible sum of the values Alice can achieve by performing the operation any number of times.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012513.png" />
// Input: nums = [1,2,1], k = 3, edges = [[0,1],[0,2]]
// Output: 6
// Explanation: Alice can achieve the maximum sum of 6 using a single operation:
// - Choose the edge [0,2]. nums[0] and nums[2] become: 1 XOR 3 = 2, and the array nums becomes: [1,2,1] -> [2,2,2].
// The total sum of values is 2 + 2 + 2 = 6.
// It can be shown that 6 is the maximum achievable sum of values.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/01/09/screenshot-2024-01-09-220017.png" />
// Input: nums = [2,3], k = 7, edges = [[0,1]]
// Output: 9
// Explanation: Alice can achieve the maximum sum of 9 using a single operation:
// - Choose the edge [0,1]. nums[0] becomes: 2 XOR 7 = 5 and nums[1] become: 3 XOR 7 = 4, and the array nums becomes: [2,3] -> [5,4].
// The total sum of values is 5 + 4 = 9.
// It can be shown that 9 is the maximum achievable sum of values.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012641.png" />
// Input: nums = [7,7,7,7,7,7], k = 3, edges = [[0,1],[0,2],[0,3],[0,4],[0,5]]
// Output: 42
// Explanation: The maximum achievable sum is 42 which can be achieved by Alice performing no operations.
 
// Constraints:
//     2 <= n == nums.length <= 2 * 10^4
//     1 <= k <= 10^9
//     0 <= nums[i] <= 10^9
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= edges[i][0], edges[i][1] <= n - 1
//     The input is generated such that edges represent a valid tree.

import "fmt"

// 树形 DP
func maximumValueSum(nums []int, k int, edges [][]int) int64 {
    n, inf := len(nums), -1 << 32 - 1
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(int, int) (int, int)
    dfs = func(x, fa int) (int, int) {
        f0, f1 := 0, inf // f[x][0] 和 f[x][1]
        for _, y := range g[x] {
            if y != fa {
                r0, r1 := dfs(y, x)
                f0, f1 = max(f0+r0, f1+r1), max(f1+r0, f0+r1)
            }
        }
        return max(f0+nums[x], f1+(nums[x]^k)), max(f1+nums[x], f0+(nums[x]^k))
    }
    res, _ := dfs(0, -1)
    return int64(res)
}

// 状态机 DP
func maximumValueSum1(nums []int, k int, _ [][]int) int64 {
    f0, f1 := 0, -1 << 32 - 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, x := range nums {
        f0, f1 = max(f0+x, f1+(x^k)), max(f1+x, f0+(x^k))
    }
    return int64(f0)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012513.png" />
    // Input: nums = [1,2,1], k = 3, edges = [[0,1],[0,2]]
    // Output: 6
    // Explanation: Alice can achieve the maximum sum of 6 using a single operation:
    // - Choose the edge [0,2]. nums[0] and nums[2] become: 1 XOR 3 = 2, and the array nums becomes: [1,2,1] -> [2,2,2].
    // The total sum of values is 2 + 2 + 2 = 6.
    // It can be shown that 6 is the maximum achievable sum of values.
    fmt.Println(maximumValueSum([]int{1,2,1},3,[][]int{{0,1},{0,2}})) // 6
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/01/09/screenshot-2024-01-09-220017.png" />
    // Input: nums = [2,3], k = 7, edges = [[0,1]]
    // Output: 9
    // Explanation: Alice can achieve the maximum sum of 9 using a single operation:
    // - Choose the edge [0,1]. nums[0] becomes: 2 XOR 7 = 5 and nums[1] become: 3 XOR 7 = 4, and the array nums becomes: [2,3] -> [5,4].
    // The total sum of values is 5 + 4 = 9.
    // It can be shown that 9 is the maximum achievable sum of values.
    fmt.Println(maximumValueSum([]int{2, 3}, 7, [][]int{{0,1}})) // 9
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012641.png" />
    // Input: nums = [7,7,7,7,7,7], k = 3, edges = [[0,1],[0,2],[0,3],[0,4],[0,5]]
    // Output: 42
    // Explanation: The maximum achievable sum is 42 which can be achieved by Alice performing no operations.
    fmt.Println(maximumValueSum([]int{7,7,7,7,7,7}, 3, [][]int{{0,1},{0,2},{0,3},{0,4},{0,5}})) // 42

    fmt.Println(maximumValueSum1([]int{1,2,1},3,[][]int{{0,1},{0,2}})) // 6
    fmt.Println(maximumValueSum1([]int{2, 3}, 7, [][]int{{0,1}})) // 9
    fmt.Println(maximumValueSum1([]int{7,7,7,7,7,7}, 3, [][]int{{0,1},{0,2},{0,3},{0,4},{0,5}})) // 42
}