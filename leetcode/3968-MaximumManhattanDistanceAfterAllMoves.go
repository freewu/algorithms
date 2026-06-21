package main

// 3968. Maximum Manhattan Distance After All Moves
// You are given a string moves consisting of the characters 'U', 'D', 'L', 'R', and '_'.

// Starting from the origin (0, 0), each character represents one move on a 2D plane:
//     1. 'U': Move up by 1 unit.
//     2. 'D': Move down by 1 unit.
//     3. 'L': Move left by 1 unit.
//     4. 'R': Move right by 1 unit.
//     5. '_': Can be independently replaced with any one of 'U', 'D', 'L', or 'R'.

// Return the maximum Manhattan distance from the origin that can be achieved after all moves have been performed.

// Example 1:
// Input: moves = "L_D_"
// Output: 4
// Explanation:
// One optimal choice is:
// 'L': (0, 0) -> (-1, 0)
// '_' treated as 'D': (-1, 0) -> (-1, -1)
// 'D': (-1, -1) -> (-1, -2)
// '_' treated as 'L': (-1, -2) -> (-2, -2)
// The final Manhattan distance from the origin is |0 - (-2)| + |0 - (-2)| = 4.

// Example 2:
// Input: moves = "U_R"
// Output: 3
// Explanation:
// One optimal choice is:
// 'U': (0, 0) -> (0, 1)
// '_' treated as 'U': (0, 1) -> (0, 2)
// 'R': (0, 2) -> (1, 2)
// The final Manhattan distance from the origin is |0 - 1| + |0 - 2| = 3.

// Constraints:
//     1 <= moves.length <= 10^5
//     moves consists of only 'U', 'D', 'L', 'R', and '_'.

import "fmt"

func maxDistance(moves string) int {
    x, y, unknown := 0, 0, 0
    for i := 0; i < len(moves); i++ {
        switch moves[i] {
        case 'L':
            x--
        case 'D':
            y--
        case 'R':
            x++
        case 'U':
            y++
        default:
            unknown++
        }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    return abs(x) + abs(y) + unknown
}

func main() {
    // Example 1:
    // Input: moves = "L_D_"
    // Output: 4
    // Explanation:
    // One optimal choice is:
    // 'L': (0, 0) -> (-1, 0)
    // '_' treated as 'D': (-1, 0) -> (-1, -1)
    // 'D': (-1, -1) -> (-1, -2)
    // '_' treated as 'L': (-1, -2) -> (-2, -2)
    // The final Manhattan distance from the origin is |0 - (-2)| + |0 - (-2)| = 4.
    fmt.Println(maxDistance("L_D_")) // 4
    // Example 2:
    // Input: moves = "U_R"
    // Output: 3
    // Explanation:
    // One optimal choice is:
    // 'U': (0, 0) -> (0, 1)
    // '_' treated as 'U': (0, 1) -> (0, 2)
    // 'R': (0, 2) -> (1, 2)
    // The final Manhattan distance from the origin is |0 - 1| + |0 - 2| = 3.
    fmt.Println(maxDistance("U_R")) // 3
}