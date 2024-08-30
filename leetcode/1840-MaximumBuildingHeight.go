package main

// 1840. Maximum Building Height
// You want to build n new buildings in a city. 
// The new buildings will be built in a line and are labeled from 1 to n.

// However, there are city restrictions on the heights of the new buildings:
//     The height of each building must be a non-negative integer.
//     The height of the first building must be 0.
//     The height difference between any two adjacent buildings cannot exceed 1.

// Additionally, there are city restrictions on the maximum height of specific buildings. 
// These restrictions are given as a 2D integer array restrictions where restrictions[i] = [idi, maxHeighti] indicates that building idi must have a height less than or equal to maxHeighti.

// It is guaranteed that each building will appear at most once in restrictions, and building 1 will not be in restrictions.
// Return the maximum possible height of the tallest building.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/08/ic236-q4-ex1-1.png" />
// Input: n = 5, restrictions = [[2,1],[4,1]]
// Output: 2
// Explanation: The green area in the image indicates the maximum allowed height for each building.
// We can build the buildings with heights [0,1,2,1,2], and the tallest building has a height of 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/08/ic236-q4-ex2.png" />
// Input: n = 6, restrictions = []
// Output: 5
// Explanation: The green area in the image indicates the maximum allowed height for each building.
// We can build the buildings with heights [0,1,2,3,4,5], and the tallest building has a height of 5.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/04/08/ic236-q4-ex3.png" />
// Input: n = 10, restrictions = [[5,3],[2,5],[7,4],[10,3]]
// Output: 5
// Explanation: The green area in the image indicates the maximum allowed height for each building.
// We can build the buildings with heights [0,1,2,3,3,4,4,5,4,3], and the tallest building has a height of 5.

// Constraints:
//     2 <= n <= 10^9
//     0 <= restrictions.length <= min(n - 1, 10^5)
//     2 <= idi <= n
//     idi is unique.
//     0 <= maxHeighti <= 10^9

import "fmt"
import "sort"

func maxBuilding(n int, restrictions [][]int) int {
    restrictions = append(restrictions, [][]int{{1, 0}, {n, n - 1}}...)
    sort.Slice(restrictions, func(i, j int) bool { 
        return restrictions[i][0] < restrictions[j][0] 
    })
    res := 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := len(restrictions) - 2; i >= 0; i-- {
        restrictions[i][1] = min(restrictions[i][1], restrictions[i+1][1] + restrictions[i+1][0] - restrictions[i][0])
    }
    for i := 1; i < len(restrictions); i++ {
        restrictions[i][1] = min(restrictions[i][1], restrictions[i-1][1] + restrictions[i][0] - restrictions[i-1][0])
        x := restrictions[i][0] - restrictions[i-1][0] + abs(restrictions[i][1] - restrictions[i-1][1])
        res = max(res, x / 2 + min(restrictions[i-1][1], restrictions[i][1]))
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/08/ic236-q4-ex1-1.png" />
    // Input: n = 5, restrictions = [[2,1],[4,1]]
    // Output: 2
    // Explanation: The green area in the image indicates the maximum allowed height for each building.
    // We can build the buildings with heights [0,1,2,1,2], and the tallest building has a height of 2.
    fmt.Println(maxBuilding(5,[][]int{{2,1},{4,1}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/08/ic236-q4-ex2.png" />
    // Input: n = 6, restrictions = []
    // Output: 5
    // Explanation: The green area in the image indicates the maximum allowed height for each building.
    // We can build the buildings with heights [0,1,2,3,4,5], and the tallest building has a height of 5.
    fmt.Println(maxBuilding(6,[][]int{})) // 5
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/04/08/ic236-q4-ex3.png" />
    // Input: n = 10, restrictions = [[5,3],[2,5],[7,4],[10,3]]
    // Output: 5
    // Explanation: The green area in the image indicates the maximum allowed height for each building.
    // We can build the buildings with heights [0,1,2,3,3,4,4,5,4,3], and the tallest building has a height of 5.
    fmt.Println(maxBuilding(10,[][]int{{5,3},{2,5},{7,4},{10,3}})) // 5
}