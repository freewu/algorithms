package main

// 3021. Alice and Bob Playing Flower Game
// Alice and Bob are playing a turn-based game on a circular field surrounded by flowers. 
// The circle represents the field, and there are x flowers in the clockwise direction between Alice and Bob, and y flowers in the anti-clockwise direction between them.

// The game proceeds as follows:
//     1. Alice takes the first turn.
//     2. In each turn, a player must choose either the clockwise or anti-clockwise direction and pick one flower from that side.
//     3. At the end of the turn, if there are no flowers left at all, the current player captures their opponent and wins the game.

// Given two integers, n and m, the task is to compute the number of possible pairs (x, y) that satisfy the conditions:
//     1. Alice must win the game according to the described rules.
//     2. The number of flowers x in the clockwise direction must be in the range [1,n].
//     3. The number of flowers y in the anti-clockwise direction must be in the range [1,m].

// Return the number of possible pairs (x, y) that satisfy the conditions mentioned in the statement.

// Example 1:
// Input: n = 3, m = 2
// Output: 3
// Explanation: The following pairs satisfy conditions described in the statement: (1,2), (3,2), (2,1).

// Example 2:
// Input: n = 1, m = 1
// Output: 0
// Explanation: No pairs satisfy the conditions described in the statement.

// Constraints:
//     1 <= n, m <= 10^5

import "fmt"

func flowerGame(n int, m int) int64 {
    return int64(n * m / 2)
}

func flowerGame1(n int, m int) int64 {
    res := 0
    if n % 2 == 0 && m % 2 == 0 {
        res = (n * m) / 2
    } else if n % 2 == 0 && m % 2 != 0 {
        res = ((n / 2) * (m / 2)) + ((n / 2) * ((m + 1) / 2))
    } else if n % 2 != 0 && m % 2 == 0 {
        res = ((n / 2) * (m / 2)) + (((n + 1) / 2) * (m / 2))
    } else {
        res = (((n + 1) / 2) * (m / 2)) + ((n / 2) * ((m + 1) / 2))
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: n = 3, m = 2
    // Output: 3
    // Explanation: The following pairs satisfy conditions described in the statement: (1,2), (3,2), (2,1).
    fmt.Println(flowerGame(3, 2)) // 3
    // Example 2:
    // Input: n = 1, m = 1
    // Output: 0
    // Explanation: No pairs satisfy the conditions described in the statement.
    fmt.Println(flowerGame(1, 1)) // 0

    fmt.Println(flowerGame(100_000, 100_000)) // 5000000000
    fmt.Println(flowerGame(1, 100_000)) // 50000
    fmt.Println(flowerGame(100_000, 1)) // 50000

    fmt.Println(flowerGame1(3, 2)) // 3
    fmt.Println(flowerGame1(1, 1)) // 0
    fmt.Println(flowerGame1(100_000, 100_000)) // 5000000000
    fmt.Println(flowerGame1(1, 100_000)) // 50000
    fmt.Println(flowerGame1(100_000, 1)) // 50000
}