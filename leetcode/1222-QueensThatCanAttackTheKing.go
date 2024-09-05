package main

// 1222. Queens That Can Attack the King
// On a 0-indexed 8 x 8 chessboard, there can be multiple black queens and one white king.

// You are given a 2D integer array queens where queens[i] = [xQueeni, yQueeni] represents the position of the ith black queen on the chessboard. 
// You are also given an integer array king of length 2 where king = [xKing, yKing] represents the position of the white king.

// Return the coordinates of the black queens that can directly attack the king. 
// You may return the answer in any order.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/12/21/chess1.jpg" />
// Input: queens = [[0,1],[1,0],[4,0],[0,4],[3,3],[2,4]], king = [0,0]
// Output: [[0,1],[1,0],[3,3]]
// Explanation: The diagram above shows the three queens that can directly attack the king and the three queens that cannot attack the king (i.e., marked with red dashes).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/12/21/chess2.jpg" />
// Input: queens = [[0,0],[1,1],[2,2],[3,4],[3,5],[4,4],[4,5]], king = [3,3]
// Output: [[2,2],[3,4],[4,4]]
// Explanation: The diagram above shows the three queens that can directly attack the king and the three queens that cannot attack the king (i.e., marked with red dashes).

// Constraints:
//     1 <= queens.length < 64
//     queens[i].length == king.length == 2
//     0 <= xQueeni, yQueeni, xKing, yKing < 8
//     All the given positions are unique.

import "fmt"

func queensAttacktheKing(queens [][]int, king []int) [][]int {
    directions := [][]int{ {0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}, } // 8个方向
    res := [][]int{}
    for _, dir := range directions {
        cur, hadCapture := []int{king[0], king[1]}, false
        for cur[0] >= 0 && cur[0] <= 8 && cur[1] >= 0 && cur[1] <= 8 {
            cur[0] += dir[0]
            cur[1] += dir[1]
            for _, q := range queens {
                if q[0] == cur[0] && q[1] == cur[1] { // 可以到达 king的位置
                    res = append(res, q)
                    hadCapture = true
                    break
                }
            }
            if hadCapture { break } // 换个方向
        }
    }
    return res
}

func queensAttacktheKing1(queens [][]int, king []int) [][]int {
    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
    res, grid := [][]int{}, make([][]int, 8)
    for i := range grid {
        grid[i] = make([]int, 8)
    }
    for _, ch := range queens {
        grid[ch[0]][ch[1]] = 1
    }
    inBoard := func(grid [][]int, x, y int) bool { // 边界检测
        m, n := len(grid), len(grid[0])
        if x < 0 || y < 0 ||  x >= m || y >= n { return false }
        return true
    }
    for _, dir := range directions {
        depth := 1
        for {
            dx, dy := dir[0] * depth, dir[1] * depth
            x, y := king[0] + dx, king[1] + dy
            if !inBoard(grid, x, y) { break }
            if grid[x][y] == 1 {
                res = append(res, []int{x, y})
                break
            }
            depth++
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/12/21/chess1.jpg" />
    // Input: queens = [[0,1],[1,0],[4,0],[0,4],[3,3],[2,4]], king = [0,0]
    // Output: [[0,1],[1,0],[3,3]]
    // Explanation: The diagram above shows the three queens that can directly attack the king and the three queens that cannot attack the king (i.e., marked with red dashes).
    fmt.Println(queensAttacktheKing([][]int{{0,1},{1,0},{4,0},{0,4},{3,3},{2,4}}, []int{0,0})) // [[0,1],[1,0],[3,3]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/12/21/chess2.jpg" />
    // Input: queens = [[0,0],[1,1],[2,2],[3,4],[3,5],[4,4],[4,5]], king = [3,3]
    // Output: [[2,2],[3,4],[4,4]]
    // Explanation: The diagram above shows the three queens that can directly attack the king and the three queens that cannot attack the king (i.e., marked with red dashes).
    fmt.Println(queensAttacktheKing([][]int{{0,0},{1,1},{2,2},{3,4},{3,5},{4,4},{4,5}}, []int{3,3})) // [[2,2],[3,4],[4,4]]

    fmt.Println(queensAttacktheKing1([][]int{{0,1},{1,0},{4,0},{0,4},{3,3},{2,4}}, []int{0,0})) // [[0,1],[1,0],[3,3]]
    fmt.Println(queensAttacktheKing1([][]int{{0,0},{1,1},{2,2},{3,4},{3,5},{4,4},{4,5}}, []int{3,3})) // [[2,2],[3,4],[4,4]]
}