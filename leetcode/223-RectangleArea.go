package main

// 223. Rectangle Area
// Given the coordinates of two rectilinear rectangles in a 2D plane, 
// return the total area covered by the two rectangles.
// The first rectangle is defined by its bottom-left corner (ax1, ay1) and its top-right corner (ax2, ay2).
// The second rectangle is defined by its bottom-left corner (bx1, by1) and its top-right corner (bx2, by2).

// Example 1:
// Rectangle Area
// <img src="https://assets.leetcode.com/uploads/2021/05/08/rectangle-plane.png" />
// Input: ax1 = -3, ay1 = 0, ax2 = 3, ay2 = 4, bx1 = 0, by1 = -1, bx2 = 9, by2 = 2
// Output: 45

// Example 2:
// Input: ax1 = -2, ay1 = -2, ax2 = 2, ay2 = 2, bx1 = -2, by1 = -2, bx2 = 2, by2 = 2
// Output: 16

// Constraints:
//     -10^4 <= ax1 <= ax2 <= 10^4
//     -10^4 <= ay1 <= ay2 <= 10^4
//     -10^4 <= bx1 <= bx2 <= 10^4
//     -10^4 <= by1 <= by2 <= 10^4

import "fmt"

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
    area1 := (ax2 - ax1) * (ay2 - ay1) // 计算出面积1
    area2 := (bx2 - bx1) * (by2 - by1) // 计算出面积2
    if ax1 >= bx2 || bx1 >= ax2 || ay1 >= by2 || by1 >= ay2 { // 判断两个矩形是否重合
        return area1 + area2
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    x1, y1 := max(ax1, bx1), max(ay1, by1)
    x2, y2 := min(ax2, bx2), min(ay2, by2)
    return area1 + area2 - ((x2 - x1) * (y2 - y1)) // 减去重合部分
}

func computeArea1(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
    area1 := (ax2 - ax1) * (ay2 - ay1)
    area2 := (bx2 - bx1) * (by2 - by1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    overlapWidth := min(ax2, bx2) - max(ax1, bx1)
    overlapHeight := min(ay2, by2) - max(ay1, by1)
    overlayArea := max(overlapWidth, 0) * max(overlapHeight, 0)
    return area1 + area2 - overlayArea
}

func main() {
    // Input: ax1 = -3, ay1 = 0, ax2 = 3, ay2 = 4, bx1 = 0, by1 = -1, bx2 = 9, by2 = 2
    // Output: 45
    fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2)) // 45
    // Example 2:
    // Input: ax1 = -2, ay1 = -2, ax2 = 2, ay2 = 2, bx1 = -2, by1 = -2, bx2 = 2, by2 = 2
    // Output: 16
    fmt.Println(computeArea(-2, -2, 2, 2, -2, -2, 2, 2)) // 16

    fmt.Println(computeArea1(-3, 0, 3, 4, 0, -1, 9, 2)) // 45
    fmt.Println(computeArea1(-2, -2, 2, 2, -2, -2, 2, 2)) // 16
}