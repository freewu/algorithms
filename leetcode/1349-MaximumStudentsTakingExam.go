package main

// 1349. Maximum Students Taking Exam
// Given a m * n matrix seats  that represent seats distributions in a classroom. 
// If a seat is broken, it is denoted by '#' character otherwise it is denoted by a '.' character.

// Students can see the answers of those sitting next to the left, right, upper left and upper right, 
// but he cannot see the answers of the student sitting directly in front or behind him. 
// Return the maximum number of students that can take the exam together without any cheating being possible.

// Students must be placed in seats in good condition.

// Example 1:
// Input: seats = [["#",".","#","#",".","#"],
//                 [".","#","#","#","#","."],
//                 ["#",".","#","#",".","#"]]
// Output: 4
// Explanation: Teacher can place 4 students in available seats so they don't cheat on the exam. 

// Example 2:
// Input: seats = [[".","#"],
//                 ["#","#"],
//                 ["#","."],
//                 ["#","#"],
//                 [".","#"]]
// Output: 3
// Explanation: Place all students in available seats. 

// Example 3:
// Input: seats = [["#",".",".",".","#"],
//                 [".","#",".","#","."],
//                 [".",".","#",".","."],
//                 [".","#",".","#","."],
//                 ["#",".",".",".","#"]]
// Output: 10
// Explanation: Place students in available seats in column 1, 3 and 5.

// Constraints:
//     seats contains only characters '.' and'#'.
//     m == seats.length
//     n == seats[i].length
//     1 <= m <= 8
//     1 <= n <= 8

import "fmt"
import "math/bits"

func maxStudents(seats [][]byte) int {
    m := len(seats[0])
    dp := make([]int, 1 << m)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, row := range seats {
        prevDP := dp
        dp = make([]int, 1 << m)
        rowMask := 0
        for i := range row {
            if row[i] =='.' {
                rowMask += 1 << i
            }
        }
        for mask := 0; mask < 1 << m; mask++ {
            if mask&rowMask != mask { continue }
            count := bits.OnesCount(uint(mask))
            for prevMask := 0; prevMask < 1 << m; prevMask++ {
                if mask << 1 & mask > 0 || prevMask << 1 & mask > 0 || prevMask >> 1 & mask > 0 { continue }
                dp[mask] = max(dp[mask], prevDP[prevMask] + count)
            }
        }
    }
    res := 0
    for i := 0; i < 1 << m; i++ {
        res = max(res, dp[i])
    }
    return res
}

func maxStudents1(seats [][]byte) int {
    m, n := len(seats), len(seats[0])
    validity := make([]int, m+1)
    for i := 0; i < m; i++ {
        rowvalid := 0
        for j := 0; j < n; j++ {
            if seats[i][j] == '.' {
                rowvalid += 1 << j
            }
        }
        validity[i+1] = rowvalid
    }
    dp := make([][]int, m + 1)
    for i := range dp {
        dp[i] = make([]int, (1<<n)+1)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    dp[0][0] = 0
    for i := 1; i <= m; i++ {
        valid := validity[i]
        for j := 0; j < (1 << n); j++ {
            if (j & valid) != j {
                continue
            }
            if (j & (j >> 1)) != 0 {
                continue
            }
            for k := 0; k < (1 << n); k++ {
                if dp[i-1][k] == -1 {
                    continue
                }
                if (j & (k >> 1)) != 0 || (k & (j >> 1)) != 0 {
                    continue
                }
                if dp[i][j] < dp[i-1][k] + int(bits.OnesCount32(uint32(j))) {
                    dp[i][j] = dp[i-1][k] + int(bits.OnesCount32(uint32(j)))
                }
            }
        }
    }
    res := -1
    for _, v := range dp[m] {
        if v > res {
            res = v
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: seats = [["#",".","#","#",".","#"],
    //                 [".","#","#","#","#","."],
    //                 ["#",".","#","#",".","#"]]
    // Output: 4
    // Explanation: Teacher can place 4 students in available seats so they don't cheat on the exam. 
    seats1 := [][]byte{
        {'#','.','#','#','.','#'},
        {'.','#','#','#','#','.'},
        {'#','.','#','#','.','#'},
    }
    fmt.Println(maxStudents(seats1)) // 4
    // Example 2:
    // Input: seats = [[".","#"],
    //                 ["#","#"],
    //                 ["#","."],
    //                 ["#","#"],
    //                 [".","#"]]
    // Output: 3
    // Explanation: Place all students in available seats. 
    seats2 := [][]byte{
        {'.','#'},
        {'#','#'},
        {'#','.'},
        {'#','#'},
        {'.','#'},
    }
    fmt.Println(maxStudents(seats2)) // 3
    // Example 3:
    // Input: seats = [["#",".",".",".","#"],
    //                 [".","#",".","#","."],
    //                 [".",".","#",".","."],
    //                 [".","#",".","#","."],
    //                 ["#",".",".",".","#"]]
    // Output: 10
    // Explanation: Place students in available seats in column 1, 3 and 5.
    seats3 := [][]byte{
        {'#','.','.','.','#'},
        {'.','#','.','#','.'},
        {'.','.','#','.','.'},
        {'.','#','.','#','.'},
        {'#','.','.','.','#'},
    }
    fmt.Println(maxStudents(seats3)) // 10

    fmt.Println(maxStudents1(seats1)) // 4
    fmt.Println(maxStudents1(seats2)) // 3
    fmt.Println(maxStudents1(seats3)) // 10
}