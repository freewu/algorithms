package main

// 3623. Count Number of Trapezoids I
// You are given a 2D integer array points, where points[i] = [xi, yi] represents the coordinates of the ith point on the Cartesian plane.

// A horizontal trapezoid is a convex quadrilateral with at least one pair of horizontal sides (i.e. parallel to the x-axis). 
// Two lines are parallel if and only if they have the same slope.

// Return the number of unique horizontal trapezoids that can be formed by choosing any four distinct points from points.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: points = [[1,0],[2,0],[3,0],[2,2],[3,2]]
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/05/01/desmos-graph-7.png" />
// There are three distinct ways to pick four points that form a horizontal trapezoid:
// Using points [1,0], [2,0], [3,2], and [2,2].
// Using points [2,0], [3,0], [3,2], and [2,2].
// Using points [1,0], [3,0], [3,2], and [2,2].

// Example 2:
// Input: points = [[0,0],[1,0],[0,1],[2,1]]
// Output: 1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/29/desmos-graph-5.png" />
// There is only one horizontal trapezoid that can be formed.

// Constraints:
//     4 <= points.length <= 10^5
//     –10^8 <= xi, yi <= 10^8
//     All points are pairwise distinct.

import "fmt"

func countTrapezoids(points [][]int) int {
    res, sum, mod := 0, 0, 1_000_000_007
    count := make(map[int]int)
    for _, v := range points {
        count[v[1]]++
    }
    for _, v := range count {
        k := v * (v - 1) / 2 % mod
        res = (res + sum * k) % mod
        sum = (sum + k) % mod
    }
    return res
}

func countTrapezoids1(points [][]int) int {
    res, sum, mod := 0, 0, 1_000_000_007
    count, arr := make(map[int]int), make([]int,0)
    for _, v := range points {
        count[v[1]]++
    }
    for _, v := range count {
        if v < 2 { continue }
        arr = append(arr, v * (v - 1) / 2)
    }
    for i := len(arr) - 1; i >= 0; i-- {
        res = (res + (arr[i]) * sum) % mod
        sum = (sum + (arr[i])) % mod
    }
    return res
}

func countTrapezoids2(points [][]int) int {
    mp := make(map[int]int) // y -> count  评选与x轴即取两个y值相同的点
    res, sum := 0,0  // 统计结果, 统计一个有多少个水平边
    for _, point := range points {
        count := mp[point[1]]
        // sum - (count * (count - 1) / 2 表示除去当前y值有多少个水平与x轴的线
        // *count 是因为每多一个点就多count个相同y值的水平线
        res = (res + (sum - (count * (count - 1) / 2)) * count)
        sum += count
        mp[point[1]]++
    }
    return res %  1_000_000_007
}

func main() {
    // Example 1:
    // Input: points = [[1,0],[2,0],[3,0],[2,2],[3,2]]
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/05/01/desmos-graph-7.png" />
    // There are three distinct ways to pick four points that form a horizontal trapezoid:
    // Using points [1,0], [2,0], [3,2], and [2,2].
    // Using points [2,0], [3,0], [3,2], and [2,2].
    // Using points [1,0], [3,0], [3,2], and [2,2].
    fmt.Println(countTrapezoids([][]int{{1,0},{2,0},{3,0},{2,2},{3,2}})) // 3
    // Example 2:
    // Input: points = [[0,0],[1,0],[0,1],[2,1]]
    // Output: 1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/04/29/desmos-graph-5.png" />
    // There is only one horizontal trapezoid that can be formed.
    fmt.Println(countTrapezoids([][]int{{0,0},{1,0},{0,1},{2,1}})) // 1

    fmt.Println(countTrapezoids1([][]int{{1,0},{2,0},{3,0},{2,2},{3,2}})) // 3
    fmt.Println(countTrapezoids1([][]int{{0,0},{1,0},{0,1},{2,1}})) // 1

    fmt.Println(countTrapezoids2([][]int{{1,0},{2,0},{3,0},{2,2},{3,2}})) // 3
    fmt.Println(countTrapezoids2([][]int{{0,0},{1,0},{0,1},{2,1}})) // 1
}