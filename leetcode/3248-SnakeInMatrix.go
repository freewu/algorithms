package main

// 3248. Snake in Matrix
// There is a snake in an n x n matrix grid and can move in four possible directions. 
// Each cell in the grid is identified by the position: grid[i][j] = (i * n) + j.

// The snake starts at cell 0 and follows a sequence of commands.

// You are given an integer n representing the size of the grid and an array of strings commands 
// where each command[i] is either "UP", "RIGHT", "DOWN", and "LEFT". 
// It's guaranteed that the snake will remain within the grid boundaries throughout its movement.

// Return the position of the final cell where the snake ends up after executing commands.

// Example 1:
// Input: n = 2, commands = ["RIGHT","DOWN"]
// Output: 3
// Explanation:
// [0] 1 | 0 [1] | 0  1 
//  2  3 | 2  3  | 2 [3]

// Example 2:
// Input: n = 3, commands = ["DOWN","RIGHT","UP"]
// Output: 1
// Explanation:
// [0] 1  2 |  0  1  2 | 0  1  2 | 0 [1] 2
//  3  4  5 | [3] 4  5 | 3 [4] 5 | 3  4  5
//  6  7  8 |  6  7  8 | 6  7  8 | 6  7  8

// Constraints:
//     2 <= n <= 10
//     1 <= commands.length <= 100
//     commands consists only of "UP", "RIGHT", "DOWN", and "LEFT".
//     The input is generated such the snake will not move outside of the boundaries.

import "fmt"

func finalPositionOfSnake(n int, commands []string) int {
    row, col := 0, 0
    for _, command := range commands {
        switch command {
            case "RIGHT": col++
            case "LEFT":  col--
            case "DOWN":  row++
            case "UP":    row--
        }
    }
    return row * n + col
}

func finalPositionOfSnake1(n int, commands []string) int {
    res := 0
    for _, command := range commands {
        switch command {
            case "RIGHT": res++
            case "LEFT":  res--
            case "DOWN":  res += n
            case "UP":    res -= n
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2, commands = ["RIGHT","DOWN"]
    // Output: 3
    // Explanation:
    // [0] 1 | 0 [1] | 0  1 
    //  2  3 | 2  3  | 2 [3]
    fmt.Println(finalPositionOfSnake(2, []string{"RIGHT","DOWN"})) // 3
    // Example 2:
    // Input: n = 3, commands = ["DOWN","RIGHT","UP"]
    // Output: 1
    // Explanation:
    // [0] 1  2 |  0  1  2 | 0  1  2 | 0 [1] 2
    //  3  4  5 | [3] 4  5 | 3 [4] 5 | 3  4  5
    //  6  7  8 |  6  7  8 | 6  7  8 | 6  7  8
    fmt.Println(finalPositionOfSnake(3, []string{"DOWN","RIGHT","UP"})) // 1

    fmt.Println(finalPositionOfSnake1(2, []string{"RIGHT","DOWN"})) // 3
    fmt.Println(finalPositionOfSnake1(3, []string{"DOWN","RIGHT","UP"})) // 1
}