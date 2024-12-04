package main

// 2249. Count Lattice Points Inside a Circle
// Given a 2D integer array circles where circles[i] = [xi, yi, ri] represents the center (xi, yi) 
// and radius ri of the ith circle drawn on a grid, 
// return the number of lattice points that are present inside at least one circle.

// Note:
//     A lattice point is a point with integer coordinates.
//     Points that lie on the circumference of a circle are also considered to be inside it.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/02/exa-11.png" />
// Input: circles = [[2,2,1]]
// Output: 5
// Explanation:
// The figure above shows the given circle.
// The lattice points present inside the circle are (1, 2), (2, 1), (2, 2), (2, 3), and (3, 2) and are shown in green.
// Other points such as (1, 1) and (1, 3), which are shown in red, are not considered inside the circle.
// Hence, the number of lattice points present inside at least one circle is 5.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/02/exa-22.png" />
// Input: circles = [[2,2,2],[3,4,1]]
// Output: 16
// Explanation:
// The figure above shows the given circles.
// There are exactly 16 lattice points which are present inside at least one circle. 
// Some of them are (0, 2), (2, 0), (2, 4), (3, 2), and (4, 4).

// Constraints:
//     1 <= circles.length <= 200
//     circles[i].length == 3
//     1 <= xi, yi <= 100
//     1 <= ri <= min(xi, yi)

import "fmt"

// brute force
func countLatticePoints(circles [][]int) int {
    type Point struct { X, Y int }
    res, seen := 0, make(map[Point]bool)
    for _, circle := range circles {
        x, y, r := circle[0], circle[1], circle[2]
        minX, minY, maxX, maxY, rSquare := x - r, y - r, x + r, y + r, r * r
        for i := minX; i <= maxX; i++ {
            for j := minY; j <= maxY; j++ {
                newPoint := Point{X: i, Y: j}
                if !seen[newPoint] {
                    if ((i - x) * (i - x) + (j - y) * (j - y)) <= rSquare {
                        seen[newPoint] = true
                        res++
                    }
                }
            }
        }
    }
    return res
}

func countLatticePoints1(circles [][]int) int {
    seen := make([][]int, 201)
    for i := 0; i < 201; i++ {
        seen[i] = make([]int, 201)
    }
    res := 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, circle := range circles {
        x, y, r := circle[0], circle[1], circle[2]
        for i := x - r; i <= x + r; i++ {
            for j := y - r; j <= y + r; j++ {
                a, b := abs(i - x), abs(j - y)
                if (a * a + b * b) <= (r * r) && seen[i][j] == 0 {
                    res++
                    seen[i][j] = 1
                }
            }
        }
    }
    return res
}

func countLatticePoints2(circles [][]int) (ans int) {
    res, mx, my := 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, c := range circles {
        mx, my = max(mx, c[0] + c[2]), max(my, c[1] + c[2])
    }
    for i := 0; i <= mx; i++ {
        for j := 0; j <= my; j++ {
            for _, c := range circles {
                dx, dy, r := i - c[0], j - c[1], c[2]
                if dx * dx + dy * dy <=  r * r {
                    res++
                    break
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/02/exa-11.png" />
    // Input: circles = [[2,2,1]]
    // Output: 5
    // Explanation:
    // The figure above shows the given circle.
    // The lattice points present inside the circle are (1, 2), (2, 1), (2, 2), (2, 3), and (3, 2) and are shown in green.
    // Other points such as (1, 1) and (1, 3), which are shown in red, are not considered inside the circle.
    // Hence, the number of lattice points present inside at least one circle is 5.
    fmt.Println(countLatticePoints([][]int{{2,2,1}})) // 5
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/02/exa-22.png" />
    // Input: circles = [[2,2,2],[3,4,1]]
    // Output: 16
    // Explanation:
    // The figure above shows the given circles.
    // There are exactly 16 lattice points which are present inside at least one circle. 
    // Some of them are (0, 2), (2, 0), (2, 4), (3, 2), and (4, 4).
    fmt.Println(countLatticePoints([][]int{{2,2,2},{3,4,1}})) // 16

    fmt.Println(countLatticePoints1([][]int{{2,2,1}})) // 5
    fmt.Println(countLatticePoints1([][]int{{2,2,2},{3,4,1}})) // 16

    fmt.Println(countLatticePoints2([][]int{{2,2,1}})) // 5
    fmt.Println(countLatticePoints2([][]int{{2,2,2},{3,4,1}})) // 16
}