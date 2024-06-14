package main

// 489. Robot Room Cleaner
// You are controlling a robot that is located somewhere in a room. 
// The room is modeled as an m x n binary grid where 0 represents a wall and 1 represents an empty slot.

// The robot starts at an unknown location in the room that is guaranteed to be empty, 
// nd you do not have access to the grid, but you can move the robot using the given API Robot.

// You are tasked to use the robot to clean the entire room (i.e., clean every empty cell in the room). 
// The robot with the four given APIs can move forward, turn left, or turn right. Each turn is 90 degrees.

// When the robot tries to move into a wall cell, its bumper sensor detects the obstacle, and it stays on the current cell.
// Design an algorithm to clean the entire room using the following APIs:
//     interface Robot {
//     // returns true if next cell is open and robot moves into the cell.
//     // returns false if next cell is obstacle and robot stays on the current cell.
//     boolean move();

//     // Robot will stay on the same cell after calling turnLeft/turnRight.
//     // Each turn will be 90 degrees.
//     void turnLeft();
//     void turnRight();

//     // Clean the current cell.
//     void clean();
//     }

// Note that the initial direction of the robot will be facing up. 
// You can assume all four edges of the grid are all surrounded by a wall.
// Custom testing:
//     The input is only given to initialize the room and the robot's position internally. 
//     You must solve this problem "blindfolded". 
//     In other words, you must control the robot using only the four mentioned APIs without knowing the room layout and the initial robot's position.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/17/lc-grid.jpg" />
// Input: room = [[1,1,1,1,1,0,1,1],[1,1,1,1,1,0,1,1],[1,0,1,1,1,1,1,1],[0,0,0,1,0,0,0,0],[1,1,1,1,1,1,1,1]], row = 1, col = 3
// Output: Robot cleaned all rooms.
// Explanation: All grids in the room are marked by either 0 or 1.
// 0 means the cell is blocked, while 1 means the cell is accessible.
// The robot initially starts at the position of row=1, col=3.
// From the top left corner, its position is one row below and three columns right.

// Example 2:
// Input: room = [[1]], row = 0, col = 0
// Output: Robot cleaned all rooms.
 
// Constraints:
//     m == room.length
//     n == room[i].length
//     1 <= m <= 100
//     1 <= n <= 200
//     room[i][j] is either 0 or 1.
//     0 <= row < m
//     0 <= col < n
//     room[row][col] == 1
//     All the empty cells can be visited from the starting position.

import "fmt"

/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 * type Robot struct {
 * }
 * 
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */

// dfs
func cleanRoom(robot *Robot) {
    dirs := [][]int{{-1,0}, {0,1}, {1,0}, {0,-1}}
    visited := make(map[int]map[int]bool)
    back := func() {
        robot.TurnRight()
        robot.TurnRight()
        robot.Move();
        robot.TurnRight()
        robot.TurnRight()
    }
    var backtrack func(r,c,d int)
    backtrack = func(r,c,d int){
        if _,ok := visited[r]; !ok {
            visited[r] = make(map[int]bool)
        }
        visited[r][c] = true
        robot.Clean()
        for i := 0; i < 4; i++ {
            nd := (d+i) % 4
            nr, nc := r + dirs[nd][0], c + dirs[nd][1]
            if v,ok := visited[nr];(!ok||!v[nc]) && robot.Move() {
                backtrack(nr,nc,nd)
                back()
            }
            robot.TurnRight();
        }
    }
    backtrack(0,0,0)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/07/17/lc-grid.jpg" />
    // Input: room = [[1,1,1,1,1,0,1,1],[1,1,1,1,1,0,1,1],[1,0,1,1,1,1,1,1],[0,0,0,1,0,0,0,0],[1,1,1,1,1,1,1,1]], row = 1, col = 3
    // Output: Robot cleaned all rooms.
    // Explanation: All grids in the room are marked by either 0 or 1.
    // 0 means the cell is blocked, while 1 means the cell is accessible.
    // The robot initially starts at the position of row=1, col=3.
    // From the top left corner, its position is one row below and three columns right.
    fmt.Println()
    // Example 2:
    // Input: room = [[1]], row = 0, col = 0
    // Output: Robot cleaned all rooms.
}