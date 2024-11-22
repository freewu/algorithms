package main

// 2814. Minimum Time Takes to Reach Destination Without Drowning
// You are given an n * m 0-indexed grid of string land. 
// Right now, you are standing at the cell that contains "S", and you want to get to the cell containing "D". 
// There are three other types of cells in this land:
//     ".": These cells are empty.
//     "X": These cells are stone.
//     "*": These cells are flooded.

// At each second, you can move to a cell that shares a side with your current cell (if it exists). 
// Also, at each second, every empty cell that shares a side with a flooded cell becomes flooded as well.

// There are two problems ahead of your journey:
//     1. You can't step on stone cells.
//     2. You can't step on flooded cells since you will drown (also, you can't step on a cell that will be flooded at the same time as you step on it).

// Return the minimum time it takes you to reach the destination in seconds, or -1 if it is impossible.

// Note that the destination will never be flooded.

// Example 1:
// Input: land = [["D",".","*"],[".",".","."],[".","S","."]]
// Output: 3
// Explanation: The picture below shows the simulation of the land second by second. 
// The blue cells are flooded, and the gray cells are stone.
// Picture (0) shows the initial state and picture (3) shows the final state when we reach destination. 
// As you see, it takes us 3 second to reach destination and the answer would be 3.
// It can be shown that 3 is the minimum time needed to reach from S to D.
// <img src="https://assets.leetcode.com/uploads/2023/08/09/ex1.png" />

// Example 2:
// Input: land = [["D","X","*"],[".",".","."],[".",".","S"]]
// Output: -1
// Explanation: The picture below shows the simulation of the land second by second. 
// The blue cells are flooded, and the gray cells are stone.
// Picture (0) shows the initial state. As you see, no matter which paths we choose, we will drown at the 3rd second. 
// Also the minimum path takes us 4 seconds to reach from S to D.
// So the answer would be -1.
// <img src="https://assets.leetcode.com/uploads/2023/08/09/ex2-2.png" /> 

// Example 3:
// Input: land = [["D",".",".",".","*","."],[".","X",".","X",".","."],[".",".",".",".","S","."]]
// Output: 6
// Explanation: It can be shown that we can reach destination in 6 seconds.
// Also it can be shown that 6 is the minimum seconds one need to reach from S to D.

// Constraints:
//     2 <= n, m <= 100
//     land consists only of "S", "D", ".", "*" and "X".
//     Exactly one of the cells is equal to "S".
//     Exactly one of the cells is equal to "D".

import "fmt"

func minimumSeconds(land [][]string) int {
    m, n, si, sj := len(land), len(land[0]), 0 ,0
    visited, queue := make([][]bool, m), [][2]int{}
    graph := make([][]int, m)
    for i, row := range land {
        visited[i] = make([]bool, n)
        graph[i] = make([]int, n)
        for j := range graph[i] {
            graph[i][j] = 1 << 31
        }
        for j, c := range row {
            if c == "*" {// flooded
                queue = append(queue, [2]int{i, j})
            } else if c == "S" {
                si, sj = i, j
            }
        }
    }
    directions := [5]int{-1, 0, 1, 0, -1}
    for t := 0; len(queue) > 0; t++ {
        for k := len(queue); k > 0; k-- {
            p := queue[0]
            queue = queue[1:]
            i, j := p[0], p[1]
            graph[i][j] = t
            for d := 0; d < 4; d++ {
                x, y := i + directions[d], j + directions[d + 1]
                if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] { // border check
                    empty, start := (land[x][y] == "."), (land[x][y] == "S")
                    if empty || start {
                        visited[x][y] = true
                        queue = append(queue, [2]int{x, y})
                    }
                }
            }
        }
    }
    queue = append(queue, [2]int{si, sj})
    visited = make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    visited[si][sj] = true
    for t := 0; len(queue) > 0; t++ {
        for k := len(queue); k > 0; k-- {
            p := queue[0]
            queue = queue[1:]
            i, j := p[0], p[1]
            if land[i][j] == "D" { return t }
            for d := 0; d < 4; d++ {
                x, y := i + directions[d], j + directions[d+1]
                if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] && graph[x][y] > t+1 {
                    empty, dest := (land[x][y] == "."), (land[x][y] == "D")
                    if empty || dest {
                        visited[x][y] = true
                        queue = append(queue, [2]int{x, y})
                    }
                }
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: land = [["D",".","*"],[".",".","."],[".","S","."]]
    // Output: 3
    // Explanation: The picture below shows the simulation of the land second by second. 
    // The blue cells are flooded, and the gray cells are stone.
    // Picture (0) shows the initial state and picture (3) shows the final state when we reach destination. 
    // As you see, it takes us 3 second to reach destination and the answer would be 3.
    // It can be shown that 3 is the minimum time needed to reach from S to D.
    // <img src="https://assets.leetcode.com/uploads/2023/08/09/ex1.png" />
    fmt.Println(minimumSeconds([][]string{{"D",".","*"},{".",".","."},{".","S","."}})) // 3
    // Example 2:
    // Input: land = [["D","X","*"],[".",".","."],[".",".","S"]]
    // Output: -1
    // Explanation: The picture below shows the simulation of the land second by second. 
    // The blue cells are flooded, and the gray cells are stone.
    // Picture (0) shows the initial state. As you see, no matter which paths we choose, we will drown at the 3rd second. 
    // Also the minimum path takes us 4 seconds to reach from S to D.
    // So the answer would be -1.
    // <img src="https://assets.leetcode.com/uploads/2023/08/09/ex2-2.png" /> 
    fmt.Println(minimumSeconds([][]string{{"D","X","*"},{".",".","."},{".",".","S"}})) // -1
    // Example 3:
    // Input: land = [["D",".",".",".","*","."],[".","X",".","X",".","."],[".",".",".",".","S","."]]
    // Output: 6
    // Explanation: It can be shown that we can reach destination in 6 seconds.
    // Also it can be shown that 6 is the minimum seconds one need to reach from S to D.
    fmt.Println(minimumSeconds([][]string{{"D",".",".",".","*","."},{".","X",".","X",".","."},{".",".",".",".","S","."}})) // 6
}