package main

// 3468. Find the Number of Copy Arrays
// You are given an array original of length n and a 2D array bounds of length n x 2, where bounds[i] = [ui, vi].

// You need to find the number of possible arrays copy of length n such that:
//     1. (copy[i] - copy[i - 1]) == (original[i] - original[i - 1]) for 1 <= i <= n - 1.
//     2. ui <= copy[i] <= vi for 0 <= i <= n - 1.

// Return the number of such arrays.

// Example 1:
// Input: original = [1,2,3,4], bounds = [[1,2],[2,3],[3,4],[4,5]]
// Output: 2
// Explanation:
// The possible arrays are:
// [1, 2, 3, 4]
// [2, 3, 4, 5]

// Example 2:
// Input: original = [1,2,3,4], bounds = [[1,10],[2,9],[3,8],[4,7]]
// Output: 4
// Explanation:
// The possible arrays are:
// [1, 2, 3, 4]
// [2, 3, 4, 5]
// [3, 4, 5, 6]
// [4, 5, 6, 7]

// Example 3:
// Input: original = [1,2,1,2], bounds = [[1,1],[2,3],[3,3],[2,3]]
// Output: 0
// Explanation:
// No array is possible.

// Constraints:
//     2 <= n == original.length <= 10^5
//     1 <= original[i] <= 10^9
//     bounds.length == n
//     bounds[i].length == 2
//     1 <= bounds[i][0] <= bounds[i][1] <= 10^9

import "fmt"

func countArrays(original []int, bounds [][]int) int {
    left, right, res := bounds[0][0], bounds[0][1], 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(bounds); i++ {
        diff := original[i] - original[i - 1]
        left, right = max(left + diff, bounds[i][0]), min(right + diff, bounds[i][1])
        res = min(res, right - left + 1)
    }
    return max(res, 0)
}

func countArrays1(original []int, bounds [][]int) int {
    mx, mn := -1 << 31, 1 << 31
    for i, v := range bounds {
        diff := original[i] - original[0]
        mx, mn = max(mx, v[0] - diff), min(mn, v[1] - diff)
    }
    return max(mn - mx + 1, 0)
}

func main() {
    // Example 1:
    // Input: original = [1,2,3,4], bounds = [[1,2],[2,3],[3,4],[4,5]]
    // Output: 2
    // Explanation:
    // The possible arrays are:
    // [1, 2, 3, 4]
    // [2, 3, 4, 5]
    fmt.Println(countArrays([]int{1,2,3,4}, [][]int{{1,2},{2,3},{3,4},{4,5}})) // 2
    // Example 2:
    // Input: original = [1,2,3,4], bounds = [[1,10],[2,9],[3,8],[4,7]]
    // Output: 4
    // Explanation:
    // The possible arrays are:
    // [1, 2, 3, 4]
    // [2, 3, 4, 5]
    // [3, 4, 5, 6]
    // [4, 5, 6, 7]
    fmt.Println(countArrays([]int{1,2,3,4}, [][]int{{1,10},{2,9},{3,8},{4,7}})) // 4
    // Example 3:
    // Input: original = [1,2,1,2], bounds = [[1,1],[2,3],[3,3],[2,3]]
    // Output: 0
    // Explanation:
    // No array is possible.
    fmt.Println(countArrays([]int{1,2,1,2}, [][]int{{1,1},{2,3},{3,3},{2,3}})) // 0

    fmt.Println(countArrays1([]int{1,2,3,4}, [][]int{{1,2},{2,3},{3,4},{4,5}})) // 2
    fmt.Println(countArrays1([]int{1,2,3,4}, [][]int{{1,10},{2,9},{3,8},{4,7}})) // 4
    fmt.Println(countArrays1([]int{1,2,1,2}, [][]int{{1,1},{2,3},{3,3},{2,3}})) // 0
}