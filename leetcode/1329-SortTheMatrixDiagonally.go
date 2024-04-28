package main

// 1329. Sort the Matrix Diagonally
// A matrix diagonal is a diagonal line of cells starting from some cell in either the topmost row or leftmost column 
// and going in the bottom-right direction until reaching the matrix's end. 
// For example, the matrix diagonal starting from mat[2][0], 
// where mat is a 6 x 3 matrix, includes cells mat[2][0], mat[3][1], and mat[4][2].

// Given an m x n matrix mat of integers, sort each matrix diagonal in ascending order and return the resulting matrix.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/01/21/1482_example_1_2.png" />
// Input: mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
// Output: [[1,1,1,1],[1,2,2,2],[1,2,3,3]]

// Example 2:
// Input: mat = [[11,25,66,1,69,7],[23,55,17,45,15,52],[75,31,36,44,58,8],[22,27,33,25,68,4],[84,28,14,11,5,50]]
// Output: [[5,17,4,1,52,7],[11,11,25,45,8,69],[14,23,25,44,58,15],[22,27,31,36,50,66],[84,28,75,33,55,68]]
 
// Constraints:
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 100
//     1 <= mat[i][j] <= 100

import "fmt"
import "sort"
import "slices"

func printMatrix(matrix [][]int) {
    for _,v := range matrix {
        fmt.Println(v)
    }
    fmt.Println()
}

func diagonalSort(mat [][]int) [][]int {
    n, m := len(mat), len(mat[0])
    modify := func(i, j int) {
        arr := make([]int, 0)
        for r, c := i, j; r < n && c < m; r, c = r+1, c+1 {
            arr = append(arr, mat[r][c])
        }
        sort.Ints(arr)
        for r, c := i, j; r < n && c < m; r, c = r+1, c+1 {
            mat[r][c] = arr[0]
            arr = arr[1:]
        }
    }
    for i := 0; i < m; i++ {
        modify(0, i)
    }
    for i := 1; i < n; i++ {
        modify(i, 0)
    }
    return mat
}

func diagonalSort1(mat [][]int) [][]int {
    t, n, m := make(map[int][]int), len(mat), len(mat[0])
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            t[i-j] = append(t[i-j], mat[i][j])
        }
    }
    for _, v := range t {
        sort.Ints(v)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            mat[i][j] = t[i-j][0]
            t[i-j] = t[i-j][1:]
        }
    }
    return mat
}

func diagonalSort2(mat [][]int) [][]int {
    m, n := len(mat), len(mat[0])
    arr := make([]int, min(m, n))
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for k := 1 - n; k < m; k++ { // k = i - j
        a := arr[:0]
        mn := max(k, 0)
        mx := min(k+n, m)
        for i := mn; i < mx; i++ {
            a = append(a, mat[i][i-k])
        }
        slices.Sort(a)
        for i := mn; i < mx; i++ {
            mat[i][i-k] = a[i-mn]
        }
    }
    return mat
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/01/21/1482_example_1_2.png" />
    // Input: mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
    // Output: [[1,1,1,1],[1,2,2,2],[1,2,3,3]]
    matrix1 := [][]int{{3,3,1,1},{2,2,1,2},{1,1,1,2}}
    printMatrix(matrix1)
    printMatrix(diagonalSort(matrix1))

    // Example 2:
    // Input: mat = [[11,25,66,1,69,7],[23,55,17,45,15,52],[75,31,36,44,58,8],[22,27,33,25,68,4],[84,28,14,11,5,50]]
    // Output: [[5,17,4,1,52,7],[11,11,25,45,8,69],[14,23,25,44,58,15],[22,27,31,36,50,66],[84,28,75,33,55,68]]
    matrix2 := [][]int{{11,25,66,1,69,7},{23,55,17,45,15,52},{75,31,36,44,58,8},{22,27,33,25,68,4},{84,28,14,11,5,50}}
    printMatrix(matrix2)
    printMatrix(diagonalSort(matrix2))

    matrix11 := [][]int{{3,3,1,1},{2,2,1,2},{1,1,1,2}}
    printMatrix(matrix11)
    printMatrix(diagonalSort1(matrix11))
    matrix12 := [][]int{{11,25,66,1,69,7},{23,55,17,45,15,52},{75,31,36,44,58,8},{22,27,33,25,68,4},{84,28,14,11,5,50}}
    printMatrix(matrix12)
    printMatrix(diagonalSort1(matrix12))

    matrix21 := [][]int{{3,3,1,1},{2,2,1,2},{1,1,1,2}}
    printMatrix(matrix21)
    printMatrix(diagonalSort1(matrix21))
    matrix22 := [][]int{{11,25,66,1,69,7},{23,55,17,45,15,52},{75,31,36,44,58,8},{22,27,33,25,68,4},{84,28,14,11,5,50}}
    printMatrix(matrix22)
    printMatrix(diagonalSort1(matrix22))
}