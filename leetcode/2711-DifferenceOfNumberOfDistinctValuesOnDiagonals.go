package main

// 2711. Difference of Number of Distinct Values on Diagonals
// Given a 2D grid of size m x n, you should find the matrix answer of size m x n.

// The cell answer[r][c] is calculated by looking at the diagonal values of the cell grid[r][c]:
//     1. Let leftAbove[r][c] be the number of distinct values on the diagonal to the left 
//        and above the cell grid[r][c] not including the cell grid[r][c] itself.
//     2. Let rightBelow[r][c] be the number of distinct values on the diagonal to the right 
//        and below the cell grid[r][c], not including the cell grid[r][c] itself.
//     3. Then answer[r][c] = |leftAbove[r][c] - rightBelow[r][c]|.

// A matrix diagonal is a diagonal line of cells starting from some cell in either the topmost row 
// or leftmost column and going in the bottom-right direction until the end of the matrix is reached.
//     For example, in the below diagram the diagonal is highlighted using the cell with indices (2, 3) colored gray:
//         Red-colored cells are left and above the cell.
//         Blue-colored cells are right and below the cell.

// Return the matrix answer.
// <img src="https://assets.leetcode.com/uploads/2024/05/26/diagonal.png">

// Example 1:
// Input: grid = [[1,2,3],[3,1,5],[3,2,1]]
// Output: Output: [[1,1,0],[1,0,1],[0,1,1]]
// Explanation:
// To calculate the answer cells:
// answer  left-above elements         leftAbove       right-below elements       rightBelow      |leftAbove - rightBelow|
// [0][0]  []                              0           [grid[1][1], grid[2][2]]    |{1, 1}| = 1            1
// [0][1]  []                              0           [grid[1][2]]                |{5}| = 1               1
// [0][2]  []                              0           []	                        0	                    0
// [1][0]  []                              0	        [grid[2][1]]                |{2}| = 1               1
// [1][1]  [grid[0][0]]                |{1}| = 1       [grid[2][2]]                |{1}| = 1               0
// [1][2]  [grid[0][1]]                |{2}| = 1       []                          0                       1
// [2][0]  []                              0           []                          0                       0
// [2][1]  [grid[1][0]]                |{3}| = 1	    []                          0                       1
// [2][2]  [grid[0][0], grid[1][1]]    |{1, 1}| = 1    []                          0                       1

// Example 2:
// Input: grid = [[1]]
// Output: Output: [[0]]

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n, grid[i][j] <= 50

import "fmt"
import "math/bits"

func differenceOfDistinctValues(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    res := make([][]int, m)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := range grid {
        res[i] = make([]int, n)
        for j := range grid[i] {
            x, y := i, j
            seen := make(map[int]bool)
            for x > 0 && y > 0 {
                x, y = x-1, y-1
                seen[grid[x][y]] = true
            }
            topLeft := len(seen)
            x, y = i, j
            seen = make(map[int]bool)
            for x + 1 < m && y + 1 < n {
                x, y = x + 1, y + 1
                seen[grid[x][y]] = true
            }
            bottomRight := len(seen)
            res[i][j] = abs(topLeft - bottomRight)
        }
    }
    return res
}

func differenceOfDistinctValues1(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    res, dp1, dp2 := make([][]int, m), make([][]int, m), make([][]int, m) // dp1 TopLeft  | dp2 BottomRight
    for i := range res {
        res[i], dp1[i], dp2[i] = make([]int, n), make([]int, n), make([]int, n)
    }
    for i := 0; i < m; i++ {
        x, y := i + 1, 1
        mp := make(map[int]struct{})
        for x < m && y < n {
            mp[grid[x-1][y-1]] = struct{}{}
            dp1[x][y] = len(mp)
            x++
            y++
        }
    }
    for i := 0; i < n; i++ {
        x, y := 1, i + 1
        mp := make(map[int]struct{})
        for x < m && y < n {
            mp[grid[x-1][y-1]] = struct{}{}
            dp1[x][y] = len(mp)
            x++
            y++
        }
    }
    for i := m - 1; i >=0; i-- {
        x, y := i - 1, n - 2
        mp := make(map[int]struct{})
        for x >= 0 && y >= 0 {
            mp[grid[x+1][y+1]] = struct{}{}
            dp2[x][y] = len(mp)
            x--
            y--
        }
    }
    for i := n - 1; i >=0; i-- {
        x, y := m - 2, i - 1
        mp := make(map[int]struct{})
        for  x >= 0 && y >= 0 {
            mp[grid[x+1][y+1]] = struct{}{}
            dp2[x][y] = len(mp)
            x--
            y--
        }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            res[i][j] = abs(dp1[i][j] - dp2[i][j])
        }
    }
    return res
}

func differenceOfDistinctValues2(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    res := make([][]int, m)
    for i := range res {
        res[i] = make([]int, n)
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 第一排在右上，最后一排在左下
    // 每排从左上到右下
    // 令 k= i-j+n，那么右上角 k=1，左下角 k=m+n-1
    for k := 1; k < m + n; k++ {
        // 核心：计算 j 的最小值和最大值
        mn := max(n - k, 0) // i = 0 的时候，j = n - k，但不能是负数
        mx := min(m + n - 1 - k, n - 1) // i = m - 1 的时候，j=m+n-1-k，但不能超过 n-1
        set := uint(0)
        for j := mn; j <= mx; j++ {
            i := k + j - n
            res[i][j] = bits.OnesCount(set) // set 的大小
            set |= 1 << grid[i][j] // 把 grid[i][j] 加到 set 中
        }
        set = 0
        for j := mx; j >= mn; j-- {
            i := k + j - n
            res[i][j] = abs(res[i][j] - bits.OnesCount(set))
            set |= 1 << grid[i][j]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[1,2,3],[3,1,5],[3,2,1]]
    // Output: [[1,1,0],[1,0,1],[0,1,1]]
    // Explanation:
    // To calculate the answer cells:
    // answer  left-above elements         leftAbove       right-below elements       rightBelow      |leftAbove - rightBelow|
    // [0][0]  []                              0           [grid[1][1], grid[2][2]]    |{1, 1}| = 1            1
    // [0][1]  []                              0           [grid[1][2]]                |{5}| = 1               1
    // [0][2]  []                              0           []	                        0	                    0
    // [1][0]  []                              0	        [grid[2][1]]                |{2}| = 1               1
    // [1][1]  [grid[0][0]]                |{1}| = 1       [grid[2][2]]                |{1}| = 1               0
    // [1][2]  [grid[0][1]]                |{2}| = 1       []                          0                       1
    // [2][0]  []                              0           []                          0                       0
    // [2][1]  [grid[1][0]]                |{3}| = 1	    []                          0                       1
    // [2][2]  [grid[0][0], grid[1][1]]    |{1, 1}| = 1    []                          0                       1
    fmt.Println(differenceOfDistinctValues([][]int{{1,2,3},{3,1,5},{3,2,1}})) // [[1,1,0],[1,0,1],[0,1,1]]
    // Example 2:
    // Input: grid = [[1]]
    // Output: Output: [[0]]
    fmt.Println(differenceOfDistinctValues([][]int{{1}})) // [[0]]

    fmt.Println(differenceOfDistinctValues([][]int{{1,2,3,4,5,6,7,8,9}, {1,2,3,4,5,6,7,8,9}})) // [[1 1 1 1 1 1 1 1 0] [0 1 1 1 1 1 1 1 1]]
    fmt.Println(differenceOfDistinctValues([][]int{{1,2,3,4,5,6,7,8,9}, {9,8,7,6,5,4,3,2,1}})) // [[1 1 1 1 1 1 1 1 0] [0 1 1 1 1 1 1 1 1]]
    fmt.Println(differenceOfDistinctValues([][]int{{9,8,7,6,5,4,3,2,1}, {1,2,3,4,5,6,7,8,9}})) // [[1 1 1 1 1 1 1 1 0] [0 1 1 1 1 1 1 1 1]]
    fmt.Println(differenceOfDistinctValues([][]int{{9,8,7,6,5,4,3,2,1}, {9,8,7,6,5,4,3,2,1}})) // [[1 1 1 1 1 1 1 1 0] [0 1 1 1 1 1 1 1 1]]

    fmt.Println(differenceOfDistinctValues1([][]int{{1,2,3},{3,1,5},{3,2,1}})) // [[1,1,0],[1,0,1],[0,1,1]]
    fmt.Println(differenceOfDistinctValues1([][]int{{1}})) // [[0]]
    fmt.Println(differenceOfDistinctValues1([][]int{{1,2,3,4,5,6,7,8,9}, {1,2,3,4,5,6,7,8,9}})) // [[1 1 1 1 1 1 1 1 0] [0 1 1 1 1 1 1 1 1]]
    fmt.Println(differenceOfDistinctValues1([][]int{{1,2,3,4,5,6,7,8,9}, {9,8,7,6,5,4,3,2,1}})) // [[1 1 1 1 1 1 1 1 0] [0 1 1 1 1 1 1 1 1]]
    fmt.Println(differenceOfDistinctValues1([][]int{{9,8,7,6,5,4,3,2,1}, {1,2,3,4,5,6,7,8,9}})) // [[1 1 1 1 1 1 1 1 0] [0 1 1 1 1 1 1 1 1]]
    fmt.Println(differenceOfDistinctValues1([][]int{{9,8,7,6,5,4,3,2,1}, {9,8,7,6,5,4,3,2,1}})) // [[1 1 1 1 1 1 1 1 0] [0 1 1 1 1 1 1 1 1]]

    fmt.Println(differenceOfDistinctValues2([][]int{{1,2,3},{3,1,5},{3,2,1}})) // [[1,1,0],[1,0,1],[0,1,1]]
    fmt.Println(differenceOfDistinctValues2([][]int{{1}})) // [[0]]
    fmt.Println(differenceOfDistinctValues2([][]int{{1,2,3,4,5,6,7,8,9}, {1,2,3,4,5,6,7,8,9}})) // [[1 1 1 1 1 1 1 1 0] [0 1 1 1 1 1 1 1 1]]
    fmt.Println(differenceOfDistinctValues2([][]int{{1,2,3,4,5,6,7,8,9}, {9,8,7,6,5,4,3,2,1}})) // [[1 1 1 1 1 1 1 1 0] [0 1 1 1 1 1 1 1 1]]
    fmt.Println(differenceOfDistinctValues2([][]int{{9,8,7,6,5,4,3,2,1}, {1,2,3,4,5,6,7,8,9}})) // [[1 1 1 1 1 1 1 1 0] [0 1 1 1 1 1 1 1 1]]
    fmt.Println(differenceOfDistinctValues2([][]int{{9,8,7,6,5,4,3,2,1}, {9,8,7,6,5,4,3,2,1}})) // [[1 1 1 1 1 1 1 1 0] [0 1 1 1 1 1 1 1 1]]

}