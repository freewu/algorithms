package main

// 499. The Maze III
// There is a ball in a maze with empty spaces (represented as 0) and walls (represented as 1). 
// The ball can go through the empty spaces by rolling up, down, left or right, but it won't stop rolling until hitting a wall. 
// When the ball stops, it could choose the next direction. 
// There is also a hole in this maze. 
// The ball will drop into the hole if it rolls onto the hole.

// Given the m x n maze, the ball's position ball and the hole's position hole, 
// where ball = [ballrow, ballcol] and hole = [holerow, holecol], return a string instructions of all the instructions 
// that the ball should follow to drop in the hole with the shortest distance possible. 
// If there are multiple valid instructions, return the lexicographically minimum one. 
// If the ball can't drop in the hole, return "impossible".

// If there is a way for the ball to drop in the hole, 
// the answer instructions should contain the characters 'u' (i.e., up), 'd' (i.e., down), 'l' (i.e., left), and 'r' (i.e., right).

// The distance is the number of empty spaces traveled by the ball from the start position (excluded) to the destination (included).
// You may assume that the borders of the maze are all walls (see examples).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/31/maze3-1-grid.jpg" />
// Input: maze = [[0,0,0,0,0],[1,1,0,0,1],[0,0,0,0,0],[0,1,0,0,1],[0,1,0,0,0]], ball = [4,3], hole = [0,1]
// Output: "lul"
// Explanation: There are two shortest ways for the ball to drop into the hole.
// The first way is left -> up -> left, represented by "lul".
// The second way is up -> left, represented by 'ul'.
// Both ways have shortest distance 6, but the first way is lexicographically smaller because 'l' < 'u'. So the output is "lul".

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/03/31/maze3-2-grid.jpg" />
// Input: maze = [[0,0,0,0,0],[1,1,0,0,1],[0,0,0,0,0],[0,1,0,0,1],[0,1,0,0,0]], ball = [4,3], hole = [3,0]
// Output: "impossible"
// Explanation: The ball cannot reach the hole.

// Example 3:
// Input: maze = [[0,0,0,0,0,0,0],[0,0,1,0,0,1,0],[0,0,0,0,1,0,0],[0,0,0,0,0,0,1]], ball = [0,4], hole = [3,5]
// Output: "dldr"
 
// Constraints:
//     m == maze.length
//     n == maze[i].length
//     1 <= m, n <= 100
//     maze[i][j] is 0 or 1.
//     ball.length == 2
//     hole.length == 2
//     0 <= ballrow, holerow <= m
//     0 <= ballcol, holecol <= n
//     Both the ball and the hole exist in an empty space, and they will not be in the same position initially.
//     The maze contains at least 2 empty spaces.

import "fmt"
import "strings"

// dfs
func findShortestWay(maze [][]int, ball []int, hole []int) string {
    type PathNode struct{
        x,y,distance int
        path string
    }
    m, n := len(maze),len(maze[0])
    dirs := [][]int{{1,0},{0,-1},{0,1},{-1,0}}
    dirstring := []string{"d","l","r","u"}
    distance := make([][]PathNode,m)
    for i := 0; i < m;i++ {
        distance[i] = make([]PathNode,n)
        for j := 0; j < n; j++ {
            distance[i][j] = PathNode{0,0,0x3f3f3f3f,""}
        }
    }
    distance[ball[0]][ball[1]] = PathNode{ball[0],ball[1],0,""}
    var dfs func(x,y int)
    dfs = func(x,y int){
        for i:=0;i<4;i++{
            dir := dirs[i]
            nx,ny := x+dir[0],y+dir[1]
            count := 0
            for 0<=nx && nx<m && 0<=ny && ny<n && maze[nx][ny] == 0{
                nx += dir[0];ny+=dir[1];count++
                if nx-dir[0] == hole[0] && ny-dir[1] == hole[1]{
                    break
                }
            }
            nx-=dir[0];ny-=dir[1]
            if distance[x][y].distance + count < distance[nx][ny].distance || 
               ( distance[x][y].distance + count == distance[nx][ny].distance && 
                 strings.Compare(distance[nx][ny].path,distance[x][y].path + dirstring[i]) > 0) {
                distance[nx][ny].distance = distance[x][y].distance + count
                distance[nx][ny].path = distance[x][y].path + dirstring[i]
                dfs(nx,ny)
            }
        }
    }
    dfs(ball[0],ball[1])
    if distance[hole[0]][hole[1]].distance == 0x3f3f3f3f {
        return "impossible"
    }
    return distance[hole[0]][hole[1]].path
}

// bfs
func findShortestWay1(maze [][]int, ball []int, hole []int) string {
    type PathNode struct{
        x,y,distance int
        path string
    }
    m,n := len(maze),len(maze[0])
    dirs := [][]int{{1,0},{0,-1},{0,1},{-1,0}}
    dirstring:=[]string{"d","l","r","u"}
    distance := make([][]PathNode,m)
    for i:=0;i<m;i++{
        distance[i] = make([]PathNode,n)
        for j:=0;j<n;j++{
            distance[i][j] = PathNode{0,0,0x3f3f3f3f,""}
        }
    }
    distance[ball[0]][ball[1]] = PathNode{ball[0],ball[1],0,""}
    deque := [][]int{ball}
    for len(deque)>0{
        q := deque[0];deque = deque[1:]
        x,y := q[0],q[1]
        for i:=0;i<4;i++{
            dir := dirs[i]
            nx,ny := x+dir[0],y+dir[1]
            count := 0
            for 0<=nx && nx<m && 0<=ny && ny<n && maze[nx][ny] == 0{
                nx += dir[0];ny+=dir[1];count++
                if nx-dir[0] == hole[0] && ny-dir[1] == hole[1]{break}
            }
            nx-=dir[0];ny-=dir[1]
            if distance[x][y].distance+count < distance[nx][ny].distance || (distance[x][y].distance+count == distance[nx][ny].distance && strings.Compare(distance[nx][ny].path,distance[x][y].path + dirstring[i])>0){
                distance[nx][ny].distance = distance[x][y].distance+count
                distance[nx][ny].path = distance[x][y].path + dirstring[i]
                deque = append(deque, []int{nx,ny})
            }
        }
    }
    if distance[hole[0]][hole[1]].distance==0x3f3f3f3f{
        return "impossible"
    }
    return distance[hole[0]][hole[1]].path
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/31/maze3-1-grid.jpg" />
    // Input: maze = [[0,0,0,0,0],[1,1,0,0,1],[0,0,0,0,0],[0,1,0,0,1],[0,1,0,0,0]], ball = [4,3], hole = [0,1]
    // Output: "lul"
    // Explanation: There are two shortest ways for the ball to drop into the hole.
    // The first way is left -> up -> left, represented by "lul".
    // The second way is up -> left, represented by 'ul'.
    // Both ways have shortest distance 6, but the first way is lexicographically smaller because 'l' < 'u'. So the output is "lul".
    maze1 := [][]int{
        {0,0,0,0,0},
        {1,1,0,0,1},
        {0,0,0,0,0},
        {0,1,0,0,1},
        {0,1,0,0,0},
    }
    fmt.Println(findShortestWay(maze1,[]int{4,3},[]int{0,1})) // "lul"
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/03/31/maze3-2-grid.jpg" />
    // Input: maze = [[0,0,0,0,0],[1,1,0,0,1],[0,0,0,0,0],[0,1,0,0,1],[0,1,0,0,0]], ball = [4,3], hole = [3,0]
    // Output: "impossible"
    // Explanation: The ball cannot reach the hole.
    maze2 := [][]int{
        {0,0,0,0,0},
        {1,1,0,0,1},
        {0,0,0,0,0},
        {0,1,0,0,1},
        {0,1,0,0,0},
    }
    fmt.Println(findShortestWay(maze2,[]int{4,3},[]int{3,0})) // "impossible"
    // Example 3:
    // Input: maze = [[0,0,0,0,0,0,0],[0,0,1,0,0,1,0],[0,0,0,0,1,0,0],[0,0,0,0,0,0,1]], ball = [0,4], hole = [3,5]
    // Output: "dldr"
    maze3 := [][]int{
        {0,0,0,0,0,0,0},
        {0,0,1,0,0,1,0},
        {0,0,0,0,1,0,0},
        {0,0,0,0,0,0,1},
    }
    fmt.Println(findShortestWay(maze3,[]int{0,4},[]int{3,5})) // "dldr"

    fmt.Println(findShortestWay1(maze1,[]int{4,3},[]int{0,1})) // "lul"
    fmt.Println(findShortestWay1(maze2,[]int{4,3},[]int{3,0})) // "impossible"
    fmt.Println(findShortestWay1(maze3,[]int{0,4},[]int{3,5})) // "dldr"
}