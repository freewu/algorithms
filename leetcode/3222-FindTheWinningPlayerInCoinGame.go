package main

// 3222. Find the Winning Player in Coin Game
// You are given two positive integers x and y, denoting the number of coins with values 75 and 10 respectively.

// Alice and Bob are playing a game. 
// Each turn, starting with Alice, the player must pick up coins with a total value 115. 
// If the player is unable to do so, they lose the game.

// Return the name of the player who wins the game if both players play optimally.

// Example 1:
// Input: x = 2, y = 7
// Output: "Alice"
// Explanation:
// The game ends in a single turn:
// Alice picks 1 coin with a value of 75 and 4 coins with a value of 10.

// Example 2:
// Input: x = 4, y = 11
// Output: "Bob"
// Explanation:
// The game ends in 2 turns:
// Alice picks 1 coin with a value of 75 and 4 coins with a value of 10.
// Bob picks 1 coin with a value of 75 and 4 coins with a value of 10.

// Constraints:
//     1 <= x, y <= 100

import "fmt"

func losingPlayer(x int, y int) string {
    n := 0
    for x > 0 { // x 和 y ，分别表示价值为 75 和 10 
        // 每次拿走  1 枚价值为 75 的硬币和 4 枚价值为 10 的硬币 75 + 4 * 10 == 115
        x--
        y -= 4
        if x >= 0 && y >= 0 {
            n++
        }
    }
    if n % 2 == 0 {
        return "Bob"
    }
    return "Alice"
}

func main() {
    // Example 1:
    // Input: x = 2, y = 7
    // Output: "Alice"
    // Explanation:
    // The game ends in a single turn:
    // Alice picks 1 coin with a value of 75 and 4 coins with a value of 10.
    fmt.Println(losingPlayer(2, 7)) // "Alice"
    // Example 2:
    // Input: x = 4, y = 11
    // Output: "Bob"
    // Explanation:
    // The game ends in 2 turns:
    // Alice picks 1 coin with a value of 75 and 4 coins with a value of 10.
    // Bob picks 1 coin with a value of 75 and 4 coins with a value of 10.
    fmt.Println(losingPlayer(4, 11)) // "Bob"
}