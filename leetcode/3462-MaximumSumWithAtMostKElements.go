package main

// 3462. Maximum Sum With at Most K Elements
// You are given a 2D integer matrix grid of size n x m, an integer array limits of length n, and an integer 
// The task is to find the maximum sum of at most k elements from the matrix grid such that:

// The number of elements taken from the ith row of grid does not exceed limits[i].

// Return the maximum sum.

// Example 1:
// Input: grid = [[1,2],[3,4]], limits = [1,2], k = 2
// Output: 7
// Explanation:
// From the second row, we can take at most 2 elements. The elements taken are 4 and 3.
// The maximum possible sum of at most 2 selected elements is 4 + 3 = 7.

// Example 2:
// Input: grid = [[5,3,7],[8,2,6]], limits = [2,2], k = 3
// Output: 21
// Explanation:
// From the first row, we can take at most 2 elements. The element taken is 7.
// From the second row, we can take at most 2 elements. The elements taken are 8 and 6.
// The maximum possible sum of at most 3 selected elements is 7 + 8 + 6 = 21.

// Constraints:
//     n == grid.length == limits.length
//     m == grid[i].length
//     1 <= n, m <= 500
//     0 <= grid[i][j] <= 10^5
//     0 <= limits[i] <= m
//     0 <= k <= min(n * m, sum(limits))

import "fmt"
import "sort"

func maxSum(grid [][]int, limits []int, k int) int64 {
    arr := make([]int, 0, len(grid[0]))
    for i, row := range grid {
        sort.Ints(row) // sorts in Ascending order (left-to-right)
        arr = append(arr, row[len(row) - limits[i] :]...) // so have to take max-values from the right-side
    }    
    sort.Ints(arr)
    res := 0
    for _, v := range arr[len(arr) - k :] { 
        res += v
    }
    return int64(res)
}

func maxSum1(grid [][]int, limits []int, k int) int64 {
    m, n := len(grid), len(grid[0])
    arr := make([]int, 0, k)
    for i := 0; i < m; i++ {
        sort.Ints(grid[i])
        arr = append(arr, grid[i][n - limits[i]:]...)
    }
    sort.Ints(arr)
    res := 0
    for ; k > 0; k-- {
        res += arr[len(arr) - k]
    }
    return int64(res)
}

// func maxSum1(grid [][]int, limits []int, k int) int64 {
//     if k == 0 { return 0 }
//     arr := []int{}
//     for i, row := range grid {
//         sort.Ints(row)
//         arr = append(arr, row[len(row) - limits[i]:]...)
//     }
//     sort.Ints(arr)
//     res, count := 0, 0
//     for i := len(arr) - 1; i > 0; i-- {
//         res += arr[i]
//         count++
//         if count == k { break }
//     }
//     // for _, v := range slices.Backward(arr) {
//     //     res += v
//     //     count++
//     //     if count == k { break }
//     // }
//     return int64(res)
// }

func main() {
    // Example 1:
    // Input: grid = [[1,2],[3,4]], limits = [1,2], k = 2
    // Output: 7
    // Explanation:
    // From the second row, we can take at most 2 elements. The elements taken are 4 and 3.
    // The maximum possible sum of at most 2 selected elements is 4 + 3 = 7.
    fmt.Println(maxSum([][]int{{1,2},{3,4}}, []int{1,2}, 2)) // 7
    // Example 2:
    // Input: grid = [[5,3,7],[8,2,6]], limits = [2,2], k = 3
    // Output: 21
    // Explanation:
    // From the first row, we can take at most 2 elements. The element taken is 7.
    // From the second row, we can take at most 2 elements. The elements taken are 8 and 6.
    // The maximum possible sum of at most 3 selected elements is 7 + 8 + 6 = 21.
    fmt.Println(maxSum([][]int{{5,3,7},{8,2,6}}, []int{2,2}, 3)) // 21

    fmt.Println(maxSum([][]int{{5,8,6,1,6,4}}, []int{3}, 3)) // 20

    fmt.Println(maxSum1([][]int{{1,2},{3,4}}, []int{1,2}, 2)) // 7
    fmt.Println(maxSum1([][]int{{5,3,7},{8,2,6}}, []int{2,2}, 3)) // 21
    fmt.Println(maxSum1([][]int{{5,8,6,1,6,4}}, []int{3}, 3)) // 20
}