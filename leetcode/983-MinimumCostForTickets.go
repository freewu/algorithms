package main

// 983. Minimum Cost For Tickets
// You have planned some train traveling one year in advance. 
// The days of the year in which you will travel are given as an integer array days. 
// Each day is an integer from 1 to 365.

// Train tickets are sold in three different ways:
//     a 1-day pass is sold for costs[0] dollars,
//     a 7-day pass is sold for costs[1] dollars, and
//     a 30-day pass is sold for costs[2] dollars.

// The passes allow that many days of consecutive travel.
//      For example, if we get a 7-day pass on day 2, then we can travel for 7 days: 2, 3, 4, 5, 6, 7, and 8.

// Return the minimum number of dollars you need to travel every day in the given list of days.

// Example 1:
// Input: days = [1,4,6,7,8,20], costs = [2,7,15]
// Output: 11
// Explanation: For example, here is one way to buy passes that lets you travel your travel plan:
// On day 1, you bought a 1-day pass for costs[0] = $2, which covered day 1.
// On day 3, you bought a 7-day pass for costs[1] = $7, which covered days 3, 4, ..., 9.
// On day 20, you bought a 1-day pass for costs[0] = $2, which covered day 20.
// In total, you spent $11 and covered all the days of your travel.

// Example 2:
// Input: days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]
// Output: 17
// Explanation: For example, here is one way to buy passes that lets you travel your travel plan:
// On day 1, you bought a 30-day pass for costs[2] = $15 which covered days 1, 2, ..., 30.
// On day 31, you bought a 1-day pass for costs[0] = $2 which covered day 31.
// In total, you spent $17 and covered all the days of your travel.
 
// Constraints:
//     1 <= days.length <= 365
//     1 <= days[i] <= 365
//     days is in strictly increasing order.
//     costs.length == 3
//     1 <= costs[i] <= 1000

import "fmt"

// dp
func mincostTickets(days []int, costs []int) int {
    if len(days) == 0 {
        return 0
    }
    magic := make([]int, 365 + 1 + 30)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, day := range days {
        c1 := magic[day + 30 - 1]  + costs[0]
        c2 := magic[day + 30 - 7]  + costs[1]
        c3 := magic[day + 30 - 30] + costs[2]
        c := min(c1, min(c2, c3))
        for i := day + 30; i < len(magic); i++ {
            magic[i] = c
        }
    }
    return magic[days[len(days)-1] + 30]
}

func mincostTickets1(days []int, costs []int) int {
    n := len(days)
    dp := make([]int, n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    dp[n-1] = min(costs[0], min(costs[1], costs[2]))
    for i:= n - 2; i >= 0; i-- {
        dp[i] = costs[0] + dp[i+1]
        next7 := 0
        for j := i + 1; j < n; j++ {
            if days[j] - days[i] >= 7 {
                next7 = dp[j]
                break
            }
        }
        dp[i] = min(dp[i], costs[1] + next7)
        next30 := 0
        for j:=i+1; j<n; j++ {
            if days[j] - days[i] >= 30 {
                next30 = dp[j]
                break
            }
        }
        dp[i] = min(dp[i], costs[2] + next30)
    }
    return dp[0]
}

func main() {
    // Explanation: For example, here is one way to buy passes that lets you travel your travel plan:
    // On day 1, you bought a 1-day pass for costs[0] = $2, which covered day 1.
    // On day 3, you bought a 7-day pass for costs[1] = $7, which covered days 3, 4, ..., 9.
    // On day 20, you bought a 1-day pass for costs[0] = $2, which covered day 20.
    // In total, you spent $11 and covered all the days of your travel.
    fmt.Println(mincostTickets([]int{1,4,6,7,8,20},[]int{2,7,15})) // 11
    // Explanation: For example, here is one way to buy passes that lets you travel your travel plan:
    // On day 1, you bought a 30-day pass for costs[2] = $15 which covered days 1, 2, ..., 30.
    // On day 31, you bought a 1-day pass for costs[0] = $2 which covered day 31.
    // In total, you spent $17 and covered all the days of your travel.
    fmt.Println(mincostTickets([]int{1,2,3,4,5,6,7,8,9,10,30,31},[]int{2,7,15})) // 17

    fmt.Println(mincostTickets1([]int{1,4,6,7,8,20},[]int{2,7,15})) // 11
    fmt.Println(mincostTickets1([]int{1,2,3,4,5,6,7,8,9,10,30,31},[]int{2,7,15})) // 17
}