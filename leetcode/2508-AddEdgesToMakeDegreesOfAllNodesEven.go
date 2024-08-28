package main

// 2508. Add Edges to Make Degrees of All Nodes Even
// There is an undirected graph consisting of n nodes numbered from 1 to n. 
// You are given the integer n and a 2D array edges where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi. 
// The graph can be disconnected.

// You can add at most two additional edges (possibly none) to this graph so that there are no repeated edges and no self-loops.

// Return true if it is possible to make the degree of each node in the graph even, otherwise return false.

// The degree of a node is the number of edges connected to it.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/10/26/agraphdrawio.png" />
// Input: n = 5, edges = [[1,2],[2,3],[3,4],[4,2],[1,4],[2,5]]
// Output: true
// Explanation: The above diagram shows a valid way of adding an edge.
// Every node in the resulting graph is connected to an even number of edges.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/10/26/aagraphdrawio.png" />
// Input: n = 4, edges = [[1,2],[3,4]]
// Output: true
// Explanation: The above diagram shows a valid way of adding two edges.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/10/26/aaagraphdrawio.png" />
// Input: n = 4, edges = [[1,2],[1,3],[1,4]]
// Output: false
// Explanation: It is not possible to obtain a valid graph with adding at most 2 edges.

// Constraints:
//     3 <= n <= 10^5
//     2 <= edges.length <= 10^5
//     edges[i].length == 2
//     1 <= ai, bi <= n
//     ai != bi
//     There are no repeated edges.

import "fmt"

func isPossible(n int, edges [][]int) bool {
    g := map[int]map[int]bool{}
    for _, e := range edges {
        x, y := e[0], e[1]
        if g[x] == nil {
            g[x] = map[int]bool{}
        }
        g[x][y] = true
        if g[y] == nil {
            g[y] = map[int]bool{}
        }
        g[y][x] = true
    }
    odd := []int{}
    for i, v := range g {
        if len(v) % 2 > 0 { // 把度数为奇数的节点记到 odd 
            odd = append(odd, i)
        }
    }
    m := len(odd)
    if m == 0 { // 已经符合要求
        return true
    }
    if m == 2 {
        x, y := odd[0], odd[1]
        if !g[x][y] { // 如果 x 和 y 之间没有边，那么连边之后就符合要求
            return true
        }
        // 如果 x 和 y 之间有边，那么枚举 [1,n] 的所有不为 x 和 y 的点 i，由于 i 的度数一定是偶数，如果 i 和 x 以及 i 和 y 之间没有边，那么连边之后就符合要求
        for i := 1; i <= n; i++ {
            if i != x && i != y && !g[i][x] && !g[i][y] {
                return true
            }
        }
        return false
    }
    if m == 4 {
        a, b, c, d := odd[0], odd[1], odd[2], odd[3]
        // 如果 a 和 b 以及 c 和 d 之间没有边，那么连边之后就符合要求了。
        // 如果 a 和 c 以及 b 和 d 之间没有边，那么连边之后就符合要求了。
        // 如果 a 和 d 以及 b 和 c 之间没有边，那么连边之后就符合要求了
        return !g[a][b] && !g[c][d] || !g[a][c] && !g[b][d] || !g[a][d] && !g[b][c]
    }
    return false
}

func isPossible1(n int, edges [][]int) bool {
    d := make([]int, n)
    for _, p := range edges {
        d[p[0] - 1]++
        d[p[1] - 1]++
    }
    odd := []int{}
    for u, c := range d {
        d[u] = 0
        switch {
            case c == n - 1 && n & 1 == 0:
                return false
            case c & 1 == 1:
                if len(odd) == 4 {
                    return false
                }
                odd = append(odd, u)
                d[u] = -len(odd)
        }
    }
    switch len(odd) {
        case 0:
        return true
        case 2:
        f := false
        for _, p := range edges {
            switch a, b := d[p[0] - 1], d[p[1] - 1]; {
                case a < 0 && b < 0:
                    f = true
                case a < 0:
                    d[p[1] - 1]++
                case b < 0:
                    d[p[0] - 1]++
            }
        }
        if !f {
            return true
        }
        for _, c := range d {
            if c == 0 {
                return true
            }
        }
        return false
        case 4:
        s := make([]bool, 16)
        for _, p := range edges {
            a, b := d[p[0] - 1], d[p[1] - 1]
            if a < 0 && b < 0 {
                s[(1 << (4 + a)) | (1 << (4 + b))] = true
            }
        }
        for _, i := range []int{3,5,6} {
            if !s[i] && !s[15 ^ i] {
                return true
            }
        }
        return false
    }
    return false
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/10/26/agraphdrawio.png" />
    // Input: n = 5, edges = [[1,2],[2,3],[3,4],[4,2],[1,4],[2,5]]
    // Output: true
    // Explanation: The above diagram shows a valid way of adding an edge.
    // Every node in the resulting graph is connected to an even number of edges.
    fmt.Println(isPossible(5,[][]int{{1,2},{2,3},{3,4},{4,2},{1,4},{2,5}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/10/26/aagraphdrawio.png" />
    // Input: n = 4, edges = [[1,2],[3,4]]
    // Output: true
    // Explanation: The above diagram shows a valid way of adding two edges.
    fmt.Println(isPossible(4,[][]int{{1,2},{3,4}})) // true
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/10/26/aaagraphdrawio.png" />
    // Input: n = 4, edges = [[1,2],[1,3],[1,4]]
    // Output: false
    // Explanation: It is not possible to obtain a valid graph with adding at most 2 edges.
    fmt.Println(isPossible(4,[][]int{{1,2},{1,3},{1,4}})) // false

    fmt.Println(isPossible1(5,[][]int{{1,2},{2,3},{3,4},{4,2},{1,4},{2,5}})) // true
    fmt.Println(isPossible1(4,[][]int{{1,2},{3,4}})) // true
    fmt.Println(isPossible1(4,[][]int{{1,2},{1,3},{1,4}})) // false
}