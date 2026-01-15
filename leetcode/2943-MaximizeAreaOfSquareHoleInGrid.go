package main

// 2943. Maximize Area of Square Hole in Grid
// You are given the two integers, n and m and two integer arrays, hBars and vBars. 
// The grid has n + 2 horizontal and m + 2 vertical bars, creating 1 x 1 unit cells. 
// The bars are indexed starting from 1.

// You can remove some of the bars in hBars from horizontal bars and some of the bars in vBars from vertical bars. 
// Note that other bars are fixed and cannot be removed.

// Return an integer denoting the maximum area of a square-shaped hole in the grid, after removing some bars (possibly none).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/11/05/screenshot-from-2023-11-05-22-40-25.png" />
// Input: n = 2, m = 1, hBars = [2,3], vBars = [2]
// Output: 4
// Explanation:
// The left image shows the initial grid formed by the bars. The horizontal bars are [1,2,3,4], and the vertical bars are [1,2,3].
// One way to get the maximum square-shaped hole is by removing horizontal bar 2 and vertical bar 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/11/04/screenshot-from-2023-11-04-17-01-02.png" />
// Input: n = 1, m = 1, hBars = [2], vBars = [2]
// Output: 4
// Explanation:
// To get the maximum square-shaped hole, we remove horizontal bar 2 and vertical bar 2.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2024/03/12/unsaved-image-2.png" />
// Input: n = 2, m = 3, hBars = [2,3], vBars = [2,4]
// Output: 4
// Explanation:
// One way to get the maximum square-shaped hole is by removing horizontal bar 3, and vertical bar 4.

// Constraints:
//     1 <= n <= 10^9
//     1 <= m <= 10^9
//     1 <= hBars.length <= 100
//     2 <= hBars[i] <= n + 1
//     1 <= vBars.length <= 100
//     2 <= vBars[i] <= m + 1
//     All values in hBars are distinct.
//     All values in vBars are distinct.

import "fmt"
import "sort"

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    find := func(arr []int) int {
        sort.Ints(arr)
        res, i, n := 1, 0, len(arr)
        for i < n {
            count := 1
            for i+1 < n && arr[i]+1 == arr[i+1] {
                i++
                count++
            }
            i++
            res = max(res, count)
        }
        return res
    }
    x, y := find(hBars), find(vBars)
    res := min(x, y) + 1
    return res * res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/11/05/screenshot-from-2023-11-05-22-40-25.png" />
    // Input: n = 2, m = 1, hBars = [2,3], vBars = [2]
    // Output: 4
    // Explanation:
    // The left image shows the initial grid formed by the bars. The horizontal bars are [1,2,3,4], and the vertical bars are [1,2,3].
    // One way to get the maximum square-shaped hole is by removing horizontal bar 2 and vertical bar 2.
    fmt.Println(maximizeSquareHoleArea(2,1,[]int{2,3},[]int{2})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/11/04/screenshot-from-2023-11-04-17-01-02.png" />
    // Input: n = 1, m = 1, hBars = [2], vBars = [2]
    // Output: 4
    // Explanation:
    // To get the maximum square-shaped hole, we remove horizontal bar 2 and vertical bar 2.
    fmt.Println(maximizeSquareHoleArea(1,1,[]int{2},[]int{2})) // 4
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2024/03/12/unsaved-image-2.png" />
    // Input: n = 2, m = 3, hBars = [2,3], vBars = [2,4]
    // Output: 4
    // Explanation:
    // One way to get the maximum square-shaped hole is by removing horizontal bar 3, and vertical bar 4.
    fmt.Println(maximizeSquareHoleArea(2,3,[]int{2,3},[]int{2,4})) // 4

    fmt.Println(maximizeSquareHoleArea(2,3,[]int{1,2,3,4,5,6,7,8,9},[]int{9,8,7,6,5,4,3,2,1})) // 100
    fmt.Println(maximizeSquareHoleArea(2,3,[]int{1,2,3,4,5,6,7,8,9},[]int{1,2,3,4,5,6,7,8,9})) // 100
    fmt.Println(maximizeSquareHoleArea(2,3,[]int{9,8,7,6,5,4,3,2,1},[]int{9,8,7,6,5,4,3,2,1})) // 100
    fmt.Println(maximizeSquareHoleArea(2,3,[]int{9,8,7,6,5,4,3,2,1},[]int{1,2,3,4,5,6,7,8,9})) // 100
}