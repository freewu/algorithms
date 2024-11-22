package main

// 3238. Find the Number of Winning Players
// You are given an integer n representing the number of players in a game 
// and a 2D array pick where pick[i] = [xi, yi] represents that the player xi picked a ball of color yi.

// Player i wins the game if they pick strictly more than i balls of the same color. 
// In other words,
//     Player 0 wins if they pick any ball.
//     Player 1 wins if they pick at least two balls of the same color.
//     ...
//     Player i wins if they pick at leasti + 1 balls of the same color.
//     Return the number of players who win the game.

// Note that multiple players can win the game.

// Example 1:
// Input: n = 4, pick = [[0,0],[1,0],[1,0],[2,1],[2,1],[2,0]]
// Output: 2
// Explanation:
// Player 0 and player 1 win the game, while players 2 and 3 do not win.

// Example 2:
// Input: n = 5, pick = [[1,1],[1,2],[1,3],[1,4]]
// Output: 0
// Explanation:
// No player wins the game.

// Example 3:
// Input: n = 5, pick = [[1,1],[2,4],[2,4],[2,4]]
// Output: 1
// Explanation:
// Player 2 wins the game by picking 3 balls with color 4.

// Constraints:
//     2 <= n <= 10
//     1 <= pick.length <= 100
//     pick[i].length == 2
//     0 <= xi <= n - 1
//     0 <= yi <= 10

import "fmt"

func winningPlayerCount(n int, pick [][]int) int {
    // Step 1: Create a 2D slice to count colors picked by each player
    colorCounts := make([][]int, n)
    for i := range colorCounts {
        colorCounts[i] = make([]int, 11)
    }
    // Step 2: Update the color counts based on the picks
    for _, v := range pick {
        player, color := v[0], v[1]
        colorCounts[player][color]++
    }
    // Step 3: Count the number of winning players
    res := 0
    for i := 0; i < n; i++ {
        for _, v := range colorCounts[i] {
            if v > i {
                res++
                break
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, pick = [[0,0],[1,0],[1,0],[2,1],[2,1],[2,0]]
    // Output: 2
    // Explanation:
    // Player 0 and player 1 win the game, while players 2 and 3 do not win.
    fmt.Println(winningPlayerCount(4, [][]int{{0,0},{1,0},{1,0},{2,1},{2,1},{2,0}})) // 2
    // Example 2:
    // Input: n = 5, pick = [[1,1],[1,2],[1,3],[1,4]]
    // Output: 0
    // Explanation:
    // No player wins the game.
    fmt.Println(winningPlayerCount(5, [][]int{{1,1},{1,2},{1,3},{1,4}})) // 0
    // Example 3:
    // Input: n = 5, pick = [[1,1],[2,4],[2,4],[2,4]]
    // Output: 1
    // Explanation:
    // Player 2 wins the game by picking 3 balls with color 4.
    fmt.Println(winningPlayerCount(5, [][]int{{1,1},{2,4},{2,4},{2,4}})) // 1
}