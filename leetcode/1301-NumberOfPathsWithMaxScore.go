package main

// 1301. Number of Paths with Max Score
// You are given a square board of characters. 
// You can move on the board starting at the bottom right square marked with the character 'S'.

// You need to reach the top left square marked with the character 'E'. 
// The rest of the squares are labeled either with a numeric character 1, 2, ..., 9 or with an obstacle 'X'. 
// In one move you can go up, left or up-left (diagonally) only if there is no obstacle there.

// Return a list of two integers: 
// the first integer is the maximum sum of numeric characters you can collect, 
// and the second is the number of such paths that you can take to get that maximum sum, taken modulo 10^9 + 7.

// In case there is no path, return [0, 0].

// Example 1:
// Input: board = ["E23","2X2","12S"]
// Output: [7,1]

// Example 2:
// Input: board = ["E12","1X1","21S"]
// Output: [4,2]

// Example 3:
// Input: board = ["E11","XXX","11S"]
// Output: [0,0]

// Constraints:
//     2 <= board.length == board[i].length <= 100

import "fmt"

func pathsWithMaxScore(board []string) []int {
    n, mod := len(board), 1_000_000_007
    dp, count := make([][]int, n, n), make([][]int, n, n)
    for i := range dp {
        dp[i] = make([]int, n, n)
    }
    for i := range count {
        count[i] = make([]int, n, n)
    }
    dp[0][0], count[0][0] = 0, 1
    for i := 1; i < n; i++ {
        if board[i][0] == 'X' || dp[i-1][0] == -1 {
            dp[i][0], count[i][0] = -1, 0
        } else {
            dp[i][0] = dp[i-1][0] + int(board[i][0]-'0')
            count[i][0] = 1
        }
    }
    for i := 1; i < n; i++ {
        if board[0][i] == 'X' || dp[0][i-1] == -1 {
            dp[0][i] = -1
            count[0][i] = 0
        } else {
            dp[0][i] = dp[0][i-1] + int(board[0][i]-'0')
            count[0][i] = 1
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        for j := 1; j < n; j++ {
            if board[i][j] == 'X' {
                dp[i][j], count[i][j] = -1, 0
            } else {
                prevMax := max(dp[i-1][j-1], max(dp[i-1][j], dp[i][j-1]))
                if prevMax == -1 {
                    dp[i][j], count[i][j] = -1, 0
                } else {
                    dp[i][j] = prevMax + int(board[i][j]-'0')
                    if prevMax == dp[i-1][j] {
                        count[i][j] += count[i-1][j]
                        count[i][j] %= mod
                    }
                    if prevMax == dp[i-1][j-1] {
                        count[i][j] += count[i-1][j-1]
                        count[i][j] %= mod
                    }
                    if prevMax == dp[i][j-1] {
                        count[i][j] += count[i][j-1]
                        count[i][j] %= mod
                    }
                }
            }
        }
    }
    if dp[n-1][n-1] == -1 {
        return []int{0, 0}
    }
    return []int{ dp[n-1][n-1] - int('S'-'0'), count[n-1][n-1]}
}

func main() {
    // Example 1:
    // Input: board = ["E23","2X2","12S"]
    // Output: [7,1]
    fmt.Println(pathsWithMaxScore([]string{"E23","2X2","12S"})) // [7,1]
    // Example 2:
    // Input: board = ["E12","1X1","21S"]
    // Output: [4,2]
    fmt.Println(pathsWithMaxScore([]string{"E12","1X1","21S"})) //  [4,2]
    // Example 3:
    // Input: board = ["E11","XXX","11S"]
    // Output: [0,0]
    fmt.Println(pathsWithMaxScore([]string{"E11","XXX","11S"})) //  [0,0]
}