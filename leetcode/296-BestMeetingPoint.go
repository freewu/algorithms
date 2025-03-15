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

func minTotalDistance1(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    if m == 0 {
        return 0
    }
    rows, cols := []int{}, []int{}
    // 按行遍历收集所有朋友的行坐标，天然有序
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                rows = append(rows, i)
            }
        }
    }
    // 按列遍历收集所有朋友的列坐标，保证有序
    for j := 0; j < n; j++ {
        for i := 0; i < m; i++ {
            if grid[i][j] == 1 {
                cols = append(cols, j)
            }
        }
    }
    getDistance := func(coords []int) int { // 使用双指针方法计算坐标列表中各点到中位数的距离和
        i, j := 0, len(coords)-1
        distance := 0
        for i < j {
            distance += coords[j] - coords[i]
            i++
            j--
        }
        return distance
    }
    // 计算行和列方向的总距离
    return getDistance(rows) + getDistance(cols)
}


func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/14/meetingpoint-grid.jpg" />
    // Input: grid = [[1,0,0,0,1],[0,0,0,0,0],[0,0,1,0,0]]
    // Output: 6
    // Explanation: Given three friends living at (0,0), (0,4), and (2,2).
    // The point (0,2) is an ideal meeting point, as the total travel distance of 2 + 2 + 2 = 6 is minimal.
    // So return 6.
    fmt.Println(minTotalDistance([][]int{{1,0,0,0,1},{0,0,0,0,0},{0,0,1,0,0}})) // 6
    // Example 2:
    // Input: grid = [[1,1]]
    // Output: 1
    fmt.Println(minTotalDistance([][]int{ { 1,1 } })) // 1

    fmt.Println(minTotalDistance1([][]int{{1,0,0,0,1},{0,0,0,0,0},{0,0,1,0,0}})) // 6
    fmt.Println(minTotalDistance1([][]int{ { 1,1 } })) // 1
}