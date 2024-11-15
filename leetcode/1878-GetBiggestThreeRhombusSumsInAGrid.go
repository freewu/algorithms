package main

// 1878. Get Biggest Three Rhombus Sums in a Grid
// You are given an m x n integer matrix grid​​​.

// A rhombus sum is the sum of the elements that form the border of a regular rhombus shape in grid​​​. 
// The rhombus must have the shape of a square rotated 45 degrees with each of the corners centered in a grid cell. 
// Below is an image of four valid rhombus shapes with the corresponding colored cells that should be included in each rhombus sum:
// <img src="https://assets.leetcode.com/uploads/2021/04/23/pc73-q4-desc-2.png" />

// Note that the rhombus can have an area of 0, which is depicted by the purple rhombus in the bottom right corner.

// Return the biggest three distinct rhombus sums in the grid in descending order. 
// If there are less than three distinct values, return all of them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/23/pc73-q4-ex1.png" />
// Input: grid = [[3,4,5,1,3],[3,3,4,2,3],[20,30,200,40,10],[1,5,5,4,1],[4,3,2,2,5]]
// Output: [228,216,211]
// Explanation: The rhombus shapes for the three biggest distinct rhombus sums are depicted above.
// - Blue: 20 + 3 + 200 + 5 = 228
// - Red: 200 + 2 + 10 + 4 = 216
// - Green: 5 + 200 + 4 + 2 = 211

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/23/pc73-q4-ex1.png" />
// Input: grid = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [20,9,8]
// Explanation: The rhombus shapes for the three biggest distinct rhombus sums are depicted above.
// - Blue: 4 + 2 + 6 + 8 = 20
// - Red: 9 (area 0 rhombus in the bottom right corner)
// - Green: 8 (area 0 rhombus in the bottom middle)

// Example 3:
// Input: grid = [[7,7,7]]
// Output: [7]
// Explanation: All three possible rhombus sums are the same, so return [7].

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 50
//     1 <= grid[i][j] <= 10^5

import "fmt"
import "sort"

func getBiggestThree(grid [][]int) []int {
    m, n := len(grid), len(grid[0])
    v1, v2, v3 := 0, 0, 0
    findRhombusPerimeter := func(grid [][]int,i,j,i1,j1 int) int {
        i2, j2, i3, j3, p := i, j - 2 * (j - j1), i1 + 2 * (i - i1), j1, 0
        if i3 < len(grid) && j2 >= 0 {
            h, k := i1, j1
            for ; h< i && k < j; {
                p += grid[h][k]
                h++
                k++
            }
            h, k = i, j
            for ; h< i3 && k > j3; {
                p += grid[h][k]
                h++
                k--
            }
            h, k = i3, j3
            for ; h> i2 && k > j2; {
                p += grid[h][k]
                h--
                k--
            }
            h, k = i2, j2
            for ; h> i1 && k < j1; {
                p += grid[h][k]
                h--
                k++
            } 
        }
        return p
    }
    for i := 0;i < m; i++ {
        for j := 0;j < n;j++ {
            i1, j1, v := i - 1, j - 1, -1
            for i1 >= 0 && j1 >= 0 && v != 0 {
                v = findRhombusPerimeter(grid, i, j, i1, j1)
                i1--
                j1--
                if v > v1 {
                    v3, v2, v1 = v2, v1, v
                } else if v > v2 && v != v1 {
                    v3, v2 = v2, v
                } else if v > v3 && v != v2 && v != v1 {
                    v3 = v
                }
            }
            v = grid[i][j]
            if v > v1 {
                v3, v2, v1 = v2, v1, v
            } else if v > v2 && v != v1 {
                v3, v2 = v2, v
            } else if v > v3 && v != v2 && v != v1 {
                v3 = v
            }
        }
    }
    res :=[]int{}
    if v1 != 0 { res = append(res, v1) }
    if v2 != 0 { res = append(res, v2) }
    if v3 != 0 { res = append(res, v3) }
    return res
}

func getBiggestThree1(grid [][]int) []int {
    m, n, mp := len(grid), len(grid[0]), make(map[int]bool)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            mp[grid[i][j]] = true
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for l := 2; l <= min((m + 1) / 2, (n + 1) / 2); l++ {
        for i := l - 1; i + l - 1 < m; i++ {
            for j := 0; j + l * 2 - 2 < n; j++ {
                sum :=0 
                for k := 0; k < l; k++ {
                    sum += grid[i-k][j+k] + grid[i+k][j+k] + grid[i-k][j + l * 2 - 2 - k] + grid[i+k][j+l * 2-2-k]
                }
                mp[sum - grid[i][j] - grid[i][j+l*2-2] - grid[i+l-1][j+l-1] - grid[i-l+1][j+l-1]] = true
            }
        }
    }
    arr := []int{}
    for k, _ := range mp {
        arr = append(arr, k)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(arr)))
    return arr[:min(3, len(arr))]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/23/pc73-q4-ex1.png" />
    // Input: grid = [[3,4,5,1,3],[3,3,4,2,3],[20,30,200,40,10],[1,5,5,4,1],[4,3,2,2,5]]
    // Output: [228,216,211]
    // Explanation: The rhombus shapes for the three biggest distinct rhombus sums are depicted above.
    // - Blue: 20 + 3 + 200 + 5 = 228
    // - Red: 200 + 2 + 10 + 4 = 216
    // - Green: 5 + 200 + 4 + 2 = 211
    fmt.Println(getBiggestThree([][]int{{3,4,5,1,3},{3,3,4,2,3},{20,30,200,40,10},{1,5,5,4,1},{4,3,2,2,5}})) // [228,216,211]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/23/pc73-q4-ex1.png" />
    // Input: grid = [[1,2,3],[4,5,6],[7,8,9]]
    // Output: [20,9,8]
    // Explanation: The rhombus shapes for the three biggest distinct rhombus sums are depicted above.
    // - Blue: 4 + 2 + 6 + 8 = 20
    // - Red: 9 (area 0 rhombus in the bottom right corner)
    // - Green: 8 (area 0 rhombus in the bottom middle)
    fmt.Println(getBiggestThree([][]int{{1,2,3},{4,5,6},{7,8,9}})) // [20,9,8]
    // Example 3:
    // Input: grid = [[7,7,7]]
    // Output: [7]
    // Explanation: All three possible rhombus sums are the same, so return [7].
    fmt.Println(getBiggestThree([][]int{{7,7,7}})) // [7]

    fmt.Println(getBiggestThree1([][]int{{3,4,5,1,3},{3,3,4,2,3},{20,30,200,40,10},{1,5,5,4,1},{4,3,2,2,5}})) // [228,216,211]
    fmt.Println(getBiggestThree1([][]int{{1,2,3},{4,5,6},{7,8,9}})) // [20,9,8]
    fmt.Println(getBiggestThree1([][]int{{7,7,7}})) // [7]
}