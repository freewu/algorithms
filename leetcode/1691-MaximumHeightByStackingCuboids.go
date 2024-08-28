package main

// 1691. Maximum Height by Stacking Cuboids
// Given n cuboids where the dimensions of the ith cuboid is cuboids[i] = [widthi, lengthi, heighti] (0-indexed). 
// Choose a subset of cuboids and place them on each other.

// You can place cuboid i on cuboid j if widthi <= widthj and lengthi <= lengthj and heighti <= heightj. 
// You can rearrange any cuboid's dimensions by rotating it to put it on another cuboid.

// Return the maximum height of the stacked cuboids.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/10/21/image.jpg" />
// Input: cuboids = [[50,45,20],[95,37,53],[45,23,12]]
// Output: 190
// Explanation:
// Cuboid 1 is placed on the bottom with the 53x37 side facing down with height 95.
// Cuboid 0 is placed next with the 45x20 side facing down with height 50.
// Cuboid 2 is placed next with the 23x12 side facing down with height 45.
// The total height is 95 + 50 + 45 = 190.

// Example 2:
// Input: cuboids = [[38,25,45],[76,35,3]]
// Output: 76
// Explanation:
// You can't place any of the cuboids on the other.
// We choose cuboid 1 and rotate it so that the 35x3 side is facing down and its height is 76.

// Example 3:
// Input: cuboids = [[7,11,17],[7,17,11],[11,7,17],[11,17,7],[17,7,11],[17,11,7]]
// Output: 102
// Explanation:
// After rearranging the cuboids, you can see that all cuboids have the same dimension.
// You can place the 11x7 side down on all cuboids so their heights are 17.
// The maximum height of stacked cuboids is 6 * 17 = 102.
 
// Constraints:
//     n == cuboids.length
//     1 <= n <= 100
//     1 <= widthi, lengthi, heighti <= 100

import "fmt"
import "sort"

func maxHeight(cuboids [][]int) int {
    for _, c := range cuboids {
        sort.Ints(c) // sort each cuboid's dimensions
    }
    sort.Slice(cuboids, func(i, j int) bool { // sort cuboids by first non-equal dimension
        if cuboids[i][0] != cuboids[j][0] {
            return cuboids[i][0] < cuboids[j][0]
        }
        if cuboids[i][1] != cuboids[j][1] {
            return cuboids[i][1] < cuboids[j][1]
        }
        return cuboids[i][2] < cuboids[j][2]
    })
    // our base case is each cuboid's height
    // e.g. if no cuboids can be stacked together
    // we will just take the tallest of them
    dp := make([]int, len(cuboids))
    for i, c := range cuboids {
        dp[i] = c[2] // 取最大的
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(cuboids); i++ {
        // look over all previous cuboids
        for j := 0; j < i; j++ {
            // if one of the previous cuboids is smaller
            // than the current one
            if cuboids[j][0] <= cuboids[i][0] && cuboids[j][1] <= cuboids[i][1] && cuboids[j][2] <= cuboids[i][2] {
                // check of placing it on top of the current cuboid
                // will be better than previous attempts
                dp[i] = max(dp[i], dp[j] + cuboids[i][2])
            }
        }
    }
    res := -1
    for _, v := range dp {
        if v > res {
            res = v
        }
    }
    return res
}

func maxHeight1(cuboids [][]int) int {
    type Pair struct {  w, l, h int }
    res, n := -1, len(cuboids)
    pairs := make([]Pair, n)
    for i := range cuboids {
        sort.Ints(cuboids[i])
        pairs[i] = Pair{ cuboids[i][0], cuboids[i][1], cuboids[i][2] }
    }
    sort.Slice(pairs, func(i, j int) bool {
        if pairs[i].w == pairs[j].w {
            if pairs[i].l == pairs[j].l {
                return pairs[i].h < pairs[j].h
            }
            return pairs[i].l < pairs[j].l
        }
        return pairs[i].w < pairs[j].w
    })
    dp := make([]int, n)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := range pairs {
        dp[i] = pairs[i].h
        if i == 0 { // 第一块不需要做判断
            continue
        }
        for j := i - 1; j >= 0; j-- {
            if pairs[i].l >= pairs[j].l && pairs[i].h >= pairs[j].h { // 可以叠上去
                dp[i] = max(dp[i], dp[j] + pairs[i].h)
            }
        }
    }
    for i := range dp {
        res = max(res, dp[i])
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/10/21/image.jpg" />
    // Input: cuboids = [[50,45,20],[95,37,53],[45,23,12]]
    // Output: 190
    // Explanation:
    // Cuboid 1 is placed on the bottom with the 53x37 side facing down with height 95.
    // Cuboid 0 is placed next with the 45x20 side facing down with height 50.
    // Cuboid 2 is placed next with the 23x12 side facing down with height 45.
    // The total height is 95 + 50 + 45 = 190.
    fmt.Println(maxHeight([][]int{{50,45,20},{95,37,53},{45,23,12}})) // 190
    // Example 2:
    // Input: cuboids = [[38,25,45],[76,35,3]]
    // Output: 76
    // Explanation:
    // You can't place any of the cuboids on the other.
    // We choose cuboid 1 and rotate it so that the 35x3 side is facing down and its height is 76.
    fmt.Println(maxHeight([][]int{{38,25,45},{76,35,3}})) // 76
    // Example 3:
    // Input: cuboids = [[7,11,17],[7,17,11],[11,7,17],[11,17,7],[17,7,11],[17,11,7]]
    // Output: 102
    // Explanation:
    // After rearranging the cuboids, you can see that all cuboids have the same dimension.
    // You can place the 11x7 side down on all cuboids so their heights are 17.
    // The maximum height of stacked cuboids is 6 * 17 = 102.
    fmt.Println(maxHeight([][]int{{7,11,17},{7,17,11},{11,7,17},{11,17,7},{17,7,11},{17,11,7}})) // 102

    fmt.Println(maxHeight1([][]int{{50,45,20},{95,37,53},{45,23,12}})) // 190
    fmt.Println(maxHeight1([][]int{{38,25,45},{76,35,3}})) // 76
    fmt.Println(maxHeight1([][]int{{7,11,17},{7,17,11},{11,7,17},{11,17,7},{17,7,11},{17,11,7}})) // 102
}