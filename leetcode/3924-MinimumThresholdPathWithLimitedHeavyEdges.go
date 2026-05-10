package main

// 3924. Minimum Threshold Path With Limited Heavy Edges
// There is an undirected weighted graph with n nodes labeled from 0 to n - 1.

// The graph is represented by a 2D integer array edges, where each edge edges[i] = [ui, vi, w​​​​​​​i] indicates that there is an undirected edge between nodes ui and vi with weight w​​​​​​​i.

// You are also given integers source, target and k.

// A threshold value determines whether an edge is considered light or heavy:
//     1. An edge is light if its weight is less than or equal to threshold.
//     2. An edge is heavy if its weight is greater than threshold.

// A path from source to target is valid if it contains at most k heavy edges.

// Return the minimum integer threshold such that at least one valid path exists from source to target. 
// If no such path exists, return -1.

// Example 1:​​​​​​​​​​​​​​
// ​​​​​​<img src="https://assets.leetcode.com/uploads/2025/10/13/g6.png" />
// Input: n = 6, edges = [[0,1,5],[1,2,3],[3,4,4],[4,5,1],[1,4,2]], source = 0, target = 3, k = 1
// Output: 4
// Explanation:
// The minimum threshold such that a path from node 0 to node 3 uses at most 1 heavy edge is 4.
// Light edges: [1, 2, 3], [3, 4, 4], [4, 5, 1], [1, 4, 2]
// Heavy edges: [0, 1, 5]
// A valid path is 0 → 1 → 4 → 3. It uses only 1 heavy edge ([0, 1, 5]), which satisfies the limit k = 1.
// Any smaller threshold would make it impossible to reach node 3 without exceeding 1 heavy edge.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2025/10/12/g3_f.png" />
// Input: n = 6, edges = [[0,1,3],[1,2,4],[3,4,5],[4,5,6]], source = 0, target = 4, k = 1
// Output: -1
// Explanation:
// There is no path from node 0 to node 4. Since the target cannot be reached, the output is -1.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2025/10/12/g5.png" />
// Input: n = 4, edges = [[0,1,2],[1,2,2],[2,3,2],[3,0,2]], source = 0, target = 0, k = 0
// Output: 0
// Explanation:
// The source and target are the same node. No edges need to be traversed, so the minimum threshold is 0.

// Constraints:
//     1 <= n <= 10^3​​​​​​​
//     0 <= edges.length <= 10^3​​​​​​​
//     edges[i] = [ui, vi, wi]
//     0 <= ui, vi​​​​​​​ <= n - 1
//     1 <= wi​​​​​​​ <= 10^9​​​​​​​
//     0 <= source, target <= n - 1
//     0 <= k <= edges.length

import "fmt"
import "sort"

func minimumThreshold(n int, edges [][]int, source int, target int, k int) int {
    type Edge struct{ to, weight int }
    graph := make([][]Edge, n)
    mx := 0
    for _, e := range edges {
        x, y, wt := e[0], e[1], e[2]
        graph[x] = append(graph[x], Edge{y, wt})
        graph[y] = append(graph[y], Edge{x, wt})
        mx = max(mx, wt)
    }
    dis := make([]int, n)
    res := sort.Search(mx + 1, func(threshold int) bool {
        for i := range dis {
            dis[i] = 1 << 61
        }
        dis[source] = 0
        type Pair struct{ x, distance int }
        ql, qr := []Pair{{source, dis[source]}}, []Pair{} // 模拟双端队列

        for len(ql) > 0 || len(qr) > 0 {
            var p Pair
            if len(ql) > 0 {
                ql, p = ql[:len(ql)-1], ql[len(ql)-1] // 队首出
            } else {
                p, qr = qr[0], qr[1:] // 队尾出
            }
            x := p.x
            if x == target {
                return true
            }
            if p.distance > dis[x] {
                continue
            }
            for _, e := range graph[x] {
                y, weight := e.to, 0
                if e.weight > threshold {
                    weight = 1
                }
                newDis := p.distance + weight
                if newDis < dis[y] {
                    dis[y] = newDis
                    if weight == 0 {
                        ql = append(ql, Pair{y, newDis}) // 加到队首
                    } else if newDis <= k {
                        qr = append(qr, Pair{y, newDis}) // 加到队尾
                    }
                }
            }
        }
        return false
    })
    if res > mx { // 图不连通
        return -1
    }
    return res
}

func main() {
    // Example 1:​​​​​​​​​​​​​​
    // ​​​​​​<img src="https://assets.leetcode.com/uploads/2025/10/13/g6.png" />
    // Input: n = 6, edges = [[0,1,5],[1,2,3],[3,4,4],[4,5,1],[1,4,2]], source = 0, target = 3, k = 1
    // Output: 4
    // Explanation:
    // The minimum threshold such that a path from node 0 to node 3 uses at most 1 heavy edge is 4.
    // Light edges: [1, 2, 3], [3, 4, 4], [4, 5, 1], [1, 4, 2]
    // Heavy edges: [0, 1, 5]
    // A valid path is 0 → 1 → 4 → 3. It uses only 1 heavy edge ([0, 1, 5]), which satisfies the limit k = 1.
    // Any smaller threshold would make it impossible to reach node 3 without exceeding 1 heavy edge.
    fmt.Println(minimumThreshold(6, [][]int{{0,1,5},{1,2,3},{3,4,4},{4,5,1},{1,4,2}}, 0, 3, 1)) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2025/10/12/g3_f.png" />
    // Input: n = 6, edges = [[0,1,3],[1,2,4],[3,4,5],[4,5,6]], source = 0, target = 4, k = 1
    // Output: -1
    // Explanation:
    // There is no path from node 0 to node 4. Since the target cannot be reached, the output is -1.
    fmt.Println(minimumThreshold(6, [][]int{{0,1,3},{1,2,4},{3,4,5},{4,5,6}}, 0, 4, 1)) // -1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2025/10/12/g5.png" />
    // Input: n = 4, edges = [[0,1,2],[1,2,2],[2,3,2],[3,0,2]], source = 0, target = 0, k = 0
    // Output: 0
    // Explanation:
    // The source and target are the same node. No edges need to be traversed, so the minimum threshold is 0.
    fmt.Println(minimumThreshold(4, [][]int{{0,1,2},{1,2,2},{2,3,2},{3,0,2}}, 0, 0, 0)) // 0
}