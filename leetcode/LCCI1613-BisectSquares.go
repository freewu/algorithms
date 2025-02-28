package main

// 面试题 16.13. Bisect Squares LCCI
// Given two squares on a two-dimensional plane, find a line that would cut these two squares in half. 
// Assume that the top and the bottom sides of the square run parallel to the x-axis.

// Each square consists of three values, the coordinate of bottom left corner [X,Y] = [square[0],square[1]], and the side length of the square square[2]. 
// The line will intersect to the two squares in four points. 
// Return the coordinates of two intersection points [X1,Y1] and [X2,Y2] that the forming segment covers the other two intersection points in format of {X1,Y1,X2,Y2}. 
// If X1 != X2, there should be X1 < X2, otherwise there should be Y1 <= Y2.

// If there are more than one line that can cut these two squares in half, 
// return the one that has biggest slope (slope of a line parallel to the y-axis is considered as infinity).

// Example:
// Input: 
// square1 = {-1, -1, 2}
// square2 = {0, -1, 2}
// Output: {-1,0,2,0}
// Explanation: y = 0 is the line that can cut these two squares in half.

// Note:
//     square.length == 3
//     square[2] > 0

import "fmt"

func cutSquares(square1 []int, square2 []int) []float64 {
    min := func (x, y int) float64 { if x < y { return float64(x); }; return float64(y); }
    max := func (x, y int) float64 { if x > y { return float64(x); }; return float64(y); }
    abs := func(x float64) float64 { if x < 0 { return -x; }; return x; }
    equal := func(x, y float64) bool { return abs(x - y) < 1e-6  }
    o1x, o1y := float64(square1[0]) + float64(square1[2]) / 2, float64(square1[1]) + float64(square1[2]) / 2
    o2x, o2y := float64(square2[0]) + float64(square2[2]) / 2, float64(square2[1]) + float64(square2[2]) / 2
    minX, maxX := min(square1[0], square2[0]), max(square1[0] + square1[2], square2[0] + square2[2])
    minY, maxY := min(square1[1], square2[1]), max(square1[1] + square1[2], square2[1] + square2[2])
    if equal(o1x, o2x) { return []float64{ o1x, minY, o1x, maxY } }
    if equal(o1y, o2y) { return []float64{ minX, o1y, maxX, o2y } }
    k := (o1y - o2y) / (o1x - o2x)
    // 与上下两边交
    // 由两点式： (x-x1)(y2-y1)=(y-y1)(x2-x1)知，y=y0是，x=(y0-y1)(x2-x1)/(y2-y1) + x1
    if k > 1 { // 左下右上走势
        return []float64{(minY - o1y) * (o2x - o1x) / (o2y - o1y) + o1x, minY, (maxY - o1y) * (o2x - o1x) / (o2y - o1y) + o1x, maxY}
    }
    if k < -1 { // 左上右下走势
        return []float64{(maxY - o1y) * (o2x - o1x) / (o2y - o1y) + o1x, maxY, (minY - o1y) * (o2x - o1x) / (o2y - o1y) + o1x, minY}
    }
    // 与左右两边交
    // 由两点式： (x-x1)(y2-y1)=(y-y1)(x2-x1)知，x=x0时， y=(x0-x1)(y2-y1)/(x2-x1) + y1
    return []float64{minX, (minX - o1x) * (o2y - o1y) / (o2x - o1x) + o1y, maxX, (maxX - o1x) * (o2y - o1y) / (o2x - o1x) + o1y }
}

func main() {
    // Example:
    // Input: 
    // square1 = {-1, -1, 2}
    // square2 = {0, -1, 2}
    // Output: {-1,0,2,0}
    // Explanation: y = 0 is the line that can cut these two squares in half.
    fmt.Println(cutSquares([]int{-1, -1, 2}, []int{0, -1, 2})) // [-1 0 2 0]

    fmt.Println(cutSquares([]int{1, 2, 3}, []int{1, 2, 3})) // [2.5 2 2.5 5]
    fmt.Println(cutSquares([]int{1, 2, 3}, []int{-1, -2, -3})) // [-1.4285714285714284 -2 3.571428571428571 5]
    fmt.Println(cutSquares([]int{1, 2, 3}, []int{3, 2, 1})) // [1 5 4 2]
    fmt.Println(cutSquares([]int{1, 2, 3}, []int{7, 8, 9})) // [1 2 16 17]
}