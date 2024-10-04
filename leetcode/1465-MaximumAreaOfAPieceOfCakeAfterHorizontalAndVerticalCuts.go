package main

// 1465. Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts
// You are given a rectangular cake of size h x w and two arrays of integers horizontalCuts and verticalCuts where:
//     1. horizontalCuts[i] is the distance from the top of the rectangular cake to the ith horizontal cut and similarly, and
//     2. verticalCuts[j] is the distance from the left of the rectangular cake to the jth vertical cut.

// Return the maximum area of a piece of cake after you cut at each horizontal and vertical position provided in the arrays horizontalCuts and verticalCuts. 
// Since the answer can be a large number, return this modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/05/14/leetcode_max_area_2.png" />
// Input: h = 5, w = 4, horizontalCuts = [1,2,4], verticalCuts = [1,3]
// Output: 4 
// Explanation: The figure above represents the given rectangular cake. Red lines are the horizontal and vertical cuts. After you cut the cake, the green piece of cake has the maximum area.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/05/14/leetcode_max_area_3.png" />
// Input: h = 5, w = 4, horizontalCuts = [3,1], verticalCuts = [1]
// Output: 6
// Explanation: The figure above represents the given rectangular cake. Red lines are the horizontal and vertical cuts. After you cut the cake, the green and yellow pieces of cake have the maximum area.

// Example 3:
// Input: h = 5, w = 4, horizontalCuts = [3], verticalCuts = [3]
// Output: 9

// Constraints:
//     2 <= h, w <= 10^9
//     1 <= horizontalCuts.length <= min(h - 1, 10^5)
//     1 <= verticalCuts.length <= min(w - 1, 10^5)
//     1 <= horizontalCuts[i] < h
//     1 <= verticalCuts[i] < w
//     All the elements in horizontalCuts are distinct.
//     All the elements in verticalCuts are distinct.

import "fmt"
import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
    horizontalCuts, verticalCuts = append(horizontalCuts, []int{0, h}...), append(verticalCuts, []int{0, w}...)
    sort.Ints(horizontalCuts)
    sort.Ints(verticalCuts)
    mxH, mxV := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(horizontalCuts); i++ {
        mxH = max(mxH, horizontalCuts[i] - horizontalCuts[i - 1])
    }
    for i := 1; i < len(verticalCuts); i++ {
        mxV = max(mxV, verticalCuts[i] - verticalCuts[i - 1])
    }
    return (mxH * mxV) % 1_000_000_007
}

func maxArea1(h int, w int, horizontalCuts []int, verticalCuts []int) int {
    sort.Ints(horizontalCuts)
    sort.Ints(verticalCuts)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    calMax := func(arr []int, boardr int) int {
        res, pre := 0, 0
        for _, v := range arr {
            res, pre = max(v - pre, res), v
        }
        return max(res, boardr - pre)
    }
    return calMax(horizontalCuts, h) * calMax(verticalCuts, w) % 1_000_000_007
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/05/14/leetcode_max_area_2.png" />
    // Input: h = 5, w = 4, horizontalCuts = [1,2,4], verticalCuts = [1,3]
    // Output: 4 
    // Explanation: The figure above represents the given rectangular cake. Red lines are the horizontal and vertical cuts. After you cut the cake, the green piece of cake has the maximum area.
    fmt.Println(maxArea(5, 4, []int{1,2,4}, []int{1,3})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/05/14/leetcode_max_area_3.png" />
    // Input: h = 5, w = 4, horizontalCuts = [3,1], verticalCuts = [1]
    // Output: 6
    // Explanation: The figure above represents the given rectangular cake. Red lines are the horizontal and vertical cuts. After you cut the cake, the green and yellow pieces of cake have the maximum area.
    fmt.Println(maxArea(5, 4, []int{3,1}, []int{1})) // 6
    // Example 3:
    // Input: h = 5, w = 4, horizontalCuts = [3], verticalCuts = [3]
    // Output: 9
    fmt.Println(maxArea(5, 4, []int{3}, []int{3})) // 9

    fmt.Println(maxArea1(5, 4, []int{1,2,4}, []int{1,3})) // 4
    fmt.Println(maxArea1(5, 4, []int{3,1}, []int{1})) // 6
    fmt.Println(maxArea1(5, 4, []int{3}, []int{3})) // 9
}