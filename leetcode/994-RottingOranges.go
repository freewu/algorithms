package main

// 994. Rotting Oranges
// You are given an m x n grid where each cell can have one of three values:
//     0 representing an empty cell,
//     1 representing a fresh orange, or
//     2 representing a rotten orange.

// Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

// Return the minimum number of minutes that must elapse until no cell has a fresh orange. 
// If this is impossible, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/02/16/oranges.png" />
// Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
// [2,1,1]         [2,2,1]         [2,2,2]        [2,2,2]        [2,2,2]
// [1,1,0] = 1 =>  [2,1,0] = 2 =>  [2,2,0] = 3 => [2,2,0] = 4 => [2,2,0] 
// [0,1,1]         [0,1,1]         [0,1,1]        [0,2,1]        [0,2,2] 
// Output: 4

// Example 2:
// Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
// [2,1,1]         [2,2,1]         [2,2,2]        [2,2,2]        [2,2,2]
// [0,1,1] = 1 =>  [0,1,1] = 2 =>  [0,2,1] = 3 => [0,2,2] = 4 => [0,2,2] .... -1
// [1,0,1]         [1,0,1]         [1,0,1]        [1,0,1]        [1,0,2] 
// Output: -1
// Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.

// Example 3:
// Input: grid = [[0,2]]
// Output: 0
// Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 10
//     grid[i][j] is 0, 1, or 2.

import "fmt"

// bfs
func orangesRotting(grid [][]int) int {
    directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    m, n, rotten, fresh := len(grid), len(grid[0]),[][2]int{}, 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            switch grid[i][j] {
            case 1: // 新鲜的
                fresh++
            case 2: // 记录烂橘子的坐标
                rotten = append(rotten, [2]int{i, j})
            }
        }
    }
    if fresh == 0 { // 没有新鲜的橘子直接返回 0 不需要花时间去感染了
        return 0
    }
    minutes := 0
    for len(rotten) > 0 {
        for _, o := range rotten {
            for _, dir := range directions {
                x, y := o[0]+dir[0], o[1]+dir[1]
                // 感染新鲜的橘子
                if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
                    grid[x][y] = 2
                    rotten = append(rotten, [2]int{x, y})
                    fresh--
                }
            }
            // 出队列
            rotten = rotten[1:]
        }
        minutes++
        if fresh == 0 {
            return minutes
        }
    }
    return -1
}

func orangesRotting1(grid [][]int) int {
    rotten, fresh := [][2]int{}, 0
    for i, row := range grid {
        for j, flag := range row {
            switch flag {
            case 1: fresh++; // 统计新鲜橘子数
            case 2: rotten = append(rotten, [2]int{i, j}); // 烂橘子坐标
            }
        }
    }
    minute := 0
    for fresh > 0 && len(rotten) > 0 {
        // 每分钟所有的烂橘子都要出来,感染自己四方的,
        // 如果没有可以感染的新鲜橘子 rotten 队列就会归 0
        minute++
        l := len(rotten)
        for i := 0; i < l; i++ {
            node := rotten[0]
            rotten = rotten[1:] // 出队列
            r, c := node[0], node[1]
            if r - 1 >= 0 && grid[r-1][c] == 1 { // 左边存在新鲜橘子，感染它
                grid[r-1][c] = 3
                fresh--
                rotten = append(rotten, [2]int{r - 1, c})
            }
            if r + 1 < len(grid) && grid[r+1][c] == 1 {  // 右边存在新鲜橘子，感染它
                grid[r+1][c] = 3
                fresh--
                rotten = append(rotten, [2]int{r + 1, c})
            }
            if c-1 >= 0 && grid[r][c-1] == 1 { // 上方存在新鲜橘子，感染它
                grid[r][c-1] = 3
                fresh--
                rotten = append(rotten, [2]int{r, c - 1})
            }
            if c+1 < len(grid[0]) && grid[r][c+1] == 1 { // 下方存在新鲜橘子，感染它
                grid[r][c+1] = 3
                fresh--
                rotten = append(rotten, [2]int{r, c + 1})
            }
        }
    }
    if fresh > 0 { // 存在没办法感染的橘子
        return -1
    }
    return minute
}

func main() {
    fmt.Println(orangesRotting([][]int{{2,1,1},{1,1,0},{0,1,1}})) // 4
    // Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
    fmt.Println(orangesRotting([][]int{{2,1,1},{0,1,1},{1,0,1}})) // -1
    // Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
    fmt.Println(orangesRotting([][]int{{0,2}})) // 0

    fmt.Println(orangesRotting1([][]int{{2,1,1},{1,1,0},{0,1,1}})) // 4
    fmt.Println(orangesRotting1([][]int{{2,1,1},{0,1,1},{1,0,1}})) // -1
    fmt.Println(orangesRotting1([][]int{{0,2}})) // 0
}