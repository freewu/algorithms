package main

// 2061. Number of Spaces Cleaning Robot Cleaned
// A room is represented by a 0-indexed 2D binary matrix room where a 0 represents an empty space and a 1 represents a space with an object. 
// The top left corner of the room will be empty in all test cases.

// A cleaning robot starts at the top left corner of the room and is facing right. 
// The robot will continue heading straight until it reaches the edge of the room or it hits an object, 
// after which it will turn 90 degrees clockwise and repeat this process. 
// The starting space and all spaces that the robot visits are cleaned by it.

// Return the number of clean spaces in the room if the robot runs indefinitely.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/01/image-20211101204703-1.png" />
// Input: room = [[0,0,0],[1,1,0],[0,0,0]]
// Output: 7
// Explanation:
// ​​​​​​​The robot cleans the spaces at (0, 0), (0, 1), and (0, 2).
// The robot is at the edge of the room, so it turns 90 degrees clockwise and now faces down.
// The robot cleans the spaces at (1, 2), and (2, 2).
// The robot is at the edge of the room, so it turns 90 degrees clockwise and now faces left.
// The robot cleans the spaces at (2, 1), and (2, 0).
// The robot has cleaned all 7 empty spaces, so return 7.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/11/01/image-20211101204736-2.png" />
// Input: room = [[0,1,0],[1,0,0],[0,0,0]]
// Output: 1
// Explanation:
// The robot cleans the space at (0, 0).
// The robot hits an object, so it turns 90 degrees clockwise and now faces down.
// The robot hits an object, so it turns 90 degrees clockwise and now faces left.
// The robot is at the edge of the room, so it turns 90 degrees clockwise and now faces up.
// The robot is at the edge of the room, so it turns 90 degrees clockwise and now faces right.
// The robot is back at its starting position.
// The robot has cleaned 1 space, so return 1.

// Example 3:
// Input: room = [[0,0,0],[0,0,0],[0,0,0]]
// Output: 8​​​​​​​

// Constraints:
//     m == room.length
//     n == room[r].length
//     1 <= m, n <= 300
//     room[r][c] is either 0 or 1.
//     room[0][0] == 0

import "fmt"

func numberOfCleanRooms(room [][]int) int {
    directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    m, n := len(room), len(room[0])
    x, y, dirIndex := 0 ,0, 0
    visited, cleaned :=map[int]bool{}, map[int]bool{}
    for {
        if visited[x * n + y + dirIndex * m * n] { break }
        visited[x * n + y + dirIndex * m * n] = true
        cleaned[x * n + y] = true
        nx, ny := x + directions[dirIndex][0], y + directions[dirIndex][1]
        if nx >= 0 && nx < m && ny >= 0 && ny < n && room[nx][ny] == 0 {
            x, y = nx, ny
        } else {
            dirIndex = (dirIndex + 1) % 4
        }
    }
    return len(cleaned)
}

// dfs
func numberOfCleanRooms1(room [][]int) int {
    directions := []struct{ X, Y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右->下->左->上->(右->...无限重复)
    type Pair struct{ X, Y int }
    visited := map[Pair]*[4]bool{}
    m, n := len(room), len(room[0])
    var dfs func(x, y, dir int)
    dfs = func(x, y, dir int) {
        p := Pair{x, y}
        flag := visited[p]
        if flag == nil { // 注意!! map不允许直接对它管理的val_struct进行赋值(防止扩容带来的问题), 所以只能声明为指针类型,将val_struct搬到map之外
            flag = &[4]bool{}
            visited[p] = flag
        }
        if flag[dir] { return } // 发现重复访问了,后续都是循环
        flag[dir] = true
        nx, ny := x + directions[dir].X, y + directions[dir].Y
        // 注意!! dfs走的下一步可以在当前层处理,也可以交由后续dfs自身判断.这里有两种不同的类型,父节点负责判断较方便
        if 0 <= nx && nx < m && 0 <= ny && ny < n && room[nx][ny] == 0 {
            dfs(nx, ny, dir)
        } else {
            dfs(x, y, (dir + 1) % 4) // trick!! 换个方向,无需for循环换到可走的方向,由接下来的格子负责下一个方向即可.
        }
    }
    dfs(0, 0, 0)
    return len(visited)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/11/01/image-20211101204703-1.png" />
    // Input: room = [[0,0,0],[1,1,0],[0,0,0]]
    // Output: 7
    // Explanation:
    // ​​​​​​​The robot cleans the spaces at (0, 0), (0, 1), and (0, 2).
    // The robot is at the edge of the room, so it turns 90 degrees clockwise and now faces down.
    // The robot cleans the spaces at (1, 2), and (2, 2).
    // The robot is at the edge of the room, so it turns 90 degrees clockwise and now faces left.
    // The robot cleans the spaces at (2, 1), and (2, 0).
    // The robot has cleaned all 7 empty spaces, so return 7.
    fmt.Println(numberOfCleanRooms([][]int{{0,0,0},{1,1,0},{0,0,0}})) // 7
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/11/01/image-20211101204736-2.png" />
    // Input: room = [[0,1,0],[1,0,0],[0,0,0]]
    // Output: 1
    // Explanation:
    // The robot cleans the space at (0, 0).
    // The robot hits an object, so it turns 90 degrees clockwise and now faces down.
    // The robot hits an object, so it turns 90 degrees clockwise and now faces left.
    // The robot is at the edge of the room, so it turns 90 degrees clockwise and now faces up.
    // The robot is at the edge of the room, so it turns 90 degrees clockwise and now faces right.
    // The robot is back at its starting position.
    // The robot has cleaned 1 space, so return 1.
    fmt.Println(numberOfCleanRooms([][]int{{0,1,0},{1,0,0},{0,0,0}})) // 1
    // Example 3:
    // Input: room = [[0,0,0],[0,0,0],[0,0,0]]
    // Output: 8​​​​​​​
    fmt.Println(numberOfCleanRooms([][]int{{0,0,0},{0,0,0},{0,0,0}})) // 8

    fmt.Println(numberOfCleanRooms1([][]int{{0,0,0},{1,1,0},{0,0,0}})) // 7
    fmt.Println(numberOfCleanRooms1([][]int{{0,1,0},{1,0,0},{0,0,0}})) // 1
    fmt.Println(numberOfCleanRooms1([][]int{{0,0,0},{0,0,0},{0,0,0}})) // 8
}