package main

// 2022. Convert 1D Array Into 2D Array
// You are given a 0-indexed 1-dimensional (1D) integer array original, and two integers, m and n. 
// You are tasked with creating a 2-dimensional (2D) array with m rows and n columns using all the elements from original.

// The elements from indices 0 to n - 1 (inclusive) of original should form the first row of the constructed 2D array, 
// the elements from indices n to 2 * n - 1 (inclusive) should form the second row of the constructed 2D array, and so on.

// Return an m x n 2D array constructed according to the above procedure, or an empty 2D array if it is impossible.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/08/26/image-20210826114243-1.png" />
// Input: original = [1,2,3,4], m = 2, n = 2
// Output: [[1,2],[3,4]]
// Explanation: The constructed 2D array should contain 2 rows and 2 columns.
// The first group of n=2 elements in original, [1,2], becomes the first row in the constructed 2D array.
// The second group of n=2 elements in original, [3,4], becomes the second row in the constructed 2D array.

// Example 2:
// Input: original = [1,2,3], m = 1, n = 3
// Output: [[1,2,3]]
// Explanation: The constructed 2D array should contain 1 row and 3 columns.
// Put all three elements in original into the first row of the constructed 2D array.

// Example 3:
// Input: original = [1,2], m = 1, n = 1
// Output: []
// Explanation: There are 2 elements in original.
// It is impossible to fit 2 elements in a 1x1 2D array, so return an empty 2D array.

// Constraints:
//     1 <= original.length <= 5 * 10^4
//     1 <= original[i] <= 10^5
//     1 <= m, n <= 4 * 10^4

import "fmt"

// func construct2DArray(original []int, m int, n int) [][]int {
//     i, j, l, res := 0, 0, len(original), [][]int{}
//     if l != n * m {
//         return res
//     }
//     for k := 0; k < l, k++ {
//         res[]
//     }
//     return res
// }

func construct2DArray(original []int, m int, n int) [][]int {
    res, rows := [][]int{}, []int{}
    if len(original) != m * n {
        return res
    }
    for k, v := range original {
        rows = append(rows, v)
        if (k + 1) % n == 0 { // 需要换一行了
            res = append(res, rows)
            rows = []int{}
        }    
    }
    return res
}

func construct2DArray1(original []int, m int, n int) [][]int {
    if len(original) != m * n {
        return [][]int{}
    }
    pos, res := 0, make([][]int, m)
    for i := range res {
        res[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if pos >= len(original) {
                res[i][j] = 0
            } else {
                res[i][j] = original[pos]
            }
            pos++
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/08/26/image-20210826114243-1.png" />
    // Input: original = [1,2,3,4], m = 2, n = 2
    // Output: [[1,2],[3,4]]
    // Explanation: The constructed 2D array should contain 2 rows and 2 columns.
    // The first group of n=2 elements in original, [1,2], becomes the first row in the constructed 2D array.
    // The second group of n=2 elements in original, [3,4], becomes the second row in the constructed 2D array.
    fmt.Println(construct2DArray([]int{1,2,3,4},2,2)) // [[1,2],[3,4]]
    // Example 2:
    // Input: original = [1,2,3], m = 1, n = 3
    // Output: [[1,2,3]]
    // Explanation: The constructed 2D array should contain 1 row and 3 columns.
    // Put all three elements in original into the first row of the constructed 2D array.
    fmt.Println(construct2DArray([]int{1,2,3},1,3)) // [[1,2,3]]
    // Example 3:
    // Input: original = [1,2], m = 1, n = 1
    // Output: []
    // Explanation: There are 2 elements in original.
    // It is impossible to fit 2 elements in a 1x1 2D array, so return an empty 2D array.
    fmt.Println(construct2DArray([]int{1,2},1,1)) // []

    fmt.Println(construct2DArray1([]int{1,2,3,4},2,2)) // [[1,2],[3,4]]
    fmt.Println(construct2DArray1([]int{1,2,3},1,3)) // [[1,2,3]]
    fmt.Println(construct2DArray1([]int{1,2},1,1)) // []
}