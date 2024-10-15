package main

// 1453. Maximum Number of Darts Inside of a Circular Dartboard
// Alice is throwing n darts on a very large wall. 
// You are given an array darts where darts[i] = [xi, yi] is the position of the ith dart that Alice threw on the wall.

// Bob knows the positions of the n darts on the wall. 
// He wants to place a dartboard of radius r on the wall so that the maximum number of darts that Alice throws lie on the dartboard.

// Given the integer r, return the maximum number of darts that can lie on the dartboard.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/04/29/sample_1_1806.png" />
// Input: darts = [[-2,0],[2,0],[0,2],[0,-2]], r = 2
// Output: 4
// Explanation: Circle dartboard with center in (0,0) and radius = 2 contain all points.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/04/29/sample_2_1806.png" />
// Input: darts = [[-3,0],[3,0],[2,6],[5,4],[0,9],[7,8]], r = 5
// Output: 5
// Explanation: Circle dartboard with center in (0,4) and radius = 5 contain all points except the point (7,8).

// Constraints:
//     1 <= darts.length <= 100
//     darts[i].length == 2
//     -10^4 <= xi, yi <= 10^4
//     All the darts are unique
//     1 <= r <= 5000

import "fmt"
import "math"

func numPoints(points [][]int, r int) int {
    n, res := len(points), 1 // 肯定有一个点在他的圆内
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 是否在同一个圆里
    calTogetherCir := func(pointX, pointY []int,  r int ) bool {
        dis := (pointX[0] - pointY[0]) * (pointX[0] - pointY[0]) + (pointX[1] - pointY[1]) * (pointX[1] - pointY[1])
        return 4 * r * r >= dis
    }
    for i := 0; i < n; i++ {
        for  j := 0; j < n; j++ {
            if points[i][0] == points[j][0] && points[i][1] == points[j][1] { continue }
            if !calTogetherCir(points[i], points[j], r) { continue }
            //(cx, cy)
            // (cx - midX) ^2 + (cy - midY) ^ 2 = high ^ 2
            // ji向量是(points[i][0] - points[j][0], points[i][1] - points[j][1])
            // ji向量和mc向量（圆心和中点相乘为0） ==> mc向量和(points[i][1] - points[j][1], points[j][0] - points[i][0]) 在同一个方向上
            //(points[i][1] - points[j][1]) * (points[i][0] - points[j][0]) + (points[j][0] - points[i][0]) * (points[i][1] - points[j][1])
            // (points[i][1] - points[j][1]) * (points[i][0] - points[j][0]) - (points[i][0] - points[j][0]) * (points[i][1] - points[j][1]) = 0
            //(cx - midx) / (points[i][1] - points[j][1]) = (cy - midy) / (points[j][0] - points[i][0])
            midx, midy := float64(points[i][0] + points[j][0]) * 0.5, float64(points[i][1] + points[j][1]) * 0.5
            // 长度的平方
            midDis := (midx - float64(points[i][0])) * (midx - float64(points[i][0])) + (midy - float64(points[i][1])) * (midy - float64(points[i][1]))
            // 高度的距离
            high := math.Sqrt(float64(r * r) - midDis)
            dx, dy := float64(points[j][0] - points[i][0]), float64(points[i][1] - points[j][1])
            //(cx - midx) /dy = (cy - midy) / dx  ==== >  (cx - midx)  = (cy - midy) / dx * dy
            //(cx - midX) ^2 + (cy - midY) ^ 2 = high ^ 2 ====> ((cy - midy) / dx * dy) ^ 2 + (cy - midY) ^ 2 = high ^ 2
            //(cy - midy) = t ====> (t * dy / dx) ^ 2 + t ^ 2 = high ====> t ^ 2 = high ^ 2 * dx ^ 2 /(dx ^ 2 + dy ^ 2)
            //==>t = high  * dx / sqrt(dx ^ 2 + dy ^ 2);
            disxy := math.Sqrt(dx * dx + dy * dy)
            cx, cy := high * dy / disxy  + midx, high * dx / disxy + midy
            count := 0
            for k := 0; k < n; k++ {
                if (float64(points[k][0]) - cx) * (float64(points[k][0]) - cx) + (float64(points[k][1]) - cy) * (float64(points[k][1]) - cy) <= float64(r * r) {
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
    // <img src="https://assets.leetcode.com/uploads/2020/04/29/sample_1_1806.png" />
    // Input: darts = [[-2,0],[2,0],[0,2],[0,-2]], r = 2
    // Output: 4
    // Explanation: Circle dartboard with center in (0,0) and radius = 2 contain all points.
    fmt.Println(numPoints([][]int{{-2,0},{2,0},{0,2},{0,-2}}, 2)) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/04/29/sample_2_1806.png" />
    // Input: darts = [[-3,0],[3,0],[2,6],[5,4],[0,9],[7,8]], r = 5
    // Output: 5
    // Explanation: Circle dartboard with center in (0,4) and radius = 5 contain all points except the point (7,8).
    fmt.Println(numPoints([][]int{{-3,0},{3,0},{2,6},{5,4},{0,9},{7,8}}, 5)) // 5
}