package main

// 1886. Determine Whether Matrix Can Be Obtained By Rotation
// Given two n x n binary matrices mat and target, 
// return true if it is possible to make mat equal to target by rotating mat in 90-degree increments, or false otherwise.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/20/grid3.png" />
// Input: mat = [[0,1],[1,0]], target = [[1,0],[0,1]]
// Output: true
// Explanation: We can rotate mat 90 degrees clockwise to make mat equal target.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/05/20/grid4.png" />
// Input: mat = [[0,1],[1,1]], target = [[1,0],[0,1]]
// Output: false
// Explanation: It is impossible to make mat equal to target by rotating mat.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/05/26/grid4.png" />
// Input: mat = [[0,0,0],[0,1,0],[1,1,1]], target = [[1,1,1],[0,1,0],[0,0,0]]
// Output: true
// Explanation: We can rotate mat 90 degrees clockwise two times to make mat equal target.

// Constraints:
//     n == mat.length == target.length
//     n == mat[i].length == target[i].length
//     1 <= n <= 10
//     mat[i][j] and target[i][j] are either 0 or 1.

import "fmt"
import "reflect"

func findRotation(mat [][]int, target [][]int) bool {
    rotate90 := func (arr [][]int) [][]int {
        n, m := len(arr), len(arr[0])
        res := make([][]int, n)
        for i := 0; i < m; i++ {
            res[i] = make([]int, n)
            for j := 0; j < n; j++ {
                res[i][j] = arr[n-1-j][i]
            }
        }
        return res
    }
    for i := 0; i < 4; i++ {
        if reflect.DeepEqual(mat, target) {
            return true
        }
        mat = rotate90(mat)
    }
    return false
}

func findRotation1(mat [][]int, target [][]int) bool {
    n, fs := len(mat) ,[4]bool{true, true, true, true}
    rotate := func (i, j, n int) (int, int) {
        return j, n-i-1
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            i2, j2 := i, j
            for k := 0; k < 4; k++ {
                i2, j2 = rotate(i2, j2, n)
                if target[i][j] != mat[i2][j2] {
                    fs[k] = false
                }
            }
        }
    }
    return fs[1] || fs[2] || fs[3] || fs[0]
}

// func findRotation(mat [][]int, target [][]int) bool {
//     if len(mat) != len(target) || len(mat[0]) != len(target[0]) {
//         return false
//     }
//     i, j, n := 0, len(mat) - 1, len(mat[0])
//     for i <= j {
//         for k := 0; k < n; k++ {
//             if mat[i][k] != target[j][k] {
//                 return false
//             }
//         }
//         i++
//         j--
//     }
//     return true
// }

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/20/grid3.png" />
    // Input: mat = [[0,1],[1,0]], target = [[1,0],[0,1]]
    // Output: true
    // Explanation: We can rotate mat 90 degrees clockwise to make mat equal target.
    fmt.Println(findRotation([][]int{{0,1},{1,0}},[][]int{{1,0},{0,1}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/05/20/grid4.png" />
    // Input: mat = [[0,1],[1,1]], target = [[1,0],[0,1]]
    // Output: false
    // Explanation: It is impossible to make mat equal to target by rotating mat.
    fmt.Println(findRotation([][]int{{0,1},{1,1}},[][]int{{1,0},{0,1}})) // false
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/05/26/grid4.png" />
    // Input: mat = [[0,0,0],[0,1,0],[1,1,1]], target = [[1,1,1],[0,1,0],[0,0,0]]
    // Output: true
    // Explanation: We can rotate mat 90 degrees clockwise two times to make mat equal target.
    fmt.Println(findRotation([][]int{{0,0,0},{0,1,0},{1,1,1}},[][]int{{1,1,1},{0,1,0},{0,0,0}})) // true

    fmt.Println(findRotation1([][]int{{0,1},{1,0}},[][]int{{1,0},{0,1}})) // true
    fmt.Println(findRotation1([][]int{{0,1},{1,1}},[][]int{{1,0},{0,1}})) // false
    fmt.Println(findRotation1([][]int{{0,0,0},{0,1,0},{1,1,1}},[][]int{{1,1,1},{0,1,0},{0,0,0}})) // true
}