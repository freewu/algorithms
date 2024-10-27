package main

// 1637. Widest Vertical Area Between Two Points Containing No Points
// Given n points on a 2D plane where points[i] = [xi, yi], 
// Return the widest vertical area between two points such that no points are inside the area.

// A vertical area is an area of fixed-width extending infinitely along the y-axis (i.e., infinite height). 
// The widest vertical area is the one with the maximum width.

// Note that points on the edge of a vertical area are not considered included in the area.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/19/points3.png" />
// Input: points = [[8,7],[9,9],[7,4],[9,7]]
// Output: 1
// Explanation: Both the red and the blue area are optimal.

// Example 2:
// Input: points = [[3,1],[9,0],[1,0],[1,4],[5,3],[8,8]]
// Output: 3

// Constraints:
//     n == points.length
//     2 <= n <= 10^5
//     points[i].length == 2
//     0 <= xi, yi <= 10^9

import "fmt"
import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
    sort.Slice(points, func(i, j int) bool { 
        return points[i][0] < points[j][0]
    })
    res := 0
    for i := 1; i < len(points); i++ {
        width := points[i][0] - points[i - 1][0]
        if width > res {
            res = width
        }
    }
    return res
}

func maxWidthOfVerticalArea1(points [][]int) int {
    res, arr := 0, make([]int, 0, len(points))
    for _, p := range points {
        arr = append(arr, p[0])
    }
    sort.Ints(arr)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(arr); i++ {
        res = max(res, arr[i] - arr[i-1])
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/19/points3.png" />
    // Input: points = [[8,7],[9,9],[7,4],[9,7]]
    // Output: 1
    // Explanation: Both the red and the blue area are optimal.
    fmt.Println(maxWidthOfVerticalArea([][]int{{8,7},{9,9},{7,4},{9,7}})) // 1
    // Example 2:
    // Input: points = [[3,1],[9,0],[1,0],[1,4],[5,3],[8,8]]
    // Output: 3
    fmt.Println(maxWidthOfVerticalArea([][]int{{3,1},{9,0},{1,0},{1,4},{5,3},{8,8}})) // 3

    fmt.Println(maxWidthOfVerticalArea1([][]int{{8,7},{9,9},{7,4},{9,7}})) // 1
    fmt.Println(maxWidthOfVerticalArea1([][]int{{3,1},{9,0},{1,0},{1,4},{5,3},{8,8}})) // 3
}