package main

// 1366. Rank Teams by Votes
// In a special ranking system, each voter gives a rank from highest to lowest to all teams participating in the competition.

// The ordering of teams is decided by who received the most position-one votes. 
// If two or more teams tie in the first position, we consider the second position to resolve the conflict, if they tie again, we continue this process until the ties are resolved. 
// If two or more teams are still tied after considering all positions, we rank them alphabetically based on their team letter.

// You are given an array of strings votes which is the votes of all voters in the ranking systems. 
// Sort all teams according to the ranking system described above.

// Return a string of all teams sorted by the ranking system.

// Example 1:
// Input: votes = ["ABC","ACB","ABC","ACB","ACB"]
// Output: "ACB"
// Explanation: 
// Team A was ranked first place by 5 voters. No other team was voted as first place, so team A is the first team.
// Team B was ranked second by 2 voters and ranked third by 3 voters.
// Team C was ranked second by 3 voters and ranked third by 2 voters.
// As most of the voters ranked C second, team C is the second team, and team B is the third.

// Example 2:
// Input: votes = ["WXYZ","XYZW"]
// Output: "XWYZ"
// Explanation:
// X is the winner due to the tie-breaking rule. X has the same votes as W for the first position, but X has one vote in the second position, while W does not have any votes in the second position. 

// Example 3:
// Input: votes = ["ZMNAGUEDSJYLBOPHRQICWFXTVK"]
// Output: "ZMNAGUEDSJYLBOPHRQICWFXTVK"
// Explanation: Only one voter, so their votes are used for the ranking.

// Constraints:
//     1 <= votes.length <= 1000
//     1 <= votes[i].length <= 26
//     votes[i].length == votes[j].length for 0 <= i, j < votes.length.
//     votes[i][j] is an English uppercase letter.
//     All characters of votes[i] are unique.
//     All the characters that occur in votes[0] also occur in votes[j] where 1 <= j < votes.length.

import "fmt"
import "sort"

func rankTeams(votes []string) string {
    type VoteInfo struct {
        char int
        count []int
    }
    res, n := []byte{}, len(votes[0])
    visited := make([]VoteInfo, 26)
    for i := range visited {
        visited[i] = VoteInfo {
            char: i,
            count: make([]int, n),
        }
    }
    for _, v := range votes {
        for i := 0; i < n; i++ {
            visited[int(v[i] - 'A')].count[i]++
        }
    }
    sort.Slice(visited, func(i, j int) bool {
        for k := 0; k < n; k++ {
            if visited[i].count[k] == visited[j].count[k] { continue }
            if visited[i].count[k] > visited[j].count[k]  { return true }
            return false
        }
        if visited[i].char < visited[j].char { return true }
        return false
    })
    for i := 0; i < n; i++ {
        res = append(res, byte('A' + visited[i].char))
    }
    return string(res)
}

func rankTeams1(votes []string) string {
    dp := [26][26]int{}
    for _, vote := range votes {
        for i, c := range vote {
            dp[int(c - 'A')][i]++
        }
    }
    res := []byte(votes[0])
    sort.Slice(res, func(i, j int) bool {
        x, y := int(res[i] - 'A'), int(res[j] - 'A')
        for i := 0; i < 26; i++ {
            if dp[x][i] != dp[y][i] {
                return dp[x][i] > dp[y][i]
            }
        }
        return x < y
    })
    return string(res)
}

func main() {
    // Example 1:
    // Input: votes = ["ABC","ACB","ABC","ACB","ACB"]
    // Output: "ACB"
    // Explanation: 
    // Team A was ranked first place by 5 voters. No other team was voted as first place, so team A is the first team.
    // Team B was ranked second by 2 voters and ranked third by 3 voters.
    // Team C was ranked second by 3 voters and ranked third by 2 voters.
    // As most of the voters ranked C second, team C is the second team, and team B is the third.
    fmt.Println(rankTeams([]string{"ABC","ACB","ABC","ACB","ACB"})) // "ACB"
    // Example 2:
    // Input: votes = ["WXYZ","XYZW"]
    // Output: "XWYZ"
    // Explanation:
    // X is the winner due to the tie-breaking rule. X has the same votes as W for the first position, but X has one vote in the second position, while W does not have any votes in the second position. 
    fmt.Println(rankTeams([]string{"WXYZ","XYZW"})) // "XWYZ"
    // Example 3:
    // Input: votes = ["ZMNAGUEDSJYLBOPHRQICWFXTVK"]
    // Output: "ZMNAGUEDSJYLBOPHRQICWFXTVK"
    // Explanation: Only one voter, so their votes are used for the ranking.
    fmt.Println(rankTeams([]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"})) // "ZMNAGUEDSJYLBOPHRQICWFXTVK"

    fmt.Println(rankTeams1([]string{"ABC","ACB","ABC","ACB","ACB"})) // "ACB"
    fmt.Println(rankTeams1([]string{"WXYZ","XYZW"})) // "XWYZ"
    fmt.Println(rankTeams1([]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"})) // "ZMNAGUEDSJYLBOPHRQICWFXTVK"
}