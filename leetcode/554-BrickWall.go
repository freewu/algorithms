package main

// 554. Brick Wall
// There is a rectangular brick wall in front of you with n rows of bricks. 
// The ith row has some number of bricks each of the same height (i.e., one unit) but they can be of different widths.
// The total width of each row is the same.

// Draw a vertical line from the top to the bottom and cross the least bricks. 
// If your line goes through the edge of a brick, then the brick is not considered as crossed. 
// You cannot draw a line just along one of the two vertical edges of the wall, in which case the line will obviously cross no bricks.

// Given the 2D array wall that contains the information about the wall, 
// return the minimum number of crossed bricks after drawing such a vertical line.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/cutwall-grid.jpg" />
// Input: wall = [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]
// Output: 2

// Example 2:
// Input: wall = [[1],[1],[1]]
// Output: 3

// Constraints:
//     n == wall.length
//     1 <= n <= 10^4
//     1 <= wall[i].length <= 10^4
//     1 <= sum(wall[i].length) <= 2 * 10^4
//     sum(wall[i]) is the same for each row i.
//     1 <= wall[i][j] <= 2^31 - 1

import "fmt"

func leastBricks(wall [][]int) int {
    mx, m := 0, make(map[int]int)
    for _, row := range wall {
        sum := 0
        for i := 0; i < len(row)-1; i++ {
            sum += row[i] // 累加砖块宽度
            m[sum]++ // 记录每个砖块缝隙的数量
        }
    }
    for _, v := range m { // 找到砖块缝隙最大的值
        if v > mx {
            mx = v
        }
    }
    return len(wall) - mx // 既然穿过砖块中缝不算穿过砖块，那么穿过最少砖块数量一定是穿过很多中缝
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/cutwall-grid.jpg" />
    // Input: wall = [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]
    // Output: 2
    wall1 := [][]int{
        {1,2,2,1},
        {3,1,2},
        {1,3,2},
        {2,4},
        {3,1,2},
        {1,3,1,1},
    }
    fmt.Println(leastBricks(wall1)) // 2
    // Example 2:
    // Input: wall = [[1],[1],[1]]
    // Output: 3
    wall2 := [][]int{
        {1},
        {1},
        {1},
    }
    fmt.Println(leastBricks(wall2)) // 3
}