package main

// 573. Squirrel Simulation
// You are given two integers height and width representing a garden of size height x width.
// You are also given:
//     an array tree where tree = [treer, treec] is the position of the tree in the garden,
//     an array squirrel where squirrel = [squirrelr, squirrelc] is the position of the squirrel in the garden,
//     and an array nuts where nuts[i] = [nutir, nutic] is the position of the ith nut in the garden.

// The squirrel can only take at most one nut at one time and can move in four directions: 
//     up, down, left, and right, to the adjacent cell.

// Return the minimal distance for the squirrel to collect all the nuts and put them under the tree one by one.
// The distance is the number of moves.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/squirrel1-grid.jpg" />
// Input: height = 5, width = 7, tree = [2,2], squirrel = [4,4], nuts = [[3,0], [2,5]]
// Output: 12
// Explanation: The squirrel should go to the nut at [2, 5] first to achieve a minimal distance.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/squirrel2-grid.jpg" />
// Input: height = 1, width = 3, tree = [0,1], squirrel = [0,0], nuts = [[0,2]]
// Output: 3
 
// Constraints:
//     1 <= height, width <= 100
//     tree.length == 2
//     squirrel.length == 2
//     1 <= nuts.length <= 5000
//     nuts[i].length == 2
//     0 <= treer, squirrelr, nutir <= height
//     0 <= treec, squirrelc, nutic <= width

import "fmt"

func minDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, distance := 0, -1 << 32 -1
    for _, nut := range nuts {
        dist := abs(tree[0] - nut[0]) + abs(tree[1] - nut[1])
        res += dist * 2
        distance = max(distance, dist - abs(squirrel[0] - nut[0]) - abs(squirrel[1] - nut[1]))
    }
    return res - distance
}

func minDistance1(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res, sum, m := 1 << 32 -1, 0, make([]int, len(nuts))
    for i, nut := range nuts {
        m[i] = abs(nut[0] - tree[0]) + abs(nut[1] - tree[1])
        sum += m[i] * 2
    }
    for i, nut := range nuts {
        res = min(res, sum-m[i] + abs(nut[0] - squirrel[0]) + abs(nut[1] - squirrel[1]))
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/squirrel1-grid.jpg" />
    // Input: height = 5, width = 7, tree = [2,2], squirrel = [4,4], nuts = [[3,0], [2,5]]
    // Output: 12
    // Explanation: The squirrel should go to the nut at [2, 5] first to achieve a minimal distance.
    fmt.Println(minDistance(5, 7,[]int{2,2},[]int{4,4},[][]int{{3,0},{2,5}})) // 12
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/squirrel2-grid.jpg" />
    // Input: height = 1, width = 3, tree = [0,1], squirrel = [0,0], nuts = [[0,2]]
    // Output: 3
    fmt.Println(minDistance(1, 3, []int{0,1},[]int{0, 0},[][]int{{0,2}})) // 3

    fmt.Println(minDistance1(5, 7,[]int{2,2},[]int{4,4},[][]int{{3,0},{2,5}})) // 12
    fmt.Println(minDistance1(1, 3, []int{0,1},[]int{0, 0},[][]int{{0,2}})) // 3
}