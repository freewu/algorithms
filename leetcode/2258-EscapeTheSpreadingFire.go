package main

// 2258. Escape the Spreading Fire
// You are given a 0-indexed 2D integer array grid of size m x n which represents a field. 
// Each cell has one of three values:
//     0 represents grass,
//     1 represents fire,
//     2 represents a wall that you and fire cannot pass through.

// You are situated in the top-left cell, (0, 0), and you want to travel to the safehouse at the bottom-right cell, (m - 1, n - 1). 
// Every minute, you may move to an adjacent grass cell. 
// After your move, every fire cell will spread to all adjacent cells that are not walls.

// Return the maximum number of minutes that you can stay in your initial position before moving while still safely reaching the safehouse. 
// If this is impossible, return -1. If you can always reach the safehouse regardless of the minutes stayed, return 10^9.

// Note that even if the fire spreads to the safehouse immediately after you have reached it, it will be counted as safely reaching the safehouse.

// A cell is adjacent to another cell if the former is directly north, east, south, or west of the latter (i.e., their sides are touching).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/10/ex1new.jpg" />
// Input: grid = [[0,2,0,0,0,0,0],[0,0,0,2,2,1,0],[0,2,0,0,1,2,0],[0,0,2,2,2,0,2],[0,0,0,0,0,0,0]]
// Output: 3
// Explanation: The figure above shows the scenario where you stay in the initial position for 3 minutes.
// You will still be able to safely reach the safehouse.
// Staying for more than 3 minutes will not allow you to safely reach the safehouse.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/10/ex2new2.jpg" />
// Input: grid = [[0,0,0,0],[0,1,2,0],[0,2,0,0]]
// Output: -1
// Explanation: The figure above shows the scenario where you immediately move towards the safehouse.
// Fire will spread to any cell you move towards and it is impossible to safely reach the safehouse.
// Thus, -1 is returned.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/03/10/ex3new.jpg" />
// Input: grid = [[0,0,0],[2,2,0],[1,2,0]]
// Output: 1000000000
// Explanation: The figure above shows the initial grid.
// Notice that the fire is contained by walls and you will always be able to safely reach the safehouse.
// Thus, 10^9 is returned.
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     2 <= m, n <= 300
//     4 <= m * n <= 2 * 10^4
//     grid[i][j] is either 0, 1, or 2.
//     grid[0][0] == grid[m - 1][n - 1] == 0

import "fmt"

func maximumMinutes(grid [][]int) int {
    noFire := int(1e9)
    // 1. 计算每个格子被火烧到的时间
    times := make([][]int, len(grid))      // 记录该格子多久会着火
    minRecord := make([][]int, len(times)) // 记录到达该格子时的最大安全等待时间(用于dfs剪枝)
    queue := [][]int{}
    for i := 0; i < len(times); i++ {
        times[i] = make([]int, len(grid[i]))
        minRecord[i] = make([]int, len(grid[i]))
        for j := 0; j < len(times[i]); j++ {
            if grid[i][j] == 0 {
                times[i][j] = noFire // 没有火
            } else if grid[i][j] == 1 {
                queue = append(queue, []int{i, j})
            }
            minRecord[i][j] = -1
        }
    }
    time := 0
    nextQueue := [][]int{}
    dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        times[cur[0]][cur[1]] = time
        for _, dir := range dirs {
            nextI, nextJ := cur[0]+dir[0], cur[1]+dir[1]
            if 0 <= nextI && nextI < len(grid) && 0 <= nextJ && nextJ < len(grid[0]) && times[nextI][nextJ] == noFire {
                nextQueue = append(nextQueue, []int{nextI, nextJ})
            }
        }
        if len(queue) == 0 {
            queue = nextQueue
            nextQueue = [][]int{}
            time++
        }
    }
    // 2. 计算最长停留分钟
    res := -1 // 若没有可达路径，就是 -1
    var dfs func(i, j, cost, time int)
    dfs = func(i, j, cost, time int) {
        if i == len(times)-1 && j == len(times[0])-1 {
            res = max(res, time)
            return
        }
        if time < res || time <= minRecord[i][j] { // 剪枝，当前路径的最大等待时间已经 < res 或 <= minRecord[i][j]，无须往下继续测试
            return
        }
        minRecord[i][j] = time // 更新到达当前格子的安全等待时间
        for _, dir := range dirs {
            nextI, nextJ := i+dir[0], j+dir[1]
            if 0 <= nextI && nextI < len(times) && 0 <= nextJ && nextJ < len(times[0]) && times[nextI][nextJ] > cost {
                temp := time
                if times[nextI][nextJ] != noFire {  // noFire 的格子无需参与更新 time
                    if nextI == len(times)-1 && nextJ == len(times[0])-1 { // 最后一格比较特殊，最大等待时间可以多 1 min
                        temp = min(temp, times[nextI][nextJ]-cost-1)	
                    } else {
                        temp = min(temp, times[nextI][nextJ]-cost-2)	// 本来应该在 for 循环之前 cost+1 的，但没有，所以这里是 -2
                    }
                }
                dfs(nextI, nextJ, cost+1, temp)	// 这里才更新了 cost
            }
        }

    }
    dfs(0, 0, 0, times[0][0])
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/10/ex1new.jpg" />
    // Input: grid = [[0,2,0,0,0,0,0],[0,0,0,2,2,1,0],[0,2,0,0,1,2,0],[0,0,2,2,2,0,2],[0,0,0,0,0,0,0]]
    // Output: 3
    // Explanation: The figure above shows the scenario where you stay in the initial position for 3 minutes.
    // You will still be able to safely reach the safehouse.
    // Staying for more than 3 minutes will not allow you to safely reach the safehouse.
    grid1 := [][]int{
        {0,2,0,0,0,0,0},
        {0,0,0,2,2,1,0},
        {0,2,0,0,1,2,0},
        {0,0,2,2,2,0,2},
        {0,0,0,0,0,0,0},
    }
    fmt.Println(maximumMinutes(grid1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/10/ex2new2.jpg" />
    // Input: grid = [[0,0,0,0],[0,1,2,0],[0,2,0,0]]
    // Output: -1
    // Explanation: The figure above shows the scenario where you immediately move towards the safehouse.
    // Fire will spread to any cell you move towards and it is impossible to safely reach the safehouse.
    // Thus, -1 is returned.
    grid2 := [][]int{
        {0,0,0,0},
        {0,1,2,0},
        {0,2,0,0},
    }
    fmt.Println(maximumMinutes(grid2)) // 3
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/03/10/ex3new.jpg" />
    // Input: grid = [[0,0,0],[2,2,0],[1,2,0]]
    // Output: 1000000000
    // Explanation: The figure above shows the initial grid.
    // Notice that the fire is contained by walls and you will always be able to safely reach the safehouse.
    // Thus, 10^9 is returned.
    grid3 := [][]int{
        {0,0,0},
        {2,2,0},
        {1,2,0},
    }
    fmt.Println(maximumMinutes(grid3)) // 1000000000
}