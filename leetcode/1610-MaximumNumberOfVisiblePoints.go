package main

// 1610. Maximum Number of Visible Points
// You are given an array points, an integer angle, and your location, 
// where location = [posx, posy] and points[i] = [xi, yi] both denote integral coordinates on the X-Y plane.

// Initially, you are facing directly east from your position. 
// You cannot move from your position, but you can rotate. 
// In other words, posx and posy cannot be changed. 
// Your field of view in degrees is represented by angle, determining how wide you can see from any given view direction. 
// Let d be the amount in degrees that you rotate counterclockwise. 
// Then, your field of view is the inclusive range of angles [d - angle/2, d + angle/2].

// You can see some set of points if, for each point, the angle formed by the point, your position, 
// and the immediate east direction from your position is in your field of view.

// There can be multiple points at one coordinate. 
// There may be points at your location, and you can always see these points regardless of your rotation. 
// Points do not obstruct your vision to other points.

// Return the maximum number of points you can see.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/30/89a07e9b-00ab-4967-976a-c723b2aa8656.png" />
// Input: points = [[2,1],[2,2],[3,3]], angle = 90, location = [1,1]
// Output: 3
// Explanation: The shaded region represents your field of view. All points can be made visible in your field of view, including [3,3] even though [2,2] is in front and in the same line of sight.

// Example 2:
// Input: points = [[2,1],[2,2],[3,4],[1,1]], angle = 90, location = [1,1]
// Output: 4
// Explanation: All points can be made visible in your field of view, including the one at your location.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/09/30/5010bfd3-86e6-465f-ac64-e9df941d2e49.png" />
// Input: points = [[1,0],[2,1]], angle = 13, location = [1,1]
// Output: 1
// Explanation: You can only see one of the two points, as shown above.

// Constraints:
//     1 <= points.length <= 10^5
//     points[i].length == 2
//     location.length == 2
//     0 <= angle < 360
//     0 <= posx, posy, xi, yi <= 100

import "fmt"
import "math"
import "sort"

func visiblePoints(points [][]int, angle int, location []int) int {
    if angle == 360 { return len(points) }
    samePoint := 0
    angles := make([]float64, 0, len(points)) // Transform points into relative angles
    for _, v := range points {
        x, y  := float64(v[0] - location[0]), float64(v[1] - location[1])
        if x == 0 && y == 0 {
            samePoint++
        } else {
            angles = append(angles, math.Atan2(y, x))
        }
    }
    sort.Float64s(angles)
    mx, loop, addPi, vision := 0, false, 0.0, float64(angle) * math.Pi / 180.0
    for lead, follow := 0, 0; follow < len(angles); {
        if loop { addPi = 2 * math.Pi }
        if angles[lead] + addPi - angles[follow] <= vision {
            if lead >= follow {
                mx = max(mx, lead - follow + 1)
            } else {
                mx = max(mx, len(angles) - follow + 1 + lead)
            }
            lead++
        } else {
            follow++
        }
        if lead >= len(angles) {
            lead, loop = 0, true
        }
    }
    return samePoint + mx
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/30/89a07e9b-00ab-4967-976a-c723b2aa8656.png" />
    // Input: points = [[2,1],[2,2],[3,3]], angle = 90, location = [1,1]
    // Output: 3
    // Explanation: The shaded region represents your field of view. All points can be made visible in your field of view, including [3,3] even though [2,2] is in front and in the same line of sight.
    fmt.Println(visiblePoints([][]int{{2,1},{2,2},{3,3}}, 90, []int{1,1})) // 3
    // Example 2:
    // Input: points = [[2,1],[2,2],[3,4],[1,1]], angle = 90, location = [1,1]
    // Output: 4
    // Explanation: All points can be made visible in your field of view, including the one at your location.
    fmt.Println(visiblePoints([][]int{{2,1},{2,2},{3,4},{1,1}}, 90, []int{1,1})) // 4
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/09/30/5010bfd3-86e6-465f-ac64-e9df941d2e49.png" />
    // Input: points = [[1,0],[2,1]], angle = 13, location = [1,1]
    // Output: 1
    // Explanation: You can only see one of the two points, as shown above.
    fmt.Println(visiblePoints([][]int{{1,0},{2,1}}, 13, []int{1,1})) // 3
}