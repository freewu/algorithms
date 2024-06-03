package main

// 3067. Count Pairs of Connectable Servers in a Weighted Tree Network
// You are given an unrooted weighted tree with n vertices representing servers numbered from 0 to n - 1, 
// an array edges where edges[i] = [ai, bi, weighti] represents a bidirectional edge between vertices ai and bi of weight weighti. 
// You are also given an integer signalSpeed.

// Two servers a and b are connectable through a server c if:
//     a < b, a != c and b != c.
//     The distance from c to a is divisible by signalSpeed.
//     The distance from c to b is divisible by signalSpeed.
//     The path from c to b and the path from c to a do not share any edges.

// Return an integer array count of length n where count[i] is the number of server pairs that are connectable through the server i.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/01/21/example22.png" />
// Input: edges = [[0,1,1],[1,2,5],[2,3,13],[3,4,9],[4,5,2]], signalSpeed = 1
// Output: [0,4,6,6,4,0]
// Explanation: Since signalSpeed is 1, count[c] is equal to the number of pairs of paths that start at c and do not share any edges.
// In the case of the given path graph, count[c] is equal to the number of servers to the left of c multiplied by the servers to the right of c.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/01/21/example11.png" />
// Input: edges = [[0,6,3],[6,5,3],[0,3,1],[3,2,7],[3,1,6],[3,4,2]], signalSpeed = 3
// Output: [2,0,0,0,0,0,2]
// Explanation: Through server 0, there are 2 pairs of connectable servers: (4, 5) and (4, 6).
// Through server 6, there are 2 pairs of connectable servers: (4, 5) and (0, 5).
// It can be shown that no two servers are connectable through servers other than 0 and 6.
 
// Constraints:
//     2 <= n <= 1000
//     edges.length == n - 1
//     edges[i].length == 3
//     0 <= ai, bi < n
//     edges[i] = [ai, bi, weighti]
//     1 <= weighti <= 10^6
//     1 <= signalSpeed <= 10^6
//     The input is generated such that edges represents a valid tree.

import "fmt"

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
    n := len(edges) + 1
    g, res := make([][][]int, n), make([]int, 0, n)
    for i := range g {
        g[i] = [][]int{}
    }
    for _, e := range edges {
        g[e[0]] = append(g[e[0]], []int{e[1],e[2]})
        g[e[1]] = append(g[e[1]], []int{e[0],e[2]})
    }
    var dfs func(p, i, d int) int
    dfs = func(p, i, d int) int {
        res := 0
        if d % signalSpeed == 0{
            res++
        }
        for _, e := range g[i] {
            if e[0] == p {
                continue
            }
            res += dfs(i,e[0],d+e[1])
        }
        return res
    }
    for i := range g {
        sqSum,sum := 0,0
        for _, e := range g[i]{
            count := dfs(i,e[0],e[1])
            sqSum += count * count
            sum += count
        }
        res = append(res, (sum * sum - sqSum ) / 2)
    }
    return res
}

func countPairsOfConnectableServers1(edges [][]int, signalSpeed int) []int {
    type edge struct{ id, dis int }
    n := len(edges) + 1
    res, g := make([]int, n), make([][]edge, n)
    for _, e := range edges {
        g[e[0]] = append(g[e[0]], edge{e[1], e[2]})
        g[e[1]] = append(g[e[1]], edge{e[0], e[2]})
    }
    // dfs1(0, 0)
    for i := 0; i < n; i++ {
        glen := len(g[i])
        if glen == 1 {
            continue
        }
        var dfs func(int, int, int) int //当前结点，父节点，距离
        dfs = func(now_node, fa, dist int) int {
            tres := 0
            if dist % signalSpeed == 0 {
                tres++
            }
            for _, nextNode := range g[now_node] {
                if nextNode.id != fa {
                    tres += dfs(nextNode.id, now_node, dist+nextNode.dis)
                }
            }
            return tres
        }
        sum := 0
        for _, jdata := range g[i] {
            cnt := dfs(jdata.id, i, jdata.dis)
            res[i] += cnt * sum
            sum += cnt
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/01/21/example22.png" />
    // Input: edges = [[0,1,1],[1,2,5],[2,3,13],[3,4,9],[4,5,2]], signalSpeed = 1
    // Output: [0,4,6,6,4,0]
    // Explanation: Since signalSpeed is 1, count[c] is equal to the number of pairs of paths that start at c and do not share any edges.
    // In the case of the given path graph, count[c] is equal to the number of servers to the left of c multiplied by the servers to the right of c.
    fmt.Println(countPairsOfConnectableServers([][]int{{0,1,1},{1,2,5},{2,3,13},{3,4,9},{4,5,2}}, 1)) // [0,4,6,6,4,0]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/01/21/example11.png" />
    // Input: edges = [[0,6,3],[6,5,3],[0,3,1],[3,2,7],[3,1,6],[3,4,2]], signalSpeed = 3
    // Output: [2,0,0,0,0,0,2]
    // Explanation: Through server 0, there are 2 pairs of connectable servers: (4, 5) and (4, 6).
    // Through server 6, there are 2 pairs of connectable servers: (4, 5) and (0, 5).
    // It can be shown that no two servers are connectable through servers other than 0 and 6.
    fmt.Println(countPairsOfConnectableServers([][]int{{0,6,3},{6,5,3},{0,3,1},{3,2,7},{3,1,6},{3,4,2}}, 3)) // [2,0,0,0,0,0,2]
    
    fmt.Println(countPairsOfConnectableServers1([][]int{{0,1,1},{1,2,5},{2,3,13},{3,4,9},{4,5,2}}, 1)) // [0,4,6,6,4,0]
    fmt.Println(countPairsOfConnectableServers1([][]int{{0,6,3},{6,5,3},{0,3,1},{3,2,7},{3,1,6},{3,4,2}}, 3)) // [2,0,0,0,0,0,2]

}