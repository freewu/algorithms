package main

// 3323. Minimize Connected Groups by Inserting Interval
// You are given a 2D array intervals, where intervals[i] = [starti, endi] represents the start and the end of interval i. 
// You are also given an integer k.

// You must add exactly one new interval [startnew, endnew] to the array such that:
//     The length of the new interval, endnew - startnew, is at most k.
//     After adding, the number of connected groups in intervals is minimized.

// A connected group of intervals is a maximal collection of intervals that, 
// when considered together, cover a continuous range from the smallest point to the largest point with no gaps between them. 
// Here are some examples:
//     A group of intervals [[1, 2], [2, 5], [3, 3]] is connected because together they cover the range from 1 to 5 without any gaps.
//     However, a group of intervals [[1, 2], [3, 4]] is not connected because the segment (2, 3) is not covered.

// Return the minimum number of connected groups after adding exactly one new interval to the array.

// Example 1:
// Input: intervals = [[1,3],[5,6],[8,10]], k = 3
// Output: 2
// Explanation:
// After adding the interval [3, 5], we have two connected groups: [[1, 3], [3, 5], [5, 6]] and [[8, 10]].

// Example 2:
// Input: intervals = [[5,10],[1,1],[3,3]], k = 1
// Output: 3
// Explanation:
// After adding the interval [1, 1], we have three connected groups: [[1, 1], [1, 1]], [[3, 3]], and [[5, 10]].

// Constraints:
//     1 <= intervals.length <= 10^5
//     intervals[i] == [starti, endi]
//     1 <= starti <= endi <= 10^9
//     1 <= k <= 10^9

import "fmt"
import "sort"

func minConnectedGroups(intervals [][]int, k int) int {
    sort.Slice(intervals, func(i, j int) bool { 
        return intervals[i][0] < intervals[j][0] 
    })
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    merged := [][]int{}
    for _, interval := range intervals {
        start, end := interval[0], interval[1]
        if len(merged) == 0 || merged[len(merged) - 1][1] < start {
            merged = append(merged, interval)
        } else {
            merged[len(merged)-1][1] = max(merged[len(merged)-1][1], end)
        }
    }
    res := len(merged)
    for i, interval := range merged {
        j := sort.Search(len(merged), func(j int) bool { 
            return merged[j][0] >= interval[1] + k + 1 
        })
        res = min(res, len(merged) - (j-i-1))
    }
    return res
}

func minConnectedGroups1(intervals [][]int, k int) int {
    n := len(intervals)
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, from, to := n, 0, 0
    temp := make([][]int, 0, n)
    for i := 0; i < n; i++ {
        if i == 0 {
            from, to = intervals[i][0], intervals[i][1]
        } else {
            if to >= intervals[i][0] {
                to = max(to, intervals[i][1])
            } else {
                temp = append(temp, []int{from, to})
                from, to = intervals[i][0], intervals[i][1]
            }
        }
    }
    temp = append(temp, []int{ from, to })
    for i := 0; i < len(temp); i++ {
        from, to = temp[i][0], temp[i][1]
        newTo := to + k
        low, high := 0, len(temp)-1
        for low <= high {
            mid := (low + high) / 2
            if temp[mid][0] <= newTo {
                low = mid + 1
            } else {
                high = mid - 1
            }
        }
        low--
        if low >= 0 {
            res = min(res, len(temp)-(low-i+1)+1)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: intervals = [[1,3],[5,6],[8,10]], k = 3
    // Output: 2
    // Explanation:
    // After adding the interval [3, 5], we have two connected groups: [[1, 3], [3, 5], [5, 6]] and [[8, 10]].
    fmt.Println(minConnectedGroups([][]int{{1,3},{5,6},{8,10}}, 3)) // 2
    // Example 2:
    // Input: intervals = [[5,10],[1,1],[3,3]], k = 1
    // Output: 3
    // Explanation:
    // After adding the interval [1, 1], we have three connected groups: [[1, 1], [1, 1]], [[3, 3]], and [[5, 10]].
    fmt.Println(minConnectedGroups([][]int{{5,10},{1,1},{3,3}}, 1)) // 3

    fmt.Println(minConnectedGroups1([][]int{{1,3},{5,6},{8,10}}, 3)) // 2
    fmt.Println(minConnectedGroups1([][]int{{5,10},{1,1},{3,3}}, 1)) // 3
}