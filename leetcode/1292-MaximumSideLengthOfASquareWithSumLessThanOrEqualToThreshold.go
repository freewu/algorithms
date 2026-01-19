package main

// 1292. Maximum Side Length of a Square with Sum Less than or Equal to Threshold
// Given a m x n matrix mat and an integer threshold, 
// return the maximum side-length of a square with a sum less than or equal to threshold or return 0 if there is no such square.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/12/05/e1.png" />
// Input: mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4
// Output: 2
// Explanation: The maximum side length of square with sum less than 4 is 2 as shown.

// Example 2:
// Input: mat = [[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]], threshold = 1
// Output: 0

// Constraints:
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 300
//     0 <= mat[i][j] <= 10^4
//     0 <= threshold <= 10^5

import "fmt"

// prefixSum + binarySearch
func maxSideLength(mat [][]int, threshold int) int {
    m, n := len(mat), len(mat[0])
    sum := make([][]int, m + 1, m + 1)
    for i := range sum {
        sum[i] = make([]int, n + 1, n + 1)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            sum[i+1][j+1] = sum[i][j+1] + sum[i+1][j] - sum[i][j] + mat[i][j]
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    inThreshold := func(threshold int, length int) bool {
        for i := 1; i < len(sum); i++ {
            for j := 1; j < len(sum[0]); j++ {
                if i < length || j < length {
                    continue
                }
                if sum[i][j]+ sum[i-length][j-length]- sum[i-length][j]- sum[i][j-length] <= threshold {
                    return true
                }
            }
        }
        return false
    }
    l, r := 0, min(m, n)
    for l+1 < r {
        mid := (l + r) / 2
        if !inThreshold(threshold, mid) {
            r = mid
        } else {
            l = mid
        }
    }
    if inThreshold(threshold, r) {
        return r
    } else {
        return l
    }
}

func maxSideLength1(mat [][]int, threshold int) int {
    m, n := len(mat), len(mat[0])
    prefix := make([][]int, m+1)
    for i := range prefix {
        prefix[i] = make([]int, n+1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + mat[i-1][j-1]
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    getRect := func(x1, y1, x2, y2 int) int {
        return prefix[x2][y2] - prefix[x1-1][y2] - prefix[x2][y1-1] + prefix[x1-1][y1-1]
    }
    res, r := 0, min(m, n)
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            for c := res + 1; c <= r; c++ {
                if i+c-1 <= m && j+c-1 <= n && getRect(i, j, i+c-1, j+c-1) <= threshold {
                    res = c
                } else {
                    break
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/12/05/e1.png" />
    // Input: mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4
    // Output: 2
    // Explanation: The maximum side length of square with sum less than 4 is 2 as shown.
    mat1 := [][]int{
        {1,1,3,2,4,3,2},
        {1,1,3,2,4,3,2},
        {1,1,3,2,4,3,2},
    }
    fmt.Println(maxSideLength(mat1, 4)) // 2
    // Example 2:
    // Input: mat = [[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]], threshold = 1
    // Output: 0
    mat2 := [][]int{
        {2,2,2,2,2},
        {2,2,2,2,2},
        {2,2,2,2,2},
        {2,2,2,2,2},
        {2,2,2,2,2},
    }
    fmt.Println(maxSideLength(mat2, 1)) // 0

    fmt.Println(maxSideLength1(mat1, 4)) // 2
    fmt.Println(maxSideLength1(mat2, 1)) // 0
}