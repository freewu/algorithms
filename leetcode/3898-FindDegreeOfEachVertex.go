package  main

// 3898. Find the Degree of Each Vertex
// You are given a 2D integer array matrix of size n x n representing the adjacency matrix of an undirected graph with n vertices labeled from 0 to n - 1.
//     1. matrix[i][j] = 1 indicates that there is an edge between vertices i and j.
//     2. matrix[i][j] = 0 indicates that there is no edge between vertices i and j.

// The degree of a vertex is the number of edges connected to it.

// Return an integer array ans of size n where ans[i] represents the degree of vertex i.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2026/03/14/g41f.png" />
// Input: matrix = [[0,1,1],[1,0,1],[1,1,0]]
// Output: [2,2,2]
// Explanation:
// Vertex 0 is connected to vertices 1 and 2, so its degree is 2.
// Vertex 1 is connected to vertices 0 and 2, so its degree is 2.
// Vertex 2 is connected to vertices 0 and 1, so its degree is 2.
// Thus, the answer is [2, 2, 2].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2026/03/14/g42f.png" />
// Input: matrix = [[0,1,0],[1,0,0],[0,0,0]]
// Output: [1,1,0]
// Explanation:
// Vertex 0 is connected to vertex 1, so its degree is 1.
// Vertex 1 is connected to vertex 0, so its degree is 1.
// Vertex 2 is not connected to any vertex, so its degree is 0.
// Thus, the answer is [1, 1, 0].

// Example 3:
// Input: matrix = [[0]]
// Output: [0]
// Explanation:
// There is only one vertex and it has no edges connected to it. Thus, the answer is [0].

// Constraints:
//     1 <= n == matrix.length == matrix[i].length <= 100​​​​​​​
//     ​​​​​​​matrix[i][i] == 0
//     matrix[i][j] is either 0 or 1
//     matrix[i][j] == matrix[j][i]

import "fmt"

func findDegrees(matrix [][]int) []int {
    n := len(matrix)
    res := make([]int, n)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 1 {
                res[i]++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2026/03/14/g41f.png" />
    // Input: matrix = [[0,1,1],[1,0,1],[1,1,0]]
    // Output: [2,2,2]
    // Explanation:
    // Vertex 0 is connected to vertices 1 and 2, so its degree is 2.
    // Vertex 1 is connected to vertices 0 and 2, so its degree is 2.
    // Vertex 2 is connected to vertices 0 and 1, so its degree is 2.
    // Thus, the answer is [2, 2, 2].
    fmt.Println(findDegrees([][]int{{0,1,1},{1,0,1},{1,1,0}})) // [2,2,2]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2026/03/14/g42f.png" />
    // Input: matrix = [[0,1,0],[1,0,0],[0,0,0]]
    // Output: [1,1,0]
    // Explanation:
    // Vertex 0 is connected to vertex 1, so its degree is 1.
    // Vertex 1 is connected to vertex 0, so its degree is 1.
    // Vertex 2 is not connected to any vertex, so its degree is 0.
    // Thus, the answer is [1, 1, 0].
        fmt.Println(findDegrees([][]int{{0,1,0},{1,0,0},{0,0,0}})) // [1,1,0]
    // Example 3:
    // Input: matrix = [[0]]
    // Output: [0]
    // Explanation:
    // There is only one vertex and it has no edges connected to it. Thus, the answer is [0].
    fmt.Println(findDegrees([][]int{{0}})) // [0]
}