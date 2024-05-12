package main

// 1466. Reorder Routes to Make All Paths Lead to the City Zero
// There are n cities numbered from 0 to n - 1 and n - 1 roads such that there is only one way to travel between two different cities (this network form a tree). 
// Last year, The ministry of transport decided to orient the roads in one direction because they are too narrow.

// Roads are represented by connections where connections[i] = [ai, bi] represents a road from city ai to city bi.

// This year, there will be a big event in the capital (city 0), and many people want to travel to this city.

// Your task consists of reorienting some roads such that each city can visit the city 0. 
// Return the minimum number of edges changed.

// It's guaranteed that each city can reach city 0 after reorder.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/05/13/sample_1_1819.png" />
// Input: n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]
// Output: 3
// Explanation: Change the direction of edges show in red such that each node can reach the node 0 (capital).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/05/13/sample_2_1819.png" />
// Input: n = 5, connections = [[1,0],[1,2],[3,2],[3,4]]
// Output: 2
// Explanation: Change the direction of edges show in red such that each node can reach the node 0 (capital).

// Example 3:
// Input: n = 3, connections = [[1,0],[2,0]]
// Output: 0
 
// Constraints:
//     2 <= n <= 5 * 10^4
//     connections.length == n - 1
//     connections[i].length == 2
//     0 <= ai, bi <= n - 1
//     ai != bi

import "fmt"

// dfs
func minReorder(n int, connections [][]int) int {
    res, visited, graph := 0, make([]bool, n), make([][]int, n)
    for _, val := range connections {
        graph[val[0]] = append(graph[val[0]], val[1])
        graph[val[1]] = append(graph[val[1]], -val[0])
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    var dfs func(number int)
    dfs = func(number int) {
        visited[number] = true
        for i := 0; i < len(graph[number]); i++ {
            if !visited[abs(graph[number][i])] {
                dfs(abs(graph[number][i]))
                if graph[number][i] > 0 {
                    res++
                }
            }
        }
    }
    for i := 0; i < len(graph); i++ {
        if !visited[i] {
            dfs(i)
        }
    }
    return res
}

// bfs
func minReorder1(n int, connections [][]int) int {
    res, visited, graph, queue := 0, make([]bool, n), make([][]int, n), []int{}
    for _, val := range connections {
        graph[val[0]] = append(graph[val[0]], val[1])
        graph[val[1]] = append(graph[val[1]], -val[0])
    }
    for _, val := range graph[0] {
        queue = append(queue, val)
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            cur := queue[0]
            queue = queue[1:]
            if cur > 0 && !visited[abs(cur)] {
                res++
            }
            if len(graph[abs(cur)]) > 0 && !visited[abs(cur)] {
                for _, val := range graph[abs(cur)] {
                    queue = append(queue, val)
                }
            }
            visited[abs(cur)] = true
        }

    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/05/13/sample_1_1819.png" />
    // Input: n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]
    // Output: 3
    // Explanation: Change the direction of edges show in red such that each node can reach the node 0 (capital).
    fmt.Println(minReorder(6,[][]int{{0,1},{1,3},{2,3},{4,0},{4,5}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/05/13/sample_2_1819.png" />
    // Input: n = 5, connections = [[1,0],[1,2],[3,2],[3,4]]
    // Output: 2
    // Explanation: Change the direction of edges show in red such that each node can reach the node 0 (capital).
    fmt.Println(minReorder(5,[][]int{{1,0},{1,2},{3,2},{3,4}})) // 2
    // Example 3:
    // Input: n = 3, connections = [[1,0],[2,0]]
    // Output: 0
    fmt.Println(minReorder(3,[][]int{{1,0},{2,0}})) // 0

    fmt.Println(minReorder1(6,[][]int{{0,1},{1,3},{2,3},{4,0},{4,5}})) // 3
    fmt.Println(minReorder1(5,[][]int{{1,0},{1,2},{3,2},{3,4}})) // 2
    fmt.Println(minReorder1(3,[][]int{{1,0},{2,0}})) // 0
}