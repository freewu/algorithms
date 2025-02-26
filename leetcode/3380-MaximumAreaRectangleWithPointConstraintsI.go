package main

// 3380. Maximum Area Rectangle With Point Constraints I
// You are given an array points where points[i] = [xi, yi] represents the coordinates of a point on an infinite plane.

// Your task is to find the maximum area of a rectangle that:
//     1. Can be formed using four of these points as its corners.
//     2. Does not contain any other point inside or on its border.
//     3. Has its edges parallel to the axes.

// Return the maximum area that you can obtain or -1 if no such rectangle is possible.

// Example 1:
// Input: points = [[1,1],[1,3],[3,1],[3,3]]
// Output: 4
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/11/02/example1.png" />
// We can make a rectangle with these 4 points as corners and there is no other point that lies inside or on the border. Hence, the maximum possible area would be 4.

// Example 2:
// Input: points = [[1,1],[1,3],[3,1],[3,3],[2,2]]
// Output: -1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/11/02/example2.png" />
// There is only one rectangle possible is with points [1,1], [1,3], [3,1] and [3,3] but [2,2] will always lie inside it. Hence, returning -1.

// Example 3:
// Input: points = [[1,1],[1,3],[3,1],[3,3],[1,2],[3,2]]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/11/02/example3.png" />
// The maximum area rectangle is formed by the points [1,3], [1,2], [3,2], [3,3], which has an area of 2. Additionally, the points [1,1], [1,2], [3,1], [3,2] also form a valid rectangle with the same area.

// Constraints:
//     1 <= points.length <= 10
//     points[i].length == 2
//     0 <= xi, yi <= 100
//     All the given points are unique.

import "fmt"
import "sort"

func maxRectangleArea(points [][]int) int {
    // sort by horizontal then vertical position
    sort.Slice(points, func(i, j int) bool {
        if points[i][0] == points[j][0] { return points[i][1] < points[j][1] }
        return points[i][0] < points[j][0]
    })
    res := -1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(points) - 2; i++ {
        // continue if not alligned on y axis
        if points[i][0] != points[i+1][0] { continue }
        // find the next adjacent x-alligned points that fall into the vertical range of the selected leftside points
        // if they do not form a valid rectangle, give up because they will be inside the next rectangle guaranteed
        for j := i + 2; j < len(points) - 1; j++ {
            // break if a point exists within the vertical distance (it ruins this rectangle)...
            if points[j][1] >= points[i][1] && points[j][1] <= points[i+1][1] {
                // ...unless they are validly alligned and match the vertical borders
                if points[j][0] == points[j+1][0] && points[j][1] == points[i][1] &&
                    points[j + 1][1] == points[i+1][1] {
                    res = max(res, ((points[j][0] - points[i][0]) * (points[i+1][1] - points[i][1])))
                }
                break
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: points = [[1,1],[1,3],[3,1],[3,3]]
    // Output: 4
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/11/02/example1.png" />
    // We can make a rectangle with these 4 points as corners and there is no other point that lies inside or on the border. Hence, the maximum possible area would be 4.
    fmt.Println(maxRectangleArea([][]int{{1,1},{1,3},{3,1},{3,3}})) // 4
    // Example 2:
    // Input: points = [[1,1],[1,3],[3,1],[3,3],[2,2]]
    // Output: -1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/11/02/example2.png" />
    // There is only one rectangle possible is with points [1,1], [1,3], [3,1] and [3,3] but [2,2] will always lie inside it. Hence, returning -1.
    fmt.Println(maxRectangleArea([][]int{{1,1},{1,3},{3,1},{3,3},{2,2}})) // -1
    // Example 3:
    // Input: points = [[1,1],[1,3],[3,1],[3,3],[1,2],[3,2]]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/11/02/example3.png" />
    // The maximum area rectangle is formed by the points [1,3], [1,2], [3,2], [3,3], which has an area of 2. Additionally, the points [1,1], [1,2], [3,1], [3,2] also form a valid rectangle with the same area.
    fmt.Println(maxRectangleArea([][]int{{1,1},{1,3},{3,1},{3,3},{1,2},{3,2}})) // 2
}