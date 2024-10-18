package main

// 1820. Maximum Number of Accepted Invitations
// There are m boys and n girls in a class attending an upcoming party.

// You are given an m x n integer matrix grid, where grid[i][j] equals 0 or 1. 
// If grid[i][j] == 1, then that means the ith boy can invite the jth girl to the party. 
// A boy can invite at most one girl, and a girl can accept at most one invitation from a boy.

// Return the maximum possible number of accepted invitations.

// Example 1:
// Input: grid = [[1,1,1],
//                [1,0,1],
//                [0,0,1]]
// Output: 3
// Explanation: The invitations are sent as follows:
// - The 1st boy invites the 2nd girl.
// - The 2nd boy invites the 1st girl.
// - The 3rd boy invites the 3rd girl.

// Example 2:
// Input: grid = [[1,0,1,0],
//                [1,0,0,0],
//                [0,0,1,0],
//                [1,1,1,0]]
// Output: 3
// Explanation: The invitations are sent as follows:
// -The 1st boy invites the 3rd girl.
// -The 2nd boy invites the 1st girl.
// -The 3rd boy invites no one.
// -The 4th boy invites the 2nd girl.

// Constraints:
//     grid.length == m
//     grid[i].length == n
//     1 <= m, n <= 200
//     grid[i][j] is either 0 or 1.

import "fmt"

func maximumInvitations(grid [][]int) int {
    // 每行每列最多只能保留1个1，返回最多的1的个数
    // 类似八皇后，回溯
    // row从上到下，一个colMap保存已经有1的行回溯
    // 尝试每个行的1，加上剪枝即可，每行也可以不加1，那就有2^m次方，越界了
    // 固定一个起始行，往下走？好像也会越界 m <= 200
    mp := make(map[int]int)
    res, m := 0, len(grid)
    var match func(i int, visited map[int]bool) bool
    match = func(i int, visited map[int]bool) bool {
        if _, ok := visited[i]; ok { return false }
        visited[i] = true
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] != 1 { continue  }
            if _, ok := mp[j]; !ok || match(mp[j], visited) {
                mp[j] = i
                return true
            }
        }
        return false
    }
    for i := 0; i < m; i++ {
        visited := make(map[int]bool)
        if match(i, visited) {
            res++
        }
    }
    return res
}

func maximumInvitations1(grid [][]int) int {
    res, m, n := 0, len(grid), len(grid[0])
    myBoy := make([]int, n)
    for i := range myBoy { myBoy[i] = -1 }
    var findGirlForBoy func(boy int, visited []bool) bool
    findGirlForBoy = func(boy int, visited []bool) bool {
        for girl := 0; girl < n; girl++ {
            if !visited[girl] && grid[boy][girl] == 1 {
                visited[girl] = true
                if myBoy[girl] == -1 || findGirlForBoy(myBoy[girl], visited) {
                    myBoy[girl] = boy
                    return true
                }
            }
        }
        return false
    }
    for i := 0; i < m; i++ {
        visited := make([]bool, n)
        if findGirlForBoy(i, visited) { res++ }
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[1,1,1],
    //                [1,0,1],
    //                [0,0,1]]
    // Output: 3
    // Explanation: The invitations are sent as follows:
    // - The 1st boy invites the 2nd girl.
    // - The 2nd boy invites the 1st girl.
    // - The 3rd boy invites the 3rd girl.
    grid1 := [][]int{
        {1,1,1},
        {1,0,1},
        {0,0,1},
    }
    fmt.Println(maximumInvitations(grid1)) // 3
    // Example 2:
    // Input: grid = [[1,0,1,0],
    //                [1,0,0,0],
    //                [0,0,1,0],
    //                [1,1,1,0]]
    // Output: 3
    // Explanation: The invitations are sent as follows:
    // -The 1st boy invites the 3rd girl.
    // -The 2nd boy invites the 1st girl.
    // -The 3rd boy invites no one.
    // -The 4th boy invites the 2nd girl.
    grid2 := [][]int{
        {1,0,1,0},
        {1,0,0,0},
        {0,0,1,0},
        {1,1,1,0},
    }
    fmt.Println(maximumInvitations(grid2)) // 3

    fmt.Println(maximumInvitations1(grid1)) // 3
    fmt.Println(maximumInvitations1(grid2)) // 3
}