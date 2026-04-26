package main

// 3910. Count Connected Subgraphs with Even Node Sum
// You are given an undirected graph with n nodes labeled from 0 to n - 1. 
// Node i has a value of nums[i], which is either 0 or 1. 
// The edges of the graph are given by a 2D array edges where edges[i] = [ui, vi] represents an edge between node ui and node vi.

// For a non-empty subset s of nodes in the graph, we consider the induced subgraph of s generated as follows:
//     1. We keep only the nodes in s.
//     2. We keep only the edges whose two endpoints are both in s.

// Return an integer representing the number of non-empty subsets s of nodes in the graph such that:
//     1. The induced subgraph of s is connected.
//     2. The sum of node values in s is even.
 

// Example 1:
// Input: nums = [1,0,1], edges = [[0,1],[1,2]]
// Output: 2
// Explanation:
// s	    | connected?                                | sum of node values    | counted?
// [0]	    | Yes	                                    | 1	                    | No
// [1]	    | Yes	                                    | 0	                    | Yes
// [2]	    | Yes	                                    | 1	                    | No
// [0,1]	| Yes	                                    | 1	                    | No
// [0,2]	| No, node 0 and node 2 are disconnected.	| 2	                    | No
// [1,2]	| Yes	                                    | 1	                    | No
// [0,1,2]	| Yes	                                    | 2	                    | Yes   

// Example 2:
// Input: nums = [1], edges = []
// Output: 0
// Explanation:
// s	    | connected?                                | sum of node values    | counted?
// [0]	    | Yes	                                    | 1	                    | No

// Constraints:
//     1 <= n == nums.length <= 13
//     nums[i] is 0 or 1.
//     0 <= edges.length <= n * (n - 1) / 2
//     edges[i] = [ui, vi]
//     0 <= ui < vi < n
//     All edges are distinct.

import "fmt"
import "math/bits"

func evenSumSubgraphs(nums []int, edges [][]int) int {
    res, n := 0, len(nums)
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    // 枚举节点集合 U = {0,1,2,...,n-1} 的非空子集 sub
    u := 1 << n - 1
    for sub := 1; sub <= u; sub++ {
        // 计算子图的点权异或和
        xor := 0
        for i, x := range nums {
            if sub >> i & 1 > 0 { // i 在 sub 中
                xor ^= x
            }
        }
        if xor != 0 {
            continue
        }
        // 判断子图是否连通
        visited := u ^ sub // 技巧：把不在子图中的节点都标记为已访问
        var dfs func(int)
        dfs = func(x int) {
            visited |= 1 << x // 标记 x 已访问
            for _, y := range g[x] {
                if visited >> y & 1 == 0 { // y 没有访问过
                    dfs(y)
                }
            }
        }
        dfs(bits.TrailingZeros(uint(sub))) // 随便选一个在子图中的节点，开始 DFS
        if visited == u { // 所有节点都已访问，子图是连通的
            res++
        }
    }
    return res  
}

func evenSumSubgraphs1(nums []int, edges [][]int) int {
    res, n := 0, len(nums)
    adj := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    on := make([]bool, n)
    var dfs func(int) int
    dfs = func(u int) int {
        on[u] = false
        res := 1
        for _, v := range adj[u] {
            if on[v] {
                res += dfs(v)
            }
        }
        return res
    }
    bit := func(n int, x int) bool { return ((n >> x) & 1) == 1 }
    for mask := 1; mask < (1 << n); mask++ {
        sum, root, count := 0, 0, 0
        for i := 0; i < n; i++ {
            if !bit(mask, i) {
                on[i] = false
            } else {
                on[i] = true
                sum += nums[i]
                root = i
                count++
            }
        }
        if (sum & 1) == 1 {
            continue
        }
        if count == dfs(root) {
            res++
        }
    }
    return res
}

func evenSumSubgraphs2(nums []int, edges [][]int) int {
    res, n := 0, len(nums)
    total := 1 << n
    index := make([]int, total)
    for i := 0; i < n; i++ {
        index[1<<i] = i
    }
    adj := make([]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] |= 1 << v
        adj[v] |= 1 << u
    }
    parity := make([]int, total)
    for mask := 1; mask < total; mask++ {
        low := mask & -mask
        parity[mask] = parity[mask^low] ^ nums[index[low]]
    }
    for mask := 1; mask < total; mask++ {
        if parity[mask] != 0 { continue }
        start := mask & -mask
        stack := start
        remain := mask ^ start
        for stack != 0 {
            b := stack & -stack
            stack ^= b
            u := index[b]
            next := adj[u] & remain
            for next != 0 {
                nb := next & -next
                next ^= nb
                remain ^= nb
                stack |= nb
            }
        }
        if remain == 0 {
            res++
        }
    }
    return res  
}

func main() {
    // Example 1:
    // Input: nums = [1,0,1], edges = [[0,1],[1,2]]
    // Output: 2
    // Explanation:
    // s	    | connected?                                | sum of node values    | counted?
    // [0]	    | Yes	                                    | 1	                    | No
    // [1]	    | Yes	                                    | 0	                    | Yes
    // [2]	    | Yes	                                    | 1	                    | No
    // [0,1]	| Yes	                                    | 1	                    | No
    // [0,2]	| No, node 0 and node 2 are disconnected.	| 2	                    | No
    // [1,2]	| Yes	                                    | 1	                    | No
    // [0,1,2]	| Yes	                                    | 2	                    | Yes   
    fmt.Println(evenSumSubgraphs([]int{1,0,1}, [][]int{{0,1},{1,2}})) // 2
    // Example 2:
    // Input: nums = [1], edges = []
    // Output: 0
    // Explanation:
    // s	    | connected?                                | sum of node values    | counted?
    // [0]	    | Yes	                                    | 1	                    | No
    fmt.Println(evenSumSubgraphs([]int{1}, [][]int{})) // 0

    fmt.Println(evenSumSubgraphs([]int{1,1,1,1,1,1,1,1,1,1}, [][]int{{0,1},{1,2}})) // 2
    fmt.Println(evenSumSubgraphs([]int{0,0,0,0,0,0,0,0,0,0}, [][]int{{0,1},{1,2}})) // 17
    fmt.Println(evenSumSubgraphs([]int{1,1,1,1,1,0,0,0,0,0}, [][]int{{0,1},{1,2}})) // 7
    fmt.Println(evenSumSubgraphs([]int{0,0,0,0,0,1,1,1,1,1}, [][]int{{0,1},{1,2}})) // 8
    fmt.Println(evenSumSubgraphs([]int{0,1,0,1,0,1,0,1,0,1}, [][]int{{0,1},{1,2}})) // 5
    fmt.Println(evenSumSubgraphs([]int{1,0,1,0,1,0,1,0,1,0}, [][]int{{0,1},{1,2}})) // 6

    fmt.Println(evenSumSubgraphs1([]int{1,0,1}, [][]int{{0,1},{1,2}})) // 2
    fmt.Println(evenSumSubgraphs1([]int{1}, [][]int{})) // 0
    fmt.Println(evenSumSubgraphs1([]int{1,1,1,1,1,1,1,1,1,1}, [][]int{{0,1},{1,2}})) // 2
    fmt.Println(evenSumSubgraphs1([]int{0,0,0,0,0,0,0,0,0,0}, [][]int{{0,1},{1,2}})) // 17
    fmt.Println(evenSumSubgraphs1([]int{1,1,1,1,1,0,0,0,0,0}, [][]int{{0,1},{1,2}})) // 7
    fmt.Println(evenSumSubgraphs1([]int{0,0,0,0,0,1,1,1,1,1}, [][]int{{0,1},{1,2}})) // 8
    fmt.Println(evenSumSubgraphs1([]int{0,1,0,1,0,1,0,1,0,1}, [][]int{{0,1},{1,2}})) // 5
    fmt.Println(evenSumSubgraphs1([]int{1,0,1,0,1,0,1,0,1,0}, [][]int{{0,1},{1,2}})) // 6

    fmt.Println(evenSumSubgraphs2([]int{1,0,1}, [][]int{{0,1},{1,2}})) // 2
    fmt.Println(evenSumSubgraphs2([]int{1}, [][]int{})) // 0
    fmt.Println(evenSumSubgraphs2([]int{1,1,1,1,1,1,1,1,1,1}, [][]int{{0,1},{1,2}})) // 2
    fmt.Println(evenSumSubgraphs2([]int{0,0,0,0,0,0,0,0,0,0}, [][]int{{0,1},{1,2}})) // 17
    fmt.Println(evenSumSubgraphs2([]int{1,1,1,1,1,0,0,0,0,0}, [][]int{{0,1},{1,2}})) // 7
    fmt.Println(evenSumSubgraphs2([]int{0,0,0,0,0,1,1,1,1,1}, [][]int{{0,1},{1,2}})) // 8
    fmt.Println(evenSumSubgraphs2([]int{0,1,0,1,0,1,0,1,0,1}, [][]int{{0,1},{1,2}})) // 5
    fmt.Println(evenSumSubgraphs2([]int{1,0,1,0,1,0,1,0,1,0}, [][]int{{0,1},{1,2}})) // 6
}