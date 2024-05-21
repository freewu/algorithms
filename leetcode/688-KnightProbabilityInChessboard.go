package main

// 688. Knight Probability in Chessboard
// On an n x n chessboard, a knight starts at the cell (row, column) and attempts to make exactly k moves.
// The rows and columns are 0-indexed, so the top-left cell is (0, 0), and the bottom-right cell is (n - 1, n - 1).

// A chess knight has eight possible moves it can make, as illustrated below. 
// Each move is two cells in a cardinal direction, then one cell in an orthogonal direction.

// <img src="https://assets.leetcode.com/uploads/2018/10/12/knight.png" />

// Each time the knight is to move, it chooses one of eight possible moves uniformly at random (even if the piece would go off the chessboard) and moves there.
// The knight continues moving until it has made exactly k moves or has moved off the chessboard.
// Return the probability that the knight remains on the board after it has stopped moving.

// Example 1:
// Input: n = 3, k = 2, row = 0, column = 0
// Output: 0.06250
// Explanation: There are two moves (to (1,2), (2,1)) that will keep the knight on the board.
// From each of those positions, there are also two moves that will keep the knight on the board.
// The total probability the knight stays on the board is 0.0625.

// Example 2:
// Input: n = 1, k = 0, row = 0, column = 0
// Output: 1.00000
 
// Constraints:
//     1 <= n <= 25
//     0 <= k <= 100
//     0 <= row, column <= n - 1

import "fmt"

func knightProbability(n int, k int, row int, column int) float64 {    
    if k == 0 { return 1.0 }
    prob, curr_dp, prev_dp := float64(0), make([][]float32, n), make([][]float32, n)
    for i := 0; i < n; i++ {
        curr_dp[i] = make([]float32, n)
        prev_dp[i] = make([]float32, n)
    }
    prev_dp[row][column] = 1
    moves := [][2]int8 { // 向88个方向的走
        {-2, -1}, {-2, 1},
        {-1, -2}, {-1, 2},
        { 2, -1}, { 2, 1},
        { 1, -2}, { 1, 2},
    }
    for ; k > 0; k-- { // 走 k 步
        for i := byte(0); i < byte(n); i++ {
            for j := byte(0); j < byte(n); j++ {
                curr_dp[i][j] = 0 
                for _, move := range moves { // 向 8 个方向
                    x, y := i + byte(move[0]), j + byte(move[1])
                    if x >= byte(n) || y >= byte(n) {
                        continue
                    }
                    curr_dp[i][j] += prev_dp[x][y] / 8
                } 
                if k == 1 {
                    prob += float64(curr_dp[i][j])
                }
            }
        }
        curr_dp, prev_dp = prev_dp, curr_dp
    }
    return prob
}

func main() {
    // Example 1:
    // Input: n = 3, k = 2, row = 0, column = 0
    // Output: 0.06250
    // Explanation: There are two moves (to (1,2), (2,1)) that will keep the knight on the board.
    // From each of those positions, there are also two moves that will keep the knight on the board.
    // The total probability the knight stays on the board is 0.0625.
    fmt.Println(knightProbability(3,2,0,0)) // 0.06250
    // Example 2:
    // Input: n = 1, k = 0, row = 0, column = 0
    // Output: 1.00000
    fmt.Println(knightProbability(1,0,0,0)) // 1.00000
}