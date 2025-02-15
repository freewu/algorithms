package main

// 3360. Stone Removal Game
// Alice and Bob are playing a game where they take turns removing stones from a pile, with Alice going first.
//     1. Alice starts by removing exactly 10 stones on her first turn.
//     2. For each subsequent turn, each player removes exactly 1 fewer stone than the previous opponent.

// The player who cannot make a move loses the game.

// Given a positive integer n, return true if Alice wins the game and false otherwise.

// Example 1:
// Input: n = 12
// Output: true
// Explanation:
// Alice removes 10 stones on her first turn, leaving 2 stones for Bob.
// Bob cannot remove 9 stones, so Alice wins.

// Example 2:
// Input: n = 1
// Output: false
// Explanation:
// Alice cannot remove 10 stones, so Alice loses.

// Constraints:
//     1 <= n <= 50

import "fmt"

func canAliceWin(n int) bool {
    stones := 10 // Alice 在第一次操作中移除 恰好 10 个石头。
    for n >= stones {
        n -= stones
        stones-- // 接下来的每次操作中，每位玩家移除的石头数 恰好 为另一位玩家上一次操作的石头数减 1 
    }     
    return stones % 2 == 1
}

func main() {
    // Example 1:
    // Input: n = 12
    // Output: true
    // Explanation:
    // Alice removes 10 stones on her first turn, leaving 2 stones for Bob.
    // Bob cannot remove 9 stones, so Alice wins.
    fmt.Println(canAliceWin(12)) // true
    // Example 2:
    // Input: n = 1
    // Output: false
    // Explanation:
    // Alice cannot remove 10 stones, so Alice loses.
    fmt.Println(canAliceWin(1)) // false

    fmt.Println(canAliceWin(2)) // false
    fmt.Println(canAliceWin(3)) // false
    fmt.Println(canAliceWin(8)) // false
    fmt.Println(canAliceWin(25)) // false
    fmt.Println(canAliceWin(49)) // true
    fmt.Println(canAliceWin(50)) // true
}