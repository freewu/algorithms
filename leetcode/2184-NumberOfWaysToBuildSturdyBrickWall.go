package main

// 2184. Number of Ways to Build Sturdy Brick Wall
// You are given integers height and width which specify the dimensions of a brick wall you are building. 
// You are also given a 0-indexed array of unique integers bricks, 
// where the ith brick has a height of 1 and a width of bricks[i]. 
// You have an infinite supply of each type of brick and bricks may not be rotated.

// Each row in the wall must be exactly width units long. 
// For the wall to be sturdy, adjacent rows in the wall should not join bricks at the same location, except at the ends of the wall.

// Return the number of ways to build a sturdy wall. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/02/20/image-20220220190749-1.png" />
// Input: height = 2, width = 3, bricks = [1,2]
// Output: 2
// Explanation:
// The first two walls in the diagram show the only two ways to build a sturdy brick wall.
// Note that the third wall in the diagram is not sturdy because adjacent rows join bricks 2 units from the left.

// Example 2:
// Input: height = 1, width = 1, bricks = [5]
// Output: 0
// Explanation:
// There are no ways to build a sturdy wall because the only type of brick we have is longer than the width of the wall.

// Constraints:
//     1 <= height <= 100
//     1 <= width <= 10
//     1 <= bricks.length <= 10
//     1 <= bricks[i] <= 10
//     All the values of bricks are unique.

import "fmt"

func buildWall(height int, width int, bricks []int) (ans int) {
    // 铺成宽度为w的状态
    masks_dp := make([][]int, width + 1)
    masks_dp[0] = []int{0}
    for w := 1; w <= width; w++ {
        for _, b := range bricks {
            if w-b >= 0 {
                for _, mask := range masks_dp[w-b] {
                    masks_dp[w] = append(masks_dp[w], mask|(1<<w))
                }
            }
        }
    }
    // 每种铺砖状态对应的与之没有重叠砖缝的铺砖状态
    prev_masks := map[int][]int{}
    for _, mask := range masks_dp[width] {
        prev_masks[mask] = []int{}
        for _, mask1 := range masks_dp[width] {
            if mask1&mask == (1 << width) {
                prev_masks[mask] = append(prev_masks[mask], mask1)
            }
        }
    }
    // 第 i 层 mask 的种数
    dp := make([]map[int]int, height)
    for i := range dp {
        dp[i] = map[int]int{}
    }
    for mask := range prev_masks {
        dp[0][mask] = 1
    }
    res, mod := 0, 1_000_000_007
    for i := 1; i < height; i++ {
        // 每一层 mask 下的种数可以由前一层状态不为 mask 的种数求出
        for mask := range dp[i-1] {
            for _, prev_mask := range prev_masks[mask] {
                dp[i][mask] = (dp[i][mask] + dp[i-1][prev_mask]) % mod
            }
        }
    }
    for _, v := range dp[height-1] {
        res = (res + v) % mod
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/02/20/image-20220220190749-1.png" />
    // Input: height = 2, width = 3, bricks = [1,2]
    // Output: 2
    // Explanation:
    // The first two walls in the diagram show the only two ways to build a sturdy brick wall.
    // Note that the third wall in the diagram is not sturdy because adjacent rows join bricks 2 units from the left.
    fmt.Println(buildWall(2,3,[]int{1,2})) // 2
    // Example 2:
    // Input: height = 1, width = 1, bricks = [5]
    // Output: 0
    // Explanation:
    // There are no ways to build a sturdy wall because the only type of brick we have is longer than the width of the wall.
    fmt.Println(buildWall(1,1,[]int{5})) // 0
}