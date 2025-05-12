package main

// 3544. Subtree Inversion Sum
// You are given an undirected tree rooted at node 0, with n nodes numbered from 0 to n - 1. 
// The tree is represented by a 2D integer array edges of length n - 1, where edges[i] = [ui, vi] indicates an edge between nodes ui and vi.

// You are also given an integer array nums of length n, where nums[i] represents the value at node i, and an integer k.

// You may perform inversion operations on a subset of nodes subject to the following rules:
//     1. Subtree Inversion Operation:
//         1.1 When you invert a node, every value in the subtree rooted at that node is multiplied by -1.
//     2. Distance Constraint on Inversions:
//         2.1 You may only invert a node if it is "sufficiently far" from any other inverted node.
//         2.2 Specifically, if you invert two nodes a and b such that one is an ancestor of the other (i.e., if LCA(a, b) = a or LCA(a, b) = b), 
//             then the distance (the number of edges on the unique path between them) must be at least k.

// Return the maximum possible sum of the tree's node values after applying inversion operations.

// Example 1:
// Input: edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]], nums = [4,-8,-6,3,7,-2,5], k = 2
// Output: 27
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/03/29/tree1-3.jpg" />
// Apply inversion operations at nodes 0, 3, 4 and 6.
// The final nums array is [-4, 8, 6, 3, 7, 2, 5], and the total sum is 27.

// Example 2:
// Input: edges = [[0,1],[1,2],[2,3],[3,4]], nums = [-1,3,-2,4,-5], k = 2
// Output: 9
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/03/29/tree2-1.jpg" />
// Apply the inversion operation at node 4.
// The final nums array becomes [-1, 3, -2, 4, 5], and the total sum is 9.

// Example 3:
// Input: edges = [[0,1],[0,2]], nums = [0,-1,-2], k = 3
// Output: 3
// Explanation:
// Apply inversion operations at nodes 1 and 2.

// Constraints:
//     2 <= n <= 5 * 10^4
//     edges.length == n - 1
//     edges[i] = [ui, vi]
//     0 <= ui, vi < n
//     nums.length == n
//     -5 * 10^4 <= nums[i] <= 5 * 10^4
//     1 <= k <= 50
//     The input is generated such that edges represents a valid tree.

import "fmt"

// 记忆化搜索
func subtreeInversionSum(edges [][]int, nums []int, k int) int64 {
    n,inf := len(nums), 1 << 31
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    memo := make([][][2]int, n)
    for i := range memo {
        memo[i] = make([][2]int, k)
        for j := range memo[i] {
            for p := range memo[i][j] {
                memo[i][j][p] = -inf
            }
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(x, fa, cd, parity int) int
    dfs = func(x, fa, cd, parity int) int {
        p := &memo[x][cd][parity]
        if *p != -inf {
            return *p
        }
        // 不反转
        res := nums[x] * (1 - parity * 2)
        for _, y := range g[x] {
            if y != fa {
                res += dfs(y, x, max(cd - 1, 0), parity)
            }
        }
        // 反转
        if cd == 0 {
            s := nums[x] * (parity*2 - 1)
            for _, y := range g[x] {
                if y != fa {
                    s += dfs(y, x, k-1, parity^1) // 重置 CD
                }
            }
            res = max(res, s)
        }
        *p = res
        return res
    }
    return int64(dfs(0, -1, 0, 0))
}

func subtreeInversionSum1(edges [][]int, nums []int, k int) int64 {
    n := len(nums)
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    f := [][2]int{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(x, fa int) (int, int, int)
    dfs = func(x, fa int) (int, int, int) {
        f = append(f, [2]int{}) // 用于刷表
        s := nums[x] // 子树和
        notInv0, notInv1 := 0, 0 // 不反转 x 时的额外增量（0 表示上面反转了偶数次，1 表示上面反转了奇数次）
        for _, y := range g[x] {
            if y == fa { continue }
            sy, y0, y1 := dfs(y, x)
            s += sy
            // 不反转 x，反转次数的奇偶性不变
            notInv0 += y0
            notInv1 += y1
        }
        subRes := f[len(f) - 1] // 被刷表后的结果
        f = f[:len(f) - 1]
        // 反转 x
        // x 上面反转了偶数次，反转 x 会带来 -2 倍子树和的增量，且对于 x 的 k 级后代来说，上面反转了奇数次（所以是 subRes1）
        inv0 := subRes[1] - s*2
        // x 上面反转了奇数次，反转 x 会带来 2 倍子树和的增量，且对于 x 的 k 级后代来说，上面反转了偶数次（所以是 subRes0）
        inv1 := subRes[0] + s*2
        res0, res1 := max(notInv0, inv0), max(notInv1, inv1)
        // 刷表法：更新 x 的 k 级祖先的状态
        if len(f) >= k {
            f[len(f)-k][0] += res0
            f[len(f)-k][1] += res1
        }
        return s, res0, res1
    }
    s, res0, _ := dfs(0, -1)
    return int64(s + res0) // 对于根节点来说，上面一定反转了偶数次（0 次）
}

func main() {
    // Example 1:
    // Input: edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]], nums = [4,-8,-6,3,7,-2,5], k = 2
    // Output: 27
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/03/29/tree1-3.jpg" />
    // Apply inversion operations at nodes 0, 3, 4 and 6.
    // The final nums array is [-4, 8, 6, 3, 7, 2, 5], and the total sum is 27.
    fmt.Println(subtreeInversionSum([][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{2,6}}, []int{4,-8,-6,3,7,-2,5}, 2)) // 27
    // Example 2:
    // Input: edges = [[0,1],[1,2],[2,3],[3,4]], nums = [-1,3,-2,4,-5], k = 2
    // Output: 9
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/03/29/tree2-1.jpg" />
    // Apply the inversion operation at node 4.
    // The final nums array becomes [-1, 3, -2, 4, 5], and the total sum is 9.
    fmt.Println(subtreeInversionSum([][]int{{0,1},{1,2},{2,3},{3,4}}, []int{-1,3,-2,4,-5}, 2)) // 9
    // Example 3:
    // Input: edges = [[0,1],[0,2]], nums = [0,-1,-2], k = 3
    // Output: 3
    // Explanation:
    // Apply inversion operations at nodes 1 and 2.
    fmt.Println(subtreeInversionSum([][]int{{0,1},{0,2}}, []int{0,-1,-2}, 3)) // 3

    fmt.Println(subtreeInversionSum1([][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{2,6}}, []int{4,-8,-6,3,7,-2,5}, 2)) // 27
    fmt.Println(subtreeInversionSum1([][]int{{0,1},{1,2},{2,3},{3,4}}, []int{-1,3,-2,4,-5}, 2)) // 9
    fmt.Println(subtreeInversionSum1([][]int{{0,1},{0,2}}, []int{0,-1,-2}, 3)) // 3
}