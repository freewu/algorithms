package main

// 3047. Find the Largest Area of Square Inside Two Rectangles
// There exist n rectangles in a 2D plane with edges parallel to the x and y axis. 
// You are given two 2D integer arrays bottomLeft and topRight 
// where bottomLeft[i] = [a_i, b_i] and topRight[i] = [c_i, d_i] represent the bottom-left and top-right coordinates of the ith rectangle, respectively.

// You need to find the maximum area of a square that can fit inside the intersecting region of at least two rectangles. 
// Return 0 if such a square does not exist.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/01/05/example12.png" />
// Input: bottomLeft = [[1,1],[2,2],[3,1]], topRight = [[3,3],[4,4],[6,6]]
// Output: 1
// Explanation:
// A square with side length 1 can fit inside either the intersecting region of rectangles 0 and 1 or the intersecting region of rectangles 1 and 2. Hence the maximum area is 1. It can be shown that a square with a greater side length can not fit inside any intersecting region of two rectangles.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/07/15/diag.png" />
// Input: bottomLeft = [[1,1],[1,3],[1,5]], topRight = [[5,5],[5,7],[5,9]]
// Output: 4
// Explanation:
// A square with side length 2 can fit inside either the intersecting region of rectangles 0 and 1 or the intersecting region of rectangles 1 and 2. Hence the maximum area is 2 * 2 = 4. It can be shown that a square with a greater side length can not fit inside any intersecting region of two rectangles.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2024/01/04/rectanglesexample2.png" />
// Input: bottomLeft = [[1,1],[2,2],[1,2]], topRight = [[3,3],[4,4],[3,4]]
// Output: 1
// Explanation:
// A square with side length 1 can fit inside the intersecting region of any two rectangles. Also, no larger square can, so the maximum area is 1. Note that the region can be formed by the intersection of more than 2 rectangles.

// Example 4:
// <img src="https://assets.leetcode.com/uploads/2024/01/04/rectanglesexample3.png" />
// Input: bottomLeft = [[1,1],[3,3],[3,1]], topRight = [[2,2],[4,4],[4,2]]
// Output: 0
// Explanation:
// No pair of rectangles intersect, hence, the answer is 0.

// Constraints:
//     n == bottomLeft.length == topRight.length
//     2 <= n <= 10^3
//     bottomLeft[i].length == topRight[i].length == 2
//     1 <= bottomLeft[i][0], bottomLeft[i][1] <= 10^7
//     1 <= topRight[i][0], topRight[i][1] <= 10^7
//     bottomLeft[i][0] < topRight[i][0]
//     bottomLeft[i][1] < topRight[i][1]

import "fmt"

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
    res, n := 0, len(bottomLeft)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            blX, blY := max(bottomLeft[i][0], bottomLeft[j][0]), max(bottomLeft[i][1], bottomLeft[j][1])
            trX, trY := min(topRight[i][0], topRight[j][0]), min(topRight[i][1], topRight[j][1])
            if blX < trX && blY < trY {
                side := min(trX - blX, trY - blY)
                res = max(res, side * side)
            }
        }
    }
    return int64(res)
}

func largestSquareArea1(bottomLeft [][]int, topRight [][]int) int64 {
    res := 0
    for i, b1 := range bottomLeft {
        t1 := topRight[i]
        x1, y1, x2, y2 := b1[0], b1[1], t1[0], t1[1]
        for j := i + 1; j < len(bottomLeft); j++ {
            x3, y3, x4, y4 := bottomLeft[j][0], bottomLeft[j][1], topRight[j][0], topRight[j][1]
            weight, height := min(x2, x4) - max(x1, x3), min(y2, y4) - max(y1, y3)
            side := min(weight, height)
            if side > 0 {
                res = max(res, side * side)
            }
        }
    }
    return int64(res)
}

func largestSquareArea2(bottomLeft [][]int, topRight [][]int) int64 {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    f := func(i, j, k int) int {
        return max(min(topRight[i][k], topRight[j][k])-max(bottomLeft[i][k], bottomLeft[j][k]), 0)
    }
    w, n := 0, len(bottomLeft)
    for i := range n {
        for j := range i {
            w = max(w, min(f(i, j, 0), f(i, j, 1)))
        }
    }
    return int64(w * w)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/01/05/example12.png" />
    // Input: bottomLeft = [[1,1],[2,2],[3,1]], topRight = [[3,3],[4,4],[6,6]]
    // Output: 1
    // Explanation:
    // A square with side length 1 can fit inside either the intersecting region of rectangles 0 and 1 or the intersecting region of rectangles 1 and 2. Hence the maximum area is 1. It can be shown that a square with a greater side length can not fit inside any intersecting region of two rectangles.
    fmt.Println(largestSquareArea([][]int{{1,1},{2,2},{3,1}}, [][]int{{3,3},{4,4},{6,6}})) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/07/15/diag.png" />
    // Input: bottomLeft = [[1,1],[1,3],[1,5]], topRight = [[5,5],[5,7],[5,9]]
    // Output: 4
    // Explanation:
    // A square with side length 2 can fit inside either the intersecting region of rectangles 0 and 1 or the intersecting region of rectangles 1 and 2. Hence the maximum area is 2 * 2 = 4. It can be shown that a square with a greater side length can not fit inside any intersecting region of two rectangles.
    fmt.Println(largestSquareArea([][]int{{1,1},{1,3},{1,5}}, [][]int{{5,5},{5,7},{5,9}})) // 4
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2024/01/04/rectanglesexample2.png" />
    // Input: bottomLeft = [[1,1],[2,2],[1,2]], topRight = [[3,3],[4,4],[3,4]]
    // Output: 1
    // Explanation:
    // A square with side length 1 can fit inside the intersecting region of any two rectangles. Also, no larger square can, so the maximum area is 1. Note that the region can be formed by the intersection of more than 2 rectangles.
    fmt.Println(largestSquareArea([][]int{{1,1},{2,2},{1,2}}, [][]int{{3,3},{4,4},{3,4}})) // 1
    // Example 4:
    // <img src="https://assets.leetcode.com/uploads/2024/01/04/rectanglesexample3.png" />
    // Input: bottomLeft = [[1,1],[3,3],[3,1]], topRight = [[2,2],[4,4],[4,2]]
    // Output: 0
    // Explanation:
    // No pair of rectangles intersect, hence, the answer is 0.
    fmt.Println(largestSquareArea([][]int{{1,1},{3,3},{3,1}}, [][]int{{2,2},{4,4},{4,2}})) // 0

    fmt.Println(largestSquareArea1([][]int{{1,1},{2,2},{3,1}}, [][]int{{3,3},{4,4},{6,6}})) // 1
    fmt.Println(largestSquareArea1([][]int{{1,1},{1,3},{1,5}}, [][]int{{5,5},{5,7},{5,9}})) // 4
    fmt.Println(largestSquareArea1([][]int{{1,1},{2,2},{1,2}}, [][]int{{3,3},{4,4},{3,4}})) // 1
    fmt.Println(largestSquareArea1([][]int{{1,1},{3,3},{3,1}}, [][]int{{2,2},{4,4},{4,2}})) // 0

    fmt.Println(largestSquareArea2([][]int{{1,1},{2,2},{3,1}}, [][]int{{3,3},{4,4},{6,6}})) // 1
    fmt.Println(largestSquareArea2([][]int{{1,1},{1,3},{1,5}}, [][]int{{5,5},{5,7},{5,9}})) // 4
    fmt.Println(largestSquareArea2([][]int{{1,1},{2,2},{1,2}}, [][]int{{3,3},{4,4},{3,4}})) // 1
    fmt.Println(largestSquareArea2([][]int{{1,1},{3,3},{3,1}}, [][]int{{2,2},{4,4},{4,2}})) // 0
}