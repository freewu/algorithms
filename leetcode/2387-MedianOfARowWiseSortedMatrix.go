package main

// 2387. Median of a Row Wise Sorted Matrix
// Given an m x n matrix grid containing an odd number of integers where each row is sorted in non-decreasing order, 
// return the median of the matrix.

// You must solve the problem in less than O(m * n) time complexity.

// Example 1:
// Input: grid = [[1,1,2],[2,3,3],[1,3,4]]
// Output: 2
// Explanation: The elements of the matrix in sorted order are 1,1,1,2,2,3,3,3,4. The median is 2.

// Example 2:
// Input: grid = [[1,1,3,3,4]]
// Output: 3
// Explanation: The elements of the matrix in sorted order are 1,1,3,3,4. The median is 3.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 500
//     m and n are both odd.
//     1 <= grid[i][j] <= 10^6
//     grid[i] is sorted in non-decreasing order.

import "fmt"

func matrixMedian(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    res, target := -1, (m * n / 2) + 1
    left, right := 1, 1_000_000 // 外二分
    for left <= right {
        count, mid := 0, left + (right - left) / 2
        exists := false  // 判断是否存在，不存在的元素不能用的
        for _, row := range grid {
            c, l, r := -1, 0, n - 1 // 内二分
            for l <= r {
                mm := l + (r - l) / 2
                if row[mm] <= mid {
                    c, l = mm, mm + 1
                    if row[mm] == mid { // 找到了
                        exists = true
                    }
                } else {
                    r = mm - 1
                }
            }
            count += (c + 1)
        }
        if count < target {
            left = mid + 1
        } else {
            if exists { res = mid }
            right = mid - 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[1,1,2],[2,3,3],[1,3,4]]
    // Output: 2
    // Explanation: The elements of the matrix in sorted order are 1,1,1,2,2,3,3,3,4. The median is 2.
    fmt.Println(matrixMedian([][]int{{1,1,2},{2,3,3},{1,3,4}})) // 2
    // Example 2:
    // Input: grid = [[1,1,3,3,4]]
    // Output: 3
    // Explanation: The elements of the matrix in sorted order are 1,1,3,3,4. The median is 3.
    fmt.Println(matrixMedian([][]int{{1,1,3,3,4}})) // 3
}