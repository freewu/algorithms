package main

// 3219. Minimum Cost for Cutting Cake II
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
//     1 <= m, n <= 10^5
//     horizontalCut.length == m - 1
//     verticalCut.length == n - 1
//     1 <= horizontalCut[i], verticalCut[i] <= 10^3

import "fmt"
import "sort"

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
    res, hindex, vindex, m, n := 0, 0, 0, len(horizontalCut), len(verticalCut) 
    sort.Slice(horizontalCut, func(i, j int) bool { return horizontalCut[i] > horizontalCut[j] })
    sort.Slice(verticalCut, func(i, j int) bool { return verticalCut[i] > verticalCut[j] })
    for hindex < m && vindex < n {
        if horizontalCut[hindex] > verticalCut[vindex] {
            res += horizontalCut[hindex] * (vindex + 1)
            hindex++
        } else {
            res += verticalCut[vindex] * (hindex + 1)
            vindex++
        }
    }
    for hindex < m {
        res += horizontalCut[hindex] * (vindex + 1)
        hindex++
    }
    for vindex < n {
        res += verticalCut[vindex] * (hindex + 1)
        vindex++
    }
    return int64(res)
}

func minimumCost1(m int, n int, horizontalCut []int, verticalCut []int) int64 {
    type ValueTimes struct { Value, Times int }
    countSortAndMerge := func(arr []int) []ValueTimes {
        if len(arr) == 0 { return nil }
        mx, mn := arr[0], arr[0]
        for _, v := range arr { // 找到数组中的最大值和最小值
            if v > mx { mx = v }
            if v < mn { mn = v }
        } 
        count := make([]int, mx - mn + 1)
        for _, v := range arr { // 填充计数数组
            count[v - mn]++
        }
        res := []ValueTimes{}
        for i := len(count) - 1; i >= 0; i-- { // 从大到小遍历计数数组并填充结果数组
            if count[i] > 0 {
                res = append(res, ValueTimes{Value: i + mn, Times: count[i]})
            }
        }
        return res
    }
    h, v := countSortAndMerge(horizontalCut), countSortAndMerge(verticalCut)
    res, hindex, vindex, htimes, vtimes := 0, 0, 0, 1, 1
    for hindex < len(h) || vindex < len(v) {
        times, cost := 0, 0
        if hindex == len(h) {
            times = vtimes * v[vindex].Times
            cost = v[vindex].Value
            htimes += v[vindex].Times
            vindex++
        } else if vindex == len(v) {
            times = htimes * h[hindex].Times
            cost = h[hindex].Value
            vtimes += h[hindex].Times
            hindex++
        } else if v[vindex].Value >= h[hindex].Value {
            times = vtimes * v[vindex].Times
            cost = v[vindex].Value
            htimes += v[vindex].Times
            vindex++
        } else {
            times = htimes * h[hindex].Times
            cost = h[hindex].Value
            vtimes += h[hindex].Times
            hindex++
        }
        res += (times * cost)
    }
    return int64(res)
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

    fmt.Println(minimumCost1(3, 2, []int{1,3}, []int{5})) // 13
    fmt.Println(minimumCost1(2, 2, []int{7}, []int{4})) // 15
}