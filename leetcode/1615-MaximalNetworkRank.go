package main

// 1615. Maximal Network Rank
// There is an infrastructure of n cities with some number of roads connecting these cities. 
// Each roads[i] = [ai, bi] indicates that there is a bidirectional road between cities ai and bi.

// The network rank of two different cities is defined as the total number of directly connected roads to either city. 
// If a road is directly connected to both cities, it is only counted once.

// The maximal network rank of the infrastructure is the maximum network rank of all pairs of different cities.
// Given the integer n and the array roads, return the maximal network rank of the entire infrastructure.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/21/ex1.png" />
// Input: n = 4, roads = [[0,1],[0,3],[1,2],[1,3]]
// Output: 4
// Explanation: The network rank of cities 0 and 1 is 4 as there are 4 roads that are connected to either 0 or 1. The road between 0 and 1 is only counted once.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/09/21/ex2.png" />
// Input: n = 5, roads = [[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]
// Output: 5
// Explanation: There are 5 roads that are connected to cities 1 or 2.

// Example 3:
// Input: n = 8, roads = [[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]
// Output: 5
// Explanation: The network rank of 2 and 5 is 5. Notice that all the cities do not have to be connected.

// Constraints:
//     2 <= n <= 100
//     0 <= roads.length <= n * (n - 1) / 2
//     roads[i].length == 2
//     0 <= ai, bi <= n-1
//     ai != bi
//     Each pair of cities has at most one road connecting them.

import "fmt"

func maximalNetworkRank(n int, roads [][]int) int {
    res, connected, count := 0, make([][]bool, n), make([]int, n)
    for i := 0; i < n; i++ {
        connected[i] = make([]bool, n)
    }
    for _, r := range roads {
        count[r[0]]++
        count[r[1]]++
        connected[r[0]][r[1]], connected[r[1]][r[0]] = true, true
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if connected[i][j] {
                res = max(res, count[i] + count[j] - 1)
            } else {
                res = max(res, count[i] + count[j])
            }
        }
    }
    return res
}

func maximalNetworkRank1(n int, roads [][]int) int {
    length := len(roads)
    groupA, groupB := make([]int, length), make([]int, length)
    for i := 0; i < length; i++ {
        groupA[i], groupB[i] = roads[i][0], roads[i][1]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    helper := func(A []int, B []int, n int) int {
        maxRank, connect, degree := 0, make([][]bool, n), make([]int, n)
        for i := 0; i < n; i++ {
            connect[i] = make([]bool, n)
        }
        for i := 0; i < len(A); i++ {
            connect[A[i]][B[i]], connect[B[i]][A[i]] = true, true
            degree[A[i]]++
            degree[B[i]]++
        }
        for i := 0; i < n; i++ {
            for j := i+1; j < n; j++ {
                cur := 0
                if connect[i][j] {
                    cur = degree[i] + degree[j] - 1
                } else {
                    cur = degree[i] + degree[j]
                }
                maxRank = max(cur, maxRank)
            }
        }
        return maxRank
    }
    return helper(groupA, groupB, n)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/21/ex1.png" />
    // Input: n = 4, roads = [[0,1],[0,3],[1,2],[1,3]]
    // Output: 4
    // Explanation: The network rank of cities 0 and 1 is 4 as there are 4 roads that are connected to either 0 or 1. The road between 0 and 1 is only counted once.
    fmt.Println(maximalNetworkRank(4, [][]int{{0,1},{0,3},{1,2},{1,3}})) // 4
    // Example 2: 
    // <img src="https://assets.leetcode.com/uploads/2020/09/21/ex2.png" />
    // Input: n = 5, roads = [[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]
    // Output: 5
    // Explanation: There are 5 roads that are connected to cities 1 or 2.
    fmt.Println(maximalNetworkRank(5, [][]int{{0,1},{0,3},{1,2},{1,3},{2,3},{2,4}})) // 5
    // Example 3:
    // Input: n = 8, roads = [[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]
    // Output: 5
    // Explanation: The network rank of 2 and 5 is 5. Notice that all the cities do not have to be connected.
    fmt.Println(maximalNetworkRank(8, [][]int{{0,1},{1,2},{2,3},{2,4},{5,6},{5,7}})) // 5

    fmt.Println(maximalNetworkRank1(4, [][]int{{0,1},{0,3},{1,2},{1,3}})) // 4
    fmt.Println(maximalNetworkRank1(5, [][]int{{0,1},{0,3},{1,2},{1,3},{2,3},{2,4}})) // 5
    fmt.Println(maximalNetworkRank1(8, [][]int{{0,1},{1,2},{2,3},{2,4},{5,6},{5,7}})) // 5
}