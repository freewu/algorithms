package main

// 3453. Separate Squares I
// You are given a 2D integer array squares. 
// Each squares[i] = [xi, yi, li] represents the coordinates of the bottom-left point and the side length of a square parallel to the x-axis.

// Find the minimum y-coordinate value of a horizontal line such that the total area of the squares above the line equals the total area of the squares below the line.

// Answers within 10^-5 of the actual answer will be accepted.

// Note: Squares may overlap. Overlapping areas should be counted multiple times.

// Example 1:
// Input: squares = [[0,0,1],[2,2,1]]
// Output: 1.00000
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/01/06/4062example1drawio.png" />
// Any horizontal line between y = 1 and y = 2 will have 1 square unit above it and 1 square unit below it. 
// The lowest option is 1.

// Example 2:
// Input: squares = [[0,0,2],[1,1,1]]
// Output: 1.16667
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/01/15/4062example2drawio.png" />
// The areas are:
// Below the line: 7/6 * 2 (Red) + 1/6 (Blue) = 15/6 = 2.5.
// Above the line: 5/6 * 2 (Red) + 5/6 (Blue) = 15/6 = 2.5.
// Since the areas above and below the line are equal, the output is 7/6 = 1.16667.

// Constraints:
//     1 <= squares.length <= 5 * 10^4
//     squares[i] = [xi, yi, li]
//     squares[i].length == 3
//     0 <= xi, yi <= 10^9
//     1 <= li <= 10^9
//     The total area of all the squares will not exceed 10^12.

import "fmt"
import "math"
import "sort"

func separateSquares(squares [][]int) float64 {
    start, end := 0.0, 2e9
    helper := func(mid float64) int {
        upper, lower := 0.0, 0.0
        for _, sq := range squares {
            y, l := float64(sq[1]), float64(sq[2])
            if y > mid {
                upper += l * l
            } else if y+l <= mid {
                lower += l * l
            } else {
                lower += (mid - y) * l
                upper += (l - (mid - y)) * l
            }
        }
        if upper == lower { return 0 }
        if upper > lower  { return 1 }
        return -1
    }
    for (end - start) > 1e-5 {
        mid := (start + end) / 2.0
        if helper(mid) == 1 {
            start = mid
        } else {
            end = mid
        }
    }
    return math.Round(end * 1e5) / 1e5
}

func separateSquares1(squares [][]int) float64 {
    sum, mx := 0. , 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, sq := range squares {
        sum += float64(sq[2] * sq[2])
        mx = max(mx, sq[1] + sq[2])
    }
    calcArea := func(y int) float64 {
        res := 0.
        for _, sq := range squares {
            yi := sq[1]
            if yi < y {
                l := sq[2]
                res += float64(l * min(y - yi, l))
            }
        }
        return res
    }
    y := sort.Search(mx, func(y int) bool { 
        return calcArea(y) >= sum / 2 
    })
    areaY := calcArea(y)
    sumL := areaY - calcArea(y - 1)
    return float64(y) - (areaY - sum / 2) / sumL
}

func main() {
    // Example 1:
    // Input: squares = [[0,0,1],[2,2,1]]
    // Output: 1.00000
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/01/06/4062example1drawio.png" />
    // Any horizontal line between y = 1 and y = 2 will have 1 square unit above it and 1 square unit below it. 
    // The lowest option is 1.
    fmt.Println(separateSquares([][]int{{0,0,1},{2,2,1}})) // 1.00000
    // Example 2:
    // Input: squares = [[0,0,2],[1,1,1]]
    // Output: 1.16667
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/01/15/4062example2drawio.png" />
    // The areas are:
    // Below the line: 7/6 * 2 (Red) + 1/6 (Blue) = 15/6 = 2.5.
    // Above the line: 5/6 * 2 (Red) + 5/6 (Blue) = 15/6 = 2.5.
    // Since the areas above and below the line are equal, the output is 7/6 = 1.16667.
    fmt.Println(separateSquares([][]int{{0,0,2},{1,1,1}})) // 1.16667

    fmt.Println(separateSquares([][]int{{1,2,3},{1,2,3}})) // 3.50001
    fmt.Println(separateSquares([][]int{{1,2,3},{3,2,1}})) // 3.33333
    fmt.Println(separateSquares([][]int{{3,2,1},{1,2,3}})) // 3.33333
    fmt.Println(separateSquares([][]int{{3,2,1},{3,2,1}})) // 2.5
    fmt.Println(separateSquares([][]int{{7,8,9},{3,2,1}})) // 12.44445
    fmt.Println(separateSquares([][]int{{3,2,1},{3,2,1}})) // 2.5
    fmt.Println(separateSquares([][]int{{7,8,9},{7,8,9}})) // 12.5
    fmt.Println(separateSquares([][]int{{3,2,1},{7,8,9}})) // 12.44445

    fmt.Println(separateSquares1([][]int{{0,0,1},{2,2,1}})) // 1.00000
    fmt.Println(separateSquares1([][]int{{0,0,2},{1,1,1}})) // 1.16667
    fmt.Println(separateSquares1([][]int{{1,2,3},{1,2,3}})) // 3.50001
    fmt.Println(separateSquares1([][]int{{1,2,3},{3,2,1}})) // 3.33333
    fmt.Println(separateSquares1([][]int{{3,2,1},{1,2,3}})) // 3.33333
    fmt.Println(separateSquares1([][]int{{3,2,1},{3,2,1}})) // 2.5
    fmt.Println(separateSquares1([][]int{{7,8,9},{3,2,1}})) // 12.44445
    fmt.Println(separateSquares1([][]int{{3,2,1},{3,2,1}})) // 2.5
    fmt.Println(separateSquares1([][]int{{7,8,9},{7,8,9}})) // 12.5
    fmt.Println(separateSquares1([][]int{{3,2,1},{7,8,9}})) // 12.44445
}