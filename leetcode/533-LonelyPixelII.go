package main

// 533. Lonely Pixel II
// Given an m x n picture consisting of black 'B' and white 'W' pixels and an integer target, 
// return the number of black lonely pixels.

// A black lonely pixel is a character 'B' that located at a specific position (r, c) where:
//     Row r and column c both contain exactly target black pixels.
//     For all rows that have a black pixel at column c, they should be exactly the same as row r.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/pixel2-1-grid.jpg" />
// Input: picture = [["W","B","W","B","B","W"],["W","B","W","B","B","W"],["W","B","W","B","B","W"],["W","W","B","W","B","W"]], target = 3
// Output: 6
// Explanation: All the green 'B' are the black pixels we need (all 'B's at column 1 and 3).
// Take 'B' at row r = 0 and column c = 1 as an example:
//  - Rule 1, row r = 0 and column c = 1 both have exactly target = 3 black pixels. 
//  - Rule 2, the rows have black pixel at column c = 1 are row 0, row 1 and row 2. They are exactly the same as row r = 0.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/pixel2-2-grid.jpg" />
// Input: picture = [["W","W","B"],["W","W","B"],["W","W","B"]], target = 1
// Output: 0

// Constraints:
//     m == picture.length
//     n == picture[i].length
//     1 <= m, n <= 200
//     picture[i][j] is 'W' or 'B'.
//     1 <= target <= min(m, n)

import "fmt"

func findBlackPixel(picture [][]byte, target int) int {
    if 0 == len(picture) || 0 == len(picture[0]) {
        return 0
    }
    m, n := len(picture), len(picture[0])
    colCount, rowHashCount := make([]int, n), map[string]int{}
    for row := 0; row < m; row++ {
        rowCount := 0
        for col := 0; col < n; col++ {
            if picture[row][col] == 'B' { // 统计每一列的B数量
                rowCount++
                colCount[col]++
            }
        }
        if rowCount == target { // 统计每一行的B数量，达到 target 的，将当前行的哈希值 rowHashCount + 1
            rowHashCount[string(picture[row])]++
        }
    }
    res := 0
    for k, v := range rowHashCount { // 遍历 rowHashCount，碰到值为 target 的继续判断
        if v != target {
            continue
        }
        for col := 0; col < n; col++ { 
            if k[col] == 'B' && colCount[col] == target { // 对于每一列，如果本列为B，且计数为N，则结果+N
                res += target
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/pixel2-1-grid.jpg" />
    // Input: picture = [["W","B","W","B","B","W"],["W","B","W","B","B","W"],["W","B","W","B","B","W"],["W","W","B","W","B","W"]], target = 3
    // Output: 6
    // Explanation: All the green 'B' are the black pixels we need (all 'B's at column 1 and 3).
    // Take 'B' at row r = 0 and column c = 1 as an example:
    //  - Rule 1, row r = 0 and column c = 1 both have exactly target = 3 black pixels. 
    //  - Rule 2, the rows have black pixel at column c = 1 are row 0, row 1 and row 2. They are exactly the same as row r = 0.
    picture1 := [][]byte{
        {'W','B','W','B','B','W'},
        {'W','B','W','B','B','W'},
        {'W','B','W','B','B','W'},
        {'W','W','B','W','B','W'},
    }
    fmt.Println(findBlackPixel(picture1,3)) // 6
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/pixel2-2-grid.jpg" />
    // Input: picture = [["W","W","B"],["W","W","B"],["W","W","B"]], target = 1
    // Output: 0
    picture2 := [][]byte{
        {'W','W','B'},
        {'W','W','B'},
        {'W','W','B'},
    }
    fmt.Println(findBlackPixel(picture2,1)) // 0
}