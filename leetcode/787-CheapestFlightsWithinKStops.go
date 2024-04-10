package main

// 787. Cheapest Flights Within K Stops
// There are n cities connected by some number of flights. 
// You are given an array flights where flights[i] = [fromi, toi, pricei] indicates that there is a flight from city fromi to city toi with cost pricei.

// You are also given three integers src, dst, and k, return the cheapest price from src to dst with at most k stops. 
// If there is no such route, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-3drawio.png" / >
// Input: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1
// Output: 700
// Explanation:
// The graph is shown above.
// The optimal path with at most 1 stop from city 0 to 3 is marked in red and has cost 100 + 600 = 700.
// Note that the path through cities [0,1,2,3] is cheaper but is invalid because it uses 2 stops.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-1drawio.png" / >
// Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
// Output: 200
// Explanation:
// The graph is shown above.
// The optimal path with at most 1 stop from city 0 to 2 is marked in red and has cost 100 + 100 = 200.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-2drawio.png" / >
// Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0
// Output: 500
// Explanation:
// The graph is shown above.
// The optimal path with no stops from city 0 to 2 is marked in red and has cost 500.

// Constraints:
//         1 <= n <= 100
//         0 <= flights.length <= (n * (n - 1) / 2)
//         flights[i].length == 3
//         0 <= fromi, toi < n
//         fromi != toi
//         1 <= pricei <= 10^4
//         There will not be any multiple flights between two cities.
//         0 <= src, dst, k < n
//         src != dst

import "fmt"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    min := func (a, b int) int {
        if a < b {
            return a
        }
        return b
    }
    // 航班的花费不超过 10^4 最多搭乘航班的次数 k+1 不超过 101
    const inf = 10000*101 + 1
    f := make([]int, n)
    for i := range f {
        f[i] = inf
    }
    f[src] = 0
    ans := inf
    for t := 1; t <= k+1; t++ {
        g := make([]int, n)
        for i := range g {
            g[i] = inf
        }
        for _, flight := range flights {
            j, i, cost := flight[0], flight[1], flight[2]
            g[i] = min(g[i], f[j]+cost)
        }
        f = g
        ans = min(ans, f[dst])
    }
    if ans == inf {
        ans = -1
    }
    return ans
}

func findCheapestPrice1(n int, flights [][]int, src int, dst int, k int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    // 航班的花费不超过 10^4 最多搭乘航班的次数 k+1 不超过 101
    const inf = 10000*101 + 1
    dp := make([][]int, k+2)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = inf
        }
    }
    dp[0][src] = 0
    for t := 1; t <= k+1; t++ {
        for _, flight := range flights {
            j, i, cost := flight[0], flight[1], flight[2]
            dp[t][i] = min(dp[t][i], dp[t-1][j]+cost)
        }
    }
    res := inf
    for t := 1; t <= k+1; t++ {
        res = min(res, dp[t][dst])
    }
    if res == inf {
        res = -1
    }
    return res
}

func main() {

    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-3drawio.png" / >
    // Input: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1
    // Output: 700
    // Explanation:
    // The graph is shown above.
    // The optimal path with at most 1 stop from city 0 to 3 is marked in red and has cost 100 + 600 = 700.
    // Note that the path through cities [0,1,2,3] is cheaper but is invalid because it uses 2 stops.
    fmt.Println(findCheapestPrice(
        4,
        [][]int{[]int{0,1,100},[]int{1,2,100},[]int{2,0,100},[]int{1,3,600},[]int{2,3,200}}, // flights
        0, // src
        3, // dst
        1, // k
    )) // 700

    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-1drawio.png" / >
    // Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
    // Output: 200
    // Explanation:
    // The graph is shown above.
    // The optimal path with at most 1 stop from city 0 to 2 is marked in red and has cost 100 + 100 = 200.
    fmt.Println(findCheapestPrice(
        3,
        [][]int{[]int{0,1,100},[]int{1,2,100},[]int{0,2,100}}, // flights
        0, // src
        2, // dst
        1, // k
    )) // 200

    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-2drawio.png" / >
    // Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0
    // Output: 500
    fmt.Println(findCheapestPrice(
        3,
        [][]int{[]int{0,1,100},[]int{1,2,100},[]int{0,2,500}}, // flights
        0, // src
        2, // dst
        0, // k
    )) // 500



    fmt.Println(findCheapestPrice1(
        4,
        [][]int{[]int{0,1,100},[]int{1,2,100},[]int{2,0,100},[]int{1,3,600},[]int{2,3,200}}, // flights
        0, // src
        3, // dst
        1, // k
    )) // 700
    
    fmt.Println(findCheapestPrice1(
        3,
        [][]int{[]int{0,1,100},[]int{1,2,100},[]int{0,2,100}}, // flights
        0, // src
        2, // dst
        1, // k
    )) // 200

    fmt.Println(findCheapestPrice1(
        3,
        [][]int{[]int{0,1,100},[]int{1,2,100},[]int{0,2,500}}, // flights
        0, // src
        2, // dst
        0, // k
    )) // 500
}