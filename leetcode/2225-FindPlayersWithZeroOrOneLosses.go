package main

// 2225. Find Players With Zero or One Losses
// You are given an integer array matches where matches[i] = [winneri, loseri] indicates 
// that the player winneri defeated player loseri in a match.

// Return a list answer of size 2 where:
//     answer[0] is a list of all players that have not lost any matches.
//     answer[1] is a list of all players that have lost exactly one match.

// The values in the two lists should be returned in increasing order.

// Note:
//     You should only consider the players that have played at least one match.
//     The testcases will be generated such that no two matches will have the same outcome.

// Example 1:
// Input: matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
// Output: [[1,2,10],[4,5,7,8]]
// Explanation:
// Players 1, 2, and 10 have not lost any matches.
// Players 4, 5, 7, and 8 each have lost one match.
// Players 3, 6, and 9 each have lost two matches.
// Thus, answer[0] = [1,2,10] and answer[1] = [4,5,7,8].

// Example 2:
// Input: matches = [[2,3],[1,3],[5,4],[6,4]]
// Output: [[1,2,5,6],[]]
// Explanation:
// Players 1, 2, 5, and 6 have not lost any matches.
// Players 3 and 4 each have lost two matches.
// Thus, answer[0] = [1,2,5,6] and answer[1] = [].
 
// Constraints:
//     1 <= matches.length <= 10^5
//     matches[i].length == 2
//     1 <= winneri, loseri <= 10^5
//     winneri != loseri
//     All matches[i] are unique.

import "fmt"
import "sort"

func findWinners(matches [][]int) [][]int {
    res, losses := make([][]int, 2), make([]int, 100_001) //  1 <= winneri, loseri <= 10^5
    res[0] = make([]int,0)
    res[1] = make([]int,0)
    for _, match := range matches {
        if losses[match[1]] == 0 { // 输
            losses[match[1]] = 1
        }
        if losses[match[0]] == 0 { // 赢
            losses[match[0]] = 1
        }
        losses[match[1]]++ // 累加输的场次
    }
    for player, score := range losses { // 只处理 > 1 的值
        if score == 1 { // 没有输过
            res[0] = append(res[0], player)
        } else if score == 2 { // 只输一场
            res[1] = append(res[1], player)
        }
    }
    return res
}

func findWinners1(matches [][]int) [][]int {
    res, losses := make([][]int, 2), make(map[int]int) 
    res[0], res[1] = []int{}, []int{}
    for _, match := range matches {
        if _, ok := losses[match[0]]; !ok { losses[match[0]] = 0; } // 赢
        if _, ok := losses[match[1]]; !ok { losses[match[1]] = 0; } // 输
        losses[match[1]]++ // 累加输的场次
    }
    for player, score := range losses { // 只处理 > 1 的值
        if score == 0 { // 没有输过
            res[0] = append(res[0], player)
        } else if score == 1 { // 只输一场
            res[1] = append(res[1], player)
        }
    }
    sort.Ints(res[0])
    sort.Ints(res[1])
    return res
}

func main() {
    // Example 1:
    // Input: matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
    // Output: [[1,2,10],[4,5,7,8]]
    // Explanation:
    // Players 1, 2, and 10 have not lost any matches.
    // Players 4, 5, 7, and 8 each have lost one match.
    // Players 3, 6, and 9 each have lost two matches.
    // Thus, answer[0] = [1,2,10] and answer[1] = [4,5,7,8].
    fmt.Println(findWinners([][]int{{1,3},{2,3},{3,6},{5,6},{5,7},{4,5},{4,8},{4,9},{10,4},{10,9}})) // [[1,2,10],[4,5,7,8]]
    // Example 2:
    // Input: matches = [[2,3],[1,3],[5,4],[6,4]]
    // Output: [[1,2,5,6],[]]
    // Explanation:
    // Players 1, 2, 5, and 6 have not lost any matches.
    // Players 3 and 4 each have lost two matches.
    // Thus, answer[0] = [1,2,5,6] and answer[1] = [].
    fmt.Println(findWinners([][]int{{2,3},{1,3},{5,4},{6,4}})) // [[1,2,5,6],[]]

    fmt.Println(findWinners1([][]int{{1,3},{2,3},{3,6},{5,6},{5,7},{4,5},{4,8},{4,9},{10,4},{10,9}})) // [[1,2,10],[4,5,7,8]]
    fmt.Println(findWinners1([][]int{{2,3},{1,3},{5,4},{6,4}})) // [[1,2,5,6],[]]
}