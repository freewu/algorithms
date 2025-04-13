package main

// 3515. Shortest Path in a Weighted Tree
// You are given an integer n and an undirected, weighted tree rooted at node 1 with n nodes numbered from 1 to n. 
// This is represented by a 2D array edges of length n - 1, where edges[i] = [ui, vi, wi] indicates an undirected edge from node ui to vi with weight wi.

// You are also given a 2D integer array queries of length q, where each queries[i] is either:
//     1. [1, u, v, w'] – Update the weight of the edge between nodes u and v to w', 
//        where (u, v) is guaranteed to be an edge present in edges.
//     2. [2, x] – Compute the shortest path distance from the root node 1 to node x.

// Return an integer array answer, where answer[i] is the shortest path distance from node 1 to x for the ith query of [2, x].

// Example 1:
// Input: n = 2, edges = [[1,2,7]], queries = [[2,2],[1,1,2,4],[2,2]]
// Output: [7,4]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/03/13/screenshot-2025-03-13-at-133524.png" />
// Query [2,2]: The shortest path from root node 1 to node 2 is 7.
// Query [1,1,2,4]: The weight of edge (1,2) changes from 7 to 4.
// Query [2,2]: The shortest path from root node 1 to node 2 is 4.

// Example 2:
// Input: n = 3, edges = [[1,2,2],[1,3,4]], queries = [[2,1],[2,3],[1,1,3,7],[2,2],[2,3]]
// Output: [0,4,2,7]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/03/13/screenshot-2025-03-13-at-132247.png" />
// Query [2,1]: The shortest path from root node 1 to node 1 is 0.
// Query [2,3]: The shortest path from root node 1 to node 3 is 4.
// Query [1,1,3,7]: The weight of edge (1,3) changes from 4 to 7.
// Query [2,2]: The shortest path from root node 1 to node 2 is 2.
// Query [2,3]: The shortest path from root node 1 to node 3 is 7.

// Example 3:
// Input: n = 4, edges = [[1,2,2],[2,3,1],[3,4,5]], queries = [[2,4],[2,3],[1,2,3,3],[2,2],[2,3]]
// Output: [8,3,2,5]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/03/13/screenshot-2025-03-13-at-133306.png" />
// Query [2,4]: The shortest path from root node 1 to node 4 consists of edges (1,2), (2,3), and (3,4) with weights 2 + 1 + 5 = 8.
// Query [2,3]: The shortest path from root node 1 to node 3 consists of edges (1,2) and (2,3) with weights 2 + 1 = 3.
// Query [1,2,3,3]: The weight of edge (2,3) changes from 1 to 3.
// Query [2,2]: The shortest path from root node 1 to node 2 is 2.
// Query [2,3]: The shortest path from root node 1 to node 3 consists of edges (1,2) and (2,3) with updated weights 2 + 3 = 5.

// Constraints:
//     1 <= n <= 10^5
//     edges.length == n - 1
//     edges[i] == [ui, vi, wi]
//     1 <= ui, vi <= n
//     1 <= wi <= 10^4
//     The input is generated such that edges represents a valid tree.
//     1 <= queries.length == q <= 10^5
//     queries[i].length == 2 or 4
//     queries[i] == [1, u, v, w'] or,
//     queries[i] == [2, x]
//     1 <= u, v, x <= n
//     (u, v) is always an edge from edges.
//     1 <= w' <= 10^4

import "fmt"

type Fenwick []int

func newFenwickTree(n int) Fenwick {
    return make(Fenwick, n + 1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 1 <= i <= n
func (f Fenwick) update(i, val int) {
    for ; i < len(f); i += i & -i {
        f[i] += val
    }
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
func (f Fenwick) pre(i int) int {
    sum := 0
    for ; i > 0; i &= i - 1 {
        sum += f[i]
    }
    return sum
}

func treeQueries(n int, edges [][]int, queries [][]int) []int {
    g := make([][]int, n + 1)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    clock, in, out := 0, make([]int, n+1), make([]int, n+1)
    var dfs func(int, int)
    dfs = func(x, fa int) {
        clock++
        in[x] = clock // 进来的时间
        for _, y := range g[x] {
            if y != fa {
                dfs(y, x)
            }
        }
        out[x] = clock // 离开的时间
    }
    dfs(1, 0)

    // 对于一条边 x-y（y 是 x 的儿子），把边权保存在 weight[y] 中
    res, weight := []int{}, make([]int, n + 1)
    diff := newFenwickTree(n)
    update := func(x, y, w int) {
        // 保证 y 是 x 的儿子
        if in[x] > in[y] {
            y = x
        }
        d := w - weight[y] // 边权的增量
        weight[y] = w
        // 把子树 y 中的最短路长度都增加 d（用差分树状数组维护）
        diff.update(in[y], d)
        diff.update(out[y]+1, -d)
    }
    for _, e := range edges {
        update(e[0], e[1], e[2])
    }
    for _, q := range queries {
        if q[0] == 1 {
            update(q[1], q[2], q[3])
        } else {
            res = append(res, diff.pre(in[q[1]]))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2, edges = [[1,2,7]], queries = [[2,2],[1,1,2,4],[2,2]]
    // Output: [7,4]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/03/13/screenshot-2025-03-13-at-133524.png" />
    // Query [2,2]: The shortest path from root node 1 to node 2 is 7.
    // Query [1,1,2,4]: The weight of edge (1,2) changes from 7 to 4.
    // Query [2,2]: The shortest path from root node 1 to node 2 is 4.
    fmt.Println(treeQueries(2, [][]int{{1,2,7}}, [][]int{{2,2},{1,1,2,4},{2,2}})) // [7,4]
    // Example 2:
    // Input: n = 3, edges = [[1,2,2],[1,3,4]], queries = [[2,1],[2,3],[1,1,3,7],[2,2],[2,3]]
    // Output: [0,4,2,7]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/03/13/screenshot-2025-03-13-at-132247.png" />
    // Query [2,1]: The shortest path from root node 1 to node 1 is 0.
    // Query [2,3]: The shortest path from root node 1 to node 3 is 4.
    // Query [1,1,3,7]: The weight of edge (1,3) changes from 4 to 7.
    // Query [2,2]: The shortest path from root node 1 to node 2 is 2.
    // Query [2,3]: The shortest path from root node 1 to node 3 is 7.
    fmt.Println(treeQueries(3, [][]int{{1,2,2},{1,3,4}}, [][]int{{2,1},{2,3},{1,1,3,7},{2,2},{2,3}})) // [0,4,2,7]
    // Example 3:
    // Input: n = 4, edges = [[1,2,2],[2,3,1],[3,4,5]], queries = [[2,4],[2,3],[1,2,3,3],[2,2],[2,3]]
    // Output: [8,3,2,5]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/03/13/screenshot-2025-03-13-at-133306.png" />
    // Query [2,4]: The shortest path from root node 1 to node 4 consists of edges (1,2), (2,3), and (3,4) with weights 2 + 1 + 5 = 8.
    // Query [2,3]: The shortest path from root node 1 to node 3 consists of edges (1,2) and (2,3) with weights 2 + 1 = 3.
    // Query [1,2,3,3]: The weight of edge (2,3) changes from 1 to 3.
    // Query [2,2]: The shortest path from root node 1 to node 2 is 2.
    // Query [2,3]: The shortest path from root node 1 to node 3 consists of edges (1,2) and (2,3) with updated weights 2 + 3 = 5.
    fmt.Println(treeQueries(4, [][]int{{1,2,2},{2,3,1},{3,4,5}}, [][]int{{2,4},{2,3},{1,2,3,3},{2,2},{2,3}})) // [8,3,2,5]
}