package main

// 1037. Valid Boomerang
// Given an array points where points[i] = [xi, yi] represents a point on the X-Y plane, 
// return true if these points are a boomerang.

// A boomerang is a set of three points that are all distinct and not in a straight line.

// Example 1:
// Input: points = [[1,1],[2,3],[3,2]]
// Output: true

// Example 2:
// Input: points = [[1,1],[2,2],[3,3]]
// Output: false

// Constraints:
//     points.length == 3
//     points[i].length == 2
//     0 <= xi, yi <= 100

import "fmt"

// func isBoomerang(points [][]int) bool {
//     diff := points[0][0] - points[0][1]
//     for i := 1; i < len(points); i++ {
//         if (points[i][0] - points[i][1]) != diff {
//             return true
//         }
//     }
//     return false // 一条直线
// }

func isBoomerang(points [][]int) bool {
    return !((points[0][0] - points[1][0]) * (points[0][1] - points[2][1]) == 
             (points[0][0] - points[2][0]) * (points[0][1] - points[1][1]))
}

func main() {
    // Example 1:
    // Input: points = [[1,1],[2,3],[3,2]]
    // Output: true
    fmt.Println(isBoomerang([][]int{{1,1},{2,3},{3,2}})) // true
    // Example 2:
    // Input: points = [[1,1],[2,2],[3,3]]
    // Output: false
    fmt.Println(isBoomerang([][]int{{1,1},{2,2},{3,3}})) // false

    fmt.Println(isBoomerang([][]int{{0,0},{2,1},{2,1}})) // false
}