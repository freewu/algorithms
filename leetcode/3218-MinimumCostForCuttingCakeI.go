package main

// 3218. Minimum Cost for Cutting Cake I
// There is an m x n cake that needs to be cut into 1 x 1 pieces.

// You are given integers m, n, and two arrays:
//     horizontalCut of size m - 1, where horizontalCut[i] represents the cost to cut along the horizontal line i.
//     verticalCut of size n - 1, where verticalCut[j] represents the cost to cut along the vertical line j.

// In one operation, you can choose any piece of cake that is not yet a 1 x 1 square and perform one of the following cuts:
//     Cut along a horizontal line i at a cost of horizontalCut[i].
//     Cut along a vertical line j at a cost of verticalCut[j].

// After the cut, the piece of cake is divided into two distinct pieces.

// The cost of a cut depends only on the initial cost of the line and does not change.

// Return the minimum total cost to cut the entire cake into 1 x 1 pieces.

// Example 1:
// Input: m = 3, n = 2, horizontalCut = [1,3], verticalCut = [5]
// Output: 13
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/04/ezgifcom-animated-gif-maker-1.gif" />
// Perform a cut on the vertical line 0 with cost 5, current total cost is 5.
// Perform a cut on the horizontal line 0 on 3 x 1 subgrid with cost 1.
// Perform a cut on the horizontal line 0 on 3 x 1 subgrid with cost 1.
// Perform a cut on the horizontal line 1 on 2 x 1 subgrid with cost 3.
// Perform a cut on the horizontal line 1 on 2 x 1 subgrid with cost 3.
// The total cost is 5 + 1 + 1 + 3 + 3 = 13.

// Example 2:
// Input: m = 2, n = 2, horizontalCut = [7], verticalCut = [4]
// Output: 15
// Explanation:
// Perform a cut on the horizontal line 0 with cost 7.
// Perform a cut on the vertical line 0 on 1 x 2 subgrid with cost 4.
// Perform a cut on the vertical line 0 on 1 x 2 subgrid with cost 4.
// The total cost is 7 + 4 + 4 = 15.

// Constraints:
//     1 <= m, n <= 20
//     horizontalCut.length == m - 1
//     verticalCut.length == n - 1
//     1 <= horizontalCut[i], verticalCut[i] <= 10^3

import "fmt"
import "sort"

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
    sort.Slice(horizontalCut, func(i, j int) bool { return horizontalCut[i] > horizontalCut[j] })
    sort.Slice(verticalCut, func(i, j int) bool { return verticalCut[i] > verticalCut[j] })
    res, hcount, vcount, m, n, inf := 0, 0, 0, len(horizontalCut), len(verticalCut),  -1 << 31
    for i, j := 0, 0; i < m || j < n; {
        hcut, vcut := inf, inf
        if i < m { hcut = horizontalCut[i] }
        if j < n { vcut = verticalCut[j] }
        if hcut > vcut {
            res += (hcut * (vcount + 1))
            hcount++
            i++
        } else {
            res += (vcut * (hcount + 1))
            vcount++
            j++
        }
    }
    return res
}

func minimumCost1(m int, n int, horizontalCut []int, verticalCut []int) int {
    // slices.SortFunc(horizontalCut, func(a, b int) int { return b - a })
    // slices.SortFunc(verticalCut, func(a, b int) int { return b - a })
    sort.Slice(horizontalCut, func(i, j int) bool { return horizontalCut[i] > horizontalCut[j] })
    sort.Slice(verticalCut, func(i, j int) bool { return verticalCut[i] > verticalCut[j] })

    res, i, j, hcount, vcount := 0, 0, 0, 1, 1
    for i < m - 1 || j < n - 1 {
        if j == n - 1 || i < m - 1 && horizontalCut[i] > verticalCut[j] {
            res += horizontalCut[i] * hcount
            i++
            vcount++
        } else {
            res += verticalCut[j] * vcount
            j++
            hcount++
        }
    } 
    return res
}

func main() {
    // Example 1:
    // Input: m = 3, n = 2, horizontalCut = [1,3], verticalCut = [5]
    // Output: 13
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/04/ezgifcom-animated-gif-maker-1.gif" />
    // Perform a cut on the vertical line 0 with cost 5, current total cost is 5.
    // Perform a cut on the horizontal line 0 on 3 x 1 subgrid with cost 1.
    // Perform a cut on the horizontal line 0 on 3 x 1 subgrid with cost 1.
    // Perform a cut on the horizontal line 1 on 2 x 1 subgrid with cost 3.
    // Perform a cut on the horizontal line 1 on 2 x 1 subgrid with cost 3.
    // The total cost is 5 + 1 + 1 + 3 + 3 = 13.
    fmt.Println(minimumCost(3, 2, []int{1,3}, []int{5})) // 13
    // Example 2:
    // Input: m = 2, n = 2, horizontalCut = [7], verticalCut = [4]
    // Output: 15
    // Explanation:
    // Perform a cut on the horizontal line 0 with cost 7.
    // Perform a cut on the vertical line 0 on 1 x 2 subgrid with cost 4.
    // Perform a cut on the vertical line 0 on 1 x 2 subgrid with cost 4.
    // The total cost is 7 + 4 + 4 = 15.
    fmt.Println(minimumCost(2, 2, []int{7}, []int{4})) // 15

    fmt.Println(minimumCost(6, 3, []int{2,3,2,3,1}, []int{1,2})) // 28

    fmt.Println(minimumCost1(3, 2, []int{1,3}, []int{5})) // 13
    fmt.Println(minimumCost1(2, 2, []int{7}, []int{4})) // 15
    fmt.Println(minimumCost1(6, 3, []int{2,3,2,3,1}, []int{1,2})) // 28
}