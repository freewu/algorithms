package main

// 3009. Maximum Number of Intersections on the Chart
// There is a line chart consisting of n points connected by line segments. 
// You are given a 1-indexed integer array y. The kth point has coordinates (k, y[k]). 
// There are no horizontal lines; that is, no two consecutive points have the same y-coordinate.

// We can draw an infinitely long horizontal line. 
// Return the maximum number of points of intersection of the line with the chart.

// Example 1:
// <img src="https://assets.leetcode.com/static_assets/others/20231208-020549.jpeg" />
// Input: y = [1,2,1,2,1,3,2]
// Output: 5
// Explanation: As you can see in the image above, the line y = 1.5 has 5 intersections with the chart (in red crosses). You can also see the line y = 2 which intersects the chart in 4 points (in red crosses). It can be shown that there is no horizontal line intersecting the chart at more than 5 points. So the answer would be 5.

// Example 2:
// <img src="https://assets.leetcode.com/static_assets/others/20231208-020557.jpeg" />
// Input: y = [2,1,3,4,5]
// Output: 2
// Explanation: As you can see in the image above, the line y = 1.5 has 2 intersections with the chart (in red crosses). You can also see the line y = 2 which intersects the chart in 2 points (in red crosses). It can be shown that there is no horizontal line intersecting the chart at more than 2 points. So the answer would be 2.
 
// Constraints:
//     2 <= y.length <= 10^5
//     1 <= y[i] <= 10^9
//     y[i] != y[i + 1] for i in range [1, n - 1]

import "fmt"
import "sort"

func maxIntersectionCount(y []int) int {
    type Point struct { v, delta int }
    points, left, right, n := []*Point{}, 0, 0, len(y)
    for i := 1; i < n; i++ {
        if y[i] > y[i - 1] {
            left, right = y[i - 1] * 2, y[i] * 2 - 1
        } else {
            left, right = y[i] * 2 + 1, y[i - 1] * 2
        }
        points = append(points, &Point{ left, -1 })
        points = append(points, &Point{ right, 1 })
    }
    points = append(points, &Point{ y[n - 1] * 2, -1 })
    points = append(points, &Point{ y[n - 1] * 2, 1 })

    sort.Slice(points, func(i, j int) bool {
        if  points[i].v == points[j].v {
            return points[i].delta  < points[j].delta
        }
        return points[i].v  < points[j].v
    })
    res, now := 0, 0
    for _, p := range points {
        now -= p.delta
        res = max(now, res)
    }
    return res
}

// class Solution(object):
//     def maxIntersectionCount(self, y):
//         """
//         :type y: List[int]
//         :rtype: int
//         """

//         points = []
//         for index in range(1, len(y)):
//             if y[index] > y[index - 1]:
//                 left, right = y[index - 1] * 2, y[index] * 2 - 1
//             else:
//                 left, right = y[index] * 2 + 1, y[index - 1] * 2
            
//             points.append([left, -1])
//             points.append([right, 1])
        
//         points.append([y[-1] * 2, -1])
//         points.append([y[-1] * 2, 1])

//         points.sort()
//         now = 0
//         res = 0
//         for v, delta in points:
//             now -= delta
//             res = max(now, res)
        
//         return res

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/static_assets/others/20231208-020549.jpeg" />
    // Input: y = [1,2,1,2,1,3,2]
    // Output: 5
    // Explanation: As you can see in the image above, the line y = 1.5 has 5 intersections with the chart (in red crosses). You can also see the line y = 2 which intersects the chart in 4 points (in red crosses). It can be shown that there is no horizontal line intersecting the chart at more than 5 points. So the answer would be 5.
    fmt.Println(maxIntersectionCount([]int{1,2,1,2,1,3,2})) // 5
    // Example 2:
    // <img src="https://assets.leetcode.com/static_assets/others/20231208-020557.jpeg" />
    // Input: y = [2,1,3,4,5]
    // Output: 2
    // Explanation: As you can see in the image above, the line y = 1.5 has 2 intersections with the chart (in red crosses). You can also see the line y = 2 which intersects the chart in 2 points (in red crosses). It can be shown that there is no horizontal line intersecting the chart at more than 2 points. So the answer would be 2.
    fmt.Println(maxIntersectionCount([]int{2,1,3,4,5})) // 2
}