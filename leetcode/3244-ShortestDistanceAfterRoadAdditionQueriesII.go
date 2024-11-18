package main

// 3244. Shortest Distance After Road Addition Queries II
// You are given an integer n and a 2D integer array queries.

// There are n cities numbered from 0 to n - 1. 
// Initially, there is a unidirectional road from city i to city i + 1 for all 0 <= i < n - 1.

// queries[i] = [ui, vi] represents the addition of a new unidirectional road from city ui to city vi. 
// After each query, you need to find the length of the shortest path from city 0 to city n - 1.

// Return an array answer where for each i in the range [0, queries.length - 1], 
// answer[i] is the length of the shortest path from city 0 to city n - 1 after processing the first i + 1 queries.

// Example 1:
// Input: n = 5, queries = [[2,4],[0,2],[0,4]]
// Output: [3,2,1]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/28/image8.jpg" />
// After the addition of the road from 2 to 4, the length of the shortest path from 0 to 4 is 3.
// <img src="https://assets.leetcode.com/uploads/2024/06/28/image9.jpg" />
// After the addition of the road from 0 to 2, the length of the shortest path from 0 to 4 is 2.
// <img src="https://assets.leetcode.com/uploads/2024/06/28/image10.jpg" />
// After the addition of the road from 0 to 4, the length of the shortest path from 0 to 4 is 1.

// Example 2:
// Input: n = 4, queries = [[0,3],[0,2]]
// Output: [1,1]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/28/image11.jpg" />
// After the addition of the road from 0 to 3, the length of the shortest path from 0 to 3 is 1.
// <img src="https://assets.leetcode.com/uploads/2024/06/28/image12.jpg" />
// After the addition of the road from 0 to 2, the length of the shortest path remains 1.

// Constraints:
//     3 <= n <= 10^5
//     1 <= queries.length <= 10^5
//     queries[i].length == 2
//     0 <= queries[i][0] < queries[i][1] < n
//     1 < queries[i][1] - queries[i][0]
//     There are no repeated roads among the queries.
//     There are no two queries such that i != j and queries[i][0] < queries[j][0] < queries[i][1] < queries[j][1].

import "fmt"
import "sort"

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
    shortestPaths := make([]int, n)
    for idx := range shortestPaths {
        shortestPaths[idx] = idx
    }
    helper := func(shortestPaths *[]int, left, right int) {
        start := sort.Search(len(*shortestPaths), func(i int) bool { return (*shortestPaths)[i] > left })
        end   := sort.Search(len(*shortestPaths), func(i int) bool { return (*shortestPaths)[i] >= right })
        *shortestPaths = append((*shortestPaths)[:start], (*shortestPaths)[end:]...)
    }
    res := make([]int, 0)
    for _, query := range queries {
        left, right := query[0], query[1]
        helper(&shortestPaths, left, right)
        res = append(res, len(shortestPaths) - 1)
    }
    return res
}

func shortestDistanceAfterQueries1(n int, queries [][]int) []int {
    res, roads := []int{}, make([]int, n)
    for i := 0; i < n; i++ {
        roads[i] = i + 1;
    }
    dist := n
    for _, query := range queries {
        k := roads[query[0]]
        roads[query[0]] = query[1]
        for k != -1 && k < query[1] {
            k, roads[k] = roads[k], -1
            dist--
        }
        res = append(res, dist - 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 5, queries = [[2,4],[0,2],[0,4]]
    // Output: [3,2,1]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/28/image8.jpg" />
    // After the addition of the road from 2 to 4, the length of the shortest path from 0 to 4 is 3.
    // <img src="https://assets.leetcode.com/uploads/2024/06/28/image9.jpg" />
    // After the addition of the road from 0 to 2, the length of the shortest path from 0 to 4 is 2.
    // <img src="https://assets.leetcode.com/uploads/2024/06/28/image10.jpg" />
    // After the addition of the road from 0 to 4, the length of the shortest path from 0 to 4 is 1.
    fmt.Println(shortestDistanceAfterQueries(5,[][]int{{2,4},{0,2},{0,4}})) // [3,2,1]
    // Example 2:
    // Input: n = 4, queries = [[0,3],[0,2]]
    // Output: [1,1]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/28/image11.jpg" />
    // After the addition of the road from 0 to 3, the length of the shortest path from 0 to 3 is 1.
    // <img src="https://assets.leetcode.com/uploads/2024/06/28/image12.jpg" />
    // After the addition of the road from 0 to 2, the length of the shortest path remains 1.
    fmt.Println(shortestDistanceAfterQueries(4,[][]int{{0,3},{0,2}})) // [1,1]

    fmt.Println(shortestDistanceAfterQueries1(5,[][]int{{2,4},{0,2},{0,4}})) // [3,2,1]
    fmt.Println(shortestDistanceAfterQueries1(4,[][]int{{0,3},{0,2}})) // [1,1]
}