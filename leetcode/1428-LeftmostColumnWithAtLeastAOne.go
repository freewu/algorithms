package main

// 1428. Leftmost Column with at Least a One
// A row-sorted binary matrix means that all elements are 0 or 1 and each row of the matrix is sorted in non-decreasing order.

// Given a row-sorted binary matrix binaryMatrix, 
// return the index (0-indexed) of the leftmost column with a 1 in it. 
// If such an index does not exist, return -1.

// You can't access the Binary Matrix directly. You may only access the matrix using a BinaryMatrix interface:
//     BinaryMatrix.get(row, col) 
//         returns the element of the matrix at index (row, col) (0-indexed).
//     BinaryMatrix.dimensions() 
//         returns the dimensions of the matrix as a list of 2 elements [rows, cols], which means the matrix is rows x cols.

// Submissions making more than 1000 calls to BinaryMatrix.get will be judged Wrong Answer. 
// Also, any solutions that attempt to circumvent the judge will result in disqualification.

// For custom testing purposes, the input will be the entire binary matrix mat. 
// You will not have access to the binary matrix directly.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/10/25/untitled-diagram-5.jpg" />
// Input: mat = [[0,0],[1,1]]
// Output: 0

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/10/25/untitled-diagram-4.jpg" />
// Input: mat = [[0,0],[0,1]]
// Output: 1

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/10/25/untitled-diagram-3.jpg" />
// Input: mat = [[0,0],[0,0]]
// Output: -1

// Constraints:
//     rows == mat.length
//     cols == mat[i].length
//     1 <= rows, cols <= 100
//     mat[i][j] is either 0 or 1.
//     mat[i] is sorted in non-decreasing order.

import "fmt"

/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get func(int, int) int
 *     Dimensions func() []int
 * }
 */
func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
    data := binaryMatrix.Dimensions()
    m,n := data[0], data[1]
    res := n
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < m; i++ {
        if binaryMatrix.Get(i, n - 1) == 0 {
            continue
        }
        l, r := 0, n
        for l < r {
            mid := (r - l) >> 1 + l
            if binaryMatrix.Get(i, mid) == 0 {
                l = mid + 1
            } else {
                r = mid
            }
        }
        if l < n {
            if l == 0 {
                return 0
            } else {
                res = min(res, l)
            }
        }
    }
    if res == n {
        return -1
    }
    return res
}

func leftMostColumnWithOne1(binaryMatrix BinaryMatrix) int {
    dims := binaryMatrix.Dimensions()
    res := dims[1];
    for i := 0 ; i < dims[0] ; i++ {
        left, right := 0, dims[1] - 1
        for left < right {
            mid := left + (right - left) / 2
            v := binaryMatrix.Get(i , mid)
            if v == 1 {
                right = mid
            } else {
                left = mid + 1
            }
        }
        v := binaryMatrix.Get(i , left)
        if v == 1 && left < res {
            res = left
        }
    }
    if res == dims[1] {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/10/25/untitled-diagram-5.jpg" />
    // Input: mat = [[0,0],[1,1]]
    // Output: 0

    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/10/25/untitled-diagram-4.jpg" />
    // Input: mat = [[0,0],[0,1]]
    // Output: 1

    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/10/25/untitled-diagram-3.jpg" />
    // Input: mat = [[0,0],[0,0]]
    // Output: -1
    fmt.Println()
}