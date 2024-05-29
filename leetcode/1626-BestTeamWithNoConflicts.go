package main

// 1626. Best Team With No Conflicts
// You are the manager of a basketball team. 
// For the upcoming tournament, you want to choose the team with the highest overall score. 
// The score of the team is the sum of scores of all the players in the team.

// However, the basketball team is not allowed to have conflicts. 
// A conflict exists if a younger player has a strictly higher score than an older player. 
// A conflict does not occur between players of the same age.

// Given two lists, scores and ages, where each scores[i] and ages[i] represents the score and age of the ith player, 
// respectively, return the highest overall score of all possible basketball teams.

// Example 1:
// Input: scores = [1,3,5,10,15], ages = [1,2,3,4,5]
// Output: 34
// Explanation: You can choose all the players.

// Example 2:
// Input: scores = [4,5,6,5], ages = [2,1,2,1]
// Output: 16
// Explanation: 
// It is best to choose the last 3 players. 
// Notice that you are allowed to choose multiple people of the same age.

// Example 3:
// Input: scores = [1,2,3,5], ages = [8,9,10,1]
// Output: 6
// Explanation: It is best to choose the first 3 players. 
 
// Constraints:
//     1 <= scores.length, ages.length <= 1000
//     scores.length == ages.length
//     1 <= scores[i] <= 10^6
//     1 <= ages[i] <= 1000

import "fmt"
import "sort"

func bestTeamScore(scores []int, ages []int) int {
    n, player := len(ages), [][]int{}
    for i := 0; i < n; i++ { // 把年龄和得分关联上
        player = append(player, []int{ ages[i], scores[i] })
    }
    sort.Slice(player, func(i, j int) bool { // 按年龄 从小 --> 大 如果年龄一样则 按得分 小 --> 大 
        return player[i][0] < player[j][0] || (player[i][0] == player[j][0] && player[i][1] < player[j][1])
    })
    res, dp := 0, make([]int, n + 1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        dp[i + 1] = player[i][1]
        for j := 0; j < i; j++ {
            if player[j][1] > player[i][1] { continue } //  得分大，年龄小 产冲突 跳过
            dp[i + 1] = max(dp[i + 1], player[i][1] + dp[j + 1])  
        }
        res = max(res, dp[i + 1])
    }
    return res
}

func bestTeamScore1(scores []int, ages []int) int {
    n, m := len(scores), 0
    type pair struct {
        score, age int
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp := make([]pair, n)
    for i := range dp {
        m = max(m, ages[i]) // 得到最大的年龄
        dp[i] = pair{scores[i], ages[i]}
    }
    sort.Slice(dp, func(i, j int) bool {
        if dp[i].score == dp[j].score {
            return dp[i].age < dp[j].age
        }
        return dp[i].score < dp[j].score
    })
    s := make([]int, m + 1)
    update := func(idx, val int) {
        for i := idx; i < len(s); i += i & (-i) {
            s[i] = max(s[i], val)
        }
    }
    query := func(idx int) (res int) {
        for i := idx; i > 0; i -= i & (-i) {
            res = max(res, s[i])
        }
        return
    }
    for _, v := range dp {
        update(v.age, query(v.age) + v.score)
    }
    return query(m)
}

func main() {
    // Example 1:
    // Input: scores = [1,3,5,10,15], ages = [1,2,3,4,5]
    // Output: 34
    // Explanation: You can choose all the players.
    fmt.Println(bestTeamScore([]int{1,3,5,10,15},[]int{1,2,3,4,5})) // 34
    // Example 2:
    // Input: scores = [4,5,6,5], ages = [2,1,2,1]
    // Output: 16
    // Explanation: 
    // It is best to choose the last 3 players. 
    // Notice that you are allowed to choose multiple people of the same age.
    fmt.Println(bestTeamScore([]int{4,5,6,5},[]int{2,1,2,1})) // 16
    // Example 3:
    // Input: scores = [1,2,3,5], ages = [8,9,10,1]
    // Output: 6
    // Explanation: It is best to choose the first 3 players. 
    fmt.Println(bestTeamScore([]int{1,2,3,5},[]int{8,9,10,1})) // 6

    fmt.Println(bestTeamScore1([]int{1,3,5,10,15},[]int{1,2,3,4,5})) // 34
    fmt.Println(bestTeamScore1([]int{4,5,6,5},[]int{2,1,2,1})) // 16
    fmt.Println(bestTeamScore1([]int{1,2,3,5},[]int{8,9,10,1})) // 6
}