package main

// 1986. Minimum Number of Work Sessions to Finish the Tasks
// There are n tasks assigned to you. 
// The task times are represented as an integer array tasks of length n, where the ith task takes tasks[i] hours to finish. 
// A work session is when you work for at most sessionTime consecutive hours and then take a break.

// You should finish the given tasks in a way that satisfies the following conditions:
//     If you start a task in a work session, you must complete it in the same work session.
//     You can start a new task immediately after finishing the previous one.
//     You may complete the tasks in any order.

// Given tasks and sessionTime, return the minimum number of work sessions needed to finish all the tasks following the conditions above.

// The tests are generated such that sessionTime is greater than or equal to the maximum element in tasks[i].

// Example 1:
// Input: tasks = [1,2,3], sessionTime = 3
// Output: 2
// Explanation: You can finish the tasks in two work sessions.
// - First work session: finish the first and the second tasks in 1 + 2 = 3 hours.
// - Second work session: finish the third task in 3 hours.

// Example 2:
// Input: tasks = [3,1,3,1,1], sessionTime = 8
// Output: 2
// Explanation: You can finish the tasks in two work sessions.
// - First work session: finish all the tasks except the last one in 3 + 1 + 3 + 1 = 8 hours.
// - Second work session: finish the last task in 1 hour.

// Example 3:
// Input: tasks = [1,2,3,4,5], sessionTime = 15
// Output: 1
// Explanation: You can finish all the tasks in one work session.

// Constraints:
//     n == tasks.length
//     1 <= n <= 14
//     1 <= tasks[i] <= 10
//     max(tasks[i]) <= sessionTime <= 15

import "fmt"
import "sort"

// // 解答错误 52 / 96 
// func minSessions(tasks []int, sessionTime int) int {
//     res, n := 1 << 31, len(tasks)
//     visited := make([]bool, n)
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     var dfs func(currTime, sessionsTook, done int) 
//     dfs = func(currTime, sessionsTook, done int) {
//         if done == n {
//             res = min(res, sessionsTook)
//             return
//         }
//         if currTime == 0 {
//             currTime = sessionTime
//             sessionsTook++
//         }
//         for i := 0; i < n; i++ {
//             if !visited[i] && tasks[i] <= currTime {
//                 visited[i] = true
//                 dfs(currTime - tasks[i], sessionsTook, done + 1)
//                 visited[i] = false
//             }
//         }
//     }
//     dfs(sessionTime, 1, 0)
//     return res
// }

func minSessions(tasks []int, sessionTime int) int {
    n := len(tasks)
    visited, dp := make([]bool, 1 << n), make([]int, 1 << n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < 1 << n; i++ {
        t := 0
        dp[i] = 1 << 31
        for j, v := range tasks {
            if i >> j & 1 == 1 {
                t += v
            }
        }
        visited[i] = t <= sessionTime
    }
    for i := 1; i < 1 << n; i++ {
        for j := i; j > 0; j = (j - 1) & i {
            if visited[j] {
                dp[i] = min(dp[i], dp[i ^ j] + 1)
            }
        }
    }
    return dp[1 << n - 1]
}

func minSessions1(tasks []int, sessionTime int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(tasks))) // 按降序排列任务，减少深度
    var backTrace func(sessions []int, index int) bool
    backTrace = func(sessions []int, index int) bool {
        if index >= len(tasks) { return true }
        for i := 0; i < len(sessions); i++ {
            if sessions[i] + tasks[index] <= sessionTime {
                sessions[i] += tasks[index]
                if backTrace(sessions, index + 1) { return true
                }
                sessions[i] -= tasks[index]
            }
            if sessions[i] == 0 { break }
        }
        return false
    }
    left, right := 1, len(tasks)
    for left < right { // 二分查找
        mid := (left + right) / 2
        if backTrace(make([]int, mid), 0) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func main() {
    // Example 1:
    // Input: tasks = [1,2,3], sessionTime = 3
    // Output: 2
    // Explanation: You can finish the tasks in two work sessions.
    // - First work session: finish the first and the second tasks in 1 + 2 = 3 hours.
    // - Second work session: finish the third task in 3 hours.
    fmt.Println(minSessions([]int{1,2,3}, 3)) // 2
    // Example 2:
    // Input: tasks = [3,1,3,1,1], sessionTime = 8
    // Output: 2
    // Explanation: You can finish the tasks in two work sessions.
    // - First work session: finish all the tasks except the last one in 3 + 1 + 3 + 1 = 8 hours.
    // - Second work session: finish the last task in 1 hour.
    fmt.Println(minSessions([]int{3,1,3,1,1}, 8)) // 2
    // Example 3:
    // Input: tasks = [1,2,3,4,5], sessionTime = 15
    // Output: 1
    // Explanation: You can finish all the tasks in one work session.
    fmt.Println(minSessions([]int{1,2,3,4,5}, 15)) // 1

    fmt.Println(minSessions([]int{3,9}, 10)) // 2

    fmt.Println(minSessions1([]int{1,2,3}, 3)) // 2
    fmt.Println(minSessions1([]int{3,1,3,1,1}, 8)) // 2
    fmt.Println(minSessions1([]int{1,2,3,4,5}, 15)) // 1
    fmt.Println(minSessions1([]int{3,9}, 10)) // 2
}