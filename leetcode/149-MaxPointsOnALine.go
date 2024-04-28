package main

// 149. Max Points on a Line
// Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane, 
// return the maximum number of points that lie on the same straight line.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/25/plane1.jpg" />
// Input: points = [[1,1],[2,2],[3,3]]
// Output: 3

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/02/25/plane2.jpg" />
// Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
// Output: 4
 
// Constraints:
//     1 <= points.length <= 300
//     points[i].length == 2
//     -10^4 <= xi, yi <= 10^4
//     All the points are unique.

import "fmt"
import "math"

func maxPoints(points [][]int) int {
    res := 0
    for i := 0; i < len(points); i++ {
        slopes := make(map[float64]int)
        for j := i + 1; j < len(points); j++ {
            x1, y1, x2, y2 := float64(points[i][0]),float64(points[i][1]), float64(points[j][0]),float64(points[j][1])
            slope := float64(0)
            if x1 == x2 {
                slope = math.Inf(1)
            } else {
                slope = (y2 - y1) / (x2 - x1)
            }
            slopes[slope]++
            if res < slopes[slope] {
                res = slopes[slope]
            }
        }
    }
    return res + 1
}

func maxPoints1(points [][]int) int {
    if len(points) <= 2 {
        return len(points)
    }
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(points); i++ {
        for j := i + 1; j < len(points); j++ {
            y1 := points[j][1] - points[i][1]
            x1 := points[j][0] - points[i][0]
            count := 2
            for k := j + 1; k < len(points); k++ {
                y2 := points[k][1] - points[j][1]
                x2 := points[k][0] - points[j][0]
                if y1 * x2 == y2 * x1 {
                    count++
                }
            }
            res = max(res, count)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: points = [[1,1],[2,2],[3,3]]
    // Output: 3
    fmt.Println(maxPoints([][]int{{1,1},{2,2},{3,3}})) // 3
    // Example 2:
    // Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
    // Output: 4
    fmt.Println(maxPoints([][]int{{1,1},{3,2},{5,3},{4,1},{2,3},{1,4}})) // 4

    fmt.Println(maxPoints1([][]int{{1,1},{2,2},{3,3}})) // 3
    fmt.Println(maxPoints1([][]int{{1,1},{3,2},{5,3},{4,1},{2,3},{1,4}})) // 4
}