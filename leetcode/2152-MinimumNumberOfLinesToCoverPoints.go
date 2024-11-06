package main

// 2152. Minimum Number of Lines to Cover Points
// You are given an array points where points[i] = [xi, yi] represents a point on an X-Y plane.

// Straight lines are going to be added to the X-Y plane, such that every point is covered by at least one line.

// Return the minimum number of straight lines needed to cover all the points.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/01/23/image-20220123200023-1.png" />
// Input: points = [[0,1],[2,3],[4,5],[4,3]]
// Output: 2
// Explanation: The minimum number of straight lines needed is two. One possible solution is to add:
// - One line connecting the point at (0, 1) to the point at (4, 5).
// - Another line connecting the point at (2, 3) to the point at (4, 3).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/01/23/image-20220123200057-3.png" />
// Input: points = [[0,2],[-2,-2],[1,4]]
// Output: 1
// Explanation: The minimum number of straight lines needed is one. The only solution is to add:
// - One line connecting the point at (-2, -2) to the point at (1, 4).

// Constraints:
//     1 <= points.length <= 10
//     points[i].length == 2
//     -100 <= xi, yi <= 100
//     All the points are unique.

import "fmt"
import "math"

func minimumLines(points [][]int) int {
    if len(points) <= 2 { return 1 }
    dp := make(map[int]int)
    setBit := func(n, bit int) int { return n | (1 << bit) }
    getBit := func(n, bit int) int { return (n >> bit) & 1 }
    check := func(x1, y1, x2, y2, x3, y3 int) bool { return (y3 - y1) * (x2 - x1) == (y2 - y1) * (x3 - x1) }
    var dfs func(condition, nowCount int) int
    dfs = func(condition, nowCount int ) int {
        if condition == int(math.Pow(2, float64(len(points)))) - 1 { return nowCount  }
        if v, ok := dp[condition]; ok { return nowCount + v }
        res, set := 0, false
        for i := 0; i < len(points); i++ {
            if getBit(condition, i) == 0 {
                x1, y1 := points[i][0], points[i][1]
                for j := 0; j < len(points); j++ {
                    if i == j { continue }
                    x2, y2 := points[j][0], points[j][1]
                    newCondition := condition
                    newCondition = setBit(newCondition, i)
                    newCondition = setBit(newCondition, j)
                    for k := 0; k < len(points); k++ {
                        if i == k || j == k { continue }
                        x3, y3 := points[k][0], points[k][1]
                        if check(x1, y1, x2, y2, x3, y3) {
                            newCondition = setBit(newCondition, k)
                        }
                    }
                    r := dfs(newCondition, nowCount + 1)
                    if !set || r < res {
                        res, set = r, true
                    }
                }
            }
        }
        dp[condition] = res - nowCount
        return res
    }
    return dfs(0, 0)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/01/23/image-20220123200023-1.png" />
    // Input: points = [[0,1],[2,3],[4,5],[4,3]]
    // Output: 2
    // Explanation: The minimum number of straight lines needed is two. One possible solution is to add:
    // - One line connecting the point at (0, 1) to the point at (4, 5).
    // - Another line connecting the point at (2, 3) to the point at (4, 3).
    fmt.Println(minimumLines([][]int{{0,1},{2,3},{4,5},{4,3}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/01/23/image-20220123200057-3.png" />
    // Input: points = [[0,2],[-2,-2],[1,4]]
    // Output: 1
    // Explanation: The minimum number of straight lines needed is one. The only solution is to add:
    // - One line connecting the point at (-2, -2) to the point at (1, 4).
    fmt.Println(minimumLines([][]int{{0,2},{-2,-2},{1,4}})) // 1
}