package main

// 2120. Execution of All Suffix Instructions Staying in a Grid
// There is an n x n grid, with the top-left cell at (0, 0) and the bottom-right cell at (n - 1, n - 1). 
// You are given the integer n and an integer array startPos where startPos = [startrow, startcol] indicates that a robot is initially at cell (startrow, startcol).

// You are also given a 0-indexed string s of length m where s[i] is the ith instruction for the robot: 
//     'L' (move left), 'R' (move right), 'U' (move up), and 'D' (move down).

// The robot can begin executing from any ith instruction in s. 
// It executes the instructions one by one towards the end of s but it stops if either of these conditions is met:
//     The next instruction will move the robot off the grid.
//     There are no more instructions left to execute.

// Return an array answer of length m where answer[i] is the number of instructions the robot can execute if the robot begins executing from the ith instruction in s.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/09/1.png" />
// Input: n = 3, startPos = [0,1], s = "RRDDLU"
// Output: [1,5,4,3,1,0]
// Explanation: Starting from startPos and beginning execution from the ith instruction:
// - 0th: "RRDDLU". Only one instruction "R" can be executed before it moves off the grid.
// - 1st:  "RDDLU". All five instructions can be executed while it stays in the grid and ends at (1, 1).
// - 2nd:   "DDLU". All four instructions can be executed while it stays in the grid and ends at (1, 0).
// - 3rd:    "DLU". All three instructions can be executed while it stays in the grid and ends at (0, 0).
// - 4th:     "LU". Only one instruction "L" can be executed before it moves off the grid.
// - 5th:      "U". If moving up, it would move off the grid.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/12/09/2.png" />
// Input: n = 2, startPos = [1,1], s = "LURD"
// Output: [4,1,0,0]
// Explanation:
// - 0th: "LURD".
// - 1st:  "URD".
// - 2nd:   "RD".
// - 3rd:    "D".

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/12/09/3.png" />
// Input: n = 1, startPos = [0,0], s = "LRUD"
// Output: [0,0,0,0]
// Explanation: No matter which instruction the robot begins execution from, it would move off the grid.
 
// Constraints:
//     m == s.length
//     1 <= n, m <= 500
//     startPos.length == 2
//     0 <= startrow, startcol < n
//     s consists of 'L', 'R', 'U', and 'D'.

import "fmt"

func executeInstructions(n int, startPos []int, s string) []int {
    res := make([]int,len(s))
    for i := 0; i < len(s);i++{
        count, x, y := 0, startPos[0], startPos[1]
        for j := i; j < len(s);j++{
            if s[j] == 'U' { // 上
                x -= 1
            } else if s[j] == 'D' { // 下
                x += 1
            } else if s[j] == 'R' { // 左
                y += 1
            } else if s[j] == 'L' { // 右
                y -= 1
            }
            if x < 0 || x >= n || y < 0 || y >= n { // 超出边界
                break
            } else {
                count += 1
            }
        }
        res[i] = count
    }
    return res
}

func executeInstructions1(n int, startPos []int, s string) []int {
    dirs := map[byte][]int{ 'L': {0, -1}, 'R': {0, 1}, 'U': {-1, 0}, 'D': {1, 0} }
    dx, dy := map[int]int{}, map[int]int{}
    x, y := startPos[0], startPos[1]
    for _, c := range s {
        d := dirs[byte(c)]
        x += d[0]
        y += d[1]
    }
    m := len(s)
    res := make([]int, m)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= m; i++ {
        dx[x], dy[y] = i, i
        d := dirs[s[m-i]]
        x -= d[0]
        y -= d[1]
        x1, y1 := x-startPos[0], y-startPos[1]
        t := max(max(dx[x1-1], dx[x1+n]), max(dy[y1-1], dy[y1+n]))
        res[m-i] = i - t
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/09/1.png" />
    // Input: n = 3, startPos = [0,1], s = "RRDDLU"
    // Output: [1,5,4,3,1,0]
    // Explanation: Starting from startPos and beginning execution from the ith instruction:
    // - 0th: "RRDDLU". Only one instruction "R" can be executed before it moves off the grid.
    // - 1st:  "RDDLU". All five instructions can be executed while it stays in the grid and ends at (1, 1).
    // - 2nd:   "DDLU". All four instructions can be executed while it stays in the grid and ends at (1, 0).
    // - 3rd:    "DLU". All three instructions can be executed while it stays in the grid and ends at (0, 0).
    // - 4th:     "LU". Only one instruction "L" can be executed before it moves off the grid.
    // - 5th:      "U". If moving up, it would move off the grid.
    fmt.Println(executeInstructions(3, []int{0,1}, "RRDDLU")) // [1,5,4,3,1,0]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/12/09/2.png" />
    // Input: n = 2, startPos = [1,1], s = "LURD"
    // Output: [4,1,0,0]
    // Explanation:
    // - 0th: "LURD".
    // - 1st:  "URD".
    // - 2nd:   "RD".
    // - 3rd:    "D".
    fmt.Println(executeInstructions(2, []int{1,1}, "LURD")) // [4,1,0,0]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/12/09/3.png" />
    // Input: n = 1, startPos = [0,0], s = "LRUD"
    // Output: [0,0,0,0]
    // Explanation: No matter which instruction the robot begins execution from, it would move off the grid.
    fmt.Println(executeInstructions(2, []int{0,0}, "LRUD")) // [0,0,0,0]

    fmt.Println(executeInstructions1(3, []int{0,1}, "RRDDLU")) // [1,5,4,3,1,0]
    fmt.Println(executeInstructions1(2, []int{1,1}, "LURD")) // [4,1,0,0]
    fmt.Println(executeInstructions1(2, []int{0,0}, "LRUD")) // [0,0,0,0]
}