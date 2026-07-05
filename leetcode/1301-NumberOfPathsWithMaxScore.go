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

func pathsWithMaxScore1(board []string) []int {
    n, mod := len(board), 1_000_000_007
    pre, now := make([][]int, n), make([][]int, n)
    for i := 0; i < n; i++ {
        now[i], pre[i] = make([]int, 2), []int{ -1, 0 }
    }
    now[n - 1][1] = 1
    getScore := func(ch byte) int { // 获取格子的数字得分
        if ch == 'E' || ch == 'S' { 
            return 0
        }
        return int(ch - '0')
    }
    for i := n - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if i == n - 1 && j == n - 1 { // 起点已经初始化过
                continue
            }
            // 障碍物不可达
            if board[i][j] == 'X' {
                now[j][0] = -1 // 标记为不可达
                continue
            }
            // 从三个方向收集候选
            down, right, downR := []int{}, []int{}, []int{}
            if i < n - 1 {
                down = pre[j]
            }
            if j < n - 1 {
                right = now[j+1]
            }
            if i < n - 1 && j < n - 1 {
                downR = pre[j + 1]
            }
            best, totalWays := -1, 0
            if len(down) != 0 {
                if down[0] > best {
                    best, totalWays = down[0], down[1]
                } else if down[0] == best {
                    totalWays += down[1]
                    totalWays %= mod
                }
            }
            if len(right) != 0 {
                if right[0] > best {
                    best, totalWays = right[0], right[1]
                } else if right[0] == best {
                    totalWays += right[1]
                    totalWays %= mod
                }
            }
            if len(downR) != 0 {
                if downR[0] > best {
                    best, totalWays = downR[0], downR[1]    
                } else if downR[0] == best {
                    totalWays += downR[1]
                    totalWays %= mod
                }
            }
            if best == -1 {
                now[j][0] = -1
                continue
            }
            now[j][0], now[j][1] = best + getScore(board[i][j]), totalWays
        }
        now, pre = pre, now
    }
    if pre[0][0] == -1 {
        return []int{0, 0}
    }
    return pre[0]
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

    fmt.Println(pathsWithMaxScore1([]string{"E23","2X2","12S"})) // [7,1]
    fmt.Println(pathsWithMaxScore1([]string{"E12","1X1","21S"})) //  [4,2]
    fmt.Println(pathsWithMaxScore1([]string{"E11","XXX","11S"})) //  [0,0]
}