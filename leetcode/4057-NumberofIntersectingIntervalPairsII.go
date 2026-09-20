package main

// 4057. Number of Intersecting Interval Pairs II
// You are given a 2D integer array intervals of n elements, 
// where intervals[i] = [starti, endi] represents the closed interval from starti to endi.

// Return the number of pairs of indices (i, j) such that 0 <= i < j < n and intervals[i] and intervals[j] intersect.

// Two intervals intersect if they have at least one point in common, including when they only share an endpoint.

// Example 1:
// Input: intervals = [[1,2],[2,3],[3,4]]
// Output: 2
// Explanation:
// There are 2 intersecting interval pairs:
// Intervals [1, 2] and [2, 3] intersect at the point 2.
// Intervals [2, 3] and [3, 4] intersect at the point 3.

// Example 2:
// Input: intervals = [[1,5],[2,4],[3,6]]
// Output: 3
// Explanation:
// There are 3 intersecting interval pairs:
// The intersection of [1, 5] and [2, 4] is [2, 4].
// The intersection of [1, 5] and [3, 6] is [3, 5].
// The intersection of [2, 4] and [3, 6] is [3, 4].

// Example 3:
// Input: intervals = [[1,2],[3,4],[5,6]]
// Output: 0
// Explanation:
// There are no intersecting interval pairs. Hence, the answer is 0.

// Constraints:
//     2 <= n == intervals.length <= 10^5
//     intervals[i] = [starti, endi]
//     0 <= starti <= endi <= 10^9

import "fmt"
import "slices"
import "sort"

func countIntersectingIntervals(intervals [][]int) int64 {
    n := len(intervals) 
    slices.SortFunc(intervals, func(a, b []int) int { // 按照右端点升序排序
        return a[1] - b[1] 
    })
    res := n * (n - 1) / 2
    for _, p := range intervals {
        start := p[0]
        // 设 j 是最小的满足 intervals[j][1] >= start 的下标
        // 那么 [0, j-1] 中的区间右端点都 < start，这有 j 个
        res -= sort.Search(n, func(j int) bool { 
            return intervals[j][1] >= start 
        })
    }
    return int64(res)
}

func countIntersectingIntervals1(intervals [][]int) int64 {
    n := len(intervals)
    starts, ends := make([]int, n), make([]int, n)
    for i, p := range intervals {
        starts[i],ends[i] = p[0], p[1]
    }
    sort.Ints(starts)
    sort.Ints(ends)
    res, j := n * (n - 1) / 2, 0
    // 对于每个左端点 start，右端点 < start 的区间都与之不相交
    for _, start := range starts {
        for j < n && ends[j] < start {
            j++
        }
        // [0, j-1] 的区间与当前区间不相交，这有 j 个
        res -= j
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: intervals = [[1,2],[2,3],[3,4]]
    // Output: 2
    // Explanation:
    // There are 2 intersecting interval pairs:
    // Intervals [1, 2] and [2, 3] intersect at the point 2.
    // Intervals [2, 3] and [3, 4] intersect at the point 3.
    fmt.Println(countIntersectingIntervals([][]int{{1,2},{2,3},{3,4}})) // 2
    // Example 2:
    // Input: intervals = [[1,5],[2,4],[3,6]]
    // Output: 3
    // Explanation:
    // There are 3 intersecting interval pairs:
    // The intersection of [1, 5] and [2, 4] is [2, 4].
    // The intersection of [1, 5] and [3, 6] is [3, 5].
    // The intersection of [2, 4] and [3, 6] is [3, 4].
    fmt.Println(countIntersectingIntervals([][]int{{1,5},{2,4},{3,6}})) // 3
    // Example 3:
    // Input: intervals = [[1,2],[3,4],[5,6]]
    // Output: 0
    // Explanation:
    // There are no intersecting interval pairs. Hence, the answer is 0.
    fmt.Println(countIntersectingIntervals([][]int{{1,2},{3,4},{5,6}})) // 0

    fmt.Println(countIntersectingIntervals1([][]int{{1,2},{2,3},{3,4}})) // 2
    fmt.Println(countIntersectingIntervals1([][]int{{1,5},{2,4},{3,6}})) // 3
    fmt.Println(countIntersectingIntervals1([][]int{{1,2},{3,4},{5,6}})) // 0
}