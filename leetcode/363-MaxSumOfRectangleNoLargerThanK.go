package main

// 363. Max Sum of Rectangle No Larger Than K
// Given an m x n matrix matrix and an integer k, 
// return the max sum of a rectangle in the matrix such that its sum is no larger than k.
// It is guaranteed that there will be a rectangle with a sum no larger than k.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/18/sum-grid.jpg" />
// Input: matrix = [[1,0,1],[0,-2,3]], k = 2
// Output: 2
// Explanation: Because the sum of the blue rectangle [[0, 1], [-2, 3]] is 2, and 2 is the max number no larger than k (k = 2).

// Example 2:
// Input: matrix = [[2,2,-1]], k = 3
// Output: 3
 
// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 100
//     -100 <= matrix[i][j] <= 100
//     -10^5 <= k <= 10^5

// Follow up: What if the number of rows is much larger than the number of columns?

import "fmt"

func maxSumSubmatrix(matrix [][]int, k int) int {
    res, m, n := -1 << 32 - 1, len(matrix), len(matrix[0])
    sum := make([][]int, m + 1)
    for i := range sum {
        sum[i] = make([]int, n + 1)
    }
    for i := range matrix {
        for j := range matrix[i] {
            sum[i+1][j+1] = matrix[i][j] + sum[i+1][j] + sum[i][j+1] - sum[i][j]
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            for p := i; p <= m; p++ {
                for q := j; q <= n; q++ {
                    now := sum[p][q] - sum[i-1][q] - sum[p][j-1] + sum[i-1][j-1]
                    if now <= k {
                        res = max(res, now)
                    } 
                }
            }
        }
    }
    return res
}

func maxSumSubmatrix1(matrix [][]int, k int) int {
    res, inf, col, row := -1 << 32 - 1, -1 << 32 - 1, len(matrix[0]), len(matrix)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    getMaxNumBelowK := func (prefixCount []int, k int) int {
        mx, count := inf, 0 // O(n) 用了贪心的思路，找最大的子串和
        for _, v := range prefixCount {
            if count > 0 {
                count += v
            } else {
                count = v
            }
            mx = max(mx, count)
        }
        if mx <= k {
            return mx
        }
        // O(n^2)
        length := len(prefixCount)
        mx = inf
        for i := 0; i < length; i ++ {
            count := 0
            for j := i; j < length; j ++ {
                count += prefixCount[j]
                if count <= k {
                    mx = max(mx, count)
                }
                if mx == k {
                    return k
                }
            }
        }
        return mx
    }
    for left := 0; left < col; left ++ {
        prefixCount := make([]int, row)
        for right := left; right < col; right ++ {
            for i := 0; i < row; i ++ {
                prefixCount[i] += matrix[i][right]
            }
            res = max(res, getMaxNumBelowK(prefixCount, k))
            if res == k {
                return k
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/18/sum-grid.jpg" />
    // Input: matrix = [[1,0,1],[0,-2,3]], k = 2
    // Output: 2
    // Explanation: Because the sum of the blue rectangle [[0, 1], [-2, 3]] is 2, and 2 is the max number no larger than k (k = 2).
    fmt.Println(maxSumSubmatrix([][]int{{1,0,1},{0,-2,3}}, 2)) // 2
    // Example 2:
    // Input: matrix = [[2,2,-1]], k = 3
    // Output: 3
    fmt.Println(maxSumSubmatrix([][]int{{2,2,-1}}, 3)) // 3

    fmt.Println(maxSumSubmatrix1([][]int{{1,0,1},{0,-2,3}}, 2)) // 2
    fmt.Println(maxSumSubmatrix1([][]int{{2,2,-1}}, 3)) // 3
}