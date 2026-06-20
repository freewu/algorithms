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

// 单调栈
func maxBuilding1(n int, restrictions [][]int) int {
    type Pair struct{ i, h int } // h:被限制的高度
    res, stack := 0, []Pair{{1, 0}}
    sort.Slice(restrictions, func(i, j int) bool { 
        return restrictions[i][0] < restrictions[j][0] 
    })
    for _, p := range restrictions {
        i, h := p[0], p[1]
        for stack[len(stack) - 1].h >= h + i - stack[len(stack) - 1].i { // 栈顶元素的限制的能力弱,不如当前元素有效(当前元素可以反向影响它)
            stack = stack[:len(stack) - 1]
        }
        top := stack[len(stack) - 1]
        if h < i-top.i + top.h { // 比栈顶, 当前元素的限制才有意义
            stack = append(stack, Pair{p[0], p[1]})
        }
    }
    for i, p := range stack[:len(stack) - 1] {
        q := stack[i+1]
        res = max(res, (p.h + q.h + q.i - p.i) / 2)
    }
    if p := stack[len(stack) - 1]; p.i != n {
        res = max(res, p.h + (n - p.i)) //  n 点此时也受限制,从stack[end]限制增  长而来
    }
    return res
}

func maxBuilding2(n int, restrictions [][]int) int {
    m := len(restrictions)
    if m == 0 {
        return n - 1
    }
    sort.Slice(restrictions, func(i, j int) bool { 
        return restrictions[i][0] < restrictions[j][0] 
    })
    // h[i] 表示编号为 id[i] 的建筑的最大高度
    h := make([]int, m)
    h[0] = min(restrictions[0][0] - 1, restrictions[0][1])
    for i := 1; i < m; i++ {
        h[i] = min(h[i - 1]+restrictions[i][0]-restrictions[i - 1][0], restrictions[i][1])
    }
    for i := m - 2; i >= 0; i-- {
        h[i] = min(h[i], h[i+1]+restrictions[i+1][0]-restrictions[i][0])
    }
    res := max((restrictions[0][0] - 1 + h[0]) / 2, h[m-1]+n-restrictions[m - 1][0])
    for i := range m - 1 {
        res = max(res, (restrictions[i + 1][0]-restrictions[i][0] + h[i] + h[i + 1]) / 2)
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

    fmt.Println(maxBuilding1(5,[][]int{{2,1},{4,1}})) // 2
    fmt.Println(maxBuilding1(6,[][]int{})) // 5
    fmt.Println(maxBuilding1(10,[][]int{{5,3},{2,5},{7,4},{10,3}})) // 5

    fmt.Println(maxBuilding2(5,[][]int{{2,1},{4,1}})) // 2
    fmt.Println(maxBuilding2(6,[][]int{})) // 5
    fmt.Println(maxBuilding2(10,[][]int{{5,3},{2,5},{7,4},{10,3}})) // 5
}