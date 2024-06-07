package main

// 490. The Maze
// There is a ball in a maze with empty spaces (represented as 0) and walls (represented as 1). 
// The ball can go through the empty spaces by rolling up, down, left or right, 
// but it won't stop rolling until hitting a wall. 
// When the ball stops, it could choose the next direction.

// Given the m x n maze, the ball's start position and the destination, 
// where start = [startrow, startcol] and destination = [destinationrow, destinationcol], 
// return true if the ball can stop at the destination, otherwise return false.

// You may assume that the borders of the maze are all walls (see examples).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/31/maze1-1-grid.jpg" />
// Input: maze = [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]], start = [0,4], destination = [4,4]
// Output: true
// Explanation: One possible way is : left -> down -> left -> down -> right -> down -> right.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/03/31/maze1-2-grid.jpg" />
// Input: maze = [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]], start = [0,4], destination = [3,2]
// Output: false
// Explanation: There is no way for the ball to stop at the destination. Notice that you can pass through the destination but you cannot stop there.

// Example 3:
// Input: maze = [[0,0,0,0,0],[1,1,0,0,1],[0,0,0,0,0],[0,1,0,0,1],[0,1,0,0,0]], start = [4,3], destination = [0,1]
// Output: false

// Constraints:
//     m == maze.length
//     n == maze[i].length
//     1 <= m, n <= 100
//     maze[i][j] is 0 or 1.
//     start.length == 2
//     destination.length == 2
//     0 <= startrow, destinationrow <= m
//     0 <= startcol, destinationcol <= n
//     Both the ball and the destination exist in an empty space, and they will not be in the same position initially.
//     The maze contains at least 2 empty spaces.

import "fmt"

// bfs
func hasPath(maze [][]int, start []int, destination []int) bool {
    queue ,rowNum, colNow:= [][]int{start}, len(maze), len(maze[0])
    dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
    visited := make([][]bool, rowNum)
    for i := 0; i < rowNum;i++{
        visited[i] = make([]bool, colNow)
    }
    visited[start[0]][start[1]] = true
    for len(queue) != 0 {
        top := queue[0]
        if top[0] == destination[0] && top[1] == destination[1] {
            return true
        }
        queue = queue[1:]
        for _, v := range dir {
            newRow, newCol := top[0] + v[0], top[1] + v[1]
            // 冲到方向最深处
            for newRow >= 0 && newRow < rowNum && newCol >=0 && newCol < colNow &&
                maze[newRow][newCol] == 0 {
                newRow, newCol = newRow + v[0], newCol + v[1]
            }
            newRow, newCol = newRow - v[0], newCol - v[1]
            if visited[newRow][newCol] {
                continue
            }
            queue = append(queue, []int{newRow, newCol})
            visited[newRow][newCol] = true
        }
    }
    return false
}

// dfs
func hasPath1(maze [][]int, start []int, destination []int) bool {
    m, n := len(maze), len(maze[0])
    visit := make([][]bool, m)
    for i, _ := range visit {
        visit[i] = make([]bool, n)
    }
    var dfs func(i,j int)
    dfs = func(i, j int) {
        if visit[i][j] { // 已访问不走了
            return
        } 
        visit[i][j] = true // 标记 
        if i == destination[0] && j == destination[1] { // 到达目的地
            return
        }
        up, down, left, right := i-1, i+1, j-1, j+1
        for right < n && maze[i][right] == 0 {
            right++
        }
        dfs(i, right-1)
        for left>=0 && maze[i][left] == 0 {
            left--
        }
        dfs(i, left+1)
        for up >= 0 && maze[up][j] == 0 {
            up--
        }
        dfs(up+1, j)
        for down < m && maze[down][j] == 0 {
            down++
        }
        dfs(down-1, j)
    }
    dfs(start[0], start[1])
    if visit[destination[0]][destination[1]] == true {
        return true
    }
    return false
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/31/maze1-1-grid.jpg" />
    // Input: maze = [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]], start = [0,4], destination = [4,4]
    // Output: true
    // Explanation: One possible way is : left -> down -> left -> down -> right -> down -> right.
    maze1 := [][]int{
        {0,0,1,0,0},
        {0,0,0,0,0},
        {0,0,0,1,0},
        {1,1,0,1,1},
        {0,0,0,0,0},
    }
    fmt.Println(hasPath(maze1,[]int{0,4},[]int{4,4})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/03/31/maze1-2-grid.jpg" />
    // Input: maze = [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]], start = [0,4], destination = [3,2]
    // Output: false
    // Explanation: There is no way for the ball to stop at the destination. Notice that you can pass through the destination but you cannot stop there.
    maze2 := [][]int{
        {0,0,1,0,0},
        {0,0,0,0,0},
        {0,0,0,1,0},
        {1,1,0,1,1},
        {0,0,0,0,0},
    }
    fmt.Println(hasPath(maze2,[]int{0,4},[]int{3,2})) // false
    // Example 3:
    // Input: maze = [[0,0,0,0,0],[1,1,0,0,1],[0,0,0,0,0],[0,1,0,0,1],[0,1,0,0,0]], start = [4,3], destination = [0,1]
    // Output: false
    maze3 := [][]int{
        {0,0,0,0,0},
        {1,1,0,0,1},
        {0,0,0,0,0},
        {0,1,0,0,1},
        {0,1,0,0,0},
    }
    fmt.Println(hasPath(maze3,[]int{0,4},[]int{3,2})) // false

    fmt.Println(hasPath1(maze1,[]int{0,4},[]int{4,4})) // true
    fmt.Println(hasPath1(maze2,[]int{0,4},[]int{3,2})) // false
    fmt.Println(hasPath1(maze3,[]int{0,4},[]int{3,2})) // false
}