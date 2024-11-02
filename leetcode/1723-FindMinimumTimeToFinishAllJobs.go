package main

// 1723. Find Minimum Time to Finish All Jobs
// You are given an integer array jobs, where jobs[i] is the amount of time it takes to complete the ith job.

// There are k workers that you can assign jobs to. 
// Each job should be assigned to exactly one worker. 
// The working time of a worker is the sum of the time it takes to complete all jobs assigned to them. 
// Your goal is to devise an optimal assignment such that the maximum working time of any worker is minimized.

// Return the minimum possible maximum working time of any assignment.

// Example 1:
// Input: jobs = [3,2,3], k = 3
// Output: 3
// Explanation: By assigning each person one job, the maximum time is 3.

// Example 2:
// Input: jobs = [1,2,4,7,8], k = 2
// Output: 11
// Explanation: Assign the jobs the following way:
// Worker 1: 1, 2, 8 (working time = 1 + 2 + 8 = 11)
// Worker 2: 4, 7 (working time = 4 + 7 = 11)
// The maximum working time is 11.

// Constraints:
//     1 <= k <= jobs.length <= 12
//     1 <= jobs[i] <= 10^7

import "fmt"
import "sort"
import "math/bits"

// backtracking
func minimumTimeRequired(jobs []int, k int) int {
    workerTimes, n, res := make([]int, k), len(jobs), 120_000_001 // all jobs max length to one worker
    sort.Slice(jobs, func(i, j int) bool { 
        return jobs[i] > jobs[j] 
    })
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(index int, workerTimes []int, maxOnPass int) 
    dfs = func (index int, workerTimes []int, maxOnPass int) {
        if index >= n || maxOnPass >= res {
            res = min(res, maxOnPass)
            return
        }
        // otherwise there is more work to do, including a backtrack
        for worker, v := range workerTimes {
            workerTimes[worker] = v + jobs[index]
            dfs(index + 1, workerTimes, max(maxOnPass, workerTimes[worker]))
            workerTimes[worker] -= jobs[index] // backtrack
        }
    }
    dfs(0, workerTimes, 0)
    return res
}

// dp
func minimumTimeRequired1(jobs []int, k int) int {
    n, inf := len(jobs), 1 << 31
    // dp(S, k) = min( dp(S, k), dp(S|sub, k-1)+sum(sub) )
    sum := make([]int, (1 << n))
    for i := 1; i < (1 << n); i++ {
        x := bits.TrailingZeros(uint(i))
        y := i ^ (1 << x)
        sum[i] = sum[y] + jobs[x]
    }
    dp := make([][]int, k)
    for i := range dp {
        dp[i] = make([]int, (1 << n))
    }
    for i, s := range sum {
        dp[0][i] = s
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < k; i++ {
        for j := 0; j < (1 << n); j++ {
            tem := inf
            for x := j; x > 0; x = (x - 1) & j {
                tem = min(tem, max(dp[i-1][j-x], sum[x]))
            }
            dp[i][j] = tem
        }
    }
    return dp[k-1][(1 << n)-1]
}

func main() {
    // Example 1:
    // Input: jobs = [3,2,3], k = 3
    // Output: 3
    // Explanation: By assigning each person one job, the maximum time is 3.
    fmt.Println(minimumTimeRequired([]int{3,2,3}, 3)) // 3
    // Example 2:
    // Input: jobs = [1,2,4,7,8], k = 2
    // Output: 11
    // Explanation: Assign the jobs the following way:
    // Worker 1: 1, 2, 8 (working time = 1 + 2 + 8 = 11)
    // Worker 2: 4, 7 (working time = 4 + 7 = 11)
    // The maximum working time is 11.
    fmt.Println(minimumTimeRequired([]int{1,2,4,7,8}, 2)) // 11

    fmt.Println(minimumTimeRequired1([]int{3,2,3}, 3)) // 3
    fmt.Println(minimumTimeRequired1([]int{1,2,4,7,8}, 2)) // 11
}