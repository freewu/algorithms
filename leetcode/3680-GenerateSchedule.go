package main

// 3680. Generate Schedule
// You are given an integer n representing n teams. 
// You are asked to generate a schedule such that:
//     1. Each team plays every other team exactly twice: once at home and once away.
//     2. There is exactly one match per day; the schedule is a list of consecutive days and schedule[i] is the match on day i.
//     3. No team plays on consecutive days.

// Return a 2D integer array schedule, where schedule[i][0] represents the home team and schedule[i][1] represents the away team. 
// If multiple schedules meet the conditions, return any one of them.

// If no schedule exists that meets the conditions, return an empty array.

// Example 1:
// Input: n = 3
// Output: []
// Explanation:
// ​​​​​​​Since each team plays every other team exactly twice, a total of 6 matches need to be played: [0,1],[0,2],[1,2],[1,0],[2,0],[2,1].
// It's not possible to create a schedule without at least one team playing consecutive days.

// Example 2:
// Input: n = 5
// Output: [[0,1],[2,3],[0,4],[1,2],[3,4],[0,2],[1,3],[2,4],[0,3],[1,4],[2,0],[3,1],[4,0],[2,1],[4,3],[1,0],[3,2],[4,1],[3,0],[4,2]]
// Explanation:
// Since each team plays every other team exactly twice, a total of 20 matches need to be played.
// The output shows one of the schedules that meet the conditions. No team plays on consecutive days.

// Constraints:
//     2 <= n <= 50

import "fmt"
import "slices"

func generateSchedule(n int) [][]int {
    if n <= 4 { return [][]int{} }
    res := [][]int{{0, 1}, {2, 3}, {0, 4}, {1, 2}, {3, 4}, {0, 2}, {1, 3}, {2, 4}, {0, 3}, {1, 4}, {2, 0}, {3, 1}, {4, 0}, {2, 1}, {4, 3}, {1, 0}, {3, 2}, {4, 1}, {3, 0}, {4, 2}}
    if n == 5 { return res }
    addTeam := func(arr [][]int, n int) [][]int {
        adding := make([][]int, 0)
        for i := 0; i < n; i++ {
            adding = append(adding, []int{i, n}, []int{n, i})
        }
        for len(adding) > 0 {
            addingNow := adding[0]
            a1, a2 := addingNow[0], addingNow[1]
            for i := 0; i < len(arr)-1; i++ {
                if (arr)[i][0] != a1 && (arr)[i][0] != a2 && (arr)[i][1] != a1 && (arr)[i][1] != a2 &&
                    (arr)[i+1][0] != a1 && (arr)[i+1][0] != a2 && (arr)[i+1][1] != a1 && (arr)[i+1][1] != a2 {
                    arr = slices.Concat(arr[:i+1], [][]int{{a1, a2}}, arr[i+1:])
                    adding = adding[1:]
                    break
                }
            }
        }
        return arr
    }
    for i := 5; i <= n - 1; i++ {
        res = addTeam(res, i)
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: []
    // Explanation:
    // ​​​​​​​Since each team plays every other team exactly twice, a total of 6 matches need to be played: [0,1],[0,2],[1,2],[1,0],[2,0],[2,1].
    // It's not possible to create a schedule without at least one team playing consecutive days.
    fmt.Println(generateSchedule(3)) // []
    // Example 2:
    // Input: n = 5
    // Output: [[0,1],[2,3],[0,4],[1,2],[3,4],[0,2],[1,3],[2,4],[0,3],[1,4],[2,0],[3,1],[4,0],[2,1],[4,3],[1,0],[3,2],[4,1],[3,0],[4,2]]
    // Explanation:
    // Since each team plays every other team exactly twice, a total of 20 matches need to be played.
    // The output shows one of the schedules that meet the conditions. No team plays on consecutive days.
    fmt.Println(generateSchedule(5)) // [[0,1],[2,3],[0,4],[1,2],[3,4],[0,2],[1,3],[2,4],[0,3],[1,4],[2,0],[3,1],[4,0],[2,1],[4,3],[1,0],[3,2],[4,1],[3,0],[4,2]]

    fmt.Println(generateSchedule(2)) // []
    fmt.Println(generateSchedule(50)) // 
}