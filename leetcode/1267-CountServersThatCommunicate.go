package main

// 1267. Count Servers that Communicate
// You are given a map of a server center, represented as a m * n integer matrix grid, 
// where 1 means that on that cell there is a server and 0 means that it is no server. 
// Two servers are said to communicate if they are on the same row or on the same column.

// Return the number of servers that communicate with any other server.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/11/14/untitled-diagram-6.jpg" />
// Input: grid = [[1,0],[0,1]]
// Output: 0
// Explanation: No servers can communicate with others.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/11/13/untitled-diagram-4.jpg" />
// Input: grid = [[1,0],[1,1]]
// Output: 3
// Explanation: All three servers can communicate with at least one other server.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/11/14/untitled-diagram-1-3.jpg" />
// Input: grid = [[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]
// Output: 4
// Explanation: The two servers in the first row can communicate with each other. The two servers in the third column can communicate with each other. The server at right bottom corner can't communicate with any other server.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m <= 250
//     1 <= n <= 250
//     grid[i][j] == 0 or 1

import "fmt"

func countServers(grid [][]int) int {
    res, m, n := 0, len(grid), len(grid[0])
    mark := make([][]int, m) //mark a spot 1 if that row has more than one servers
    for i := 0; i < m; i++ { // mark a spot 2 if the column has more than one servers
        mark[i] = make([]int, n)
    }
    for i := 0; i < m; i ++{
        sum := 0
        for j := 0; j < n; j ++{
            sum += grid[i][j]
        }
        if sum > 1 {
            for j := 0; j < n; j ++{
                if grid[i][j] == 1 {
                    mark[i][j] = 1
                }
            }
        }
    }    
    for j := 0; j < n; j ++{
        sum := 0
        for i := 0; i < m; i ++{
            sum += grid[i][j]
        }
        if sum > 1 {
            for i := 0; i < m; i ++{
                if grid[i][j] == 1 {
                    mark[i][j] = 2
                }
            }
        } 
    }    
    for i := 0; i < m; i ++{
        for j := 0; j < n; j ++{
            if mark[i][j] != 0 {
                res++
            }
        }
    }
    return res
}

func countServers1(grid [][]int) int {
    res, m, n := 0, len(grid), len(grid[0])
    rowServers, colServers := make([]int, m), make([]int, n) // Create arrays to store the count of servers in each row and column
    for i := 0; i < m; i++ { // Count servers in each row and column
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                rowServers[i]++
                colServers[j]++
            }
        }
    }
    for i := 0; i < m; i++ { // Count servers that can communicate with at least one other server
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 && (rowServers[i] > 1 || colServers[j] > 1) {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/11/14/untitled-diagram-6.jpg" />
    // Input: grid = [[1,0],[0,1]]
    // Output: 0
    // Explanation: No servers can communicate with others.
    fmt.Println(countServers([][]int{{1,0},{0,1}})) // 0
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/11/13/untitled-diagram-4.jpg" />
    // Input: grid = [[1,0],[1,1]]
    // Output: 3
    // Explanation: All three servers can communicate with at least one other server.
    fmt.Println(countServers([][]int{{1,0},{1,1}})) // 3
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/11/14/untitled-diagram-1-3.jpg" />
    // Input: grid = [[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]
    // Output: 4
    // Explanation: The two servers in the first row can communicate with each other. The two servers in the third column can communicate with each other. The server at right bottom corner can't communicate with any other server.
    fmt.Println(countServers([][]int{{1,1,0,0},{0,0,1,0},{0,0,1,0},{0,0,0,1}})) // 4

    fmt.Println(countServers([][]int{{1,0},{0,1}})) // 0
    fmt.Println(countServers([][]int{{1,0},{1,1}})) // 3
    fmt.Println(countServers([][]int{{1,1,0,0},{0,0,1,0},{0,0,1,0},{0,0,0,1}})) // 4
}