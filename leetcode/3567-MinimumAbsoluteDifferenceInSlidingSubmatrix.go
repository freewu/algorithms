package main

// 3567. Minimum Absolute Difference in Sliding Submatrix
// You are given an m x n integer matrix grid and an integer k.

// For every contiguous k x k submatrix of grid, 
// compute the minimum absolute difference between any two distinct values within that submatrix.

// Return a 2D array ans of size (m - k + 1) x (n - k + 1), 
// where ans[i][j] is the minimum absolute difference in the submatrix whose top-left corner is (i, j) in grid.

// Note: If all elements in the submatrix have the same value, the answer will be 0.

// A submatrix (x1, y1, x2, y2) is a matrix that is formed by choosing all cells matrix[x][y] where x1 <= x <= x2 and y1 <= y <= y2.

// Example 1:
// Input: grid = [[1,8],[3,-2]], k = 2
// Output: [[2]]
// Explanation:
// There is only one possible k x k submatrix: [[1, 8], [3, -2]].
// Distinct values in the submatrix are [1, 8, 3, -2].
// The minimum absolute difference in the submatrix is |1 - 3| = 2. Thus, the answer is [[2]].

// Example 2:
// Input: grid = [[3,-1]], k = 1
// Output: [[0,0]]
// Explanation:
// Both k x k submatrix has only one distinct element.
// Thus, the answer is [[0, 0]].

// Example 3:
// Input: grid = [[1,-2,3],[2,3,5]], k = 2
// Output: [[1,2]]
// Explanation:
// There are two possible k × k submatrix:
// Starting at (0, 0): [[1, -2], [2, 3]].
// Distinct values in the submatrix are [1, -2, 2, 3].
// The minimum absolute difference in the submatrix is |1 - 2| = 1.
// Starting at (0, 1): [[-2, 3], [3, 5]].
// Distinct values in the submatrix are [-2, 3, 5].
// The minimum absolute difference in the submatrix is |3 - 5| = 2.
// Thus, the answer is [[1, 2]].
 
// Constraints:
//     1 <= m == grid.length <= 30
//     1 <= n == grid[i].length <= 30
//     -10^5 <= grid[i][j] <= 10^5
//     1 <= k <= min(m, n)

import "fmt"
import "sort"

func minAbsDiff(grid [][]int, k int) [][]int {
    n, m := len(grid), len(grid[0])
    // Initialize the result matrix with dimensions (rows-k+1) x (cols-k+1)
    res := make([][]int,  n - k + 1)
    for r, _ := range res {
        res[r] = make([]int, m - k + 1)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    find := func(x, y int) int {
        flattern := make([]int, 0, k * k)
        for i := 0; i < k; i++ {
            flattern = append(flattern, grid[x+i][y:y+k]...) // Extract the k elements from each row and append to the slice
        }
        sort.Ints(flattern) // Sort the flattened array to compute absolute differences efficiently
        res := 1 << 31 // Initialize with the maximum possible integer value
        // Compute the minimum absolute difference between consecutive unique elements
        for i := 1; i < len(flattern); i++ {
            if flattern[i] == flattern[i-1] { continue } // Skip duplicate elements
            res = min(res, abs(flattern[i]-flattern[i-1]))
        }
        if res == 1 << 31 { return 0 } // If no valid difference was found, return 0
        return res
    }
    // Iterate through all possible k x k submatrices
    for i := 0; i + k <= n; i++ {
        for j := 0; j + k <= m; j++ {
            res[i][j] = find(i, j)  // Compute the minimum absolute difference for the current submatrix
        }
    }
    return res
}

func minAbsDiff1(grid [][]int, k int) [][]int {
    n, m := len(grid), len(grid[0])
    arr := make([]int, 0, k * k)
    res := make([][]int, n - k + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    sum := func(grid [][]int, cols, k int, arr []int) int {
        for i := range grid {
            for j := cols; j < cols + k; j++ {
                arr = append(arr, grid[i][j])
            }
        }
        sort.Ints(arr)
        res := 1 << 31
        for i := range arr {
            if i > 0 && arr[i] != arr[i-1] {
                res = min(res, abs(arr[i] - arr[i-1]))
            }
        }
        if res == 1 << 31 { return 0 }
        return res
    }
    for i := 0; i < n - k + 1; i++ {
        res[i] = make([]int, m - k + 1)
        for j := 0; j <  m - k + 1; j++ {
            res[i][j] = sum(grid[i:i + k], j, k, arr)
            arr = arr[:0]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[1,8],[3,-2]], k = 2
    // Output: [[2]]
    // Explanation:
    // There is only one possible k x k submatrix: [[1, 8], [3, -2]].
    // Distinct values in the submatrix are [1, 8, 3, -2].
    // The minimum absolute difference in the submatrix is |1 - 3| = 2. Thus, the answer is [[2]].
    fmt.Println(minAbsDiff([][]int{{1,8},{3,-2}}, 2)) // [[2]]
    // Example 2:
    // Input: grid = [[3,-1]], k = 1
    // Output: [[0,0]]
    // Explanation:
    // Both k x k submatrix has only one distinct element.
    // Thus, the answer is [[0, 0]].
    fmt.Println(minAbsDiff([][]int{{3,-1}}, 1)) // [[0,0]]
    // Example 3:
    // Input: grid = [[1,-2,3],[2,3,5]], k = 2
    // Output: [[1,2]]
    // Explanation:
    // There are two possible k × k submatrix:
    // Starting at (0, 0): [[1, -2], [2, 3]].
    // Distinct values in the submatrix are [1, -2, 2, 3].
    // The minimum absolute difference in the submatrix is |1 - 2| = 1.
    // Starting at (0, 1): [[-2, 3], [3, 5]].
    // Distinct values in the submatrix are [-2, 3, 5].
    // The minimum absolute difference in the submatrix is |3 - 5| = 2.
    // Thus, the answer is [[1, 2]].
    fmt.Println(minAbsDiff([][]int{{1,-2,3},{2,3,5}}, 2)) // [[1,2]]

    fmt.Println(minAbsDiff1([][]int{{1,8},{3,-2}}, 2)) // [[2]]
    fmt.Println(minAbsDiff1([][]int{{3,-1}}, 1)) // [[0,0]]
    fmt.Println(minAbsDiff1([][]int{{1,-2,3},{2,3,5}}, 2)) // [[1,2]]
}