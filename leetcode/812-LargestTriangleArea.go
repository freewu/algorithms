package main

// 812. Largest Triangle Area
// Given an array of points on the X-Y plane points where points[i] = [xi, yi], 
// return the area of the largest triangle that can be formed by any three different points. 
// Answers within 10^-5 of the actual answer will be accepted.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/04/1027.png" />
// Input: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
// Output: 2.00000
// Explanation: The five points are shown in the above figure. The red triangle is the largest.

// Example 2:
// Input: points = [[1,0],[0,0],[0,1]]
// Output: 0.50000

// Constraints:
//     3 <= points.length <= 50
//     -50 <= xi, yi <= 50
//     All the given points are unique.

import "fmt"

func largestTriangleArea(points [][]int) float64 {
    res, n := float64(0), len(points)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    calc := func(x1, y1, x2, y2, x3, y3 int) float64 { return float64(abs(x1*y2+x2*y3+x3*y1-y1*x2-y2*x3-y3*x1)) / 2.0 }
    for i := 0; i < n - 2; i++ {
        for j := i + 1; j < n - 1; j++ {
            for k := 0; k < n; k++ {
                if v := calc(points[i][0], points[i][1], points[j][0], points[j][1], points[k][0], points[k][1]); v > res {
                    res = v
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/04/1027.png" />
    // Input: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
    // Output: 2.00000
    // Explanation: The five points are shown in the above figure. The red triangle is the largest.
    fmt.Println(largestTriangleArea([][]int{{0,0},{0,1},{1,0},{0,2},{2,0}})) // 2.00000
    // Example 2:
    // Input: points = [[1,0],[0,0],[0,1]]
    // Output: 0.50000
    fmt.Println(largestTriangleArea([][]int{{1,0},{0,0},{0,1}})) // 0.50000
}