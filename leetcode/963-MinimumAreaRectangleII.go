package main

// 963. Minimum Area Rectangle II
// You are given an array of points in the X-Y plane points where points[i] = [xi, yi].

// Return the minimum area of any rectangle formed from these points, with sides not necessarily parallel to the X and Y axes. 
// If there is not any such rectangle, return 0.

// Answers within 10^-5 of the actual answer will be accepted.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/12/21/1a.png" />
// Input: points = [[1,2],[2,1],[1,0],[0,1]]
// Output: 2.00000
// Explanation: The minimum area rectangle occurs at [1,2],[2,1],[1,0],[0,1], with an area of 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2018/12/22/2.png" />
// Input: points = [[0,1],[2,1],[1,1],[1,0],[2,0]]
// Output: 1.00000
// Explanation: The minimum area rectangle occurs at [1,0],[1,1],[2,1],[2,0], with an area of 1.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2018/12/22/3.png" />
// Input: points = [[0,3],[1,2],[3,1],[1,3],[2,1]]
// Output: 0
// Explanation: There is no possible rectangle to form from these points.

// Constraints:
//     1 <= points.length <= 50
//     points[i].length == 2
//     0 <= xi, yi <= 4 * 10^4
//     All the given points are unique.

import "fmt"
import "math"

func minAreaFreeRect(points [][]int) float64 {
    res, centers := 0, make(map[[3]int][][2]int) // group pointer pairs by key [3]int{center_x*2, center_y*2, dis2}
    distance := func(p0, p1 []int) int {
        dx, dy := p0[0] - p1[0], p0[1] - p1[1]
        return dx*dx + dy*dy
    }
    for i, p1 := range points {
        for j := i + 1; j < len(points); j++ {
            p2 := points[j]
            cp := [3]int{p1[0] + p2[0], p1[1] + p2[1], distance(p1, p2) }
            v := centers[cp]
            v = append(v, [2]int{i, j})
            centers[cp] = v
        }
    }
    for _, l := range centers {
        for i := 0; i < len(l); i++ {
            for j := i+1; j < len(l); j++ {
                p0, p2, p3 := points[l[i][0]], points[l[j][0]], points[l[j][1]]
                area2 := distance(p0, p2) * distance(p0, p3)
                if res == 0 || res > area2 {
                    res = area2
                }
            }
        }
    }
    return math.Sqrt(float64(res))
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/12/21/1a.png" />
    // Input: points = [[1,2],[2,1],[1,0],[0,1]]
    // Output: 2.00000
    // Explanation: The minimum area rectangle occurs at [1,2],[2,1],[1,0],[0,1], with an area of 2.
    fmt.Println(minAreaFreeRect([][]int{{1,2},{2,1},{1,0},{0,1}})) // 2.00000
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2018/12/22/2.png" />
    // Input: points = [[0,1],[2,1],[1,1],[1,0],[2,0]]
    // Output: 1.00000
    // Explanation: The minimum area rectangle occurs at [1,0],[1,1],[2,1],[2,0], with an area of 1.
    fmt.Println(minAreaFreeRect([][]int{{0,1},{2,1},{1,1},{1,0},{2,0}})) // 1.00000
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2018/12/22/3.png" />
    // Input: points = [[0,3],[1,2],[3,1],[1,3],[2,1]]
    // Output: 0
    // Explanation: There is no possible rectangle to form from these points.
    fmt.Println(minAreaFreeRect([][]int{{0,3},{1,2},{3,1},{1,3},{2,1}})) // 0
}