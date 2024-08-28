package main

// 2492. Minimum Score of a Path Between Two Cities
// You are given a positive integer n representing n cities numbered from 1 to n. 
// You are also given a 2D array roads where roads[i] = [ai, bi, distancei] indicates 
// that there is a bidirectional road between cities ai and bi with a distance equal to distancei. 
// The cities graph is not necessarily connected.

// The score of a path between two cities is defined as the minimum distance of a road in this path.
// Return the minimum possible score of a path between cities 1 and n.

// Note:
//     A path is a sequence of roads between two cities.
//     It is allowed for a path to contain the same road multiple times, and you can visit cities 1 and n multiple times along the path.
//     The test cases are generated such that there is at least one path between 1 and n.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/10/12/graph11.png" />
// Input: n = 4, roads = [[1,2,9],[2,3,6],[2,4,5],[1,4,7]]
// Output: 5
// Explanation: The path from city 1 to 4 with the minimum score is: 1 -> 2 -> 4. The score of this path is min(9,5) = 5.
// It can be shown that no other path has less score.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/10/12/graph22.png" />
// Input: n = 4, roads = [[1,2,2],[1,3,4],[3,4,7]]
// Output: 2
// Explanation: The path from city 1 to 4 with the minimum score is: 1 -> 2 -> 1 -> 3 -> 4. The score of this path is min(2,2,4,7) = 2.

// Constraints:
//     2 <= n <= 10^5
//     1 <= roads.length <= 10^5
//     roads[i].length == 3
//     1 <= ai, bi <= n
//     ai != bi
//     1 <= distancei <= 10^4
//     There are no repeated edges.
//     There is at least one path between 1 and n.

import "fmt"

// 并查集
func minScore(n int, roads [][]int) int {
    parent, score, inf := make([]int, n), make([]int, n), 1 << 31
    for i := range parent {
        parent[i] = i
        score[i] = inf
    }
    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    union := func(x, y int, w int) {
        fx, fy := find(x), find(y)
        if fx != fy {
            parent[fx] = fy
        }
        score[fy] = min(score[fy], score[fx], w)
    }
    for _, e := range roads {
        u, v, w := e[0]-1, e[1]-1, e[2]
        union(u, v, w)
    }
    return score[find(0)]
}

// bfs
func minScore1(n int, roads [][]int) int {
    type Edge struct { to, dist int }
    graph := make([][]Edge, n+1)
    for _, road := range roads {
        a, b, dist := road[0], road[1], road[2]
        graph[a] = append(graph[a], Edge{b, dist})
        graph[b] = append(graph[b], Edge{a, dist})
    }
    res, visited, queue := 1 << 31, make(map[int]bool), []int{1}
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[0] // pop
            queue = queue[1:]
            if visited[node] {
                continue
            }
            visited[node] = true
            for _, edge := range graph[node] {
                if edge.dist < res {
                    res = edge.dist
                }
                if !visited[edge.to] {
                    queue = append(queue, edge.to)
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/10/12/graph11.png" />
    // Input: n = 4, roads = [[1,2,9],[2,3,6],[2,4,5],[1,4,7]]
    // Output: 5
    // Explanation: The path from city 1 to 4 with the minimum score is: 1 -> 2 -> 4. The score of this path is min(9,5) = 5.
    // It can be shown that no other path has less score.
    fmt.Println(minScore(4,[][]int{{1,2,9},{2,3,6},{2,4,5},{1,4,7}})) // 5
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/10/12/graph22.png" />
    // Input: n = 4, roads = [[1,2,2],[1,3,4],[3,4,7]]
    // Output: 2
    // Explanation: The path from city 1 to 4 with the minimum score is: 1 -> 2 -> 1 -> 3 -> 4. The score of this path is min(2,2,4,7) = 2.
    fmt.Println(minScore(4,[][]int{{1,2,2},{1,3,4},{3,4,7}})) // 2

    fmt.Println(minScore1(4,[][]int{{1,2,9},{2,3,6},{2,4,5},{1,4,7}})) // 5
    fmt.Println(minScore1(4,[][]int{{1,2,2},{1,3,4},{3,4,7}})) // 2
}