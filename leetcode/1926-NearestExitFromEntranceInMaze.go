package main

// 1926. Nearest Exit from Entrance in Maze
// You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.') and walls (represented as '+'). 
// You are also given the entrance of the maze, 
// where entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are initially standing at.

// In one step, you can move one cell up, down, left, or right. 
// You cannot step into a cell with a wall, and you cannot step outside the maze. 
// Your goal is to find the nearest exit from the entrance. 
// An exit is defined as an empty cell that is at the border of the maze. The entrance does not count as an exit.

// Return the number of steps in the shortest path from the entrance to the nearest exit, or -1 if no such path exists.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/04/nearest1-grid.jpg" />
// Input: maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
// Output: 1
// Explanation: There are 3 exits in this maze at [1,0], [0,2], and [2,3].
// Initially, you are at the entrance cell [1,2].
// - You can reach [1,0] by moving 2 steps left.
// - You can reach [0,2] by moving 1 step up.
// It is impossible to reach [2,3] from the entrance.
// Thus, the nearest exit is [0,2], which is 1 step away.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/04/nearesr2-grid.jpg" />
// Input: maze = [["+","+","+"],[".",".","."],["+","+","+"]], entrance = [1,0]
// Output: 2
// Explanation: There is 1 exit in this maze at [1,2].
// [1,0] does not count as an exit since it is the entrance cell.
// Initially, you are at the entrance cell [1,0].
// - You can reach [1,2] by moving 2 steps right.
// Thus, the nearest exit is [1,2], which is 2 steps away.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/06/04/nearest3-grid.jpg" />
// Input: maze = [[".","+"]], entrance = [0,0]
// Output: -1
// Explanation: There are no exits in this maze.

// Constraints:
//     maze.length == m
//     maze[i].length == n
//     1 <= m, n <= 100
//     maze[i][j] is either '.' or '+'.
//     entrance.length == 2
//     0 <= entrancerow < m
//     0 <= entrancecol < n
//     entrance will always be an empty cell.

import "fmt"

// bfs
func nearestExit(maze [][]byte, entrance []int) int {
    res, m, n, q := 0, len(maze), len(maze[0]), [][2]int{{entrance[0], entrance[1]}}
    maze[entrance[0]][entrance[1]] = '+'
    for len(q) > 0 {
        for _, v := range q {
            q = q[1:]
            if (v[0] == m-1 || v[1] == n-1 || v[0] == 0 || v[1] == 0) && !(v[0] == entrance[0] && v[1] == entrance[1]) { // 到达出口
                return res
            }
            for _, div := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
                x, y := v[0] + div[0], v[1]+div[1]
                if x >= 0 && y >= 0 && x < m && y < n && maze[x][y] != '+' {  // 不是墙且未出界加入到队列中
                    q = append(q, [2]int{x, y})
                    maze[x][y] = '+'
                }
            }
        }
        res++
    }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/04/nearest1-grid.jpg" />
    // Input: maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
    // Output: 1
    // Explanation: There are 3 exits in this maze at [1,0], [0,2], and [2,3].
    // Initially, you are at the entrance cell [1,2].
    // - You can reach [1,0] by moving 2 steps left.
    // - You can reach [0,2] by moving 1 step up.
    // It is impossible to reach [2,3] from the entrance.
    // Thus, the nearest exit is [0,2], which is 1 step away.
    fmt.Println(nearestExit([][]byte{{'+','+','.','+'},{'.','.','.','+'},{'+','+','+','.'}},[]int{1,2})) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/04/nearesr2-grid.jpg" />
    // Input: maze = [["+","+","+"],[".",".","."],["+","+","+"]], entrance = [1,0]
    // Output: 2
    // Explanation: There is 1 exit in this maze at [1,2].
    // [1,0] does not count as an exit since it is the entrance cell.
    // Initially, you are at the entrance cell [1,0].
    // - You can reach [1,2] by moving 2 steps right.
    // Thus, the nearest exit is [1,2], which is 2 steps away.
    fmt.Println(nearestExit([][]byte{{'+','+','+'},{'.','.','.'},{'+','+','+'}},[]int{1,0})) // 2
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/06/04/nearest3-grid.jpg" />
    // Input: maze = [[".","+"]], entrance = [0,0]
    // Output: -1
    // Explanation: There are no exits in this maze.
    fmt.Println(nearestExit([][]byte{{'.','+'}},[]int{0,0})) // -1
}