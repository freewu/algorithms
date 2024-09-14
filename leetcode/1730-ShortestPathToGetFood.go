package main

// 1730. Shortest Path to Get Food
// You are starving and you want to eat food as quickly as possible. 
// You want to find the shortest path to arrive at any food cell.

// You are given an m x n character matrix, grid, of these different types of cells:
//     '*' is your location. There is exactly one '*' cell.
//     '#' is a food cell. There may be multiple food cells.
//     'O' is free space, and you can travel through these cells.
//     'X' is an obstacle, and you cannot travel through these cells.

// You can travel to any adjacent cell north, east, south, or west of your current location if there is not an obstacle.

// Return the length of the shortest path for you to reach any food cell. 
// If there is no path for you to reach food, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/21/img1.jpg" />
// Input: grid = [["X","X","X","X","X","X"],["X","*","O","O","O","X"],["X","O","O","#","O","X"],["X","X","X","X","X","X"]]
// Output: 3
// Explanation: It takes 3 steps to reach the food.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/09/21/img2.jpg" />
// Input: grid = [["X","X","X","X","X"],["X","*","X","O","X"],["X","O","X","#","X"],["X","X","X","X","X"]]
// Output: -1
// Explanation: It is not possible to reach the food.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/09/21/img3.jpg" />
// Input: grid = [["X","X","X","X","X","X","X","X"],["X","*","O","X","O","#","O","X"],["X","O","O","X","O","O","X","X"],["X","O","O","O","O","#","O","X"],["X","X","X","X","X","X","X","X"]]
// Output: 6
// Explanation: There can be multiple food cells. It only takes 6 steps to reach the bottom food.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 200
//     grid[row][col] is '*', 'X', 'O', or '#'.
//     The grid contains exactly one '*'.

import "fmt"

// bfs + memo
func getFood(grid [][]byte) int {
    dirs := [][]int{{-1,0},{0,1},{1,0},{0,-1}}
    n, m := len(grid),len(grid[0])
    if n == 1 && m == 1{
        return -1
    }
    
    x, y := -1, -1 // 找到出发点
    visited := make([][]bool,n)
    for i := range visited { 
        visited[i] = make([]bool,m)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == '*' { // 找到起始点
                x, y = i, j
                break
            }
        }
        if x != -1 { break }
    }
    queue := make([][]int,0)
    queue = append(queue,[]int{x,y})
    visited[x][y] = true
    for step := 1; len(queue) > 0; step++ {
        node := len(queue)
        for i := 0; i < node; i++ {
            cur := queue[0] // pop
            queue = queue[1:]
            for _, d := range dirs {
                x, y = cur[0] + d[0], cur[1] + d[1]
                if x >= 0 && x < n &&  y >= 0 && y < m { // 边界内
                    if grid[x][y] != 'X' && !visited[x][y] { // 不是墙且没有访问过
                        if grid[x][y] == '#' { // 找到目标
                            return step
                        }
                        visited[x][y] = true
                        queue = append(queue,[]int{x,y})
                    }
                }
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/21/img1.jpg" />
    // Input: grid = [["X","X","X","X","X","X"],["X","*","O","O","O","X"],["X","O","O","#","O","X"],["X","X","X","X","X","X"]]
    // Output: 3
    // Explanation: It takes 3 steps to reach the food.
    grid1 := [][]byte{
        {'X','X','X','X','X','X'},
        {'X','*','O','O','O','X'},
        {'X','O','O','#','O','X'},
        {'X','X','X','X','X','X'},
    }
    fmt.Println(getFood(grid1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/09/21/img2.jpg" />
    // Input: grid = [["X","X","X","X","X"],["X","*","X","O","X"],["X","O","X","#","X"],["X","X","X","X","X"]]
    // Output: -1
    // Explanation: It is not possible to reach the food.
    grid2 := [][]byte{
        {'X','X','X','X','X'},
        {'X','*','X','O','X'},
        {'X','O','X','#','X'},
        {'X','X','X','X','X'},
    }
    fmt.Println(getFood(grid2)) // -1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/09/21/img3.jpg" />
    // Input: grid = [["X","X","X","X","X","X","X","X"],["X","*","O","X","O","#","O","X"],["X","O","O","X","O","O","X","X"],["X","O","O","O","O","#","O","X"],["X","X","X","X","X","X","X","X"]]
    // Output: 6
    // Explanation: There can be multiple food cells. It only takes 6 steps to reach the bottom food.
    grid3 := [][]byte{
        {'X','X','X','X','X','X','X','X'},
        {'X','*','O','X','O','#','O','X'},
        {'X','O','O','X','O','O','X','X'},
        {'X','O','O','O','O','#','O','X'},
        {'X','X','X','X','X','X','X','X'},
    }
    fmt.Println(getFood(grid3)) // 6
}