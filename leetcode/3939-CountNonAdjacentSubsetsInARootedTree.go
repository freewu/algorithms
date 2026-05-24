package main

// 3939. Count Non Adjacent Subsets in a Rooted Tree
// You are given a rooted tree with n nodes labeled from 0 to n - 1, represented by an integer array parent of length n, where:
//     1. parent[0] = -1 (node 0 is the root).
//     2. For each 1 <= i < n, parent[i] is the parent of node i (0 <= parent[i] < i).

// You are also given an integer array nums of length n, where nums[i] is the value of node i, and an integer k.

// A non-empty subset of nodes is called valid if:
//     1. The sum of the values of the selected nodes is divisible by k.
//     2. No two selected nodes are adjacent in the tree (no node and its direct parent are both included in the subset).

// Return the number of valid subsets modulo 10^9 + 7.

// Example 1:
// Input: parent = [-1,0,1], nums = [1,2,3], k = 3
// Output: 1
// Explanation:
// ​​​​​​<img src="https://assets.leetcode.com/uploads/2025/07/17/image1.png" />
// The only valid subset is {2}. It contains node 2 with value 3, which is divisible by 3.

// Example 2:
// Input: parent = [-1,0,0,0], nums = [2,1,2,1], k = 3
// Output: 2
// Explanation:
// ​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2023/08/24/3939-1.png" />
// The valid subsets are:
// {1, 2}: Nodes 1 and 2 are both children of node 0 and not directly connected to each other. Their values sum to 1 + 2 = 3, which is divisible by 3.
// {2, 3}: Nodes 2 and 3 are also non-adjacent. Their values sum to 2 + 1 = 3, which is divisible by 3.
// No other subset satisfies both conditions. Therefore, the answer is 2.

// Constraints:
//     n == parent.length == nums.length
//     1 <= n <= 1000
//     parent[0] == -1
//     For all 1 <= i < n:
//     0 <= parent[i] < i
//     1 <= nums[i] <= 10^9
//     1 <= k <= 100​​​​​​​​​​​​​​​​​​​​​
//     parent describes a valid rooted tree.

import "fmt"

func countValidSubsets(parent []int, nums []int, k int) int {
    n,mod := len(parent), 1_000_000_007
    graph := make([][]int, n)
    for i := 1; i < n; i++ {
        p := parent[i]
        graph[p] = append(graph[p], i)
    }
    var dfs func(int) ([]int, []int)
    dfs = func(x int) ([]int, []int) {
        f0 := make([]int, k) // f0[i] 表示不选 x 时，子树 x 的子集点权和模 k 为 i 的方案数
        f1 := make([]int, k) // f1[i] 表示选 x 时，子树 x 的子集点权和模 k 为 i 的方案数
        f0[0] = 1
        f1[nums[x]%k] = 1
        for _, y := range graph[x] {
            fy0, fy1 := dfs(y)
            // 不选 x，那么 y 可选可不选
            nf0 := make([]int, k)
            for i := range k { // 枚举从子树 y 中选出的点权和模 k 为 i
                v := fy0[i] + fy1[i]
                if v == 0 { continue } // 优化
                for j, w := range f0 { // 枚举从之前的子树中选出的点权和模 k 为 j
                    s := (i + j) % k
                    nf0[s] = (nf0[s] + v*w) % mod
                }
            }
            // 选 x，那么 y 不能选
            nf1 := make([]int, k)
            for i, v := range fy0 { // 枚举从子树 y 中选出的点权和模 k 为 i
                if v == 0 { continue } // 优化
                for j, w := range f1 { // 枚举从 x 以及之前的子树中选出的点权和模 k 为 j
                    s := (i + j) % k
                    nf1[s] = (nf1[s] + v*w) % mod
                }
            }
            f0, f1 = nf0, nf1
        }
        return f0, f1
    }
    f0, f1 := dfs(0)
    return (f0[0] + f1[0] - 1 + mod) % mod // 恰好被 k 整除即模 k 为 0，注意减去空集的方案数 1
}

func countValidSubsets1(parent []int, nums []int, k int) int {
    n, mod := len(parent), 1_000_000_007
    dp0, dp1 := make([]int, n * k), make([]int, n * k)
    for i := range n {
        dp0[i * k] = 1
        dp1[i * k + ((nums[i] % k) + k) % k] = 1
    }
    ndp0, ndp1 := make([]int, k), make([]int, k)
    for i := n - 1; i > 0; i-- {
        p := parent[i]
        p0, p1 := dp0[p*k : p*k+k], dp1[p*k : p*k+k]
        i0, i1 := dp0[i*k : i*k+k], dp1[i*k : i*k+k]
        for j := 0; j < k; j++ {
            ndp0[j], ndp1[j] = 0, 0
        }
        for r1 := range k {
            v0, v1 := p0[r1], p1[r1]
            if v0 == 0 && v1 == 0 { continue }
            for r2 := range k {
                s := i0[r2] + i1[r2]
                if s >= mod {
                    s -= mod
                }
                nr := r1 + r2
                if nr >= k {
                    nr -= k
                }
                if v0 > 0 && s > 0 {
                    ndp0[nr] = int((int64(ndp0[nr]) + int64(v0) * int64(s)) % int64(mod))
                }
                if v1 > 0 && i0[r2] > 0 {
                    ndp1[nr] = int((int64(ndp1[nr]) + int64(v1) * int64(i0[r2])) % int64(mod))
                }
            }
        }
        for j := 0; j < k; j++ {
            p0[j], p1[j] = ndp0[j], ndp1[j]
        }
    }
    return (dp0[0] + dp1[0] - 1 + mod) % mod
}

func main() {
    // Example 1:
    // Input: parent = [-1,0,1], nums = [1,2,3], k = 3
    // Output: 1
    // Explanation:
    // ​​​​​​<img src="https://assets.leetcode.com/uploads/2025/07/17/image1.png" />
    // The only valid subset is {2}. It contains node 2 with value 3, which is divisible by 3.
    fmt.Println(countValidSubsets([]int{-1,0,1}, []int{1,2,3}, 3)) // 1
    // Example 2:
    // Input: parent = [-1,0,0,0], nums = [2,1,2,1], k = 3
    // Output: 2
    // Explanation:
    // ​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2023/08/24/3939-1.png" />
    // The valid subsets are:
    // {1, 2}: Nodes 1 and 2 are both children of node 0 and not directly connected to each other. Their values sum to 1 + 2 = 3, which is divisible by 3.
    // {2, 3}: Nodes 2 and 3 are also non-adjacent. Their values sum to 2 + 1 = 3, which is divisible by 3.
    // No other subset satisfies both conditions. Therefore, the answer is 2.
    fmt.Println(countValidSubsets([]int{-1,0,0,0}, []int{2,1,2,1}, 3)) // 2

    fmt.Println(countValidSubsets1([]int{-1,0,1}, []int{1,2,3}, 3)) // 1
    fmt.Println(countValidSubsets1([]int{-1,0,0,0}, []int{2,1,2,1}, 3)) // 2
}