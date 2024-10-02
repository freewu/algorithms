package main

// 1928. Minimum Cost to Reach Destination in Time
// There is a country of n cities numbered from 0 to n - 1 where all the cities are connected by bi-directional roads. 
// The roads are represented as a 2D integer array edges where edges[i] = [xi, yi, timei] denotes a road between cities xi and yi that takes timei minutes to travel. 
// There may be multiple roads of differing travel times connecting the same two cities, but no road connects a city to itself.

// Each time you pass through a city, you must pay a passing fee. 
// This is represented as a 0-indexed integer array passingFees of length n where passingFees[j] is the amount of dollars you must pay when you pass through city j.

// In the beginning, you are at city 0 and want to reach city n - 1 in maxTime minutes or less. 
// The cost of your journey is the summation of passing fees for each city that you passed through at some moment of your journey (including the source and destination cities).

// Given maxTime, edges, and passingFees, return the minimum cost to complete your journey, or -1 if you cannot complete it within maxTime minutes.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/04/leetgraph1-1.png" />
// Input: maxTime = 30, edges = [[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]], passingFees = [5,1,2,20,20,3]
// Output: 11
// Explanation: The path to take is 0 -> 1 -> 2 -> 5, which takes 30 minutes and has $11 worth of passing fees.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/04/copy-of-leetgraph1-1.png" />
// Input: maxTime = 29, edges = [[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]], passingFees = [5,1,2,20,20,3]
// Output: 48
// Explanation: The path to take is 0 -> 3 -> 4 -> 5, which takes 26 minutes and has $48 worth of passing fees.
// You cannot take path 0 -> 1 -> 2 -> 5 since it would take too long.

// Example 3:
// Input: maxTime = 25, edges = [[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]], passingFees = [5,1,2,20,20,3]
// Output: -1
// Explanation: There is no way to reach city 5 from city 0 within 25 minutes.

// Constraints:
//     1 <= maxTime <= 1000
//     n == passingFees.length
//     2 <= n <= 1000
//     n - 1 <= edges.length <= 1000
//     0 <= xi, yi <= n - 1
//     1 <= timei <= 1000
//     1 <= passingFees[j] <= 1000 
//     The graph may contain multiple edges between two nodes.
//     The graph does not contain self loops.

import "fmt"

func minCost(maxTime int, edges [][]int, passingFees []int) int {
    n, inf := len(passingFees), 1 << 31
    res, dp := inf, make([][]int, maxTime + 1)
    for i := range dp { // // initialise the dp with infinity
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = inf
        }
    }
    dp[0][0] = passingFees[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for currTime := 0; currTime <= maxTime; currTime++ {
        for _, edge := range edges {  // loop through all the edges
            reachTime := currTime + edge[2] // if time to reach the next node is > maxTime, then just skip this edge
            // the edges are bidirectional, so we have to handle both directions
            for i := 0 ; reachTime <= maxTime && i <= 1; i++ {
                fromNode, toNode := edge[int(i^0)], edge[int(i^1)] // intelligent way to swap the fromNode and toNode
                if dp[currTime][fromNode] != inf && reachTime <= maxTime {
                    dp[reachTime][toNode] = min(dp[reachTime][toNode], dp[currTime][fromNode] + passingFees[toNode])
                }      
            } 
        } 
    }
    for i := 0; i <= maxTime; i++ {
        res = min(res, dp[i][n-1]) // find the minCost to reach lastNode
    }
    if res == inf { return -1 } // 无解
    return res
}

func minCost1(maxTime int, edges [][]int, passingFees []int) int {
    n, inf := len(passingFees), 1 << 31
    dp := make([][]int, maxTime + 1)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = inf
       }
    }
    dp[0][0] = passingFees[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for t := 1; t <= maxTime; t++ {
        for _, edge := range edges {
            i, j, cost := edge[0], edge[1], edge[2]
            if cost <= t {
                if dp[t - cost][j] != inf {
                    dp[t][i] = min(dp[t][i], dp[t - cost][j] + passingFees[i])
                }
                if dp[t - cost][i] != inf {
                    dp[t][j] = min(dp[t][j], dp[t - cost][i] + passingFees[j])
                }
           }
       }
   }
   res := inf
   for t := 1; t <= maxTime; t++ {
       res = min(res, dp[t][n - 1])
   }
   if res == inf { return -1 }
   return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/04/leetgraph1-1.png" />
    // Input: maxTime = 30, edges = [[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]], passingFees = [5,1,2,20,20,3]
    // Output: 11
    // Explanation: The path to take is 0 -> 1 -> 2 -> 5, which takes 30 minutes and has $11 worth of passing fees.
    fmt.Println(minCost(30, [][]int{{0,1,10},{1,2,10},{2,5,10},{0,3,1},{3,4,10},{4,5,15}}, []int{5,1,2,20,20,3})) // 11
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/04/copy-of-leetgraph1-1.png" />
    // Input: maxTime = 29, edges = [[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]], passingFees = [5,1,2,20,20,3]
    // Output: 48
    // Explanation: The path to take is 0 -> 3 -> 4 -> 5, which takes 26 minutes and has $48 worth of passing fees.
    // You cannot take path 0 -> 1 -> 2 -> 5 since it would take too long.
    fmt.Println(minCost(29, [][]int{{0,1,10},{1,2,10},{2,5,10},{0,3,1},{3,4,10},{4,5,15}}, []int{5,1,2,20,20,3})) // 48
    // Example 3:
    // Input: maxTime = 25, edges = [[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]], passingFees = [5,1,2,20,20,3]
    // Output: -1
    // Explanation: There is no way to reach city 5 from city 0 within 25 minutes.
    fmt.Println(minCost(25, [][]int{{0,1,10},{1,2,10},{2,5,10},{0,3,1},{3,4,10},{4,5,15}}, []int{5,1,2,20,20,3})) // -1

    fmt.Println(minCost1(30, [][]int{{0,1,10},{1,2,10},{2,5,10},{0,3,1},{3,4,10},{4,5,15}}, []int{5,1,2,20,20,3})) // 11
    fmt.Println(minCost1(29, [][]int{{0,1,10},{1,2,10},{2,5,10},{0,3,1},{3,4,10},{4,5,15}}, []int{5,1,2,20,20,3})) // 48
    fmt.Println(minCost1(25, [][]int{{0,1,10},{1,2,10},{2,5,10},{0,3,1},{3,4,10},{4,5,15}}, []int{5,1,2,20,20,3})) // -1
}