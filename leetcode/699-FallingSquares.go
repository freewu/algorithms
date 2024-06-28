package main

// 699. Falling Squares
// There are several squares being dropped onto the X-axis of a 2D plane.

// You are given a 2D integer array positions where positions[i] = [lefti, sideLengthi] 
// represents the ith square with a side length of sideLengthi 
// that is dropped with its left edge aligned with X-coordinate lefti.

// Each square is dropped one at a time from a height above any landed squares. 
// It then falls downward (negative Y direction) until it either lands on the top side of another square or on the X-axis. 
// A square brushing the left/right side of another square does not count as landing on it. 
// Once it lands, it freezes in place and cannot be moved.

// After each square is dropped, you must record the height of the current tallest stack of squares.
// Return an integer array ans where ans[i] represents the height described above after dropping the ith square.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/28/fallingsq1-plane.jpg" />
// Input: positions = [[1,2],[2,3],[6,1]]
// Output: [2,5,5]
// Explanation:
// After the first drop, the tallest stack is square 1 with a height of 2.
// After the second drop, the tallest stack is squares 1 and 2 with a height of 5.
// After the third drop, the tallest stack is still squares 1 and 2 with a height of 5.
// Thus, we return an answer of [2, 5, 5].

// Example 2:
// Input: positions = [[100,100],[200,100]]
// Output: [100,100]
// Explanation:
// After the first drop, the tallest stack is square 1 with a height of 100.
// After the second drop, the tallest stack is either square 1 or square 2, both with heights of 100.
// Thus, we return an answer of [100, 100].
// Note that square 2 only brushes the right side of square 1, which does not count as landing on it.
 
// Constraints:
//     1 <= positions.length <= 1000
//     1 <= lefti <= 10^8
//     1 <= sideLengthi <= 10^6

import "fmt"

func fallingSquares(positions [][]int) []int {
    heights, res, globalHighest := [][]int{}, []int{}, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, position := range positions {
        start, end := position[0], position[0]+position[1]-1
        highest := 0
        for _, height := range heights {
            if (height[0] <= end && start <= height[1]) {
                highest = max(highest, height[2])
            }
        }
        globalHighest = max(globalHighest, highest+position[1])
        res = append(res, globalHighest)
        heights = append(heights, []int { start, end, highest+position[1] })
    }
    return res
}

func fallingSquares1(positions [][]int) []int {
    n := len(positions)
    res := make([]int, n)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, p := range positions {
        left1, right1 := p[0], p[0]+p[1]-1
        res[i] = p[1]
        for j, q := range positions[:i] {
            left2, right2 := q[0], q[0]+q[1]-1
            if right1 >= left2 && right2 >= left1 {
                res[i] = max(res[i], res[j]+p[1])
            }
        }
    }
    for i := 1; i < n; i++ {
        res[i] = max(res[i], res[i-1])
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/28/fallingsq1-plane.jpg" />
    // Input: positions = [[1,2],[2,3],[6,1]]
    // Output: [2,5,5]
    // Explanation:
    // After the first drop, the tallest stack is square 1 with a height of 2.
    // After the second drop, the tallest stack is squares 1 and 2 with a height of 5.
    // After the third drop, the tallest stack is still squares 1 and 2 with a height of 5.
    // Thus, we return an answer of [2, 5, 5].
    fmt.Println(fallingSquares([][]int{{1,2},{2,3},{6,1}})) // [2,5,5]
    // Example 2:
    // Input: positions = [[100,100],[200,100]]
    // Output: [100,100]
    // Explanation:
    // After the first drop, the tallest stack is square 1 with a height of 100.
    // After the second drop, the tallest stack is either square 1 or square 2, both with heights of 100.
    // Thus, we return an answer of [100, 100].
    // Note that square 2 only brushes the right side of square 1, which does not count as landing on it.
    fmt.Println(fallingSquares([][]int{{100,100},{200,100}})) // [100,100]

    fmt.Println(fallingSquares1([][]int{{1,2},{2,3},{6,1}})) // [2,5,5]
    fmt.Println(fallingSquares1([][]int{{100,100},{200,100}})) // [100,100]
}