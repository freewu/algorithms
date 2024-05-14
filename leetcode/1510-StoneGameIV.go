package main

// 1510. Stone Game IV
// Alice and Bob take turns playing a game, with Alice starting first.
// Initially, there are n stones in a pile. On each player's turn, t
// hat player makes a move consisting of removing any non-zero square number of stones in the pile.
// Also, if a player cannot make a move, he/she loses the game.
// Given a positive integer n, return true if and only if Alice wins the game otherwise return false, assuming both players play optimally.

// Example 1:
// Input: n = 1
// Output: true
// Explanation: Alice can remove 1 stone winning the game because Bob doesn't have any moves.

// Example 2:
// Input: n = 2
// Output: false
// Explanation: Alice can only remove 1 stone, after that Bob removes the last one winning the game (2 -> 1 -> 0).

// Example 3:
// Input: n = 4
// Output: true
// Explanation: n is already a perfect square, Alice can win with one move, removing 4 stones (4 -> 0).
 
// Constraints:
//     1 <= n <= 10^5

import "fmt"
import "math"

func winnerSquareGame(n int) bool {
    m := make(map[int]bool)
    m[0] = false
    var dfs func(m map[int]bool, n int) bool
    dfs = func(m map[int]bool, n int) bool {
        if ok := m[n]; ok {
            return m[n]
        }
        s := int(math.Sqrt(float64(n)))
        for i := 1; i <= s; i++ {
            if !dfs(m, n - i * i) { 
                m[n] = true
                return true
            }
        }
        m[n] = false
        return false
    }
    return dfs(m, n)
}

func winnerSquareGame1(n int) bool {
    f := make([]bool, n + 1)
    for i := range f {
        for j := int(math.Sqrt(float64(i))); j > 0; j-- {
            if !f[i - j * j] {
                f[i] = true
                break
            }
        }
    }
    return f[n]
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: true
    // Explanation: Alice can remove 1 stone winning the game because Bob doesn't have any moves.
    fmt.Println(winnerSquareGame(1)) // true
    // Example 2:
    // Input: n = 2
    // Output: false
    // Explanation: Alice can only remove 1 stone, after that Bob removes the last one winning the game (2 -> 1 -> 0).
    fmt.Println(winnerSquareGame(2)) // false
    // Example 3:
    // Input: n = 4
    // Output: true
    // Explanation: n is already a perfect square, Alice can win with one move, removing 4 stones (4 -> 0).
    fmt.Println(winnerSquareGame(4)) // true

    fmt.Println(winnerSquareGame1(1)) // true
    fmt.Println(winnerSquareGame1(2)) // false
    fmt.Println(winnerSquareGame1(4)) // true
}