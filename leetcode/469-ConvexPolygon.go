package main

// 469. Convex Polygon
// You are given an array of points on the X-Y plane points where points[i] = [xi, yi]. 
// The points form a polygon when joined sequentially.

// Return true if this polygon is convex and false otherwise.

// You may assume the polygon formed by given points is always a simple polygon. 
// In other words, we ensure that exactly two edges intersect at each vertex and that edges otherwise don't intersect each other.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/09/covpoly1-plane.jpg" />
// Input: points = [[0,0],[0,5],[5,5],[5,0]]
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/09/covpoly2-plane.jpg" />
// Input: points = [[0,0],[0,10],[10,10],[10,0],[5,5]]
// Output: false

// Constraints:
//     3 <= points.length <= 10^4
//     points[i].length == 2
//     -10^4 <= xi, yi <= 10^4
//     All the given points are unique.

import "fmt"

func isConvex(points [][]int) bool {
    n, preCrossProduct := len(points), 0 // 用于记录前一个叉积的值
    for i := 0; i < n; i++ {
        // 向量(x1,y1)
        dx1 := points[(i+1) % n][0] - points[i][0]
        dx2 := points[(i+2) % n][0] - points[i][0]
        // 向量(x2, y2)
        dy1 := points[(i+1) % n][1] - points[i][1]
        dy2 := points[(i+2) % n][1] - points[i][1]
        // 作叉积
        crossProduct := dx1 * dy2 - dx2 * dy1
        if crossProduct != 0 {
            // 前后两个叉积异号，非凸,注意强制类型转换
            if crossProduct * preCrossProduct < 0 {
                return false
            }
            preCrossProduct = crossProduct
        }
    }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/09/covpoly1-plane.jpg" />
    // Input: points = [[0,0],[0,5],[5,5],[5,0]]
    // Output: true
    fmt.Println(isConvex([][]int{{0,0},{0,5},{5,5},{5,0}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/09/covpoly2-plane.jpg" />
    // Input: points = [[0,0],[0,10],[10,10],[10,0],[5,5]]
    // Output: false
    fmt.Println(isConvex([][]int{{0,0},{0,10},{10,10},{10,0},{5,5}})) // false
}