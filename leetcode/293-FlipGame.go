package main

// 293. Flip Game
// You are playing a Flip Game with your friend.
// You are given a string currentState that contains only '+' and '-'.
// You and your friend take turns to flip two consecutive "++" into "--". 
// The game ends when a person can no longer make a move, and therefore the other person will be the winner.

// Return all possible states of the string currentState after one valid move. 
// You may return the answer in any order. If there is no valid move, return an empty list [].
 
// Example 1:
// Input: currentState = "++++"
// Output: ["--++","+--+","++--"]

// Example 2:
// Input: currentState = "+"
// Output: []
 
// Constraints:
//     1 <= currentState.length <= 500
//     currentState[i] is either '+' or '-'.

import "fmt"

func generatePossibleNextMoves(currentState string) []string {
    res := []string{}
    if len(currentState) < 2 { // 不够翻转
        return res
    }
    src := []byte(currentState)
    for i := 1; i < len(currentState); i++ {
        if currentState[i] == '+' && currentState[i - 1] == '+' { // 如果是连续 的两个 "++" 反转成 "--"
            src[i], src[i-1] = '-', '-'
            res = append(res, string(src)) //  记录可能的状态
            src[i], src[i-1] = '+', '+' // 恢复
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: currentState = "++++"
    // Output: ["--++","+--+","++--"]
    fmt.Println(generatePossibleNextMoves("++++")) // [--++, +--+, ++--]
    // Example 2:
    // Input: currentState = "+"
    // Output: []
    fmt.Println(generatePossibleNextMoves("+")) // []

    fmt.Println(generatePossibleNextMoves("+++")) // [--+, +--]
    fmt.Println(generatePossibleNextMoves("++++++")) // [--++++ +--+++ ++--++ +++--+ ++++--]
}