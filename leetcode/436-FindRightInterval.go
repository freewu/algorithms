package main

// 436. Find Right Interval
// You are given an array of intervals, where intervals[i] = [starti, endi] and each starti is unique.
// The right interval for an interval i is an interval j such that startj >= endi and startj is minimized. 
// Note that i may equal j.
// Return an array of right interval indices for each interval i. 
// If no right interval exists for interval i, then put -1 at index i.

// Example 1:
// Input: intervals = [[1,2]]
// Output: [-1]
// Explanation: There is only one interval in the collection, so it outputs -1.

// Example 2:
// Input: intervals = [[3,4],[2,3],[1,2]]
// Output: [-1,0,1]
// Explanation: There is no right interval for [3,4].
// The right interval for [2,3] is [3,4] since start0 = 3 is the smallest start that is >= end1 = 3.
// The right interval for [1,2] is [2,3] since start1 = 2 is the smallest start that is >= end2 = 2.

// Example 3:
// Input: intervals = [[1,4],[2,3],[3,4]]
// Output: [-1,2,-1]
// Explanation: There is no right interval for [1,4] and [3,4].
// The right interval for [2,3] is [3,4] since start2 = 3 is the smallest start that is >= end1 = 3.

// Constraints:
//     1 <= intervals.length <= 2 * 10^4
//     intervals[i].length == 2
//     -10^6 <= starti <= endi <= 10^6
//     The start point of each interval is unique.

import "fmt"
import "sort"

func findRightInterval(intervals [][]int) []int {
    indexes, starts := make(map[int]int), make([]int, len(intervals))
    for i, interval := range intervals {
        start := interval[0]
        indexes[start] = i
        starts[i] = start
    }
    sort.Ints(starts)
    search := func (arr []int, target int) int {
        low, high := 0, len(arr) - 1
        for low < high {
            mid := low + ((high - low) >> 1)
            if arr[mid] >= target {
                high = mid
            } else {
                low = mid + 1
            }
        }
        if arr[low] >= target {
            return low
        }
        return -1
    }
    res := make([]int, len(intervals))
    for i, interval := range intervals {
        end := interval[1]
        found := search(starts, end)
        if found == -1 {
            res[i] = -1
        } else {
            res[i] = indexes[starts[found]]
        }
    }
    return res
}

func findRightInterval1(intervals [][]int) []int {
    for i := range intervals {
        intervals[i] = append(intervals[i],i)
    }
    sort.Slice(intervals, func(i,j int) bool { return intervals[i][0] < intervals[j][0]; })
    res := make([]int,len(intervals))
    for _, p := range intervals {
        left, right := 0, len(intervals)
        for left < right {
            mid := (left + right) / 2
            if intervals[mid][0] >= p[1] {
                right = mid
            }else {
                left = mid+1
            }
        }
        if left < len(intervals) {
            res[p[2]] = intervals[left][2]
        }else{
            res[p[2]] = -1
        }
    }
    return res
}

func main() {
    // Explanation: There is only one interval in the collection, so it outputs -1.
    fmt.Println(findRightInterval([][]int{{1,2}})) // [-1]

    // Explanation: There is no right interval for [3,4].
    // The right interval for [2,3] is [3,4] since start0 = 3 is the smallest start that is >= end1 = 3.
    // The right interval for [1,2] is [2,3] since start1 = 2 is the smallest start that is >= end2 = 2.
    fmt.Println(findRightInterval([][]int{{3,4}, {2,3}, {1,2}})) // [-1,0,1]

    // Explanation: There is no right interval for [1,4] and [3,4].
    // The right interval for [2,3] is [3,4] since start2 = 3 is the smallest start that is >= end1 = 3.
    fmt.Println(findRightInterval([][]int{{1,4}, {2,3}, {3,4}})) // [-1,2,-1]

    fmt.Println(findRightInterval1([][]int{{1,2}})) // [-1]
    fmt.Println(findRightInterval1([][]int{{3,4}, {2,3}, {1,2}})) // [-1,0,1]
    fmt.Println(findRightInterval1([][]int{{1,4}, {2,3}, {3,4}})) // [-1,2,-1]

}