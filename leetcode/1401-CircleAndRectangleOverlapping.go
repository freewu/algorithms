package main

// 1401. Circle and Rectangle Overlapping
// You are given a circle represented as (radius, xCenter, yCenter) and an axis-aligned rectangle represented as (x1, y1, x2, y2), 
// where (x1, y1) are the coordinates of the bottom-left corner, and (x2, y2) are the coordinates of the top-right corner of the rectangle.

// Return true if the circle and rectangle are overlapped otherwise return false. 
// In other words, check if there is any point (xi, yi) that belongs to the circle and the rectangle at the same time.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/02/20/sample_4_1728.png" />
// Input: radius = 1, xCenter = 0, yCenter = 0, x1 = 1, y1 = -1, x2 = 3, y2 = 1
// Output: true
// Explanation: Circle and rectangle share the point (1,0).

// Example 2:
// Input: radius = 1, xCenter = 1, yCenter = 1, x1 = 1, y1 = -3, x2 = 2, y2 = -1
// Output: false

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/02/20/sample_2_1728.png" />
// Input: radius = 1, xCenter = 0, yCenter = 0, x1 = -1, y1 = 0, x2 = 0, y2 = 1
// Output: true

// Constraints:
//     1 <= radius <= 2000
//     -10^4 <= xCenter, yCenter <= 10^4
//     -10^4 <= x1 < x2 <= 10^4
//     -10^4 <= y1 < y2 <= 10^4

import "fmt"
import "math"

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
    if xCenter >= x1 && xCenter <= x2 && 
       yCenter >= y1 && yCenter <= y2 {
        return true
    } else if xCenter >= x1 && xCenter <= x2 && (yCenter < y1 || yCenter > y2) {
        if (y1 - yCenter > 0 && y1 - yCenter <= radius) || 
           (yCenter - y2 > 0 && yCenter - y2 <= radius) {
            return true
        }
    } else if (xCenter < x1 || xCenter > x2) && yCenter >= y1 && yCenter <= y2 {
        if (x1 - xCenter > 0 && x1 - xCenter <= radius) || 
           (xCenter - x2 > 0 && xCenter - x2 <= radius) {
            return true
        }
    } else if ((xCenter - x1) * (xCenter - x1) + (yCenter - y1) * (yCenter - y1) <= radius * radius) ||
              ((xCenter - x2) * (xCenter - x2) + (yCenter - y1) * (yCenter - y1) <= radius * radius) ||
              ((xCenter - x1) * (xCenter - x1) + (yCenter - y2) * (yCenter - y2) <= radius * radius) ||
              ((xCenter - x2) * (xCenter - x2) + (yCenter - y2) * (yCenter - y2) <= radius * radius) {
        return true
    }
    return false
}

func checkOverlap1(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
    for x := x1; x <= x2; x++ {
        for y := y1; y <= y2; y++ {
            a, b := x - xCenter, y - yCenter
            if float64(radius) >= math.Sqrt(float64(a * a + b * b)) {
                return true
            }
        }
    }
    return false
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/02/20/sample_4_1728.png" />
    // Input: radius = 1, xCenter = 0, yCenter = 0, x1 = 1, y1 = -1, x2 = 3, y2 = 1
    // Output: true
    // Explanation: Circle and rectangle share the point (1,0).
    fmt.Println(checkOverlap(1, 0, 0, 1, -1, 3, 1)) // true
    // Example 2:
    // Input: radius = 1, xCenter = 1, yCenter = 1, x1 = 1, y1 = -3, x2 = 2, y2 = -1
    // Output: false
    fmt.Println(checkOverlap(1, 1, 1, 1, -3, 2, -1)) // false
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/02/20/sample_2_1728.png" />
    // Input: radius = 1, xCenter = 0, yCenter = 0, x1 = -1, y1 = 0, x2 = 0, y2 = 1
    // Output: true
    fmt.Println(checkOverlap(1, 0, 0, -1, 0, 0, 1)) // true

    fmt.Println(checkOverlap1(1, 0, 0, 1, -1, 3, 1)) // true
    fmt.Println(checkOverlap1(1, 1, 1, 1, -3, 2, -1)) // false
    fmt.Println(checkOverlap1(1, 0, 0, -1, 0, 0, 1)) // true
}