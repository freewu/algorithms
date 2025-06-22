package main

// 3588. Find Maximum Area of a Triangle
// You are given a 2D array coords of size n x 2, representing the coordinates of n points in an infinite Cartesian plane.

// Find twice the maximum area of a triangle with its corners at any three elements from coords, such that at least one side of this triangle is parallel to the x-axis or y-axis. 
// Formally, if the maximum area of such a triangle is A, return 2 * A.

// If no such triangle exists, return -1.

// Note that a triangle cannot have zero area.

// Example 1:
// Input: coords = [[1,1],[1,2],[3,2],[3,3]]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/19/image-20250420010047-1.png" />
// The triangle shown in the image has a base 1 and height 2. Hence its area is 1/2 * base * height = 1.

// Example 2:
// Input: coords = [[1,1],[2,2],[3,3]]
// Output: -1
// Explanation:
// The only possible triangle has corners (1, 1), (2, 2), and (3, 3). None of its sides are parallel to the x-axis or the y-axis.

// Constraints:
//     1 <= n == coords.length <= 10^5
//     1 <= coords[i][0], coords[i][1] <= 10^6
//     All coords[i] are unique.

import "fmt"
import "slices"

func maxArea(coords [][]int) int64 {
    res := 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    calc := func() {
        minX, maxX := 1 << 31, 0
        minY := map[int]int{}
        maxY := map[int]int{}
        for _, p := range coords {
            x, y := p[0], p[1]
            minX, maxX = min(minX, x), max(maxX, x)
            maxY[x] = max(maxY[x], y)
            mn, ok := minY[x]
            if !ok {
                minY[x] = y
            } else {
                minY[x] = min(mn, y)
            }
        }
        for x, y := range minY {
            res = max(res, (maxY[x] - y) * max(maxX - x, x - minX))
        }
    }
    calc()
    for _, p := range coords {
        p[0], p[1] = p[1], p[0]
    }
    calc()
    if res == 0 {
        return -1
    }
    return int64(res)
}

func maxArea1(coords [][]int) int64 {
    if len(coords) < 3 { return -1 }
    up, down, left, right  := make([]int, 2), make([]int, 2), make([]int, 2),  make([]int, 2)
    for _, c := range coords {
        if up[0] == 0    || c[1] > up[1]    { copy(up, c)  }
        if down[0] == 0  || c[1] < down[1]  { copy(down, c) }
        if left[0] == 0  || c[0] < left[0]  { copy(left, c)  }
        if right[0] == 0 || c[0] > right[0] { copy(right, c) }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, n := 0, len(coords)
    slices.SortFunc(coords, func(a, b []int) int {
        if a[0] == b[0] {
            return a[1] - b[1]
        }
        return a[0] - b[0]
    })
    i, j := 0, 1
    for i < n {
        for j < n && coords[j][0] == coords[i][0] {
            j++
        }
        if j > i + 1 {
            res = max(res,(coords[j-1][1] - coords[i][1]) * max(coords[i][0]-left[0], right[0]-coords[i][0]))
        }
            j++
            i = j - 1
    }
    slices.SortFunc(coords, func(a, b []int) int {
        if a[1] == b[1] { return a[0] - b[0] }
        return a[1] - b[1]
    })
    i, j = 0, 1
    for i < n {
        for j < n && coords[j][1] == coords[i][1] {
            j++
        }
        if j > i + 1 {
            res = max(res,(coords[j-1][0] - coords[i][0]) * max(coords[i][1]-down[1], up[1]-coords[i][1]))
        }
            j++
            i = j - 1
    }
    if res == 0 { return -1 }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: coords = [[1,1],[1,2],[3,2],[3,3]]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/04/19/image-20250420010047-1.png" />
    // The triangle shown in the image has a base 1 and height 2. Hence its area is 1/2 * base * height = 1.
    fmt.Println(maxArea([][]int{{1,1},{1,2},{3,2},{3,3}})) // 2
    // Example 2:
    // Input: coords = [[1,1],[2,2],[3,3]]
    // Output: -1
    // Explanation:
    // The only possible triangle has corners (1, 1), (2, 2), and (3, 3). None of its sides are parallel to the x-axis or the y-axis.
    fmt.Println(maxArea([][]int{{1,1},{2,2},{3,3}})) // -1

    fmt.Println(maxArea1([][]int{{1,1},{1,2},{3,2},{3,3}})) // 2
    fmt.Println(maxArea1([][]int{{1,1},{2,2},{3,3}})) // -1
}