package main

// 2322. Minimum Score After Removals on a Tree
// There is an undirected connected tree with n nodes labeled from 0 to n - 1 and n - 1 edges.

// You are given a 0-indexed integer array nums of length n where nums[i] represents the value of the ith node. 
// You are also given a 2D integer array edges of length n - 1 where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

// Remove two distinct edges of the tree to form three connected components. 
// For a pair of removed edges, the following steps are defined:
//     1. Get the XOR of all the values of the nodes for each of the three components respectively.
//     2. The difference between the largest XOR value and the smallest XOR value is the score of the pair.

//     For example, say the three components have the node values: [4,5,7], [1,9], and [3,3,3]. 
//     The three XOR values are 4 ^ 5 ^ 7 = 6, 1 ^ 9 = 8, and 3 ^ 3 ^ 3 = 3. 
//     The largest XOR value is 8 and the smallest XOR value is 3. The score is then 8 - 3 = 5.

// Return the minimum score of any possible pair of edge removals on the given tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/05/03/ex1drawio.png" />
// Input: nums = [1,5,5,4,11], edges = [[0,1],[1,2],[1,3],[3,4]]
// Output: 9
// Explanation: The diagram above shows a way to make a pair of removals.
// - The 1st component has nodes [1,3,4] with values [5,4,11]. Its XOR value is 5 ^ 4 ^ 11 = 10.
// - The 2nd component has node [0] with value [1]. Its XOR value is 1 = 1.
// - The 3rd component has node [2] with value [5]. Its XOR value is 5 = 5.
// The score is the difference between the largest and smallest XOR value which is 10 - 1 = 9.
// It can be shown that no other pair of removals will obtain a smaller score than 9.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/05/03/ex2drawio.png" />
// Input: nums = [5,5,2,4,4,2], edges = [[0,1],[1,2],[5,2],[4,3],[1,3]]
// Output: 0
// Explanation: The diagram above shows a way to make a pair of removals.
// - The 1st component has nodes [3,4] with values [4,4]. Its XOR value is 4 ^ 4 = 0.
// - The 2nd component has nodes [1,0] with values [5,5]. Its XOR value is 5 ^ 5 = 0.
// - The 3rd component has nodes [2,5] with values [2,2]. Its XOR value is 2 ^ 2 = 0.
// The score is the difference between the largest and smallest XOR value which is 0 - 0 = 0.
// We cannot obtain a smaller score than 0.

// Constraints:
//     n == nums.length
//     3 <= n <= 1000
//     1 <= nums[i] <= 10^8
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= ai, bi < n
//     ai != bi
//     edges represents a valid tree.

import "fmt"

func minimumScore(nums []int, edges [][]int) int {
    res, adj, dp := 1 << 31, make([][]int, len(nums)), make(map[[2]int][]int)
    for _, v := range edges {
        adj[v[0]] = append(adj[v[0]], v[1])
        adj[v[1]] = append(adj[v[1]], v[0])
    }
    var dfs func (i, j int) []int
    dfs = func (i, j int) []int {
        e := [2]int{i, j}
        if _, ok := dp[e]; !ok {
            x := nums[j]
            dp[e] = []int{}
            for _, k := range adj[j] {
                if k != i {
                    dp[e] = append(dp[e], dfs(j, k)...)
                    x ^= dp[e][len(dp[e])-1]
                }
            }
            dp[e] = append(dp[e], x)
        }
        return dp[e]
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, e := range edges {
        x, y := dfs(e[0], e[1]), dfs(e[1], e[0])
        for _, z := range [2][2][]int{[2][]int{x, y}, [2][]int{y, x}} {
            m, n := len(z[0])-1, len(z[1])-1
            for _, a := range z[0][:m] {
                b := z[0][m] ^ a
                res = min(res, max(max(a, b), z[1][n]) - min(min(a, b), z[1][n]))
            }
        }
    }
    return res
}

func minimumScore1(nums []int, edges [][]int) int {
    res, n := 1 << 31, len(nums)
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    xor, in, out, clock := make([]int, n), make([]int, n), make([]int, n), 0
    var dfs func(int, int)
    dfs = func(x, fa int) {
        clock++
        in[x] = clock
        xor[x] = nums[x]
        for _, y := range g[x] {
            if y != fa {
                dfs(y, x)
                xor[x] ^= xor[y]
            }
        }
        out[x] = clock
    }
    dfs(0, -1)
    isAncestor := func(x, y int) bool { return in[x] < in[y] && in[y] <= out[x] }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        for j := i+1; j < n; j++ {
            var x, y, z int
            if isAncestor(i, j) { // i 是 j 的祖先节点
                x, y, z = xor[j], xor[i]^xor[j], xor[0]^xor[i]
            } else if isAncestor(j, i) { // j 是 i 的祖先节点
                x, y, z = xor[i], xor[i]^xor[j], xor[0]^xor[j]
            } else { // 删除的两条边分别属于两颗不相交的子树
                x, y, z = xor[i], xor[j], xor[0]^xor[i]^xor[j]
            }
            res = min(res, max(max(x, y), z)-min(min(x, y), z))
            if res == 0 {
                return 0 // 提前退出
            }
        }
    }
    return res
}

func minimumScore2(nums []int, edges [][]int) int {
    res, count, n := 1 << 61, 0, len(nums)
    adj := make([][]int, n)
    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]], e[1])
        adj[e[1]] = append(adj[e[1]], e[0])
    }
    sum, in, out := make([]int, n), make([]int, n), make([]int, n)
    var dfs func(x, fa int)
    dfs = func(x, fa int) {
        in[x] = count
        count++
        sum[x] = nums[x]
        for _, y := range adj[x] {
            if y == fa {
                continue
            }
            dfs(y, x)
            sum[x] ^= sum[y]
        }
        out[x] = count
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    calc := func(part1, part2, part3 int) int { return max(part1, max(part2, part3)) - min(part1, min(part2, part3)) }
    dfs(0, -1)
    for u := 1; u < n; u++ {
        for v := u + 1; v < n; v++ {
            if in[v] > in[u] && in[v] < out[u] {
                res = min(res, calc(sum[0]^sum[u], sum[u]^sum[v], sum[v]))
            } else if in[u] > in[v] && in[u] < out[v] {
                res = min(res, calc(sum[0]^sum[v], sum[v]^sum[u], sum[u]))
            } else {
                res = min(res, calc(sum[0]^sum[u]^sum[v], sum[u], sum[v]))
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/05/03/ex1drawio.png" />
    // Input: nums = [1,5,5,4,11], edges = [[0,1],[1,2],[1,3],[3,4]]
    // Output: 9
    // Explanation: The diagram above shows a way to make a pair of removals.
    // - The 1st component has nodes [1,3,4] with values [5,4,11]. Its XOR value is 5 ^ 4 ^ 11 = 10.
    // - The 2nd component has node [0] with value [1]. Its XOR value is 1 = 1.
    // - The 3rd component has node [2] with value [5]. Its XOR value is 5 = 5.
    // The score is the difference between the largest and smallest XOR value which is 10 - 1 = 9.
    // It can be shown that no other pair of removals will obtain a smaller score than 9.
    fmt.Println(minimumScore([]int{1,5,5,4,11}, [][]int{{0,1},{1,2},{1,3},{3,4}})) // 9
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/05/03/ex2drawio.png" />
    // Input: nums = [5,5,2,4,4,2], edges = [[0,1],[1,2],[5,2],[4,3],[1,3]]
    // Output: 0
    // Explanation: The diagram above shows a way to make a pair of removals.
    // - The 1st component has nodes [3,4] with values [4,4]. Its XOR value is 4 ^ 4 = 0.
    // - The 2nd component has nodes [1,0] with values [5,5]. Its XOR value is 5 ^ 5 = 0.
    // - The 3rd component has nodes [2,5] with values [2,2]. Its XOR value is 2 ^ 2 = 0.
    // The score is the difference between the largest and smallest XOR value which is 0 - 0 = 0.
    // We cannot obtain a smaller score than 0.
    fmt.Println(minimumScore([]int{5,5,2,4,4,2}, [][]int{{0,1},{1,2},{5,2},{4,3},{1,3}})) // 0

    fmt.Println(minimumScore1([]int{1,5,5,4,11}, [][]int{{0,1},{1,2},{1,3},{3,4}})) // 9
    fmt.Println(minimumScore1([]int{5,5,2,4,4,2}, [][]int{{0,1},{1,2},{5,2},{4,3},{1,3}})) // 0

    fmt.Println(minimumScore2([]int{1,5,5,4,11}, [][]int{{0,1},{1,2},{1,3},{3,4}})) // 9
    fmt.Println(minimumScore2([]int{5,5,2,4,4,2}, [][]int{{0,1},{1,2},{5,2},{4,3},{1,3}})) // 0
}