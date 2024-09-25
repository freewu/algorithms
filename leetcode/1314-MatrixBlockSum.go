package main

// 1314. Matrix Block Sum
// Given a m x n matrix mat and an integer k, 
// return a matrix answer where each answer[i][j] is the sum of all elements mat[r][c] for:
//     i - k <= r <= i + k,
//     j - k <= c <= j + k, and
//     (r, c) is a valid position in the matrix.

// Example 1:
// Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1
// Output: [[12,21,16],[27,45,33],[24,39,28]]

// Example 2:
// Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 2
// Output: [[45,45,45],[45,45,45],[45,45,45]]

// Constraints:
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n, k <= 100
//     1 <= mat[i][j] <= 100

import "fmt"

func matrixBlockSum(mat [][]int, k int) [][]int {
    m, n, total := len(mat), len(mat[0]), make([][]int, len(mat)+1)
    for i := 0; i <= m; i++ {
        total[i] = make([]int, n+1)
    }
    for r := 1; r <= m; r++ {
        for c := 1; c <= n; c++ {
            total[r][c] = mat[r-1][c-1] + total[r-1][c] + total[r][c-1] - total[r-1][c-1]
        }
    }
    res := make([][]int, m)
    for i := 0; i < m; i++ {
        res[i] = make([]int, n)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            r1, c1 := max(0, r-k), max(0, c-k)
            r2, c2 := min(m-1, r+k), min(n-1, c+k)
            r1, c1 = r1+1, c1+1
            r2, c2 = r2+1, c2+1
            res[r][c] = total[r2][c2] - total[r2][c1-1] - total[r1-1][c2] + total[r1-1][c1-1]
        }
    }
    return res
}

func matrixBlockSum1(mat [][]int, k int) [][]int {
    m, n := len(mat), len(mat[0])
    matrix := make([][]int, m + 1)
    for i := range matrix {
        matrix[i] = make([]int, n + 1)
    }
    for i := 1; i < m + 1; i++{
        for j := 1; j < n + 1; j++{
            matrix[i][j] = matrix[i-1][j] + matrix[i][j-1] - matrix[i-1][j-1] + mat[i-1][j-1]
        }
    }
    res := make([][]int, m)
    for i := 0; i < len(res); i++{
        res[i] = make([]int, n)
    }
    for i := 0; i < len(res); i++{
        for j := 0; j < len(res[0]); j++{
            ltVal := matrix[max(i - k, 0)][max(j - k, 0)]
            ldVal := matrix[min(i + k + 1, m)][max(j - k, 0)]
            rtVal := matrix[max(i - k, 0)][min(j + k + 1, n)]
            rdVal := matrix[min(i + k + 1, m)][min(j + k + 1, n)]
            res[i][j] = rdVal - ldVal - rtVal + ltVal
        }
    }
    return res 
}

func main() {
    // Example 1:
    // Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1
    // Output: [[12,21,16],[27,45,33],[24,39,28]]
    mat1 := [][]int{
        {1,2,3},
        {4,5,6},
        {7,8,9},
    }
    fmt.Println(matrixBlockSum(mat1, 1)) // [[12,21,16],[27,45,33],[24,39,28]]
    // Example 2:
    // Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 2
    // Output: [[45,45,45],[45,45,45],[45,45,45]]
    mat2 := [][]int{
        {1,2,3},
        {4,5,6},
        {7,8,9},
    }
    fmt.Println(matrixBlockSum(mat2, 2)) // [[45,45,45],[45,45,45],[45,45,45]]

    fmt.Println(matrixBlockSum1(mat1, 1)) // [[12,21,16],[27,45,33],[24,39,28]]
    fmt.Println(matrixBlockSum1(mat2, 2)) // [[45,45,45],[45,45,45],[45,45,45]]
}