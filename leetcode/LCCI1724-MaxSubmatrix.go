package main

// 面试题 17.24. Max Submatrix LCCI
// Given an NxM matrix of positive and negative integers, write code to find the submatrix with the largest possible sum.

// Return an array [r1, c1, r2, c2], where r1, c1 are the row number and the column number of the submatrix's upper left corner respectively, and r2, c2 are the row number of and the column number of lower right corner. 
// If there are more than one answers, return any one of them.

// Note: This problem is slightly different from the original one in the book.

// Example:
// Input:
// [
//    [-1,0],
//    [0,-1]
// ]
// Output: [0,1,0,1]

// Note:
//     1 <= matrix.length, matrix[0].length <= 200

import "fmt"

func getMaxMatrix(matrix [][]int) []int {
    n, m := len(matrix), len(matrix[0])
    // 预处理：求出每一列的前缀和，从而O(1)求出任何一列的任何一段的元素和
    prefix := [][]int{}
    for j := 0; j < m; j++ {
        col := make([]int, n + 1)
        col[0] = 0 
        for i := 1; i <= n; i++ {
            col[i] = col[i - 1] + matrix[i - 1][j]
        }
        prefix = append(prefix, col)
    }
    // 二维压缩到一维，需要考虑所有可能的情况
    res := [5]int{0,0,0,0,matrix[0][0]} // 最后一个数是子矩阵元素和
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ { // 对于行号为[i,j]的矩阵压缩到一维，& 一维最大子序列和
            dp, c := -1, 0 //c记录起点列号
            for k := 0; k < m; k++ {
                temp := prefix[k][j+1] - prefix[k][i]
                if dp < 0 {
                    dp = temp // 另起炉灶
                    c = k
                } else {
                    dp += temp
                }
                if dp > res[4] {
                    res = [5]int{ i,c,j,k,dp }
                }
            }
        }
    }
    return res[:4]
}

func getMaxMatrix1(matrix [][]int) []int {
    m, n , mx := len(matrix), len(matrix[0]), -1 << 31
    res := []int{0, 0, 0, 0}
    for i := 0; i < m; i++ {
        arr := make([]int, n)
        for j := i; j < m; j++ {
            cur, start := 0, -1
            for k := 0; k < n; k++ {
                arr[k] += matrix[j][k]
                cur += arr[k]
                if cur > mx {
                    mx = cur
                    res = []int{i, start + 1, j, k}
                }
                if cur < 0 {
                    cur, start = 0, k
                }
            }
        }
    }
    return res
}

func main() {
    // Example:
    // Input:
    // [
    //    [-1,0],
    //    [0,-1]
    // ]
    // Output: [0,1,0,1]
    fmt.Println(getMaxMatrix([][]int{{-1,0}, {0,-1}})) // [0,1,0,1]

    fmt.Println(getMaxMatrix([][]int{{1,0}, {0,1}})) // [0 0 1 1]
    fmt.Println(getMaxMatrix([][]int{{1,1}, {1,1}})) // [0 0 1 1]
    fmt.Println(getMaxMatrix([][]int{{0,0}, {0,0}})) // [0 0 0 0]
    fmt.Println(getMaxMatrix([][]int{{-1,-1}, {-1,-1}})) // [0 0 0 0]

    fmt.Println(getMaxMatrix1([][]int{{-1,0}, {0,-1}})) // [0,1,0,1]
    fmt.Println(getMaxMatrix1([][]int{{1,0}, {0,1}})) // [0 0 1 1]
    fmt.Println(getMaxMatrix1([][]int{{1,1}, {1,1}})) // [0 0 1 1]
    fmt.Println(getMaxMatrix1([][]int{{0,0}, {0,0}})) // [0 0 0 0]
    fmt.Println(getMaxMatrix1([][]int{{-1,-1}, {-1,-1}})) // [0 0 0 0]
}