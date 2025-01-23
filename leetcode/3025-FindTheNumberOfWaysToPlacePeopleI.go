package main

// 3025. Find the Number of Ways to Place People I
// You are given a 2D array points of size n x 2 representing integer coordinates of some points on a 2D plane, 
// where points[i] = [xi, yi].

// Count the number of pairs of points (A, B), where:
//     1. A is on the upper left side of B, and
//     2. there are no other points in the rectangle (or line) they make (including the border).

// Return the count.

// Example 1:
// Input: points = [[1,1],[2,2],[3,3]]
// Output: 0
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/01/04/example1alicebob.png" />
// There is no way to choose A and B so A is on the upper left side of B.

// Example 2:
// Input: points = [[6,2],[4,4],[2,6]]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/25/t2.jpg" />
// The left one is the pair (points[1], points[0]), where points[1] is on the upper left side of points[0] and the rectangle is empty.
// The middle one is the pair (points[2], points[1]), same as the left one it is a valid pair.
// The right one is the pair (points[2], points[0]), where points[2] is on the upper left side of points[0], but points[1] is inside the rectangle so it's not a valid pair.

// Example 3:
// Input: points = [[3,1],[1,3],[1,1]]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/25/t3.jpg" />
// The left one is the pair (points[2], points[0]), where points[2] is on the upper left side of points[0] and there are no other points on the line they form. Note that it is a valid state when the two points form a line.
// The middle one is the pair (points[1], points[2]), it is a valid pair same as the left one.
// The right one is the pair (points[1], points[0]), it is not a valid pair as points[2] is on the border of the rectangle.

// Constraints:
//     2 <= n <= 50
//     points[i].length == 2
//     0 <= points[i][0], points[i][1] <= 50
//     All points[i] are distinct.

import "fmt"
import "sort"

func numberOfPairs(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        if points[i][1] == points[j][1] { return points[i][0] < points[j][0] }
        return points[i][1] > points[j][1]
    })
    res, n := 0, len(points)
    for i := 0; i < n; i++ {
        prev := 1 << 31
        for j := i + 1; j < n; j++ {
            diff := points[j][0] - points[i][0]
            if diff >= 0 && diff < prev {
                prev = diff
                res++
            }
        }
    }
    return res
}

func numberOfPairs1(points [][]int) int {
    res, n := 0, len(points)
    for i := 0; i < n; i++ {
        x1, y1 := points[i][0], points[i][1]
        for j := 0; j < n; j++ {
            if i == j { continue }
            f, x2, y2 := 0, points[j][0], points[j][1]
            if y1 >= y2 && x1 <= x2 {
                for k := 0; k < n; k++ {
                    if k == i || k == j { continue }
                    x3, y3 := points[k][0], points[k][1]
                    if x3 >= x1 && x3 <= x2 && y3 >= y2 && y3 <= y1 {
                        f = 1
                        break
                    }
                }
                if f == 0 {
                    res++
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: points = [[1,1],[2,2],[3,3]]
    // Output: 0
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/01/04/example1alicebob.png" />
    // There is no way to choose A and B so A is on the upper left side of B.
    fmt.Println(numberOfPairs([][]int{{1,1},{2,2},{3,3}})) // 0
    // Example 2:
    // Input: points = [[6,2],[4,4],[2,6]]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/25/t2.jpg" />
    // The left one is the pair (points[1], points[0]), where points[1] is on the upper left side of points[0] and the rectangle is empty.
    // The middle one is the pair (points[2], points[1]), same as the left one it is a valid pair.
    // The right one is the pair (points[2], points[0]), where points[2] is on the upper left side of points[0], but points[1] is inside the rectangle so it's not a valid pair.
    fmt.Println(numberOfPairs([][]int{{6,2},{4,4},{2,6}})) // 2
    // Example 3:
    // Input: points = [[3,1],[1,3],[1,1]]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/25/t3.jpg" />
    // The left one is the pair (points[2], points[0]), where points[2] is on the upper left side of points[0] and there are no other points on the line they form. Note that it is a valid state when the two points form a line.
    // The middle one is the pair (points[1], points[2]), it is a valid pair same as the left one.
    // The right one is the pair (points[1], points[0]), it is not a valid pair as points[2] is on the border of the rectangle.
    fmt.Println(numberOfPairs([][]int{{3,1},{1,3},{1,1}})) // 2

    fmt.Println(numberOfPairs1([][]int{{1,1},{2,2},{3,3}})) // 0
    fmt.Println(numberOfPairs1([][]int{{6,2},{4,4},{2,6}})) // 2
    fmt.Println(numberOfPairs1([][]int{{3,1},{1,3},{1,1}})) // 2
}