package main

// 1901. Find a Peak Element II
// A peak element in a 2D grid is an element that is strictly greater than all of its adjacent neighbors to the left, right, top, and bottom.

// Given a 0-indexed m x n matrix mat where no two adjacent cells are equal, 
// find any peak element mat[i][j] and return the length 2 array [i,j].

// You may assume that the entire matrix is surrounded by an outer perimeter with the value -1 in each cell.

// You must write an algorithm that runs in O(m log(n)) or O(n log(m)) time.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/08/1.png" />
// Input: mat = [[1,4],[3,2]]
// Output: [0,1]
// Explanation: Both 3 and 4 are peak elements so [1,0] and [0,1] are both acceptable answers.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/07/3.png" />
// Input: mat = [[10,20,15],[21,30,14],[7,16,32]]
// Output: [1,1]
// Explanation: Both 30 and 32 are peak elements so [1,1] and [2,2] are both acceptable answers.

// Constraints:
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 500
//     1 <= mat[i][j] <= 10^5
//     No two adjacent cells are equal.

import "fmt"

// func findPeakGrid(mat [][]int) []int {
//     m, n := len(mat), len(mat[0])
//     l, r := 0, n - 1
//     isPeak := func(mat [][]int, i int, j int) bool {
//         m, n := len(mat), len(mat[0])
//         return (i == 0 || mat[i - 1][j] < mat[i][j]) && 
//                (i == m - 1 || mat[i + 1][j] < mat[i][j]) && 
//                (j == 0 || mat[i][j - 1] < mat[i][j]) && 
//                (j == n - 1 || mat[i][j + 1] < mat[i][j])
//     }
//     for l <= r {
//         mid := (l + r) / 2
//         for i := 0; i < m; i++ {
//             if isPeak(mat, i, mid) {
//                 return []int{  mid, i }
//             }
//         }
//         if mid == 0 {
//             l = mid + 1
//             continue 
//         }
//         for i := 0; i < m; i++ {
//             if mat[i][mid - 1] > mat[i][mid] && (i == 0 || 
//                mat[i - 1][mid] < mat[i][mid] && (i == m - 1 || mat[i + 1][mid] < mat[i][mid])) {
//                 r = mid - 1
//                 break
//             }
//         }       
//         l = mid + 1
//     }
//     return []int{ -1, -1 }
// }

func findPeakGrid(mat [][]int) []int {
    n, m := len(mat), len(mat[0])
    low, high, mid := 0, m - 1, 0
    findMaxIndex := func (mat [][]int, n int, col int) int {
        mx, index := -1,  -1
        for i := 0; i < n; i++ {
            if mat[i][col] > mx { mx, index = mat[i][col], i }
        }
        return index
    }
    for low <= high {
        mid = (low + high) / 2
        left, right, maxRowIndex := -1, -1, findMaxIndex(mat, n, mid)
        if mid - 1 >= 0 { left = mat[maxRowIndex][mid - 1] }
        if mid + 1 < m  { right = mat[maxRowIndex][mid + 1] }
        if mat[maxRowIndex][mid] > left && mat[maxRowIndex][mid] > right {
            return []int{ maxRowIndex, mid }
        } else if mat[maxRowIndex][mid] < left {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return []int{ -1, -1 }
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/08/1.png" />
    // Input: mat = [[1,4],[3,2]]
    // Output: [0,1]
    // Explanation: Both 3 and 4 are peak elements so [1,0] and [0,1] are both acceptable answers.
    fmt.Println(findPeakGrid([][]int{{1,4},{3,2}})) // [0,1]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/07/3.png" />
    // Input: mat = [[10,20,15],[21,30,14],[7,16,32]]
    // Output: [1,1]
    // Explanation: Both 30 and 32 are peak elements so [1,1] and [2,2] are both acceptable answers.
    fmt.Println(findPeakGrid([][]int{{10,20,15},{21,30,14},{7,16,32}})) // [1,1]

    fmt.Println(findPeakGrid([][]int{{1,1},{1,1}})) // [-1 -1]
    fmt.Println(findPeakGrid([][]int{{10,10,10},{10,2,10},{10,10,10}})) // [-1 -1]
    fmt.Println(findPeakGrid([][]int{{10,10,10,10,10},{10,2,2,2,10},{10,2,10,2,10},{10,2,2,2,10},{10,10,10,10,10}})) // [2 2]
    fmt.Println(findPeakGrid([][]int{{70,50,40,30,20},{100,1,2,3,4}})) //[1,0]
}