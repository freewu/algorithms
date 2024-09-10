package main

// 3197. Find the Minimum Area to Cover All Ones II
// You are given a 2D binary array grid. 
// You need to find 3 non-overlapping rectangles having non-zero areas with horizontal 
// and vertical sides such that all the 1's in grid lie inside these rectangles.

// Return the minimum possible sum of the area of these rectangles.

// Note that the rectangles are allowed to touch.

// Example 1:
// Input: grid = [[1,0,1],[1,1,1]]
// Output: 5
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/14/example0rect21.png" />
// The 1's at (0, 0) and (1, 0) are covered by a rectangle of area 2.
// The 1's at (0, 2) and (1, 2) are covered by a rectangle of area 2.
// The 1 at (1, 1) is covered by a rectangle of area 1.

// Example 2:
// Input: grid = [[1,0,1,0],[0,1,0,1]]
// Output: 5
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/14/example1rect2.png" />
// The 1's at (0, 0) and (0, 2) are covered by a rectangle of area 3.
// The 1 at (1, 1) is covered by a rectangle of area 1.
// The 1 at (1, 3) is covered by a rectangle of area 1.

// Constraints:
//     1 <= grid.length, grid[i].length <= 30
//     grid[i][j] is either 0 or 1.
//     The input is generated such that there are at least three 1's in grid.

import "fmt"

func minimumSum(grid [][]int) int {
    m, n, inf := len(grid), len(grid[0]), 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    getMin := func (a, b int) int {
        if a == -1 || b == -1 { return a + b + 1 }
        return min(a, b)
    }
    getMax := func(a, b int) int {
        if a == -1 || b == -1 { return a + b + 1 }
        return max(a, b)
    }
    calc := func(v []int) int {
        if v[0] == -1 || v[1] == -1 || v[2] == -1 || v[3] == -1 { return inf }
        return (v[1] - v[0] + 1) * (v[3] - v[2] + 1)
    }
    get := func(arr [][][]int, x, y int) (int, int, int, int) {
        if x < 0 || x >= m || y < 0 || y >= n { return -1, -1, -1, -1 }
        return arr[x][y][0], arr[x][y][1], arr[x][y][2], arr[x][y][3]
    }
    makePrefix := func(x1, x2, y1, y2, stepX, stepY int, a, b, c, d int) [][][]int {
        f := make([][][]int, m)
        for i := range f {
            f[i] = make([][]int, n)
        }
        for x := x1; x != x2; x += stepX {
            for y := y1; y != y2; y += stepY {
                l1, r1, u1, d1 := get(f, x+a, y+b)
                l2, r2, u2, d2 := get(f, x+c, y+d)
                vx, vy := x, y
                if grid[x][y] == 0 {
                    vx, vy = -1, -1
                }
                f[x][y] = []int{
                    getMin(getMin(l1, l2), vy),
                    getMax(getMax(r1, r2), vy),
                    getMin(getMin(u1, u2), vx),
                    getMax(getMax(d1, d2), vx),
                }
            }
        }
        return f
    }
    getArea := func(grid [][]int, x1, x2, y1, y2 int) int {
        l, r, u, d := -1, -1, -1, -1
        for i := x1; i <= x2; i++ {
            for j := y1; j <= y2; j++ {
                if grid[i][j] == 1 {
                    l = getMin(l, j)
                    r = getMax(r, j)
                    u = getMin(u, i)
                    d = getMax(d, i)
                }
            }
        }
        return calc([]int{l, r, u, d})
    }
    // dp[i][y] = {矩形最左,矩形最右,矩形最上,矩形最下}
    leftUp := makePrefix(0, m, 0, n, 1, 1, -1, 0, 0, -1)
    rightUp := makePrefix(0, m, n-1, -1, 1, -1, -1, 0, 0, 1)
    leftDown := makePrefix(m-1, -1, 0, n, -1, 1, 1, 0, 0, -1)
    rightDown := makePrefix(m-1, -1, n-1, -1, -1, -1, 1, 0, 0, 1)

    res := m * n
    // 1. x/y 在边上: 横向画两条线
    // [0, x1] [x1 + 1, x2] [x2 + 1, m - 1]
    for x1 := 0; x1 < m; x1++ {
        for x2 := x1 + 1; x2 < m-1; x2++ {
            a := calc(leftUp[x1][n-1])
            b := calc(leftDown[x2+1][n-1])
            c := getArea(grid, x1+1, x2, 0, n-1)
            res = min(res, a+b+c)
        }
    }
    // 2. x/y 在边上: 纵向画两条线
    // [0, y1] [y1 + 1, y2] [y2 + 1, n - 1]
    for y1 := 0; y1 < n; y1++ {
        for y2 := y1 + 1; y2 < n-1; y2++ {
            a := calc(leftUp[m-1][y1])
            b := calc(rightUp[m-1][y2+1])
            c := getArea(grid, 0, m-1, y1+1, y2)
            res = min(res, a+b+c)
        }
    }
    // 3. x/y 都不在边上: 分成四个矩形
    // leftUp[x, y] rightUp[x, y+1] leftDown[x+1, y] rightDown[x+1,y+1]
    for x := 0; x < m-1; x++ {
        for y := 0; y < n-1; y++ {
            a := calc(leftUp[m-1][y]) + calc(rightUp[x][y+1]) + calc(rightDown[x+1][y+1])
            b := calc(leftUp[x][n-1]) + calc(leftDown[x+1][y]) + calc(rightDown[x+1][y+1])
            c := calc(rightDown[0][y+1]) + calc(leftUp[x][y]) + calc(leftDown[x+1][y])
            d := calc(leftDown[x+1][n-1]) + calc(leftUp[x][y]) + calc(rightUp[x][y+1])
            res = getMin(res, getMin(a, getMin(b, getMin(c, d))))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[1,0,1],[1,1,1]]
    // Output: 5
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/14/example0rect21.png" />
    // The 1's at (0, 0) and (1, 0) are covered by a rectangle of area 2.
    // The 1's at (0, 2) and (1, 2) are covered by a rectangle of area 2.
    // The 1 at (1, 1) is covered by a rectangle of area 1.
    fmt.Println(minimumSum([][]int{{1,0,1},{1,1,1}})) // 5
    // Example 2:
    // Input: grid = [[1,0,1,0],[0,1,0,1]]
    // Output: 5
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/14/example1rect2.png" />
    // The 1's at (0, 0) and (0, 2) are covered by a rectangle of area 3.
    // The 1 at (1, 1) is covered by a rectangle of area 1.
    // The 1 at (1, 3) is covered by a rectangle of area 1.
    fmt.Println(minimumSum([][]int{{1,0,1,0},{0,1,0,1}})) // 5
}