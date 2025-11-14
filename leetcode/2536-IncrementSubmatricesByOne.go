package main

// 2536. Increment Submatrices by One
// You are given a positive integer n, indicating that we initially have an n x n 0-indexed integer matrix mat filled with zeroes.

// You are also given a 2D integer array query. 
// For each query[i] = [row1i, col1i, row2i, col2i], you should do the following operation:
//     Add 1 to every element in the submatrix with the top left corner (row1i, col1i) and the bottom right corner (row2i, col2i). 
//     That is, add 1 to mat[x][y] for all row1i <= x <= row2i and col1i <= y <= col2i.

// Return the matrix mat after performing every query.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/11/24/p2example11.png" />
// Input: n = 3, queries = [[1,1,2,2],[0,0,1,1]]
// Output: [[1,1,0],[1,2,1],[0,1,1]]
// Explanation: The diagram above shows the initial matrix, the matrix after the first query, and the matrix after the second query.
// - In the first query, we add 1 to every element in the submatrix with the top left corner (1, 1) and bottom right corner (2, 2).
// - In the second query, we add 1 to every element in the submatrix with the top left corner (0, 0) and bottom right corner (1, 1).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/11/24/p2example22.png" />
// Input: n = 2, queries = [[0,0,1,1]]
// Output: [[1,1],[1,1]]
// Explanation: The diagram above shows the initial matrix and the matrix after the first query.
// - In the first query we add 1 to every element in the matrix.

// Constraints:
//     1 <= n <= 500
//     1 <= queries.length <= 10^4
//     0 <= row1i <= row2i < n
//     0 <= col1i <= col2i < n

import "fmt"

func rangeAddQueries(n int, queries [][]int) [][]int {
    res := make([][]int, n)
    for i := 0; i < n; i++ {
        res[i] = make([]int, n)
    }
    for _, v := range queries {
        for i := v[0]; i <= v[2]; i++ {
            for j := v[1]; j <= v[3]; j++ {
                res[i][j] = res[i][j] + 1
            }
        }
    }
    return res 
}

func rangeAddQueries1(n int, queries [][]int) [][]int {
    res := make([][]int, n + 2)
    for i := range res {
        res[i] = make([]int, n + 2)
    }
    for _, q := range queries {
        r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
        res[r1+1][c1+1]++
        res[r1+1][c2+2]--
        res[r2+2][c1+1]--
        res[r2+2][c2+2]++
    }
    for i := 1; i <= n + 1; i++ {
        for j := 1; j <= n+1; j++ {
            res[i][j] += (res[i-1][j] + res[i][j-1] - res[i-1][j-1])
        }
    }
    res = res[1 : n + 1]
    for i, row := range res {
        res[i] = row[1 : n + 1]
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/11/24/p2example11.png" />
    // Input: n = 3, queries = [[1,1,2,2],[0,0,1,1]]
    // Output: [[1,1,0],[1,2,1],[0,1,1]]
    // Explanation: The diagram above shows the initial matrix, the matrix after the first query, and the matrix after the second query.
    // - In the first query, we add 1 to every element in the submatrix with the top left corner (1, 1) and bottom right corner (2, 2).
    // - In the second query, we add 1 to every element in the submatrix with the top left corner (0, 0) and bottom right corner (1, 1).
    fmt.Println(rangeAddQueries(3, [][]int{{1,1,2,2},{0,0,1,1}})) // [[1,1,0],[1,2,1],[0,1,1]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/11/24/p2example22.png" />
    // Input: n = 2, queries = [[0,0,1,1]]
    // Output: [[1,1],[1,1]]
    // Explanation: The diagram above shows the initial matrix and the matrix after the first query.
    // - In the first query we add 1 to every element in the matrix.
    fmt.Println(rangeAddQueries(2, [][]int{{0,0,1,1}})) // [[1,1],[1,1]]

    fmt.Println(rangeAddQueries(9, [][]int{{1,2,3,4,5,6,7,8,9}})) // [[0 0 0 0 0 0 0 0 0] [0 0 1 1 1 0 0 0 0] [0 0 1 1 1 0 0 0 0] [0 0 1 1 1 0 0 0 0] [0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0]]

    fmt.Println(rangeAddQueries1(3, [][]int{{1,1,2,2},{0,0,1,1}})) // [[1,1,0],[1,2,1],[0,1,1]]
    fmt.Println(rangeAddQueries1(2, [][]int{{0,0,1,1}})) // [[1,1],[1,1]]
    fmt.Println(rangeAddQueries1(9, [][]int{{1,2,3,4,5,6,7,8,9}})) // [[0 0 0 0 0 0 0 0 0] [0 0 1 1 1 0 0 0 0] [0 0 1 1 1 0 0 0 0] [0 0 1 1 1 0 0 0 0] [0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0]]
}