package main

// 3710. Maximum Partition Factor
// You are given a 2D integer array points, where points[i] = [xi, yi] represents the coordinates of the ith point on the Cartesian plane.

// The Manhattan distance between two points points[i] = [xi, yi] and points[j] = [xj, yj] is |xi - xj| + |yi - yj|.

// Split the n points into exactly two non-empty groups. 
// The partition factor of a split is the minimum Manhattan distance among all unordered pairs of points that lie in the same group.

// Return the maximum possible partition factor over all valid splits.

// Note: A group of size 1 contributes no intra-group pairs. 
// When n = 2 (both groups size 1), there are no intra-group pairs, so define the partition factor as 0.

// Example 1:
// Input: points = [[0,0],[0,2],[2,0],[2,2]]
// Output: 4
// Explanation:
// We split the points into two groups: {[0, 0], [2, 2]} and {[0, 2], [2, 0]}.
// In the first group, the only pair has Manhattan distance |0 - 2| + |0 - 2| = 4.
// In the second group, the only pair also has Manhattan distance |0 - 2| + |2 - 0| = 4.
// The partition factor of this split is min(4, 4) = 4, which is maximal.

// Example 2:
// Input: points = [[0,0],[0,1],[10,0]]
// Output: 11
// Explanation:​​​​​​​
// We split the points into two groups: {[0, 1], [10, 0]} and {[0, 0]}.
// In the first group, the only pair has Manhattan distance |0 - 10| + |1 - 0| = 11.
// The second group is a singleton, so it contributes no pairs.
// The partition factor of this split is 11, which is maximal.

// Constraints:
//     2 <= points.length <= 500
//     points[i] = [xi, yi]
//     -10^8 <= xi, yi <= 10^8

import "fmt"
import "sort"

func maxPartitionFactor(points [][]int) int {
    n, mx := len(points), int(4e8)
    if n == 2 { return 0 }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    // 判断二分图
    isBipartite := func(points [][]int, low int) bool {
        colors := make([]int8, len(points))
        var dfs func(x int, c int8) bool 
        dfs = func(x int, c int8) bool {
            colors[x] = c
            p := points[x]
            for y, q := range points {
                if y == x || abs(p[0]-q[0])+abs(p[1]-q[1]) >= low { // 符合要求
                    continue
                }
                if colors[y] == c || colors[y] == 0 && !dfs(y, -c) {
                    return false // 不是二分图
                }
            }
            return true
        }
        // 可能有多个连通块
        for i, c := range colors {
            if c == 0 && !dfs(i, 1) {
                return false
            }
        }
        return true
    }
    return sort.Search(mx, func(low int) bool {
        // 二分最小的不满足要求的 low+1，就可以得到最大的满足要求的 low
        return !isBipartite(points, low + 1)
    })
}

func main() {
    // Example 1:
    // Input: points = [[0,0],[0,2],[2,0],[2,2]]
    // Output: 4
    // Explanation:
    // We split the points into two groups: {[0, 0], [2, 2]} and {[0, 2], [2, 0]}.
    // In the first group, the only pair has Manhattan distance |0 - 2| + |0 - 2| = 4.
    // In the second group, the only pair also has Manhattan distance |0 - 2| + |2 - 0| = 4.
    // The partition factor of this split is min(4, 4) = 4, which is maximal.
    fmt.Println(maxPartitionFactor([][]int{{0,0},{0,2},{2,0},{2,2}})) // 4
    // Example 2:
    // Input: points = [[0,0],[0,1],[10,0]]
    // Output: 11
    // Explanation:​​​​​​​
    // We split the points into two groups: {[0, 1], [10, 0]} and {[0, 0]}.
    // In the first group, the only pair has Manhattan distance |0 - 10| + |1 - 0| = 11.
    // The second group is a singleton, so it contributes no pairs.
    // The partition factor of this split is 11, which is maximal.
    fmt.Println(maxPartitionFactor([][]int{{0,0},{0,1},{10,0}})) // 11
}