package main

// 2660. Determine the Winner of a Bowling Game
// You are given two 0-indexed integer arrays player1 and player2, 
// representing the number of pins that player 1 and player 2 hit in a bowling game, respectively.

// The bowling game consists of n turns, and the number of pins in each turn is exactly 10.

// Assume a player hits xi pins in the ith turn. The value of the ith turn for the player is:
//     1. 2xi if the player hits 10 pins in either (i - 1)th or (i - 2)th turn.
//     2. Otherwise, it is xi.

// The score of the player is the sum of the values of their n turns.

// Return
//     1 if the score of player 1 is more than the score of player 2,
//     2 if the score of player 2 is more than the score of player 1, and
//     0 in case of a draw.

// Example 1:
// Input: player1 = [5,10,3,2], player2 = [6,5,7,3]
// Output: 1
// Explanation:
// The score of player 1 is 5 + 10 + 2*3 + 2*2 = 25.
// The score of player 2 is 6 + 5 + 7 + 3 = 21.

// Example 2:
// Input: player1 = [3,5,7,6], player2 = [8,10,10,2]
// Output: 2
// Explanation:
// The score of player 1 is 3 + 5 + 7 + 6 = 21.
// The score of player 2 is 8 + 10 + 2*10 + 2*2 = 42.

// Example 3:
// Input: player1 = [2,3], player2 = [4,1]
// Output: 0
// Explanation:
// The score of player1 is 2 + 3 = 5.
// The score of player2 is 4 + 1 = 5.

// Example 4:
// Input: player1 = [1,1,1,10,10,10,10], player2 = [10,10,10,10,1,1,1]
// Output: 2
// Explanation:
// The score of player1 is 1 + 1 + 1 + 10 + 2*10 + 2*10 + 2*10 = 73.
// The score of player2 is 10 + 2*10 + 2*10 + 2*10 + 2*1 + 2*1 + 1 = 75.

// Constraints:
//     n == player1.length == player2.length
//     1 <= n <= 1000
//     0 <= player1[i], player2[i] <= 10

import "fmt"

func isWinner(player1 []int, player2 []int) int {
    score1, score2 := 0,  0
    for i, v := range player1 {
        score1 += v
        score2 += player2[i]
        if i - 1 >= 0 && player1[i - 1] == 10 || i - 2 >= 0 && player1[i - 2] == 10 {
            score1 += v
        } 
        if i - 1 >= 0 && player2[i - 1] == 10 || i - 2 >= 0 && player2[i - 2] == 10 {
            score2 += player2[i]
        }
    }
    if score1 > score2 { return 1 } 
    if score1 < score2 { return 2 } 
    return 0
}

func isWinner1(player1 []int, player2 []int) int {
    calc := func(arr []int) int {
        res, double := 0, 0
        for _, v := range arr {
            if double > 0 {
                res += 2 * v
                double--
            } else {
                res += v
            }
            if v == 10 {
                double = 2
            }
        }
        return res
    }
    score1, score2 := calc(player1), calc(player2)
    if score1 > score2 { return 1 }
    if score1 < score2 { return 2 }
    return 0
}

func main() {
    // Example 1:
    // Input: player1 = [5,10,3,2], player2 = [6,5,7,3]
    // Output: 1
    // Explanation:
    // The score of player 1 is 5 + 10 + 2*3 + 2*2 = 25.
    // The score of player 2 is 6 + 5 + 7 + 3 = 21.
    fmt.Println(isWinner([]int{5,10,3,2}, []int{6,5,7,3})) // 1
    // Example 2:
    // Input: player1 = [3,5,7,6], player2 = [8,10,10,2]
    // Output: 2
    // Explanation:
    // The score of player 1 is 3 + 5 + 7 + 6 = 21.
    // The score of player 2 is 8 + 10 + 2*10 + 2*2 = 42.
    fmt.Println(isWinner([]int{3,5,7,6}, []int{8,10,10,2})) // 2
    // Example 3:
    // Input: player1 = [2,3], player2 = [4,1]
    // Output: 0
    // Explanation:
    // The score of player1 is 2 + 3 = 5.
    // The score of player2 is 4 + 1 = 5.
    fmt.Println(isWinner([]int{2,3}, []int{4,1})) // 0
    // Example 4:
    // Input: player1 = [1,1,1,10,10,10,10], player2 = [10,10,10,10,1,1,1]
    // Output: 2
    // Explanation:
    // The score of player1 is 1 + 1 + 1 + 10 + 2*10 + 2*10 + 2*10 = 73.
    // The score of player2 is 10 + 2*10 + 2*10 + 2*10 + 2*1 + 2*1 + 1 = 75.
    fmt.Println(isWinner([]int{1,1,1,10,10,10,10}, []int{10,10,10,10,1,1,1})) // 2

    fmt.Println(isWinner([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(isWinner1([]int{5,10,3,2}, []int{6,5,7,3})) // 1
    fmt.Println(isWinner1([]int{3,5,7,6}, []int{8,10,10,2})) // 2
    fmt.Println(isWinner1([]int{2,3}, []int{4,1})) // 0
    fmt.Println(isWinner1([]int{1,1,1,10,10,10,10}, []int{10,10,10,10,1,1,1})) // 2
    fmt.Println(isWinner1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 0
}