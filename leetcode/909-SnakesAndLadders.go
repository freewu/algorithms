package main

// 909. Snakes and Ladders
// You are given an n x n integer matrix board where the cells are labeled from 1 to n2 in a Boustrophedon style starting from the bottom left of the board (i.e. board[n - 1][0]) and alternating direction each row.
// You start on square 1 of the board. In each move, starting from square curr, do the following:
//     Choose a destination square next with a label in the range [curr + 1, min(curr + 6, n2)].
//         This choice simulates the result of a standard 6-sided die roll: i.e., 
//         there are always at most 6 destinations, regardless of the size of the board.
//     If next has a snake or ladder, you must move to the destination of that snake or ladder. Otherwise, you move to next.
//     The game ends when you reach the square n2.

// A board square on row r and column c has a snake or ladder if board[r][c] != -1. 
// The destination of that snake or ladder is board[r][c]. Squares 1 and n2 do not have a snake or ladder.

// Note that you only take a snake or ladder at most once per move. 
// If the destination to a snake or ladder is the start of another snake or ladder, you do not follow the subsequent snake or ladder.
//     For example, suppose the board is [[-1,4],[-1,3]], and on the first move, your destination square is 2. 
//     You follow the ladder to square 3, but do not follow the subsequent ladder to 4.

// Return the least number of moves required to reach the square n2. If it is not possible to reach the square, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/09/23/snakes.png" />
// Input: board = [[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,35,-1,-1,13,-1],[-1,-1,-1,-1,-1,-1],[-1,15,-1,-1,-1,-1]]
// Output: 4
// Explanation: 
// In the beginning, you start at square 1 (at row 5, column 0).
// You decide to move to square 2 and must take the ladder to square 15.
// You then decide to move to square 17 and must take the snake to square 13.
// You then decide to move to square 14 and must take the ladder to square 35.
// You then decide to move to square 36, ending the game.
// This is the lowest possible number of moves to reach the last square, so return 4.

// Example 2:
// Input: board = [[-1,-1],[-1,3]]
// Output: 1

// Constraints:
//     n == board.length == board[i].length
//     2 <= n <= 20
//     board[i][j] is either -1 or in the range [1, n^2].
//     The squares labeled 1 and n2 do not have any ladders or snakes.

import "fmt"

func printMatrix(matrix [][]int) {
    for _,v := range matrix {
        fmt.Println(v)
    }
    fmt.Println()
}

// bfs
func snakesAndLadders(board [][]int) int {
    n := len(board)
    dp, move := make([]int, n * n + 1), make([]int, n * n + 1)
    for i:=1; i <= n*n; i++{
        dp[i] = n*n
    }
    //Init move
    count, cur := 0, 1
    for i := n-1; i >= 0; i-- {
        if count == 0{
            for count < n {
                move[cur] = board[i][count]
                cur++
                count++
            }
        } else {
            for count > 0 {
                count--
                move[cur] = board[i][count]
                cur++
            }
        }
    }
    // BFS Queue
    q := make([]int, 0)
    q = append(q, 1)
    dp[1] = 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for len(q) > 0 {
        i := q[0]
        q = q[1:]
        for j := i+1; j <= min(i+6, n*n); j++ {
            if move[j] > 0 {
                if dp[move[j]] > dp[i] + 1{
                    q = append(q, move[j])
                    dp[move[j]] = dp[i] + 1
                }
                continue
            }
            if dp[j] > dp[i] + 1 {
                q = append(q, j)
                dp[j] = dp[i] + 1
            }
            
        }
    }
    if dp[n*n] == n*n {
        return -1
    }
    return dp[n*n]
}

func snakesAndLadders1(board [][]int) int {
    n := len(board)
    if n <= 1{
        return n
    }
    visited, queue, res := make([]bool, n*n), []int{1}, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    iToIdx := func (i int, n int)(int, int){
        x := (n - 1) - (i - 1) / n
        y := (i - 1) % n
        if x % 2 == n % 2{
            y = n - 1 - y
        }
        return x, y
    }
    for len(queue) > 0 {
        tmp := []int{}
        res += 1  
        for _, i := range queue {
            skip := false
            if i == n*n {
                return res 
            }
            if visited[i-1] {
                continue
            }
            visited[i-1] = true
            mn, mx := i + 1, min(i + 6, n*n)
            for j := mn; j <= mx; j++{
                if skip {
                    tmp = append(tmp, j)
                } else {
                    x, y := iToIdx(j, n)
                    if board[x][y] != -1{
                        tmp = append(tmp, board[x][y])
                    }else{
                        tmp = append(tmp, j)
                    }
                }
                if tmp[len(tmp) - 1] == n * n {
                    return res
                }
            }
        }
        queue = tmp 
    }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/09/23/snakes.png" />
    // Input: board = [[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,35,-1,-1,13,-1],[-1,-1,-1,-1,-1,-1],[-1,15,-1,-1,-1,-1]]
    // Output: 4
    // Explanation: 
    // In the beginning, you start at square 1 (at row 5, column 0).
    // You decide to move to square 2 and must take the ladder to square 15.
    // You then decide to move to square 17 and must take the snake to square 13.
    // You then decide to move to square 14 and must take the ladder to square 35.
    // You then decide to move to square 36, ending the game.
    // This is the lowest possible number of moves to reach the last square, so return 4.
    board1 := [][]int{{-1,-1,-1,-1,-1,-1},{-1,-1,-1,-1,-1,-1},{-1,-1,-1,-1,-1,-1},{-1,35,-1,-1,13,-1},{-1,-1,-1,-1,-1,-1},{-1,15,-1,-1,-1,-1}}
    printMatrix(board1)
    fmt.Println(snakesAndLadders(board1)) // 4
    // Example 2:
    // Input: board = [[-1,-1],[-1,3]]
    // Output: 1
    board2 := [][]int{{-1,-1},{-1,3}}
    printMatrix(board2)
    fmt.Println(snakesAndLadders(board2)) // 1 

    fmt.Println(snakesAndLadders1(board1)) // 4
    fmt.Println(snakesAndLadders1(board2)) // 1 
}