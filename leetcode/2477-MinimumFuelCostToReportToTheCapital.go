package main

// 2477. Minimum Fuel Cost to Report to the Capital
// There is a tree (i.e., a connected, undirected graph with no cycles) structure country network consisting of n cities numbered from 0 to n - 1 and exactly n - 1 roads. 
// The capital city is city 0. You are given a 2D integer array roads where roads[i] = [ai, bi] denotes that there exists a bidirectional road connecting cities ai and bi.

// There is a meeting for the representatives of each city. 
// The meeting is in the capital city.

// There is a car in each city. 
// You are given an integer seats that indicates the number of seats in each car.

// A representative can use the car in their city to travel or change the car and ride with another representative. 
// The cost of traveling between two cities is one liter of fuel.

// Return the minimum number of liters of fuel to reach the capital city.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/09/22/a4c380025e3ff0c379525e96a7d63a3.png" />
// Input: roads = [[0,1],[0,2],[0,3]], seats = 5
// Output: 3
// Explanation: 
// - Representative1 goes directly to the capital with 1 liter of fuel.
// - Representative2 goes directly to the capital with 1 liter of fuel.
// - Representative3 goes directly to the capital with 1 liter of fuel.
// It costs 3 liters of fuel at minimum. 
// It can be proven that 3 is the minimum number of liters of fuel needed.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/11/16/2.png" />
// Input: roads = [[3,1],[3,2],[1,0],[0,4],[0,5],[4,6]], seats = 2
// Output: 7
// Explanation: 
// - Representative2 goes directly to city 3 with 1 liter of fuel.
// - Representative2 and representative3 go together to city 1 with 1 liter of fuel.
// - Representative2 and representative3 go together to the capital with 1 liter of fuel.
// - Representative1 goes directly to the capital with 1 liter of fuel.
// - Representative5 goes directly to the capital with 1 liter of fuel.
// - Representative6 goes directly to city 4 with 1 liter of fuel.
// - Representative4 and representative6 go together to the capital with 1 liter of fuel.
// It costs 7 liters of fuel at minimum. 
// It can be proven that 7 is the minimum number of liters of fuel needed.

// Example 3:
// Input: roads = [], seats = 1
// Output: 0
// Explanation: No representatives need to travel to the capital city.

// Constraints:
//     1 <= n <= 10^5
//     roads.length == n - 1
//     roads[i].length == 2
//     0 <= ai, bi < n
//     ai != bi
//     roads represents a valid tree.
//     1 <= seats <= 10^5

import "fmt"
import "math"

// bfs
func minimumFuelCost(roads [][]int, seats int) int64 {
    res, n := 0, len(roads) + 1
    adj, inDegree := make([][]int, n), make([]int, n)
    for _, road := range roads {
        adj[road[0]] = append(adj[road[0]], road[1])
        adj[road[1]] = append(adj[road[1]], road[0])
        inDegree[road[0]]++
        inDegree[road[1]]++
    }
    reps, queue := make([]int, n), []int{}
    // Don't check 0
    for i := 1; i < n; i++ {
        if inDegree[i] == 1 {
            queue = append(queue, i)
        }
        reps[i] = 1
    }
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        res += int(math.Ceil(float64(reps[node]) / float64(seats)))
        for _, next := range adj[node] {
            inDegree[next]--
            reps[next] += reps[node]
            if inDegree[next] == 1 && next != 0 {
                queue = append(queue, next)
            }
        }
    }
    return int64(res)
}

// dfs
func minimumFuelCost1(roads [][]int, seats int) int64 {
    res, n := 0, len(roads) + 1
    route := make([][]int, n)
    for _, v := range roads {
        route[v[0]] = append(route[v[0]], v[1])
        route[v[1]] = append(route[v[1]], v[0])
    }
    var dfs func(last, index int) int
    dfs = func(last, index int) int {
        val := 0
        for _, i := range route[index] {
            if i == last { continue }
            val += dfs(index, i)
        }
        res += (val / seats + 1)
        return val + 1
    }
    for _, v := range route[0] {
        dfs(0, v)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/09/22/a4c380025e3ff0c379525e96a7d63a3.png" />
    // Input: roads = [[0,1],[0,2],[0,3]], seats = 5
    // Output: 3
    // Explanation: 
    // - Representative1 goes directly to the capital with 1 liter of fuel.
    // - Representative2 goes directly to the capital with 1 liter of fuel.
    // - Representative3 goes directly to the capital with 1 liter of fuel.
    // It costs 3 liters of fuel at minimum. 
    // It can be proven that 3 is the minimum number of liters of fuel needed.
    fmt.Println(minimumFuelCost([][]int{{0,1},{0,2},{0,3}}, 5)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/11/16/2.png" />
    // Input: roads = [[3,1],[3,2],[1,0],[0,4],[0,5],[4,6]], seats = 2
    // Output: 7
    // Explanation: 
    // - Representative2 goes directly to city 3 with 1 liter of fuel.
    // - Representative2 and representative3 go together to city 1 with 1 liter of fuel.
    // - Representative2 and representative3 go together to the capital with 1 liter of fuel.
    // - Representative1 goes directly to the capital with 1 liter of fuel.
    // - Representative5 goes directly to the capital with 1 liter of fuel.
    // - Representative6 goes directly to city 4 with 1 liter of fuel.
    // - Representative4 and representative6 go together to the capital with 1 liter of fuel.
    // It costs 7 liters of fuel at minimum. 
    // It can be proven that 7 is the minimum number of liters of fuel needed.
    fmt.Println(minimumFuelCost([][]int{{3,1},{3,2},{1,0},{0,4},{0,5},{4,6}}, 2)) // 7
    // Example 3:
    // Input: roads = [], seats = 1
    // Output: 0
    // Explanation: No representatives need to travel to the capital city.
    fmt.Println(minimumFuelCost([][]int{}, 1)) // 0

    fmt.Println(minimumFuelCost1([][]int{{0,1},{0,2},{0,3}}, 5)) // 3
    fmt.Println(minimumFuelCost1([][]int{{3,1},{3,2},{1,0},{0,4},{0,5},{4,6}}, 2)) // 7
    fmt.Println(minimumFuelCost1([][]int{}, 1)) // 0
}