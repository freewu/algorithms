package main

// 1900. The Earliest and Latest Rounds Where Players Compete
// There is a tournament where n players are participating. 
// The players are standing in a single row and are numbered from 1 to n based on their initial standing position 
// (player 1 is the first player in the row, player 2 is the second player in the row, etc.).

// The tournament consists of multiple rounds (starting from round number 1). 
// In each round, the ith player from the front of the row competes against the ith player from the end of the row, 
// and the winner advances to the next round. 
// When the number of players is odd for the current round, the player in the middle automatically advances to the next round.
//     For example, if the row consists of players 1, 2, 4, 6, 7
//         Player 1 competes against player 7.
//         Player 2 competes against player 6.
//         Player 4 automatically advances to the next round.

// After each round is over, the winners are lined back up in the row based on the original ordering assigned to them initially (ascending order).

// The players numbered firstPlayer and secondPlayer are the best in the tournament. 
// They can win against any other player before they compete against each other. 
// If any two other players compete against each other, either of them might win, and thus you may choose the outcome of this round.

// Given the integers n, firstPlayer, and secondPlayer, 
// return an integer array containing two values,
// the earliest possible round number and the latest possible round number in which these two players will compete against each other, respectively.

// Example 1:
// Input: n = 11, firstPlayer = 2, secondPlayer = 4
// Output: [3,4]
// Explanation:
// One possible scenario which leads to the earliest round number:
// First round: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
// Second round: 2, 3, 4, 5, 6, 11
// Third round: 2, 3, 4
// One possible scenario which leads to the latest round number:
// First round: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
// Second round: 1, 2, 3, 4, 5, 6
// Third round: 1, 2, 4
// Fourth round: 2, 4

// Example 2:
// Input: n = 5, firstPlayer = 1, secondPlayer = 5
// Output: [1,1]
// Explanation: The players numbered 1 and 5 compete in the first round.
// There is no way to make them compete in any other round.

// Constraints:
//     2 <= n <= 28
//     1 <= firstPlayer < secondPlayer <= n

import "fmt"

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    p1, p2 := min(firstPlayer, secondPlayer), max(firstPlayer, secondPlayer)
    if p1 + p2 == n + 1 { // p1 and p2 compete in the first round
        return []int{1, 1} 
    }
    if n == 3 || n == 4 { // p1 and p2 must compete in the second round (only two rounds).
        return []int{2, 2}
    }
    if p1 - 1 > n - p2 {  // Flip to make p1 be more closer to left than p2 to right end for convenience
        p1, p2 = n + 1 - p2, n + 1 - p1
    }
    mid, mn, mx := (n + 1) / 2, n, 1
    if p2 * 2 <= n + 1 { // p2 is in the first half (n odd or even) or exact middle (n odd)
        a, b := p1 - 1, p2 - p1 - 1
        // i represents the number of front players in A wins
        // j represents the number of front players in B wins
        for i := 0; i <= a; i++ {
            for j := 0; j <= b; j++ {
                res := earliestAndLatest(mid, i + 1, i + j + 2)
                mn, mx  = min(mn, res[0] + 1), max(mx, res[1] + 1)
            }
        }
    } else { // p2 is in the later half (and has >= p1 distance to the end)
        p4 := n + 1 - p2;
        // Group C are players between p4 and p2, (c+1)/2 will advance to the next round.
        a, b, c := p1 - 1,  p4 - p1 - 1,  p2 - p4 - 1;
        for i := 0; i <= a; i++ {
            for j := 0; j <= b; j++ {
                res := earliestAndLatest(mid, i + 1, i + j + 1 + (c + 1) / 2 + 1)
                mn, mx = min(mn, res[0] + 1), max(mx, res[1] + 1)
            }
        }
    }
    return []int{ mn, mx }
}

func main() {
    // Example 1:
    // Input: n = 11, firstPlayer = 2, secondPlayer = 4
    // Output: [3,4]
    // Explanation:
    // One possible scenario which leads to the earliest round number:
    // First round: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
    // Second round: 2, 3, 4, 5, 6, 11
    // Third round: 2, 3, 4
    // One possible scenario which leads to the latest round number:
    // First round: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
    // Second round: 1, 2, 3, 4, 5, 6
    // Third round: 1, 2, 4
    // Fourth round: 2, 4
    fmt.Println(earliestAndLatest(11, 2, 4)) // [3,4]
    // Example 2:
    // Input: n = 5, firstPlayer = 1, secondPlayer = 5
    // Output: [1,1]
    // Explanation: The players numbered 1 and 5 compete in the first round.
    // There is no way to make them compete in any other round.
    fmt.Println(earliestAndLatest(5, 1, 5)) // [1,1]

    fmt.Println(earliestAndLatest(2, 1, 2)) // [1,1]
    fmt.Println(earliestAndLatest(28, 1, 28)) // [1,1]
    fmt.Println(earliestAndLatest(28, 2, 28)) // [2 2]
}