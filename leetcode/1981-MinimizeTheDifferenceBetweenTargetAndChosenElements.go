package main

// 1981. Minimize the Difference Between Target and Chosen Elements
// You are given an m x n integer matrix mat and an integer target.
// Choose one integer from each row in the matrix such that the absolute difference between target and the sum of the chosen elements is minimized.
// Return the minimum absolute difference.
// The absolute difference between two numbers a and b is the absolute value of a - b.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/08/03/matrix1.png" />
// Input: mat = [[1,2,3],[4,5,6],[7,8,9]], target = 13
// Output: 0
// Explanation: One possible choice is to:
// - Choose 1 from the first row.
// - Choose 5 from the second row.
// - Choose 7 from the third row.
// The sum of the chosen elements is 13, which equals the target, so the absolute difference is 0.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/08/03/matrix1-1.png" />
// Input: mat = [[1],[2],[3]], target = 100
// Output: 94
// Explanation: The best possible choice is to:
// - Choose 1 from the first row.
// - Choose 2 from the second row.
// - Choose 3 from the third row.
// The sum of the chosen elements is 6, and the absolute difference is 94.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/08/03/matrix1-3.png" />
// Input: mat = [[1,2,9,8,7]], target = 6
// Output: 1
// Explanation: The best choice is to choose 7 from the first row.
// The absolute difference is 1.

// Constraints:
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 70
//     1 <= mat[i][j] <= 70
//     1 <= target <= 800

import "fmt"

// 超出时间限制 80 / 81
func minimizeTheDifference(mat [][]int, target int) int {
    res, totalMin := 1 << 31, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := range mat {
        relMin := 1 << 31
        for j := range mat[i] {
            relMin = min(relMin, mat[i][j])
        }
        totalMin += relMin
    }
    if totalMin >= target {
        return totalMin - target
    }
    candidate := make(map[int]bool)
    candidate[0] = true
    for i := range mat {
        tmp := make(map[int]bool)
        for j := range mat[i] {
            for c := range candidate {
                if c+mat[i][j] < 2 * target - totalMin {
                    tmp[c+mat[i][j]] = true
                }
            }
        }
        candidate = tmp
    }
    for c := range candidate {
        res = min(res, abs(target - c))
    }
    return res
}

func minimizeTheDifference1(mat [][]int, target int) int {
    dp := make([]bool, min(len(mat)*70, target*2)+1) // 需要枚举的重量不会超过 target*2
    dp[0] = true
    minSum, maxSum := 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, row := range mat {
        mi, mx := row[0], row[0]
        for _, v := range row[1:] {
            if v > mx {
                mx = v
            } else if v < mi {
                mi = v
            }
        }
        minSum += mi // 求 minSum 是为了防止 target 过小导致 dp 没有记录到
        maxSum = min(maxSum+mx, target*2) // 前 i 组的最大重量，优化枚举时 j 的初始值
        for j := maxSum; j >= 0; j-- {
            dp[j] = false
            for _, v := range row {
                if v <= j && dp[j-v] {
                    dp[j] = true
                    break
                }
            }
        }
    }
    res := abs(minSum - target)
    for i, ok := range dp {
        if ok {
            res = min(res, abs(i-target))
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/08/03/matrix1.png" />
    // Input: mat = [[1,2,3],[4,5,6],[7,8,9]], target = 13
    // Output: 0
    // Explanation: One possible choice is to:
    // - Choose 1 from the first row.
    // - Choose 5 from the second row.
    // - Choose 7 from the third row.
    // The sum of the chosen elements is 13, which equals the target, so the absolute difference is 0.
    fmt.Println(minimizeTheDifference([][]int{{1,2,3},{4,5,6},{7,8,9}}, 13)) // 0
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/08/03/matrix1-1.png" />
    // Input: mat = [[1],[2],[3]], target = 100
    // Output: 94
    // Explanation: The best possible choice is to:
    // - Choose 1 from the first row.
    // - Choose 2 from the second row.
    // - Choose 3 from the third row.
    // The sum of the chosen elements is 6, and the absolute difference is 94.
    fmt.Println(minimizeTheDifference([][]int{{1},{2},{3}}, 100)) // 94
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/08/03/matrix1-3.png" />
    // Input: mat = [[1,2,9,8,7]], target = 6
    // Output: 1
    // Explanation: The best choice is to choose 7 from the first row.
    // The absolute difference is 1.
    fmt.Println(minimizeTheDifference([][]int{{1,2,9,8,7}}, 6)) // 1

    fmt.Println(minimizeTheDifference1([][]int{{1,2,3},{4,5,6},{7,8,9}}, 13)) // 0
    fmt.Println(minimizeTheDifference1([][]int{{1},{2},{3}}, 100)) // 94
    fmt.Println(minimizeTheDifference1([][]int{{1,2,9,8,7}}, 6)) // 1
}