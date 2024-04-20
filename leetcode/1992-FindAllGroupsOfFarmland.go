package main

// 1992. Find All Groups of Farmland
// You are given a 0-indexed m x n binary matrix land where a 0 represents a hectare of forested land and a 1 represents a hectare of farmland.

// To keep the land organized, there are designated rectangular areas of hectares that consist entirely of farmland. 
// These rectangular areas are called groups.
//  No two groups are adjacent, meaning farmland in one group is not four-directionally adjacent to another farmland in a different group.

// land can be represented by a coordinate system where the top left corner of land is (0, 0) and the bottom right corner of land is (m-1, n-1). 
// Find the coordinates of the top left and bottom right corner of each group of farmland. 
// A group of farmland with a top left corner at (r1, c1) and a bottom right corner at (r2, c2) is represented by the 4-length array [r1, c1, r2, c2].

// Return a 2D array containing the 4-length arrays described above for each group of farmland in land. 
// If there are no groups of farmland, return an empty array. You may return the answer in any order.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/27/screenshot-2021-07-27-at-12-23-15-copy-of-diagram-drawio-diagrams-net.png" />
// Input: land = [[1,0,0],[0,1,1],[0,1,1]]
// Output: [[0,0,0,0],[1,1,2,2]]
// Explanation:
// The first group has a top left corner at land[0][0] and a bottom right corner at land[0][0].
// The second group has a top left corner at land[1][1] and a bottom right corner at land[2][2].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/07/27/screenshot-2021-07-27-at-12-30-26-copy-of-diagram-drawio-diagrams-net.png" />
// Input: land = [[1,1],[1,1]]
// Output: [[0,0,1,1]]
// Explanation:
// The first group has a top left corner at land[0][0] and a bottom right corner at land[1][1].

// Example 3:
// Input: land = [[0]]
// Output: []
// Explanation:
// There are no groups of farmland.
 
// Constraints:
//     m == land.length
//     n == land[i].length
//     1 <= m, n <= 300
//     land consists of only 0's and 1's.
//     Groups of farmland are rectangular in shape.

import "fmt"

func findFarmland(a [][]int) [][]int {
    res, n, m := [][]int{}, len(a), len(a[0])
    for x, row := range a {
        for y, col := range row {
            // 判断是否为矩形左上角，需要满足三个条件：
            //     元素值为 1
            //     左边是边界或者是 0
            //     上边是边界或者是 0
            if col == 0 || y > 0 && row[y-1] == 1 || x > 0 && a[x-1][y] == 1 {
                continue
            }
            // 遍历找到矩形的右边界和下边界
            ii := x
            for ; ii + 1 < n && a[ii + 1][y] == 1; ii++ {}
            jj := y
            for ; jj + 1 < m && a[ii][jj + 1] == 1; jj++ {}
            res = append(res, []int{x, y, ii, jj})
        }
    }
    return res
}

func main() {
    // The first group has a top left corner at land[0][0] and a bottom right corner at land[0][0].
    // The second group has a top left corner at land[1][1] and a bottom right corner at land[2][2].
    fmt.Println(findFarmland([][]int{{1,0,0},{0,1,1},{0,1,1}})) // [[0,0,0,0],[1,1,2,2]]
    // The first group has a top left corner at land[0][0] and a bottom right corner at land[1][1].
    fmt.Println(findFarmland([][]int{{1,1},{1,1}})) // [[0,0,1,1]]
    // There are no groups of farmland.
    fmt.Println(findFarmland([][]int{{0}})) // []
}