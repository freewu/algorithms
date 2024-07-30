package main

// 773. Sliding Puzzle
// On an 2 x 3 board, there are five tiles labeled from 1 to 5, and an empty square represented by 0. 
// A move consists of choosing 0 and a 4-directionally adjacent number and swapping it.

// The state of the board is solved if and only if the board is [[1,2,3],[4,5,0]].

// Given the puzzle board board, return the least number of moves required so that the state of the board is solved. 
// If it is impossible for the state of the board to be solved, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/29/slide1-grid.jpg" />
// Input: board = [[1,2,3],[4,0,5]]
// Output: 1
// Explanation: Swap the 0 and the 5 in one move.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/29/slide2-grid.jpg" />
// Input: board = [[1,2,3],[5,4,0]]
// Output: -1
// Explanation: No number of moves will make the board solved.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/06/29/slide3-grid.jpg" />
// Input: board = [[4,1,2],[5,0,3]]
// Output: 5
// Explanation: 5 is the smallest number of moves that solves the board.
// An example path:
// After move 0: [[4,1,2],[5,0,3]]
// After move 1: [[4,1,2],[0,5,3]]
// After move 2: [[0,1,2],[4,5,3]]
// After move 3: [[1,0,2],[4,5,3]]
// After move 4: [[1,2,0],[4,5,3]]
// After move 5: [[1,2,3],[4,5,0]]
 
// Constraints:
//     board.length == 2
//     board[i].length == 3
//     0 <= board[i][j] <= 5
//     Each value board[i][j] is unique.

import "fmt"

// bfs
func slidingPuzzle(board [][]int) int {
    type state struct {
        board [2][3]uint8
        i, j  int8
    }
    init := state{}
    for i := range board {
        for j, v := range board[i] {
            init.board[i][j] = uint8(v)
            if v == 0 {
                init.i = int8(i)
                init.j = int8(j)
            }
        }
    }
    directions := [][2]int8{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
    done := [2][3]uint8{{1, 2, 3}, {4, 5, 0}}
    curr := []state{init}
    next := []state{}
    seen := make(map[state]struct{})
    seen[curr[0]] = struct{}{}
    for steps := 0; len(curr) > 0; steps++ {
        next = next[:0]
        for _, x := range curr {
            if x.board == done {
                return steps
            }
            for _, d := range directions {
                ii, jj := x.i+d[0], x.j+d[1]
                if ii < 0 || jj < 0 || ii >= 2 || jj >= 3 {
                    continue
                }
                // Copy board and swap tiles
                s := x
                s.board[ii][jj], s.board[x.i][x.j] = s.board[x.i][x.j], s.board[ii][jj]
                s.i = ii
                s.j = jj
                if _, exists := seen[s]; exists {
                    continue
                }
                seen[s] = struct{}{}
                next = append(next, s)
            }
        }
        curr, next = next, curr
    }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/29/slide1-grid.jpg" />
    // Input: board = [[1,2,3],[4,0,5]]
    // Output: 1
    // Explanation: Swap the 0 and the 5 in one move.
    fmt.Println(slidingPuzzle([][]int{{1,2,3},{4,0,5}})) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/29/slide2-grid.jpg" />
    // Input: board = [[1,2,3],[5,4,0]]
    // Output: -1
    // Explanation: No number of moves will make the board solved.
    fmt.Println(slidingPuzzle([][]int{{1,2,3},{5,4,0}})) // -1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/06/29/slide3-grid.jpg" />
    // Input: board = [[4,1,2],[5,0,3]]
    // Output: 5
    // Explanation: 5 is the smallest number of moves that solves the board.
    // An example path:
    // After move 0: [[4,1,2],[5,0,3]]
    // After move 1: [[4,1,2],[0,5,3]]
    // After move 2: [[0,1,2],[4,5,3]]
    // After move 3: [[1,0,2],[4,5,3]]
    // After move 4: [[1,2,0],[4,5,3]]
    // After move 5: [[1,2,3],[4,5,0]]
    fmt.Println(slidingPuzzle([][]int{{4,1,2},{5,0,3}})) // 5
}