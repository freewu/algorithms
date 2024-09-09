package main

// 1072. Flip Columns For Maximum Number of Equal Rows
// You are given an m x n binary matrix matrix.

// You can choose any number of columns in the matrix and flip every cell in that column (i.e., Change the value of the cell from 0 to 1 or vice versa).

// Return the maximum number of rows that have all values equal after some number of flips.

// Example 1:
// Input: matrix = [[0,1],[1,1]]
// Output: 1
// Explanation: After flipping no values, 1 row has all values equal.

// Example 2:
// Input: matrix = [[0,1],[1,0]]
// Output: 2
// Explanation: After flipping values in the first column, both rows have equal values.

// Example 3:
// Input: matrix = [[0,0,0],[0,0,1],[1,1,0]]
// Output: 2
// Explanation: After flipping values in the first two columns, the last two rows have equal values.

// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 300
//     matrix[i][j] is either 0 or 1.

import "fmt"

func maxEqualRowsAfterFlips(matrix [][]int) int {
    original, flip, memo := make([]byte, len(matrix[0])), make([]byte, len(matrix[0])), map[string]int{}
    for i := range matrix {
        for j := range matrix[i] {
            original[j] = byte(matrix[i][j] + '0')
            flip[j] = byte(matrix[i][j] ^ 1 + '0')
        }
        memo[string(original)]++
        memo[string(flip)]++
    }
    res := 0
    for _, v := range memo {
        if v > res {
            res = v
        }
    }
    return res
}

func maxEqualRowsAfterFlips1(matrix [][]int) int {
    res, mp := 0, make(map[string]int)
    // 那为啥不直接使用取反运算符呢：~ 因为取反会把所有的位都取反，1取反之后不是0，而是 1111111111……0
    gen := func(arr []int) string {
        if arr[0] == 0 { // 如果是 0 开头就统一取反
            for i := 0; i < len(arr); i++ {
                arr[i] = arr[i] ^ 1
            }
        }
        res := make([]byte, len(arr))
        for i, ch := range arr {
            res[i] = byte(ch + 'a')
        }
        return string(res)
    }
    for _, item := range matrix {
        mp[gen(item)]++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range mp {
        res = max(res, v)
    }
    return res
}

func main() {
    // Example 1:
    // Input: matrix = [[0,1],[1,1]]
    // Output: 1
    // Explanation: After flipping no values, 1 row has all values equal.
    fmt.Println(maxEqualRowsAfterFlips([][]int{{0,1},{1,1}})) // 1
    // Example 2:
    // Input: matrix = [[0,1],[1,0]]
    // Output: 2
    // Explanation: After flipping values in the first column, both rows have equal values.
    fmt.Println(maxEqualRowsAfterFlips([][]int{{0,1},{1,0}})) // 2
    // Example 3:
    // Input: matrix = [[0,0,0],[0,0,1],[1,1,0]]
    // Output: 2
    // Explanation: After flipping values in the first two columns, the last two rows have equal values.
    fmt.Println(maxEqualRowsAfterFlips([][]int{{0,0,0},{0,0,1},{1,1,0}})) // 2

    fmt.Println(maxEqualRowsAfterFlips1([][]int{{0,1},{1,1}})) // 1
    fmt.Println(maxEqualRowsAfterFlips1([][]int{{0,1},{1,0}})) // 2
    fmt.Println(maxEqualRowsAfterFlips1([][]int{{0,0,0},{0,0,1},{1,1,0}})) // 2
}