package main

// 面试题 01.07. Rotate Matrix LCCI

// Given an image represented by an N x N matrix,
// where each pixel in the image is 4 bytes, write a method to rotate the image by 90 degrees. 
// Can you do this in place?

// Example 1:
// Given matrix = 
// [
//   [1,2,3],
//   [4,5,6],
//   [7,8,9]
// ],
// Rotate the matrix in place. It becomes:
// [
//   [7,4,1],
//   [8,5,2],
//   [9,6,3]
// ]

// Example 2:
// Given matrix =
// [
//   [ 5, 1, 9,11],
//   [ 2, 4, 8,10],
//   [13, 3, 6, 7],
//   [15,14,12,16]
// ], 
// Rotate the matrix in place. It becomes:
// [
//   [15,13, 2, 5],
//   [14, 3, 4, 1],
//   [12, 6, 8, 9],
//   [16, 7,10,11]
// ]

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