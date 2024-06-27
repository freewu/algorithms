package main

// 675. Cut Off Trees for Golf Event
// You are asked to cut off all the trees in a forest for a golf event. 
// The forest is represented as an m x n matrix. In this matrix:
//     0 means the cell cannot be walked through.
//     1 represents an empty cell that can be walked through.
//     A number greater than 1 represents a tree in a cell that can be walked through, and this number is the tree's height.

// In one step, you can walk in any of the four directions: north, east, south, and west. 
// If you are standing in a cell with a tree, you can choose whether to cut it off.

// You must cut off the trees in order from shortest to tallest. 
// When you cut off a tree, the value at its cell becomes 1 (an empty cell).

// Starting from the point (0, 0), return the minimum steps you need to walk to cut off all the trees. 
// If you cannot cut off all the trees, return -1.

// Note: The input is generated such that no two trees have the same height, and there is at least one tree needs to be cut off.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/26/trees1.jpg" />
// Input: forest = [[1,2,3],[0,0,4],[7,6,5]]
// Output: 6
// Explanation: Following the path above allows you to cut off the trees from shortest to tallest in 6 steps.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/26/trees2.jpg" />
// Input: forest = [[1,2,3],[0,0,0],[7,6,5]]
// Output: -1
// Explanation: The trees in the bottom row cannot be accessed as the middle row is blocked.

// Example 3:
// Input: forest = [[2,3,4],[0,0,5],[8,7,6]]
// Output: 6
// Explanation: You can follow the same path as Example 1 to cut off all the trees.
// Note that you can cut off the first tree at (0, 0) before making any steps.

// Constraints:
//     m == forest.length
//     n == forest[i].length
//     1 <= m, n <= 50
//     0 <= forest[i][j] <= 10^9
//     Heights of all trees are distinct.

import "fmt"
import "sort"

func cutOffTree(forest [][]int) int {
    type Coord struct {
        x   int
        y   int
        val int
    }
    r, c := len(forest), len(forest[0])
    trees := []Coord{}
    for i := range forest {
        for j := range forest[i] {
            if forest[i][j] > 1 {
                trees = append(trees, Coord{i, j, forest[i][j]})
            }
        }
    }
    sort.Slice(trees, func(i, j int) bool {
        return trees[i].val < trees[j].val
    })
    dirs := [][]int{ {1, 0}, {0, 1}, {0, -1}, {-1, 0}, }
    bfs := func(sx, sy, dx, dy int) int {
        row, col := r, c
        visited, rows := make([][]bool, row), make([]bool, row * col)
        for i := range visited {
            visited[i] = rows[i*c : (i+1)*c]
        }
        queue := [][3]int{}
        queue = append(queue, [3]int{sx, sy, 0})
        visited[sx][sy] = true
        for len(queue) > 0 {
            curr := queue[0]
            queue = queue[1:]
            if curr[0] == dx && curr[1] == dy {
                return curr[2]
            }
            for _, d := range dirs {
                nx := d[0] + curr[0]
                ny := d[1] + curr[1]
                if nx >= 0 && nx < r && ny >= 0 && ny < c && !visited[nx][ny] && forest[nx][ny] != 0 {
                    visited[nx][ny] = true
                    queue = append(queue, [3]int{nx, ny, curr[2] + 1})
                }
            }
        }
        return -1
    }
    res, start := 0, Coord{x: 0, y: 0}
    for i := 0; i < len(trees); i++ {
        dist := bfs(start.x, start.y, trees[i].x, trees[i].y)
        if dist == -1 {
            return -1
        }
        start.x = trees[i].x
        start.y = trees[i].y
        res += dist
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/11/26/trees1.jpg" />
    // Input: forest = [[1,2,3],[0,0,4],[7,6,5]]
    // Output: 6
    // Explanation: Following the path above allows you to cut off the trees from shortest to tallest in 6 steps.
    forest1 := [][]int{
        {1,2,3},
        {0,0,4},
        {7,6,5},
    }
    fmt.Println(cutOffTree(forest1)) // 6
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/11/26/trees2.jpg" />
    // Input: forest = [[1,2,3],[0,0,0],[7,6,5]]
    // Output: -1
    // Explanation: The trees in the bottom row cannot be accessed as the middle row is blocked.
    forest2 := [][]int{
        {1,2,3},
        {0,0,0},
        {7,6,5},
    }
    fmt.Println(cutOffTree(forest2)) // -1
    // Example 3:
    // Input: forest = [[2,3,4],[0,0,5],[8,7,6]]
    // Output: 6
    // Explanation: You can follow the same path as Example 1 to cut off all the trees.
    // Note that you can cut off the first tree at (0, 0) before making any steps.
    forest3 := [][]int{
        {1,2,3},
        {0,0,5},
        {8,7,6},
    }
    fmt.Println(cutOffTree(forest3)) // 6
}