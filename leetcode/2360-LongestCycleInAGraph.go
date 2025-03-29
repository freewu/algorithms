package main

// 2360. Longest Cycle in a Graph
// You are given a directed graph of n nodes numbered from 0 to n - 1, where each node has at most one outgoing edge.

// The graph is represented with a given 0-indexed array edges of size n, indicating that there is a directed edge from node i to node edges[i]. 
// If there is no outgoing edge from node i, then edges[i] == -1.

// Return the length of the longest cycle in the graph. If no cycle exists, return -1.

// A cycle is a path that starts and ends at the same node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/06/08/graph4drawio-5.png" />
// Input: edges = [3,3,4,2,3]
// Output: 3
// Explanation: The longest cycle in the graph is the cycle: 2 -> 4 -> 3 -> 2.
// The length of this cycle is 3, so 3 is returned.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/06/07/graph4drawio-1.png" />
// Input: edges = [2,-1,3,1]
// Output: -1
// Explanation: There are no cycles in this graph.

// Constraints:
//     n == edges.length
//     2 <= n <= 10^5
//     -1 <= edges[i] < n
//     edges[i] != i

import "fmt"

func longestCycle(edges []int) int {
    res, visited := -1, make([]bool, len(edges)) // memory
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(edges); i++ {
        if !visited[i] { // each node has at most one outgoing edge, ignore the node which has been visited
            cur, step, path := i, 0, make(map[int]int) // track the path, calculate the length of the circle if exists
            visited[cur], path[cur] = true, step
            for edges[cur] != -1 {
                step++
                cur = edges[cur]
                if _, ok := path[cur]; ok {
                    res = max(res, step-path[cur])
                    break
                }
                if visited[cur] { break } // ignore for the same reason as above
                visited[cur], path[cur] = true, step
            }
        }
    }
    return res
}

func longestCycle1(edges []int) int {
    res, clock, n := -1, 1, len(edges)
    time := make([]int, n)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, t := range time {
        if t != 0 { continue }
        x := i
        for startTime := clock; x != -1; x = edges[x] {
            if time[x] != 0 {
                if time[x] >= startTime {
                    res = max(res, clock - time[x])
                }
                break
            }
            time[x] = clock
            clock++
        }
    }
    return res
}

func longestCycle2(edges []int) int {
    res, clock := -1, 1
    time := make([]int, len(edges)) // time[i] == 0 表示没有访问过
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, t := range time {
        if t > 0 { continue } // i 之前访问过，即使 i 在环上或者可以到达环，我们也已经更新了 res
        startTime := clock // 本轮循环的开始时间
        for ; edges[i] != -1 && time[i] == 0; i = edges[i] {
            time[i] = clock // 记录访问 i 的时间
            clock++
        }
        // 如果 time[i] < startTime，说明 i 在其他循环中访问过，无需重复计算
        // 否则说明 i 是本轮循环访问过的点，我们找到了一个环
        if time[i] >= startTime {
            res = max(res, clock - time[i]) // 当前时间减去上次访问 i 的时间，即为环长
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/06/08/graph4drawio-5.png" />
    // Input: edges = [3,3,4,2,3]
    // Output: 3
    // Explanation: The longest cycle in the graph is the cycle: 2 -> 4 -> 3 -> 2.
    // The length of this cycle is 3, so 3 is returned.
    fmt.Println(longestCycle([]int{3,3,4,2,3})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/06/07/graph4drawio-1.png" />
    // Input: edges = [2,-1,3,1]
    // Output: -1
    // Explanation: There are no cycles in this graph.
    fmt.Println(longestCycle([]int{2,-1,3,1})) // -1

    fmt.Println(longestCycle([]int{0,1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(longestCycle([]int{9,8,7,6,5,4,3,2,1,0})) // 2

    fmt.Println(longestCycle1([]int{3,3,4,2,3})) // 3
    fmt.Println(longestCycle1([]int{2,-1,3,1})) // -1
    fmt.Println(longestCycle1([]int{0,1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(longestCycle1([]int{9,8,7,6,5,4,3,2,1,0})) // 2

    fmt.Println(longestCycle2([]int{3,3,4,2,3})) // 3
    fmt.Println(longestCycle2([]int{2,-1,3,1})) // -1
    fmt.Println(longestCycle2([]int{0,1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(longestCycle2([]int{9,8,7,6,5,4,3,2,1,0})) // 2
}