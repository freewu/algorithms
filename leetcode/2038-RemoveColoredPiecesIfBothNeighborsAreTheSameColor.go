package main

// 2038. Remove Colored Pieces if Both Neighbors are the Same Color
// There are n pieces arranged in a line, and each piece is colored either by 'A' or by 'B'. 
// You are given a string colors of length n where colors[i] is the color of the ith piece.

// Alice and Bob are playing a game where they take alternating turns removing pieces from the line. 
// In this game, Alice moves first.
//     1. Alice is only allowed to remove a piece colored 'A' if both its neighbors are also colored 'A'. 
//        She is not allowed to remove pieces that are colored 'B'.
//     2. Bob is only allowed to remove a piece colored 'B' if both its neighbors are also colored 'B'. 
//        He is not allowed to remove pieces that are colored 'A'.
//     3. Alice and Bob cannot remove pieces from the edge of the line.
//     4. If a player cannot make a move on their turn, that player loses and the other player wins.

// Assuming Alice and Bob play optimally, return true if Alice wins, or return false if Bob wins.

// Example 1:
// Input: colors = "AAABABB"
// Output: true
// Explanation:
// AAABABB -> AABABB
// Alice moves first.
// She removes the second 'A' from the left since that is the only 'A' whose neighbors are both 'A'.
// Now it's Bob's turn.
// Bob cannot make a move on his turn since there are no 'B's whose neighbors are both 'B'.
// Thus, Alice wins, so return true.

// Example 2:
// Input: colors = "AA"
// Output: false
// Explanation:
// Alice has her turn first.
// There are only two 'A's and both are on the edge of the line, so she cannot move on her turn.
// Thus, Bob wins, so return false.

// Example 3:
// Input: colors = "ABBBBBBBAAA"
// Output: false
// Explanation:
// ABBBBBBBAAA -> ABBBBBBBAA
// Alice moves first.
// Her only option is to remove the second to last 'A' from the right.
// ABBBBBBBAA -> ABBBBBBAA
// Next is Bob's turn.
// He has many options for which 'B' piece to remove. He can pick any.
// On Alice's second turn, she has no more pieces that she can remove.
// Thus, Bob wins, so return false.

// Constraints:
//     1 <= colors.length <= 10^5
//     colors consists of only the letters 'A' and 'B'

import "fmt"

func winnerOfGame(colors string) bool {
    res, cur, count := [2]int{}, colors[0], 0
    for i := 0; i < len(colors); i++{
        if colors[i] != cur {
            cur, count = colors[i], 0
        }
        count++
        if count >= 3 {
            res[cur - 'A']++
        }
    }
    return res[0] > res[1]
}

func winnerOfGame1(colors string) bool {
    alice, bob, count := 0, 0, 1
    for i := 1; i < len(colors); i++{
        if(colors[i] == colors[i - 1]){
            count++
            if count >= 3 {
                if(colors[i] == 'A'){
                    alice++
                } else {
                    bob++
                }
            }
        }else{
            count = 1
        }
    }
    return alice > bob
}

func main() {
    // Example 1:
    // Input: colors = "AAABABB"
    // Output: true
    // Explanation:
    // AAABABB -> AABABB
    // Alice moves first.
    // She removes the second 'A' from the left since that is the only 'A' whose neighbors are both 'A'.
    // Now it's Bob's turn.
    // Bob cannot make a move on his turn since there are no 'B's whose neighbors are both 'B'.
    // Thus, Alice wins, so return true.
    fmt.Println(winnerOfGame("AAABABB")) // true
    // Example 2:
    // Input: colors = "AA"
    // Output: false
    // Explanation:
    // Alice has her turn first.
    // There are only two 'A's and both are on the edge of the line, so she cannot move on her turn.
    // Thus, Bob wins, so return false.
    fmt.Println(winnerOfGame("AA")) // false
    // Example 3:
    // Input: colors = "ABBBBBBBAAA"
    // Output: false
    // Explanation:
    // ABBBBBBBAAA -> ABBBBBBBAA
    // Alice moves first.
    // Her only option is to remove the second to last 'A' from the right.
    // ABBBBBBBAA -> ABBBBBBAA
    // Next is Bob's turn.
    // He has many options for which 'B' piece to remove. He can pick any.
    // On Alice's second turn, she has no more pieces that she can remove.
    // Thus, Bob wins, so return false.
    fmt.Println(winnerOfGame("ABBBBBBBAAA")) // false

    fmt.Println(winnerOfGame1("AAABABB")) // true
    fmt.Println(winnerOfGame1("AA")) // false
    fmt.Println(winnerOfGame1("ABBBBBBBAAA")) // false
}