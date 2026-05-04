package main

// 48. Rotate Image
// You are given an n x n 2D matrix representing an image.
// Rotate the image by 90 degrees (clockwise).
// Note:
//     You have to rotate the image  in-place, which means you have to modify the input 2D matrix directly. 
//     DO NOT allocate another 2D matrix and do the rotation.

// Example 1:
// Given input matrix =
// [
//     [1,2,3],
//     [4,5,6],
//     [7,8,9]
// ],
// rotate the input matrix in-place such that it becomes:
// [
//     [7,4,1],
//     [8,5,2],
//     [9,6,3]
// ]

// Example 2:
// Given input matrix =
// [
//     [ 5, 1, 9,11],
//     [ 2, 4, 8,10],
//     [13, 3, 6, 7],
//     [15,14,12,16]
// ],
// rotate the input matrix in-place such that it becomes:
// [
//     [15,13, 2, 5],
//     [14, 3, 4, 1],
//     [12, 6, 8, 9],
//     [16, 7,10,11]
// ]

// Constraints:
//     n == matrix.length == matrix[i].length
//     1 <= n <= 20
//     -1000 <= matrix[i][j] <= 1000

// 解题思路:

//      # clockwise rotate 顺时针旋转
//      first reverse up to down, then swap the symmetry
//      1 2 3     7 8 9     7 4 1
//      4 5 6  => 4 5 6  => 8 5 2
//      7 8 9     1 2 3     9 6 3

//      # anticlockwise rotate 逆时针旋转
//      first reverse left to right, then swap the symmetry
//      1 2 3     3 2 1     3 6 9
//      4 5 6  => 6 5 4  => 2 5 8
//      7 8 9     9 8 7     1 4 7

import "fmt"

// 解法一
func rotate(matrix [][]int) {
    n := len(matrix)
    // rotate by diagonal 对角线变换
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    // rotate by vertical centerline 竖直轴对称翻转
    for i := 0; i < n; i++ {
        for j := 0; j < n / 2; j++ {
            matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
        }
    }
}

// 解法二
func rotate1(matrix [][]int) {
    n := len(matrix)
    if n == 1 { return }
    // swap changes original slice's i,j position
    rotateSwap := func(nums [][]int, i, j int) { nums[i][j], nums[j][i] = nums[j][i], nums[i][j]; }
    // reverses a row of image, matrix[i]
    rotateReverse := func(nums []int) []int {
        left, right := 0, len(nums) - 1
        for left < right {
            nums[left], nums[right] = nums[right], nums[left]
            left++
            right--
        }
        return nums
    }
    /* rotate clock-wise = 1. transpose matrix => 2. reverse(matrix[i])
    1   2  3  4      1   5  9  13        13  9  5  1
    5   6  7  8  =>  2   6  10 14  =>    14  10 6  2
    9  10 11 12      3   7  11 15        15  11 7  3
    13 14 15 16      4   8  12 16        16  12 8  4
    */
    for i := 0; i < n; i++ {
        // transpose, i=rows, j=columns
        // j = i+1, coz diagonal elements didn't change in a square matrix
        for j := i + 1; j < n; j++ {
            rotateSwap(matrix, i, j)
        }
        // reverse each row of the image
        matrix[i] = rotateReverse(matrix[i])
    }
}

// 递归
func rotate2(matrix [][]int) {
    left, right, top, low := 0, len(matrix[0]) - 1, 0, len(matrix) - 1
    var swapRotate func(matrix [][]int, left int, right int, top int, low int)
    swapRotate = func(matrix [][]int, left int, right int, top int, low int) {
        if left > right || top > low { return }
        for i := 0; i < right-left; i++ {
            matrix[top][left+i], matrix[top+i][right] = matrix[top+i][right], matrix[top][left+i]
            matrix[top][left+i], matrix[low][right-i] = matrix[low][right-i], matrix[top][left+i]
            matrix[top][left+i], matrix[low-i][left] = matrix[low-i][left], matrix[top][left+i]
        }
        swapRotate(matrix, left+1, right-1, top+1, low-1)
    }
    swapRotate(matrix, left, right, top, low)
}

func main() {
    // Example 1:
    // Given input matrix =
    // [
    //     [1,2,3],
    //     [4,5,6],
    //     [7,8,9]
    // ],
    // rotate the input matrix in-place such that it becomes:
    // [
    //     [7,4,1],
    //     [8,5,2],
    //     [9,6,3]
    // ]
    fmt.Println("matrix01 rotate: ")
    matrix01 := [][]int{
        []int{1,2,3},
        []int{4,5,6},
        []int{7,8,9},
    }
    for _,m := range matrix01 {
        fmt.Printf("%v\n",m)
    }
    fmt.Println("matrix01 rotate after:")
    rotate(matrix01)
    for _,m := range matrix01 {
        fmt.Printf("%v\n",m)
    }
    // Example 2:
    // Given input matrix =
    // [
    //     [ 5, 1, 9,11],
    //     [ 2, 4, 8,10],
    //     [13, 3, 6, 7],
    //     [15,14,12,16]
    // ],
    // rotate the input matrix in-place such that it becomes:
    // [
    //     [15,13, 2, 5],
    //     [14, 3, 4, 1],
    //     [12, 6, 8, 9],
    //     [16, 7,10,11]
    // ]
    matrix02 := [][]int{
        []int{5, 1, 9,11},
        []int{2, 4, 8,10},
        []int{13, 3, 6,7},
        []int{15,14,12,16},
    }
    for _,m := range matrix02 {
        fmt.Printf("%v\n",m)
    }
    fmt.Println("matrix02 rotate after:")
    rotate(matrix02)
    for _,m := range matrix02 {
        fmt.Printf("%v\n",m)
    }

    fmt.Println("matrix11 rotate: ")
    matrix11 := [][]int{
        []int{1,2,3},
        []int{4,5,6},
        []int{7,8,9},
    }
    for _,m := range matrix11 {
        fmt.Printf("%v\n",m)
    }
    fmt.Println("matrix11 rotate after:")
    rotate1(matrix11)
    for _,m := range matrix11 {
        fmt.Printf("%v\n",m)
    }
    matrix12 := [][]int{
        []int{5, 1, 9,11},
        []int{2, 4, 8,10},
        []int{13, 3, 6,7},
        []int{15,14,12,16},
    }
    fmt.Println("matrix12 rotate: ")
    for _,m := range matrix12 {
        fmt.Printf("%v\n",m)
    }
    fmt.Println("matrix12 rotate after:")
    rotate1(matrix12)
    for _,m := range matrix12 {
        fmt.Printf("%v\n",m)
    }

    fmt.Println("matrix21 rotate: ")
    matrix21 := [][]int{
        []int{1,2,3},
        []int{4,5,6},
        []int{7,8,9},
    }
    for _,m := range matrix21 {
        fmt.Printf("%v\n",m)
    }
    fmt.Println("matrix21 rotate after:")
    rotate2(matrix21)
    for _,m := range matrix21 {
        fmt.Printf("%v\n",m)
    }
    matrix22 := [][]int{
        []int{5, 1, 9,11},
        []int{2, 4, 8,10},
        []int{13, 3, 6,7},
        []int{15,14,12,16},
    }
    fmt.Println("matrix22 rotate: ")
    for _,m := range matrix22 {
        fmt.Printf("%v\n",m)
    }
    fmt.Println("matrix22 rotate after:")
    rotate2(matrix22)
    for _,m := range matrix22 {
        fmt.Printf("%v\n",m)
    }
}