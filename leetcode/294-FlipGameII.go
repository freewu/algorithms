package main

// 294. Flip Game II
// You are playing a Flip Game with your friend.

// You are given a string currentState that contains only '+' and '-'. 
// You and your friend take turns to flip two consecutive "++" into "--". 
// The game ends when a person can no longer make a move, and therefore the other person will be the winner.

// Return true if the starting player can guarantee a win, and false otherwise.

// Example 1:
// Input: currentState = "++++"
// Output: true
// Explanation: The starting player can guarantee a win by flipping the middle "++" to become "+--+".

// Example 2:
// Input: currentState = "+"
// Output: false
 
// Constraints:
//     1 <= currentState.length <= 60
//     currentState[i] is either '+' or '-'.
    
// Follow up: Derive your algorithm's runtime complexity.

import "fmt"

func canWin(currentState string) bool {
    cache := make(map[string]bool)
    var backtrace func(cache map[string]bool, state string) bool
    backtrace = func(cache map[string]bool, state string) bool {
        if win, ok := cache[state]; ok {
            return win
        }
        for j := 1; j < len(state); j++ {
            if state[j-1:j+1] == "++" {
                newState := state[:j-1] + "--" + state[j+1:]
                if !backtrace(cache, newState) {
                    cache[state] = true
                    return true
                }
            }
        }
        cache[state] = false
        return false
    }
    return backtrace(cache, currentState)
}

func main() {
    // Example 1:
    // Input: currentState = "++++"
    // Output: true
    // Explanation: The starting player can guarantee a win by flipping the middle "++" to become "+--+".
    fmt.Println(canWin("++++")) // true
    // Example 2:
    // Input: currentState = "+"
    // Output: false
    fmt.Println(canWin("+")) // false

    fmt.Println(canWin("++")) // true
    fmt.Println(canWin("+++")) // true
}