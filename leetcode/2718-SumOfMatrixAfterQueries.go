package main

// 2718. Sum of Matrix After Queries
// You are given an integer n and a 0-indexed 2D array queries where queries[i] = [typei, indexi, vali].

// Initially, there is a 0-indexed n x n matrix filled with 0's. 
// For each query, you must apply one of the following changes:
//     if typei == 0, set the values in the row with indexi to vali, overwriting any previous values.
//     if typei == 1, set the values in the column with indexi to vali, overwriting any previous values.

// Return the sum of integers in the matrix after all queries are applied.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/05/11/exm1.png" />
// Input: n = 3, queries = [[0,0,1],[1,2,2],[0,2,3],[1,0,4]]
// Output: 23
// Explanation: The image above describes the matrix after each query. 
// The sum of the matrix after all queries are applied is 23. 

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/05/11/exm2.png" />
// Input: n = 3, queries = [[0,0,4],[0,1,2],[1,0,1],[0,2,3],[1,2,1]]
// Output: 17
// Explanation: The image above describes the matrix after each query. 
// The sum of the matrix after all queries are applied is 17.
 
// Constraints:
//     1 <= n <= 10^4
//     1 <= queries.length <= 5 * 10^4
//     queries[i].length == 3
//     0 <= typei <= 1
//     0 <= indexi < n
//     0 <= vali <= 10^5

import "fmt"

func matrixSumQueries(n int, queries [][]int) int64 {
    res, row,col, m := 0, n, n, len(queries)
    rowval,colval := make([]bool,n), make([]bool,n)
    for i := 0; i < n; i++ {
        rowval[i], colval[i] = false, false
    }
    for i := m-1; i >= 0; i-- {
        if queries[i][0] == 0 && rowval[queries[i][1]] == false {
            res += queries[i][2] * col
            rowval[queries[i][1]] = true
            row--
        }
        if queries[i][0] ==1 && colval[queries[i][1]] == false {
            res += queries[i][2] * row
            colval[queries[i][1]] = true
            col--
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/05/11/exm1.png" />
    // Input: n = 3, queries = [[0,0,1],[1,2,2],[0,2,3],[1,0,4]]
    // Output: 23
    // Explanation: The image above describes the matrix after each query. 
    // The sum of the matrix after all queries are applied is 23. 
    fmt.Println(matrixSumQueries(3, [][]int{{0,0,1},{1,2,2},{0,2,3},{1,0,4}})) // 23
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/05/11/exm2.png" />
    // Input: n = 3, queries = [[0,0,4],[0,1,2],[1,0,1],[0,2,3],[1,2,1]]
    // Output: 17
    // Explanation: The image above describes the matrix after each query. 
    // The sum of the matrix after all queries are applied is 17.
    fmt.Println(matrixSumQueries(3, [][]int{{0,0,4},{0,1,2},{1,0,1},{0,2,3},{1,2,1}})) // 17
}