package main

// 505. The Maze II
// There is a ball in a maze with empty spaces (represented as 0) and walls (represented as 1). 
// The ball can go through the empty spaces by rolling up, down, left or right, but it won't stop rolling until hitting a wall. 
// When the ball stops, it could choose the next direction.

// Given the m x n maze, the ball's start position and the destination, 
// where start = [startrow, startcol] and destination = [destinationrow, destinationcol], 
// return the shortest distance for the ball to stop at the destination. 
// If the ball cannot stop at destination, return -1.

// The distance is the number of empty spaces traveled by the ball from the start position (excluded) to the destination (included).
// You may assume that the borders of the maze are all walls (see examples).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/31/maze1-1-grid.jpg"/>
// Input: maze = [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]], start = [0,4], destination = [4,4]
// Output: 12
// Explanation: One possible way is : left -> down -> left -> down -> right -> down -> right.
// The length of the path is 1 + 1 + 3 + 1 + 2 + 2 + 2 = 12.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/03/31/maze1-2-grid.jpg"/>
// Input: maze = [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]], start = [0,4], destination = [3,2]
// Output: -1
// Explanation: There is no way for the ball to stop at the destination. Notice that you can pass through the destination but you cannot stop there.

// Example 3:
// Input: maze = [[0,0,0,0,0],[1,1,0,0,1],[0,0,0,0,0],[0,1,0,0,1],[0,1,0,0,0]], start = [4,3], destination = [0,1]
// Output: -1

// Constraints:
//     m == maze.length
//     n == maze[i].length
//     1 <= m, n <= 100
//     maze[i][j] is 0 or 1.
//     start.length == 2
//     destination.length == 2
//     0 <= startrow, destinationrow < m
//     0 <= startcol, destinationcol < n
//     Both the ball and the destination exist in an empty space, and they will not be in the same position initially.
//     The maze contains at least 2 empty spaces.

import "fmt"
import "container/heap"

// bfs
func shortestDistance(maze [][]int, start []int, destination []int) int {
    queue ,rowNum, colNow:= [][]int{start}, len(maze), len(maze[0])
    dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
    stepCount := make([][]int, rowNum)
    for i := 0; i < rowNum;i++{
        stepCount[i] = make([]int, colNow)
        for j := 0; j < colNow;j++ {
            stepCount[i][j] = -1
        }
    }
    stepCount[start[0]][start[1]] = 0
    for len(queue) != 0 {
        top := queue[0]
        queue = queue[1:]
        for _, v := range dir {
            newRow, newCol := top[0] + v[0], top[1] + v[1]
            step := 0
            for newRow >= 0 && newRow < rowNum && newCol >=0 && newCol < colNow &&
                maze[newRow][newCol] == 0 {
                newRow, newCol = newRow + v[0], newCol + v[1]
                step++
            }
            newRow, newCol = newRow - v[0], newCol - v[1]
            if stepCount[newRow][newCol] != -1 && stepCount[top[0]][top[1]] + step >= stepCount[newRow][newCol] {
                continue
            }
            queue = append(queue, []int{newRow, newCol})
            stepCount[newRow][newCol] = stepCount[top[0]][top[1]] + step
        }
    }
    return stepCount[destination[0]][destination[1]]
}

// Point represents a point in the maze with its coordinates and distance from the start point.
type Point struct {
    x, y, dist int
}

// PriorityQueue implements a priority queue for points based on their distance.
type PriorityQueue []Point
func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Point)) }
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    x := old[n-1]
    *pq = old[0 : n-1]
    return x
}

// shortestDistance finds the shortest distance from start to destination in the maze.
func shortestDistance1(maze [][]int, start []int, destination []int) int {
    m, n := len(maze), len(maze[0])
    dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Possible directions: up, right, down, left
    // Initialize distances with -1, indicating unreachable points
    distance := make([][]int, m)
    for i := range distance {
        distance[i] = make([]int, n)
        for j := range distance[i] {
            distance[i][j] = -1 // Initialize distance to -1 (unvisited)
        }
    }
    // Initialize priority queue
    pq := &PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, Point{start[0], start[1], 0}) // Push starting point with distance 0
    distance[start[0]][start[1]] = 0            // Update distance to starting point
    for pq.Len() > 0 {
        p := heap.Pop(pq).(Point) // Pop the point with smallest distance from the priority queue
        if p.x == destination[0] && p.y == destination[1] {
            return p.dist // If reached destination, return distance
        }
        for _, dir := range dirs {
            x, y, dist := p.x, p.y, p.dist
            for newX, newY := x+dir[0], y+dir[1]; newX >= 0 && newX < m && newY >= 0 && newY < n && maze[newX][newY] == 0; newX, newY = newX+dir[0], newY+dir[1] {
                x, y = newX, newY
                dist++
            }
            if distance[x][y] == -1 || dist < distance[x][y] {
                distance[x][y] = dist
                heap.Push(pq, Point{x, y, dist}) // Push new point with updated distance to the priority queue
            }
        }
    }
    return -1 // If destination cannot be reached, return -1
}

func shortestDistance2(maze [][]int, start []int, destination []int) int {
    m, n := len(maze), len(maze[0])
    distance := make([][]int, m)
    for i := range distance {
        distance[i] = make([]int, n)
        for j := range distance[i] {
            distance[i][j] =  1 << 61
        }
    }
    distance[start[0]][start[1]] = 0
    dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
    queue := [][]int{start}
    for len(queue) > 0 {
        s := queue[0]
        queue = queue[1:]
        for _, dir := range dirs {
            x, y := s[0] + dir[0], s[1] + dir[1]
            count := 0
            for x >= 0 && y >= 0 && x < m && y < n && maze[x][y] == 0 {
                x += dir[0]
                y += dir[1]
                count++
            }
            newX, newY := x - dir[0], y - dir[1]
            if distance[s[0]][s[1]] + count < distance[newX][newY] {
                distance[newX][newY] = distance[s[0]][s[1]] + count
                queue = append(queue, []int{newX, newY})
            }
        }
    }
    if distance[destination[0]][destination[1]] == 1 << 61 {
        return -1
    }
    return distance[destination[0]][destination[1]]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/31/maze1-1-grid.jpg"/>
    // Input: maze = [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]], start = [0,4], destination = [4,4]
    // Output: 12
    // Explanation: One possible way is : left -> down -> left -> down -> right -> down -> right.
    // The length of the path is 1 + 1 + 3 + 1 + 2 + 2 + 2 = 12.
    maze1 := [][]int{
        {0,0,1,0,0},
        {0,0,0,0,0},
        {0,0,0,1,0},
        {1,1,0,1,1},
        {0,0,0,0,0},
    }
    fmt.Println(shortestDistance(maze1,[]int{0,4}, []int{4,4})) // 12
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/03/31/maze1-2-grid.jpg"/>
    // Input: maze = [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]], start = [0,4], destination = [3,2]
    // Output: -1
    // Explanation: There is no way for the ball to stop at the destination. Notice that you can pass through the destination but you cannot stop there.
    maze2 := [][]int{
        {0,0,1,0,0},
        {0,0,0,0,0},
        {0,0,0,1,0},
        {1,1,0,1,1},
        {0,0,0,0,0},
    }
    fmt.Println(shortestDistance(maze2,[]int{0,4}, []int{3,2})) // -1
    // Example 3:
    // Input: maze = [[0,0,0,0,0],[1,1,0,0,1],[0,0,0,0,0],[0,1,0,0,1],[0,1,0,0,0]], start = [4,3], destination = [0,1]
    // Output: -1
    maze3 := [][]int{
        {0,0,0,0,0},
        {1,1,0,0,1},
        {0,0,0,0,0},
        {0,1,0,0,1},
        {0,1,0,0,0},
    }
    fmt.Println(shortestDistance(maze3,[]int{4,3}, []int{0,1})) // -1

    fmt.Println(shortestDistance1(maze1,[]int{0,4}, []int{4,4})) // 12
    fmt.Println(shortestDistance1(maze2,[]int{0,4}, []int{3,2})) // -1
    fmt.Println(shortestDistance1(maze3,[]int{4,3}, []int{0,1})) // -1

    fmt.Println(shortestDistance2(maze1,[]int{0,4}, []int{4,4})) // 12
    fmt.Println(shortestDistance2(maze2,[]int{0,4}, []int{3,2})) // -1
    fmt.Println(shortestDistance2(maze3,[]int{4,3}, []int{0,1})) // -1
}