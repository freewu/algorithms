package main

// 3888. Minimum Operations to Make All Grid Elements Equal
// You are given a 2D integer array grid of size m × n, and an integer k.

// In one operation, you can:
//     1. Select any k x k submatrix of grid, and
//     2. Increment all elements inside that submatrix by 1.

// Return the minimum number of operations required to make all elements in the grid equal. If it is not possible, return -1.

// A submatrix (x1, y1, x2, y2) is a matrix that forms by choosing all cells matrix[x][y] where x1 <= x <= x2 and y1 <= y <= y2.

// Example 1:
// Input: grid = [[3,3,5],[3,3,5]], k = 2
// Output: 2
// Explanation:
// Choose the left 2 x 2 submatrix (covering the first two columns) and apply the operation twice.
// After 1 operation: [[4, 4, 5], [4, 4, 5]]
// After 2 operations: [[5, 5, 5], [5, 5, 5]]
// All elements become equal to 5. Thus, the minimum number of operations is 2.

// Example 2:
// Input: grid = [[1,2],[2,3]], k = 1
// Output: 4
// Explanation:
// Since k = 1, each operation increments a single cell grid[i][j] by 1. To make all elements equal, the final value must be 3.
// Increase grid[0][0] = 1 to 3, requiring 2 operations.
// Increase grid[0][1] = 2 to 3, requiring 1 operation.
// Increase grid[1][0] = 2 to 3, requiring 1 operation.
// Thus, the minimum number of operations is 2 + 1 + 1 + 0 = 4.

// Constraints:
//     1 <= m == grid.length <= 1000
//     1 <= n == grid[i].length <= 1000
//     -105 <= grid[i][j] <= 10^5
//     1 <= k <= min(m, n)

import "fmt"

func minOperations(grid [][]int, k int) int64 {
    m, n := len(grid), len(grid[0])
    if m == 0 { return 0 }
    // 初始化列和数组
    colSum := make([]int, n)
    // 初始化操作环二维切片
    opsRing := make([][]int, k)
    for i := range opsRing {
        opsRing[i] = make([]int, n)
    }
    hasCandidate, candidate, sumOps0, sumOps1, mn := false, 0, 0, 0, -1 << 63 // 等价于 float('-inf')
    for i := 0; i < m; i++ {
        // 移除超出窗口的行
        if i >= k {
            evictRow := i % k
            for j := 0; j < n; j++ {
                colSum[j] -= opsRing[evictRow][j]
                opsRing[evictRow][j] = 0
            }
        }
        windowSum := 0
        for j := 0; j < n; j++ {
            windowSum += colSum[j]
            req0 := -grid[i][j] - windowSum
            if i <= m-k && j <= n-k {
                opsRing[i%k][j] = req0
                colSum[j] += req0
                windowSum += req0
                // 判断是否是左上角操作点
                o1 := 0
                if i % k == 0 && j % k == 0 {
                    o1 = 1
                }
                sumOps0 += req0
                sumOps1 += o1
                if o1 == 1 {
                    if -req0 > mn {
                        mn = -req0
                    }
                } else {
                    if req0 < 0 {
                        return -1
                    }
                }
            } else {
                r0 := req0
                // 计算 r1
                r1 := 0
                if (i / k) * k > m - k || (j / k) * k > n - k {
                    r1 = 1
                }
                if r1 != 0 {
                    x := -r0
                    if !hasCandidate {
                        candidate, hasCandidate = x, true
                    } else if candidate != x {
                        return -1
                    }
                } else {
                    if r0 != 0 {
                        return -1
                    }
                }
            }
            // 滑动窗口减去超出的列
            if j >= k-1 {
                windowSum -= colSum[j-k+1]
            }
        }
    }
    // 最终校验
    if hasCandidate {
        if candidate < mn {
            return -1
        }
    } else {
        candidate = mn
    }
    return int64(sumOps0) + int64(candidate) * int64(sumOps1)
}

func main() {
    // Example 1:
    // Input: grid = [[3,3,5],[3,3,5]], k = 2
    // Output: 2
    // Explanation:
    // Choose the left 2 x 2 submatrix (covering the first two columns) and apply the operation twice.
    // After 1 operation: [[4, 4, 5], [4, 4, 5]]
    // After 2 operations: [[5, 5, 5], [5, 5, 5]]
    // All elements become equal to 5. Thus, the minimum number of operations is 2.
    fmt.Println(minOperations([][]int{{3,3,5},{3,3,5}}, 2)) // 2
    // Example 2:
    // Input: grid = [[1,2],[2,3]], k = 1
    // Output: 4
    // Explanation:
    // Since k = 1, each operation increments a single cell grid[i][j] by 1. To make all elements equal, the final value must be 3.
    // Increase grid[0][0] = 1 to 3, requiring 2 operations.
    // Increase grid[0][1] = 2 to 3, requiring 1 operation.
    // Increase grid[1][0] = 2 to 3, requiring 1 operation.
    // Thus, the minimum number of operations is 2 + 1 + 1 + 0 = 4.
    fmt.Println(minOperations([][]int{{1,2},{2,3}}, 1)) // 4

    fmt.Println(minOperations([][]int{{1,2,3,4,5,6,7,8,9},{1,2,3,4,5,6,7,8,9}}, 1)) // 72
    fmt.Println(minOperations([][]int{{1,2,3,4,5,6,7,8,9},{9,8,7,6,5,4,3,2,1}}, 1)) // 72
    fmt.Println(minOperations([][]int{{9,8,7,6,5,4,3,2,1},{1,2,3,4,5,6,7,8,9}}, 1)) // 72
    fmt.Println(minOperations([][]int{{9,8,7,6,5,4,3,2,1},{9,8,7,6,5,4,3,2,1}}, 1)) // 72
}
