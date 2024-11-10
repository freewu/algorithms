package main

// 1817. Finding the Users Active Minutes
// You are given the logs for users' actions on LeetCode, and an integer k. 
// The logs are represented by a 2D integer array logs where each logs[i] = [IDi, timei] indicates 
// that the user with IDi performed an action at the minute timei.

// Multiple users can perform actions simultaneously, and a single user can perform multiple actions in the same minute.

// The user active minutes (UAM) for a given user is defined as the number of unique minutes in which the user performed an action on LeetCode. 
// A minute can only be counted once, even if multiple actions occur during it.

// You are to calculate a 1-indexed array answer of size k such that, 
// for each j (1 <= j <= k), answer[j] is the number of users whose UAM equals j.

// Return the array answer as described above.

// Example 1:
// Input: logs = [[0,5],[1,2],[0,2],[0,5],[1,3]], k = 5
// Output: [0,2,0,0,0]
// Explanation:
// The user with ID=0 performed actions at minutes 5, 2, and 5 again. Hence, they have a UAM of 2 (minute 5 is only counted once).
// The user with ID=1 performed actions at minutes 2 and 3. Hence, they have a UAM of 2.
// Since both users have a UAM of 2, answer[2] is 2, and the remaining answer[j] values are 0.

// Example 2:
// Input: logs = [[1,1],[2,2],[2,3]], k = 4
// Output: [1,1,0,0]
// Explanation:
// The user with ID=1 performed a single action at minute 1. Hence, they have a UAM of 1.
// The user with ID=2 performed actions at minutes 2 and 3. Hence, they have a UAM of 2.
// There is one user with a UAM of 1 and one with a UAM of 2.
// Hence, answer[1] = 1, answer[2] = 1, and the remaining values are 0.

// Constraints:
//     1 <= logs.length <= 10^4
//     0 <= IDi <= 10^9
//     1 <= timei <= 10^5
//     k is in the range [The maximum UAM for a user, 10^5].

import "fmt"
import "sort"

func findingUsersActiveMinutes(logs [][]int, k int) []int {
    users := make(map[int]map[int]bool)
    for _, log := range logs {
        id, time := log[0], log[1]
        times := users[id]
        if times == nil {
            times = map[int]bool{ time: true}
        } else {
            times[time] = true
        }
        users[id] = times
    }
    res := make([]int, k)
    for _, action := range users {
        res[len(action) - 1]++
    }
    return res
}

func findingUsersActiveMinutes1(logs [][]int, k int) []int {
    res := make([]int, k)
    sort.Slice(logs, func(i, j int) bool { // 按 id  & time 从小到大排序
        if logs[i][0] == logs[j][0] { return logs[i][1] < logs[j][1] }
        return logs[i][0] < logs[j][0]
    })
    id, time, count := logs[0][0], logs[0][1], 1
    for i := 1; i < len(logs); i++ {
        if id != logs[i][0] {
            res[count-1]++
            id, time, count = logs[i][0], logs[i][1], 1
        } else if time != logs[i][1] {
            time = logs[i][1]
            count++
        }
    }
    res[count - 1]++
    return res
}

func main() {
    // Example 1:
    // Input: logs = [[0,5],[1,2],[0,2],[0,5],[1,3]], k = 5
    // Output: [0,2,0,0,0]
    // Explanation:
    // The user with ID=0 performed actions at minutes 5, 2, and 5 again. Hence, they have a UAM of 2 (minute 5 is only counted once).
    // The user with ID=1 performed actions at minutes 2 and 3. Hence, they have a UAM of 2.
    // Since both users have a UAM of 2, answer[2] is 2, and the remaining answer[j] values are 0.
    fmt.Println(findingUsersActiveMinutes([][]int{{0,5},{1,2},{0,2},{0,5},{1,3}}, 5)) // [0,2,0,0,0]
    // Example 2:
    // Input: logs = [[1,1],[2,2],[2,3]], k = 4
    // Output: [1,1,0,0]
    // Explanation:
    // The user with ID=1 performed a single action at minute 1. Hence, they have a UAM of 1.
    // The user with ID=2 performed actions at minutes 2 and 3. Hence, they have a UAM of 2.
    // There is one user with a UAM of 1 and one with a UAM of 2.
    // Hence, answer[1] = 1, answer[2] = 1, and the remaining values are 0.
    fmt.Println(findingUsersActiveMinutes([][]int{{1,1},{2,2},{2,3}}, 4)) // [1,1,0,0]

    fmt.Println(findingUsersActiveMinutes1([][]int{{0,5},{1,2},{0,2},{0,5},{1,3}}, 5)) // [0,2,0,0,0]
    fmt.Println(findingUsersActiveMinutes1([][]int{{1,1},{2,2},{2,3}}, 4)) // [1,1,0,0]
}