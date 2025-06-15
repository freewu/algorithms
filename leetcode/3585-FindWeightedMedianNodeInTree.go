package main

// 3585. Find Weighted Median Node in Tree
// You are given an integer n and an undirected, weighted tree rooted at node 0 with n nodes numbered from 0 to n - 1. 
// This is represented by a 2D array edges of length n - 1, where edges[i] = [ui, vi, wi] indicates an edge from node ui to vi with weight wi.

// The weighted median node is defined as the first node x on the path from ui to vi such that the sum of edge weights from ui to x is greater than or equal to half of the total path weight.

// You are given a 2D integer array queries. 
// For each queries[j] = [uj, vj], determine the weighted median node along the path from uj to vj.

// Return an array ans, where ans[j] is the node index of the weighted median for queries[j].

// Example 1:
// Input: n = 2, edges = [[0,1,7]], queries = [[1,0],[0,1]]
// Output: [0,1]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/05/26/screenshot-2025-05-26-at-193447.png" />
// Query	|   Path	|   Edge Weights    |	Total Path Weight	| Half	| Explanation	                                | Answer
// [1, 0]	    1 → 0	    [7]	                    7	               3.5	  Sum from 1 → 0 = 7 >= 3.5, median is node 0.	    0
// [0, 1]	    0 → 1	    [7]	                    7	               3.5	  Sum from 0 → 1 = 7 >= 3.5, median is node 1.	    1

// Example 2:
// Input: n = 3, edges = [[0,1,2],[2,0,4]], queries = [[0,1],[2,0],[1,2]]
// Output: [1,0,2]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/05/26/screenshot-2025-05-26-at-193610.png" />
// Query	|   Path	|   Edge Weights    |	Total Path Weight	| Half	| Explanation	                                                                |   Answer
// [0, 1]	    0 → 1	        [2]	                2	                1	    Sum from 0 → 1 = 2 >= 1, median is node 1.	                                        1
// [2, 0]	    2 → 0	        [4]	                4	                2	    Sum from 2 → 0 = 4 >= 2, median is node 0.	                                        0
// [1, 2]	    1 → 0 → 2	    [2, 4]	            6	                3	    Sum from 1 → 0 = 2 < 3. Sum from 1 → 2 = 2 + 4 = 6 >= 3, median is node 2.	        2

// Example 3:
// Input: n = 5, edges = [[0,1,2],[0,2,5],[1,3,1],[2,4,3]], queries = [[3,4],[1,2]]
// Output: [2,2]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/05/26/screenshot-2025-05-26-at-193857.png" />
// Query	|   Path	            |   Edge Weights    |	Total Path Weight	|   Half	| Explanation	                                                | Answer
// [3, 4]	    3 → 1 → 0 → 2 → 4	    [1, 2, 5, 3]	        11	                5.5	        Sum from 3 → 1 = 1 < 5.5.                                       2
//                                                                                             Sum from 3 → 0 = 1 + 2 = 3 < 5.5.                     
//                                                                                             Sum from 3 → 2 = 1 + 2 + 5 = 8 >= 5.5, median is node 2.	
// [1, 2]	    1 → 0 → 2	            [2, 5]	                 7	                3.5	        Sum from 1 → 0 = 2 < 3.5.                                       2
//                                                                                             Sum from 1 → 2 = 2 + 5 = 7 >= 3.5, median is node 2.   
 
// Constraints:
//     2 <= n <= 10^5
//     edges.length == n - 1
//     edges[i] == [ui, vi, wi]
//     0 <= ui, vi < n
//     1 <= wi <= 10^9
//     1 <= queries.length <= 10^5
//     queries[j] == [uj, vj]
//     0 <= uj, vj < n
//     The input is generated such that edges represents a valid tree.

import "fmt"

func findMedian(n int, edges [][]int, queries [][]int) []int {
    g := make([][][]int, n)
    for _, p := range edges {
        g[p[0]] = append(g[p[0]], []int{p[1], p[2]})
        g[p[1]] = append(g[p[1]], []int{p[0], p[2]})
    }
    arr, h, w := make([]int, n), make([]int, n), make([]int64, n)
    arr[0] = -1
    var dfs func(g [][][]int, u, d int, e int64, v []bool, p, h []int, w []int64) int 
    dfs = func(g [][][]int, u, d int, e int64, v []bool, p, h []int, w []int64) int {
        v[u], h[u], w[u] = true, d, e
        res := d
        for _, c := range g[u] {
            if !v[c[0]] {
                p[c[0]] = u
                r := dfs(g, c[0], d + 1, e + int64(c[1]), v, p, h, w)
                if r > res {
                    res = r
                }
            }
        }
        return res
    }
    ancestor := func(p [][]int, u, k int) int {
        res := u
        for i, b := 0, 1; k >= b && res >= 0; i++ {
            if k & b == b {
                res = p[i][res]
            }
            b <<= 1
        }
        return res
    }
    lca := func(h []int, p [][]int, u, v int) int {
        hu, hv := h[u], h[v]
        if hv < hu {
            u, v, hu, hv = v, u, hv, hu
        }
        res, l, r := 0, 0, hu + 1
        for r - l > 1 {
            m := (l + r) / 2
            if q := ancestor(p, u, hu - m); q == ancestor(p, v, hv - m) {
                l, res = m, q
            } else {
                r = m
            }
        }
        return res
    }
    max_h := dfs(g, 0, 0, 0, make([]bool, n), arr, h, w)
    p := [][]int{arr}
    for b := 2; b <= max_h; b <<= 1 {
        arr2 := make([]int, n)
        for d, a := range arr {
            if a == -1 {
                arr2[d] = a
            } else {
                arr2[d] = arr[a]
            }
        }
        p = append(p, arr2)
        arr = arr2
    }
    res := make([]int, len(queries))
    for i, q := range queries {
        u, v := q[0], q[1]
        if u == v {
            res[i] = u
        } else {
            a := lca(h, p, u, v)
            total_path := w[u] + w[v] - 2 * w[a]
            total_len := h[u] + h[v] - 2 * h[a]
            l, r := -1, total_len
            for r - l > 1 {
                m := (l + r) / 2
                path := w[u] - w[a]
                switch d := h[u] - h[a] - m; {
                    case d < 0:
                    path = total_path - w[v] + w[ancestor(p, v, total_len - m)]
                    case d > 0:
                    path = w[u] - w[ancestor(p, u, m)]
                }
                if 2 * path >= total_path {
                    r = m
                } else {
                    l = m
                }
            }
            switch d := h[u] - h[a] - r; {
                case d < 0:
                res[i] = ancestor(p, v, total_len - r)
                case d == 0:
                res[i] = a
                default:
                res[i] = ancestor(p, u, r)
            }
        }
    }
    return res
}

const MAXN = 101010
var (
    adj   [MAXN][][2]int
    level [MAXN]int
    lca   [MAXN][22]int
    dep   [MAXN]int64
)

func findMedian1(n int, edges [][]int, queries [][]int) []int {
    helper := func(u int, d float64, fl bool) int {
        for {
            v := -1
            for i := 21; i >= 0; i-- {
                if lca[u][i] == 0 {
                    continue
                }
                if float64(dep[u]-dep[lca[u][i]]) < d {
                    v = lca[u][i]
                    break
                }
            }
            if v == -1 {
                break
            }
            d = d - float64(dep[u]) + float64(dep[v])
            u = v
        }
        if fl {
            return u - 1
        }
        return lca[u][0] - 1
    }
    lcaQuery := func(u, v int) int {
        if level[u] < level[v] {
            u, v = v, u
        }
        diff := level[u] - level[v]
        for i := 0; diff > 0; i++ {
            if diff&1 == 1 {
                u = lca[u][i]
            }
            diff >>= 1
        }
        if u != v {
            for i := 21; i >= 0; i-- {
                if lca[u][i] != lca[v][i] {
                    u = lca[u][i]
                    v = lca[v][i]
                }
            }
            u = lca[u][0]
        }
        return u
    }
    distance := func (u, v int) int64 {
        w := lcaQuery(u, v)
        return dep[u] + dep[v] - 2*dep[w]
    }
    var dfs func(u, lvl, par int) 
    dfs = func(u, lvl, par int) {
        level[u] = lvl
        lca[u][0] = par
        for i := 1; i < 22; i++ {
            lca[u][i] = lca[lca[u][i-1]][i-1]
        }
        for _, vw := range adj[u] {
            v, w := vw[0], vw[1]
            if v == par {
                continue
            }
            dep[v] = dep[u] + int64(w)
            dfs(v, lvl+1, u)
        }
    }
    for i := 1; i <= n; i++ {
        adj[i] = adj[i][:0]
        for j := 0; j < 22; j++ {
            lca[i][j] = 0
        }
        dep[i] = 0
        level[i] = 0
    }
    for _, e := range edges {
        u := e[0] + 1
        v := e[1] + 1
        w := e[2]
        adj[u] = append(adj[u], [2]int{v, w})
        adj[v] = append(adj[v], [2]int{u, w})
    }
    dfs(1, 0, 0)
    res := make([]int, len(queries))
    for i, q := range queries {
        u := q[0] + 1
        v := q[1] + 1
        w := lcaQuery(u, v)
        d := float64(distance(u, v))
        if w == u {
            res[i] = helper(v, d/2.0+0.1, true)
        } else if w == v {
            res[i] = helper(u, d/2.0, false)
        } else {
            du := float64(distance(u, w))
            if du*2.0 < d {
                rem := float64(dep[v]-dep[w]) - (d/2.0 - du)
                res[i] = helper(v, rem+0.1, true)
            } else {
                res[i] = helper(u, d/2.0, false)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2, edges = [[0,1,7]], queries = [[1,0],[0,1]]
    // Output: [0,1]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/05/26/screenshot-2025-05-26-at-193447.png" />
    // Query	|   Path	|   Edge Weights    |	Total Path Weight	| Half	| Explanation	                                | Answer
    // [1, 0]	    1 → 0	    [7]	                    7	               3.5	  Sum from 1 → 0 = 7 >= 3.5, median is node 0.	    0
    // [0, 1]	    0 → 1	    [7]	                    7	               3.5	  Sum from 0 → 1 = 7 >= 3.5, median is node 1.	    1
    fmt.Println(findMedian(2,[][]int{{0,1,7}}, [][]int{{1,0},{0,1}})) // [0,1]
    // Example 2:
    // Input: n = 3, edges = [[0,1,2],[2,0,4]], queries = [[0,1],[2,0],[1,2]]
    // Output: [1,0,2]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/05/26/screenshot-2025-05-26-at-193610.png" />
    // Query	|   Path	|   Edge Weights    |	Total Path Weight	| Half	| Explanation	                                                                |   Answer
    // [0, 1]	    0 → 1	        [2]	                2	                1	    Sum from 0 → 1 = 2 >= 1, median is node 1.	                                        1
    // [2, 0]	    2 → 0	        [4]	                4	                2	    Sum from 2 → 0 = 4 >= 2, median is node 0.	                                        0
    // [1, 2]	    1 → 0 → 2	    [2, 4]	            6	                3	    Sum from 1 → 0 = 2 < 3. Sum from 1 → 2 = 2 + 4 = 6 >= 3, median is node 2.	        2
    fmt.Println(findMedian(3, [][]int{{0,1,2},{2,0,4}}, [][]int{{0,1},{2,0},{1,2}})) // [1,0,2]
    // Example 3:
    // Input: n = 5, edges = [[0,1,2],[0,2,5],[1,3,1],[2,4,3]], queries = [[3,4],[1,2]]
    // Output: [2,2]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/05/26/screenshot-2025-05-26-at-193857.png" />
    // Query	|   Path	            |   Edge Weights    |	Total Path Weight	|   Half	| Explanation	                                                | Answer
    // [3, 4]	    3 → 1 → 0 → 2 → 4	    [1, 2, 5, 3]	        11	                5.5	        Sum from 3 → 1 = 1 < 5.5.                                       2
    //                                                                                             Sum from 3 → 0 = 1 + 2 = 3 < 5.5.                     
    //                                                                                             Sum from 3 → 2 = 1 + 2 + 5 = 8 >= 5.5, median is node 2.	
    // [1, 2]	    1 → 0 → 2	            [2, 5]	                 7	                3.5	        Sum from 1 → 0 = 2 < 3.5.                                       2
    //                                                                                             Sum from 1 → 2 = 2 + 5 = 7 >= 3.5, median is node 2.   
    fmt.Println(findMedian(5, [][]int{{0,1,2},{0,2,5},{1,3,1},{2,4,3}}, [][]int{{3,4},{1,2}})) // [2,2]

    fmt.Println(findMedian1(2,[][]int{{0,1,7}}, [][]int{{1,0},{0,1}})) // [0,1]
    fmt.Println(findMedian1(3, [][]int{{0,1,2},{2,0,4}}, [][]int{{0,1},{2,0},{1,2}})) // [1,0,2]
    fmt.Println(findMedian1(5, [][]int{{0,1,2},{0,2,5},{1,3,1},{2,4,3}}, [][]int{{3,4},{1,2}})) // [2,2]
}