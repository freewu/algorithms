package main

// 面试题 16.03. Intersection LCCI
// Given two straight line segments (represented as a start point and an end point), compute the point of intersection, if any. 
// If there's no intersection, return an empty array.

// The absolute error should not exceed 10^-6. 
// If there are more than one intersections, return the one with smallest X axis value. 
// If there are more than one intersections that have same X axis value, return the one with smallest Y axis value.

// Example 1:
// Input: 
// line1 = {0, 0}, {1, 0}
// line2 = {1, 1}, {0, -1}
// Output:  {0.5, 0}

// Example 2:
// Input: 
// line1 = {0, 0}, {3, 3}
// line2 = {1, 1}, {2, 2}
// Output:  {1, 1}

// Example 3:
// Input: 
// line1 = {0, 0}, {1, 1}
// line2 = {1, 0}, {2, 1}
// Output:  {}  (no intersection)

// Note:
//     The absolute value of coordinate value will not exceed 2^7.
//     All coordinates are valid 2D coordinates.

import "fmt"

// 使用直线的一般式Ax+By+C=0来求解。 已经两点 (x1,y1),(x2,y2) 可以得到A=y2-y1,B=x1-x2,c=x2y1-x1y2
// 已知两条直线A1x+B1y+C1=0 和 A2x+B2y+C2=0, 记D=A1B2-A2B1，若D==0则两直线平行，此时若A1C2-A2C1==0则两直线重合，否则没有交点
// 若D!=0则两直线相交，交点记为(x,y),x=(B1C2-B2C1)/D, y=(A2C1-A1C2)/D
func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
    if start1[0] > end1[0]   { start1, end1 = end1, start1 }
    if start2[0] > end2[0]   { start2, end2 = end2, start2 }
    if start1[0] > start2[0] { start1, end1, start2, end2 = start2, end2, start1, end1 }
    A1, B1, C1 := end1[1] - start1[1], start1[0] - end1[0], end1[0] * start1[1] - start1[0] * end1[1]
    A2, B2, C2 := end2[1] - start2[1], start2[0] - end2[0], end2[0] * start2[1] - start2[0] * end2[1]
    D1, D2 := A1 * B2 - A2 * B1, A1 * C2 - A2 * C1
    if D1 == 0 && D2 != 0 { // 平行
        return nil
    }
    if D1 == 0 && D2 == 0 { // 两直线重合
        if start1[0] != end1[0] { // 不垂直于x轴的情形
            if end1[0] >= start2[0] { // 线段重合
                return []float64{ float64(start2[0]), float64(start2[1]) }
            }
            return nil // 线段不重合
        }
        // 垂直于x轴的情形
        if start1[1] > end1[1] {
            start1, end1 = end1, start1
        }
        if start2[1] > end2[1] {
            start2, end2 = end2, start2
        }
        if start1[1] > start2[1] {
            start1, end1, start2, end2 = start2, end2, start1, end1
        }
        // 线段重合
        if end1[1] >= start2[1] {
            return []float64{float64(start2[0]), float64(start2[1])}
        }
        // 线段不重合
        return nil
    }
    // 余下的是直线相交的情况
    x, y := float64(B1 * C2 - B2 * C1) / float64(D1), float64(A2 * C1 - A1 * C2) / float64(D1)
    if  x >= float64(start1[0]) && 
        x <= float64(end1[0]) && 
        x >= float64(start2[0]) && 
        x <= float64(end2[0]) &&
        y >= float64(min(start1[1], end1[1])) &&
        y <= float64(max(start1[1], end1[1])) &&
        y >= float64(min(start2[1], end2[1])) &&
        y <= float64(max(start2[1], end2[1])) {
        return []float64{ x, y }
    }
    return nil
}

func main() {
    // Example 1:
    // Input: 
    // line1 = {0, 0}, {1, 0}
    // line2 = {1, 1}, {0, -1}
    // Output:  {0.5, 0}
    fmt.Println(intersection([]int{0, 0}, []int{1, 0}, []int{1, 1}, []int{0, -1})) // {0.5, 0}
    // Example 2:
    // Input: 
    // line1 = {0, 0}, {3, 3}
    // line2 = {1, 1}, {2, 2}
    // Output:  {1, 1}
    fmt.Println(intersection([]int{0, 0}, []int{3, 3}, []int{1, 1}, []int{2, 2})) // {1, 1}
    // Example 3:
    // Input: 
    // line1 = {0, 0}, {1, 1}
    // line2 = {1, 0}, {2, 1}
    // Output:  {}  (no intersection)
    fmt.Println(intersection([]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{2, 1})) // {} 
}