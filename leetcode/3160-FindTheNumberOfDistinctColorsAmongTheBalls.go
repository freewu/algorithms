package main

// 3160. Find the Number of Distinct Colors Among the Balls
// You are given an integer limit and a 2D array queries of size n x 2.

// There are limit + 1 balls with distinct labels in the range [0, limit]. Initially, all balls are uncolored. 
// For every query in queries that is of the form [x, y], you mark ball x with the color y. 
// After each query, you need to find the number of distinct colors among the balls.

// Return an array result of length n, where result[i] denotes the number of distinct colors after ith query.

// Note that when answering a query, lack of a color will not be considered as a color.

// Example 1:
// Input: limit = 4, queries = [[1,4],[2,5],[1,3],[3,4]]
// Output: [1,2,2,3]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/04/17/ezgifcom-crop.gif" />
// After query 0, ball 1 has color 4.
// After query 1, ball 1 has color 4, and ball 2 has color 5.
// After query 2, ball 1 has color 3, and ball 2 has color 5.
// After query 3, ball 1 has color 3, ball 2 has color 5, and ball 3 has color 4.

// Example 2:
// Input: limit = 4, queries = [[0,1],[1,2],[2,2],[3,4],[4,5]]
// Output: [1,2,2,3,4]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/04/17/ezgifcom-crop2.gif" />
// After query 0, ball 0 has color 1.
// After query 1, ball 0 has color 1, and ball 1 has color 2.
// After query 2, ball 0 has color 1, and balls 1 and 2 have color 2.
// After query 3, ball 0 has color 1, balls 1 and 2 have color 2, and ball 3 has color 4.
// After query 4, ball 0 has color 1, balls 1 and 2 have color 2, ball 3 has color 4, and ball 4 has color 5.
 
// Constraints:
//     1 <= limit <= 10^9
//     1 <= n == queries.length <= 10^5
//     queries[i].length == 2
//     0 <= queries[i][0] <= limit
//     1 <= queries[i][1] <= 10^9

import "fmt"

func queryResults(limit int, queries [][]int) []int {
    n := len(queries)
    res, colors, balls := make([]int, 0, n), make(map[int]int), make(map[int]int)
    for _, query := range queries {
        ball, color := query[0], query[1] // Extract ball label and color from query
        if prev, ok := balls[ball]; ok { // Check if ball is already colored
            colors[prev]-- // Decrement count of the previous color on the ball
            if colors[prev] == 0 { // If there are no balls with previous color left, remove color from color map
                delete(colors, prev)
            }
        }
        balls[ball] = color // Set color of ball to the new color
        colors[color]++ // Increment the count of the new color
        res = append(res, len(colors))
    }
    return res
}

func main() {
    // Example 1:
    // Input: limit = 4, queries = [[1,4],[2,5],[1,3],[3,4]]
    // Output: [1,2,2,3]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/04/17/ezgifcom-crop.gif" />
    // After query 0, ball 1 has color 4.
    // After query 1, ball 1 has color 4, and ball 2 has color 5.
    // After query 2, ball 1 has color 3, and ball 2 has color 5.
    // After query 3, ball 1 has color 3, ball 2 has color 5, and ball 3 has color 4.
    fmt.Println(queryResults(4,[][]int{{1,4},{2,5},{1,3},{3,4}})) // [1,2,2,3]
    // Example 2:
    // Input: limit = 4, queries = [[0,1],[1,2],[2,2],[3,4],[4,5]]
    // Output: [1,2,2,3,4]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/04/17/ezgifcom-crop2.gif" />
    // After query 0, ball 0 has color 1.
    // After query 1, ball 0 has color 1, and ball 1 has color 2.
    // After query 2, ball 0 has color 1, and balls 1 and 2 have color 2.
    // After query 3, ball 0 has color 1, balls 1 and 2 have color 2, and ball 3 has color 4.
    // After query 4, ball 0 has color 1, balls 1 and 2 have color 2, ball 3 has color 4, and ball 4 has color 5.
    fmt.Println(queryResults(4,[][]int{{0,1},{1,2},{2,2},{3,4},{4,5}})) // [1,2,2,3,4]

    // fmt.Println(queryResults1(4,[][]int{{1,4},{2,5},{1,3},{3,4}})) // [1,2,2,3]
    // fmt.Println(queryResults1(4,[][]int{{0,1},{1,2},{2,2},{3,4},{4,5}})) // [1,2,2,3,4]
}