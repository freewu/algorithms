package main

// 3320. Count The Number of Winning Sequences
// Alice and Bob are playing a fantasy battle game consisting of n rounds where they summon one of three magical creatures each round: a Fire Dragon, a Water Serpent, or an Earth Golem. 
// In each round, players simultaneously summon their creature and are awarded points as follows:
//     1. If one player summons a Fire Dragon and the other summons an Earth Golem, 
//        the player who summoned the Fire Dragon is awarded a point.
//     2. If one player summons a Water Serpent and the other summons a Fire Dragon, 
//        the player who summoned the Water Serpent is awarded a point.
//     3. If one player summons an Earth Golem and the other summons a Water Serpent, 
//        the player who summoned the Earth Golem is awarded a point.
//     4. If both players summon the same creature, no player is awarded a point.

// You are given a string s consisting of n characters 'F', 'W', and 'E', 
// representing the sequence of creatures Alice will summon in each round:
//     1. If s[i] == 'F', Alice summons a Fire Dragon.
//     2. If s[i] == 'W', Alice summons a Water Serpent.
//     3. If s[i] == 'E', Alice summons an Earth Golem.

// Bob’s sequence of moves is unknown, but it is guaranteed that Bob will never summon the same creature in two consecutive rounds. 
// Bob beats Alice if the total number of points awarded to Bob after n rounds is strictly greater than the points awarded to Alice.

// Return the number of distinct sequences Bob can use to beat Alice.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: s = "FFF"
// Output: 3
// Explanation:
// Bob can beat Alice by making one of the following sequences of moves: "WFW", "FWF", or "WEW". 
// Note that other winning sequences like "WWE" or "EWW" are invalid since Bob cannot make the same move twice in a row.

// Example 2:
// Input: s = "FWEFW"
// Output: 18
// Explanation:
// Bob can beat Alice by making one of the following sequences of moves: "FWFWF", "FWFWE", "FWEFE", "FWEWE", "FEFWF", "FEFWE", "FEFEW", "FEWFE", "WFEFE", "WFEWE", "WEFWF", "WEFWE", "WEFEF", "WEFEW", "WEWFW", "WEWFE", "EWFWE", or "EWEWE".

// Constraints:
//     1 <= s.length <= 1000
//     s[i] is one of 'F', 'W', or 'E'.

import "fmt"

func countWinningSequences(s string) int {
    res, mod := 0, 1_000_000_007
    mp, lose, win := map[rune]int{ 'F': 0, 'W': 1, 'E': 2, }, map[int]int{ 0: 2, 1: 0, 2: 1, }, map[int]int{ 2: 0, 0: 1, 1: 2, }
    dp, prev, states := [2001][3]int{}, [2001][3]int{}, [3][2]int{}
    for i, c := range s {
        l, w, d := lose[mp[c]], win[mp[c]], mp[c]
        if i == 0 {
            dp[1000 - 1][l], dp[1000 + 0][d], dp[1000 + 1][w] = 1, 1, 1
            continue
        }
        prev, dp = dp, prev // reset mem
        dp = [2001][3]int{}
        for p, pdp := range prev {
            states[0], states[1], states[2] = [2]int{ p + 1, w}, [2]int{ p - 1, l}, [2]int{ p, d }
            for old, count := range pdp {
                if count == 0 { continue }
                for _, state := range states {
                    if state[1] == old { continue}
                    dp[state[0]][state[1]] += (count % mod)
                }
            }
        }
    }
    for _, row := range dp[1001:] {
        for _, count := range row {
            res += count
        }
    }
    return res % mod
}

func countWinningSequences1(s string) int {
    n, mod := len(s), 1_000_000_007
    mp := [...]int{'F': 0, 'W': 1, 'E': 2}
    dp := make([][][3]int, n + 1)
    for i := range dp {
        dp[i] = make([][3]int, 2 * n + 1)
    }
    for i := n + 1; i <= 2 * n; i++ {
        dp[0][i] = [3]int{1, 1, 1}
    }
    for i, c := range s {
        for j := -i; j < n-i; j++ {
            for p := 0; p < 3; p++ {
                v := 0
                for q := 0; q < 3; q++ {
                    if i == n-1 || q != p {
                        a := (q - mp[c] + 3) % 3
                        if a == 2 {
                            a = -1
                        }
                        v += dp[i][j + a + n][q]
                    }
                }
                dp[i+1][j+n][p] = v % mod
            }
        }
    }
    return dp[n][n][0]
}

func main() {
    // Example 1:
    // Input: s = "FFF"
    // Output: 3
    // Explanation:
    // Bob can beat Alice by making one of the following sequences of moves: "WFW", "FWF", or "WEW". 
    // Note that other winning sequences like "WWE" or "EWW" are invalid since Bob cannot make the same move twice in a row.
    fmt.Println(countWinningSequences("FFF")) // 3
    // Example 2:
    // Input: s = "FWEFW"
    // Output: 18
    // Explanation:
    // Bob can beat Alice by making one of the following sequences of moves: "FWFWF", "FWFWE", "FWEFE", "FWEWE", "FEFWF", "FEFWE", "FEFEW", "FEWFE", "WFEFE", "WFEWE", "WEFWF", "WEFWE", "WEFEF", "WEFEW", "WEWFW", "WEWFE", "EWFWE", or "EWEWE".
    fmt.Println(countWinningSequences("FWEFW")) // 18

    fmt.Println(countWinningSequences("FFFFFFFF")) // 142
    fmt.Println(countWinningSequences("WWWWWWWW")) // 142
    fmt.Println(countWinningSequences("EEEEEEEE")) // 142

    fmt.Println(countWinningSequences1("FFF")) // 3
    fmt.Println(countWinningSequences1("FWEFW")) // 18
    fmt.Println(countWinningSequences1("FFFFFFFF")) // 142
    fmt.Println(countWinningSequences1("WWWWWWWW")) // 142
    fmt.Println(countWinningSequences1("EEEEEEEE")) // 142
}