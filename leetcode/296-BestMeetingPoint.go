package main

// 296. Best Meeting Point
// Given an m x n binary grid grid where each 1 marks the home of one friend, return the minimal total travel distance.
// The total travel distance is the sum of the distances between the houses of the friends and the meeting point.
// The distance is calculated using Manhattan Distance, where distance(p1, p2) = |p2.x - p1.x| + |p2.y - p1.y|.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/14/meetingpoint-grid.jpg" />
// Input: grid = [[1,0,0,0,1],[0,0,0,0,0],[0,0,1,0,0]]
// Output: 6
// Explanation: Given three friends living at (0,0), (0,4), and (2,2).
// The point (0,2) is an ideal meeting point, as the total travel distance of 2 + 2 + 2 = 6 is minimal.
// So return 6.

// Example 2:
// Input: grid = [[1,1]]
// Output: 1

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 200
//     grid[i][j] is either 0 or 1.
//     There will be at least two friends in the grid.

import "fmt"
import "sort"

func minTotalDistance(grid [][]int) int {
    col, lst := make([]int,0), make([]int,0)
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                col = append(col,i)
                lst = append(lst,j)
            }
        }
    }
    sort.Ints(col)
    sort.Ints(lst)
    sum, i, j := 0, 0, len(col) - 1
    for i < j {
        sum += col[j] - col[i]
        sum += lst[j] - lst[i]
        i++
        j--
    }
    return sum
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/14/meetingpoint-grid.jpg" />
    // Input: grid = [[1,0,0,0,1],[0,0,0,0,0],[0,0,1,0,0]]
    // Output: 6
    // Explanation: Given three friends living at (0,0), (0,4), and (2,2).
    // The point (0,2) is an ideal meeting point, as the total travel distance of 2 + 2 + 2 = 6 is minimal.
    // So return 6.
    grid1 := [][]int{
        { 1,0,0,0,1 },
        { 0,0,0,0,0 },
        { 0,0,1,0,0 },
    }
    fmt.Println(minTotalDistance(grid1)) // 6
    // Example 2:
    // Input: grid = [[1,1]]
    // Output: 1
    grid2 := [][]int{ { 1,1 } }
    fmt.Println(minTotalDistance(grid2)) // 1
}