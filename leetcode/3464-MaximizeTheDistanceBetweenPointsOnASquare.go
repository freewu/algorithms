package main

// 3464. Maximize the Distance Between Points on a Square
// You are given an integer side, representing the edge length of a square with corners at (0, 0), (0, side), (side, 0), and (side, side) on a Cartesian plane.

// You are also given a positive integer k and a 2D integer array points, where points[i] = [xi, yi] represents the coordinate of a point lying on the boundary of the square.

// You need to select k elements among points such that the minimum Manhattan distance between any two points is maximized.

// Return the maximum possible minimum Manhattan distance between the selected k points.

// The Manhattan Distance between two cells (xi, yi) and (xj, yj) is |xi - xj| + |yi - yj|.

// Example 1:
// Input: side = 2, points = [[0,2],[2,0],[2,2],[0,0]], k = 4
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/01/28/4080_example0_revised.png" />
// Select all four points.

// Example 2:
// Input: side = 2, points = [[0,0],[1,2],[2,0],[2,2],[2,1]], k = 4
// Output: 1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/01/28/4080_example1_revised.png" />
// Select the points (0, 0), (2, 0), (2, 2), and (2, 1).

// Example 3:
// Input: side = 2, points = [[0,0],[0,1],[0,2],[1,2],[2,0],[2,2],[2,1]], k = 5
// Output: 1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/01/28/4080_example2_revised.png" />
// Select the points (0, 0), (0, 1), (0, 2), (1, 2), and (2, 2).

// Constraints:
//     1 <= side <= 10^9
//     4 <= points.length <= min(4 * side, 15 * 10^3)
//     points[i] == [xi, yi]
//     The input is generated such that:
//     points[i] lies on the boundary of the square.
//     All points[i] are unique.
//     4 <= k <= min(25, points.length)

import "fmt"
import "sort"

func maxDistance(side int, points [][]int, k int) int {
    n := len(points)
    arr := make([]int, n)
    for i, p := range points {
        x, y := p[0], p[1]
        if x == 0 {
            arr[i] = y
        } else if y == side {
            arr[i] = side + x
        } else if x == side {
            arr[i] = side * 3 - y
        } else {
            arr[i] = side * 4 - x
        }
    }
    sort.Ints(arr)
    f, end := make([]int, n + 1), make([]int, n)
    return sort.Search(side, func(low int) bool {
        low++
        j := n
        for i := n - 1; i >= 0; i-- {
            for arr[j - 1] >= arr[i] + low {
                j--
            }
            f[i] = f[j] + 1
            if f[i] == 1 {
                end[i] = i // i 自己就是最后一个点
            } else {
                end[i] = end[j]
            }
            if f[i] > k || f[i] == k && arr[end[i]] - arr[i] <= side * 4-low {
                return false
            }
        }
        return true
    })
}

func main() {
    // Example 1:
    // Input: side = 2, points = [[0,2],[2,0],[2,2],[0,0]], k = 4
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/01/28/4080_example0_revised.png" />
    // Select all four points.
    fmt.Println(maxDistance(2, [][]int{{0,2},{2,0},{2,2},{0,0}}, 4)) // 4
    // Example 2:
    // Input: side = 2, points = [[0,0],[1,2],[2,0],[2,2],[2,1]], k = 4
    // Output: 1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/01/28/4080_example1_revised.png" />
    // Select the points (0, 0), (2, 0), (2, 2), and (2, 1).
    fmt.Println(maxDistance(2, [][]int{{0,0},{1,2},{2,0},{2,2},{2,1}}, 4)) // 1
    // Example 3:
    // Input: side = 2, points = [[0,0],[0,1],[0,2],[1,2],[2,0],[2,2],[2,1]], k = 5
    // Output: 1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/01/28/4080_example2_revised.png" />
    // Select the points (0, 0), (0, 1), (0, 2), (1, 2), and (2, 2).
    fmt.Println(maxDistance(2, [][]int{{0,0},{0,1},{0,2},{1,2},{2,0},{2,2},{2,1}}, 5)) // 1
}