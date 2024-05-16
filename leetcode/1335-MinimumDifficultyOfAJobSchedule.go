package main

// 1335. Minimum Difficulty of a Job Schedule
// You want to schedule a list of jobs in d days. 
// Jobs are dependent (i.e To work on the ith job, you have to finish all the jobs j where 0 <= j < i).

// You have to finish at least one task every day. 
// The difficulty of a job schedule is the sum of difficulties of each day of the d days. 
// The difficulty of a day is the maximum difficulty of a job done on that day.

// You are given an integer array jobDifficulty and an integer d. 
// The difficulty of the ith job is jobDifficulty[i].

// Return the minimum difficulty of a job schedule. 
// If you cannot find a schedule for the jobs return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/01/16/untitled.png" />
// Input: jobDifficulty = [6,5,4,3,2,1], d = 2
// Output: 7
// Explanation: First day you can finish the first 5 jobs, total difficulty = 6.
// Second day you can finish the last job, total difficulty = 1.
// The difficulty of the schedule = 6 + 1 = 7 

// Example 2:
// Input: jobDifficulty = [9,9,9], d = 4
// Output: -1
// Explanation: If you finish a job per day you will still have a free day. you cannot find a schedule for the given jobs.

// Example 3:
// Input: jobDifficulty = [1,1,1], d = 3
// Output: 3
// Explanation: The schedule is one job per day. total difficulty will be 3.

// Constraints:
//     1 <= jobDifficulty.length <= 300
//     0 <= jobDifficulty[i] <= 1000
//     1 <= d <= 10

import "fmt"

func minDifficulty(jobDifficulty []int, d int) int {
    n := len(jobDifficulty)
    if n < d { // 任务比工作天数据还少,无法制定工作计划
        return -1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp, inf := make([]int, n), 1 << 32 -1
    dp[0] = jobDifficulty[0]
    for i := 1; i < n; i++ {
        dp[i] = max(dp[i-1], jobDifficulty[i])
    }
    for i := 1; i < d; i++ {
        for j := n - 1; j >= i; j-- {
            dp[j] = inf
            mx := 0
            for k := j; k >= i; k-- {
                mx = max(mx, jobDifficulty[k])
                dp[j] = min(dp[j], dp[k-1] + mx)
            }
        }
    }
    return dp[n-1]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/01/16/untitled.png" />
    // Input: jobDifficulty = [6,5,4,3,2,1], d = 2
    // Output: 7
    // Explanation: First day you can finish the first 5 jobs, total difficulty = 6.
    // Second day you can finish the last job, total difficulty = 1.
    // The difficulty of the schedule = 6 + 1 = 7 
    fmt.Println(minDifficulty([]int{6,5,4,3,2,1}, 2)) // 7
    // Example 2:
    // Input: jobDifficulty = [9,9,9], d = 4
    // Output: -1
    // Explanation: If you finish a job per day you will still have a free day. you cannot find a schedule for the given jobs.
    fmt.Println(minDifficulty([]int{9,9,9}, 4)) // -1
    // Example 3:
    // Input: jobDifficulty = [1,1,1], d = 3
    // Output: 3
    // Explanation: The schedule is one job per day. total difficulty will be 3.
    fmt.Println(minDifficulty([]int{1,1,1}, 3)) // 3
}