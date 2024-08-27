package main

// 2146. K Highest Ranked Items Within a Price Range
// You are given a 0-indexed 2D integer array grid of size m x n that represents a map of the items in a shop. 
// The integers in the grid represent the following:
//     0 represents a wall that you cannot pass through.
//     1 represents an empty cell that you can freely move to and from.
//     All other positive integers represent the price of an item in that cell. You may also freely move to and from these item cells.

// It takes 1 step to travel between adjacent grid cells.

// You are also given integer arrays pricing and start where pricing = [low, high] and start = [row, col] indicates 
// that you start at the position (row, col) and are interested only in items with a price in the range of [low, high] (inclusive). 
// You are further given an integer k.

// You are interested in the positions of the k highest-ranked items whose prices are within the given price range. 
// The rank is determined by the first of these criteria that is different:
//     1. Distance, defined as the length of the shortest path from the start (shorter distance has a higher rank).
//     2. Price (lower price has a higher rank, but it must be in the price range).
//     3. The row number (smaller row number has a higher rank).
//     4. The column number (smaller column number has a higher rank).

// Return the k highest-ranked items within the price range sorted by their rank (highest to lowest). 
// If there are fewer than k reachable items within the price range, return all of them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/16/example1drawio.png" />
// Input: grid = [[1,2,0,1],[1,3,0,1],[0,2,5,1]], pricing = [2,5], start = [0,0], k = 3
// Output: [[0,1],[1,1],[2,1]]
// Explanation: You start at (0,0).
// With a price range of [2,5], we can take items from (0,1), (1,1), (2,1) and (2,2).
// The ranks of these items are:
// - (0,1) with distance 1
// - (1,1) with distance 2
// - (2,1) with distance 3
// - (2,2) with distance 4
// Thus, the 3 highest ranked items in the price range are (0,1), (1,1), and (2,1).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/12/16/example2drawio1.png" />
// Input: grid = [[1,2,0,1],[1,3,3,1],[0,2,5,1]], pricing = [2,3], start = [2,3], k = 2
// Output: [[2,1],[1,2]]
// Explanation: You start at (2,3).
// With a price range of [2,3], we can take items from (0,1), (1,1), (1,2) and (2,1).
// The ranks of these items are:
// - (2,1) with distance 2, price 2
// - (1,2) with distance 2, price 3
// - (1,1) with distance 3
// - (0,1) with distance 4
// Thus, the 2 highest ranked items in the price range are (2,1) and (1,2).

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/12/30/example3.png" />
// Input: grid = [[1,1,1],[0,0,1],[2,3,4]], pricing = [2,3], start = [0,0], k = 3
// Output: [[2,1],[2,0]]
// Explanation: You start at (0,0).
// With a price range of [2,3], we can take items from (2,0) and (2,1). 
// The ranks of these items are: 
// - (2,1) with distance 5
// - (2,0) with distance 6
// Thus, the 2 highest ranked items in the price range are (2,1) and (2,0). 
// Note that k = 3 but there are only 2 reachable items within the price range.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 10^5
//     1 <= m * n <= 10^5
//     0 <= grid[i][j] <= 10^5
//     pricing.length == 2
//     2 <= low <= high <= 10^5
//     start.length == 2
//     0 <= row <= m - 1
//     0 <= col <= n - 1
//     grid[row][col] > 0
//     1 <= k <= m * n

import "fmt"
import "sort"

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
    m, n := len(grid), len(grid[0])
    type Pos struct { i, j int }
    type DistAndPos struct { distance, i, j int }
    seen := make([][]bool, m)
    for i := range seen {
        seen[i] = make([]bool, n)
    }
    // BFS visits cur, adding positions to next
    cur := []Pos{{start[0], start[1]}}
    seen[start[0]][start[1]] = true
    next := []Pos{}

    arr := make([]DistAndPos, 0, k)
    for dist := 0; len(cur) > 0 && len(arr) < k; dist++ {
        next := next[:0] // re-use previous slice (note swap at end of for loop)
        // Visit each cell at the current distance
        for _, cell := range cur {
            if price := grid[cell.i][cell.j]; price >= pricing[0] && price <= pricing[1] {
                arr = append(arr, DistAndPos{ dist, cell.i, cell.j, })
            }
            // Visit adjacent cells
            for _, d := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
                ii, jj  := cell.i + d[0], cell.j + d[1]
                if ii < 0 || ii >= m || jj < 0 || jj >= n || grid[ii][jj] == 0 || seen[ii][jj] {
                    continue
                }
                seen[ii][jj] = true
                next = append(next, Pos{ii, jj})
            }
        }
        cur, next = next, cur
    }
    sort.Slice(arr, func(i, j int) bool {
        p1, p2 := arr[i], arr[j]
        if p1.distance == p2.distance { // Distance
            if grid[p1.i][p1.j] == grid[p2.i][p2.j] { // Price
                if p1.i == p2.i { // Row
                    return p1.j < p2.j // Column
                }
                return p1.i < p2.i
            }
            return grid[p1.i][p1.j] < grid[p2.i][p2.j]
        }
        return p1.distance < p2.distance
    })
    if len(arr) > k { // Trim excess items if necessary
        arr = arr[:k]
    }
    // Convert to result type
    res := make([][]int, len(arr))
    for i := range res {
        res[i] = []int{arr[i].i, arr[i].j}
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/16/example1drawio.png" />
    // Input: grid = [[1,2,0,1],[1,3,0,1],[0,2,5,1]], pricing = [2,5], start = [0,0], k = 3
    // Output: [[0,1],[1,1],[2,1]]
    // Explanation: You start at (0,0).
    // With a price range of [2,5], we can take items from (0,1), (1,1), (2,1) and (2,2).
    // The ranks of these items are:
    // - (0,1) with distance 1
    // - (1,1) with distance 2
    // - (2,1) with distance 3
    // - (2,2) with distance 4
    // Thus, the 3 highest ranked items in the price range are (0,1), (1,1), and (2,1).
    grid1 := [][]int{
        {1,2,0,1},
        {1,3,0,1},
        {0,2,5,1},
    }
    fmt.Println(highestRankedKItems(grid1, []int{2,5}, []int{0,0}, 3)) // [[0,1],[1,1],[2,1]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/12/16/example2drawio1.png" />
    // Input: grid = [[1,2,0,1],[1,3,3,1],[0,2,5,1]], pricing = [2,3], start = [2,3], k = 2
    // Output: [[2,1],[1,2]]
    // Explanation: You start at (2,3).
    // With a price range of [2,3], we can take items from (0,1), (1,1), (1,2) and (2,1).
    // The ranks of these items are:
    // - (2,1) with distance 2, price 2
    // - (1,2) with distance 2, price 3
    // - (1,1) with distance 3
    // - (0,1) with distance 4
    // Thus, the 2 highest ranked items in the price range are (2,1) and (1,2).
    grid2 := [][]int{
        {1,2,0,1},
        {1,3,3,1},
        {0,2,5,1},
    }
    fmt.Println(highestRankedKItems(grid2, []int{2,3}, []int{2,3}, 2)) // [[2,1],[1,2]]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/12/30/example3.png" />
    // Input: grid = [[1,1,1],[0,0,1],[2,3,4]], pricing = [2,3], start = [0,0], k = 3
    // Output: [[2,1],[2,0]]
    // Explanation: You start at (0,0).
    // With a price range of [2,3], we can take items from (2,0) and (2,1). 
    // The ranks of these items are: 
    // - (2,1) with distance 5
    // - (2,0) with distance 6
    // Thus, the 2 highest ranked items in the price range are (2,1) and (2,0). 
    // Note that k = 3 but there are only 2 reachable items within the price range.
    grid3 := [][]int{
        {1,1,1},
        {0,0,1},
        {2,3,4},
    }
    fmt.Println(highestRankedKItems(grid3, []int{2,3}, []int{0,0}, 3)) // [[2,1],[2,0]]
}