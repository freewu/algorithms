package main

// 2008. Maximum Earnings From Taxi
// There are n points on a road you are driving your taxi on. 
// The n points on the road are labeled from 1 to n in the direction you are going, 
// and you want to drive from point 1 to point n to make money by picking up passengers. 
// You cannot change the direction of the taxi.

// The passengers are represented by a 0-indexed 2D integer array rides, 
// where rides[i] = [starti, endi, tipi] denotes the ith passenger requesting a ride from point starti to point endi who is willing to give a tipi dollar tip.

// For each passenger i you pick up, you earn endi - starti + tipi dollars. 
// You may only drive at most one passenger at a time.

// Given n and rides, return the maximum number of dollars you can earn by picking up the passengers optimally.

// Note: You may drop off a passenger and pick up a different passenger at the same point.

// Example 1:
// Input: n = 5, rides = [[2,5,4],[1,5,1]]
// Output: 7
// Explanation: We can pick up passenger 0 to earn 5 - 2 + 4 = 7 dollars.

// Example 2:
// Input: n = 20, rides = [[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]
// Output: 20
// Explanation: We will pick up the following passengers:
// - Drive passenger 1 from point 3 to point 10 for a profit of 10 - 3 + 2 = 9 dollars.
// - Drive passenger 2 from point 10 to point 12 for a profit of 12 - 10 + 3 = 5 dollars.
// - Drive passenger 5 from point 13 to point 18 for a profit of 18 - 13 + 1 = 6 dollars.
// We earn 9 + 5 + 6 = 20 dollars in total.

// Constraints:
//     1 <= n <= 10^5
//     1 <= rides.length <= 3 * 10^4
//     rides[i].length == 3
//     1 <= starti < endi <= n
//     1 <= tipi <= 10^5

import "fmt"
import "sort"

func maxTaxiEarnings(n int, rides [][]int) int64 {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    sort.Slice(rides, func(i, j int) bool {
        if rides[i][0] != rides[j][0] { return rides[i][0] < rides[j][0] } 
        if rides[i][1] != rides[j][1] { return rides[i][1] < rides[j][1] } 
        return rides[i][2] < rides[j][2]
    })
    maxProfit := make([]int, n + 1)
    res, index := 0, 0
    for i := 1; i <= n; i++ {
        // max profit till current point is maximum of profit till previous point and current point
        maxProfit[i] = max(maxProfit[i], maxProfit[i - 1])
        for index < len(rides) && rides[index][0] <= i { // process rides which start before current point
            start, end, tip := rides[index][0], rides[index][1], rides[index][2]
            dollars := end - start + tip
            maxProfit[end] = max(maxProfit[end], maxProfit[start] + dollars) // profit of a ride is included only when its completed 
            index++
        }
        res = max(res, maxProfit[i])
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: n = 5, rides = [[2,5,4],[1,5,1]]
    // Output: 7
    // Explanation: We can pick up passenger 0 to earn 5 - 2 + 4 = 7 dollars.
    fmt.Println(maxTaxiEarnings(5, [][]int{{2,5,4},{1,5,1}})) // 7
    // Example 2:
    // Input: n = 20, rides = [[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]
    // Output: 20
    // Explanation: We will pick up the following passengers:
    // - Drive passenger 1 from point 3 to point 10 for a profit of 10 - 3 + 2 = 9 dollars.
    // - Drive passenger 2 from point 10 to point 12 for a profit of 12 - 10 + 3 = 5 dollars.
    // - Drive passenger 5 from point 13 to point 18 for a profit of 18 - 13 + 1 = 6 dollars.
    // We earn 9 + 5 + 6 = 20 dollars in total.
    fmt.Println(maxTaxiEarnings(20, [][]int{{1,6,1},{3,10,2},{10,12,3},{11,12,2},{12,15,2},{13,18,1}})) // 20
}