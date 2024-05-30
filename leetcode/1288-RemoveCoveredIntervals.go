package main

// 1288. Remove Covered Intervals
// Given an array intervals where intervals[i] = [li, ri] represent the interval [li, ri), 
// remove all intervals that are covered by another interval in the list.

// The interval [a, b) is covered by the interval [c, d) if and only if c <= a and b <= d.
// Return the number of remaining intervals.

// Example 1:
// Input: intervals = [[1,4],[3,6],[2,8]]
// Output: 2
// Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.

// Example 2:
// Input: intervals = [[1,4],[2,3]]
// Output: 1
 
// Constraints:
//     1 <= intervals.length <= 1000
//     intervals[i].length == 2
//     0 <= li < ri <= 10^5
//     All the given intervals are unique.

import "fmt"
import "sort"

func removeCoveredIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool { // sort by increasing 'start' then by decreasing 'end'
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] <= intervals[j][0]
    })
    i, count := 0, 0
    for j := 1; j < len(intervals); j += 1 {
        if intervals[j][1] > intervals[i][1] {
            i = j
            continue
        }
        count += 1 // remove jth interval 
    }
    return len(intervals) - count
}

func removeCoveredIntervals1(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })
    res, left, right := 0, intervals[0][0], intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        if left <= intervals[i][0] && right >= intervals[i][1] { // 被覆盖
            res++
        } else if right >= intervals[i][0] && right <= intervals[i][1] { // 有交集
            right = intervals[i][1]
        } else { // 无交集
            left = intervals[i][0]
            right = intervals[i][1]
        }
    }
    return len(intervals) - res
}

func main() {
    // Example 1:
    // Input: intervals = [[1,4],[3,6],[2,8]]
    // Output: 2
    // Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.
    fmt.Println(removeCoveredIntervals([][]int{{1,4},{3,6},{2,8}})) // 2
    // Example 2:
    // Input: intervals = [[1,4],[2,3]]
    // Output: 1
    fmt.Println(removeCoveredIntervals([][]int{{1,4},{2,3}})) // 1

    fmt.Println(removeCoveredIntervals1([][]int{{1,4},{3,6},{2,8}})) // 2
    fmt.Println(removeCoveredIntervals1([][]int{{1,4},{2,3}})) // 1
}