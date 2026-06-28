package main

// 3975. Filter Occupied Intervals
// You are given a 2D integer array occupiedIntervals, where occupiedIntervals[i] = [starti, endi] represents a time interval during which you are occupied. 
// Each interval starts at starti and ends at endi, inclusive. 
// These intervals may overlap.

// You are also given two integers freeStart and freeEnd, which define a free time interval from freeStart to freeEnd, inclusive.

// Your task is to merge all occupied intervals that overlap or touch, then remove all integer points in the free interval from the merged occupied intervals.

// Two intervals touch if the second interval starts immediately after the first one ends. 
// For example, [1, 1] and [2, 2] touch and should be merged into [1, 2].

// Return the remaining occupied intervals in sorted order. 
// The returned intervals must be non-overlapping and must contain the minimum number of intervals possible. 
// If there are no remaining occupied points, return an empty list.

// Example 1:
// Input: occupiedIntervals = [[2,6],[4,8],[10,10],[10,12],[14,16]], freeStart = 7, freeEnd = 11
// Output: [[2,6],[12,12],[14,16]]
// Explanation:
// After merging, the occupied intervals are [2, 8], [10, 12], and [14, 16].
// Excluding the free interval [7, 11] results in [2, 6], [12, 12], and [14, 16].

// Example 2:
// Input: occupiedIntervals = [[1,5],[2,3]], freeStart = 3, freeEnd = 8
// Output: [[1,2]]
// Explanation:
// After merging, the occupied interval is [1, 5].
// Excluding the free interval [3, 8] results in [1, 2].

// Constraints:
//     1 <= occupiedIntervals.length <= 5 * 10^4
//     occupiedIntervals[i].length == 2
//     1 <= starti <= endi <= 10^9
//     1 <= freeStart <= freeEnd <= 10^9

import "fmt"
import "slices"

func filterOccupiedIntervals(occupiedIntervals [][]int, freeStart int, freeEnd int) [][]int {
    slices.SortFunc(occupiedIntervals, func(a, b []int) int { 
        return a[0] - b[0] // 按照左端点从小到大排序
    }) 
    res := [][]int{}
    add := func(l, r int) {
        if r < freeStart || l > freeEnd { // 不相交
            res = append(res, []int{l, r})
            return
        }
        if l < freeStart {
            res = append(res, []int{l, freeStart - 1}) // 余留前缀
        }
        if r > freeEnd {
            res = append(res, []int{freeEnd + 1, r}) // 余留后缀
        }
    }
    left, mxRight := occupiedIntervals[0][0], occupiedIntervals[0][1]
    for _, p := range occupiedIntervals[1:] { // 从第二个区间开始
        l, r := p[0], p[1]
        if l-1 > mxRight { // 发现一个新区间
            add(left, mxRight) // 先把旧的加入答案
            left = l // 记录新区间左端点
        }
        mxRight = max(mxRight, r)
    }
    add(left, mxRight)
    return res
}

func filterOccupiedIntervals1(occupiedIntervals [][]int, freeStart int, freeEnd int) [][]int {
    slices.SortFunc(occupiedIntervals, func(a, b []int) int {
        return a[0] - b[0]
    })
    curr := 0
    for i := range occupiedIntervals {
        if occupiedIntervals[i][0] > occupiedIntervals[curr][1]+1 {
            curr++
            occupiedIntervals[curr] = occupiedIntervals[i]
            continue
        }
        occupiedIntervals[curr][1] = max(occupiedIntervals[i][1], occupiedIntervals[curr][1])
    }
    occupiedIntervals = occupiedIntervals[:curr + 1]
    res := make([][]int, 0, len(occupiedIntervals))
    for _, oc := range occupiedIntervals {
        if oc[1] < freeStart || oc[0] > freeEnd {
            res = append(res, oc)
            continue
        }
        if oc[0] >= freeStart && oc[1] <= freeEnd {
            continue
        }
        if oc[0] < freeStart {
            res = append(res, []int{oc[0], freeStart - 1})
        }
        if oc[1] > freeEnd {
            res = append(res, []int{freeEnd + 1, oc[1]})    
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: occupiedIntervals = [[2,6],[4,8],[10,10],[10,12],[14,16]], freeStart = 7, freeEnd = 11
    // Output: [[2,6],[12,12],[14,16]]
    // Explanation:
    // After merging, the occupied intervals are [2, 8], [10, 12], and [14, 16].
    // Excluding the free interval [7, 11] results in [2, 6], [12, 12], and [14, 16].
    fmt.Println(filterOccupiedIntervals([][]int{{2,6},{4,8},{10,10},{10,12},{14,16}}, 7, 11)) // [[2,6],[12,12],[14,16]]
    // Example 2:
    // Input: occupiedIntervals = [[1,5],[2,3]], freeStart = 3, freeEnd = 8
    // Output: [[1,2]]
    // Explanation:
    // After merging, the occupied interval is [1, 5].
    // Excluding the free interval [3, 8] results in [1, 2].   
    fmt.Println(filterOccupiedIntervals([][]int{{1,5},{2,3}}, 3, 8)) // [[1,2]]

    fmt.Println(filterOccupiedIntervals1([][]int{{2,6},{4,8},{10,10},{10,12},{14,16}}, 7, 11)) // [[2,6],[12,12],[14,16]]
    fmt.Println(filterOccupiedIntervals1([][]int{{1,5},{2,3}}, 3, 8)) // [[1,2]]
}