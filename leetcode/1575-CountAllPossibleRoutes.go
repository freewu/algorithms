package main

// 1575. Count All Possible Routes
// You are given an array of distinct positive integers locations where locations[i] represents the position of city i.
// You are also given integers start, finish and fuel representing the starting city, ending city, 
// and the initial amount of fuel you have, respectively.

// At each step, if you are at city i, you can pick any city j such that j != i and 0 <= j < locations.length 
// and move to city j. Moving from city i to city j reduces the amount of fuel you have by |locations[i] - locations[j]|. 
// Please notice that |x| denotes the absolute value of x.

// Notice that fuel cannot become negative at any point in time, 
// and that you are allowed to visit any city more than once (including start and finish).

// Return the count of all possible routes from start to finish. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: locations = [2,3,6,8,4], start = 1, finish = 3, fuel = 5
// Output: 4
// Explanation: The following are all possible routes, each uses 5 units of fuel:
// 1 -> 3
// 1 -> 2 -> 3
// 1 -> 4 -> 3
// 1 -> 4 -> 2 -> 3

// Example 2:
// Input: locations = [4,3,1], start = 1, finish = 0, fuel = 6
// Output: 5
// Explanation: The following are all possible routes:
// 1 -> 0, used fuel = 1
// 1 -> 2 -> 0, used fuel = 5
// 1 -> 2 -> 1 -> 0, used fuel = 5
// 1 -> 0 -> 1 -> 0, used fuel = 3
// 1 -> 0 -> 1 -> 0 -> 1 -> 0, used fuel = 5

// Example 3:
// Input: locations = [5,2,1], start = 0, finish = 2, fuel = 3
// Output: 0
// Explanation: 
// It is impossible to get from 0 to 2 using only 3 units of fuel since the shortest route needs 4 units of fuel.

// Constraints:
//     2 <= locations.length <= 100
//     1 <= locations[i] <= 10^9
//     All integers in locations are distinct.
//     0 <= start, finish < locations.length
//     1 <= fuel <= 200

import "fmt"

// bfs
func countRoutes(locations []int, start int, finish int, fuel int) int {
    // this is BFS + caching, aka memorisation, aka DP
    n, mod := len(locations), int(1e9) + 7
    sum, dp, usages := 0, make([][]int, n), make([][]int, n)
    diff := func (a int, b int) int {
        if a - b < 0 {
            return b - a
        }
        return a - b
    }
    for i := 0; i < n; i++ {
        usage := make([]int, n)
        for j := 0; j < n; j++ {
            usage[j] = diff(locations[i], locations[j])
        }
        usages[i] = usage
    }

    for i := 0; i < n; i++ {
        fuelLeft := make([]int, fuel + 1)
        dp[i] = fuelLeft
    }
    dp[start][fuel]++
    for f := fuel; f >= 0; f-- {
        for l := 0; l < n; l++ {
            lf := dp[l][f]
            if lf > 0 {
                for i := 0; i < n; i++ { // for each usage in the usages map
                    if usages[l][i] > 0 && usages[l][i] <= f && l != i {
                        dp[i][f - usages[l][i]] += lf
                        if dp[i][f - usages[l][i]] > mod {
                            dp[i][f - usages[l][i]] = dp[i][f - usages[l][i]] % mod
                        }
                    }
                }
            }
        }
    }
    for i := 0; i <= fuel; i++ {
        sum += dp[finish][i]
    }
    return sum % mod
}

// dfs
func countRoutes1(locations []int, start int, finish int, fuel int) int {
    n, mod := len(locations), int(1e9) + 7
    states := make([]int, n * (fuel + 1))
    for i := range states {
        states[i] = -1
    }
    toState := func(curPos, curFuel int) int {
        return curPos + curFuel * n
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    var dfs func(curPos, curFuel int) int
    dfs = func(curPos, curFuel int) int {
        if states[toState(curPos, curFuel)] != -1 {
            return states[toState(curPos, curFuel)]
        }
        res := 0
        for nextPos := range locations {
            if nextPos != curPos {
                if abs(locations[curPos] - locations[nextPos]) <= curFuel {
                    res += dfs(nextPos, curFuel - abs(locations[curPos] - locations[nextPos]))
                }
            }
        }
        if curPos == finish {
            res += 1
        }
        states[toState(curPos, curFuel)] = res % mod
        return res % mod
    }
    dfs(start, fuel)
    return states[toState(start, fuel)]
}

func main() {
    // Example 1:
    // Input: locations = [2,3,6,8,4], start = 1, finish = 3, fuel = 5
    // Output: 4
    // Explanation: The following are all possible routes, each uses 5 units of fuel:
    // 1 -> 3
    // 1 -> 2 -> 3
    // 1 -> 4 -> 3
    // 1 -> 4 -> 2 -> 3
    fmt.Println(countRoutes([]int{2,3,6,8,4}, 1, 3, 5)) // 4
    // Example 2:
    // Input: locations = [4,3,1], start = 1, finish = 0, fuel = 6
    // Output: 5
    // Explanation: The following are all possible routes:
    // 1 -> 0, used fuel = 1
    // 1 -> 2 -> 0, used fuel = 5
    // 1 -> 2 -> 1 -> 0, used fuel = 5
    // 1 -> 0 -> 1 -> 0, used fuel = 3
    // 1 -> 0 -> 1 -> 0 -> 1 -> 0, used fuel = 5
    fmt.Println(countRoutes([]int{4,3,1}, 1, 0, 6)) // 5
    // Example 3:
    // Input: locations = [5,2,1], start = 0, finish = 2, fuel = 3
    // Output: 0
    // Explanation: 
    // It is impossible to get from 0 to 2 using only 3 units of fuel since the shortest route needs 4 units of fuel.
    fmt.Println(countRoutes([]int{5,2,1}, 0, 2, 3)) // 0

    fmt.Println(countRoutes1([]int{2,3,6,8,4}, 1, 3, 5)) // 4
    fmt.Println(countRoutes1([]int{4,3,1}, 1, 0, 6)) // 5
    fmt.Println(countRoutes1([]int{5,2,1}, 0, 2, 3)) // 0
}