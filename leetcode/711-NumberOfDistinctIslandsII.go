package main

// 711. Number of Distinct Islands II
// You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) 
// You may assume all four edges of the grid are surrounded by water.
// An island is considered to be the same as another if they have the same shape, 
// or have the same shape after rotation (90, 180, or 270 degrees only) or reflection (left/right direction or up/down direction).
// Return the number of distinct islands.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/01/distinctisland2-1-grid.jpg" />
// Input: grid = [[1,1,0,0,0],[1,0,0,0,0],[0,0,0,0,1],[0,0,0,1,1]]
// Output: 1
// Explanation: The two islands are considered the same because if we make a 180 degrees clockwise rotation on the first island, then two islands will have the same shapes.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/05/01/distinctisland1-1-grid.jpg" />
// Input: grid = [[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]
// Output: 1
 
// Constraints:
//         m == grid.length
//         n == grid[i].length
//         1 <= m, n <= 50
//         grid[i][j] is either 0 or 1.

import "fmt"

// 规范化哈希
func numDistinctIslands2(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    seen := make([][]bool, m)
    for i := 0; i < m; i++ {
        seen[i] = make([]bool, n)
        for j := 0; j < n; j++ {
            seen[i][j] = false
        }
    }
    shapes := make(map[string]bool)
    shape := []int{}

    var explore func(r, c int)
    explore = func(r, c int) {
        if 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0]) && grid[r][c] == 1 && !seen[r][c] {
            seen[r][c] = true
            shape = append(shape, r * len(grid[0]) + c)
            explore(r + 1, c)
            explore(r - 1, c)
            explore(r, c + 1)
            explore(r, c - 1)
        }
    }
    var canonical func(shape []int) string
    canonical = func(shape []int) string {
        ans := ""
        lift := len(grid) + len(grid[0])
        out := make([]int, len(shape))
        xs := make([]int, len(shape))
        ys := make([]int, len(shape))

        for c := 0; c < 8; c++ {
            t := 0;
            for _, z := range shape {
                x := z / len(grid[0])
                y := z % len(grid[0])
                if c <= 1 {
                    xs[t] = x
                } else if c <= 3 {
                    xs[t] = -x
                } else if c <= 5 {
                    xs[t] = y
                } else {
                    xs[t] = -y
                }
                if c <= 3 {
                    if c % 2 == 0 {
                        ys[t] = y
                    } else {
                        ys[t] = -y
                    }
                } else {
                    if c % 2 == 0 {
                        ys[t] = x
                    } else {
                        ys[t] = -x
                    }
                }
                t++
            }
            mx, my := xs[0], ys[0]
            for i := 0; i < len(shape); i++ {
                mx = min(mx, xs[i])
                my = min(my, ys[i])
            }
            for i := 0; i < len(shape); i++ {
                out[i] = (xs[i] - mx) * lift + (ys[i] - my)
            }
            sort.Ints(out)
            candidate := ""
            for _, x := range(out) {
                candidate += strconv.Itoa(x) + " "
            }
            if ans < candidate {
                ans = candidate
            }
        }
        return ans
    }
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            shape = []int{}
            explore(r, c)
            if len(shape) > 0 {
                shapes[canonical(shape)] = true
            }
        }
    }
    return len(shapes)
}

func main() {
    fmt.Println(numDistinctIslands2(
        [][]int{
            []int{1,1,0,0,0},
            []int{1,0,0,0,0},
            []int{0,0,0,0,1},
            []int{0,0,0,1,1},
        },
    )) // 1
    fmt.Println(numDistinctIslands2(
        [][]int{
            []int{1,1,0,0,0},
            []int{1,1,0,0,0},
            []int{0,0,0,1,1},
            []int{0,0,0,1,1},
        },
    )) // 1
}