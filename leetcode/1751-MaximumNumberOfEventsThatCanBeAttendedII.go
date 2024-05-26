package main

// 1751. Maximum Number of Events That Can Be Attended II
// You are given an array of events where events[i] = [startDayi, endDayi, valuei]. 
// The ith event starts at startDayi and ends at endDayi, and if you attend this event, you will receive a value of valuei. 
// You are also given an integer k which represents the maximum number of events you can attend.

// You can only attend one event at a time. 
// If you choose to attend an event, you must attend the entire event. 
// Note that the end day is inclusive: 
//     that is, you cannot attend two events where one of them starts and the other ends on the same day.

// Return the maximum sum of values that you can receive by attending events.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/10/screenshot-2021-01-11-at-60048-pm.png" />
// Input: events = [[1,2,4],[3,4,3],[2,3,1]], k = 2
// Output: 7
// Explanation: Choose the green events, 0 and 1 (0-indexed) for a total value of 4 + 3 = 7.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/01/10/screenshot-2021-01-11-at-60150-pm.png" />
// Input: events = [[1,2,4],[3,4,3],[2,3,10]], k = 2
// Output: 10
// Explanation: Choose event 2 for a total value of 10.
// Notice that you cannot attend any other event as they overlap, and that you do not have to attend k events.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/01/10/screenshot-2021-01-11-at-60703-pm.png" />
// Input: events = [[1,1,1],[2,2,2],[3,3,3],[4,4,4]], k = 3
// Output: 9
// Explanation: Although the events do not overlap, you can only attend 3 events. Pick the highest valued three.

// Constraints:
//     1 <= k <= events.length
//     1 <= k * events.length <= 10^6
//     1 <= startDayi <= endDayi <= 10^9
//     1 <= valuei <= 10^6

import "fmt"
import "sort"
import "slices"

func maxValue(events [][]int, k int) int {
    n := len(events)
    sort.Slice(events, func(i, j int) bool {
        return events[i][0] < events[j][0]
    })
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, k + 1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        e, x := events[i][1], events[i][2]
        next := sort.Search(n, func(i int) bool {
            return events[i][0] > e
        })
        for j := 1; j <= k; j++ {
            dp[i][j] = max(dp[i+1][j], dp[next][j-1]+x)
        }
    }
    return dp[0][k]
}

func maxValue1(events [][]int, k int) int {
    slices.SortFunc(events, func(a, b []int) int { 
        return a[1] - b[1] 
    })
    n := len(events)
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, k+1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, e := range events {
        p := sort.Search(i, func(j int) bool { 
            return events[j][1] >= e[0] 
        })
        for j := 1; j <= k; j++ {
            dp[i+1][j] = max(dp[i][j], dp[p][j-1]+e[2])
        }
    }
    return dp[n][k]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/01/10/screenshot-2021-01-11-at-60048-pm.png" />
    // Input: events = [[1,2,4],[3,4,3],[2,3,1]], k = 2
    // Output: 7
    // Explanation: Choose the green events, 0 and 1 (0-indexed) for a total value of 4 + 3 = 7.
    fmt.Println(maxValue([][]int{{1,2,4},{3,4,3},{2,3,1}}, 2)) // 7
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/01/10/screenshot-2021-01-11-at-60150-pm.png" />
    // Input: events = [[1,2,4],[3,4,3],[2,3,10]], k = 2
    // Output: 10
    // Explanation: Choose event 2 for a total value of 10.
    // Notice that you cannot attend any other event as they overlap, and that you do not have to attend k events.
    fmt.Println(maxValue([][]int{{1,2,4},{3,4,3},{2,3,10}}, 2)) // 10
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/01/10/screenshot-2021-01-11-at-60703-pm.png" />
    // Input: events = [[1,1,1],[2,2,2],[3,3,3],[4,4,4]], k = 3
    // Output: 9
    // Explanation: Although the events do not overlap, you can only attend 3 events. Pick the highest valued three.
    fmt.Println(maxValue([][]int{{1,1,1},{2,2,2},{3,3,3},{4,4,4}}, 3)) // 9

    fmt.Println(maxValue1([][]int{{1,2,4},{3,4,3},{2,3,1}}, 2)) // 7
    fmt.Println(maxValue1([][]int{{1,2,4},{3,4,3},{2,3,10}}, 2)) // 10
    fmt.Println(maxValue1([][]int{{1,1,1},{2,2,2},{3,3,3},{4,4,4}}, 3)) // 9
}