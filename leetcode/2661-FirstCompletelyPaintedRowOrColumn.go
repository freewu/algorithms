package main

// 2661. First Completely Painted Row or Column
// You are given a 0-indexed integer array arr, and an m x n integer matrix mat. 
// arr and mat both contain all the integers in the range [1, m * n].

// Go through each index i in arr starting from index 0 and paint the cell in mat containing the integer arr[i].

// Return the smallest index i at which either a row or a column will be completely painted in mat.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/01/18/grid1.jpg" />
// image explanation for example 1
// Input: arr = [1,3,4,2], mat = [[1,4],[2,3]]
// Output: 2
// Explanation: The moves are shown in order, and both the first row and second column of the matrix become fully painted at arr[2].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/01/18/grid2.jpg" />
// image explanation for example 2
// Input: arr = [2,8,7,4,1,3,5,6,9], mat = [[3,2,5],[1,4,6],[8,7,9]]
// Output: 3
// Explanation: The second column becomes fully painted at arr[3].

// Constraints:
//     m == mat.length
//     n = mat[i].length
//     arr.length == m * n
//     1 <= m, n <= 10^5
//     1 <= m * n <= 10^5
//     1 <= arr[i], mat[r][c] <= m * n
//     All the integers of arr are unique.
//     All the integers of mat are unique.

import "fmt"

func firstCompleteIndex(arr []int, mat [][]int) int {
    m, n := len(mat), len(mat[0])
    num2pos, row2cnt, col2cnt := make([]int, m * n + 1), make([]int, m), make([]int, n)
    for i := range row2cnt {
        row2cnt[i] = n
    }
    for j := range col2cnt {
        col2cnt[j] = m
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            num2pos[mat[i][j]] = i * n + j
        }
    }
    for i, v := range arr {
        pos := num2pos[v]
        r, c := pos / n, pos % n
        row2cnt[r]--
        col2cnt[c]--
        if row2cnt[r] == 0 || col2cnt[c] == 0 {
            return i
        }
    }
    return -1
}

func firstCompleteIndex1(arr []int, mat [][]int) int {
    res, m, n := 0, len(mat), len(mat[0])
    positions := make([][2]int, len(arr)+1) // First we need to know which (row, col) each number refers to.
    for i := 0; i < m; i++ {
      for j := 0; j < n; j++ {
        positions[mat[i][j]] = [2]int{i, j}
      }
    }
    // rowPainted keeps track of the number of cells painted in each row, colPainted keeps track of the number of cells painted in each column
    rowPainted, colPainted := make([]int, m), make([]int, n)
    for index, id := range arr {
        i, j := positions[id][0], positions[id][1]
        rowPainted[i]++
        colPainted[j]++
        if rowPainted[i] == n || colPainted[j] == m {
            res = index
            break
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/01/18/grid1.jpg" />
    // image explanation for example 1
    // Input: arr = [1,3,4,2], mat = [[1,4],[2,3]]
    // Output: 2
    // Explanation: The moves are shown in order, and both the first row and second column of the matrix become fully painted at arr[2].
    fmt.Println(firstCompleteIndex([]int{1,3,4,2}, [][]int{{1,4},{2,3}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/01/18/grid2.jpg" />
    // image explanation for example 2
    // Input: arr = [2,8,7,4,1,3,5,6,9], mat = [[3,2,5],[1,4,6],[8,7,9]]
    // Output: 3
    // Explanation: The second column becomes fully painted at arr[3].
    fmt.Println(firstCompleteIndex([]int{2,8,7,4,1,3,5,6,9}, [][]int{{3,2,5},{1,4,6},{8,7,9}})) // 3

    fmt.Println(firstCompleteIndex1([]int{1,3,4,2}, [][]int{{1,4},{2,3}})) // 2
    fmt.Println(firstCompleteIndex1([]int{2,8,7,4,1,3,5,6,9}, [][]int{{3,2,5},{1,4,6},{8,7,9}})) // 3
}