package main

// 1927. Sum Game
// Alice and Bob take turns playing a game, with Alice starting first.

// You are given a string num of even length consisting of digits and '?' characters. 
// On each turn, a player will do the following if there is still at least one '?' in num:
//     1. Choose an index i where num[i] == '?'.
//     2. Replace num[i] with any digit between '0' and '9'.

// The game ends when there are no more '?' characters in num.

// For Bob to win, the sum of the digits in the first half of num must be equal to the sum of the digits in the second half. 
// For Alice to win, the sums must not be equal.
//     For example, if the game ended with num = "243801", then Bob wins because 2+4+3 = 8+0+1. 
//     If the game ended with num = "243803", then Alice wins because 2+4+3 != 8+0+3.

// Assuming Alice and Bob play optimally, return true if Alice will win and false if Bob will win.

// Example 1:
// Input: num = "5023"
// Output: false
// Explanation: There are no moves to be made.
// The sum of the first half is equal to the sum of the second half: 5 + 0 = 2 + 3.

// Example 2:
// Input: num = "25??"
// Output: true
// Explanation: Alice can replace one of the '?'s with '9' and it will be impossible for Bob to make the sums equal.

// Example 3:
// Input: num = "?3295???"
// Output: false
// Explanation: It can be proven that Bob will always win. One possible outcome is:
// - Alice replaces the first '?' with '9'. num = "93295???".
// - Bob replaces one of the '?' in the right half with '9'. num = "932959??".
// - Alice replaces one of the '?' in the right half with '2'. num = "9329592?".
// - Bob replaces the last '?' in the right half with '7'. num = "93295927".
// Bob wins because 9 + 3 + 2 + 9 = 5 + 9 + 2 + 7.

// Constraints:
//     2 <= num.length <= 10^5
//     num.length is even.
//     num consists of only digits and '?'.

import "fmt"

func sumGame(num string) bool {
    n, sum, move := len(num), [2]int{}, [2]int{}
    for i, v := range num {
        if v == '?' {
            move[i/(n/2)]++
        } else {
            sum[i/(n/2)] += int(v - '0')
        }
    }
    // Alice picks the last number and wins
    if (move[0] + move[1]) % 2 == 1 {
        return true
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    // For the side that has more `?` than the other,
    // Alice wins if:
    // 1. That side's sum is greater than the other.
    // 2. The distance between sums is not 9*(# moves for Bob)
    movesForBob, diff := abs(move[0] - move[1]) / 2, abs(sum[1] - sum[0])
    if move[0] > move[1] {
        return sum[0] > sum[1] || diff != 9 * movesForBob
    }
    return sum[1] > sum[0] || diff != 9 * movesForBob
}

func main() {
    // Example 1:
    // Input: num = "5023"
    // Output: false
    // Explanation: There are no moves to be made.
    // The sum of the first half is equal to the sum of the second half: 5 + 0 = 2 + 3.
    fmt.Println(sumGame("5023")) // false
    // Example 2:
    // Input: num = "25??"
    // Output: true
    // Explanation: Alice can replace one of the '?'s with '9' and it will be impossible for Bob to make the sums equal.
    fmt.Println(sumGame("25??")) // true
    // Example 3:
    // Input: num = "?3295???"
    // Output: false
    // Explanation: It can be proven that Bob will always win. One possible outcome is:
    // - Alice replaces the first '?' with '9'. num = "93295???".
    // - Bob replaces one of the '?' in the right half with '9'. num = "932959??".
    // - Alice replaces one of the '?' in the right half with '2'. num = "9329592?".
    // - Bob replaces the last '?' in the right half with '7'. num = "93295927".
    // Bob wins because 9 + 3 + 2 + 9 = 5 + 9 + 2 + 7.
    fmt.Println(sumGame("?3295???")) // false
}