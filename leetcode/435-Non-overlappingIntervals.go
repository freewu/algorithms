package main

// 435. Non-overlapping Intervals
// Given an array of intervals intervals where intervals[i] = [starti, endi], 
// return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

// Example 1:
// Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
// Output: 1
// Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.

// Example 2:
// Input: intervals = [[1,2],[1,2],[1,2]]
// Output: 2
// Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.

// Example 3:
// Input: intervals = [[1,2],[2,3]]
// Output: 0
// Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
 
// Constraints:
//     1 <= intervals.length <= 10^5
//     intervals[i].length == 2
//     -5 * 10^4 <= starti < endi <= 5 * 10^4

import "fmt"
import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func (a, b int) bool {
        return intervals[a][1] < intervals[b][1]
    })
    res, j := 0, 0
    for i := 1; i < len(intervals); i++ {
        if intervals[j][1] <= intervals[i][0] { // 不重叠
            j = i
            continue
        }
        res++ // 如果没有跳出说明 本项是重叠项需要被移除
    }
    return res
}

// 官方解法，直接取最小end
func eraseOverlapIntervals1(intervals [][]int) int {
    sort.Slice(intervals, func (a, b int) bool {
        return intervals[a][1] < intervals[b][1]
    })
    res, end := 0, intervals[0][1]
    for i, size := 1, len(intervals); i < size; i++ {
        if end <= intervals[i][0] {
            end = intervals[i][1]
        } else {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
    // Output: 1
    // Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.
    fmt.Println(eraseOverlapIntervals([][]int{{1,2},{2,3},{3,4},{1,3}})) // 1
    // Example 2:
    // Input: intervals = [[1,2],[1,2],[1,2]]
    // Output: 2
    // Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.
    fmt.Println(eraseOverlapIntervals([][]int{{1,2},{1,2},{1,2}})) // 2
    // Example 3:
    // Input: intervals = [[1,2],[2,3]]
    // Output: 0
    // Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
    fmt.Println(eraseOverlapIntervals([][]int{{1,2},{2,3}})) // 0

    fmt.Println(eraseOverlapIntervals1([][]int{{1,2},{2,3},{3,4},{1,3}})) // 1
    fmt.Println(eraseOverlapIntervals1([][]int{{1,2},{1,2},{1,2}})) // 2
    fmt.Println(eraseOverlapIntervals1([][]int{{1,2},{2,3}})) // 0
}