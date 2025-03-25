package main

// 3493. Properties Graph
// You are given a 2D integer array properties having dimensions n x m and an integer k.

// Define a function intersect(a, b) that returns the number of distinct integers common to both arrays a and b.

// Construct an undirected graph where each index i corresponds to properties[i]. 
// There is an edge between node i and node j if and only if intersect(properties[i], properties[j]) >= k, where i and j are in the range [0, n - 1] and i != j.

// Return the number of connected components in the resulting graph.

// Example 1:
// Input: properties = [[1,2],[1,1],[3,4],[4,5],[5,6],[7,7]], k = 1
// Output: 3
// Explanation:
// The graph formed has 3 connected components:
// <img src="https://assets.leetcode.com/uploads/2025/02/27/image.png" />

// Example 2:
// Input: properties = [[1,2,3],[2,3,4],[4,3,5]], k = 2
// Output: 1
// Explanation:
// The graph formed has 1 connected component:
// <img src="https://assets.leetcode.com/uploads/2025/02/27/screenshot-from-2025-02-27-23-58-34.png" />

// Example 3:
// Input: properties = [[1,1],[1,1]], k = 2
// Output: 2
// Explanation:
// intersect(properties[0], properties[1]) = 1, which is less than k. 
// This means there is no edge between properties[0] and properties[1] in the graph.

// Constraints:
//     1 <= n == properties.length <= 100
//     1 <= m == properties[i].length <= 100
//     1 <= properties[i][j] <= 100
//     1 <= k <= m

import "fmt"

func numberOfComponents(properties [][]int, k int) int {
    n := len(properties)
    ps := []map[int]struct{}{}
    for _, row := range properties {
        elem := map[int]struct{}{}
        for _, v := range row {
            elem[v] = struct{}{}
        }
        ps = append(ps, elem)
    }
    parent := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
    }
    var find func(p int) int
    find = func(p int) int {
        if (p != parent[p]) {
            parent[p] = find(parent[p])
        }
        return parent[p]
    }
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            count := 0
            for x, _ := range ps[i] {
                if _, ok := ps[j][x]; ok {
                    count++
                }
            }
            if count >= k {
                parent[find(i)] = find(j)
            }
        }
    }
    freq := map[int]int{}
    for _, v := range parent {
        freq[find(v)]++
    }
    return len(freq)
}

// bfs
func numberOfComponents1(properties [][]int, k int) int {
    res, n := 0, len(properties)
    presence := make([][101]bool, n)
    for i := 0; i < n; i++ {
        for _, v := range properties[i] {
            presence[i][v] = true
        }
    }
    adj := make([][]int, n)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            count := 0
            for l := 1; l <= 100; l++ {
                if presence[i][l] && presence[j][l] {
                    count++
                    if count >= k {
                        adj[i] = append(adj[i], j)
                        adj[j] = append(adj[j], i)
                        break
                    }
                }
            }
        }
    }
    visited := make([]bool, n)
    for i := 0; i < n; i++ {
        if !visited[i] {
            res++
            queue := []int{i}
            visited[i] = true
            for len(queue) > 0 {
                node := queue[0]
                queue = queue[1:]
                for _, next := range adj[node] {
                    if !visited[next] {
                        visited[next] = true
                        queue = append(queue, next)
                    }
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: properties = [[1,2],[1,1],[3,4],[4,5],[5,6],[7,7]], k = 1
    // Output: 3
    // Explanation:
    // The graph formed has 3 connected components:
    // <img src="https://assets.leetcode.com/uploads/2025/02/27/image.png" />
    fmt.Println(numberOfComponents([][]int{{1,2},{1,1},{3,4},{4,5},{5,6},{7,7}}, 1)) // 3
    // Example 2:
    // Input: properties = [[1,2,3],[2,3,4],[4,3,5]], k = 2
    // Output: 1
    // Explanation:
    // The graph formed has 1 connected component:
    // <img src="https://assets.leetcode.com/uploads/2025/02/27/screenshot-from-2025-02-27-23-58-34.png" />
    fmt.Println(numberOfComponents([][]int{{1,2,3},{2,3,4},{4,3,5}}, 2)) // 1
    // Example 3:
    // Input: properties = [[1,1],[1,1]], k = 2
    // Output: 2
    // Explanation:
    // intersect(properties[0], properties[1]) = 1, which is less than k. 
    // This means there is no edge between properties[0] and properties[1] in the graph.
    fmt.Println(numberOfComponents([][]int{{1,1},{1,1}}, 2)) // 2

    fmt.Println(numberOfComponents1([][]int{{1,2},{1,1},{3,4},{4,5},{5,6},{7,7}}, 1)) // 3
    fmt.Println(numberOfComponents1([][]int{{1,2,3},{2,3,4},{4,3,5}}, 2)) // 1
    fmt.Println(numberOfComponents1([][]int{{1,1},{1,1}}, 2)) // 2
}