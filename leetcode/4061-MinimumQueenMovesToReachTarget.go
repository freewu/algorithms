package main

// 4061. Minimum Queen Moves to Reach Target
// There is an 8 x 8 empty chessboard with 1-indexed rows and columns.

// You are given an array source = [sr, sc] representing the starting position of a queen, 
// and an array target = [tr, tc] representing the target position.

// In one move, the queen travels one or more squares along a single row, column, or diagonal, staying within the board.

// Return the minimum number of moves for the queen to land exactly on target.

// Example 1:
// Input: source = [8,1], target = [1,8]
// Output: 1
// Explanation:
// ​​​​​​​​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/08/31/111.png" />
// A single diagonal move takes the queen straight from (8, 1) to (1, 8).

// Example 2:
// Input: source = [4,2], target = [1,3]
// Output: 2
// Explanation:
// ​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/09/14/1e602ed4-c525-4be4-a6a9-6804cb7d2a55.png" />
// The queen moves from (4, 2) to (4, 3), then from (4, 3) to (1, 3), reaching the target in 2 moves.

// Example 3:
// Input: source = [1,1], target = [1,1]
// Output: 0
// Explanation:
// The queen is already at the target position, so no moves are needed.

// Constraints:​​​​​​​
//     source == [sr, sc]
//     target == [tr, tc]
//     1 <= sr, sc, tr, tc <= 8

import "fmt"

func minQueenMoves(source []int, target []int) int {
    sr, sc, tr, tc := source[0], source[1], target[0], target[1]
    if sr == tr && sc == tc {
        return 0
    }
    if sr == tr || sc == tc || sr + sc == tr + tc || sr - sc == tr - tc {
        return 1
    }
    return 2
}

func main() {
    // Example 1:
    // Input: source = [8,1], target = [1,8]
    // Output: 1
    // Explanation:
    // ​​​​​​​​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/08/31/111.png" />
    // A single diagonal move takes the queen straight from (8, 1) to (1, 8).
    fmt.Println(minQueenMoves([]int{8,1}, []int{1,8})) // 1
    // Example 2:
    // Input: source = [4,2], target = [1,3]
    // Output: 2
    // Explanation:
    // ​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/09/14/1e602ed4-c525-4be4-a6a9-6804cb7d2a55.png" />
    // The queen moves from (4, 2) to (4, 3), then from (4, 3) to (1, 3), reaching the target in 2 moves.
    fmt.Println(minQueenMoves([]int{4,2}, []int{1,3})) // 2
    // Example 3:
    // Input: source = [1,1], target = [1,1]
    // Output: 0
    // Explanation:
    // The queen is already at the target position, so no moves are needed.
    fmt.Println(minQueenMoves([]int{1,1}, []int{1,1})) // 0

    fmt.Println(minQueenMoves([]int{8,8}, []int{8,8})) // 0
    fmt.Println(minQueenMoves([]int{1,1}, []int{8,8})) // 1
    fmt.Println(minQueenMoves([]int{8,8}, []int{1,1})) // 1
    fmt.Println(minQueenMoves([]int{1,8}, []int{1,8})) // 0
    fmt.Println(minQueenMoves([]int{8,1}, []int{1,8})) // 1
    fmt.Println(minQueenMoves([]int{8,1}, []int{8,1})) // 0
    fmt.Println(minQueenMoves([]int{1,8}, []int{8,1})) // 1
}